FROM alpine

ARG FIREBASE_TOOLS_VERSION=14.2.1

RUN apk --no-cache add nodejs npm openjdk17-jre-headless
ENV NODE_ENV=production
RUN npm install -g firebase-tools@${FIREBASE_TOOLS_VERSION} && \
    npm cache clean --force

WORKDIR /firebase-tools

COPY firebase.json storage.rules ./

RUN firebase setup:emulators:pubsub && \
    firebase setup:emulators:storage && \
    firebase setup:emulators:ui

CMD [ "firebase", "emulators:start", "--project", "demo-emulator" ]