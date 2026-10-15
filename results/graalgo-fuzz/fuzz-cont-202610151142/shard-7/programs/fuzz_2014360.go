package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
func genFn0(p0 int) uint64 {
    var v4 int = (p0 + (p0 | (-(p0))))
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := []int{2147483647, v4, (-(fi5)), fi5}
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 42
    _ = p0
    _ = v4
    _ = v7
    return 937
}
func genFn1(p0 uint8, p1 int32, p2 uint32) int {
    v8 := []int{-9223372036854775808, -698, 2147483647, 32767}
    p0 = 255
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v9 float32 = float32(float64(float32((1e10 - (3.14159 + -2.25)))))
    v10 := GenS1{F0: true, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 32768
}

func main() {
    var v1 int64 = 273
    var v2 uint32 = 65535
    var v3 int8 = 100
    v11 := []int{}
    for fi12 := 0; fi12 < 8; fi12++ {
        v1 += (v1 >> ((416) & 63))
        _ = fi12
    }
    if false {
        var v13 uint16 = 0
        if false {
            var v14 bool = (true && false)
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
            _ = v14
        }
        v13 -= v13
        _ = v13
    } else {
        v11 = append(v11, 42)
    }
    v3 -= v3
    if false {
        if len(v11) > 0 {
            v11[0] = 0
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                v16 := GenS1{F0: ((fi15 & fi15) >= (fi15 << ((2) & 63))), F1: 0}
                _ = len(v11)
                fmt.Println(v2, fi15, v3)
                _ = v16
            }
            var v17 uint32 = 1
            _ = fi15
            _ = v17
        }
    } else {
        v11 = append(v11, 1)
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v18 uint8 = 104
    v19 := GenS1{F0: true, F1: -8}
    if (v3 == v3) {
        v2 -= v2
        v18 -= v18
        var v20 float32 = (3.14159 - 1.5)
        _ = v20
    }
    v1 += v1
    v21 := GenS1{F0: (!false), F1: -700}
    var v22 string = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
