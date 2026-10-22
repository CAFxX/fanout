package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int64; F1 int; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 uint32 }
func genFn0(p0 int8) uint8 {
    p0 += (-(p0))
    var v5 int64 = -9223372036854775808
    if false {
        if (-1 == 1) {
            v6 := []int{-491, 4294967295}
            _ = v6
        }
        fmt.Println(v5, v5)
        v7 := GenS1{F0: -9223372036854775808, F1: 128}
        _ = v7
    }
    var v8 int8 = p0
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 = (v5 & 9223372036854775807)
        _ = fi9
    }
    _ = p0
    _ = v5
    _ = v8
    return 0
}
func genFn1(p0 int, p1 bool, p2 uint) int16 {
    v10 := GenS1{F0: p0, F1: 155}
    v11 := GenS2{F0: -1, F1: (p0 >> ((18446744073709551615) & 63)), F2: (p0 / ((p0) | 1))}
    var v12 int64 = 42
    v13 := GenS2{F0: v12, F1: (p0 << ((p2) & 63)), F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{"hello", "a", (("a" + "\n\t") + "\n\t"), (("\n\t" + (("" + "hello") + "hello")) + ("a" + ("" + "foo bar")))}
    var v3 float32 = -2.25
    v4 := make(map[int]int)
    var v14 int = 127
    v3 = (v3 + 1e10)
    switch ("" + ("a" + "a")) {
    case "foo bar":
        _ = len(v2)
    case "hello":
        v3 = -2.25
    case "":
        switch v14 {
        case 42:
            var v15 uint = 0
            _ = v15
        case -9223372036854775808:
            var v16 string = ("foo bar" + "hello")
            _ = v16
        case 2147483647:
            var v17 float64 = 0.1
            _ = v17
        }
    }
    var v18 float32 = v3
    for fi19 := 0; fi19 < 20; fi19++ {
        v20 := GenS3{F0: ("" + "")}
        v3 += v3
        var v21 int32 = 2147483647
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := v4[v14]
    v23 := GenS4{F0: -480}
    var v24 uint8 = genFn0(int8(v22))
    var v25 uint = 255
    v22 += (v14 << ((686) & 63))
    for fi26 := 0; fi26 < 19; fi26++ {
        if (42 >= 415) {
            var v27 uint32 = 2
            for fi28 := 0; fi28 < 11; fi28++ {
                switch "" {
                case "\n\t":
                    v22 = (v22 & v22)
                case "foo bar":
                    var v29 string = ("\n\t" + ("" + ("\n\t" + "\n\t")))
                    _ = v29
                default:
                    switch "hello" {
                    case "hello":
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    case "\n\t":
                        v18 = ((-2.25 + 0.0) - 0.0)
                    }
                }
                _ = fi28
            }
            _ = v27
        }
        v2 = append(v2, ("" + ""))
        v30 := GenS5{F0: v14, F1: 2147483647}
        _ = fi26
        _ = v30
    }
    v2 = append(v2, "hello")
    v31 := v1[42]
    fmt.Println(v24)
    var v32 int32 = -84
    if false {
        for fi33 := 0; fi33 < 6; fi33++ {
            v24 = (v24 + (v24 * v24))
            _ = fi33
        }
        v3 = (v18 / (float32(0.1) / v18))
    } else {
        v14 = 2147483647
        fmt.Println(v31, v22, v22)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
