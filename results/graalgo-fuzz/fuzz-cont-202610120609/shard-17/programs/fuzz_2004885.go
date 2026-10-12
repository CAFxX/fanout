package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 bool; F1 int16 }
func genFn0() int32 {
    var v4 int64 = 9223372036854775807
    v5 := []int{9223372036854775807, 798, 42}
    v4 += v4
    v6 := GenS1{F0: 326, F1: 0, F2: 4294967295}
    _ = v4
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := []int{}
    var v7 uint8 = 0
    v7 = v7
    v7 = (v7 + v7)
    for fi8 := 0; fi8 < 2; fi8++ {
        switch (^(fi8)) {
        case 2147483647:
            v7 = v7
        case -9223372036854775808:
            var v9 int8 = -128
            _ = v9
        case 282:
            for fi10 := 0; fi10 < 2; fi10++ {
                v3 = append(v3, (-(fi8)))
                v7 += v7
                _ = fi10
            }
        default:
            var v11 string = ""
            _ = v11
        }
        var v12 float64 = ((0.0 + ((0.1 - ((float64(3.14159) + 1e10) + (1e10 / 1.0))) + 1.5)) * float64(1.5))
        _ = fi8
        _ = v12
    }
    fmt.Println(v7)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v13 := v2[2147483647]
    var v14 int16 = int16(v7)
    var v15 uint8 = 128
    if false {
        if true {
            fmt.Println(v13, v15, v7)
            var v16 uint64 = 498
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint16 = 255
                _ = fi17
                _ = v18
            }
            _ = v16
        } else {
            var v19 int = 2147483647
            v20 := GenS1{F0: 2, F1: 2147483647, F2: 2}
            _ = v19
            _ = v20
        }
        var v21 int8 = -128
        _ = v21
    } else {
        v22 := GenS2{F0: (2 < 0), F1: 42}
        var v23 string = "a"
        switch -702 {
        case -605:
            switch (v23 + v23) {
            case "\n\t":
                var v24 int8 = 1
                _ = v24
            }
        case -2147483648:
            if len(v3) > 0 {
                v3[0] = 0
            }
        case 0:
            v25 := GenS2{F0: (!(!true)), F1: v14}
            _ = v25
        default:
            for fi26 := 0; fi26 < 16; fi26++ {
                switch (fi26 ^ fi26) {
                case 404:
                    var v27 int64 = -1
                    _ = v27
                case -9223372036854775808:
                    v13 = v23
                case -240:
                    var v28 uint32 = 2147483648
                    _ = v28
                }
                _ = fi26
            }
        }
        _ = v22
        _ = v23
    }
    var v29 int32 = 2147483647
    var v30 uint16 = 42
    v31 := v1[0]
    if true {
        v15 = (v15 + v15)
    }
    v14 -= -1
    v13 = v31
    if true {
        for fi32 := 0; fi32 < 16; fi32++ {
            v31 = v13
            if (-1 > 42) {
                _ = len(v1)
                v29 = (v29 >> ((65535) & 63))
            }
            _ = fi32
        }
        if (1 != 42) {
            fmt.Println(v30, v30)
            fmt.Println(v31, v31)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v7
}
