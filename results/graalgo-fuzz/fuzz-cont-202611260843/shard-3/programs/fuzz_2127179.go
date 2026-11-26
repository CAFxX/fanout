package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int32; F2 int32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 string) bool {
    switch p0 {
    case "hello":
        v6 := []int{2147483647, 2147483647}
        _ = v6
    case "a":
        switch 42 {
        case 32768:
            var v7 uint16 = 42
            _ = v7
        case -47:
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: 47, F1: true}
                if (-32768 <= 350) {
                    fmt.Println(p0, fi8)
                    if true {
                        p0 = p0
                        v10 := []int{}
                        _ = v10
                    }
                }
                _ = fi8
                _ = v9
            }
        }
    case "":
        v11 := GenS2{F0: (-9223372036854775808 > -826), F1: (1 < 42)}
        _ = v11
    default:
        v12 := GenS1{F0: 0, F1: false}
        _ = v12
    }
    if (128 == 128) {
        switch 128 {
        case 127:
            p0 = "hello"
        case 888:
            var v13 int32 = 2147483647
            _ = v13
        default:
            v14 := GenS1{F0: 0, F1: (false && true)}
            _ = v14
        }
        if true {
            var v15 float64 = 0.1
            _ = v15
        } else {
            p0 = p0
            p0 = p0
        }
    } else {
        v16 := []int{}
        _ = v16
    }
    _ = p0
    return false
}
func genFn1(p0 uint16, p1 int64) int64 {
    var v17 int8 = 0
    v18 := GenS3{F0: -1, F1: 0}
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "hello"
    var v3 string = ((v2 + v2) + ((((v2 + v2) + "foo bar") + v2) + "\n\t"))
    var v4 uint64 = 126
    var v5 int64 = 1
    switch v3 {
    case "hello":
        fmt.Println(v2, v5)
    case "\n\t":
        var v19 int8 = 122
        _ = v19
    default:
        v5 = 9223372036854775807
    }
    var v20 float32 = 1.5
    var v21 bool = ((int32(v5) != 42) && ((true && (true || true)) || (!false)))
    v21 = false
    switch "\n\t" {
    case "\n\t":
        v22 := GenS4{F0: -32768}
        _ = v22
    case "":
        if v21 {
            v20 -= v20
        }
    case "a":
        var v23 float32 = ((float32(1e10) + v20) - v20)
        _ = v23
    default:
        fmt.Println(v5, v20)
    }
    var v24 float64 = 0.1
    v25 := GenS5{F0: 2147483647, F1: int32(v5), F2: 1}
    if (!(!v21)) {
        if v21 {
            v26 := GenS2{F0: (127 <= 100), F1: (v4 == v4)}
            v27 := GenS6{F0: 127}
            v5 -= (-(v5))
            _ = v26
            _ = v27
        }
        switch 2147483647 {
        case -2147483648:
            for fi28 := 0; fi28 < 2; fi28++ {
                fmt.Println(v2)
                _ = fi28
            }
        case 348:
            fmt.Println(v20, v21, v24)
        default:
            switch 0 {
            case -1:
                for fi29 := 0; fi29 < 7; fi29++ {
                    v30 := GenS3{F0: 537, F1: 13}
                    _ = fi29
                    _ = v30
                }
            default:
                v31 := GenS2{F0: true, F1: v21}
                _ = v31
            }
        }
        var v32 string = v3
        _ = v32
    } else {
        fmt.Println(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
