package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0(p0 uint) int8 {
    var v4 int16 = -800
    var v5 uint = (p0 / (((^(p0))) | 1))
    v6 := []int{772, 453}
    var v7 uint64 = 18446744073709551615
    v7 += v7
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 53
}
func genFn1(p0 int) int8 {
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 = (fi8 - fi8)
        _ = fi8
    }
    var v9 int8 = 0
    _ = p0
    _ = v9
    return (v9 + genFn0(288))
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 18446744073709551615
    v3 := make(map[uint32]bool)
    if (737 >= 4294967295) {
        var v10 int32 = 2147483647
        for fi11 := 0; fi11 < 20; fi11++ {
            v10 = (v10 | v10)
            v2 = v2
            _ = fi11
        }
        _ = v10
    } else {
        switch "\n\t" {
        case "\n\t":
            v1 = append(v1, 869)
        case "a":
            if true {
                fmt.Println(v2)
                v12 := GenS1{F0: "a"}
                if false {
                    v13 := GenS1{F0: ""}
                    _ = v13
                } else {
                    v14 := GenS1{F0: "foo bar"}
                    _ = v14
                }
                _ = v12
            } else {
                v2 += 0
                fmt.Println(v2)
                var v15 uint16 = 2
                _ = v15
            }
        case "foo bar":
            if true {
                if true {
                    if (-32768 <= 32767) {
                        var v16 uint64 = (v2 ^ 721)
                        var v17 int32 = 42
                        if (true && (!false)) {
                            v16 -= 18446744073709551615
                            var v18 uint32 = 2147483648
                            var v19 int = -1
                            _ = v18
                            _ = v19
                        } else {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                var v21 uint64 = 42
                                v17 = v17
                                _ = fi20
                                _ = v21
                            }
                        }
                        _ = v16
                        _ = v17
                    } else {
                        fmt.Println(v2)
                    }
                    fmt.Println(v2)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v2 -= v2
                        v23 := v3[62]
                        for fi24 := 0; fi24 < 14; fi24++ {
                            v25 := GenS2{F0: -1, F1: 17, F2: 42}
                            _ = len(v3)
                            _ = fi24
                            _ = v25
                        }
                        _ = fi22
                        _ = v23
                    }
                }
                var v26 int32 = 0
                for fi27 := 0; fi27 < 5; fi27++ {
                    v26 = (v26 << ((65535) & 63))
                    _ = fi27
                }
                _ = v26
            }
        default:
            fmt.Println(v2)
        }
    }
    var v28 bool = ((0.0 <= float32(-2.25)) && false)
    _ = len(v1)
    var v29 bool = true
    var v30 float32 = 0.1
    switch 128 {
    case -774:
        var v31 int = 4294967295
        _ = v31
    case 127:
        v32 := GenS2{F0: -84, F1: 128, F2: 689}
        _ = v32
    default:
        for fi33 := 0; fi33 < 13; fi33++ {
            v30 -= v30
            _ = fi33
        }
    }
    v29 = (int64(v2) <= -1)
    v2 = v2
    v28 = v29
    var v34 uint8 = 42
    var v35 bool = (832 < 628)
    v36 := GenS3{F0: 279, F1: (v2 | v2), F2: 1}
    v35 = v35
    v37 := GenS4{F0: -793}
    var v38 uint32 = 557
    v39 := GenS1{F0: "foo bar"}
    v30 = 1.5
    if (127 > (v2 * v2)) {
        v28 = (-32768 >= 32767)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
}
