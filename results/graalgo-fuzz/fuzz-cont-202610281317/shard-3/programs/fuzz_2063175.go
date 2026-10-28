package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }
func genFn0() int {
    var v7 int8 = -128
    v7 += v7
    v7 = 1
    switch 2147483647 {
    case 0:
        fmt.Println(v7)
    case -9223372036854775808:
        v7 += v7
    case 42:
        v8 := GenS1{F0: 1, F1: -2147483648}
        _ = v8
    default:
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{0, fi9, -2147483648}
            var v11 int64 = -9223372036854775808
            if false {
                switch fi9 {
                case -961:
                    v7 += int8(v11)
                case 128:
                    v12 := GenS1{F0: 42, F1: 42}
                    _ = v12
                default:
                    if (v7 < (v7 ^ v7)) {
                        var v13 uint16 = 360
                        v10 = append(v10, (fi9 | (fi9 % (((fi9 + (fi9 / ((int(v11)) | 1)))) | 1))))
                        _ = v13
                    } else {
                        if true {
                            v14 := GenS2{F0: 446, F1: v7}
                            v11 = (v11 & (v11 ^ (v11 << ((609) & 63))))
                            v11 = (v11 | v11)
                            _ = v14
                        } else {
                            if true {
                                v11 = (v11 ^ (v11 | 9223372036854775807))
                                for fi15 := 0; fi15 < 7; fi15++ {
                                    v16 := GenS2{F0: 4294967295, F1: (-(v7))}
                                    var v17 int32 = 1
                                    var v18 uint32 = 0
                                    _ = fi15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                }
                            } else {
                                var v19 int16 = 247
                                _ = v19
                            }
                            v20 := GenS1{F0: 2, F1: 2147483647}
                            fmt.Println(v11)
                            _ = v20
                        }
                        v21 := GenS3{F0: 2147483648}
                        _ = v21
                    }
                }
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        switch "a" {
        case "\n\t":
            v7 += (v7 * v7)
        default:
            v7 += (v7 % ((100) | 1))
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            if (((-2.25 * (-2.25 * 0.0)) - (float64((0.1 / 1.0)) * 1.5)) >= (-2.25 + 1.5)) {
                var v24 uint16 = 42
                _ = v24
            }
            _ = fi23
        }
        _ = fi22
    }
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint8 = 1
    var v2 int16 = -32768
    var v3 int16 = v2
    var v4 int8 = -1
    var v5 bool = ("foo bar" < "\n\t")
    v6 := make(map[uint32]int)
    fmt.Println(v2, v5, v4)
    v25 := []int{-261, 1}
    switch 905 {
    case 887:
        if len(v25) >= 2 {
            v25 = v25[1:len(v25)]
        }
    case 4294967295:
        v26 := GenS1{F0: 113, F1: -2147483648}
        _ = v26
    default:
        if v5 {
            v2 += (v3 ^ v3)
            for fi27 := 0; fi27 < 15; fi27++ {
                switch "\n\t" {
                case "foo bar":
                    for fi28 := 0; fi28 < 2; fi28++ {
                        if len(v25) > 0 {
                            v25[0] = fi27
                        }
                        if ((v1 % (((v1 - (v1 + v1))) | 1)) <= (v1 ^ (v1 >> ((255) & 63)))) {
                            v29 := GenS2{F0: 1, F1: (v4 & v4)}
                            _ = v29
                        }
                        _ = fi28
                    }
                case "hello":
                    v30 := v6[1]
                    _ = v30
                case "":
                    switch ("a" + (("a" + (("" + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t"))) + "hello")) {
                    case "":
                        var v31 uint32 = 805
                        _ = v31
                    default:
                        fmt.Println(v3, fi27)
                    }
                default:
                    fmt.Println(v3, v3)
                }
                _ = fi27
            }
            v6[2147483648] = genFn0()
        }
    }
    v32 := GenS4{F0: 2}
    if v5 {
        for fi33 := 0; fi33 < 7; fi33++ {
            var v34 uint = 4294967295
            _ = fi33
            _ = v34
        }
        var v35 uint32 = 2147483648
        _ = v35
    }
    v36 := GenS4{F0: 228}
    delete(v6, 284)
    _ = len(v25)
    var v37 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
