package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint64; F1 int32; F2 int8 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 string; F1 string; F2 uint16 }
func genFn0(p0 uint16, p1 int32) uint16 {
    var v7 uint8 = 128
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int32 = p1
        _ = fi8
        _ = v9
    }
    var v10 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return p0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 bool = (42 != -1)
    var v3 uint32 = 2147483648
    v4 := make(map[int]int)
    var v5 int32 = -1
    v6 := make(map[string]uint32)
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := v6["hello"]
        _ = fi11
        _ = v12
    }
    v1 = (v1 | v1)
    fmt.Println(v1, v1)
    for fi13 := 0; fi13 < 19; fi13++ {
        if v2 {
            fmt.Println(v2)
            for fi14 := 0; fi14 < 6; fi14++ {
                v15 := []int{-2147483648, 454, 4294967295, fi14}
                v16 := GenS1{F0: 2147483647}
                _ = fi14
                _ = v15
                _ = v16
            }
        } else {
            for fi17 := 0; fi17 < 19; fi17++ {
                if false {
                    if ((v2 || v2) && v2) {
                        switch "\n\t" {
                        case "a":
                            fmt.Println(fi13, v2)
                        case "hello":
                            v18 := GenS1{F0: 2}
                            _ = v18
                        case "\n\t":
                            v1 -= (^(v1))
                        default:
                            v19 := []int{-2147483648}
                            v2 = v2
                            v5 -= (^(v5))
                            _ = v19
                        }
                        v20 := GenS2{F0: "", F1: 42, F2: "hello"}
                        v1 -= v1
                        _ = v20
                    }
                    var v21 string = ""
                    var v22 float32 = 3.14159
                    _ = v21
                    _ = v22
                } else {
                    switch int(v1) {
                    case 128:
                        v5 = 1
                    case 0:
                        switch "\n\t" {
                        case "foo bar":
                            delete(v6, "foo bar")
                        case "":
                            if v2 {
                                fmt.Println(fi13, v2)
                                var v23 uint32 = (v3 % ((v3) | 1))
                                _ = v23
                            }
                        default:
                            v24 := GenS3{F0: (fi13 % ((fi13) | 1)), F1: 42}
                            _ = v24
                        }
                    default:
                        var v25 uint8 = 128
                        _ = v25
                    }
                    var v26 int8 = -29
                    _ = v26
                }
                var v27 uint8 = 255
                v28 := GenS4{F0: uint16(v3), F1: 0}
                v29 := GenS5{F0: (v1 + v1), F1: v5, F2: 19}
                var v30 string = "hello"
                switch -315 {
                case 0:
                    v31 := GenS3{F0: fi17, F1: v3}
                    if v2 {
                        if v2 {
                            v32 := GenS6{F0: fi13, F1: (v27 - (^(v27)))}
                            _ = v32
                        } else {
                            if (!(39 != 106)) {
                                switch v30 {
                                case "\n\t":
                                    v33 := v6[(v30 + v30)]
                                    _ = v33
                                case "":
                                    v34 := v4[fi13]
                                    _ = v34
                                default:
                                    v27 = uint8(v3)
                                }
                            } else {
                                v35 := GenS6{F0: fi13, F1: (^(v27))}
                                v36 := GenS7{F0: ((v30 + "hello") + "a"), F1: "\n\t", F2: uint16(v3)}
                                _ = v35
                                _ = v36
                            }
                            v1 = 442
                            v37 := GenS3{F0: fi13, F1: 2147483648}
                            _ = v37
                        }
                        fmt.Println(fi17, v2)
                    }
                    fmt.Println(v1, fi13, v2)
                    _ = v31
                case 2147483647:
                    v38 := GenS4{F0: 65535, F1: 42}
                    _ = v38
                case 4294967295:
                    v39 := v6[v30]
                    _ = v39
                default:
                    var v40 uint8 = (^(v27))
                    _ = v40
                }
                if true {
                    var v41 float32 = ((float32(1.5) / 1.5) / (1e10 * float32(-2.25)))
                    switch 0 {
                    case -186:
                        fmt.Println(fi13, v3, v3)
                    case 292:
                        if false {
                            fmt.Println(v30, v41)
                            v1 = uint64(v27)
                        } else {
                            fmt.Println(v3)
                            fmt.Println(v30)
                            v2 = v2
                        }
                    default:
                        v30 = (v30 + v30)
                    }
                    _ = v41
                } else {
                    var v42 uint32 = (v3 | v3)
                    v43 := []int{(fi17 << ((118) & 63)), fi13, fi13, (fi13 ^ fi17)}
                    switch (fi13 * fi13) {
                    case 127:
                        fmt.Println(v30, v30, v2)
                    case -1:
                        fmt.Println(v2)
                    default:
                        v5 = (v5 / (((v5 << ((1) & 63))) | 1))
                    }
                    _ = v42
                    _ = v43
                }
                _ = fi17
                _ = v27
                _ = v28
                _ = v29
                _ = v30
            }
            var v44 int64 = 42
            v1 += 1
            _ = v44
        }
        _ = fi13
    }
    var v45 float64 = 0.1
    if (652 != -9223372036854775808) {
        var v46 uint8 = 0
        v4[-460] = -811
        _ = v46
    }
    v5 = 1
    var v47 float64 = (3.14159 * 0.0)
    v48 := GenS2{F0: "\n\t", F1: 519, F2: ("a" + "a")}
    switch 9223372036854775807 {
    case 128:
        var v49 uint = 0
        _ = v49
    case 606:
        if v2 {
            var v50 uint8 = 33
            var v51 int64 = 9223372036854775807
            _ = v50
            _ = v51
        }
    case 0:
        _ = len(v4)
    }
    var v52 float32 = ((3.14159 - 3.14159) + (0.1 * (float32((((v47 + v47) * ((3.14159 * (1.5 * v47)) + v45)) + (v47 * v47))) + 0.0)))
    switch "a" {
    case "\n\t":
        if (v2 || false) {
            v53 := v6[("" + ("a" + "hello"))]
            v54 := []int{}
            for fi55 := 0; fi55 < 14; fi55++ {
                var v56 uint8 = 2
                var v57 int8 = 0
                _ = fi55
                _ = v56
                _ = v57
            }
            _ = v53
            _ = v54
        }
    case "foo bar":
        var v58 int32 = (v5 ^ (v5 & (-(v5))))
        _ = v58
    case "a":
        v45 = float64(v52)
    default:
        v59 := GenS7{F0: ("" + ""), F1: "hello", F2: uint16(v3)}
        _ = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v48
    _ = v5
    _ = v52
    _ = v6
}
