package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0(p0 int) uint {
    p0 += (p0 & p0)
    p0 = 128
    if false {
        var v7 string = "a"
        var v8 uint8 = 2
        v9 := []int{}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 string = "a"
        _ = v10
    }
    _ = p0
    return 1
}
func genFn1(p0 string) uint8 {
    if false {
        v11 := GenS1{F0: "\n\t", F1: -2147483648, F2: 297}
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 float64 = (0.0 - (0.0 * 0.1))
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        var v15 uint64 = 735
        _ = v11
        _ = v15
    } else {
        v16 := GenS1{F0: p0, F1: 0, F2: 0}
        p0 = "foo bar"
        var v17 uint = 18446744073709551615
        _ = v16
        _ = v17
    }
    if (0 == 0) {
        if false {
            p0 = "\n\t"
        }
    } else {
        v18 := GenS2{F0: 4294967295, F1: 603}
        _ = v18
    }
    if false {
        if true {
            v19 := []int{}
            var v20 string = p0
            v20 = ("a" + "foo bar")
            _ = v19
            _ = v20
        } else {
            var v21 string = "\n\t"
            if (3.14159 > (0.0 - 0.1)) {
                fmt.Println(v21)
                v22 := GenS2{F0: -961, F1: 448}
                _ = v22
            } else {
                switch 127 {
                case -245:
                    if false {
                        v23 := []int{1, 10, -657}
                        if false {
                            var v24 bool = (!(v21 < (v21 + v21)))
                            var v25 uint32 = 1
                            _ = v24
                            _ = v25
                        }
                        p0 = "hello"
                        _ = v23
                    } else {
                        var v26 uint = 941
                        if true {
                            v21 = (v21 + (v21 + p0))
                            v26 += 0
                        }
                        _ = v26
                    }
                case 1:
                    var v27 float32 = 1e10
                    _ = v27
                case 693:
                    switch (((p0 + p0) + ("\n\t" + v21)) + p0) {
                    case "":
                        v28 := GenS3{F0: -1}
                        _ = v28
                    case "hello":
                        if (!(-2147483648 >= 2147483647)) {
                            switch -593 {
                            case -2147483648:
                                p0 = p0
                            case 127:
                                if true {
                                    v21 = "a"
                                    v29 := []int{9223372036854775807, 32768, 494}
                                    var v30 int = 2147483647
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 uint64 = 543
                                    p0 = ((p0 + "a") + "")
                                    for fi32 := 0; fi32 < 2; fi32++ {
                                        v33 := GenS2{F0: (fi32 * fi32), F1: 2}
                                        v21 = ((p0 + v21) + p0)
                                        _ = fi32
                                        _ = v33
                                    }
                                    _ = v31
                                }
                            }
                            if (((float32((1.5 * 0.0)) * 1.5) == -2.25) && (65535 != 32767)) {
                                v34 := GenS3{F0: 2147483647}
                                v21 = v21
                                _ = v34
                            } else {
                                var v35 float64 = 0.0
                                fmt.Println(p0, p0)
                                v36 := GenS1{F0: (p0 + "hello"), F1: -1, F2: -645}
                                _ = v35
                                _ = v36
                            }
                        } else {
                            if false {
                                var v37 int8 = 127
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    v39 := []int{(fi38 % ((fi38) | 1)), -732}
                                    _ = fi38
                                    _ = v39
                                }
                                _ = v37
                            }
                        }
                    }
                default:
                    switch v21 {
                    case "a":
                        if (2147483648 == 2) {
                            for fi40 := 0; fi40 < 16; fi40++ {
                                switch (fi40 % (((fi40 + fi40)) | 1)) {
                                case 869:
                                    fmt.Println(v21, fi40)
                                case 4294967295:
                                    if (float64(1.5) > float64(0.0)) {
                                        var v41 uint16 = 0
                                        if (!(true || true)) {
                                            v42 := GenS4{F0: 2, F1: true}
                                            _ = v42
                                        } else {
                                            for fi43 := 0; fi43 < 13; fi43++ {
                                                v44 := GenS5{F0: 2, F1: (fi40 | fi40)}
                                                _ = fi43
                                                _ = v44
                                            }
                                        }
                                        _ = v41
                                    }
                                case 1:
                                    var v45 int8 = 127
                                    _ = v45
                                default:
                                    v46 := GenS6{F0: 42, F1: genFn0(fi40)}
                                    _ = v46
                                }
                                v21 = "a"
                                _ = fi40
                            }
                            fmt.Println(p0, v21)
                            v21 = v21
                        } else {
                            switch "hello" {
                            case "\n\t":
                                var v47 string = v21
                                _ = v47
                            case "a":
                                v48 := GenS7{F0: 42}
                                _ = v48
                            case "foo bar":
                                v21 = ""
                            default:
                                p0 = v21
                            }
                            v49 := GenS5{F0: 18446744073709551615, F1: -92}
                            _ = v49
                        }
                    case "\n\t":
                        p0 = (v21 + p0)
                    default:
                        var v50 uint8 = 11
                        _ = v50
                    }
                }
            }
            _ = v21
        }
    }
    _ = p0
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    var v3 int16 = 32767
    var v4 string = "hello"
    var v5 uint16 = 354
    v6 := []uint32{936, 332, 2147483648}
    v51 := GenS2{F0: -461, F1: 349}
    v52 := GenS6{F0: genFn0(int(v5)), F1: 42}
    for fi53 := 0; fi53 < 12; fi53++ {
        var v54 int32 = 1
        _ = fi53
        _ = v54
    }
    v55 := GenS6{F0: genFn0(32768), F1: 0}
    for fi56 := 0; fi56 < 7; fi56++ {
        for fi57 := 0; fi57 < 9; fi57++ {
            var v58 int32 = 1
            _ = fi57
            _ = v58
        }
        v2 -= v2
        if len(v6) > 0 {
            v6[0] = 65535
        }
        _ = fi56
    }
    _ = len(v1)
    fmt.Println(v3, v4)
    fmt.Println(v3, v2)
    var v59 uint8 = 127
    v4 = v4
    for fi60 := 0; fi60 < 14; fi60++ {
        v61 := GenS2{F0: fi60, F1: 0}
        _ = fi60
        _ = v61
    }
    v4 = v4
    switch "" {
    case "foo bar":
        v4 = (((("" + ("a" + v4)) + (v4 + (v4 + (((v4 + ("" + "hello")) + (v4 + ("hello" + ("hello" + v4)))) + v4)))) + ((("" + v4) + v4) + ("hello" + "a"))) + "\n\t")
    case "":
        if true {
            var v62 float32 = v2
            _ = v62
        } else {
            v63 := GenS8{F0: uint32(v59), F1: v3, F2: 0}
            _ = v63
        }
    default:
        if false {
            fmt.Println(v3, v5)
            if (false || (356 > 18446744073709551615)) {
                v64 := v1[4294967295]
                v65 := v1[v64]
                _ = v64
                _ = v65
            }
        }
    }
    var v66 int32 = 638
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v55
    _ = v59
    _ = v6
    _ = v66
}
