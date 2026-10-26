package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint, p2 bool) uint8 {
    fmt.Println(p0)
    var v4 int32 = (p0 ^ (p0 >> ((1) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2
}
func genFn1(p0 string, p1 uint64) string {
    var v5 int64 = 0
    var v6 int32 = -2147483648
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 int = 0
    v2 := make(map[string]uint32)
    var v3 int64 = 1
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            v3 -= (-(v3))
            v9 := GenS1{F0: (v3 ^ (^(v3))), F1: 2147483647}
            _ = fi8
            _ = v9
        }
        if false {
            var v10 bool = (127 >= 42)
            v10 = true
            var v11 float32 = 3.14159
            _ = v10
            _ = v11
        }
        var v12 int = 1
        _ = fi7
        _ = v12
    }
    var v13 string = "foo bar"
    v13 = v13
    v14 := GenS2{F0: 1}
    v13 = "\n\t"
    v1 = v1
    fmt.Println(v3, v3, v3)
    switch (v13 + ("a" + (v13 + v13))) {
    case "foo bar":
        switch v13 {
        case "foo bar":
            v13 = ((v13 + v13) + (v13 + v13))
        default:
            v13 = v13
        }
    case "\n\t":
        var v15 uint = 82
        _ = v15
    case "":
        var v16 float32 = float32(-2.25)
        _ = v16
    default:
        var v17 string = "hello"
        _ = v17
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
