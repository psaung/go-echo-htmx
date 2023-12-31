### Go Echo HTMX Boilerplate

Experimenting Echo framework with HTMX.

## Running Echo Server

To setup postgre server:

```bash
make setup-local
```

To run echo server:

```bash
make run
```

## Tailwinds and PostCss setup

In order to porvide tailwind css with postcss, this folder will suppose to create tailwindcss for our golang-echo-htmx. Postcss stuff was separated out of the golang code and organized under `postcss` directory.
This project was created using `bun init` in bun v1.0.4. [Bun](https://bun.sh) is a fast all-in-one JavaScript runtime.

To install dependencies:

```bash
bun install
```

To run:

```bash
bun run build
`or`
make build-css
```

To watch:

```bash
bun run build:watch
make watch-css
```
