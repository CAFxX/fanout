package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int; F1 int32; F2 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int64; F1 int32 }
func genFn0(p0 string, p1 uint16) uint32 {
    var v4 uint64 = 169
    v5 := GenS1{F0: -114, F1: (p1 >> ((866) & 63))}
    v6 := GenS2{F0: 0, F1: (v4 ^ (v4 / (((^(v4))) | 1)))}
    var v7 uint16 = p1
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1(p0 string, p1 uint32, p2 uint8) uint8 {
    v9 := []int{-789}
    var v10 string = p0
    _ = len(v9)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p2
}

func main() {
    var v1 uint16 = 32767
    v2 := []string{"", "foo bar"}
    var v3 int64 = -157
    switch "hello" {
    case "\n\t":
        v1 = (v1 | 32767)
    case "":
        v1 = v1
    case "hello":
        var v11 int8 = 42
        _ = v11
    default:
        fmt.Println(v1, v1)
    }
    if true {
        if len(v2) > 0 {
            v2[0] = ""
        }
        v3 = 0
        var v12 float32 = 1.5
        _ = v12
    } else {
        v13 := GenS2{F0: (v3 & -1), F1: 4294967295}
        var v14 uint8 = 128
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v3 += -437
        v1 -= v1
        _ = fi15
    }
    var v16 float32 = 1.5
    v17 := GenS2{F0: v3, F1: 0}
    for fi18 := 0; fi18 < 4; fi18++ {
        if true {
            v1 = (^(v1))
        }
        _ = fi18
    }
    var v19 int16 = -1
    if ((false && ((891 != 4294967295) || ((v16 - v16) == v16))) || false) {
        v3 += (^(v3))
        var v20 bool = false
        v2 = append(v2, ("" + "foo bar"))
        _ = v20
    } else {
        _ = len(v2)
        var v21 float64 = 1e10
        v2 = append(v2, (((((("a" + ("foo bar" + (("hello" + "a") + "\n\t"))) + "foo bar") + "") + "a") + ("\n\t" + "hello")) + "\n\t"))
        _ = v21
    }
    var v22 int = -846
    if ((-1 <= v19) && false) {
        fmt.Println(v19, v3)
    }
    if len(v2) > 0 {
        v2[0] = ("" + ("foo bar" + ("a" + ("a" + "hello"))))
    }
    if (!(!(!(255 != 0)))) {
        if true {
            switch "" {
            case "a":
                var v23 uint64 = 0
                _ = v23
            default:
                switch "hello" {
                case "foo bar":
                    if false {
                        v22 -= (v22 % (((v22 | v22)) | 1))
                        v24 := GenS1{F0: 42, F1: (^(v1))}
                        _ = v24
                    } else {
                        for fi25 := 0; fi25 < 3; fi25++ {
                            switch fi25 {
                            case 4294967295:
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    v3 = (v3 ^ (^(v3)))
                                    fmt.Println(v1)
                                    _ = fi26
                                }
                            default:
                                v27 := GenS3{F0: v22, F1: 0, F2: (v1 >> ((2) & 63))}
                                _ = v27
                            }
                            _ = len(v2)
                            _ = fi25
                        }
                        var v28 string = "foo bar"
                        fmt.Println(v28, v28)
                        _ = v28
                    }
                case "\n\t":
                    switch "" {
                    case "":
                        v29 := GenS3{F0: (v22 + v22), F1: -1, F2: 65535}
                        _ = v29
                    case "\n\t":
                        _ = len(v2)
                    }
                case "a":
                    if true {
                        v3 += v3
                        v30 := GenS4{F0: (^(v19))}
                        var v31 uint64 = 18446744073709551615
                        _ = v30
                        _ = v31
                    } else {
                        v1 = (v1 % ((v1) | 1))
                        v16 = (v16 - 1.5)
                        for fi32 := 0; fi32 < 6; fi32++ {
                            var v33 int = v22
                            for fi34 := 0; fi34 < 3; fi34++ {
                                var v35 uint16 = (v1 % (((v1 | (v1 ^ (^(v1))))) | 1))
                                v36 := GenS5{F0: v3, F1: ("" != "hello")}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                }
            }
        }
        var v37 int64 = v3
        _ = v37
    } else {
        _ = len(v2)
    }
    v38 := GenS6{F0: v3, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v38
}
