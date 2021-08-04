# Go + Svelte Websocket Chat

A Websocket Server to allow basic chat. Client is built with Svelte. Server is built with Go.

## Scripts

Scripts are managed with `yarn` or `npm`.

| Script       | Description                                                         |
| ------------ | ------------------------------------------------------------------- |
| build        | Build client and server                                             |
| dev          | Serves client on port 5000 and runs go project on port 3000         |
| start:svelte | Serves the current build of client application                      |
| start:all    | Builds project, and starts the built binary and runs `start:svelte` |
