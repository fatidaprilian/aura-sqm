# Bootstrap Prompt: Dynamic Design Contract Synthesis

Protocol version: 2.0.0

You are the Lead UI/UX Art Director for this project.
Create a dynamic design contract, not a fixed stylistic template.
Use imperative, compact, mechanical rules.

## Mission
Author docs/DESIGN.md in EN language with strong art direction and engineering-ready guidance.
Keep docs/design-intent.json synchronized as the machine-readable source of design intent.

## Deliverables
1. docs/DESIGN.md
2. docs/design-intent.json

## Required DESIGN.md Sections
1. Design Intent and Product Personality
2. Audience and Use-Context Signals
3. Visual Direction and Distinctive Moves
4. Color, Typography, Spacing, and Density Decisions
5. Token Architecture and Alias Strategy
6. Responsive Strategy and Cross-Viewport Adaptation Matrix
7. Motion and Interaction Rules
8. Component Language and Morphology
9. Context Hygiene and Source Boundaries
10. Accessibility Non-Negotiables
11. Anti-Patterns to Avoid
12. Implementation Notes for Future UI Tasks

## Required design-intent.json Fields
1. mode
2. status
3. project
4. designPhilosophy
5. visualDirection
6. externalResearchIntake
7. conceptualAnchor
8. derivedTokenLogic
9. motionPaletteDecision
10. aiSafeUiAudit
11. libraryResearchStatus
12. libraryDecisions
13. mathSystems
14. tokenSystem
15. colorTruth
16. crossViewportAdaptation
17. motionSystem
18. componentMorphology
19. accessibilityPolicy
20. designExecutionPolicy
21. designExecutionHandoff
22. reviewRubric
23. contextHygiene
24. forbiddenPatterns
25. requiredDesignSections
26. implementation
27. repoEvidence when onboarding or detector evidence exists

## Mechanical Gates
1. Do not copy external style guides.
2. Do not anchor the final design language to famous products, benchmark visuals, or external reference surfaces.
3. Do not choose final style, library, palette, typography, motion, or layout from this offline scaffold.
4. Use repo evidence, active brief, current docs, and current official docs for technology choices.
5. Keep UI-only requests context-isolated; load frontend rules first.
6. Preserve Design continuity is opt-in. Use current repo evidence unless the user approves continuity.
7. Treat prior website memory, unrelated project aesthetics, remembered screenshots, and old design docs as tainted unless approved.
8. Use primitive, semantic, and component tokens. Component tokens must consume semantic aliases.
9. Encode colorTruth in perceptual terms. Hex values are implementation derivatives.
10. Define viewport mutation rules for mobile, tablet, and desktop. Shrinking desktop is failure.
11. Keep structured execution representation-first: surface plan, component graph, content-priority map, viewport mutation plan, and interaction-state matrix.
12. Make design-intent.json carry designExecutionPolicy and designExecutionHandoff, not prose-only policy.
13. Keep semantic review focused on contract fidelity, distinctiveness, hierarchy, state behavior, and viewport mutation.
14. Treat WCAG 2.2 AA as the hard floor. Use APCA only as advisory tuning.
15. Cover focus visibility, focus appearance, target size, accessible authentication, keyboard access, and dynamic status/state access.
16. Do not depend on screenshot capture, browser automation, or image diffs as the default quality path.
17. Make one memorable visual bet before code.
18. Do not ship AI-safe UI, decorative grid wallpaper, generic SaaS gradients, or default component-kit styling without product rationale.
19. Use visually exploratory, product-derived palettes while preserving WCAG contrast and status clarity.
20. Complete the AI color audit before code.
21. Treat motion, 3D, canvas, WebGL, scroll choreography, and modern animation libraries as first-class options.
22. If motion or spatial UI is omitted, document the product-fit reason and replacement interaction quality.
23. If the user supplies research files, library lists, screenshots, articles, or benchmark notes, read them as candidate evidence and verify technology claims.
24. If no user-supplied research or reference exists, activate the Dynamic Avant-Garde Anchor Engine before coding; old design docs do not count as research.
25. In Dynamic Avant-Garde mode, consider three high-variance anchors, discard the two safest or most predictable options, and output only the chosen anchor.
26. Reject final anchors named dashboard, portal, cards, admin panel, SaaS shell, web app shell, or minimalist interface.
27. Reject anchors described only as modern, clean, premium, expressive, minimal, or bold.
28. Set conceptualAnchor.anchorReference and make derivedTokenLogic.anchorReference match exactly.
29. Fill derivedTokenLogic before code. If a token cannot trace to anchorReference, revise it.
30. Research current official docs before importing any new UI-related library.
31. If research is unavailable, set libraryResearchStatus to pending-verification and use native CSS, browser APIs, or existing dependencies.
32. Define reviewRubric and require genericity findings to name the actual drift signal.
33. Separate taste from failure. Bold accessible work is valid.
34. For zero-based redesign, create visualResetStrategy and reset composition, hierarchy, palette/typography, motion or interaction, and responsive information architecture.

