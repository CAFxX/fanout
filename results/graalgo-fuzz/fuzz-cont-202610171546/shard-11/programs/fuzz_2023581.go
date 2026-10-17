package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint; F1 int64; F2 uint16 }
func genFn0(p0 bool, p1 bool, p2 int64) int64 {
    fmt.Println(p1)
    var v5 uint16 = 977
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: v5}
        v8 := []int{}
        p1 = p0
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (^(p2))
}
func genFn1(p0 uint32) uint32 {
    var v9 bool = true
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 float64 = 0.0
        _ = fi10
        _ = v11
    }
    fmt.Println(v9)
    var v12 int32 = 2147483647
    _ = p0
    _ = v12
    _ = v9
    return (^(p0))
}

func main() {
    var v1 int32 = -711
    v2 := make(map[uint32]int)
    v3 := []int{32768, 0, 2147483647}
    var v4 int8 = -118
    v1 -= (v1 & (v1 % ((v1) | 1)))
    var v13 bool = (0 >= 0)
    v14 := GenS1{F0: 0}
    v15 := GenS1{F0: 2}
    fmt.Println(v1, v1)
    for fi16 := 0; fi16 < 16; fi16++ {
        if ("" >= (("\n\t" + ("" + "a")) + "")) {
            for fi17 := 0; fi17 < 19; fi17++ {
                switch fi16 {
                case -603:
                    v18 := GenS1{F0: 42}
                    _ = v18
                case 32768:
                    var v19 float64 = (float64(((0.1 / 1.0) / 1.0)) / 0.1)
                    _ = v19
                case -746:
                    v1 = v1
                }
                v20 := GenS1{F0: 0}
                v13 = v13
                _ = fi17
                _ = v20
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v4)
            case "a":
                var v21 uint8 = 13
                _ = v21
            }
            var v22 uint64 = 18446744073709551615
            _ = v22
        } else {
            fmt.Println(v13)
            v13 = v13
            for fi23 := 0; fi23 < 10; fi23++ {
                if (!false) {
                    v4 += v4
                    switch ("foo bar" + "") {
                    case "\n\t":
                        v1 = v1
                    case "":
                        v3 = append(v3, 199)
                    }
                    if true {
                        if true {
                            if v13 {
                                v3 = append(v3, -2147483648)
                                var v24 uint = 208
                                _ = v24
                            }
                            switch ("a" + "") {
                            case "hello":
                                v25 := v2[2147483647]
                                _ = v25
                            case "a":
                                var v26 uint8 = 42
                                _ = v26
                            case "":
                                if v13 {
                                    var v27 float32 = (0.1 * 0.0)
                                    var v28 int8 = (v4 | 42)
                                    fmt.Println(fi16, fi16)
                                    _ = v27
                                    _ = v28
                                } else {
                                    var v29 uint32 = genFn1(2147483647)
                                    switch ("hello" + "") {
                                    case "hello":
                                        var v30 uint32 = v29
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        var v31 int = -9223372036854775808
                                        v32 := GenS2{F0: ("\n\t" + "")}
                                        var v33 uint32 = v30
                                        if v13 {
                                            v33 += (v33 * v30)
                                        }
                                        v34 := GenS3{F0: -1}
                                        v1 = 0
                                        switch fi23 {
                                        case 2147483647:
                                            v2[(v30 << ((370) & 63))] = fi23
                                        case -1:
                                            _ = len(v3)
                                        default:
                                            var v35 uint = 255
                                            _ = v35
                                        }
                                        _ = v30
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    case "a":
                                        _ = len(v3)
                                        var v36 int32 = (v1 + v1)
                                        var v37 int16 = 1000
                                        _ = v36
                                        _ = v37
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v29
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        _ = len(v3)
                    }
                }
                v2[2147483648] = fi23
                v1 += (v1 | v1)
                _ = fi23
            }
        }
        _ = fi16
    }
    var v38 int16 = -32768
    var v39 int = 825
    switch "\n\t" {
    case "\n\t":
        var v40 uint16 = 32767
        _ = v40
    case "":
        v3 = append(v3, (^(v39)))
    }
    if false {
        _ = len(v2)
        fmt.Println(v1, v1)
        v2[genFn1(184)] = -489
    } else {
        fmt.Println(v13, v38)
    }
    var v41 float32 = (((3.14159 + 0.0) + 1.5) - 0.1)
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v42 int32 = 0
        _ = v42
    } else {
        delete(v2, 65535)
        delete(v2, 106)
        v43 := GenS4{F0: 2, F1: 9223372036854775807, F2: 32767}
        _ = v43
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        v1 -= -234
        _ = fi44
    }
    var v45 uint32 = 900
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v45
}
