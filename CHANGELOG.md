# CHANGELOG

## v0.1.0-alpha.4:

### Breaking Changes

- `ViewInBrowser` component now does not accept any params
- Function definition for `Image` has changed:

  ```go
  // Old ❌
  templateless.NewContent().
    Image(
      "https://example.com/image.jpg",
      Some("https://example.com"),
      Some(200), // Width
      Some(100), // Height
      Some("Alt Text"),
    ).
    Build()
  
  // New (just image) ✅
  templateless.NewContent().
    Image("https://example.com/image.jpg").
    Build()
  
  // New (clickable image with custom attributes) ✅
  url := "https://example.com"
  width := 300
  height := 200
  alt := "Alt text"
  templateless.NewContent().
    Component(
      components.NewImage(
        "https://placekitten.com/300/200",
        &url,
        &width,
        &height,
        &alt,
      ),
    ).
    Build()
  ```

### New Features
- New social icons: `Mastodon` and `Rss`
- New `StoreBadges` component
- New `QrCode` component
- New `Signature` component
- New [examples](examples)

### Enhancements
- Updated README
- Dependency updates

## v0.1.0-alpha.3:
- `README.md`: notice about test mode
- Support for test mode logging

## v0.1.0-alpha.2:
- `README.md` cleanup (listing of components)
- Added `examples/confirm_email`
- Renamed `SetFooter()` to `Footer()`

## v0.1.0-alpha.1:
- Introduced `CHANGELOG.md`
- Introduced new services as social icons:
  - `Phone` (converts into a link with `tel:` prefix)
  - `Facebook`
  - `YouTube`
  - `Instagram`
  - `LinkedIn`
  - `Slack`
  - `Discord`
  - `TikTok`
  - `Snapchat`
  - `Threads`
  - `Telegram`

## v0.1.0-alpha.0:
- Initial implementation