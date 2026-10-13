package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint32) string {
    switch "" {
    case "a":
        switch 127 {
        case 0:
            if (1.5 < -2.25) {
                p0 += 4294967295
                p0 = (p0 - p0)
            }
        }
    case "foo bar":
        var v5 bool = true
        _ = v5
    case "":
        fmt.Println(p0)
    default:
        v6 := GenS1{F0: 4294967295, F1: 2, F2: 9223372036854775807}
        _ = v6
    }
    v7 := GenS1{F0: 2, F1: 65535, F2: 9223372036854775807}
    var v8 uint64 = 18446744073709551615
    for fi9 := 0; fi9 < 20; fi9++ {
        if (255 != 4294967295) {
            p0 += uint32(v8)
            if false {
                v10 := []int{fi9}
                v11 := GenS1{F0: 592, F1: 400, F2: (fi9 ^ fi9)}
                _ = v10
                _ = v11
            }
        } else {
            var v12 uint = 65535
            v12 = v12
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 int, p1 uint64, p2 uint8) int32 {
    p1 += p1
    if false {
        p2 = p2
        v13 := GenS1{F0: 42, F1: 255, F2: (p0 ^ (^(p0)))}
        v14 := []int{}
        _ = v13
        _ = v14
    } else {
        v15 := []int{(p0 * (^(p0))), 127}
        _ = v15
    }
    v16 := GenS1{F0: p1, F1: 768, F2: p0}
    switch p0 {
    case 127:
        for fi17 := 0; fi17 < 20; fi17++ {
            p1 = (p1 / (((^(p1))) | 1))
            _ = fi17
        }
    case 32767:
        switch p0 {
        case -9223372036854775808:
            p0 = 0
        default:
            v18 := []int{p0, p0}
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return 42
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := []string{"foo bar", "", "", "foo bar"}
    v19 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: -1000}
    var v20 uint8 = 42
    _ = len(v2)
    switch "foo bar" {
    case "foo bar":
        v21 := GenS2{F0: 65535}
        _ = v21
    }
    fmt.Println(v20)
    v22 := GenS1{F0: 2, F1: 353, F2: 32767}
    switch "hello" {
    case "hello":
        if (false && true) {
            v20 += v20
        }
    case "foo bar":
        var v23 uint = 317
        _ = v23
    case "a":
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(v20, fi24)
            _ = fi24
        }
    default:
        for fi25 := 0; fi25 < 7; fi25++ {
            var v26 uint = 4294967295
            _ = fi25
            _ = v26
        }
    }
    v20 += (v20 & (v20 - v20))
    _ = len(v1)
    if (-477 <= 125) {
        if (((2147483647 >= 338) || false) || (42 == 1)) {
            v27 := GenS1{F0: 49, F1: 18446744073709551615, F2: 674}
            v28 := GenS3{F0: -765, F1: 100}
            _ = v27
            _ = v28
        } else {
            if (false && false) {
                v20 = (v20 + 211)
            } else {
                switch ("" + "\n\t") {
                case "a":
                    fmt.Println(v20)
                }
                v29 := GenS3{F0: 0, F1: 100}
                _ = v29
            }
        }
        _ = len(v3)
    } else {
        v30 := GenS2{F0: 243}
        switch -467 {
        case 42:
            v20 -= 255
        case 483:
            fmt.Println(v20)
        case -1:
            v20 = v20
        }
        var v31 bool = true
        _ = v30
        _ = v31
    }
    v32 := GenS2{F0: 899}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v32
    _ = v4
}
