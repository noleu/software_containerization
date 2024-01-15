FROM oven/bun:latest as dev

WORKDIR /app

COPY package.json /app
COPY bun.lockb /app

RUN bun install --silent

COPY . /app

CMD ["bun", "run", "dev"]

FROM oven/bun:latest as prod

WORKDIR /app

COPY package.json /app
COPY bun.lockb /app

RUN bun install -g http-server
RUN bun install --silent --production

COPY . /app
RUN bun run build

CMD ["http-server", "dist"]
