package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 uint, p2 int8) int16 {
    var v5 int32 = -1
    if true {
        v6 := GenS1{F0: (v5 >> (((^(p1))) & 63))}
        _ = v6
    }
    switch -334 {
    case 32768:
        var v7 int = 300
        _ = v7
    default:
        fmt.Println(p1, p1, p2)
    }
    var v8 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 42
}
func genFn1() uint {
    if (967 == -2147483648) {
        v9 := GenS1{F0: 1}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 int = -300
            switch (-(v11)) {
            case -1:
                var v12 string = "\n\t"
                _ = v12
            default:
                fmt.Println(v11, v11)
            }
            _ = fi10
            _ = v11
        }
    }
    v13 := []int{508, 127, 32767, -650}
    _ = v13
    return 1
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[string]string)
    var v3 float32 = -2.25
    var v4 uint16 = 32767
    v3 = (3.14159 / 3.14159)
    v14 := []int{1}
    switch -2147483648 {
    case -1:
        switch "foo bar" {
        case "hello":
            var v15 bool = false
            _ = v15
        case "":
            fmt.Println(v4, v4)
        case "a":
            if false {
                if len(v14) > 0 {
                    v14[0] = 128
                }
            }
        default:
            if len(v14) > 0 {
                v14[0] = 32767
            }
        }
    case 128:
        v3 -= v3
    }
    v4 += (v4 * 2)
    if false {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int64 = 42
            var v18 uint16 = v4
            var v19 float32 = ((1e10 + 0.0) - float32(1e10))
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        v4 = v4
        v3 -= (v3 / v3)
    } else {
        v3 += float32(((float64((v3 - v3)) - 1e10) * 1.5))
        v4 += (^(v4))
        v20 := GenS2{F0: 119, F1: 65535}
        _ = v20
    }
    v1 = (v1 / ((v1) | 1))
    for fi21 := 0; fi21 < 5; fi21++ {
        var v22 uint64 = 1
        _ = fi21
        _ = v22
    }
    if len(v14) > 0 {
        v14[0] = 312
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
