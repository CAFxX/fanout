package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 bool }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int16 }
func genFn0(p0 uint8, p1 uint) uint {
    if false {
        p0 += p0
        var v5 uint16 = 32767
        _ = v5
    } else {
        var v6 bool = false
        _ = v6
    }
    if true {
        var v7 uint16 = 883
        switch 0 {
        case 0:
            for fi8 := 0; fi8 < 9; fi8++ {
                for fi9 := 0; fi9 < 15; fi9++ {
                    var v10 int64 = 9223372036854775807
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
        default:
            var v11 uint8 = 127
            _ = v11
        }
        v12 := GenS1{F0: -366, F1: 9223372036854775807, F2: false}
        _ = v12
        _ = v7
    }
    switch "a" {
    case "":
        p1 -= (^(p1))
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := []int{(^(fi13))}
            var v15 int16 = 252
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    _ = p0
    _ = p1
    return p1
}
func genFn1(p0 int64) string {
    var v16 int32 = 0
    fmt.Println(p0)
    if false {
        fmt.Println(v16, p0)
        v16 = (^(v16))
        var v17 float32 = -2.25
        _ = v17
    } else {
        v18 := []int{128}
        if false {
            switch "\n\t" {
            case "":
                switch ("a" + "hello") {
                case "hello":
                    v19 := GenS1{F0: -60, F1: p0, F2: false}
                    _ = v19
                case "a":
                    fmt.Println(p0, p0)
                }
            default:
                p0 = -9223372036854775808
            }
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            var v20 int16 = 0
            _ = v20
        } else {
            p0 = 0
            var v21 int16 = -1
            var v22 int64 = (p0 % ((p0) | 1))
            _ = v21
            _ = v22
        }
        if (100 > 0) {
            var v23 uint = genFn0(1, 65535)
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = v23
        }
        _ = v18
    }
    _ = p0
    _ = v16
    return "\n\t"
}

func main() {
    var v1 int = 2147483647
    var v2 string = "foo bar"
    var v3 int16 = -819
    v4 := make(map[int]bool)
    v24 := GenS2{F0: 443, F1: 255}
    v25 := GenS3{F0: 1}
    v26 := GenS3{F0: v3}
    var v27 int64 = 837
    var v28 int8 = -82
    v1 = v1
    v29 := GenS4{F0: v2, F1: 1000}
    v1 = int(v3)
    v27 = v27
    if true {
        for fi30 := 0; fi30 < 1; fi30++ {
            var v31 uint64 = 18446744073709551615
            _ = fi30
            _ = v31
        }
        fmt.Println(v28, v27, v1)
        v4[v1] = true
    }
    fmt.Println(v1, v3)
    v32 := []int{v1, (v1 * v1), (-(v1)), v1}
    var v33 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