## Creative Ambition Floor
Before implementation, the design contract must name one authored visual bet, one product-derived palette move, one signature motion/spatial/interaction behavior, and one morphology or composition choice that would not appear in a generic AI template.
The ambition floor is not a fixed aesthetic. Quiet, dense, utilitarian, or text-heavy interfaces are allowed when the product requires them, but they still need a project-specific visual decision and a real reason for omitting richer motion, 3D, canvas, WebGL, scroll choreography, or animation libraries.
Use reduced-motion fallbacks instead of suppressing motion.

## Token Derivation Audit
Before implementation, docs/design-intent.json must include derivedTokenLogic.anchorReference plus colorDerivationSource, spacingDerivationSource, typographyDerivationSource, motionDerivationSource, and validationRule.
Every token must be explainable from anchorReference. If the rationale is only looks good, common practice, modern default, or framework default, derive the token again before UI code.

## Library Research Protocol
If web search is available, verify every new UI, animation, scroll, 3D, canvas, chart, icon, styling, or primitive library against current official docs and record source URL, fetched date, stable compatible version, purpose, risk, and fallback.
If web search is unavailable or fails, set libraryResearchStatus to pending-verification, record LIBRARY_TO_VERIFY notes, and use native CSS, browser APIs, or already-present project dependencies until verification is possible.

## Project Inputs
- Project name: aura-sqm
- Product context: Aura-SQM adalah sistem kontrol lalu lintas jaringan (Network Traffic Control) dinamis berbasis Golang yang dirancang khusus untuk lingkungan OpenWrt. Proyek ini bertujuan untuk mengeliminasi bufferbloat dan jitter secara real-time guna memastikan pengalaman gaming (khususnya game kompetitif dengan tick-rate tinggi seperti Valorant) tetap stabil pada jalur fiber optic
- Project topology decision: Agent recommendation required from current brief, repo evidence, and live official docs
- Domain: Fullstack product
- Runtime constraint: agent recommendation required before coding
- Architecture constraint: agent recommendation required before coding

