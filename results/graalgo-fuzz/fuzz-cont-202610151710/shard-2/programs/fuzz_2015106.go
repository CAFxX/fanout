package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0() int {
    v6 := []int{-251, -2147483648, 32768}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 int64 = 0
        for fi9 := 0; fi9 < 15; fi9++ {
            if true {
                _ = len(v6)
                v8 -= v8
                v8 -= v8
                v8 = 0
            }
            _ = fi9
        }
        switch 304 {
        case 0:
            v10 := GenS1{F0: 236, F1: 0}
            _ = v10
        case 2147483647:
            var v11 uint = 4294967295
            _ = v11
        case 1:
            fmt.Println(fi7, v8)
        default:
            if (8 >= -1) {
                var v12 string = (("hello" + "a") + ("" + ""))
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v12
            } else {
                if false {
                    v6 = append(v6, -17)
                    v13 := GenS1{F0: 983, F1: 425}
                    _ = v13
                }
            }
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 325
}
func genFn1() uint32 {
    var v14 int = -1
    v14 -= v14
    var v15 uint64 = 18446744073709551615
    _ = v14
    _ = v15
    return 0
}

func main() {
    var v1 uint8 = 255
    var v2 int16 = -10
    var v3 int32 = 707
    var v4 int8 = -117
    var v5 uint64 = 0
    v2 -= -418
    var v16 string = "hello"
    v5 -= (^(v5))
    var v17 int = genFn0()
    fmt.Println(v2, v1)
    var v18 uint = uint(v5)
    v5 += v5
    var v19 int = v17
    var v20 string = ((v16 + v16) + v16)
    fmt.Println(v19)
    if (192 != 9223372036854775807) {
        var v21 int64 = -1
        var v22 int32 = (v3 / (((v3 & (^(v3)))) | 1))
        _ = v21
        _ = v22
    }
    var v23 uint = 2
    var v24 uint64 = 18446744073709551615
    switch v16 {
    case "":
        var v25 int32 = 42
        _ = v25
    default:
        v26 := []int{0, (^(v17)), (v19 + genFn0())}
        _ = v26
    }
    v1 += 42
    v27 := []int{genFn0(), (v19 * v19), (v17 ^ v19)}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
