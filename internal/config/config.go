package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Shaper        ShaperSettings        `json:"shaper"`
	Control       ControlSettings       `json:"control"`
	Probe         ProbeSettings         `json:"probe"`
	Priority      PrioritySettings      `json:"priority"`
	Observability ObservabilitySettings `json:"observability"`
}

type ShaperSettings struct {
	WANInterface        string  `json:"wan_interface"`
	AutoDiscoverWAN     bool    `json:"auto_discover_wan"`
	UploadFloorMbps     float64 `json:"upload_floor_mbps"`
	UploadCeilingMbps   float64 `json:"upload_ceiling_mbps"`
	DownloadFloorMbps   float64 `json:"download_floor_mbps"`
	DownloadCeilingMbps float64 `json:"download_ceiling_mbps"`
	CAKEDiffserv        string  `json:"cake_diffserv"`
	CAKEIsolation       string  `json:"cake_isolation"`
}

type ControlSettings struct {
	LoopIntervalMS  int     `json:"loop_interval_ms"`
	TargetLatencyMS float64 `json:"target_latency_ms"`
	KP              float64 `json:"kp"`
	KI              float64 `json:"ki"`
	KD              float64 `json:"kd"`
	IntegralMin     float64 `json:"integral_min"`
	IntegralMax     float64 `json:"integral_max"`
	MaxRateDeltaMbps float64 `json:"max_rate_delta_mbps"`
}

type ProbeSettings struct {
	Reflectors       []Reflector `json:"reflectors"`
	Protocols        []string    `json:"protocols"`
	TimeoutMS        int         `json:"timeout_ms"`
	FastEWMAAlpha    float64     `json:"fast_ewma_alpha"`
	SlowEWMAAlpha    float64     `json:"slow_ewma_alpha"`
	OutlierThreshold float64     `json:"outlier_threshold"`
}

type Reflector struct {
	ID      string `json:"id"`
	Host    string `json:"host"`
	Kind    string `json:"kind"`
	Enabled bool   `json:"enabled"`
}

type PrioritySettings struct {
	Enabled   bool   `json:"enabled"`
	DeviceIP  string `json:"device_ip"`
	DeviceMAC string `json:"device_mac"`
	TargetTin string `json:"target_tin"`
}

type ObservabilitySettings struct {
	MetricsListen string `json:"metrics_listen"`
	LogLevel      string `json:"log_level"`
	TUIEnabled    bool   `json:"tui_enabled"`
}

func LoadFile(path string) (Config, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("AURA_CONFIG_READ: cannot read config %q: %w", path, err)
	}

	var cfg Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return Config{}, fmt.Errorf("AURA_CONFIG_PARSE: cannot parse config %q: %w", path, err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg Config) Validate() error {
	if err := cfg.Shaper.Validate(); err != nil {
		return err
	}
	if err := cfg.Control.Validate(); err != nil {
		return err
	}
	if err := cfg.Probe.Validate(); err != nil {
		return err
	}
	if err := cfg.Priority.Validate(); err != nil {
		return err
	}
	if err := cfg.Observability.Validate(); err != nil {
		return err
	}
	return nil
}

func (s ShaperSettings) Validate() error {
	if !s.AutoDiscoverWAN && s.WANInterface == "" {
		return errors.New("AURA_CONFIG_WAN_INTERFACE: wan_interface is required when auto_discover_wan is false")
	}
	if s.UploadFloorMbps <= 0 || s.UploadCeilingMbps <= 0 {
		return errors.New("AURA_CONFIG_UPLOAD_RANGE: upload floor and ceiling must be greater than zero")
	}
	if s.DownloadFloorMbps <= 0 || s.DownloadCeilingMbps <= 0 {
		return errors.New("AURA_CONFIG_DOWNLOAD_RANGE: download floor and ceiling must be greater than zero")
	}
	if s.UploadFloorMbps >= s.UploadCeilingMbps {
		return errors.New("AURA_CONFIG_UPLOAD_RANGE: upload floor must be lower than upload ceiling")
	}
	if s.DownloadFloorMbps >= s.DownloadCeilingMbps {
		return errors.New("AURA_CONFIG_DOWNLOAD_RANGE: download floor must be lower than download ceiling")
	}
	if s.CAKEDiffserv != "diffserv4" {
		return errors.New("AURA_CONFIG_CAKE_DIFFSERV: first build supports only diffserv4")
	}
	if s.CAKEIsolation != "triple-isolate" {
		return errors.New("AURA_CONFIG_CAKE_ISOLATION: first build supports only triple-isolate")
	}
	return nil
}

func (c ControlSettings) Validate() error {
	if c.LoopIntervalMS < 20 || c.LoopIntervalMS > 50 {
		return errors.New("AURA_CONFIG_LOOP_INTERVAL: loop_interval_ms must be between 20 and 50")
	}
	if c.TargetLatencyMS <= 0 {
		return errors.New("AURA_CONFIG_TARGET_LATENCY: target_latency_ms must be greater than zero")
	}
	if c.IntegralMin > c.IntegralMax {
		return errors.New("AURA_CONFIG_INTEGRAL_CLAMP: integral_min must be lower than integral_max")
	}
	if c.MaxRateDeltaMbps <= 0 {
		return errors.New("AURA_CONFIG_RATE_DELTA: max_rate_delta_mbps must be greater than zero")
	}
	return nil
}

func (p ProbeSettings) Validate() error {
	if len(p.Reflectors) == 0 {
		return errors.New("AURA_CONFIG_REFLECTORS: at least one reflector is required")
	}
	if len(p.Protocols) == 0 {
		return errors.New("AURA_CONFIG_PROTOCOLS: at least one probe protocol is required")
	}
	if p.TimeoutMS <= 0 {
		return errors.New("AURA_CONFIG_PROBE_TIMEOUT: timeout_ms must be greater than zero")
	}
	if p.FastEWMAAlpha <= 0 || p.FastEWMAAlpha >= 1 {
		return errors.New("AURA_CONFIG_FAST_EWMA: fast_ewma_alpha must be greater than 0 and lower than 1")
	}
	if p.SlowEWMAAlpha <= 0 || p.SlowEWMAAlpha >= 1 {
		return errors.New("AURA_CONFIG_SLOW_EWMA: slow_ewma_alpha must be greater than 0 and lower than 1")
	}
	if p.SlowEWMAAlpha >= p.FastEWMAAlpha {
		return errors.New("AURA_CONFIG_EWMA_ORDER: slow_ewma_alpha must be lower than fast_ewma_alpha")
	}
	for _, reflector := range p.Reflectors {
		if reflector.Enabled && (reflector.ID == "" || reflector.Host == "") {
			return errors.New("AURA_CONFIG_REFLECTOR: enabled reflectors require id and host")
		}
	}
	return nil
}

func (p PrioritySettings) Validate() error {
	if !p.Enabled {
		return nil
	}
	if p.DeviceIP == "" && p.DeviceMAC == "" {
		return errors.New("AURA_CONFIG_PRIORITY_DEVICE: device_ip or device_mac is required when priority mode is enabled")
	}
	if p.TargetTin == "" {
		return errors.New("AURA_CONFIG_PRIORITY_TIN: target_tin is required when priority mode is enabled")
	}
	return nil
}

func (o ObservabilitySettings) Validate() error {
	if o.MetricsListen == "" {
		return errors.New("AURA_CONFIG_METRICS_LISTEN: metrics_listen is required")
	}
	switch o.LogLevel {
	case "debug", "info", "warn", "error":
		return nil
	default:
		return errors.New("AURA_CONFIG_LOG_LEVEL: log_level must be debug, info, warn, or error")
	}
}
