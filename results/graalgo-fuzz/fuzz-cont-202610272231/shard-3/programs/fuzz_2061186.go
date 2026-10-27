package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1000
    v3 := []int{}
    var v4 uint32 = 942
    var v5 uint = 959
    var v6 uint16 = 32767
    v7 := GenS1{F0: 2, F1: false}
    var v8 uint16 = (^(v6))
    _ = len(v1)
    if (false && false) {
        if len(v3) > 0 {
            v3[0] = 262
        }
    }
    var v9 string = "a"
    v10 := v1[""]
    var v11 uint32 = (v4 - v4)
    switch "a" {
    case "\n\t":
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int8 = 100
            v10 = false
            _ = fi12
            _ = v13
        }
    case "foo bar":
        if (v10 || v10) {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint32 = v11
                var v16 int32 = 2147483647
                _ = fi14
                _ = v15
                _ = v16
            }
            if v10 {
                v17 := GenS2{F0: 0}
                var v18 uint64 = 96
                _ = v17
                _ = v18
            }
            v4 += v4
        } else {
            var v19 int64 = -9223372036854775808
            fmt.Println(v11)
            var v20 int8 = -128
            _ = v19
            _ = v20
        }
    case "a":
        if v10 {
            if v10 {
                var v21 int = -1
                if v10 {
                    fmt.Println(v21, v10, v4)
                    if len(v3) > 0 {
                        v3[0] = 104
                    }
                }
                _ = v21
            }
            if (v11 > v4) {
                v4 -= (v4 | 13)
            }
        }
    default:
        v22 := v1[("a" + "a")]
        _ = v22
    }
    var v23 uint64 = 2
    v24 := GenS2{F0: (v23 >> (((^(v5))) & 63))}
    v25 := GenS3{F0: 63}
    _ = len(v3)
    v26 := GenS4{F0: 156, F1: v8}
    var v27 uint16 = (v6 % ((561) | 1))
    var v28 float64 = (0.1 / 1.0)
    var v29 float64 = v28
    var v30 uint8 = 127
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
