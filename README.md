# libimagecashletter

`libimagecashletter` is a Go library which wraps [Moov's `imagecashletter`
library](https://github.com/moov-io/imagecashletter) and exports C bindings
for parsing and encoding X9 ICL (Image Cash Letter) files.

## Building

This project uses Go v1.13 or higher. See [Golang's install
instructions](https://golang.org/doc/install) for help setting up Go.

To build `libimagecashletter.so` and generate its accompanying C header file,
run:

```sh
make build
```

After a successful build, you'll be able to find the `*.so` and `*.h` files in
the `bin` directory at the root of this repository.
