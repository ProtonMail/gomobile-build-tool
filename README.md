# gomobile-build-tool

A go program to build the go libraries needed for mobile clients.

## Prerequisites

- Install golang. (the script has been tested with go 1.15 only)
- Apple platforms: You need xcode (version 12 or above) installed
- Android platforms:
  - You need android and android ndk installed.
  - You need `$ANDROID_HOME` to point to your installation of Android sdk
  - You need to set `$ANDROID_HOME/ndk-bundle` or `$ANDROID_NDK_HOME` to point
    to your android ndk


## Build instructions

- Clone the repository
- Configure the the build (see below).
- Inside the cloned repository, run
```
make build cfg=/path/to/config/file
```
This should produce the artifacts inside the `out` directory, set in the
configuration.

You can remove previous builds with `make clean`.

## Configuration of the build

You need to provide the builder program a configuration file.
We provide a few example configuration files in `examples/`.

- `"go_version"`: Must match the version number of your local golang installation
- `"build_dir"` : the directory to contain the generated the build script, default is `"build"`
- `"out_dir"` : the directory that will contain the artifacts, default is `"out"`
- `"go_mobile_dir"` : the directory where the fork of gomobile was clone, default is `"mobile"`
- `"go_mobile_flags"` : a list of flags given to the `gomobile bind` command, see [this](https://godoc.org/golang.org/x/mobile/cmd/gomobile#hdr-Build_a_library_for_Android_and_iOS) for a list of flags.
- `"build_name"` : the name of the produced artifacts
- `"build_tag"` : Used by gitlab automated builds to tag the builds
- `"java_pkg"` : (for android) the name of the java package representing the go code, it defaults to the `"build_name"`.
- `"min_ios_version"`: (for ios) the minimum version of iOS we want to support. (Forwarded to the `-iosversion` option of gomobile)
- `"targets"`: a list of platforms (`"android"` and `"apple"`) to build artifacts for, if not provided, the program builds for both platforms.
- `"requirements"` : used to specify a list of packages to include in the artifacts.
    - `"module"`: the go module containing the package
        - `"path"`: the import path of the module
        - `"version"`: specifies which module version to use, either a commit tag (e.g `"v2.0.1"`) or a commit hash, or `"latest"`.
    - `"packages"`: a list of package names from the module to include in the artifacts (without the full import path of the module). If not provided, the artifact will include the whole module.
- `"replacements"`: a list of module replacements for the artifacts
    - `"old"`: the module to replace
        - `"path"`: the path of the module
        - `"version"`: the version of the module to replace, if not included we replace all versions
    
    - `"new"`: the module to replace with
        - `"path"`: the path of the module
        - `"version"`: the version of the module to replace with.
    - `"local_path"`: a local path to the replacement module (instead of a remot go module)


