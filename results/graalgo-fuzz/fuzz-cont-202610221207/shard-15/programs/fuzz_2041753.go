package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }
type GenS4 struct { F0 uint8 }
func genFn0(p0 uint16, p1 uint32, p2 int32) uint32 {
    v4 := []int{}
    var v5 uint16 = (p0 >> ((2) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int64 = 1
    v7 := GenS1{F0: 1000, F1: 541, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (^(p1))
}
func genFn1(p0 int) int32 {
    v8 := []int{}
    fmt.Println(p0)
    var v9 uint = 18446744073709551615
    switch p0 {
    case -2147483648:
        p0 = p0
    }
    _ = p0
    _ = v8
    _ = v9
    return 254
}

func main() {
    var v1 int64 = 42
    var v2 uint32 = 799
    var v3 uint = 4294967295
    var v10 uint64 = 229
    v11 := []int{127, 9223372036854775807, 562}
    v12 := GenS2{F0: -751, F1: 788}
    v10 -= 1
    fmt.Println(v10, v2)
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS3{F0: 42, F1: uint32(v3), F2: false}
        var v15 int8 = 127
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS4{F0: 128}
    var v17 float32 = 3.14159
    v18 := GenS3{F0: 1, F1: (^(v2)), F2: ((239 != 242) || true)}
    v2 += v2
    var v19 uint8 = 71
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 float32 = float32(1.5)
        var v22 int = fi20
        _ = fi20
        _ = v21
        _ = v22
    }
    switch "" {
    case "a":
        switch "" {
        case "a":
            var v23 int = 0
            _ = v23
        default:
            var v24 int8 = 100
            _ = v24
        }
    case "":
        if false {
            if (true && false) {
                fmt.Println(v19, v2)
                switch "\n\t" {
                case "foo bar":
                    v25 := GenS1{F0: 32767, F1: v10, F2: 450}
                    _ = v25
                default:
                    var v26 int8 = 0
                    _ = v26
                }
            }
        }
    }
    var v27 uint32 = 542
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
}
