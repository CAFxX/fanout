package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
func genFn0(p0 int32) int16 {
    p0 = p0
    p0 = (-(p0))
    var v6 uint32 = 702
    _ = p0
    _ = v6
    return -405
}
func genFn1(p0 uint8) int32 {
    var v7 uint32 = 0
    v8 := []int{365, -1}
    switch "hello" {
    case "foo bar":
        v9 := GenS1{F0: 4294967295, F1: v7, F2: -1}
        _ = v9
    case "\n\t":
        v7 -= 238
    }
    _ = p0
    _ = v7
    _ = v8
    return 643
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{42, 42}
    v3 := make(map[uint32]string)
    var v4 uint = 2
    var v5 uint32 = 0
    v5 -= 2147483648
    var v10 int32 = genFn1(128)
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            fmt.Println(v1)
            fmt.Println(v1, v10, v1)
            _ = fi12
        }
        _ = fi11
    }
    v2 = append(v2, v5)
    switch 32767 {
    case 602:
        switch -582 {
        case -129:
            v13 := GenS1{F0: -877, F1: 2, F2: -64}
            _ = v13
        case 369:
            v4 -= v4
        default:
            var v14 float64 = 0.0
            _ = v14
        }
    case -9223372036854775808:
        v15 := GenS1{F0: -1, F1: v5, F2: 27}
        _ = v15
    }
    switch 228 {
    case -227:
        var v16 float32 = ((3.14159 / 1.5) / 1.5)
        _ = v16
    case 32768:
        fmt.Println(v1)
    case 2147483647:
        if true {
            var v17 uint16 = 32767
            fmt.Println(v1, v1, v10)
            switch 42 {
            case 128:
                _ = len(v2)
            }
            _ = v17
        }
    default:
        var v18 uint32 = v5
        _ = v18
    }
    v19 := GenS1{F0: 0, F1: 65535, F2: -1}
    var v20 float32 = 1e10
    v3[2147483647] = ("a" + "\n\t")
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS1{F0: 2147483647, F1: (v5 * v5), F2: -18}
        _ = fi21
        _ = v22
    }
    var v23 int32 = (v10 + v10)
    v24 := GenS1{F0: 70, F1: (v5 >> ((4294967295) & 63)), F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
