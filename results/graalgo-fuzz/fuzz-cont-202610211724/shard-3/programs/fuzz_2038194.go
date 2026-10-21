package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16) string {
    if false {
        v5 := GenS1{F0: -337, F1: p0, F2: 1}
        _ = v5
    } else {
        p0 = p0
        fmt.Println(p0)
        var v6 int8 = -64
        _ = v6
    }
    switch ("" + "") {
    case "foo bar":
        var v7 uint32 = 2147483648
        _ = v7
    }
    _ = p0
    return "a"
}
func genFn1(p0 int16, p1 uint64, p2 uint64) int16 {
    switch 127 {
    case 42:
        var v8 uint = uint(p2)
        _ = v8
    case 585:
        v9 := []int{32767, 128, 42}
        _ = v9
    case 127:
        p2 = 42
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{(fi10 | fi10), (^(fi10))}
        var v12 string = ("hello" + "\n\t")
        _ = fi10
        _ = v11
        _ = v12
    }
    p2 = p1
    switch 70 {
    case 2147483647:
        p0 = 1000
    case 128:
        var v13 uint32 = 2147483648
        _ = v13
    default:
        v14 := GenS2{F0: 1}
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint64 = 2
    v4 := []uint32{0}
    _ = len(v4)
    if true {
        switch 32767 {
        case 1:
            v15 := GenS1{F0: 32767, F1: 1, F2: -1}
            _ = v15
        case -101:
            _ = len(v1)
        case -213:
            if (2147483647 <= 0) {
                var v16 int8 = 42
                _ = v16
            } else {
                for fi17 := 0; fi17 < 18; fi17++ {
                    _ = len(v1)
                    _ = len(v1)
                    if ((!true) && true) {
                        delete(v1, uint32(v3))
                    }
                    _ = fi17
                }
                if len(v4) > 0 {
                    v4[0] = 1
                }
                v3 = 1
            }
        }
    }
    fmt.Println(v3)
    v3 -= (v3 / ((v3) | 1))
    v3 = (v3 % (((v3 << ((1) & 63))) | 1))
    v3 -= (v3 % ((784) | 1))
    var v18 uint64 = (v3 + v3)
    v3 = 4294967295
    switch "\n\t" {
    case "hello":
        if (-1 <= 107) {
            fmt.Println(v3, v3)
            switch "hello" {
            case "":
                switch -1 {
                case 32767:
                    v3 = 792
                case 1:
                    fmt.Println(v3, v3)
                default:
                    v2[2147483648] = 562
                }
            case "hello":
                var v19 uint8 = 191
                _ = v19
            default:
                v18 -= (v3 & (v18 * v3))
            }
            fmt.Println(v18, v3)
        }
    }
    var v20 uint64 = 106
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
