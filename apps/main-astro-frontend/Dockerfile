# syntax=docker/dockerfile:1

# Comments are provided throughout this file to help you get started.
# If you need more help, visit the Dockerfile reference guide at
# https://docs.docker.com/go/dockerfile-reference/

# Want to help us make this template better? Share your feedback here: https://forms.gle/ybq9Krt8jtBL3iCk7

#Define ARGs for the image versions
ARG NODE_VERSION=24-alpine
ARG PNPM_VERSION=10.10.0
ARG NGINX_VERSION=stable-alpine-slim

#Stage 1: Build the application
FROM node:${NODE_VERSION} AS builder

# Use production node environment by default.
ARG NODE_ENV=production
ENV NODE_ENV=${NODE_ENV}

# Install pnpm.
RUN --mount=type=cache,target=/root/.npm \
    npm install -g pnpm@${PNPM_VERSION}

WORKDIR /usr/src/app

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /root/.local/share/pnpm/store to speed up subsequent builds.
# Leverage a bind mounts to package.json and pnpm-lock.yaml to avoid having to copy them into
# into this layer.
RUN --mount=type=bind,source=package.json,target=package.json \
    --mount=type=bind,source=pnpm-lock.yaml,target=pnpm-lock.yaml \
    --mount=type=bind,source=pnpm-workspace.yaml,target=pnpm-workspace.yaml \
    --mount=type=bind,source=patches/@nxtensions__astro@19.0.1.patch,target=patches/@nxtensions__astro@19.0.1.patch \
    --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --prod --frozen-lockfile

# Copy the rest of the source files into the image.
COPY . .

# Run the application.
RUN pnpm nx run main-astro-frontend:build

# Stage 2: Serve the application with nginx
FROM nginx:${NGINX_VERSION} AS runtime

COPY infra/docker/nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /usr/src/app/dist/apps/main-astro-frontend /usr/share/nginx/html

EXPOSE 8080