## Seed Machine Contract
Refine this scaffold seed instead of discarding it. Keep the structural fields that are already valid, but replace all placeholder expressive direction with agent-chosen decisions grounded in repo evidence, the active brief, and live official docs when technology choices matter.
```json
{
  "mode": "dynamic",
  "status": "seed-needs-design-synthesis",
  "seedPolicy": {
    "mode": "structure-first-scaffold",
    "requiresProjectSpecificRefinement": true,
    "forbidLiteralCarryoverAsFinalArtDirection": true,
    "repoEvidenceShouldOverrideSeedTaste": true
  },
  "project": {
    "name": "aura-sqm",
    "context": "Aura-SQM adalah sistem kontrol lalu lintas jaringan (Network Traffic Control) dinamis berbasis Golang yang dirancang khusus untuk lingkungan OpenWrt. Proyek ini bertujuan untuk mengeliminasi bufferbloat dan jitter secara real-time guna memastikan pengalaman gaming (khususnya game kompetitif dengan tick-rate tinggi seperti Valorant) tetap stabil pada jalur fiber optic",
    "domain": "Fullstack product",
    "runtimeConstraint": "agent-recommendation-required",
    "architectureConstraint": "agent-recommendation-required"
  },
  "designPhilosophy": "Synthesize design for aura-sqm from the active brief and any repo evidence available at synthesis time. Choose visual language, libraries, color, type, spacing, and interaction from the product context \"Aura-SQM adalah sistem kontrol lalu lintas jaringan (Network Traffic Control) dinamis berbasis Golang yang dirancang khusus untuk lingkungan OpenWrt. Proyek ini bertujuan untuk mengeliminasi bufferbloat dan jitter secara real-time guna memastikan pengalaman gaming (khususnya game kompetitif dengan tick-rate tinggi seperti Valorant) tetap stabil pada jalur fiber optic\"; verify technology claims with official docs.",
  "visualDirection": {
    "seedMode": "scaffold-only",
    "requiresProjectSpecificSynthesis": true,
    "selectionAuthority": "agent-llm-after-current-context-repo-evidence-and-live-official-docs",
    "trendStance": "current-context-first-not-offline-preset-first",
    "distinctiveMoves": [
      "Choose one product-specific move from task, audience, content, repo evidence, and docs."
    ],
    "copiedReferenceAllowed": false
  },
  "externalResearchIntake": {
    "userSuppliedResearchPolicy": "read-as-candidate-evidence-not-final-prescription",
    "requireSummaryOfUsedSignals": true,
    "requireFitFiltering": true,
    "requireOfficialDocsVerificationForTechnologyClaims": true,
    "candidateDomains": [
      "visual-direction",
      "motion-and-scroll",
      "ui-primitives-or-rich-media",
      "typography-and-interaction"
    ],
    "finalDecisionAuthority": "project-fit-accessibility-performance-maintainability-official-docs"
  },
  "conceptualAnchor": {
    "mode": "required-when-no-external-research",
    "seedMode": "selection-policy-only",
    "anchorReference": "agent-defined-anchor-reference",
    "requiresAgentSelectionBeforeUiImplementation": true,
    "userResearchAbsencePolicy": {
      "userSuppliedResearchOnly": true,
      "scaffoldSeedDoesNotCountAsResearch": true,
      "priorUiDoesNotCountAsResearch": true,
      "requireAgentLedResearchWhenAvailable": true
    },
    "candidateSelectionPolicy": {
      "considerAtLeast": 3,
      "discardObviousCandidateCount": 2,
      "minimumCandidateDistance": "high",
      "discardPredictableCandidates": true,
      "preferDistinctiveOverSafe": true,
      "doNotRevealHiddenCandidateList": true,
      "outputOnlyChosenAnchor": true
    },
    "creativeCommitmentPolicy": {
      "requiredBeforeComplianceReview": true,
      "recordInDesignDocs": true,
      "requiredCommitmentFields": [
        "specificReferencePoint",
        "signatureMotion",
        "typographicDecision"
      ],
      "rejectGenericQualityWordsOnly": true,
      "specificityFloor": "name-real-material-instrument-artifact-architecture-editorial-genre-cinematic-behavior-exhibition-system-scientific-apparatus-or-industrial-mechanism"
    },
    "forbiddenFinalAnchorTerms": [
      "dashboard",
      "cards",
      "admin-panel",
      "saas-shell",
      "minimalist-interface",
      "safe-admin-layout"
    ],
    "sourceDomains": [
      "complex-physical-engineering",
      "cinematic-spatial-interface",
      "experimental-editorial-structure",
      "scientific-instrumentation",
      "premium-interactive-web-experiences"
    ],
    "visualRiskBudget": {
      "mode": "high-distinctiveness-with-accessibility-and-performance-guardrails",
      "allowRichMotionAndMicroInteraction": true,
      "rejectTimidDefaultWhenAnchorSupportsExpressiveUi": true,
      "requireReducedMotionFallback": true
    },
    "requiredDerivedAxes": [
      "typography",
      "morphology",
      "motion",
      "responsive-composition"
    ],
    "finalAnchorContract": {
      "requiredFields": [
        "name",
        "anchorReference",
        "agentResearchMode",
        "sourceDomain",
        "specificReferencePoint",
        "rationale",
        "signatureMotion",
        "typographicDecision",
        "derivedTokenLogic",
        "visualRiskBudget",
        "motionRiskBudget",
        "cohesionChecks"
      ],
      "derivedTokenLogicAxes": [
        "morphology",
        "motion"
      ],
      "cohesionChecks": [
        "no-dashboard-mental-model",
        "motion-derived-from-anchor"
      ]
    }
  },
  "derivedTokenLogic": {
    "anchorReference": "agent-defined-anchor-reference",
    "colorDerivationSource": "Explain semantic color roles from anchorReference; reject generic palettes without anchor evidence.",
    "spacingDerivationSource": "Explain grid, rhythm, density, and exceptions from anchorReference.",
    "typographyDerivationSource": "Explain display, body, metadata, and data roles from anchorReference.",
    "motionDerivationSource": "Explain duration, easing, choreography, and reduced-motion from anchorReference.",
    "validationRule": "Every token must trace to anchorReference; revise tokens that cannot."
  },
  "motionPaletteDecision": {
    "productCategorySignal": "agent-inferred-starting-heuristic",
    "densityDecisionSource": "Choose motion density from task, content, brand, device, performance, and accessibility. Categories are heuristics.",
    "requiredInteractionStates": [
      "default",
      "hover",
      "focus-visible",
      "active",
      "disabled",
      "loading",
      "empty",
      "error",
      "success",
      "transition"
    ],
    "paletteAutopilotRisks": [
      "dark-slate-default",
      "cream-beige-default",
      "purple-blue-gradient-default",
      "monochrome-template-default",
      "uniform-card-surface-default",
      "generic-grid-wallpaper-default",
      "soft-glow-ai-template-default",
      "cyber-neon-terminal-default"
    ],
    "spatialDecision": "State 3D/canvas/WebGL fit. If omitted, name product-fit reason and replacement interaction quality."
  },
  "aiSafeUiAudit": {
    "status": "agent-must-complete-before-ui-implementation",
    "failureDefinition": "AI-safe UI uses template cards, generic marks, decorative grid or line wallpaper, safe palettes, glow backgrounds, or copied scaffold composition.",
    "interchangeabilityTest": "If this UI can be renamed from aura-sqm to another product category without changing composition, palette, iconography, and motion, revise it.",
    "requiredProductSpecificSignals": [
      "agent-defined-product-specific-data-treatment",
      "agent-defined-product-specific-motion-or-state-behavior",
      "agent-defined-product-specific-morphology-iconography-or-spatial-structure"
    ],
    "paletteExplorationRule": "Use a visually exploratory product-derived palette with WCAG contrast and status clarity.",
    "backgroundPatternRule": "Lines, grids, scanlines, noise, glows, blobs, logos, and geometry must serve a named product function; never use grid or line backgrounds as first-output filler.",
    "aiColorAudit": {
      "status": "agent-must-complete-before-ui-implementation",
      "failureDefinition": "AI color drift uses safe defaults before deriving roles from the product anchor.",
      "autopilotRisks": [
        "cream-editorial-default",
        "dark-slate-dashboard-default",
        "purple-blue-gradient-default",
        "monochrome-minimal-default",
        "cyber-neon-terminal-default",
        "soft-glow-atmosphere-default"
      ],
      "requiredEvidence": [
        "anchor-derived-color-logic",
        "semantic-role-contrast-beyond-surface-decoration",
        "product-specific-color-behavior-that-would-not-transfer"
      ],
      "reviewQuestion": "Why does this palette belong to this product?"
    },
    "motionSpatialCourageAudit": {
      "status": "agent-must-complete-before-ui-implementation",
      "defaultStance": "Treat motion, scroll choreography, canvas, WebGL, and 3D as first-class options.",
      "requiredDecisionFields": [
        "signature-motion-or-interaction",
        "spatial-or-3d-fit",
        "performance-and-reduced-motion-fallback"
      ],
      "rejectionRule": "State a product reason and replacement interaction quality before omitting 3D/canvas.",
      "reviewQuestion": "Is the interaction as expressive as the product can responsibly support?"
    },
    "reviewQuestion": "What visible evidence proves this is product-specific?",
    "blockingByDefault": true
  },
  "libraryResearchStatus": "pending-verification",
  "libraryDecisions": [
    {
      "library": "agent-defined-or-none",
      "purpose": "Verify UI-related libraries against current official docs before imports.",
      "verifiedAt": null,
      "sourceUrl": null,
      "stableVersion": null,
      "fallbackIfUnavailable": "Use native CSS, browser APIs, or existing dependencies."
    }
  ],
  "mathSystems": {
    "typographyScaleRatio": "agent-calibrated-from-content-platform-and-readability",
    "baseGridUnit": "agent-calibrated-from-platform-density-and-implementation-stack",
    "spacingPattern": "agent-defined-from-task-flow-and-viewport-needs",
    "densityMode": "agent-defined-from-user-task-device-and-content-pressure",
    "seedValuesRequireCalibration": true
  },
  "tokenSystem": {
    "sourceOfTruth": "docs/design-intent.json",
    "taxonomyOrder": [
      "primitive",
      "semantic",
      "component"
    ],
    "primitiveColorSpace": "OKLCH",
    "requireSemanticAliases": true,
    "semanticAliasesMutableWithoutComponentRewrite": true,
    "componentTokensConsumeSemantic": true,
    "forbidDirectComponentPrimitiveBypass": true,
    "aliasReferenceStyle": "brace-reference",
    "fallbackPolicy": {
      "forbidRawHexOutsidePrimitives": true,
      "forbidRawSpacingOutsidePrimitives": true,
      "requireDocumentedExceptionForLegacyBypass": true
    },
    "namingConstraints": {
      "forbidCurlyBracesInNames": true,
      "forbidDotsInNames": true,
      "forbidSquareBracketsInNames": true
    }
  },
  "colorTruth": {
    "format": "OKLCH",
    "allowHexDerivatives": true,
    "requirePerceptualLightnessCurve": true,
    "paletteRoles": [
      "agent-defined-semantic-roles"
    ],
    "rolePolicy": "minimum-semantic-scaffold",
    "rolesAreMinimumScaffold": true,
    "rolesMustBeAgentDefined": true,
    "forbidAutopilotPalettesWithoutEvidence": true,
    "intent": "Choose semantic palette roles from the product context \"Aura-SQM adalah sistem kontrol lalu lintas jaringan (Network Traffic Control) dinamis berbasis Golang yang dirancang khusus untuk lingkungan OpenWrt. Proyek ini bertujuan untuk mengeliminasi bufferbloat dan jitter secara real-time guna memastikan pengalaman gaming (khususnya game kompetitif dengan tick-rate tinggi seperti Valorant) tetap stabil pada jalur fiber optic\", repo evidence, and accessibility. Reject scaffold or SaaS palette defaults."
  },
  "crossViewportAdaptation": {
    "adaptByRecomposition": true,
    "touchTargetMinPx": 44,
    "mutationRules": {
      "mobile": "Reorder, merge, or disclose content for mobile. Reject scale-only shrink.",
      "tablet": "Regroup for tablet instead of width-only desktop reduction.",
      "desktop": "Use space intentionally; avoid equal-weight modules without evidence."
    }
  },
  "motionSystem": {
    "allowMeaningfulMotion": true,
    "purpose": "Use expressive motion when it improves hierarchy, continuity, feedback, memorability, or confidence. Verify new motion libraries with official docs.",
    "seedToneLocked": false,
    "densitySource": "task-content-brand-device-accessibility",
    "respectReducedMotion": true
  },
  "componentMorphology": {
    "requireStateBehaviorMatrix": true,
    "preserveIdentityAcrossViewports": true,
    "seedBehaviorsRequireRefinement": true,
    "stateKeys": [
      "default",
      "hover",
      "focus-visible",
      "active",
      "disabled",
      "loading",
      "empty",
      "error",
      "success",
      "transition"
    ],
    "viewportBehavior": {
      "mobile": "Recompose for touch, task priority, and constrained attention.",
      "tablet": "Regroup surfaces for medium width without cloning desktop or mobile.",
      "desktop": "Use space for hierarchy and scanability; avoid template grids."
    }
  },
  "accessibilityPolicy": {
    "hardComplianceFloor": "WCAG-2.2-AA",
    "advisoryContrastModel": "APCA",
    "failOnHardViolations": true,
    "advisoryFindingsDoNotBlockByDefault": true,
    "hardRequirements": {
      "textContrastMinimum": true,
      "nonTextContrast": true,
      "useOfColorOnlyProhibited": true,
      "focusVisible": true,
      "focusAppearance": true,
      "targetSizeMinimum": true,
      "keyboardAccess": true,
      "reflowRequired": true,
      "accessibleAuthenticationMinimum": true,
      "statusMessagesAndDynamicStateAccess": true
    },
    "advisoryChecks": {
      "perceptualContrastReview": true,
      "darkModeContrastTuning": true,
      "typographyReadabilityTuning": true
    }
  },
  "designExecutionPolicy": {
    "representationStrategy": "surface-plan-v1",
    "seedRefinementRequiredBeforeUiImplementation": true,
    "requireSurfacePlan": true,
    "requireComponentGraph": true,
    "requireViewportMutationPlan": true,
    "requireInteractionStateMatrix": true,
    "requireContentPriorityMap": true,
    "requireTaskFlowNarrative": true,
    "requireSignatureMoveRationale": true,
    "requireCreativeCommitmentGate": true,
    "requireStructuredHandoff": true,
    "requireRepoEvidenceAlignment": true,
    "forbidScreenshotDependency": true,
    "handoffFormatVersion": "ui-handoff-v1",
    "requirePerSurfaceMutationOps": true,
    "forbidUniformSiblingSurfaceTreatment": true,
    "zeroBasedRedesignResetsPriorVisualsWhenRequested": true,
    "semanticReviewFocus": [
      "distinctiveness-vs-genericity",
      "contract-fidelity",
      "hierarchy-and-task-priority",
      "component-state-behavior",
      "cross-viewport-mutation"
    ]
  },
  "designExecutionHandoff": {
    "version": "ui-handoff-v1",
    "location": "inline-design-intent",
    "status": "seed-needs-refinement",
    "seedMode": "structure-first-scaffold",
    "requiresTaskSpecificRefinement": true,
    "primaryExperienceGoal": "Define the main fullstack product journey for aura-sqm from repo evidence, brief, and docs.",
    "surfacePlan": [
      {
        "surfaceId": "agent-defined-primary-experience",
        "role": "primary-context-synthesized-by-agent",
        "goal": "Choose the first task path from product evidence; reject template shells.",
        "antiPatterns": [
          "dashboard-default",
          "scale-only-responsive-layout"
        ]
      }
    ],
    "componentGraph": {
      "nodes": [
        {
          "id": "primary-experience",
          "role": "agent-defined-primary",
          "priority": "high"
        },
        {
          "id": "supporting-context",
          "role": "agent-defined-support",
          "priority": "medium"
        }
      ],
      "edges": [
        {
          "from": "primary-experience",
          "to": "supporting-context",
          "relationship": "task-priority-support"
        }
      ]
    },
    "contentPriorityMap": {
      "primary": [
        "agent-defined-core-task-or-reading-path"
      ],
      "secondary": [
        "agent-defined-supporting-context"
      ],
      "deferred": [
        "agent-defined-deferred-or-hidden-content"
      ]
    },
    "viewportMutationPlan": {
      "mobile": {
        "primaryOperation": "agent-defined-mobile-recomposition",
        "requiredSurfaceActions": [
          "choose-mobile-task-order",
          "disclose-or-remove-low-priority-content"
        ],
        "forbiddenPatterns": [
          "scale-only-shrink"
        ],
        "rationale": "Reorder, merge, or disclose content for mobile. Reject scale-only shrink."
      },
      "tablet": {
        "primaryOperation": "agent-defined-tablet-regrouping",
        "requiredSurfaceActions": [
          "define-medium-width-grouping",
          "preserve-task-clarity"
        ],
        "forbiddenPatterns": [
          "uniform-module-grid-without-role-change"
        ],
        "rationale": "Regroup for tablet instead of width-only desktop reduction."
      },
      "desktop": {
        "primaryOperation": "agent-defined-desktop-composition",
        "requiredSurfaceActions": [
          "use-space-to-improve-hierarchy",
          "avoid-equalizing-unrelated-content"
        ],
        "forbiddenPatterns": [
          "interchangeable-dashboard-or-landing-chrome"
        ],
        "rationale": "Use space intentionally; avoid equal-weight modules without evidence."
      }
    },
    "interactionStateMatrix": [
      {
        "componentId": "primary-interaction",
        "states": [
          "default",
          "hover",
          "focus",
          "loading",
          "error"
        ],
        "notes": "Refine states from project language and anchor; reject anonymous panels."
      }
    ],
    "taskFlowNarrative": [
      "Entry: start aura-sqm from real evidence, not a generic opener.",
      "Resolution: define proof, feedback, recovery, and next action."
    ],
    "visualResetStrategy": {
      "activatesWhenUserRequests": [
        "redesign from zero",
        "redesain dari 0"
      ],
      "existingUiAllowedAs": [
        "content-evidence",
        "behavior-evidence",
        "asset-source-evidence"
      ],
      "existingUiForbiddenAs": [
        "palette-source",
        "layout-source",
        "motion-source"
      ],
      "requiredResetAxes": [
        "composition",
        "hierarchy",
        "motion-or-interaction",
        "responsive-information-architecture"
      ]
    },
    "signatureMoveRationale": "Choose one project-specific visual, motion, type, or interaction move.",
    "creativeCommitment": {
      "status": "agent-must-complete-before-ui-implementation",
      "requiredFields": [
        "specificReferencePoint",
        "signatureMotion",
        "typographicDecision"
      ],
      "failureMode": "generic quality words without a real-world reference fail"
    },
    "implementationGuardrails": {
      "requireBuildFromHandoff": true,
      "requireGapNotesBeforeFallback": true,
      "forbidGenericLayoutFallbackWithoutReason": true
    }
  },
  "reviewRubric": {
    "version": "ui-rubric-v1",
    "genericityAutoFail": true,
    "dimensions": [
      {
        "key": "distinctiveness",
        "blockingByDefault": true,
        "question": "Is the UI authored and product-specific?"
      },
      {
        "key": "contractFidelity",
        "blockingByDefault": true,
        "question": "Does the UI follow contract, priorities, and accessibility?"
      },
      {
        "key": "visualConsistency",
        "blockingByDefault": false,
        "question": "Do type, spacing, color, and states form one system?"
      },
      {
        "key": "heuristicUxQuality",
        "blockingByDefault": false,
        "question": "Does the UI preserve clarity, feedback, and confidence?"
      },
      {
        "key": "motionDiscipline",
        "blockingByDefault": false,
        "question": "Is motion purposeful, performant, reduced-motion-safe, and on-tone?"
      }
    ],
    "genericitySignals": [
      "offline-prescribed-style-used-as-final-direction",
      "unresearched-library-or-framework-choice",
      "missing-conceptual-anchor-without-external-research",
      "visual-decisions-not-derived-from-conceptual-anchor",
      "ai-safe-ui-template-look",
      "ai-color-default-palette-without-product-role-behavior",
      "brandless-clean-template-look",
      "interchangeable-product-renaming-test-fails",
      "decorative-grid-or-glow-wallpaper-without-product-function",
      "safe-cream-slate-or-monochrome-palette-used-as-readability-excuse",
      "generic-abstract-logo-or-iconography",
      "timid-anchor-that-renames-dashboard-or-admin-shell",
      "motion-suppressed-without-accessibility-or-performance-reason",
      "motion-or-3d-omitted-from-fear-without-fit-analysis",
      "scale-only-responsive-layout",
      "zero-based-redesign-kept-prior-visual-dna",
      "restyle-instead-of-recomposition"
    ],
    "validBoldSignals": [
      "single-cohesive-conceptual-anchor",
      "high-variance-candidate-selection",
      "context-derived-visual-direction",
      "three-at-a-glance-product-specific-signals",
      "visually-exploratory-accessible-palette-derived-from-product",
      "audacious-accessible-palette-with-product-role-behavior",
      "background-or-geometry-serves-product-function",
      "motion-or-spatial-experience-derived-from-anchor",
      "explicit-3d-canvas-fit-or-nonfit-decision",
      "responsive-recomposition-by-task-priority",
      "purposeful-motion-with-reduced-motion-path"
    ],
    "reportingRules": {
      "mustExplainGenericity": true,
      "mustSeparateTasteFromFailure": true,
      "contractFidelityOverridesPersonalTaste": true
    }
  },
  "contextHygiene": {
    "continuityMode": "opt-in-only",
    "allowedSources": [
      "current-repo-evidence",
      "current-user-brief",
      "current-project-docs",
      "explicitly-approved-current-task-constraints"
    ],
    "taintedSources": [
      "prior-chat-visual-memory",
      "unrelated-project-aesthetics",
      "remembered-screenshots-without-current-approval",
      "generic-template-recall"
    ],
    "repoEvidenceOverridesMemory": true,
    "requireExplicitContinuityApproval": true,
    "forbidCarryoverWhenUnapproved": true,
    "approvedExternalConstraintUsage": "Convert approved external constraints into current-project rules; do not imitate source surfaces.",
    "driftSignals": [
      "palette-reused-without-brief-support",
      "prior-ui-visual-dna-carried-into-reset-request"
    ]
  },
  "forbiddenPatterns": [
    "offline-prescribed-style-used-as-final-direction",
    "missing-conceptual-anchor-without-external-research",
    "visual-decisions-not-derived-from-conceptual-anchor",
    "ai-safe-ui-template-look",
    "ai-color-default-palette-without-product-role-behavior",
    "brandless-clean-template-look",
    "interchangeable-product-renaming-test-fails",
    "decorative-grid-or-glow-wallpaper-without-product-function",
    "safe-cream-slate-or-monochrome-palette-used-as-readability-excuse",
    "generic-abstract-logo-or-iconography",
    "timid-anchor-that-renames-dashboard-or-admin-shell",
    "motion-suppressed-without-accessibility-or-performance-reason",
    "motion-or-3d-omitted-from-fear-without-fit-analysis",
    "scale-only-responsive-layout",
    "zero-based-redesign-kept-prior-visual-dna",
    "restyle-instead-of-recomposition",
    "single-safe-typographic-family-without-role-contrast-or-rationale"
  ],
  "validationHints": [
    "rejectArbitraryHexOnlyPalette",
    "requireViewportMutationRules",
    "requirePerceptualColorRationale",
    "requireTokenLayering",
    "requireTokenAliasingPlan",
    "allowHexDerivatives",
    "requireMotionRationale",
    "requireStateMorphology",
    "requireAccessibilitySplit",
    "requireWcagHardFloor",
    "requireStructuredDesignExecutionPolicy",
    "requireStructuredDesignHandoff",
    "requireVisualResetStrategyWhenZeroBasedRedesignRequested",
    "requireConceptualAnchorWhenNoExternalResearch",
    "requireAgentLedAnchorResearchWhenUserResearchMissing",
    "rejectTimidDashboardAnchor",
    "requireReviewRubric",
    "requireGenericityExplanation",
    "genericityAutoFail",
    "requireSignatureMove",
    "rejectTemplateNeutralLayout",
    "requireAiSafeUiAudit",
    "requireAiColorAudit",
    "rejectAiColorDefaults",
    "requireMotionSpatialCourageAudit",
    "requireExplicit3dCanvasFitDecision",
    "rejectSafetyAsCreativitySubstitute",
    "rejectAiSafeUiTemplateLook",
    "requireThreeProductSpecificSignals",
    "rejectDecorativeBackgroundPatternsWithoutProductFunction"
  ],
  "requiredDesignSections": [
    "Design Intent and Product Personality",
    "Audience and Use-Context Signals",
    "Visual Direction and Distinctive Moves",
    "Color Science and Semantic Roles",
    "Typographic Engineering and Hierarchy",
    "Spacing, Layout Rhythm, and Density Strategy",
    "Token Architecture and Alias Strategy",
    "Responsive Strategy and Cross-Viewport Adaptation Matrix",
    "Interaction, Motion, and Feedback Rules",
    "Component Language, Morphology, and Shared Patterns",
    "Context Hygiene and Source Boundaries",
    "Accessibility Non-Negotiables",
    "Anti-Patterns to Avoid",
    "Implementation Notes for Future UI Tasks"
  ],
  "implementation": {
    "requiredDeliverables": [
      "docs/DESIGN.md",
      "docs/design-intent.json"
    ],
    "requireDesignRationale": true,
    "requireDistinctVisualDirection": true,
    "requireMachineReadableContract": true,
    "requireViewportMutationRules": true,
    "requirePurposefulMotionGuidelines": true,
    "requireRecognizableVisualBet": true,
    "requireConceptualAnchor": true,
    "bootstrapPrompt": ".agent-context/prompts/bootstrap-design.md",
    "autoLoadedRuleFiles": [
      ".agent-context/prompts/bootstrap-design.md",
      ".agent-context/rules/frontend-architecture.md"
    ],
    "disallowedAutoLoadedRuleFiles": [
      ".agent-context/rules/database-design.md",
      ".agent-context/rules/docker-runtime.md",
      ".agent-context/rules/microservices.md"
    ]
  }
}
```

