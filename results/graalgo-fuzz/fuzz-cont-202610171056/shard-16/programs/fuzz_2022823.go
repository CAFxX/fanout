package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool, p1 uint64, p2 int64) uint16 {
    v5 := GenS1{F0: 18446744073709551615}
    var v6 uint = 65535
    var v7 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 8
}

func main() {
    var v1 int8 = -54
    v2 := []int{4294967295, 32768, 0}
    var v3 int8 = (v1 ^ v1)
    var v4 string = "hello"
    if false {
        v4 = ((v4 + ("\n\t" + ("hello" + v4))) + v4)
        v3 -= v3
        v1 -= 42
    }
    fmt.Println(v4)
    if (115 >= genFn0((771 > 2), 1, 271)) {
        var v8 uint8 = 128
        _ = v8
    } else {
        var v9 uint16 = 665
        var v10 int32 = 42
        var v11 float64 = 1.5
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 4294967295}
    v3 -= v1
    var v13 uint64 = 218
    if (107 != 127) {
        v14 := GenS1{F0: 0}
        if true {
            if (!true) {
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 uint = 745
                    var v17 float64 = 1.5
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                switch 32768 {
                case 0:
                    v18 := GenS2{F0: 1, F1: 1}
                    _ = v18
                default:
                    if false {
                        fmt.Println(v4)
                        v13 = 4294967295
                    } else {
                        var v19 string = (v4 + "hello")
                        _ = v19
                    }
                }
                if ((3.14159 - 0.1) != 1e10) {
                    var v20 uint64 = (^(v13))
                    v3 = (-(v1))
                    v13 += (v20 | v20)
                    _ = v20
                }
            }
            v21 := GenS3{F0: 0}
            _ = v21
        } else {
            v22 := GenS3{F0: 103}
            v4 = (v4 + ((v4 + ("foo bar" + (("\n\t" + v4) + v4))) + v4))
            var v23 float32 = -2.25
            _ = v22
            _ = v23
        }
        if (float32(3.14159) <= 0.0) {
            if (128 < 127) {
                var v24 float64 = ((3.14159 + 1e10) / 1.5)
                if len(v2) > 0 {
                    v2[0] = 1
                }
                _ = v24
            } else {
                var v25 int32 = -195
                fmt.Println(v3, v4)
                v26 := GenS3{F0: 42}
                _ = v25
                _ = v26
            }
        }
        _ = v14
    }
    var v27 int16 = 32767
    for fi28 := 0; fi28 < 20; fi28++ {
        v4 = (((v4 + (v4 + (("foo bar" + v4) + "hello"))) + (v4 + ("hello" + v4))) + (v4 + (v4 + v4)))
        var v29 uint16 = 255
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
