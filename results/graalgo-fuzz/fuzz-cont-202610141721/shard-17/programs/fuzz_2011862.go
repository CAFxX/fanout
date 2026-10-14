package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int16, p1 uint64) int64 {
    switch 32767 {
    case 32768:
        var v5 uint32 = 4294967295
        _ = v5
    case 42:
        p1 += (p1 % ((133) | 1))
    case 4294967295:
        var v6 uint16 = 51
        _ = v6
    default:
        var v7 int32 = -1
        _ = v7
    }
    v8 := []int{-2147483648, 948}
    var v9 int = 127
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 int8, p1 uint, p2 int) int32 {
    for fi10 := 0; fi10 < 10; fi10++ {
        p0 = (p0 >> ((p1) & 63))
        _ = fi10
    }
    v11 := []int{p2, p2, (p2 / ((p2) | 1))}
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{65535, 161, 33, 2147483647}
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    var v12 uint = 1
    for fi13 := 0; fi13 < 9; fi13++ {
        v12 = (^(v12))
        switch "hello" {
        case "hello":
            var v14 int64 = 42
            _ = v14
        case "":
            v12 = 255
        default:
            var v15 bool = false
            _ = v15
        }
        v12 = (v12 >> (((^(v12))) & 63))
        _ = fi13
    }
    fmt.Println(v12)
    v16 := GenS1{F0: 1, F1: 2147483647}
    v17 := GenS1{F0: 219, F1: 4294967295}
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 float32 = 1e10
        var v20 uint8 = 255
        for fi21 := 0; fi21 < 4; fi21++ {
            v1 = append(v1, 256)
            v22 := GenS1{F0: (^(v20)), F1: 2147483648}
            switch fi21 {
            case 32768:
                switch (("a" + "") + "\n\t") {
                case "hello":
                    v12 += 18446744073709551615
                case "\n\t":
                    if false {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v20, fi18)
                        case "a":
                            v12 += (^(v12))
                        case "foo bar":
                            var v23 uint = 2
                            _ = v23
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v24 := GenS1{F0: (v20 >> ((1) & 63)), F1: 65535}
                        _ = v24
                    } else {
                        var v25 uint8 = 2
                        _ = v25
                    }
                default:
                    var v26 string = ""
                    _ = v26
                }
            case -1:
                fmt.Println(fi18, v12)
            case -296:
                v19 -= -2.25
            default:
                v4[(fi21 ^ (-(fi21)))] = 864
            }
            _ = len(v1)
            _ = fi21
            _ = v22
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v12 += 4294967295
    var v27 uint = (v12 >> (((v12 << ((v12) & 63))) & 63))
    v3[-9223372036854775808] = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
}
