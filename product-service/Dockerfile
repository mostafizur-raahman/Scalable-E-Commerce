FROM node:18-alpine

WORKDIR /app

# Copy only the package.json and yarn.lock files for caching dependencies
COPY package.json yarn.lock ./

# Install dependencies with yarn
RUN yarn install

# Copy the rest of the application files
COPY . .

EXPOSE 5001
CMD [ "node", "index.js" ]
