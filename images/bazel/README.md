# Bazel

This bazel image is derived from the `gcr.io/cloud-builders/bazel` image, which is no longer being maintained by the cloud build team.


## Usage Details

This is a tool builder to simply invoke [`bazel`](https://bazel.io) commands.

Arguments passed to this builder will be passed to `bazel` directly.

The latest available version of `bazel` is used.