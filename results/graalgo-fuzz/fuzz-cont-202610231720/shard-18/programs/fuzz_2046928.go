package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0() int64 {
    var v7 string = "hello"
    v8 := GenS1{F0: 69}
    _ = v7
    _ = v8
    return -932
}
func genFn1(p0 uint) int {
    fmt.Println(p0)
    var v9 uint8 = 249
    if (v9 >= (^(v9))) {
        if false {
            var v10 float64 = 1e10
            v9 = v9
            fmt.Println(v9)
            _ = v10
        }
    }
    v11 := GenS1{F0: -32768}
    _ = p0
    _ = v11
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]bool)
    var v3 int = 32768
    v4 := []int{}
    var v5 uint = 0
    var v6 uint32 = 615
    switch 0 {
    case -42:
        v5 += 882
    case 9223372036854775807:
        if (42 == -2147483648) {
            v12 := GenS1{F0: 42}
            _ = v12
        }
    case 2147483647:
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := GenS2{F0: false, F1: v1}
            v6 += (v6 << ((v5) & 63))
            _ = fi13
            _ = v14
        }
    }
    v1 -= v1
    v15 := GenS3{F0: genFn0(), F1: 1, F2: 782}
    v16 := GenS3{F0: 1, F1: -2147483648, F2: 961}
    var v17 string = "hello"
    var v18 uint16 = 0
    var v19 uint8 = v1
    if false {
        if (float32(0.1) < 0.1) {
            for fi20 := 0; fi20 < 18; fi20++ {
                v3 = 9223372036854775807
                var v21 float64 = (-2.25 * float64(0.0))
                _ = fi20
                _ = v21
            }
            switch genFn1(v5) {
            case -1:
                var v22 int32 = 1
                _ = v22
            case -127:
                if len(v4) > 0 {
                    v4[0] = v3
                }
            default:
                var v23 bool = false
                _ = v23
            }
            var v24 float64 = 1.5
            _ = v24
        } else {
            _ = len(v2)
            v25 := GenS2{F0: true, F1: 0}
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 9; fi26++ {
            var v27 int = v3
            for fi28 := 0; fi28 < 20; fi28++ {
                v29 := GenS2{F0: true, F1: (v19 | (v19 ^ (^(v19))))}
                _ = fi28
                _ = v29
            }
            _ = fi26
            _ = v27
        }
        if len(v4) > 0 {
            v4[0] = v3
        }
    }
    v1 += 127
    v30 := GenS1{F0: 42}
    v31 := GenS3{F0: 0, F1: -87, F2: 881}
    for fi32 := 0; fi32 < 6; fi32++ {
        v5 = (^(v5))
        _ = fi32
    }
    v33 := GenS1{F0: -984}
    var v34 int64 = genFn0()
    var v35 uint16 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
