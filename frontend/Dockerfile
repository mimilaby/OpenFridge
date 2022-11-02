# syntax=docker/dockerfile:1

FROM node:lts-alpine AS base

# Development stage
# ---------------------------------------
FROM base AS development

WORKDIR /node
COPY package*.json ./
RUN \
  NODE_ENV=development && \
  npm ci && \
  npm cache clean --force

WORKDIR /node/app
EXPOSE 5173
CMD [ "npm", "run", "dev" ]
