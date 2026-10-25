package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0(p0 uint32, p1 uint16, p2 bool) string {
    var v4 float64 = 1e10
    if (!true) {
        v4 -= v4
        v5 := GenS1{F0: -2147483648, F1: -1}
        _ = v5
    }
    switch 1 {
    case 2147483647:
        var v6 int64 = 466
        _ = v6
    case 42:
        var v7 uint64 = 4294967295
        _ = v7
    default:
        v8 := GenS1{F0: -1, F1: -128}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ((("foo bar" + "foo bar") + (((("foo bar" + ("hello" + "\n\t")) + ("hello" + ("a" + ("a" + "a")))) + "") + "")) + "")
}
func genFn1() bool {
    var v9 int32 = 42
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 int32 = (v9 | v9)
        _ = fi10
        _ = v11
    }
    fmt.Println(v9)
    _ = v9
    return (!false)
}

func main() {
    var v1 float32 = (((3.14159 / 1.0) - (0.0 - -2.25)) - 3.14159)
    v2 := []int{4294967295}
    var v3 uint64 = 459
    for fi12 := 0; fi12 < 20; fi12++ {
        v1 = v1
        v3 -= v3
        _ = fi12
    }
    var v13 uint16 = 449
    v14 := GenS2{F0: 4294967295}
    v1 -= v1
    v13 += (v13 - (v13 - (v13 + v13)))
    v13 = 103
    for fi15 := 0; fi15 < 4; fi15++ {
        switch 2147483647 {
        case 128:
            v3 = v3
        case 4294967295:
            v3 += (v3 * (^(v3)))
        case 110:
            fmt.Println(v3, v1, fi15)
        default:
            v2 = append(v2, fi15)
        }
        v16 := GenS3{F0: 321, F1: 109, F2: "a"}
        fmt.Println(v1, v3)
        _ = fi15
        _ = v16
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v13 = v13
        var v18 int32 = -1
        _ = fi17
        _ = v18
    }
    var v19 uint64 = (v3 & (v3 ^ (^(v3))))
    v20 := GenS4{F0: -2147483648, F1: -1}
    v1 = (float32(1e10) + float32((0.1 * 1.5)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
