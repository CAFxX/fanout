package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int = 42
    v4 := []int{127, v3}
    v5 := make(map[uint32]uint32)
    v6 := []string{"\n\t", "\n\t", "a", ("hello" + (("" + "a") + "hello"))}
    v3 = -9223372036854775808
    var v7 string = "foo bar"
    if (0 < -890) {
        if false {
            v3 += v3
            v3 = 127
        } else {
            var v8 int32 = 0
            v9 := GenS1{F0: 128, F1: 0}
            _ = v8
            _ = v9
        }
    } else {
        switch "a" {
        case "hello":
            v2[(^(v3))] = (-128 <= -1)
        }
        v7 = ("hello" + (v7 + "\n\t"))
        var v10 uint = 42
        _ = v10
    }
    var v11 uint = 496
    v12 := GenS2{F0: 2}
    if false {
        if true {
            _ = len(v1)
            v13 := GenS1{F0: 2, F1: 0}
            v3 += v3
            _ = v13
        }
    }
    var v14 int8 = 127
    var v15 uint64 = 990
    var v16 uint32 = 4294967295
    v14 += 127
    var v17 float32 = 1.5
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint16 = 817
        _ = len(v1)
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: 304}
    v7 = (v7 + v7)
    switch v7 {
    case "a":
        _ = len(v5)
    default:
        for fi21 := 0; fi21 < 15; fi21++ {
            switch v7 {
            case "hello":
                v11 -= (v11 & (v11 % ((0) | 1)))
            case "":
                var v22 int16 = 730
                _ = v22
            }
            if true {
                v16 += v16
                v23 := GenS3{F0: (v15 / ((1) | 1)), F1: v11}
                _ = v23
            }
            _ = fi21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
