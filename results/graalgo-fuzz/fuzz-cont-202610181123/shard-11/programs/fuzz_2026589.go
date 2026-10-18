package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 bool }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 string; F1 bool }
func genFn0(p0 uint32) int8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := []int{(fi5 % ((fi5) | 1)), 747, (fi5 & fi5)}
        p0 = (p0 * (p0 ^ (^(p0))))
        _ = fi5
        _ = v6
    }
    p0 -= p0
    v7 := GenS1{F0: -725, F1: 0, F2: false}
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    _ = v7
    return 42
}
func genFn1() bool {
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: false}
        var v10 float32 = 0.0
        var v11 int8 = 100
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int64 = 71
    switch 442 {
    case -690:
        switch 0 {
        case 127:
            v13 := GenS2{F0: (32767 >= -358), F1: 18446744073709551615}
            _ = v13
        }
    case 32768:
        v14 := GenS3{F0: "", F1: (42 >= -32768)}
        _ = v14
    default:
        v15 := GenS2{F0: false, F1: 96}
        _ = v15
    }
    _ = v12
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 65535
    v3 := []int{}
    var v4 int32 = 0
    v4 += v4
    v4 = (^(v4))
    v2 -= v2
    if (v4 < v4) {
        v2 += v2
    }
    var v16 string = "\n\t"
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v2 = v2
        v4 = int32(v2)
        v17 := GenS3{F0: "", F1: (395 >= 225)}
        _ = v17
    }
    var v18 uint = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
