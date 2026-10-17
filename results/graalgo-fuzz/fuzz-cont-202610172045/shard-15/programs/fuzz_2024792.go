package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16, p1 int8) int {
    var v5 uint8 = 255
    if ((16 == int16(v5)) || (-1 < (p0 * p0))) {
        var v6 uint32 = 2
        v7 := []int{32768, -2147483648, -213, 0}
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 float64 = (0.0 * 0.1)
            var v10 int8 = p1
            if (169 == 322) {
                v11 := GenS1{F0: 633, F1: 1}
                _ = v11
            } else {
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                var v12 int = fi8
                if false {
                    _ = len(v7)
                } else {
                    fmt.Println(p0, fi8)
                    var v13 string = "hello"
                    _ = v13
                }
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    if true {
        var v14 bool = ("" > "hello")
        v14 = v14
        for fi15 := 0; fi15 < 7; fi15++ {
            var v16 float64 = 1e10
            v5 = 180
            fmt.Println(v14)
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        var v17 int8 = p1
        var v18 float32 = 3.14159
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v5
    return -694
}

func main() {
    var v1 int32 = 1
    var v2 uint32 = 42
    var v3 int = 474
    var v4 int16 = -190
    if (2147483647 > 431) {
        var v19 int = 0
        v3 = v19
        _ = v19
    } else {
        v20 := []int{v3, (v3 >> ((712) & 63)), int(v1), 32767}
        v21 := GenS2{F0: 0}
        v3 -= v3
        _ = v20
        _ = v21
    }
    var v22 uint32 = 2
    v23 := GenS2{F0: 255}
    if (0 != -919) {
        for fi24 := 0; fi24 < 18; fi24++ {
            switch "\n\t" {
            case "":
                for fi25 := 0; fi25 < 12; fi25++ {
                    v1 = (v1 ^ (v1 + v1))
                    _ = fi25
                }
            case "a":
                if false {
                    var v26 uint16 = 255
                    v27 := []int{v3}
                    _ = v26
                    _ = v27
                }
            case "\n\t":
                var v28 float32 = 1e10
                _ = v28
            }
            _ = fi24
        }
        var v29 int8 = -1
        v4 += v4
        _ = v29
    } else {
        for fi30 := 0; fi30 < 7; fi30++ {
            v22 -= v2
            _ = fi30
        }
    }
    v4 = (^(v4))
    v31 := []int{127, -719}
    fmt.Println(v22, v3)
    v3 = (v3 - genFn0((v4 - (-(v4))), -128))
    v3 -= (-(v3))
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
}
