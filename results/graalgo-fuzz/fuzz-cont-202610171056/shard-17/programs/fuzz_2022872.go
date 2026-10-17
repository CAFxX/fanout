package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8; F1 int32; F2 int32 }

func main() {
    v1 := []uint32{}
    var v2 int16 = -228
    var v3 float64 = 3.14159
    if false {
        v1 = append(v1, 207)
        v2 = v2
    } else {
        for fi4 := 0; fi4 < 14; fi4++ {
            v5 := GenS1{F0: 424, F1: 0}
            if ((-1 <= 55) && false) {
                if (true || ((2 <= 23) || true)) {
                    v6 := GenS1{F0: 399, F1: 65535}
                    var v7 float64 = float64(float32(v3))
                    var v8 int8 = 42
                    _ = v6
                    _ = v7
                    _ = v8
                }
            } else {
                v9 := GenS1{F0: 537, F1: 42}
                _ = v9
            }
            v10 := GenS2{F0: 1, F1: "hello"}
            _ = fi4
            _ = v10
            _ = v5
        }
    }
    if true {
        v3 += 0.0
        switch "a" {
        case "":
            v11 := GenS1{F0: 42, F1: 634}
            _ = v11
        }
        v2 -= 1
    }
    fmt.Println(v2, v2)
    var v12 bool = (517 <= 452)
    v2 += (v2 ^ (v2 / (((v2 & v2)) | 1)))
    v3 = v3
    v12 = (v12 || v12)
    v2 -= v2
    if len(v1) > 0 {
        v1[0] = 440
    }
    _ = len(v1)
    var v13 float32 = 3.14159
    if false {
        v12 = (141 < 227)
        for fi14 := 0; fi14 < 17; fi14++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v3)
            v13 += 0.0
            _ = fi14
        }
    }
    var v15 string = "foo bar"
    if ((!(0 < 4294967295)) || (!((-2.25 > float32(v3)) && false))) {
        var v16 float64 = 0.0
        fmt.Println(v3)
        _ = v16
    }
    if v12 {
        v17 := GenS3{F0: -128, F1: 110, F2: -2147483648}
        fmt.Println(v12, v15, v2)
        _ = v17
    } else {
        fmt.Println(v13)
    }
    v15 = v15
    var v18 uint16 = 32767
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
