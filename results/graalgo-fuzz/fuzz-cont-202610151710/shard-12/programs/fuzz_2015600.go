package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 uint32; F1 string; F2 string }
func genFn0(p0 uint32) int64 {
    v5 := []int{}
    v5 = append(v5, 4294967295)
    _ = len(v5)
    _ = p0
    _ = v5
    return -825
}
func genFn1(p0 bool) bool {
    v6 := []int{-2147483648, 127, 0}
    p0 = p0
    var v7 float32 = 0.0
    for fi8 := 0; fi8 < 4; fi8++ {
        p0 = (65535 >= 32767)
        _ = fi8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float32 = float32(0.1)
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    return (!p0)
}

func main() {
    var v1 uint = 42
    var v2 uint = (v1 << (((v1 | 2)) & 63))
    var v3 int32 = -367
    v4 := make(map[string]uint32)
    fmt.Println(v2, v2)
    v11 := []int{2147483647}
    if ("\n\t" != "a") {
        v3 = v3
    } else {
        if false {
            for fi12 := 0; fi12 < 5; fi12++ {
                if (v3 >= (v3 / ((-2147483648) | 1))) {
                    for fi13 := 0; fi13 < 9; fi13++ {
                        var v14 int64 = 42
                        _ = fi13
                        _ = v14
                    }
                    fmt.Println(v3, v1, v3)
                    v2 += v1
                } else {
                    fmt.Println(fi12, v2)
                    switch 32767 {
                    case 684:
                        v15 := v4["hello"]
                        _ = v15
                    case 463:
                        var v16 int64 = 0
                        _ = v16
                    case -717:
                        var v17 int = 127
                        _ = v17
                    default:
                        var v18 uint16 = 148
                        _ = v18
                    }
                    _ = len(v4)
                }
                var v19 int8 = 33
                v20 := GenS1{F0: (fi12 ^ fi12)}
                _ = fi12
                _ = v19
                _ = v20
            }
            var v21 uint32 = 1
            v1 -= (^(v1))
            _ = v21
        } else {
            fmt.Println(v1)
        }
        if true {
            v22 := GenS1{F0: -9223372036854775808}
            switch "a" {
            case "foo bar":
                for fi23 := 0; fi23 < 9; fi23++ {
                    if true {
                        if true {
                            var v24 uint = v1
                            _ = v24
                        }
                    }
                    v1 = v1
                    _ = fi23
                }
            case "\n\t":
                v3 = (v3 ^ v3)
            default:
                v25 := GenS1{F0: 9223372036854775807}
                _ = v25
            }
            v26 := GenS2{F0: 1, F1: -463, F2: -1}
            _ = v22
            _ = v26
        } else {
            if (true || ((0.0 * 0.0) < (((-2.25 * (3.14159 * 0.0)) / (0.1 * 3.14159)) / 1.5))) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                v1 = 65535
                v1 = v2
            } else {
                if len(v11) > 0 {
                    v11[0] = -41
                }
            }
        }
    }
    switch 801 {
    case 127:
        if (0 == -128) {
            fmt.Println(v2, v1, v3)
            var v27 bool = (((^(v1)) == (v2 >> ((v1) & 63))) && true)
            for fi28 := 0; fi28 < 13; fi28++ {
                var v29 float64 = -2.25
                v29 = (float64(1.5) + v29)
                _ = fi28
                _ = v29
            }
            _ = v27
        }
    case 9223372036854775807:
        v3 = 866
    case -9223372036854775808:
        for fi30 := 0; fi30 < 19; fi30++ {
            v1 += v1
            v2 = (v1 >> (((^(v1))) & 63))
            v1 += 1
            _ = fi30
        }
    default:
        switch "hello" {
        case "":
            v11 = append(v11, 1)
        case "\n\t":
            var v31 int16 = -1
            _ = v31
        case "a":
            v32 := GenS1{F0: 496}
            _ = v32
        default:
            v1 += (v2 - (v1 / ((v1) | 1)))
        }
    }
    v33 := GenS2{F0: 943, F1: -32768, F2: (v3 & (v3 / ((-1) | 1)))}
    v1 = v2
    v3 -= (v3 % ((42) | 1))
    v11 = append(v11, -522)
    var v34 float64 = ((1e10 + -2.25) - 0.0)
    v35 := GenS3{F0: 4294967295, F1: "\n\t", F2: "a"}
    v36 := GenS1{F0: 0}
    var v37 bool = (929 > 65535)
    v3 -= (^(v3))
    var v38 float32 = 3.14159
    v2 = v1
    v38 += ((float32(v34) / (v38 + 1e10)) - -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
