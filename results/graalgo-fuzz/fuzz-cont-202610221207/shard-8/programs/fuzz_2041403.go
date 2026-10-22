package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }

func main() {
    v1 := []int{-1}
    v2 := []uint32{286}
    var v3 int8 = 42
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := GenS1{F0: 219, F1: 42}
        v8 := GenS1{F0: 4294967295, F1: 359}
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float32 = 0.0
        v3 = 1
        _ = fi9
        _ = v10
    }
    v3 -= (v3 * (v3 % (((v3 - v3)) | 1)))
    if true {
        fmt.Println(v3)
    } else {
        if (0.1 == ((3.14159 + float64(3.14159)) / 1.0)) {
            v3 -= (v3 & (v3 / ((42) | 1)))
        }
        switch 913 {
        case 4294967295:
            v3 = (v3 % (((v3 ^ v3)) | 1))
        case 9223372036854775807:
            _ = len(v1)
        }
        v3 = v3
    }
    v11 := GenS1{F0: 2, F1: 0}
    var v12 uint16 = 42
    var v13 uint16 = (v12 / ((32767) | 1))
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 float64 = ((1e10 + (0.1 / 1.5)) * (0.1 - 1e10))
        v12 -= (v12 - (v12 ^ (v13 * (v13 % (((v13 * 803)) | 1)))))
        _ = fi14
        _ = v15
    }
    switch 42 {
    case 2147483647:
        if true {
            var v16 int8 = v3
            v17 := GenS1{F0: 834, F1: -1}
            _ = v16
            _ = v17
        }
    case 32767:
        switch ("\n\t" + ((((("foo bar" + "hello") + "a") + "a") + "") + "foo bar")) {
        case "foo bar":
            v18 := GenS1{F0: 244, F1: 42}
            _ = v18
        default:
            fmt.Println(v12, v13)
        }
    }
    v12 -= v12
    switch 9223372036854775807 {
    case -2147483648:
        v1 = append(v1, -2147483648)
    default:
        v1 = append(v1, 32767)
    }
    var v19 uint64 = 1
    var v20 uint64 = (^(v19))
    v19 += v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
