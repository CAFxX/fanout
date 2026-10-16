package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string; F1 uint; F2 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 string; F1 uint16 }
func genFn0() bool {
    if (true && false) {
        var v4 int16 = 1
        v5 := GenS1{F0: 0, F1: ""}
        _ = v4
        _ = v5
    } else {
        switch ("foo bar" + ("\n\t" + ((("hello" + (("\n\t" + "") + ("" + "a"))) + ("foo bar" + "foo bar")) + ""))) {
        case "\n\t":
            var v6 float64 = (1.5 + 0.0)
            _ = v6
        default:
            var v7 int64 = -460
            _ = v7
        }
        v8 := GenS1{F0: -1, F1: "a"}
        var v9 int64 = 232
        _ = v8
        _ = v9
    }
    var v10 int32 = 149
    var v11 bool = true
    v12 := GenS1{F0: 114, F1: (("foo bar" + "hello") + ("" + ("" + "\n\t")))}
    _ = v10
    _ = v11
    _ = v12
    return v11
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 string = (("\n\t" + "a") + "")
    v13 := GenS2{F0: v3, F1: 783, F2: v3}
    var v14 int = 42
    v15 := GenS3{F0: 869, F1: 0}
    var v16 uint64 = 42
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 = ""
            if (177 >= 0) {
                v14 = (fi17 & -9223372036854775808)
                if (100 == -128) {
                    v16 -= v16
                } else {
                    v14 = fi18
                    if true {
                        v19 := GenS3{F0: uint32(v16), F1: (^(v16))}
                        v14 += fi17
                        if ((v3 + v3) > ((v3 + "hello") + (v3 + ("hello" + (("hello" + "foo bar") + (v3 + ((v3 + "a") + v3))))))) {
                            if true {
                                _ = len(v1)
                                if (!(1 > 122)) {
                                    switch "a" {
                                    case "a":
                                        if ((0.0 > ((0.0 + 1.5) * (float32((1e10 + ((1e10 / 1e10) / -2.25))) / ((float32(3.14159) - (1.5 + 0.0)) + float32(3.14159))))) || true) {
                                            v1 = append(v1, (fi17 * 9223372036854775807))
                                            fmt.Println(v16, fi17, v3)
                                            if (true && (2 == v16)) {
                                                v20 := GenS4{F0: 1}
                                                delete(v2, 32768)
                                                switch v3 {
                                                case "hello":
                                                    v16 = (v16 * 286)
                                                    var v21 bool = true
                                                    v22 := GenS4{F0: 2147483647}
                                                    v23 := GenS4{F0: -1}
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                case "a":
                                                    var v24 uint32 = 4294967295
                                                    _ = v24
                                                }
                                                var v25 int64 = -9223372036854775808
                                                var v26 float32 = 0.1
                                                _ = v20
                                                _ = v25
                                                _ = v26
                                            } else {
                                                var v27 int32 = -1
                                                v3 = v3
                                                _ = v27
                                            }
                                        } else {
                                            var v28 int16 = 393
                                            _ = v28
                                        }
                                    case "":
                                        v29 := GenS5{F0: -1}
                                        _ = v29
                                    default:
                                        v30 := GenS3{F0: 42, F1: v16}
                                        v14 -= fi17
                                        _ = v30
                                    }
                                    v16 += (^(v16))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            } else {
                                var v31 int = (-(fi17))
                                v3 = (v3 + v3)
                                _ = v31
                            }
                            v32 := GenS4{F0: 1}
                            _ = v32
                        }
                        _ = v19
                    }
                    v1 = append(v1, (fi18 % (((fi17 / ((v14) | 1))) | 1)))
                }
                v16 = v16
            } else {
                var v33 string = v3
                _ = v33
            }
            v1 = append(v1, 42)
            _ = fi18
        }
        v1 = append(v1, (v14 << ((500) & 63)))
        _ = fi17
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v34 int32 = -190
    switch 698 {
    case 913:
        delete(v2, (v14 + v14))
    case 2147483647:
        if true {
            fmt.Println(v34, v16, v3)
        }
    case 128:
        var v35 uint32 = uint32(v16)
        _ = v35
    default:
        switch "\n\t" {
        case "hello":
            v14 = v14
        case "foo bar":
            var v36 float64 = float64(1.5)
            _ = v36
        case "\n\t":
            delete(v2, 42)
        }
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        v38 := GenS4{F0: v34}
        if len(v1) > 0 {
            v1[0] = fi37
        }
        v34 -= (v34 - 0)
        _ = fi37
        _ = v38
    }
    v39 := GenS6{F0: -828, F1: -862}
    v40 := GenS7{F0: (("\n\t" + ("hello" + "hello")) + "foo bar"), F1: 918}
    var v41 string = v3
    var v42 uint64 = v16
    var v43 uint16 = 32767
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    for fi44 := 0; fi44 < 10; fi44++ {
        fmt.Println(v42, v14)
        _ = fi44
    }
    if len(v1) > 0 {
        v1[0] = (v14 / (((v14 << ((255) & 63))) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v34
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
