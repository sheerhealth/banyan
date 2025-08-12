FROM gcr.io/gcp-runtimes/ubuntu_20_0_4

ADD bazel.sh /builder/bazel.sh
ARG BAZEL_VERSION=latest
ARG DOCKER_VERSION=5:24.0.7-1~ubuntu.20.04~focal

RUN \
    set -eux; \
    apt-get update; \
    apt-get -y install --no-install-recommends \
        python \
        python3 \
        python-pkg-resources \
        python3-pkg-resources \
        software-properties-common \
        unzip \
        gnupg \
        curl \
        ca-certificates \
        apt-transport-https \
        openjdk-8-jdk; \
    \
    # Install Git >2.0.1
    add-apt-repository -y ppa:git-core/ppa; \
    \
    # Install bazel (https://docs.bazel.build/versions/master/install-ubuntu.html)
    curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor -o /usr/share/keyrings/bazel-archive-keyring.gpg; \
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/bazel-archive-keyring.gpg] http://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list; \
    \
    # Install Docker (https://docs.docker.com/engine/installation/linux/docker-ce/ubuntu/#install-using-the-repository)
    mkdir -p /etc/apt/keyrings; \
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /etc/apt/keyrings/docker.gpg; \
    echo \
      "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
      $(lsb_release -cs) stable edge" | tee /etc/apt/sources.list.d/docker.list > /dev/null; \
    \
    # Update apt repositories and install versioned software.
    apt-get update; \
    apt-get -y install --no-install-recommends git; \
    if [ "${BAZEL_VERSION}" = "latest" ]; then \
        apt-get -y install --no-install-recommends bazel; \
    else \
        apt-get -y install --no-install-recommends bazel-${BAZEL_VERSION}; \
        ln -s /usr/bin/bazel-${BAZEL_VERSION} /usr/bin/bazel; \
    fi; \
    apt-get dist-upgrade -y; \
    apt-get install -y --no-install-recommends docker-ce=${DOCKER_VERSION} docker-ce-cli=${DOCKER_VERSION}; \
    \
    # Clean up
    apt-get autoremove -y; \
    apt-get clean; \
    rm -rf \
       /var/cache/debconf/* \
       /var/lib/apt/lists/* \
       /var/log/* \
       /tmp/* \
       /var/tmp/*; \
    \
    mv /usr/bin/bazel /builder/bazel; \
    mv /builder/bazel.sh /usr/bin/bazel; \
    \
    # Unpack bazel for future use.
    bazel version

# Store the Bazel outputs under /workspace so that the symlinks under bazel-bin (et al) are accessible
# to downstream build steps.
RUN mkdir -p /workspace
RUN echo 'startup --output_base=/workspace/.bazel' > ~/.bazelrc

ENTRYPOINT ["bazel"]
