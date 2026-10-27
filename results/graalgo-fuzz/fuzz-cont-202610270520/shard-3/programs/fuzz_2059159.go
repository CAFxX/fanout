package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint8; F1 bool; F2 int64 }

func main() {
    v1 := []uint32{4294967295, 42, 2147483648}
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    v4 := []int{}
    v5 := make(map[int]uint32)
    v6 := make(map[string]uint32)
    var v7 bool = ((9223372036854775807 > -1) || ((1.5 * -2.25) == 0.1))
    v7 = v7
    v7 = (113 != 326)
    v7 = (2147483647 >= 1)
    if (-2147483648 != -732) {
        var v8 int64 = -985
        _ = v8
    } else {
        switch 757 {
        case 32768:
            var v9 string = (("" + "a") + "a")
            _ = v9
        case 828:
            v7 = (0 <= 65535)
        }
        var v10 float32 = -2.25
        _ = v10
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if len(v1) > 0 {
            v1[0] = 42
        }
    } else {
        v11 := GenS1{F0: -871, F1: (v7 || true)}
        v12 := GenS2{F0: 200, F1: ((65535 != 840) || v7), F2: -1}
        v7 = v7
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 255, F1: (65535 != 282), F2: 0}
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if (-32768 >= 216) {
                v16 := GenS1{F0: 1, F1: false}
                var v17 int32 = -2147483648
                _ = v16
                _ = v17
            } else {
                fmt.Println(v7, fi14, v7)
            }
            _ = fi15
        }
        _ = fi14
    }
    v7 = false
    var v18 int32 = -2147483648
    fmt.Println(v18, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
