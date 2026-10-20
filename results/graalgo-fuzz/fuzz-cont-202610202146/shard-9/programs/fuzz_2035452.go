package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 bool; F1 string; F2 bool }
func genFn0(p0 int16, p1 uint8, p2 string) int16 {
    switch p2 {
    case "\n\t":
        if true {
            fmt.Println(p2, p2, p0)
            v7 := GenS1{F0: -9223372036854775808, F1: 42, F2: (p1 >> ((95) & 63))}
            var v8 float32 = -2.25
            _ = v7
            _ = v8
        }
    }
    v9 := []int{int(p1), 32768, int(p0)}
    var v10 float64 = (0.1 + (((1.5 * 1e10) / 0.1) * 0.0))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 2
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    var v5 uint = 42
    var v6 int16 = 0
    for fi11 := 0; fi11 < 20; fi11++ {
        if (false || (v6 <= (v6 - 42))) {
            var v12 uint32 = 799
            fmt.Println(v2, v5, v12)
            for fi13 := 0; fi13 < 16; fi13++ {
                if (0.1 == (3.14159 - (-2.25 - 0.1))) {
                    fmt.Println(v6, v6, v2)
                    switch "hello" {
                    case "foo bar":
                        switch fi13 {
                        case 32768:
                            switch ("a" + "foo bar") {
                            case "hello":
                                var v14 uint16 = 2
                                _ = v14
                            case "foo bar":
                                fmt.Println(v12)
                            default:
                                v15 := []int{(fi13 % ((-594) | 1)), fi13}
                                _ = v15
                            }
                        case 630:
                            fmt.Println(v5, fi13, v12)
                        }
                    case "\n\t":
                        v16 := GenS1{F0: 0, F1: 42, F2: 255}
                        _ = v16
                    case "":
                        v6 = genFn0(v6, 42, "a")
                    }
                } else {
                    fmt.Println(v12, v6, fi11)
                    v2 = v5
                    switch -224 {
                    case -9223372036854775808:
                        v6 = v6
                    }
                }
                _ = fi13
            }
            _ = v12
        }
        for fi17 := 0; fi17 < 15; fi17++ {
            if true {
                v5 = (v5 % (((^(v5))) | 1))
            } else {
                if true {
                    v2 = v5
                    v6 -= -227
                } else {
                    var v18 uint32 = 341
                    _ = v18
                }
            }
            var v19 string = "\n\t"
            v5 -= 305
            v20 := GenS2{F0: (v6 & v6), F1: 2, F2: 210}
            v6 -= -305
            _ = fi17
            _ = v19
            _ = v20
        }
        _ = fi11
    }
    fmt.Println(v6, v6, v6)
    v21 := GenS3{F0: 18446744073709551615, F1: 182, F2: -61}
    switch ("a" + "hello") {
    case "\n\t":
        var v22 int8 = 42
        _ = v22
    default:
        v2 += 919
    }
    switch 175 {
    case 127:
        var v23 int16 = v6
        _ = v23
    case 32767:
        if (44 <= 4294967295) {
            fmt.Println(v5, v5)
        }
    case 0:
        v24 := GenS1{F0: -625, F1: 1, F2: 2}
        _ = v24
    default:
        v2 -= v2
    }
    switch (("" + "a") + ("" + "\n\t")) {
    case "foo bar":
        var v25 int8 = 66
        _ = v25
    }
    v26 := []int{920}
    switch -1 {
    case 32768:
        fmt.Println(v6)
    case -546:
        v2 = 689
    default:
        v6 = (v6 ^ (v6 / ((v6) | 1)))
    }
    v27 := GenS4{F0: (1.5 != -2.25), F1: "hello", F2: false}
    v5 = v2
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(fi28)
        for fi29 := 0; fi29 < 8; fi29++ {
            v5 = v2
            v30 := GenS4{F0: (!false), F1: ("\n\t" + "a"), F2: (!(!true))}
            var v31 int8 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
        v6 = genFn0((v6 % (((v6 * int16(fi28))) | 1)), 0, ("\n\t" + "\n\t"))
        _ = fi28
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        v33 := GenS3{F0: 1, F1: 2147483647, F2: -59}
        _ = fi32
        _ = v33
    }
    v2 = (v5 >> (((^(v2))) & 63))
    v2 -= (v5 & v2)
    v6 -= v6
    v2 += (v2 ^ v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
