package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483648, 781, 0, 139}
    v3 := make(map[string]string)
    var v4 uint32 = 145
    var v5 int8 = 100
    v5 -= v5
    v6 := v3[("" + ((((("" + "foo bar") + ("hello" + (("hello" + "foo bar") + ("a" + "foo bar")))) + "a") + ("hello" + "\n\t")) + ""))]
    fmt.Println(v6, v6, v6)
    for fi7 := 0; fi7 < 9; fi7++ {
        switch (fi7 >> ((255) & 63)) {
        case 0:
            v8 := GenS1{F0: -1, F1: -32768, F2: 128}
            _ = v8
        case 24:
            var v9 float32 = 1.5
            _ = v9
        default:
            if ((!false) && (float32(1.5) >= 1.5)) {
                var v10 bool = (v5 != 33)
                v11 := v1[fi7]
                _ = v10
                _ = v11
            }
        }
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        switch v6 {
        case "":
            v6 = ""
        case "a":
            v6 = v6
        case "\n\t":
            v4 = v4
        }
        _ = fi7
    }
    var v12 uint64 = 1
    v5 = 100
    v12 = (v12 + (v12 * v12))
    switch "hello" {
    case "foo bar":
        var v13 float64 = (1.5 + (-2.25 - -2.25))
        _ = v13
    case "":
        fmt.Println(v4, v5)
    case "a":
        for fi14 := 0; fi14 < 19; fi14++ {
            if (!(((!((fi14 ^ fi14) <= (fi14 >> ((2) & 63)))) && false) && (v12 != 42))) {
                v15 := GenS1{F0: -2147483648, F1: -357, F2: 42}
                _ = v15
            }
            _ = fi14
        }
    default:
        fmt.Println(v5, v6, v6)
    }
    if (true || true) {
        v2 = append(v2, (v4 % (((v4 - v4)) | 1)))
        v5 -= v5
        var v16 bool = false
        _ = v16
    } else {
        v17 := v1[-9223372036854775808]
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
