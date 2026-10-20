package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }
func genFn0() uint8 {
    v5 := []int{-2147483648, 0, -2147483648, -833}
    var v6 uint = 0
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 float64 = -2.25
    var v2 uint8 = 127
    var v3 int = 0
    var v4 uint64 = 18446744073709551615
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: -9223372036854775808, F1: v4, F2: 1}
        for fi9 := 0; fi9 < 1; fi9++ {
            if true {
                v10 := GenS1{F0: (fi7 | (fi7 - (-(fi7)))), F1: (v4 << ((135) & 63)), F2: 255}
                _ = v10
            }
            var v11 uint8 = genFn0()
            var v12 float32 = (-2.25 * float32(v1))
            _ = fi9
            _ = v11
            _ = v12
        }
        var v13 int = fi7
        _ = fi7
        _ = v13
        _ = v8
    }
    v14 := []int{}
    if false {
        var v15 int64 = -1
        var v16 float64 = float64(0.0)
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            v3 = (v3 % ((v3) | 1))
            var v18 uint64 = (v4 & (v4 | (v4 % (((v4 | (v4 - v4))) | 1))))
            _ = fi17
            _ = v18
        }
    }
    if len(v14) > 0 {
        v14[0] = v3
    }
    fmt.Println(v2, v4)
    if false {
        v2 += (v2 / ((v2) | 1))
    } else {
        var v19 string = ("a" + ("a" + "\n\t"))
        var v20 uint32 = 4294967295
        _ = v19
        _ = v20
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if false {
        var v21 int32 = -2147483648
        fmt.Println(v3, v1)
        _ = len(v14)
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
