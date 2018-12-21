# Gotron Examples
Examples showing how to use gotron api 

Use npm install script and start the application
```
cd ui && npm install && npm run build
go build && ./example
```

## Frontend Development Workflow
Take a look into [ui/js](https://github.com/Equanox/gotron/tree/master/ui/js), [ui/react](https://github.com/Equanox/gotron/tree/master/ui/react),
[ui/typescript-react](https://github.com/Equanox/gotron/tree/master/ui/typescript) or [ui/vue](https://github.com/Equanox/gotron/tree/master/ui/vue) for details.

For plain javascript (default) use

    cd ui && npm run build  

For other frontend use

    cd ui && npm run build:${frontend}

where ${frontend} is one out of (js|react|typescript|vue).

Then type

    go build && ./example

to bring up go backend and electron frontend.

Reload updated index.js using 'r' key.
