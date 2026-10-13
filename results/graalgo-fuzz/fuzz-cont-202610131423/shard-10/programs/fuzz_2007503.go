package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64; F2 int64 }
func genFn0(p0 uint8) string {
    v4 := GenS1{F0: 0, F1: true}
    fmt.Println(p0)
    v5 := []int{-9223372036854775808, 127}
    v6 := GenS2{F0: 255, F1: 1, F2: 42}
    p0 = p0
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ""
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1e10
    v3 := []int{4294967295}
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v2, v2)
        switch (-(fi7)) {
        case 427:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case -952:
            v1[("a" + "hello")] = "hello"
        case 9223372036854775807:
            delete(v1, "\n\t")
        default:
            v2 = v2
        }
        _ = fi7
    }
    v2 = v2
    var v8 int32 = 42
    var v9 uint32 = 476
    v9 = (^(v9))
    v10 := GenS1{F0: 65535, F1: true}
    switch (("foo bar" + "a") + ("\n\t" + ("hello" + ("\n\t" + "\n\t")))) {
    case "\n\t":
        fmt.Println(v8, v9)
    default:
        if true {
            v11 := GenS2{F0: 1, F1: 4294967295, F2: -1}
            var v12 float64 = 1.5
            v1["foo bar"] = "\n\t"
            _ = v11
            _ = v12
        } else {
            for fi13 := 0; fi13 < 8; fi13++ {
                v2 -= v2
                v8 -= 667
                _ = fi13
            }
        }
    }
    v9 = (v9 * 42)
    for fi14 := 0; fi14 < 5; fi14++ {
        v2 = (v2 / v2)
        v2 = (float32((1.5 + 1e10)) - 0.0)
        _ = fi14
    }
    if false {
        fmt.Println(v8, v8, v8)
    } else {
        var v15 float32 = v2
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
