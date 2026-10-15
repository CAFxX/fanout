package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 bool }
func genFn0() string {
    var v6 float64 = (0.1 * 1e10)
    v6 = v6
    var v7 float32 = -2.25
    var v8 int8 = -102
    _ = v6
    _ = v7
    _ = v8
    return ""
}
func genFn1(p0 int32, p1 int8) int64 {
    p1 -= (p1 << ((1) & 63))
    var v9 uint32 = 2
    var v10 int8 = (p1 + 21)
    if false {
        fmt.Println(p0, v9, v9)
        v11 := []int{2147483647, 4294967295, 32767}
        _ = v11
    } else {
        v12 := GenS1{F0: 1, F1: 64, F2: (^(v9))}
        if true {
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 uint32 = (v9 % ((967) | 1))
                _ = fi13
                _ = v14
            }
        }
        _ = v12
    }
    v15 := GenS1{F0: 25, F1: p0, F2: v9}
    _ = p0
    _ = p1
    _ = v10
    _ = v15
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint64 = 42
    v4 := []uint32{65535}
    v5 := make(map[int]uint32)
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 7; fi17++ {
            if (fi16 < fi16) {
                var v18 float64 = 1e10
                fmt.Println(fi16)
                switch fi16 {
                case -1:
                    v18 = ((v18 + v18) - 0.1)
                case 9223372036854775807:
                    _ = len(v4)
                case 0:
                    var v19 int64 = 9223372036854775807
                    _ = v19
                default:
                    v20 := GenS2{F0: true}
                    _ = v20
                }
                _ = v18
            } else {
                v3 = v3
                fmt.Println(fi17, fi17)
                var v21 int = (fi16 * 32767)
                _ = v21
            }
            var v22 float32 = 1.5
            _ = fi17
            _ = v22
        }
        _ = fi16
    }
    var v23 string = ("a" + ("foo bar" + "a"))
    v23 = v23
    var v24 uint64 = 42
    fmt.Println(v24, v24)
    var v25 float32 = (float32((1.5 + 0.0)) + 1.5)
    if (!true) {
        for fi26 := 0; fi26 < 10; fi26++ {
            v25 += float32(1e10)
            for fi27 := 0; fi27 < 17; fi27++ {
                fmt.Println(v23, v3)
                v3 -= v24
                _ = fi27
            }
            fmt.Println(fi26)
            _ = fi26
        }
    } else {
        v28 := v2[("" + v23)]
        _ = v28
    }
    var v29 int8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
