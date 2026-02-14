# Initial Implementation Plan

## Goal

Create a runnable Go + Ebitengine project that builds all Galaxians-style game sprites in code for the first iteration.

## Scope (Iteration 1)

- Build sprite assets in code only (no external image files)
- Render sprites in-app and verify animation frames
- Establish interfaces and architecture for future gameplay iterations

## Build Steps

1. Initialize module and dependencies.
2. Create app skeleton (`main`, game loop, fixed logical screen).
3. Implement `spritegen` core:
   - palette constants
   - mask-to-image builder
   - deterministic sprite generation entrypoint
4. Implement sprite builders:
   - player (2 frames + explosion)
   - enemy variants (red/purple/flagship/escort with flight+dive frames)
   - bullets (player + 2 enemy styles)
   - effects (small/large explosions, sparkle)
   - UI glyphs (digits + needed letters)
5. Implement sprite registry:
   - IDs, sets, frame metadata, lookup/list APIs
6. Implement gallery scene:
   - grid preview by category
   - animated frames
   - labels using generated glyphs
   - debug counters
7. Add tests:
   - presence of expected IDs
   - frame counts per sprite
   - dimension checks
   - non-empty frame checks
   - deterministic hash checks
8. Run `go test` and `go build` as acceptance gates.

## Interfaces

- `SpriteID`
- `Frame`
- `SpriteSet`
- `Registry`
- `GenerateAll(cfg Config) (map[SpriteID]SpriteSet, error)`

## Acceptance Criteria

1. App launches and displays all sprite categories.
2. Frame animations visibly cycle.
3. No external image assets exist.
4. Tests pass for shape, counts, and deterministic output.

## Assumptions

1. “Pixel-perfect” means faithful Galaxian-style silhouettes, not ROM-byte extraction.
2. Startup-time generation only.
3. Iteration 1 excludes gameplay systems (waves, collision, scoring, audio).
