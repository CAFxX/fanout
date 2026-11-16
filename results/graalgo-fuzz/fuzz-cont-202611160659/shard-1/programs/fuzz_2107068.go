package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 uint8) int16 {
    var v7 int16 = -1
    if (!((^(p0)) >= 2)) {
        if (((v7 == 1000) && ((p0 & (p0 >> ((1) & 63))) != (p0 % (((p0 * 2)) | 1)))) && false) {
            v7 -= -32768
            if (-1 >= -531) {
                var v8 int8 = 1
                v9 := GenS1{F0: 1}
                switch 42 {
                case 42:
                    if (42 <= 2) {
                        var v10 uint32 = 383
                        _ = v10
                    }
                default:
                    v11 := GenS1{F0: 97}
                    _ = v11
                }
                _ = v8
                _ = v9
            } else {
                fmt.Println(v7, p0)
                var v12 uint32 = 2
                var v13 float32 = float32((0.0 + (1e10 * ((3.14159 * ((3.14159 - 1.5) - float64(-2.25))) - 0.0))))
                _ = v12
                _ = v13
            }
        } else {
            p0 += p0
            fmt.Println(v7, p0)
        }
    } else {
        v14 := []int{32767}
        v7 = v7
        _ = v14
    }
    v15 := GenS2{F0: 4294967295, F1: 24}
    _ = p0
    _ = v15
    _ = v7
    return 1000
}

func main() {
    var v1 int = 0
    var v2 int8 = 1
    var v3 uint = 413
    var v4 uint32 = 0
    var v5 uint64 = 42
    v6 := make(map[int]string)
    var v16 bool = (-1 < 42)
    var v17 float64 = 1.5
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 string = ("" + (("" + "a") + (("foo bar" + "hello") + "hello")))
        _ = len(v6)
        _ = fi18
        _ = v19
    }
    var v20 float64 = (1e10 + (v17 + (v17 + (v17 / v17))))
    for fi21 := 0; fi21 < 19; fi21++ {
        if (3.14159 == -2.25) {
            var v22 uint = (v3 % (((^(v3))) | 1))
            _ = v22
        } else {
            var v23 int32 = 0
            _ = v23
        }
        _ = fi21
    }
    if false {
        v2 = -1
    }
    v3 = (v3 - v3)
    var v24 int = (v1 - -144)
    v25 := GenS2{F0: uint(v5), F1: -1}
    v16 = ((^(v4)) <= (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
