FROM node:lts as development

WORKDIR /app
COPY package.json .
COPY . .
RUN yarn install

CMD [ "yarn", "run", "dev" ]

FROM node:lts as test

WORKDIR /app
COPY package.json .
COPY . .
RUN yarn install

CMD [ "yarn", "run", "test" ]