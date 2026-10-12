package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int8; F1 uint; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8 }
func genFn0() int {
    if false {
        var v4 int32 = -1
        v4 += v4
        _ = v4
    } else {
        var v5 bool = false
        var v6 int32 = -1
        _ = v5
        _ = v6
    }
    switch "foo bar" {
    case "":
        var v7 uint64 = 2
        _ = v7
    case "hello":
        var v8 int64 = -9223372036854775808
        _ = v8
    default:
        var v9 string = ("\n\t" + (("\n\t" + (("\n\t" + "foo bar") + "")) + ""))
        _ = v9
    }
    v10 := GenS1{F0: 42, F1: 988}
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{}
        var v13 float64 = 3.14159
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = v10
    return 0
}

func main() {
    var v1 bool = false
    var v2 float64 = (1.5 / 1.0)
    var v3 int8 = 65
    switch ("foo bar" + "\n\t") {
    case "hello":
        fmt.Println(v1, v2)
    case "a":
        v14 := []int{971, 2147483647, -464, -9223372036854775808}
        _ = v14
    case "\n\t":
        v15 := []int{genFn0()}
        _ = v15
    default:
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 string = "\n\t"
                _ = fi17
                _ = v18
            }
            v2 = v2
            var v19 int32 = 874
            _ = fi16
            _ = v19
        }
    }
    var v20 int8 = 1
    var v21 uint16 = 255
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            if v1 {
                var v24 string = "\n\t"
                _ = v24
            } else {
                switch "foo bar" {
                case "hello":
                    if (18446744073709551615 != 1) {
                        if (0 != 2) {
                            v25 := GenS2{F0: 0, F1: 0}
                            v2 = 1.5
                            fmt.Println(v2, v21)
                            _ = v25
                        } else {
                            v2 += 1.5
                            var v26 int = fi22
                            _ = v26
                        }
                    }
                case "foo bar":
                    v27 := []int{(fi22 >> ((4294967295) & 63)), fi23, -743}
                    _ = v27
                case "":
                    var v28 float32 = (0.1 * -2.25)
                    _ = v28
                }
            }
            v21 = (v21 >> ((0) & 63))
            if ((v3 << ((919) & 63)) <= 116) {
                v29 := []int{(-(fi23)), genFn0()}
                v20 = v20
                v21 = 0
                v30 := GenS1{F0: 89, F1: 979}
                var v31 uint8 = 128
                v32 := GenS2{F0: 1, F1: (v31 ^ (v31 + (v31 % ((1) | 1))))}
                if true {
                    fmt.Println(v20, v31)
                    if v1 {
                        v33 := GenS3{F0: 0, F1: genFn0()}
                        v34 := GenS4{F0: v3, F1: 65535, F2: 42}
                        switch fi22 {
                        case 32768:
                            if len(v29) > 0 {
                                v29[0] = fi23
                            }
                        case 1:
                            if v1 {
                                fmt.Println(fi22)
                                var v35 uint32 = 42
                                var v36 uint = uint(v31)
                                _ = v35
                                _ = v36
                            } else {
                                fmt.Println(v1, v1, fi22)
                                v31 = v31
                            }
                            fmt.Println(v31, fi22)
                        case -360:
                            switch "" {
                            case "a":
                                if len(v29) > 0 {
                                    v29[0] = (fi22 | int(v3))
                                }
                            case "\n\t":
                                fmt.Println(fi22)
                            case "foo bar":
                                if len(v29) >= 2 {
                                    v29 = v29[1:len(v29)]
                                }
                            default:
                                v20 += 1
                            }
                        default:
                            v21 = (^(v21))
                        }
                        fmt.Println(v20, v20, v31)
                        var v37 int64 = 9223372036854775807
                        _ = v33
                        _ = v34
                        _ = v37
                    } else {
                        v31 = v31
                        var v38 uint = 18446744073709551615
                        fmt.Println(v1, fi23, v20)
                        _ = v38
                    }
                } else {
                    var v39 int16 = -1
                    _ = v39
                }
                _ = v29
                _ = v30
                _ = v31
                _ = v32
            }
            var v40 float32 = 0.1
            var v41 float32 = 1e10
            _ = fi23
            _ = v40
            _ = v41
        }
        v42 := GenS1{F0: 42, F1: 423}
        fmt.Println(v1, v20)
        _ = fi22
        _ = v42
    }
    var v43 int16 = 426
    if true {
        v44 := []int{}
        v21 -= v21
        _ = v44
    } else {
        var v45 uint64 = 1
        _ = v45
    }
    v46 := []int{-9223372036854775808, genFn0(), 1}
    _ = len(v46)
    fmt.Println(v43, v2, v20)
    var v47 float64 = (v2 / (v2 + 0.0))
    v48 := GenS5{F0: 2147483648}
    v43 = v43
    var v49 int = -75
    if (149 == 0) {
        if v1 {
            var v50 uint = 598
            if true {
                if ((!v1) || (!(42 <= 2))) {
                    switch "\n\t" {
                    case "":
                        var v51 uint = (v50 | (v50 << ((v50) & 63)))
                        _ = v51
                    }
                    fmt.Println(v21)
                    fmt.Println(v21)
                } else {
                    var v52 uint16 = (^(v21))
                    v2 += float64(((0.0 * float32(v47)) * -2.25))
                    if (v1 && (!v1)) {
                        switch "hello" {
                        case "\n\t":
                            v50 -= (v50 << (((v50 | (v50 - 18446744073709551615))) & 63))
                        case "foo bar":
                            fmt.Println(v52, v50, v3)
                        }
                        if (v49 > 455) {
                            for fi53 := 0; fi53 < 13; fi53++ {
                                var v54 int8 = (v20 | v20)
                                v55 := GenS6{F0: v50}
                                _ = fi53
                                _ = v54
                                _ = v55
                            }
                            if true {
                                fmt.Println(v47)
                                switch (v49 % ((-9223372036854775808) | 1)) {
                                case -247:
                                    var v56 int = v49
                                    _ = v56
                                case -2147483648:
                                    var v57 bool = v1
                                    _ = v57
                                case 0:
                                    if ((^(v21)) < 65535) {
                                        switch "\n\t" {
                                        case "a":
                                            v58 := GenS3{F0: 901, F1: (v49 >> (((^(v50))) & 63))}
                                            _ = v58
                                        default:
                                            for fi59 := 0; fi59 < 16; fi59++ {
                                                var v60 bool = (!(-32768 >= (v43 >> (((v50 & 260)) & 63))))
                                                _ = fi59
                                                _ = v60
                                            }
                                        }
                                    } else {
                                        v61 := GenS6{F0: 65535}
                                        _ = v61
                                    }
                                default:
                                    for fi62 := 0; fi62 < 3; fi62++ {
                                        v63 := GenS7{F0: v20}
                                        _ = fi62
                                        _ = v63
                                    }
                                }
                            } else {
                                v21 += (v21 >> ((v50) & 63))
                                var v64 bool = (false && v1)
                                v43 = 0
                                _ = v64
                            }
                        } else {
                            var v65 float32 = 3.14159
                            v43 = (-(v43))
                            _ = v65
                        }
                    } else {
                        if v1 {
                            var v66 int64 = -1
                            _ = v66
                        } else {
                            v1 = false
                        }
                    }
                    _ = v52
                }
                v67 := GenS2{F0: (v50 << (((^(v50))) & 63)), F1: 0}
                _ = len(v46)
                _ = v67
            } else {
                if (!((((("hello" + "foo bar") + ("\n\t" + "a")) + "") >= "") || v1)) {
                    if len(v46) > 0 {
                        v46[0] = -1
                    }
                } else {
                    var v68 float64 = (v47 / -2.25)
                    _ = v68
                }
            }
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v47:", v47)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v43
    _ = v46
    _ = v47
    _ = v48
    _ = v49
}
