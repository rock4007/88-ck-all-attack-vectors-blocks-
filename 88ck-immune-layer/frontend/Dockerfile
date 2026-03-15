FROM node:20-bookworm-slim AS builder
WORKDIR /app
COPY package.json .
RUN npm install
COPY . .
RUN npm run build

FROM gcr.io/distroless/nodejs20-debian12:nonroot
WORKDIR /app
COPY --from=builder /app/dist /app/dist
COPY --from=builder /app/server.mjs /app/server.mjs
EXPOSE 4173
USER nonroot:nonroot
CMD ["server.mjs"]
