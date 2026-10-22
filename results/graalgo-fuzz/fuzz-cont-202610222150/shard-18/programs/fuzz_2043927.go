package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    var v1 int16 = 403
    var v2 uint32 = 0
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{(v2 * v2)}
    var v6 uint32 = (v2 >> ((4294967295) & 63))
    v2 = 2147483648
    if (335 < (-(v1))) {
        var v7 float32 = 0.1
        _ = v7
    } else {
        var v8 string = ("foo bar" + "hello")
        _ = v8
    }
    v2 = 540
    v1 = 42
    switch ("a" + ((((("" + "") + "a") + ("foo bar" + ("" + ("\n\t" + "\n\t")))) + "") + "hello")) {
    case "foo bar":
        var v9 bool = false
        _ = v9
    case "a":
        v10 := GenS1{F0: 255, F1: 128, F2: 32767}
        _ = v10
    }
    v6 += (v6 - (v2 << ((42) & 63)))
    var v11 uint64 = 798
    v1 -= v1
    switch "" {
    case "a":
        var v12 uint32 = v2
        _ = v12
    case "foo bar":
        var v13 uint32 = (v6 & v6)
        _ = v13
    case "hello":
        var v14 float64 = (1.5 / -2.25)
        _ = v14
    }
    var v15 float32 = ((((((3.14159 - (float32((1.5 + 1.5)) / 1.0)) * float32((((1e10 - 1.5) / 3.14159) + 3.14159))) - float32((0.1 * (3.14159 * -2.25)))) - (1e10 * 0.1)) + 3.14159) - -2.25)
    var v16 uint16 = 419
    v6 = 225
    if (3.14159 > 0.1) {
        v6 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
