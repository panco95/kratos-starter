# Micro Service Starter

Run Server:
- cd app/gateway/interface && kratos run
- cd app/service/user && kratos run

Make:
- make api
- make config
- make errors
- make validate

Wire Inject:
- cd app/gateway/interface/cmd/server && wire
- cd app/user/service/cmd/server && wire

Build & Run:
- make build
- ./server -conf configs/config.yml -log logs