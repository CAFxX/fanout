package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint32; F2 int }
type GenS3 struct { F0 int; F1 bool }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{fi7, fi7}
        fmt.Println(fi7)
        var v9 uint16 = 255
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{(fi10 ^ fi10), 1}
        switch "\n\t" {
        case "":
            v12 := GenS1{F0: 4294967295}
            _ = v12
        case "\n\t":
            fmt.Println(fi10)
        default:
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 float64 = 0.1
            v15 := GenS2{F0: 1, F1: 1, F2: 42}
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v16 := []int{-9223372036854775808, -694, 42, 0}
    _ = v16
    return 0
}

func main() {
    var v1 int32 = 575
    var v2 int32 = (v1 % ((v1) | 1))
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 string = "foo bar"
    v6 := make(map[string]uint32)
    var v17 uint = 255
    var v18 int64 = 42
    var v19 string = v5
    switch 128 {
    case -273:
        v20 := GenS3{F0: -475, F1: (!true)}
        _ = v20
    default:
        fmt.Println(v17)
    }
    if (true && (1 != 0)) {
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint64 = 18446744073709551615
            var v23 string = (("a" + v19) + v19)
            _ = fi21
            _ = v22
            _ = v23
        }
        v18 -= v18
        _ = len(v3)
    } else {
        for fi24 := 0; fi24 < 7; fi24++ {
            fmt.Println(v2, v1)
            _ = fi24
        }
    }
    fmt.Println(v17, v19)
    switch v19 {
    case "\n\t":
        var v25 float64 = ((1e10 - (((1e10 - 1.5) - 3.14159) + 1.5)) / 1e10)
        _ = v25
    case "":
        v5 = (v19 + v19)
    default:
        switch v19 {
        case "foo bar":
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 int8 = 100
                _ = fi26
                _ = v27
            }
        case "a":
            v28 := GenS3{F0: 32767, F1: (4294967295 < 883)}
            _ = v28
        default:
            v19 = v5
        }
    }
    var v29 int32 = (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
