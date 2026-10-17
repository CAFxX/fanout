package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0() uint16 {
    var v6 int64 = 0
    var v7 bool = (!true)
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8) int8 {
    p0 = p0
    var v8 uint32 = 150
    v8 += 1
    _ = p0
    _ = v8
    return 113
}

func main() {
    var v1 uint32 = 42
    var v2 uint16 = 42
    var v3 int = -517
    var v4 uint16 = (^(v2))
    var v5 float64 = 0.0
    v9 := GenS1{F0: 2147483647}
    if (true && true) {
        var v10 int32 = -2147483648
        if (false && (uint(v4) == 4294967295)) {
            v11 := GenS2{F0: 0, F1: -1}
            _ = v11
        } else {
            var v12 bool = false
            fmt.Println(v10)
            var v13 int16 = -1
            _ = v12
            _ = v13
        }
        fmt.Println(v4)
        _ = v10
    } else {
        if ((v1 ^ (v1 & (v1 & v1))) > v1) {
            v3 = (v3 << ((956) & 63))
            v3 = (^(v3))
            fmt.Println(v3)
        }
    }
    v14 := []int{v3, (v3 - v3), (v3 << ((4294967295) & 63)), v3}
    switch v3 {
    case 2147483647:
        switch 154 {
        case 9223372036854775807:
            _ = len(v14)
        }
    case 736:
        _ = len(v14)
    case 32768:
        if true {
            fmt.Println(v4, v2)
            v2 += 217
        } else {
            if false {
                v4 = genFn0()
                var v15 int8 = 63
                _ = v15
            }
        }
    default:
        v3 -= v3
    }
    var v16 int32 = int32(v3)
    var v17 int = -9223372036854775808
    v14 = append(v14, (-(v17)))
    v4 += genFn0()
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v1, v4)
    case "a":
        v17 += v3
    default:
        switch 128 {
        case 4294967295:
            var v18 int8 = genFn1(127)
            _ = v18
        }
    }
    v17 = 2147483647
    v4 = genFn0()
    var v19 int32 = 2147483647
    var v20 uint = 566
    for fi21 := 0; fi21 < 7; fi21++ {
        v22 := GenS3{F0: v16}
        var v23 int8 = -1
        v5 += 1.5
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 2, F1: -140}
    if false {
        fmt.Println(v5, v4)
        v25 := GenS1{F0: v16}
        var v26 float32 = 0.1
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
