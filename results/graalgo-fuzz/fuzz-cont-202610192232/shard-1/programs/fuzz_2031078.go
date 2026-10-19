package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 string; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 bool; F1 uint; F2 uint }
type GenS7 struct { F0 uint64 }
func genFn0() string {
    v6 := GenS1{F0: 1, F1: 1}
    var v7 uint16 = 42
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1() int16 {
    var v8 float64 = 3.14159
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 string = (("" + "") + (("\n\t" + "hello") + ""))
        _ = fi9
        _ = v10
    }
    _ = v8
    return -242
}

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 486
    v3 := make(map[string]uint32)
    v4 := make(map[string]uint32)
    v5 := make(map[string]uint32)
    v11 := GenS2{F0: 127}
    switch "foo bar" {
    case "hello":
        if true {
            if true {
                switch ("a" + (("hello" + ("foo bar" + "hello")) + "foo bar")) {
                case "a":
                    if true {
                        v12 := []int{32767, 42, 1}
                        _ = v12
                    }
                case "foo bar":
                    if true {
                        if (float64(0.1) != -2.25) {
                            var v13 uint32 = 855
                            _ = v13
                        }
                        v14 := GenS3{F0: -816, F1: 32767}
                        v15 := []int{-565, -190}
                        _ = v14
                        _ = v15
                    } else {
                        v2 = v2
                        var v16 int8 = 0
                        _ = v16
                    }
                default:
                    var v17 int32 = -399
                    _ = v17
                }
                v2 = v2
                v18 := []int{}
                _ = v18
            } else {
                for fi19 := 0; fi19 < 4; fi19++ {
                    if false {
                        v2 -= v2
                    } else {
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 float64 = -2.25
                            _ = fi20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 8; fi22++ {
                            if (v2 <= (v2 & (v2 ^ (v2 | v2)))) {
                                delete(v5, "\n\t")
                            }
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 13; fi23++ {
                            if (!true) {
                                var v24 float64 = ((3.14159 + 0.1) / 1.0)
                                switch "\n\t" {
                                case "\n\t":
                                    _ = len(v1)
                                default:
                                    var v25 int64 = 0
                                    _ = v25
                                }
                                var v26 int64 = 541
                                _ = v24
                                _ = v26
                            }
                            var v27 uint8 = 128
                            v28 := GenS4{F0: (("foo bar" + "foo bar") + ""), F1: (fi19 & (fi19 << ((42) & 63))), F2: -32768}
                            switch "a" {
                            case "\n\t":
                                if ((2147483647 >= 201) && (3.14159 < 1.5)) {
                                    v29 := []int{fi19, (fi19 & 32768), (fi19 % (((fi23 ^ fi19)) | 1)), fi23}
                                    if (!false) {
                                        fmt.Println(v27, fi23, fi19)
                                        if false {
                                            switch "a" {
                                            case "":
                                                v30 := GenS5{F0: 1, F1: 42}
                                                _ = v30
                                            default:
                                                v31 := GenS6{F0: false, F1: 255, F2: 2}
                                                _ = v31
                                            }
                                            v27 = v27
                                            if true {
                                                if false {
                                                    switch "\n\t" {
                                                    case "":
                                                        v32 := GenS3{F0: 42, F1: -457}
                                                        _ = v32
                                                    case "a":
                                                        if (42 < -1) {
                                                            var v33 int32 = 58
                                                            if false {
                                                                fmt.Println(v27, fi23)
                                                                if (v27 <= (^(v27))) {
                                                                    if false {
                                                                        var v34 int = (-(fi19))
                                                                        _ = v34
                                                                    }
                                                                    var v35 int32 = (v33 & 368)
                                                                    _ = v35
                                                                } else {
                                                                    v36 := GenS3{F0: (v33 & v33), F1: 1}
                                                                    var v37 bool = (-1 >= -32768)
                                                                    _ = v36
                                                                    _ = v37
                                                                }
                                                            } else {
                                                                var v38 uint64 = (v2 << ((uint(v2)) & 63))
                                                                if ("" >= "") {
                                                                    v2 += (^(v38))
                                                                    v39 := GenS5{F0: v33, F1: 1}
                                                                    _ = v39
                                                                }
                                                                var v40 int = 128
                                                                _ = v38
                                                                _ = v40
                                                            }
                                                            _ = v33
                                                        }
                                                    case "\n\t":
                                                        var v41 uint8 = 1
                                                        _ = v41
                                                    default:
                                                        v42 := GenS2{F0: (v27 & (^(v27)))}
                                                        _ = v42
                                                    }
                                                    var v43 uint32 = 65535
                                                    _ = v43
                                                }
                                                if (3.14159 < float32(-2.25)) {
                                                    v27 = (v27 & 2)
                                                } else {
                                                    var v44 uint64 = (v2 * 908)
                                                    _ = v44
                                                }
                                            }
                                        }
                                        var v45 int8 = 0
                                        _ = v45
                                    }
                                    _ = v29
                                } else {
                                    v46 := GenS4{F0: "hello", F1: (fi19 % (((fi19 - (fi23 - fi23))) | 1)), F2: genFn1()}
                                    v2 += (v2 >> ((42) & 63))
                                    fmt.Println(fi19, v2, v2)
                                    v47 := GenS7{F0: v2}
                                    _ = v46
                                    _ = v47
                                }
                            default:
                                v48 := GenS5{F0: 691, F1: 2}
                                _ = v48
                            }
                            v2 += (v2 - 4294967295)
                            v49 := GenS6{F0: true, F1: 2, F2: uint(v2)}
                            var v50 float64 = float64((1e10 + (3.14159 * 3.14159)))
                            _ = fi23
                            _ = v27
                            _ = v28
                            _ = v49
                            _ = v50
                        }
                    }
                    v2 -= v2
                    v2 = 951
                    _ = fi19
                }
                var v51 uint16 = 32767
                _ = v51
            }
            v52 := []int{-9223372036854775808, -2147483648, 545}
            _ = v52
        }
    default:
        for fi53 := 0; fi53 < 6; fi53++ {
            for fi54 := 0; fi54 < 10; fi54++ {
                fmt.Println(fi54, fi54)
                _ = len(v5)
                _ = fi54
            }
            for fi55 := 0; fi55 < 15; fi55++ {
                var v56 uint = 0
                v2 += (v2 ^ (^(v2)))
                _ = fi55
                _ = v56
            }
            _ = fi53
        }
    }
    switch "a" {
    case "\n\t":
        var v57 int = 4294967295
        _ = v57
    case "foo bar":
        var v58 uint8 = 2
        _ = v58
    case "hello":
        v2 -= (v2 << ((0) & 63))
    default:
        var v59 uint32 = 2147483647
        _ = v59
    }
    if (true && (genFn1() == 397)) {
        v60 := GenS7{F0: v2}
        _ = v60
    }
    v2 += (v2 >> ((435) & 63))
    if true {
        v61 := []int{2147483647, 127, 9223372036854775807}
        v4["a"] = 456
        _ = v61
    }
    v62 := GenS5{F0: -2147483648, F1: 717}
    var v63 int8 = -128
    delete(v1, 128)
    v64 := GenS1{F0: 9223372036854775807, F1: 622}
    v65 := []int{32768, 32767}
    switch ("" + "hello") {
    case "foo bar":
        var v66 int = 42
        _ = v66
    }
    fmt.Println(v63)
    fmt.Println(v2, v63)
    fmt.Println(v63)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v65:", len(v65), v65)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
