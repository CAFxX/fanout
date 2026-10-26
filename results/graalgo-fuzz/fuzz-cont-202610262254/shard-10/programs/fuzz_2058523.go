package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 uint64; F2 uint16 }
func genFn0(p0 uint32, p1 uint64) int32 {
    var v6 int16 = 32767
    var v7 uint32 = 65535
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = -24
    var v2 uint8 = 128
    var v3 uint32 = 4294967295
    var v4 uint = 385
    var v5 uint16 = 255
    if (-78 < -128) {
        for fi8 := 0; fi8 < 17; fi8++ {
            v4 += 42
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(fi8, v2)
                _ = fi9
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint16 = (v5 * (v5 - (v5 - 32767)))
                v4 = v4
                _ = fi10
                _ = v11
            }
            _ = fi8
        }
    }
    var v12 int64 = -9223372036854775808
    v13 := GenS1{F0: (v4 % (((v4 >> ((v4) & 63))) | 1)), F1: (v2 * v2), F2: false}
    if (2 < 1) {
        v14 := GenS2{F0: v12}
        var v15 int8 = -128
        _ = v14
        _ = v15
    } else {
        var v16 uint64 = 2
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        switch -9223372036854775808 {
        case 0:
            if (!(!false)) {
                v18 := GenS2{F0: v12}
                _ = v18
            } else {
                if (("hello" != ("" + "hello")) && ((v2 > 127) || (42 >= 127))) {
                    v19 := GenS1{F0: v4, F1: (v2 & (v2 + v2)), F2: (-2.25 >= float32(((1.5 + 0.1) + -2.25)))}
                    fmt.Println(v2)
                    _ = v19
                } else {
                    v20 := []int{}
                    v1 = (v1 & 891)
                    switch 42 {
                    case 553:
                        v21 := GenS3{F0: (v2 % ((v2) | 1))}
                        _ = v21
                    case -1:
                        var v22 string = ""
                        _ = v22
                    case 9223372036854775807:
                        switch "a" {
                        case "\n\t":
                            if (-32768 > 528) {
                                v23 := GenS2{F0: (^(v12))}
                                _ = v23
                            }
                        }
                    }
                    _ = v20
                }
                var v24 int16 = 42
                _ = v24
            }
        default:
            if false {
                v2 = (v2 >> ((2) & 63))
            } else {
                for fi25 := 0; fi25 < 9; fi25++ {
                    v12 = v12
                    var v26 uint8 = (v2 | (v2 ^ (v2 << ((uint(v2)) & 63))))
                    var v27 bool = true
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        }
        v28 := GenS4{F0: "foo bar", F1: 2, F2: (v5 + v5)}
        _ = fi17
        _ = v28
    }
    var v29 int64 = v12
    v30 := []int{-556, -1, 128, 1}
    fmt.Println(v3, v4)
    for fi31 := 0; fi31 < 11; fi31++ {
        if true {
            if ((!((v4 << (((^(v4))) & 63)) != v4)) && (!false)) {
                if false {
                    var v32 int64 = (v29 & -9223372036854775808)
                    for fi33 := 0; fi33 < 5; fi33++ {
                        v3 = v3
                        v29 += v29
                        v34 := GenS4{F0: (("" + "hello") + "hello"), F1: 101, F2: v5}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                } else {
                    for fi35 := 0; fi35 < 10; fi35++ {
                        fmt.Println(v2, v2, fi35)
                        v12 += v12
                        var v36 int32 = -2147483648
                        v37 := GenS3{F0: 218}
                        v36 = genFn0(v3, 4294967295)
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                    var v38 bool = (v5 == v5)
                    _ = v38
                }
                v39 := GenS3{F0: uint8(v3)}
                _ = v39
            }
        } else {
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
        }
        _ = fi31
    }
    fmt.Println(v4)
    if (true || true) {
        v3 += (^(v3))
        _ = len(v30)
        v5 -= v5
    }
    var v40 uint8 = (v2 ^ v2)
    var v41 int8 = -107
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
