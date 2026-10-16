package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64 }
func genFn0() string {
    var v5 int16 = 1000
    v6 := []int{625, 42}
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 0.0
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    v2 = float32((((0.0 - (1e10 / (((0.1 - 1e10) / (-2.25 * ((3.14159 * -2.25) * -2.25))) - (0.0 - 1e10)))) * -2.25) - (-2.25 + -2.25)))
    if (-949 == 234) {
        v7 := GenS1{F0: 639, F1: 0}
        _ = v7
    } else {
        fmt.Println(v2)
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint64 = 18446744073709551615
            v9 = (^(v9))
            _ = fi8
            _ = v9
        }
    }
    if false {
        v2 = float32(0.0)
    } else {
        if true {
            v10 := []int{-21, 729, -2147483648}
            switch 32768 {
            case 9223372036854775807:
                v2 -= float32(1e10)
            case 1:
                var v11 int = 128
                _ = v11
            default:
                v12 := GenS1{F0: 2, F1: -1}
                _ = v12
            }
            _ = v10
        }
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 18446744073709551615, F1: 2147483647}
        _ = fi13
        _ = v14
    }
    v2 += 0.1
    v15 := []int{9223372036854775807, 732}
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            v15 = append(v15, -9223372036854775808)
            var v18 uint = 2
            _ = len(v3)
            var v19 int32 = 42
            delete(v4, (((("" + ("a" + "hello")) + "") + "a") + ((("hello" + "\n\t") + "") + "hello")))
            _ = fi17
            _ = v18
            _ = v19
        }
        v2 = (float32(((0.0 / 1.5) + 0.0)) / v2)
        var v20 int64 = -1
        _ = fi16
        _ = v20
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        _ = len(v15)
        v22 := GenS2{F0: 2}
        v23 := GenS1{F0: 65535, F1: 42}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v2)
    var v24 int16 = 42
    v25 := GenS1{F0: 1, F1: -1}
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 uint = 865
        _ = fi26
        _ = v27
    }
    var v28 uint32 = 42
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    fmt.Println(v24, v24)
    for fi29 := 0; fi29 < 6; fi29++ {
        v2 -= (v2 * 3.14159)
        v1[252] = "hello"
        v30 := GenS2{F0: 149}
        _ = fi29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
