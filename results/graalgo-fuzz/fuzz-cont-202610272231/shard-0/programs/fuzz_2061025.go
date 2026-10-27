package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0(p0 uint8, p1 int32, p2 int64) string {
    v4 := []int{-9223372036854775808, 32767}
    p0 -= p0
    p1 -= 2147483647
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int8 = 98
        fmt.Println(p0, p1)
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return "a"
}
func genFn1(p0 int32) int64 {
    v7 := GenS1{F0: 32767, F1: "foo bar"}
    p0 += 42
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 46
    var v2 int64 = -1
    v3 := []uint32{1}
    var v8 uint = 4294967295
    v9 := GenS1{F0: 1, F1: "a"}
    v8 -= v8
    if (67 >= 127) {
        v1 += (v1 * 2)
    } else {
        fmt.Println(v8, v1)
        switch 42 {
        case 0:
            switch 2147483647 {
            case 1:
                fmt.Println(v1)
            case 638:
                var v10 float32 = float32(1.5)
                _ = v10
            case 0:
                if (false && true) {
                    v8 -= v8
                } else {
                    for fi11 := 0; fi11 < 10; fi11++ {
                        _ = len(v3)
                        var v12 float32 = 1.5
                        v13 := GenS1{F0: fi11, F1: "hello"}
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = len(v3)
                    v3 = append(v3, 1)
                }
            default:
                v2 = (v2 & (v2 / (((v2 / (((v2 % ((-1) | 1))) | 1))) | 1)))
            }
        case 127:
            _ = len(v3)
        default:
            v2 = v2
        }
        v2 -= 42
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        _ = len(v3)
        _ = fi14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 uint8 = (v1 - v1)
        v8 -= v8
        _ = fi15
        _ = v16
    }
    v17 := GenS1{F0: 32767, F1: ("hello" + "hello")}
    fmt.Println(v1)
    v8 = v8
    var v18 uint32 = uint32(v8)
    v19 := GenS1{F0: -1, F1: "a"}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
