package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 uint8 }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint, p1 uint16, p2 int32) int8 {
    v7 := []int{32768, 9223372036854775807, -976}
    p0 = p0
    v8 := GenS1{F0: 2, F1: 4294967295}
    v9 := GenS1{F0: 255, F1: 775}
    p0 = 945
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 20
}

func main() {
    var v1 int16 = -40
    var v2 int16 = v1
    var v3 int16 = v2
    var v4 uint64 = 0
    var v5 float32 = 1e10
    var v6 bool = (160 < 0)
    switch 1 {
    case 32768:
        v1 -= (v2 / (((^(v3))) | 1))
    case 1:
        v10 := []int{9223372036854775807}
        _ = v10
    case -9223372036854775808:
        var v11 int16 = (v2 * v3)
        _ = v11
    default:
        v12 := GenS2{F0: 65535, F1: 2}
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 bool = true
        v2 = v1
        _ = fi13
        _ = v14
    }
    var v15 int8 = 0
    v16 := []int{-1, -711, -2147483648}
    var v17 uint32 = 1
    v17 = v17
    switch "hello" {
    case "a":
        if false {
            for fi18 := 0; fi18 < 9; fi18++ {
                v4 += v4
                _ = fi18
            }
        }
    default:
        if true {
            if (v6 || v6) {
                v19 := GenS1{F0: 127, F1: v4}
                if (!false) {
                    v2 -= -1
                    if len(v16) > 0 {
                        v16[0] = 128
                    }
                    switch ("\n\t" + "") {
                    case "hello":
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v21 := GenS3{F0: (v1 * 1000), F1: "foo bar", F2: 127}
                            for fi22 := 0; fi22 < 9; fi22++ {
                                var v23 string = ("\n\t" + ("" + ""))
                                _ = fi22
                                _ = v23
                            }
                            _ = fi20
                            _ = v21
                        }
                    case "foo bar":
                        var v24 int64 = -9223372036854775808
                        _ = v24
                    case "\n\t":
                        fmt.Println(v4, v4)
                    default:
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    }
                } else {
                    v5 -= 0.1
                    v1 -= (v1 & (v1 * v1))
                    v3 = v3
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    if v6 {
                        fmt.Println(v6, v4)
                        v3 = (v1 << ((uint(v4)) & 63))
                        fmt.Println(v4, v4, v4)
                    }
                    switch "foo bar" {
                    case "\n\t":
                        v26 := GenS1{F0: 169, F1: uint64(v17)}
                        _ = v26
                    case "a":
                        var v27 int8 = v15
                        _ = v27
                    case "":
                        v1 = v1
                    default:
                        if len(v16) > 0 {
                            v16[0] = 32767
                        }
                    }
                    _ = fi25
                }
                _ = v19
            } else {
                v6 = (1 > genFn0(1, 65535, 1))
                v28 := GenS4{F0: "", F1: "a"}
                _ = len(v16)
                _ = v28
            }
        } else {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    }
    switch "foo bar" {
    case "foo bar":
        v29 := GenS5{F0: uint8(v17)}
        _ = v29
    default:
        if (genFn0(218, 65535, -1) <= v15) {
            fmt.Println(v17, v15, v6)
            if (((v15 | (v15 - (-(v15)))) < (v15 - v15)) || true) {
                var v30 int64 = -9223372036854775808
                v3 += v1
                v31 := GenS6{F0: genFn0(255, 42, 42)}
                _ = v30
                _ = v31
            }
        } else {
            var v32 int64 = -9223372036854775808
            for fi33 := 0; fi33 < 6; fi33++ {
                if len(v16) > 0 {
                    v16[0] = fi33
                }
                v34 := GenS7{F0: v32, F1: v17}
                _ = fi33
                _ = v34
            }
            if len(v16) > 0 {
                v16[0] = 4294967295
            }
            _ = v32
        }
    }
    _ = len(v16)
    if true {
        if len(v16) > 0 {
            v16[0] = 127
        }
        if len(v16) > 0 {
            v16[0] = -1
        }
    }
    v35 := GenS3{F0: (v3 ^ (v2 % (((-(v2))) | 1))), F1: "foo bar", F2: 249}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
