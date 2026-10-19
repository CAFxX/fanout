package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 int32) bool {
    var v4 uint64 = 0
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(p0, v4)
        _ = fi5
    }
    if (255 <= 1) {
        switch "a" {
        case "\n\t":
            p0 = p0
        case "foo bar":
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: "hello", F1: ((((-2.25 >= 1.5) || (!(fi6 == (-(fi6))))) || false) && ((p0 >= (p0 * (p0 >> ((73) & 63)))) && true))}
                var v8 int32 = (p0 % ((42) | 1))
                switch 673 {
                case 32768:
                    if true {
                        for fi9 := 0; fi9 < 14; fi9++ {
                            v8 = (v8 % ((v8) | 1))
                            _ = fi9
                        }
                        v4 = v4
                    } else {
                        var v10 uint8 = 255
                        fmt.Println(v4, v10, v4)
                        _ = v10
                    }
                default:
                    var v11 string = "foo bar"
                    _ = v11
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        case "":
            p0 += (p0 + (p0 ^ p0))
        }
        if true {
            for fi12 := 0; fi12 < 7; fi12++ {
                v4 += 2
                _ = fi12
            }
        }
    } else {
        for fi13 := 0; fi13 < 3; fi13++ {
            v4 += v4
            v14 := GenS2{F0: -1}
            var v15 uint32 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v4 = (v4 / (((v4 & (v4 | (v4 * v4)))) | 1))
    if (!true) {
        p0 += (p0 * (p0 * p0))
        fmt.Println(p0, v4)
    }
    _ = p0
    _ = v4
    return (!false)
}
func genFn1(p0 uint32, p1 uint, p2 int32) int64 {
    v16 := GenS3{F0: 509, F1: (p2 >> (((^(p1))) & 63))}
    v17 := GenS1{F0: "foo bar", F1: ((2147483647 != 787) || true)}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return -9223372036854775808
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]string)
    v3 := []uint32{2, 561, 813, 0}
    v1 = v1
    v1 -= v1
    v1 += genFn1(59, 615, 42)
    v1 = v1
    v18 := GenS3{F0: 745, F1: 962}
    fmt.Println(v1)
    var v19 uint64 = 217
    var v20 int64 = (v1 ^ (v1 >> ((0) & 63)))
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 uint8 = 42
        if false {
            var v23 int8 = int8(v19)
            v19 = 504
            switch (fi21 & fi21) {
            case 32767:
                if ((!false) || (595 < 197)) {
                    v22 -= (v22 * v22)
                    v24 := GenS4{F0: (v1 / ((1) | 1)), F1: "foo bar"}
                    var v25 float64 = 0.1
                    _ = v24
                    _ = v25
                }
            default:
                switch "foo bar" {
                case "foo bar":
                    v23 -= (v23 >> ((1) & 63))
                case "hello":
                    if ((v19 >= 662) || ((!(42 >= -691)) || true)) {
                        var v26 float64 = ((-2.25 * (0.1 - -2.25)) - (3.14159 / 3.14159))
                        _ = v26
                    } else {
                        switch (("hello" + "a") + "\n\t") {
                        case "":
                            fmt.Println(v20)
                        }
                        v23 += -128
                    }
                case "a":
                    fmt.Println(v20, v1, v20)
                }
            }
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v1 = 1
    var v27 int16 = 32767
    var v28 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v3
}
