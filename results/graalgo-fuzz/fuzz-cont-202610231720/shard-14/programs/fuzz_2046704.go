package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{432, 266, 1}
    var v3 int16 = 32767
    for fi4 := 0; fi4 < 14; fi4++ {
        v5 := GenS1{F0: (fi4 | (^(fi4))), F1: 0, F2: fi4}
        _ = fi4
        _ = v5
    }
    if (true || true) {
        v3 = (v3 << ((794) & 63))
    } else {
        v1 += v1
        var v6 int = 32767
        _ = v6
    }
    v7 := GenS1{F0: -2147483648, F1: 2147483647, F2: 128}
    var v8 uint16 = 2
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := GenS1{F0: -9223372036854775808, F1: (v9 - (v9 ^ (v9 ^ (v9 * 2147483647)))), F2: (fi10 ^ fi11)}
            _ = fi11
            _ = v12
        }
        var v13 float32 = float32(((0.0 + 1.5) * (float64(-2.25) + (0.0 * float64(1.5)))))
        _ = fi10
        _ = v13
    }
    if true {
        var v14 int = -2147483648
        _ = v14
    } else {
        v9 = (v9 * (^(v9)))
        if len(v2) > 0 {
            v2[0] = (^(v9))
        }
    }
    v15 := GenS1{F0: 128, F1: v9, F2: 4294967295}
    var v16 float32 = 1e10
    switch "a" {
    case "a":
        var v17 float64 = (0.1 + (((((1.5 / -2.25) + 0.0) + 0.1) - 1e10) / 3.14159))
        _ = v17
    case "hello":
        v1 += 0
    case "foo bar":
        var v18 uint8 = 127
        _ = v18
    }
    v9 = (^(v9))
    v1 -= (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
