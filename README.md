# Linking multiple modules in Extism

This example shows how you can link multiple modules in the same Manifest. `lib` provides the `Capitalize` function, which is used by `main`.

To run the example:

1. Build each plugin:

```
cd lib
make
```

```
cd main
make
```

2. Run the host app:

```
cd host
go run .
```

You should get something like this:
```
Output: Hello, MO
```