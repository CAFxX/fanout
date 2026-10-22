package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 bool; F2 string }
type GenS3 struct { F0 uint32; F1 int64 }
func genFn0(p0 int32) uint8 {
    v4 := []int{558, -667, 42}
    p0 += -983
    _ = p0
    _ = v4
    return 127
}
func genFn1(p0 uint8) int64 {
    for fi5 := 0; fi5 < 16; fi5++ {
        switch "a" {
        case "a":
            p0 = (p0 | (p0 & (p0 & p0)))
        default:
            var v6 uint32 = 0
            _ = v6
        }
        v7 := []int{-2147483648, 262, (fi5 / ((fi5) | 1))}
        _ = fi5
        _ = v7
    }
    p0 = 255
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := []int{fi8, 2147483647, fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    return -1
}

func main() {
    v1 := []int{-763, 32768}
    var v2 uint = 288
    v3 := []string{("foo bar" + ("" + ("" + "\n\t")))}
    if (false || (-2147483648 <= -960)) {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: 255, F1: 720}
            if (false && true) {
                v2 -= (v2 >> (((v2 | v2)) & 63))
                v3 = append(v3, "")
                var v12 string = (("hello" + (((("\n\t" + ("a" + ("hello" + "foo bar"))) + "\n\t") + "hello") + "a")) + "\n\t")
                _ = v12
            } else {
                v13 := GenS1{F0: 1, F1: -1}
                var v14 uint64 = 974
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
    }
    v2 -= v2
    if ((0 != 185) || false) {
        fmt.Println(v2)
        for fi15 := 0; fi15 < 6; fi15++ {
            switch ("a" + "hello") {
            case "a":
                v1 = append(v1, (^(fi15)))
            default:
                var v16 uint64 = 42
                _ = v16
            }
            var v17 uint8 = 32
            var v18 uint32 = 2
            _ = fi15
            _ = v17
            _ = v18
        }
    }
    v19 := GenS1{F0: 42, F1: 1}
    switch 0 {
    case -299:
        switch ((("\n\t" + "a") + ((((("\n\t" + "a") + "foo bar") + "\n\t") + "hello") + "a")) + "foo bar") {
        case "foo bar":
            for fi20 := 0; fi20 < 8; fi20++ {
                if len(v1) > 0 {
                    v1[0] = fi20
                }
                _ = fi20
            }
        case "\n\t":
            fmt.Println(v2)
        case "":
            var v21 float64 = -2.25
            _ = v21
        default:
            var v22 uint8 = 178
            _ = v22
        }
    case 4294967295:
        for fi23 := 0; fi23 < 12; fi23++ {
            v2 -= v2
            v2 = v2
            _ = fi23
        }
    }
    v24 := GenS2{F0: "hello", F1: (1e10 < 1.5), F2: "a"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v25 := GenS3{F0: 20, F1: 576}
    var v26 uint = 74
    if (false && true) {
        v2 = v2
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 int8 = 0
        _ = fi27
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
