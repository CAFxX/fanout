package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 int }
type GenS4 struct { F0 int64; F1 string; F2 string }
type GenS5 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 uint) uint64 {
    v5 := GenS1{F0: 42, F1: 55, F2: 2}
    var v6 int = -2147483648
    fmt.Println(v6, p0)
    _ = p0
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -1
    v3 := make(map[uint32]uint32)
    var v4 uint = 1
    v2 -= v2
    var v7 float32 = 0.1
    var v8 uint64 = genFn0(v4)
    var v9 float32 = (0.0 * v7)
    if ((0 <= 1000) && ((true && (!true)) || false)) {
        for fi10 := 0; fi10 < 5; fi10++ {
            fmt.Println(v2, v4)
            fmt.Println(fi10)
            _ = fi10
        }
        v11 := []int{}
        var v12 int = -829
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if (false && true) {
            v9 += v7
            fmt.Println(v1, fi13, v2)
            v9 -= 0.0
        } else {
            v9 += float32(v1)
            var v14 int = 4294967295
            _ = v14
        }
        _ = fi13
    }
    switch "hello" {
    case "\n\t":
        if (!true) {
            v15 := GenS1{F0: 152, F1: 127, F2: 368}
            v2 -= v2
            _ = v15
        }
    case "foo bar":
        _ = len(v3)
    default:
        fmt.Println(v9)
    }
    var v16 int32 = (v2 % (((^(v2))) | 1))
    v3[1] = 2147483648
    switch 0 {
    case 1:
        v4 -= 65535
    case 254:
        if ((true && (true && true)) || false) {
            v17 := GenS2{F0: 108, F1: v8, F2: 633}
            _ = v17
        } else {
            v3[65535] = 4294967295
            _ = len(v3)
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS3{F0: 255, F1: 127, F2: (^(fi18))}
                fmt.Println(v7, v7)
                _ = fi18
                _ = v19
            }
        }
    case -2147483648:
        for fi20 := 0; fi20 < 10; fi20++ {
            if false {
                v1 -= v1
                var v21 int32 = v16
                if (((v4 << ((v4) & 63)) >= (v4 + (v4 - v4))) && (229 <= v21)) {
                    var v22 int16 = -403
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v24 := GenS2{F0: 128, F1: 1, F2: 791}
                        if (!true) {
                            var v25 int8 = -122
                            v26 := v3[2147483647]
                            if ((v26 % (((v26 / ((2) | 1))) | 1)) < v26) {
                                if false {
                                    var v27 bool = (uint16(v8) != 497)
                                    _ = v27
                                } else {
                                    v28 := GenS4{F0: 1, F1: "hello", F2: ((("hello" + (((("\n\t" + "a") + "") + "foo bar") + ("\n\t" + "foo bar"))) + "foo bar") + "a")}
                                    v1 = (v1 + (-2.25 + v1))
                                    v29 := GenS2{F0: 120, F1: 2, F2: 65535}
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(fi20)
                            }
                            _ = v25
                            _ = v26
                        }
                        var v30 string = ""
                        _ = fi23
                        _ = v24
                        _ = v30
                    }
                    _ = v22
                }
                _ = v21
            }
            v16 = (v2 & v16)
            var v31 uint8 = 42
            _ = fi20
            _ = v31
        }
    default:
        v7 -= v7
    }
    v2 += v16
    var v32 uint16 = 610
    v4 -= v4
    var v33 int64 = 167
    v9 -= (-2.25 / -2.25)
    var v34 uint8 = 127
    v35 := GenS3{F0: v32, F1: v34, F2: 32767}
    v36 := GenS5{F0: -9223372036854775808, F1: 42, F2: 9223372036854775807}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
