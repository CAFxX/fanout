package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int; F1 int }
func genFn0() uint64 {
    var v6 float64 = (float64((-2.25 - -2.25)) + 1.5)
    switch 42 {
    case 32768:
        var v7 int16 = -1
        _ = v7
    default:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint16 = 2
            switch fi8 {
            case 32768:
                var v10 int = fi8
                _ = v10
            case 47:
                v11 := GenS1{F0: 42, F1: 0}
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    _ = v6
    return 0
}
func genFn1(p0 string, p1 int, p2 uint8) int {
    v12 := GenS2{F0: -1, F1: 0, F2: 0}
    p0 = p0
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = 9223372036854775807
        var v15 int8 = 109
        var v16 int64 = 0
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 int8 = 42
    var v18 int64 = -809
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v17
    _ = v18
    return -301
}

func main() {
    var v1 int64 = 42
    var v2 int8 = -82
    v3 := make(map[uint32]string)
    var v4 uint16 = 2
    v5 := make(map[uint32]bool)
    var v19 float64 = 0.1
    if false {
        if true {
            switch 9223372036854775807 {
            case -9223372036854775808:
                for fi20 := 0; fi20 < 14; fi20++ {
                    var v21 int16 = 614
                    if (((-2147483648 > 128) || (false && false)) && false) {
                        if false {
                            v22 := GenS3{F0: 2147483647, F1: 84}
                            v4 = v4
                            _ = v22
                        } else {
                            var v23 float32 = -2.25
                            _ = v23
                        }
                    }
                    v24 := []int{(fi20 / ((-9223372036854775808) | 1)), 127, 32768, fi20}
                    _ = fi20
                    _ = v21
                    _ = v24
                }
            default:
                v25 := GenS4{F0: 127, F1: 42}
                _ = v25
            }
        }
        v26 := GenS1{F0: 4294967295, F1: 1}
        var v27 uint32 = 65535
        _ = v26
        _ = v27
    }
    var v28 bool = (false || (!false))
    v4 = v4
    if true {
        v4 = (v4 & 2)
        for fi29 := 0; fi29 < 7; fi29++ {
            v19 = 0.1
            _ = fi29
        }
        var v30 string = "a"
        _ = v30
    }
    v1 = v1
    var v31 uint8 = 165
    var v32 int16 = 68
    v4 -= (v4 ^ (v4 / ((v4) | 1)))
    _ = len(v5)
    v1 = v1
    var v33 int16 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
