# Simple, Basic, Stupid Reverse Proxy

Hi there! This is a project that I'm coding to learn how reverse proxies work. This project highly affect by Microsoft's [YARP](https://microsoft.github.io/reverse-proxy/) project

I know, there's a lot of bad stuff in this warehouse. But you may want to review

All configs are placed into the `config.yaml` file

This is an example config

```yaml
routes:
  route-1:
    port: 8083
    address: 'http://localhost'
    type: header
    headers:
      header-1:
        values:
          - '8083'
          - first
  route-2:
    port: 8084
    address: 'http://localhost'
    type: header
    headers:
      header-2:
        values:
          - '8084'
          - second
  route-3:
    port: 8085
    address: 'http://localhost'
    type: header
    headers:
      header-3:
        values:
          - '8085'
          - third
  route-4:
    port: 8086
    address: 'http://localhost'
    type: path
    paths:
      - '/'
      - '/about'
      - '/contact'
```

According to these configs, from **route-1** to **route-3**, looks for header configs. The client should send a header to communicate with servers.


For example, if a client sends header `header-1` and with the value `first`, the request will redirect to the `http://localhost:8083`

If **route-1** and **route-2** have the same headers and values, the first one will be using.

## Run

To run this project;

`go run .`

## Help

**Note:** I'm looking for help to make [matcher.go](/internal/matcher.go) file clean and fast :)

## Resources

- https://www.youtube.com/watch?v=vlPCAEUCCa0
- https://gist.github.com/thurt/2ae1be5fd12a3501e7f049d96dc68bb9
- https://pkg.go.dev/gopkg.in/yaml.v3
- https://microsoft.github.io/reverse-proxy/articles/header-routing.html