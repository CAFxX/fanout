package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    var v1 int = 32768
    var v2 int16 = 0
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    v5 := GenS1{F0: v1, F1: -1}
    switch (v1 ^ (v1 * v1)) {
    case -592:
        if (false || (!(42 != v1))) {
            var v6 uint64 = 341
            if false {
                v7 := []int{v1, 1}
                v2 += (v2 - -32768)
                v1 += -2147483648
                _ = v7
            } else {
                fmt.Println(v1)
                var v8 int16 = v2
                _ = v8
            }
            v9 := GenS2{F0: v1}
            _ = v6
            _ = v9
        }
    case 9223372036854775807:
        v10 := v3[(v1 << ((65535) & 63))]
        _ = v10
    case -919:
        var v11 uint64 = 524
        _ = v11
    }
    if (float32(1.5) <= -2.25) {
        switch (-(v1)) {
        case -1:
            var v12 uint64 = 0
            _ = v12
        case 219:
            v2 += v2
        case 128:
            var v13 int = (v1 * -2147483648)
            _ = v13
        }
        var v14 int8 = 100
        _ = v14
    } else {
        v15 := GenS3{F0: 1, F1: 42, F2: 1}
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = v1
        _ = fi16
        _ = v17
    }
    var v18 int8 = -128
    if true {
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS2{F0: (v1 | -2147483648)}
            v1 -= -753
            _ = fi19
            _ = v20
        }
        v21 := GenS1{F0: (v1 ^ 32767), F1: 42}
        _ = v21
    }
    var v22 uint8 = 225
    v23 := []int{}
    var v24 int8 = v18
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
