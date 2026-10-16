package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int; F2 uint64 }
func genFn0(p0 uint32) int64 {
    v6 := []int{128, -368, -9223372036854775808}
    v6 = append(v6, 9223372036854775807)
    var v7 int32 = 1
    _ = len(v6)
    p0 = (p0 | 2147483647)
    _ = p0
    _ = v6
    _ = v7
    return -470
}

func main() {
    v1 := []uint32{2}
    v2 := []int{-2147483648, 32768}
    var v3 int32 = 203
    var v4 uint8 = 1
    v5 := []string{"foo bar"}
    v4 = (v4 * v4)
    switch 32768 {
    case 32767:
        var v8 float64 = 0.1
        _ = v8
    case 4294967295:
        var v9 uint16 = 32767
        _ = v9
    case 2147483647:
        v10 := GenS1{F0: 0}
        _ = v10
    }
    if false {
        if true {
            var v11 uint16 = 42
            _ = v11
        } else {
            switch 128 {
            case 348:
                fmt.Println(v3)
            }
            var v12 uint64 = 42
            if false {
                v12 += (v12 & (^(v12)))
            } else {
                fmt.Println(v12, v12)
                v12 += v12
                var v13 bool = false
                _ = v13
            }
            _ = v12
        }
        v3 = v3
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
    } else {
        switch "\n\t" {
        case "a":
            v3 = -1
        case "":
            var v14 uint32 = 65535
            _ = v14
        case "\n\t":
            v4 += 96
        }
    }
    var v15 int8 = 100
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 18; fi17++ {
            v18 := GenS1{F0: -13}
            _ = fi17
            _ = v18
        }
        v4 -= v4
        _ = fi16
    }
    v19 := GenS2{F0: 2, F1: -615, F2: 1}
    var v20 float32 = 1e10
    v21 := GenS1{F0: 790}
    fmt.Println(v20, v15, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
