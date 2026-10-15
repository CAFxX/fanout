package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0() uint32 {
    var v7 float32 = 0.1
    if (-371 <= 1000) {
        var v8 uint = 255
        fmt.Println(v8)
        v9 := GenS1{F0: 45, F1: 1}
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "":
        v7 = (v7 / 1e10)
    case "hello":
        var v10 uint16 = 42
        _ = v10
    case "foo bar":
        v11 := []int{-2147483648, 814, 1, 0}
        _ = v11
    }
    fmt.Println(v7)
    _ = v7
    return 2147483648
}
func genFn1(p0 int, p1 uint64, p2 uint16) uint16 {
    if false {
        for fi12 := 0; fi12 < 13; fi12++ {
            if true {
                p2 = p2
            } else {
                v13 := GenS1{F0: 1, F1: (p2 >> ((18446744073709551615) & 63))}
                var v14 uint = 65535
                var v15 int64 = -1
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    }
    v16 := GenS1{F0: 1, F1: p2}
    var v17 int = 32768
    if (p2 >= p2) {
        if false {
            v18 := []int{}
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(v17)
        v20 := GenS2{F0: 255, F1: 255}
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 float64 = (((1e10 + 1.5) * float64(0.1)) * 1.5)
    var v2 uint8 = 42
    v3 := []uint32{}
    var v4 int16 = 42
    v5 := []int{}
    v6 := make(map[string]string)
    v1 = v1
    v21 := GenS1{F0: (^(v2)), F1: 1}
    v22 := GenS1{F0: v2, F1: 2}
    if false {
        if true {
            v23 := GenS1{F0: (v2 * 127), F1: 709}
            v24 := GenS3{F0: 86, F1: 100}
            v2 = v2
            _ = v23
            _ = v24
        }
        v25 := GenS3{F0: 58, F1: -128}
        v4 = (v4 ^ (v4 ^ v4))
        _ = v25
    } else {
        var v26 int32 = -838
        v4 += v4
        var v27 uint64 = 527
        _ = v26
        _ = v27
    }
    var v28 int64 = 9223372036854775807
    var v29 uint8 = v2
    v28 += 9223372036854775807
    v1 = float64(-2.25)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1, v1, v29)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
