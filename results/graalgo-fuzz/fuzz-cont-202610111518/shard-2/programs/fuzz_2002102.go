package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 uint64 = 42
    var v4 int16 = 1000
    var v5 float32 = 1e10
    delete(v2, "\n\t")
    if false {
        for fi6 := 0; fi6 < 2; fi6++ {
            fmt.Println(fi6)
            v7 := GenS1{F0: fi6, F1: 42, F2: (-(fi6))}
            _ = fi6
            _ = v7
        }
        v5 = float32(((0.0 / float64(1.5)) - (0.0 + 0.0)))
    } else {
        v8 := GenS1{F0: 42, F1: 2147483647, F2: -2147483648}
        _ = v8
    }
    switch "foo bar" {
    case "\n\t":
        var v9 int = 42
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(fi10, v3)
            var v11 uint = 2
            var v12 float32 = (float32(3.14159) + (3.14159 - v5))
            _ = fi10
            _ = v11
            _ = v12
        }
    case "a":
        fmt.Println(v4)
    default:
        var v13 string = "foo bar"
        _ = v13
    }
    v3 = (v3 ^ v3)
    if (255 >= 65535) {
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS1{F0: -135, F1: 1, F2: (fi14 % ((fi14) | 1))}
            var v16 uint = 18446744073709551615
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        fmt.Println(v5, v3)
        if (404 <= (v3 << ((2) & 63))) {
            var v17 bool = (("" > ("hello" + "")) || ((v5 < (0.1 + float32(-2.25))) || (42 != 1)))
            _ = v17
        } else {
            v18 := GenS1{F0: 42, F1: -1, F2: 127}
            var v19 uint64 = (^(v3))
            _ = v18
            _ = v19
        }
    }
    var v20 float32 = -2.25
    if ((v3 / ((1) | 1)) < (v3 / ((935) | 1))) {
        switch 1 {
        case 42:
            v3 = (^(v3))
        case -746:
            var v21 uint = 18446744073709551615
            _ = v21
        }
    } else {
        if (0 != 65535) {
            v5 -= (v20 - 3.14159)
            _ = len(v1)
        }
    }
    v22 := GenS2{F0: (4294967295 != 375), F1: 32767}
    fmt.Println(v5)
    v23 := GenS1{F0: 772, F1: 0, F2: -2147483648}
    fmt.Println(v3, v20, v20)
    var v24 uint32 = 2147483647
    var v25 int8 = 42
    for fi26 := 0; fi26 < 11; fi26++ {
        switch (fi26 % ((fi26) | 1)) {
        case 9223372036854775807:
            var v27 int = 128
            _ = v27
        default:
            var v28 int16 = (v4 >> ((160) & 63))
            _ = v28
        }
        var v29 int32 = 42
        _ = fi26
        _ = v29
    }
    v25 -= 0
    var v30 int = -631
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
