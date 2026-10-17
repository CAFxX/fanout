package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int32; F1 uint64; F2 uint64 }
func genFn0(p0 int) uint32 {
    v7 := []int{p0, (p0 | p0)}
    for fi8 := 0; fi8 < 13; fi8++ {
        if true {
            if true {
                p0 += (p0 ^ p0)
                if false {
                    var v9 uint8 = 42
                    for fi10 := 0; fi10 < 15; fi10++ {
                        fmt.Println(fi10, fi10)
                        if len(v7) > 0 {
                            v7[0] = fi8
                        }
                        _ = len(v7)
                        v11 := GenS1{F0: (true && false), F1: 42, F2: "a"}
                        v9 = v9
                        _ = fi10
                        _ = v11
                    }
                    _ = v9
                }
            }
            v7 = append(v7, (fi8 * -323))
        } else {
            v12 := GenS1{F0: (127 > 42), F1: -482, F2: "\n\t"}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 float32 = 1e10
            p0 -= fi8
            _ = fi13
            _ = v14
        }
        _ = fi8
    }
    if true {
        switch (p0 ^ p0) {
        case 9223372036854775807:
            if (947 == 873) {
                p0 += (p0 + p0)
                if (!false) {
                    p0 -= p0
                    p0 += p0
                    p0 = p0
                } else {
                    p0 = 128
                }
            }
        default:
            if ("a" == "foo bar") {
                fmt.Println(p0)
            } else {
                v15 := GenS1{F0: true, F1: -574, F2: (((("" + "a") + ("" + "\n\t")) + "") + "\n\t")}
                switch p0 {
                case 42:
                    var v16 int16 = 1000
                    _ = v16
                case 4294967295:
                    p0 = (-(p0))
                case 32768:
                    if false {
                        p0 = (p0 >> ((142) & 63))
                    }
                }
                _ = v15
            }
        }
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        if false {
            _ = len(v7)
            fmt.Println(p0)
        }
        p0 += p0
        for fi18 := 0; fi18 < 17; fi18++ {
            p0 = fi17
            _ = fi18
        }
        _ = fi17
    }
    v19 := GenS2{F0: ("hello" + "foo bar")}
    _ = p0
    _ = v19
    _ = v7
    return 0
}

func main() {
    v1 := []uint32{292}
    var v2 int8 = -95
    var v3 uint32 = 42
    var v4 float64 = -2.25
    v5 := make(map[int]string)
    var v6 uint16 = 399
    v4 = v4
    v20 := GenS1{F0: true, F1: 1, F2: ""}
    delete(v5, 0)
    var v21 uint8 = 42
    for fi22 := 0; fi22 < 2; fi22++ {
        if (v2 < v2) {
            if false {
                var v23 int32 = 1
                _ = v23
            } else {
                v3 += 68
                for fi24 := 0; fi24 < 11; fi24++ {
                    v3 = (v3 / ((uint32(v21)) | 1))
                    _ = fi24
                }
            }
        } else {
            _ = len(v5)
            v25 := GenS1{F0: false, F1: 1, F2: "\n\t"}
            for fi26 := 0; fi26 < 11; fi26++ {
                v2 = (^(v2))
                fmt.Println(v2)
                switch (("" + "foo bar") + "") {
                case "foo bar":
                    v27 := GenS3{F0: v2}
                    _ = v27
                case "":
                    switch "foo bar" {
                    case "hello":
                        v28 := GenS4{F0: (v3 % (((v3 & genFn0(4294967295))) | 1)), F1: 42, F2: uint8(v6)}
                        switch "foo bar" {
                        case "foo bar":
                            if ((!false) && true) {
                                v6 = v6
                            } else {
                                v4 += ((1.5 + v4) + (v4 - (0.1 * (v4 - v4))))
                                v29 := GenS5{F0: 2}
                                var v30 uint16 = (v6 - (v6 + v6))
                                v31 := GenS5{F0: 18446744073709551615}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        default:
                            v32 := GenS5{F0: 676}
                            if ((1 < 2) || true) {
                                v21 = v21
                                var v33 int32 = 2147483647
                                var v34 int64 = -111
                                _ = v33
                                _ = v34
                            } else {
                                var v35 uint = 642
                                v6 -= (^(v6))
                                _ = v35
                            }
                            _ = v32
                        }
                        v36 := GenS6{F0: (v2 | (v2 % ((100) | 1))), F1: 1, F2: (((("foo bar" + ("\n\t" + "foo bar")) + "") + "\n\t") + "\n\t")}
                        _ = v28
                        _ = v36
                    case "\n\t":
                        fmt.Println(v21, v2)
                    default:
                        var v37 int = fi26
                        _ = v37
                    }
                default:
                    v3 += (v3 * 2)
                }
                var v38 uint8 = v21
                _ = fi26
                _ = v38
            }
            _ = v25
        }
        _ = fi22
    }
    v39 := GenS7{F0: 2147483647, F1: 1, F2: uint64(v21)}
    var v40 uint32 = v3
    v41 := v5[-513]
    if ((v41 + v41) <= v41) {
        var v42 uint64 = 518
        var v43 int8 = 1
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        v40 -= v40
        _ = fi44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