## Required Execution
1. Create or update docs/DESIGN.md with complete content.
2. Create or update docs/design-intent.json with machine-readable design intent.
3. Keep both files synchronized: the markdown explains the rationale and the JSON captures the machine-readable contract.
4. Encode accessibility as a split policy: WCAG 2.2 AA hard floor, APCA advisory readability tuning, and explicit hard-vs-advisory checks.
5. Encode structured design execution as policy: representation strategy, surface plan, component graph, content-priority map, viewport mutation plan, interaction-state matrix, semantic review focus, and structured handoff requirements.
6. Encode an explicit structured handoff inside docs/design-intent.json: surface plan, component graph, content-priority map, viewport mutation plan, interaction-state matrix, task-flow narrative, and signature move rationale.
7. Encode a stable review rubric: required dimensions, genericity signals, valid bold signals, and reporting rules that separate taste from real failure.
8. Make the handoff executable without screenshot dependency. The contract must still guide high-precision UI generation from repo evidence and changed code alone.
9. Preserve repoEvidence.designEvidenceSummary when onboarding or detector evidence exists instead of discarding it.
10. If repoEvidence.designEvidenceSummary.structuredInspection exists, use it as stronger evidence for class surfaces, inline style bypasses, and expression-backed UI structure before defaulting to generic assumptions.
11. Ensure both files stay project-specific, dynamic, and practical for implementation and review. The seed may guide structure, but it must not decide style offline.
12. Keep visualResetStrategy in the machine-readable handoff so reset-language tasks cannot quietly become restyles of the previous UI.
13. Preserve externalResearchIntake so user-provided research becomes reviewed evidence without turning into an offline style or dependency preset.
14. Preserve conceptualAnchor so prompt-only UI work has one cohesive non-template concept instead of a mixed collection of bold but unrelated visual decisions.
15. Record conceptualAnchor.agentResearchMode, specificReferencePoint, signatureMotion, typographicDecision, visualRiskBudget, motionRiskBudget, and cohesionChecks so the final UI cannot quietly fall back to a timid dashboard/admin mental model.
16. Preserve derivedTokenLogic, libraryResearchStatus, and libraryDecisions so token choices and dependency uncertainty stay visible before implementation.
17. After the contract exists, use it as a first-class source for future UI tasks.
