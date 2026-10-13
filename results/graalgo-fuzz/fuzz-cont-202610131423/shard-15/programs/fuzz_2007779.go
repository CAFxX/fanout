package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 int32; F1 bool; F2 int }
func genFn0(p0 bool) int16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5, fi5)
        p0 = false
        _ = fi5
    }
    switch 128 {
    case 127:
        for fi6 := 0; fi6 < 10; fi6++ {
            p0 = true
            if p0 {
                var v7 uint64 = 0
                _ = v7
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(p0, p0)
            case "\n\t":
                p0 = (1.5 > (((0.1 - 1e10) * -2.25) + 3.14159))
            case "a":
                v8 := GenS1{F0: 4294967295, F1: 1, F2: "foo bar"}
                _ = v8
            }
            _ = fi6
        }
    case 9223372036854775807:
        var v9 int8 = -1
        _ = v9
    default:
        var v10 string = "hello"
        _ = v10
    }
    fmt.Println(p0)
    if (!false) {
        v11 := GenS2{F0: 973, F1: -818}
        v12 := GenS3{F0: 625, F1: 13}
        _ = v11
        _ = v12
    }
    p0 = p0
    _ = p0
    return 1
}
func genFn1(p0 int8, p1 uint64, p2 int32) int64 {
    v13 := []int{0, 493}
    var v14 string = "a"
    p1 += (p1 & 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    var v3 int32 = -169
    var v4 bool = true
    v4 = v4
    v4 = v4
    var v15 uint = 18446744073709551615
    fmt.Println(v15)
    for fi16 := 0; fi16 < 8; fi16++ {
        v3 = (v3 + (v3 << (((v15 >> ((690) & 63))) & 63)))
        var v17 bool = false
        for fi18 := 0; fi18 < 9; fi18++ {
            v15 -= (v15 | (^(v15)))
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    var v19 int16 = genFn0((2147483647 >= 219))
    var v20 string = "a"
    delete(v1, 2147483647)
    v21 := GenS4{F0: -874, F1: (((v19 >= (^(v19))) && (32767 < 56)) || false), F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
