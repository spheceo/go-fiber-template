# {{ project_name }}

This repository provides a minimal Go API template built with the Fiber framework and intended for deployment on Vercel. 

It uses the [Just](https://github.com/casey/just) task runner for local development, but this can be replaced with alternatives.

## Getting Started

First, run the development server:

```
just dev
```
or
```
air
```

`air` is a live-reload utility for Go applications. It watches your project files and automatically rebuilds/restarts the app whenever code changes, which speeds up local development and shortens the edit-run-test loop.

Open [http://localhost:8080](http://localhost:8080) with your browser or API client to see the result.

You can start editing the API by modifying `api/index.go`. The API auto-updates as you edit the file.

## Deploy on Vercel

The recommended way to deploy this Go Fiber app is to use the [Vercel Platform](https://vercel.com/new).

Check out the [Vercel functions documentation](https://vercel.com/docs/functions/runtimes/go) for more details.
