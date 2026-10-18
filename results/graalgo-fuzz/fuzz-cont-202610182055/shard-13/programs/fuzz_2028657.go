package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint64, p1 int16) int {
    fmt.Println(p0, p1)
    var v6 string = "\n\t"
    switch 32767 {
    case 0:
        v7 := GenS1{F0: 42, F1: 202, F2: 996}
        _ = v7
    case 32767:
        switch 1 {
        case 924:
            if (((false || (-557 == 127)) || (((136 < 412) && (true || (!(18446744073709551615 > 255)))) || false)) || false) {
                for fi8 := 0; fi8 < 10; fi8++ {
                    p0 += p0
                    _ = fi8
                }
                var v9 uint8 = 128
                for fi10 := 0; fi10 < 5; fi10++ {
                    var v11 uint64 = 18446744073709551615
                    var v12 int64 = 0
                    v13 := []int{(^(fi10)), fi10, fi10, (fi10 | (fi10 - fi10))}
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
                _ = v9
            } else {
                var v14 int64 = 9223372036854775807
                _ = v14
            }
        }
    case 2147483647:
        v6 = (v6 + "foo bar")
    default:
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 int = fi15
            v17 := GenS1{F0: 42, F1: 1, F2: int16(p0)}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v18 uint64 = p0
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    v5 := []uint32{4294967295}
    v2 = 0
    v2 = 255
    switch 4294967295 {
    case 42:
        var v19 string = ("" + "")
        _ = v19
    case -95:
        delete(v3, genFn0(18446744073709551615, -913))
    default:
        if false {
            v2 = 42
        }
    }
    fmt.Println(v4, v4)
    v2 -= (v2 & (v2 + (v2 / ((v2) | 1))))
    switch 42 {
    case 0:
        var v20 uint16 = 65535
        _ = v20
    case 267:
        v21 := GenS2{F0: int16(v2), F1: -1}
        _ = v21
    case 868:
        v2 = 127
    }
    if true {
        v22 := GenS3{F0: true, F1: -7}
        _ = v22
    } else {
        fmt.Println(v2)
        fmt.Println(v2, v2)
    }
    delete(v1, "foo bar")
    if false {
        var v23 bool = false
        _ = v23
    }
    v2 += (v2 % ((1) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
