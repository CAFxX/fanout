package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0(p0 int8, p1 uint8, p2 uint64) uint64 {
    p2 += (p2 >> ((887) & 63))
    p1 += (p1 - 4)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 65535
    v3 := make(map[int]string)
    v4 := []int{42}
    var v5 uint = 255
    v6 := GenS1{F0: -1, F1: 790, F2: v5}
    v2 = (^(v2))
    v7 := GenS1{F0: 0, F1: -474, F2: 18}
    v8 := GenS2{F0: 4294967295, F1: ("" + v1), F2: genFn0(0, 83, 0)}
    v1 = (("\n\t" + v1) + v1)
    if (1 >= 9223372036854775807) {
        var v9 uint = (v5 / (((v5 >> ((2) & 63))) | 1))
        _ = v9
    } else {
        v3[32767] = "hello"
        v2 += 32767
        v4 = append(v4, -9223372036854775808)
    }
    if (v1 >= (("foo bar" + v1) + (v1 + "hello"))) {
        if (-1 <= -32768) {
            var v10 uint16 = v2
            _ = v10
        }
        v11 := GenS3{F0: -660, F1: 9223372036854775807}
        _ = v11
    } else {
        v5 = v5
        _ = len(v3)
        v2 = (v2 >> (((v5 | (v5 ^ (v5 | (^(v5)))))) & 63))
    }
    if len(v4) > 0 {
        v4[0] = 42
    }
    switch "" {
    case "a":
        if (-340 > 1) {
            v1 = (v1 + "a")
            var v12 int16 = 1000
            _ = v12
        } else {
            var v13 int16 = -32768
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                v13 -= 42
            } else {
                v13 = 1
                var v14 int16 = -977
                _ = v14
            }
            _ = v13
        }
    default:
        var v15 float32 = (1e10 * -2.25)
        _ = v15
    }
    v5 -= v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
