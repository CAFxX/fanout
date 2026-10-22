package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float32 = float32(0.1)
    var v2 string = ("foo bar" + "\n\t")
    var v3 bool = (42 > 42)
    if false {
        v4 := GenS1{F0: 510, F1: 32767}
        _ = v4
    }
    v5 := []int{-9223372036854775808, 127}
    v2 = ((("hello" + v2) + (v2 + "hello")) + v2)
    v5 = append(v5, -363)
    fmt.Println(v3, v1)
    fmt.Println(v2)
    v5 = append(v5, 32767)
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS2{F0: 127}
        _ = fi6
        _ = v7
    }
    v2 = v2
    if (0 <= -619) {
        fmt.Println(v3, v3, v3)
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                if (!(100 > 0)) {
                    v1 += v1
                    v10 := GenS1{F0: (fi8 >> ((2) & 63)), F1: -32768}
                    v3 = v3
                    _ = v10
                } else {
                    _ = len(v5)
                    var v11 bool = (0 != 65535)
                    v1 = v1
                    v11 = (false || ((1e10 >= 0.0) && v11))
                    var v12 uint32 = 2147483647
                    _ = v11
                    _ = v12
                }
                var v13 float32 = -2.25
                v5 = append(v5, fi8)
                _ = fi9
                _ = v13
            }
            var v14 float32 = v1
            v1 = (float32(((1.5 - (1e10 + 0.1)) * -2.25)) / v1)
            _ = fi8
            _ = v14
        }
        if (-687 != -1) {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        } else {
            var v15 uint = 748
            _ = v15
        }
    }
    v5 = append(v5, -734)
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(v1, v3)
        _ = fi16
    }
    var v17 int32 = -300
    var v18 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
}
