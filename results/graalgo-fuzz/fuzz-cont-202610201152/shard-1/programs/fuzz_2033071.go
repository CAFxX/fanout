package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
func genFn0(p0 int64, p1 int32, p2 int8) int64 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint = 255
        var v8 uint32 = 2147483648
        for fi9 := 0; fi9 < 18; fi9++ {
            p2 = p2
            fmt.Println(p0, v8)
            _ = fi9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch -2147483648 {
    case 42:
        fmt.Println(p0, p0, p2)
    default:
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(p2, p0)
            var v11 float32 = (float32(float64((float32(float64(float32(0.0))) * 0.1))) / (1.5 * (0.1 - (3.14159 * 3.14159))))
            _ = fi10
            _ = v11
        }
    }
    var v12 int32 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return (p0 >> ((0) & 63))
}
func genFn1() int64 {
    var v13 uint32 = 2147483647
    var v14 uint16 = 0
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("a" + ""), (("a" + "a") + "foo bar"), "\n\t", ""}
    var v3 float64 = 1e10
    var v4 uint = 42
    var v5 uint = v4
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(fi15)
        fmt.Println(fi15)
        v4 += (v4 | v4)
        _ = fi15
    }
    if true {
        v5 -= v4
        var v16 uint = 1
        _ = v16
    } else {
        var v17 float32 = float32((v3 * -2.25))
        _ = v17
    }
    v1[691] = -1
    fmt.Println(v3, v3, v3)
    if (-1 <= -611) {
        var v18 uint16 = 42
        _ = v18
    }
    var v19 uint64 = 18446744073709551615
    for fi20 := 0; fi20 < 11; fi20++ {
        var v21 float32 = 1.5
        v2 = append(v2, (("\n\t" + "") + ("" + "\n\t")))
        var v22 float64 = v3
        _ = fi20
        _ = v21
        _ = v22
    }
    v5 = (^(v5))
    if (!(1 > -128)) {
        v23 := GenS1{F0: 42, F1: -32768, F2: (!true)}
        if false {
            v24 := GenS1{F0: 318, F1: 42, F2: (!false)}
            _ = v24
        } else {
            v25 := GenS1{F0: -60, F1: -32768, F2: (float64(0.0) != float64(0.0))}
            _ = v25
        }
        v3 += (3.14159 * 0.1)
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
