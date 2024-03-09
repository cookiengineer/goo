
# goo.js

Opinionated minimalistic, static and semantic UI framework.


## Motivation

This UI library is an opinionated take on how to use semantic HTML5
and modern CSS5 to implement typical web application layouts.

Its intended use is a `go` based tech stack, which uses the `webview/webview`
bindings to spawn a local webview that points to the self-hosted local
backend API.


## Features

- Compiles to Web Assembly.
- Embraces Go-specific language patterns.
- Embraces static, semantic HTML5 syntax.
- Separates between Content, Design, and Layout.
- Embraces CSS Layers (`@layer` in CSS5) for separation of framework-specific and application-specific stylesheets.
- Aims to be 100% JavaScript free, meaning everything should be implemented using `syscall/js`.


## Design Choices

- The HTML _always_ uses the identical semantic structure.
- The `data-layout` property on `body` sets the App Layout (e.g. `body[data-layout="app"]`)
- The `data-theme` property on `body` sets the App Theme (e.g. `body[data-theme="default"]`)
- The `enctype` property on `form` supports `application/json` to send generated structs to the server.
- The optional `data-struct` property on `form` will validate the struct before it is send to the API.


## Usage

Examples can be found inside the [examples](/examples) folder. The Basic Example's
[build.go](/examples/basic/build.go) shows how to use the exposed `goo` API.

Inside the HTML file, you can statically include the bundled `css` and `js` files.
Afterwards, you can decide to either develop your Web App in a static manner, or
with the exposed `go` APIs.

```html
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="/goo/app-default.css">
		<script defer src="goo/app-default.js"></script>
	</head>
	<body>
		(...)
	</body>
</html>
```

## TODO (Work in Progress)

- [ ] Data validation via given `folder` parameter, which contains the folder exposing a `structs` package.


## License

MIT

