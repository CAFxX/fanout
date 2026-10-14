package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 bool; F1 uint8; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64, p1 int32, p2 uint64) int8 {
    p0 -= 42
    switch "" {
    case "a":
        var v7 int8 = -92
        _ = v7
    case "":
        var v8 uint8 = 226
        _ = v8
    default:
        var v9 float64 = 0.1
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        switch (fi10 & fi10) {
        case 32768:
            var v11 int16 = -1
            _ = v11
        case 986:
            v12 := GenS1{F0: "a", F1: p1}
            _ = v12
        default:
            fmt.Println(p0, p2)
        }
        _ = fi10
    }
    p0 = p0
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 uint16) uint8 {
    v13 := []int{1, 32767, -9223372036854775808}
    v14 := GenS2{F0: -1, F1: 197, F2: ""}
    var v15 int = -241
    if len(v13) > 0 {
        v13[0] = (v15 * v15)
    }
    fmt.Println(v15, v15)
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]bool)
    v3 := []int{}
    var v4 int64 = -9223372036854775808
    v5 := make(map[string]uint32)
    var v6 int16 = 42
    var v16 int8 = 42
    if (820 != 4294967295) {
        v6 = v6
        v17 := GenS2{F0: 850, F1: 571, F2: (((("foo bar" + "foo bar") + "foo bar") + "hello") + (("a" + ("foo bar" + "hello")) + ""))}
        _ = v17
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        var v19 uint64 = 18446744073709551615
        _ = fi18
        _ = v19
    }
    var v20 int8 = (-(v16))
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v4)
    case "hello":
        if true {
            if (!(1 <= 305)) {
                var v21 uint32 = 4294967295
                if len(v3) > 0 {
                    v3[0] = 462
                }
                var v22 bool = false
                _ = v21
                _ = v22
            } else {
                v23 := GenS3{F0: 2, F1: false}
                v4 -= v4
                _ = v23
            }
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                v6 -= -1
                switch "hello" {
                case "hello":
                    if true {
                        v25 := GenS2{F0: v6, F1: fi24, F2: "\n\t"}
                        v6 -= (v6 / (((v6 / (((-(v6))) | 1))) | 1))
                        v5[("\n\t" + "hello")] = 2147483647
                        _ = v25
                    }
                case "a":
                    switch fi24 {
                    case 4294967295:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            v4 = v4
                            _ = fi26
                        }
                    case 2147483647:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 153:
                        _ = len(v3)
                    }
                case "":
                    v27 := GenS3{F0: 0, F1: ("" >= "hello")}
                    _ = v27
                default:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        if (!true) {
                            if false {
                                v16 -= v20
                            } else {
                                fmt.Println(fi28, v20, fi28)
                                fmt.Println(v20, fi24, v4)
                                if (77 != 42) {
                                    var v29 float32 = (1.5 - float32(1e10))
                                    var v30 float64 = float64(float32((1.5 - -2.25)))
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 uint16 = 65535
                                    _ = v31
                                }
                            }
                            v32 := GenS4{F0: true, F1: 127, F2: v6}
                            var v33 int16 = v6
                            _ = v32
                            _ = v33
                        }
                        var v34 int64 = v4
                        var v35 int8 = (^(v16))
                        if len(v3) > 0 {
                            v3[0] = fi28
                        }
                        _ = fi28
                        _ = v34
                        _ = v35
                    }
                }
                v16 -= genFn0((v4 >> ((319) & 63)), -545, 4294967295)
                _ = fi24
            }
            if (false || true) {
                v36 := GenS4{F0: ("\n\t" <= ""), F1: 110, F2: -32768}
                fmt.Println(v4, v16)
                _ = v36
            }
        }
    case "":
        v37 := GenS5{F0: v4}
        _ = v37
    }
    v38 := GenS5{F0: (v4 / ((v4) | 1))}
    if len(v3) > 0 {
        v3[0] = 498
    }
    if true {
        v16 += v16
    } else {
        v1[594] = "hello"
        v16 -= 42
        v39 := GenS6{F0: 127, F1: v16}
        _ = v39
    }
    _ = len(v5)
    fmt.Println(v4, v16, v20)
    var v40 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
