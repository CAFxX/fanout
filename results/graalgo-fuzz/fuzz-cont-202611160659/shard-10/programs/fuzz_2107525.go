package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 int64 }
func genFn0(p0 int, p1 uint32, p2 int16) int16 {
    if (42 > 42) {
        v7 := []int{}
        v8 := GenS1{F0: ((("hello" + (("" + "") + ("a" + ""))) + ("foo bar" + ("\n\t" + ("" + "")))) + ("\n\t" + ("\n\t" + ("a" + ("hello" + ""))))), F1: 65535}
        v7 = append(v7, p0)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        if (65535 < 959) {
            p1 -= (p1 << ((255) & 63))
        }
        p1 = (^(p1))
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 1000
}

func main() {
    v1 := []uint32{2147483648, 2147483648}
    v2 := []uint32{2147483648, 2}
    var v3 int64 = 9223372036854775807
    var v4 float32 = (1e10 / 0.1)
    v5 := []uint32{4294967295, 42, 65535, 2}
    v6 := make(map[string]int)
    fmt.Println(v3, v4)
    v10 := GenS2{F0: -1}
    v11 := GenS1{F0: "a", F1: 1}
    if len(v1) > 0 {
        v1[0] = 2
    }
    v12 := GenS3{F0: 65535, F1: -1, F2: 0}
    v13 := GenS3{F0: 42, F1: 0, F2: v3}
    fmt.Println(v4, v4)
    if (!false) {
        v4 += v4
        _ = len(v1)
        fmt.Println(v3, v3)
    } else {
        if (((false || (801 != 120)) && (42 < 1)) && (false && false)) {
            var v14 uint64 = 1
            _ = v14
        } else {
            var v15 bool = ((!false) || false)
            _ = len(v6)
            _ = v15
        }
        v3 += 817
    }
    var v16 int32 = 119
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
