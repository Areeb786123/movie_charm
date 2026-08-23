FROM node:20-alpine AS dependencies

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci

FROM dependencies AS build

COPY . .
RUN npm run build

FROM node:20-alpine AS runner

WORKDIR /app
ENV NODE_ENV=production

COPY --from=build /app/package.json ./
COPY --from=build /app/package-lock.json ./
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/.next ./.next

EXPOSE 3000
CMD ["npm", "start"]
