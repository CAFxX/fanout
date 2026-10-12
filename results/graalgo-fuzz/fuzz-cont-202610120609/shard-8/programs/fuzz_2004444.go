package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 string }
func genFn0() bool {
    var v7 uint64 = 42
    if false {
        v8 := []int{128}
        switch "hello" {
        case "":
            v7 = v7
        case "\n\t":
            _ = len(v8)
        }
        if (!false) {
            _ = len(v8)
            if (109 < 127) {
                if true {
                    v7 += (^(v7))
                    v7 = (v7 >> ((4294967295) & 63))
                    v8 = append(v8, 32768)
                }
                if (!false) {
                    var v9 bool = (!(68 == 127))
                    fmt.Println(v9, v9)
                    if false {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            for fi11 := 0; fi11 < 8; fi11++ {
                                fmt.Println(v7, v7, v9)
                                _ = fi11
                            }
                            _ = fi10
                        }
                        fmt.Println(v9, v7)
                        if len(v8) > 0 {
                            v8[0] = -229
                        }
                    }
                    _ = v9
                } else {
                    v7 -= 104
                    switch 32767 {
                    case 42:
                        if len(v8) > 0 {
                            v8[0] = 4294967295
                        }
                    case 127:
                        var v12 int16 = 32767
                        _ = v12
                    case -677:
                        v7 = 2
                    default:
                        v7 = (v7 / (((v7 ^ (v7 - v7))) | 1))
                    }
                }
            }
        } else {
            v13 := GenS1{F0: 0}
            _ = v13
        }
        _ = v8
    } else {
        switch ("a" + "\n\t") {
        case "foo bar":
            v7 = 71
        case "\n\t":
            v14 := []int{32768, 127, 9223372036854775807, 602}
            _ = v14
        default:
            v15 := GenS1{F0: 2}
            _ = v15
        }
    }
    if (-128 != 84) {
        switch "hello" {
        case "":
            for fi16 := 0; fi16 < 13; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v18 := []int{-917, (-(fi16))}
                    _ = fi17
                    _ = v18
                }
                v7 -= v7
                _ = fi16
            }
        case "hello":
            v7 = (v7 - 4294967295)
        default:
            v7 -= v7
        }
    }
    _ = v7
    return (!(2 != 398))
}

func main() {
    var v1 int64 = 1
    var v2 int32 = -1
    var v3 uint32 = 635
    var v4 float64 = 1.5
    var v5 uint32 = v3
    v6 := []string{"hello", "foo bar", "hello"}
    v4 += (((v4 / (v4 * 1e10)) + 1e10) / v4)
    _ = len(v6)
    var v19 int64 = v1
    var v20 uint64 = 18446744073709551615
    if (-996 > 127) {
        if (-2147483648 != 9223372036854775807) {
            v21 := GenS2{F0: -311, F1: v1}
            var v22 int64 = 397
            v23 := GenS3{F0: true}
            _ = v21
            _ = v22
            _ = v23
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    _ = len(v6)
    for fi24 := 0; fi24 < 1; fi24++ {
        v2 += (v2 + int32(v1))
        var v25 int32 = v2
        var v26 int16 = 1000
        _ = fi24
        _ = v25
        _ = v26
    }
    v5 += (v5 ^ v5)
    v27 := GenS2{F0: 127, F1: (v1 - v1)}
    v28 := GenS4{F0: "hello", F1: ("a" + "hello")}
    var v29 uint = uint(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
