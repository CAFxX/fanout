package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0(p0 int16) uint {
    switch 128 {
    case -1:
        fmt.Println(p0)
    }
    v5 := []int{}
    _ = p0
    _ = v5
    return 1
}

func main() {
    v1 := []string{"", "hello", ""}
    var v2 uint16 = 649
    var v3 uint16 = (v2 - (v2 / ((v2) | 1)))
    v4 := make(map[string]uint32)
    v3 += (v2 ^ 255)
    v2 = (^(v2))
    var v6 int32 = 2147483647
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    if (90 != 1) {
        v2 -= 42
        v4[("hello" + (("a" + "a") + "hello"))] = 377
    } else {
        var v7 bool = ((0.0 + ((((1e10 - float32(-2.25)) + float32(float64(-2.25))) - (1.5 * ((0.1 - float32(3.14159)) - -2.25))) - (1.5 + (0.1 / (-2.25 - 0.1))))) < float32(float64((0.1 / 1e10))))
        _ = len(v4)
        fmt.Println(v7, v2, v2)
        _ = v7
    }
    v8 := GenS1{F0: 4294967295, F1: -661, F2: -1}
    v3 += (v2 / ((v3) | 1))
    fmt.Println(v6)
    v9 := GenS2{F0: 127}
    if len(v1) > 0 {
        v1[0] = ""
    }
    if true {
        var v10 float32 = 0.0
        if (-401 == 42) {
            v6 += (v6 % ((v6) | 1))
            _ = len(v4)
            v11 := GenS3{F0: -1, F1: 100}
            _ = v11
        }
        _ = v10
    } else {
        if false {
            delete(v4, ("foo bar" + ""))
            if (0 == -649) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v12 uint64 = 0
                _ = v12
            }
            var v13 int32 = (v6 & v6)
            _ = v13
        }
        v4[("\n\t" + "")] = 571
        fmt.Println(v3, v3)
    }
    var v14 int32 = (-(v6))
    switch "\n\t" {
    case "":
        var v15 int8 = -1
        _ = v15
    default:
        var v16 int = -648
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
    _ = v9
}
