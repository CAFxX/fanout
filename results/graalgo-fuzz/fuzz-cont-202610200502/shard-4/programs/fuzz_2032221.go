package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 int) uint {
    v5 := GenS1{F0: "", F1: -9223372036854775808}
    for fi6 := 0; fi6 < 6; fi6++ {
        p0 -= p0
        _ = fi6
    }
    _ = p0
    _ = v5
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 bool = ((-128 <= 42) || true)
    var v4 bool = (0 == 0)
    fmt.Println(v3, v3)
    switch "\n\t" {
    case "foo bar":
        var v7 uint32 = 222
        _ = v7
    default:
        var v8 uint32 = 2147483647
        _ = v8
    }
    v9 := GenS1{F0: "hello", F1: -143}
    if (v4 && v4) {
        var v10 int = -623
        _ = len(v2)
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v4, v3)
        _ = fi11
    }
    v12 := GenS1{F0: "", F1: 42}
    if (0.0 != (1e10 + (float32(3.14159) / (0.0 - (3.14159 / 1.0))))) {
        var v13 float32 = ((0.0 * 0.0) + 0.1)
        _ = v13
    }
    var v14 float64 = 1e10
    v4 = (!(!true))
    if (950 <= 1) {
        v15 := []int{}
        _ = v15
    } else {
        _ = len(v2)
        if (2 > 895) {
            fmt.Println(v3, v14)
        } else {
            if (true && (42 != 42)) {
                v14 = 1.5
                v4 = true
                var v16 bool = ((false && v4) && ((-1 < 1) && v4))
                _ = v16
            }
            v17 := GenS1{F0: "foo bar", F1: 357}
            _ = v17
        }
        v18 := GenS2{F0: 128, F1: ""}
        _ = v18
    }
    var v19 int16 = 695
    var v20 int64 = 0
    switch 1 {
    case 9223372036854775807:
        if ((v4 || (0 != (v19 << ((genFn0(2147483647)) & 63)))) && (!(-294 <= v20))) {
            for fi21 := 0; fi21 < 6; fi21++ {
                fmt.Println(fi21, v3)
                _ = fi21
            }
            v22 := []int{-355}
            var v23 uint16 = 1
            _ = v22
            _ = v23
        }
    case 128:
        if (0.0 == (-2.25 + (1.5 - 3.14159))) {
            var v24 int16 = int16(v20)
            fmt.Println(v14, v20, v3)
            _ = v24
        }
    }
    v25 := []int{}
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint8 = 0
        _ = fi26
        _ = v27
    }
    if (218 <= 640) {
        var v28 int = 2147483647
        var v29 uint = 42
        _ = v28
        _ = v29
    } else {
        switch ("a" + "hello") {
        case "hello":
            v4 = ((float32(v14) / 3.14159) > 0.1)
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            if len(v25) > 0 {
                v25[0] = fi30
            }
            _ = fi30
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v9
}
