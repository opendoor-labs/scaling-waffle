# Scaling Waffle

Available Envvars for Client (defined in `client.sh`:

| Name | Default | Notes |
|---|---|---|
| `SERVER_ADDRESS` | `server:8050` | Location of `server` container. |
| `CONN_COUNT` | 60000  | Number of connections to create. |
| `CONN_PERIOD_MS`  | 10 | How long to wait between connection creations. |

# Notes
- Ensure target device has sufficent RAM available, or the client process may be OOM-killed.


# Deploy - Balena
```
balena push
```

# Run - Locally
```
docker-compse up --build
```
