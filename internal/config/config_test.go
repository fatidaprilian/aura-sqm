package config

import "testing"

func TestValidateAcceptsMinimalConfig(t *testing.T) {
	cfg := validConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected valid config, got %v", err)
	}
}

func TestValidateRejectsUnsafeUploadRange(t *testing.T) {
	cfg := validConfig()
	cfg.Shaper.UploadFloorMbps = cfg.Shaper.UploadCeilingMbps

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected upload range error")
	}
}

func TestValidateRequiresPriorityIdentity(t *testing.T) {
	cfg := validConfig()
	cfg.Priority.Enabled = true
	cfg.Priority.DeviceIP = ""
	cfg.Priority.DeviceMAC = ""

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected priority identity error")
	}
}

func validConfig() Config {
	return Config{
		Shaper: ShaperSettings{
			AutoDiscoverWAN:     true,
			UploadFloorMbps:     50,
			UploadCeilingMbps:   100,
			DownloadFloorMbps:   50,
			DownloadCeilingMbps: 100,
			CAKEDiffserv:        "diffserv4",
			CAKEIsolation:       "triple-isolate",
		},
		Control: ControlSettings{
			LoopIntervalMS:  50,
			TargetLatencyMS: 8,
			KP:              0.2,
			KI:              0.02,
			KD:              0.01,
			IntegralMin:     -1,
			IntegralMax:     1,
			MaxRateDeltaMbps: 5,
		},
		Probe: ProbeSettings{
			Reflectors: []Reflector{
				{ID: "cloudflare", Host: "1.1.1.1", Kind: "public", Enabled: true},
			},
			Protocols:        []string{"icmp", "udp"},
			TimeoutMS:        500,
			FastEWMAAlpha:    0.3,
			SlowEWMAAlpha:    0.1,
			OutlierThreshold: 2,
		},
		Priority: PrioritySettings{
			Enabled:   true,
			DeviceIP:  "192.168.10.50",
			TargetTin: "voice",
		},
		Observability: ObservabilitySettings{
			MetricsListen: "127.0.0.1:9108",
			LogLevel:      "info",
			TUIEnabled:    true,
		},
	}
}
