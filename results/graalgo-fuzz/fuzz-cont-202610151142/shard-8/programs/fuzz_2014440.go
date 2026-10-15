package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool, p1 uint16) uint32 {
    switch "" {
    case "\n\t":
        switch "a" {
        case "hello":
            v4 := []int{42}
            _ = v4
        case "":
            p1 -= p1
        }
    default:
        v5 := GenS1{F0: 118, F1: "hello"}
        _ = v5
    }
    p1 = (p1 & (p1 - p1))
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := []int{-240, -276, -9223372036854775808, 42}
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int64, p1 uint) int64 {
    v8 := GenS1{F0: -1, F1: ""}
    var v9 float32 = 0.1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (-(p0))
}

func main() {
    var v1 int8 = -99
    v2 := make(map[uint32]string)
    var v3 uint = 4294967295
    v10 := []int{-2147483648, 32768}
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            v1 = 1
            _ = fi12
        }
        v1 -= (v1 % ((v1) | 1))
        _ = fi11
    }
    v3 -= v3
    v13 := GenS2{F0: 759, F1: 128}
    v1 = (v1 - v1)
    var v14 int16 = 32767
    v3 = 4294967295
    v15 := GenS1{F0: (v1 % (((^(v1))) | 1)), F1: ("hello" + "a")}
    v16 := GenS2{F0: 1, F1: 56}
    switch "a" {
    case "\n\t":
        v3 = 840
    default:
        _ = len(v10)
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    v17 := GenS3{F0: (v3 & 857), F1: 32767}
    fmt.Println(v1, v14, v14)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
