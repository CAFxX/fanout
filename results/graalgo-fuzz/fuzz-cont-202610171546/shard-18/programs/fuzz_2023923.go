package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint64, p1 uint8) int64 {
    v5 := GenS1{F0: (128 <= -443)}
    p0 = (p0 / ((p0) | 1))
    v6 := []int{-2147483648, -854, -2147483648, -762}
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 += v7
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 uint64, p2 uint64) bool {
    var v9 bool = false
    v9 = true
    fmt.Println(p0)
    var v10 int16 = -571
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return true
}

func main() {
    var v1 int8 = -115
    var v2 int16 = -565
    var v3 int8 = v1
    v4 := []string{"", ""}
    if true {
        var v11 bool = false
        if (v11 && false) {
            fmt.Println(v11, v11)
        } else {
            var v12 int16 = (v2 | v2)
            _ = v12
        }
        _ = v11
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int16 = v2
                var v16 float32 = float32(0.0)
                _ = fi14
                _ = v15
                _ = v16
            }
            fmt.Println(fi13, fi13)
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = fi13
        }
        _ = len(v4)
    }
    v2 = v2
    if len(v4) > 0 {
        v4[0] = (("\n\t" + "foo bar") + "\n\t")
    }
    var v17 int16 = v2
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 uint = 42
        fmt.Println(v2, v3)
        _ = fi18
        _ = v19
    }
    v2 -= v17
    var v20 int = 128
    switch v20 {
    case -1:
        if len(v4) > 0 {
            v4[0] = (("\n\t" + ("a" + "")) + "foo bar")
        }
    case 1:
        if false {
            if (0.1 >= ((3.14159 * 3.14159) + 3.14159)) {
                switch (v20 | (v20 + v20)) {
                case 32768:
                    v21 := GenS1{F0: false}
                    _ = v21
                case 0:
                    if (((genFn0(42, 54) != 1) && (2 >= 2)) || (!false)) {
                        fmt.Println(v17, v17, v1)
                        v1 = (v1 % ((int8(v2)) | 1))
                    } else {
                        v2 = v17
                        _ = len(v4)
                    }
                case 761:
                    v22 := GenS2{F0: 1, F1: 1}
                    _ = v22
                }
                v2 -= v17
            }
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 int64 = 1
                _ = fi23
                _ = v24
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
