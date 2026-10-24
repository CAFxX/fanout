package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0() uint8 {
    v5 := GenS1{F0: 1}
    var v6 int32 = 0
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int8) bool {
    v7 := GenS2{F0: p0}
    var v8 uint = 945
    switch "hello" {
    case "":
        var v9 uint32 = 584
        _ = v9
    default:
        v10 := GenS2{F0: (^(p0))}
        _ = v10
    }
    var v11 uint64 = 2
    if (p0 <= 127) {
        v12 := []int{-9223372036854775808, 2147483647, 0}
        _ = v12
    }
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    return (!((p0 ^ p0) < (p0 % ((p0) | 1))))
}

func main() {
    var v1 int8 = -1
    var v2 uint16 = 282
    var v3 uint = 1
    v4 := make(map[uint32]bool)
    var v13 int16 = 1
    if ((^(v1)) == -76) {
        _ = len(v4)
        v14 := []int{0}
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        if false {
            switch "\n\t" {
            case "":
                for fi16 := 0; fi16 < 19; fi16++ {
                    v1 = (v1 | v1)
                    v17 := GenS2{F0: v1}
                    if (-1 < 9223372036854775807) {
                        v4[1] = true
                    }
                    _ = fi16
                    _ = v17
                }
            default:
                v18 := GenS2{F0: (v1 * int8(v3))}
                _ = v18
            }
        }
        fmt.Println(v13, fi15, v1)
        delete(v4, 300)
        _ = fi15
    }
    switch ("foo bar" + "hello") {
    case "a":
        v3 -= (v3 / (((v3 % ((v3) | 1))) | 1))
    }
    var v19 float32 = 3.14159
    switch 684 {
    case -301:
        var v20 int16 = -32768
        _ = v20
    default:
        delete(v4, 42)
    }
    switch 1 {
    case 2147483647:
        var v21 uint32 = 4294967295
        _ = v21
    case -2147483648:
        v3 = (v3 * v3)
    }
    v22 := GenS3{F0: -1, F1: v1}
    v13 = (v13 + 1000)
    v2 = 42
    var v23 float64 = 0.1
    var v24 int = -569
    v1 += 1
    fmt.Println(v24, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
