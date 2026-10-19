package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }

func main() {
    var v1 int8 = 1
    var v2 int32 = -810
    var v3 int8 = (v1 / ((int8(v2)) | 1))
    v4 := []int{9223372036854775807, -2147483648, int(v2)}
    var v5 int16 = int16(v1)
    var v6 int8 = v3
    v2 += int32(v1)
    var v7 bool = (v3 == v1)
    v4 = append(v4, -917)
    var v8 int16 = (-(v5))
    var v9 int64 = 1
    var v10 uint8 = 250
    var v11 float32 = 0.0
    v8 += v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 float64 = 0.0
        if false {
            var v14 int64 = (v9 | 0)
            if v7 {
                if ((2 <= 197) && (!(!v7))) {
                    fmt.Println(v8, v5, v6)
                    v15 := GenS1{F0: v2, F1: 100}
                    _ = v15
                }
                if len(v4) > 0 {
                    v4[0] = fi12
                }
            } else {
                _ = len(v4)
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                v2 -= (v2 % ((v2) | 1))
                _ = fi16
            }
            _ = v14
        }
        v5 = v5
        _ = fi12
        _ = v13
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v2 = v2
        v10 += v10
        if (352 == 623) {
            var v18 uint16 = 32767
            v10 = v10
            if len(v4) > 0 {
                v4[0] = 42
            }
            _ = v18
        }
        _ = fi17
    }
    var v19 uint = 589
    if len(v4) > 0 {
        v4[0] = -977
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
