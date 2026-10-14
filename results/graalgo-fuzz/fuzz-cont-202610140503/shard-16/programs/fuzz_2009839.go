package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 float64 = 1.5
    var v5 uint32 = 2147483647
    var v6 float32 = (0.1 - 0.1)
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 int = fi7
        fmt.Println(fi7, v6)
        for fi9 := 0; fi9 < 17; fi9++ {
            v5 = v5
            v5 -= 2147483647
            var v10 uint64 = 604
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v5 -= 4294967295
    fmt.Println(v4, v5)
    switch -729 {
    case -1:
        v1 = append(v1, 128)
    default:
        v11 := GenS1{F0: -98, F1: 52}
        _ = v11
    }
    if false {
        var v12 bool = true
        if true {
            var v13 uint32 = v5
            var v14 uint16 = 1
            var v15 bool = v12
            _ = v13
            _ = v14
            _ = v15
        }
        switch -2147483648 {
        case -1:
            fmt.Println(v6)
        case 0:
            if (2 != 18446744073709551615) {
                if ((!v12) || (380 >= 1)) {
                    v1 = append(v1, -796)
                    v16 := GenS1{F0: 19, F1: 267}
                    _ = v16
                } else {
                    var v17 uint16 = 255
                    _ = v17
                }
            }
        case -86:
            v5 = (v5 + v5)
        default:
            switch 612 {
            case 127:
                var v18 uint64 = 2
                _ = v18
            case 32767:
                var v19 uint32 = v5
                _ = v19
            }
        }
        _ = v12
    } else {
        var v20 uint16 = 0
        v21 := GenS2{F0: ""}
        var v22 uint = 18446744073709551615
        _ = v20
        _ = v21
        _ = v22
    }
    var v23 int16 = -240
    switch 1 {
    case 9223372036854775807:
        var v24 int64 = 0
        _ = v24
    }
    var v25 uint8 = 235
    var v26 uint = 65535
    for fi27 := 0; fi27 < 20; fi27++ {
        if (((false || (42 == 42)) && true) && true) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi28 := 0; fi28 < 10; fi28++ {
                var v29 float64 = ((v4 / (v4 / v4)) + v4)
                v4 = float64(float32(3.14159))
                _ = fi28
                _ = v29
            }
        }
        _ = fi27
    }
    v4 -= v4
    v30 := GenS3{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
