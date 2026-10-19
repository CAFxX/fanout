package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int16; F1 int8; F2 uint }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int64) uint32 {
    p0 -= (p0 | p0)
    fmt.Println(p0, p1)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 2147483647
}
func genFn1(p0 int, p1 int64) uint8 {
    var v4 int = 4294967295
    var v5 int = 9223372036854775807
    v6 := []int{(p0 * v4), 127, (-(p0))}
    v6 = append(v6, (-(p0)))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 142
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    var v7 int8 = -125
    v7 += 0
    var v8 int32 = 2147483647
    var v9 int32 = -1
    v7 += (v7 >> ((255) & 63))
    switch "hello" {
    case "\n\t":
        v10 := GenS1{F0: 255, F1: 4294967295}
        _ = v10
    case "":
        var v11 int8 = 100
        _ = v11
    }
    fmt.Println(v7)
    v12 := GenS2{F0: "foo bar", F1: true, F2: 381}
    if false {
        v8 = v8
        v9 -= -482
        switch (("" + "foo bar") + "") {
        case "foo bar":
            var v13 uint32 = 941
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := []int{(fi14 - fi14), (^(fi14)), (^(fi14))}
            v7 -= (v7 | v7)
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 uint32 = 65535
            _ = fi16
            _ = v17
        }
        v8 -= (v8 ^ -2147483648)
    }
    v18 := v3[617]
    if (!true) {
        v8 -= 563
        v9 -= (v8 ^ v8)
        var v19 string = ""
        _ = v19
    }
    if v18 {
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS3{F0: -172, F1: v7, F2: 0}
            _ = fi20
            _ = v21
        }
        v22 := GenS4{F0: 42}
        var v23 int8 = (v7 % ((v7) | 1))
        _ = v22
        _ = v23
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
