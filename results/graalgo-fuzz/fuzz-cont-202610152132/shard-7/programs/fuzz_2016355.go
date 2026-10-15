package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32; F2 int64 }
type GenS4 struct { F0 int32; F1 uint; F2 uint16 }
func genFn0(p0 int32, p1 uint) uint64 {
    if false {
        var v5 int64 = 9223372036854775807
        p0 = p0
        _ = v5
    } else {
        v6 := []int{0}
        _ = v6
    }
    switch "a" {
    case "":
        fmt.Println(p0)
    case "\n\t":
        if false {
            switch "\n\t" {
            case "hello":
                var v7 int16 = 1000
                _ = v7
            case "":
                if true {
                    p0 = p0
                }
            }
        } else {
            p0 += (p0 | p0)
            var v8 int8 = 0
            _ = v8
        }
    default:
        switch 32768 {
        case 32768:
            if (575 != 18446744073709551615) {
                var v9 uint32 = 2147483647
                _ = v9
            } else {
                p1 -= p1
                for fi10 := 0; fi10 < 17; fi10++ {
                    p1 = (^(p1))
                    _ = fi10
                }
                if (32767 <= 553) {
                    var v11 uint16 = 255
                    v12 := GenS1{F0: -1, F1: -128}
                    _ = v11
                    _ = v12
                } else {
                    v13 := GenS1{F0: 0, F1: 42}
                    p0 -= p0
                    v14 := GenS2{F0: 18446744073709551615, F1: 42}
                    _ = v13
                    _ = v14
                }
            }
        case -9223372036854775808:
            v15 := []int{-2147483648}
            _ = v15
        default:
            v16 := []int{-211}
            _ = v16
        }
    }
    if (!false) {
        var v17 int64 = -9223372036854775808
        v18 := []int{127, 2147483647, -9223372036854775808, 2147483647}
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 uint8 = 132
    v2 := make(map[uint32]string)
    v3 := []int{-1, 345}
    v4 := []int{0, -9223372036854775808, 333, -473}
    v1 += 0
    v19 := GenS1{F0: 0, F1: -1}
    for fi20 := 0; fi20 < 15; fi20++ {
        v21 := GenS2{F0: 582, F1: 1}
        if false {
            for fi22 := 0; fi22 < 19; fi22++ {
                v1 -= v1
                _ = fi22
            }
            for fi23 := 0; fi23 < 5; fi23++ {
                switch "hello" {
                case "":
                    v24 := GenS3{F0: 890, F1: -2147483648, F2: 0}
                    _ = v24
                case "foo bar":
                    v25 := GenS4{F0: int32(v1), F1: 2, F2: 32767}
                    var v26 string = ("hello" + ("hello" + (("" + (("a" + "foo bar") + "hello")) + "hello")))
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        } else {
            var v27 int64 = -1
            v27 = v27
            if (0 == -1) {
                v1 = v1
                v1 = (v1 >> ((65535) & 63))
                _ = len(v2)
            } else {
                for fi28 := 0; fi28 < 17; fi28++ {
                    v4 = append(v4, -244)
                    v27 += -1
                    _ = fi28
                }
                v1 = (v1 & (v1 % ((191) | 1)))
            }
            _ = v27
        }
        v1 -= v1
        _ = fi20
        _ = v21
    }
    if (!(v1 <= 1)) {
        v29 := GenS3{F0: 42, F1: -235, F2: 971}
        _ = v29
    } else {
        v1 = v1
        v1 -= (v1 - v1)
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v1 += 1
        _ = fi30
    }
    switch 2147483647 {
    case 9223372036854775807:
        var v31 uint16 = 2
        _ = v31
    default:
        var v32 bool = false
        _ = v32
    }
    fmt.Println(v1)
    var v33 uint8 = (v1 ^ (^(v1)))
    var v34 int8 = -128
    var v35 int64 = 1
    v1 -= (v1 ^ (^(v1)))
    v36 := GenS1{F0: 0, F1: -128}
    var v37 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
