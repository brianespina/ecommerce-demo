# --- Stage 1: Build the SvelteKit app ---
FROM node:20-alpine AS builder

# Set working directory inside the container
WORKDIR /app

# Copy package files first (better caching for dependencies)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the project
COPY . .

# Build the app (output goes to /app/build)
RUN npm run build


# --- Stage 2: Run the built app with Node ---
FROM node:20-alpine

WORKDIR /app

# Copy only the built files and production deps
COPY --from=builder /app/build ./build
COPY --from=builder /app/package*.json ./

# Install only production dependencies
RUN npm install --omit=dev

# Expose the port SvelteKit will run on
EXPOSE 3000

# Start the app
CMD ["node", "build"]
