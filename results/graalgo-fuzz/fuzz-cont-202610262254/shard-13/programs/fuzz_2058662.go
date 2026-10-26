package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }

func main() {
    v1 := []string{(("\n\t" + "a") + (("foo bar" + "\n\t") + "foo bar")), "hello"}
    var v2 uint8 = 2
    v3 := []uint32{1, 397, 65535}
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 float64 = -2.25
        v6 := GenS1{F0: 100}
        _ = fi4
        _ = v5
        _ = v6
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v7 int16 = 0
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 float32 = -2.25
        fmt.Println(v2, v9)
        switch (("foo bar" + ("\n\t" + "a")) + "\n\t") {
        case "foo bar":
            if true {
                if (fi8 == 4294967295) {
                    fmt.Println(v9, v9)
                }
                v7 = v7
                var v10 int32 = -2147483648
                _ = v10
            }
        case "":
            v11 := GenS2{F0: 0, F1: 0, F2: 4294967295}
            _ = v11
        default:
            var v12 float64 = 0.1
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    var v13 int = -2147483648
    if len(v3) > 0 {
        v3[0] = 1
    }
    if true {
        var v14 int8 = 42
        _ = v14
    } else {
        var v15 uint16 = 42
        v2 = (v2 * (v2 | (^(v2))))
        _ = v15
    }
    v2 = v2
    v13 -= int(v2)
    v2 = 2
    var v16 float32 = (1.5 + (((0.0 - 0.0) / 1.0) + (0.0 * 3.14159)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
