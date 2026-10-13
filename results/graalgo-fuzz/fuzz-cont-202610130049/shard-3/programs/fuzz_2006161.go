package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint; F1 string; F2 int }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string) int {
    switch "hello" {
    case "":
        p0 = (((p0 + p0) + "foo bar") + "")
    case "a":
        p0 = ("" + p0)
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -9223372036854775808
}

func main() {
    var v1 bool = (-2147483648 >= 4294967295)
    var v2 float32 = (0.0 + -2.25)
    var v3 uint32 = 65535
    var v4 int8 = -128
    var v5 uint16 = 32767
    var v6 bool = false
    var v7 int32 = 42
    if true {
        var v8 uint32 = (^(v3))
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int32 = (v7 % (((v7 / ((v7) | 1))) | 1))
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        v11 := []int{1, 32768}
        v12 := GenS1{F0: (v3 << ((42) & 63)), F1: (v7 >> ((744) & 63))}
        _ = v11
        _ = v12
    }
    var v13 int64 = 9223372036854775807
    if true {
        if v6 {
            v14 := []int{-970, genFn0((("foo bar" + ("" + "hello")) + "\n\t")), 1, 116}
            if ((v7 + v7) >= -50) {
                v15 := GenS1{F0: v3, F1: (v7 / ((v7) | 1))}
                var v16 uint8 = 16
                _ = v15
                _ = v16
            } else {
                var v17 string = ("hello" + "a")
                var v18 int64 = -1
                _ = v17
                _ = v18
            }
            _ = v14
        }
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 float64 = 0.1
            fmt.Println(v1)
            v21 := []int{127, fi19}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if (-287 > 1) {
        v7 -= v7
        v5 += 255
        v22 := []int{1, int(v3), 9223372036854775807, -715}
        _ = v22
    } else {
        switch 32767 {
        case -509:
            if v1 {
                fmt.Println(v5, v4)
            } else {
                if v1 {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if true {
                            var v24 int64 = (-(v13))
                            if (((("hello" > "hello") && v6) && false) || (v6 && v1)) {
                                fmt.Println(v3, v3, v7)
                            }
                            _ = v24
                        }
                        var v25 int8 = v4
                        _ = fi23
                        _ = v25
                    }
                    v26 := []int{2147483647, 4294967295, -942, genFn0("hello")}
                    v27 := GenS2{F0: "foo bar", F1: 0, F2: (v13 + 1)}
                    _ = v26
                    _ = v27
                }
                var v28 uint16 = v5
                _ = v28
            }
        case 511:
            switch -26 {
            case -2147483648:
                v13 = v13
            case 0:
                for fi29 := 0; fi29 < 8; fi29++ {
                    for fi30 := 0; fi30 < 4; fi30++ {
                        if (v7 != -465) {
                            if v6 {
                                var v31 int32 = v7
                                _ = v31
                            } else {
                                if v1 {
                                    v32 := GenS3{F0: fi30, F1: 128, F2: v4}
                                    _ = v32
                                } else {
                                    v33 := GenS4{F0: 1, F1: "foo bar", F2: (^(fi30))}
                                    _ = v33
                                }
                                var v34 string = "\n\t"
                                switch genFn0(v34) {
                                case 4294967295:
                                    v5 += v5
                                case 42:
                                    fmt.Println(v3, v4)
                                case 127:
                                    v35 := GenS3{F0: fi29, F1: 2, F2: (v4 * v4)}
                                    v36 := GenS2{F0: (v34 + v34), F1: 42, F2: v13}
                                    _ = v35
                                    _ = v36
                                default:
                                    if (4294967295 >= 481) {
                                        var v37 int8 = (v4 & (v4 + -70))
                                        var v38 int = fi30
                                        var v39 uint16 = 0
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = v34
                            }
                            v40 := []int{fi30, (fi30 - (fi29 - (fi29 % ((fi30) | 1)))), (fi30 | fi30), genFn0("hello")}
                            var v41 uint32 = (v3 >> ((18446744073709551615) & 63))
                            _ = v40
                            _ = v41
                        }
                        _ = fi30
                    }
                    for fi42 := 0; fi42 < 3; fi42++ {
                        if (v1 && false) {
                            v43 := []int{int(v4), 32767}
                            _ = v43
                        }
                        _ = fi42
                    }
                    _ = fi29
                }
            }
        default:
            v44 := GenS3{F0: 2147483647, F1: 1, F2: (v4 + (v4 / ((v4) | 1)))}
            _ = v44
        }
        v45 := GenS3{F0: 32767, F1: 1, F2: -128}
        _ = v45
    }
    var v46 uint8 = 69
    v6 = v6
    v7 += v7
    v47 := GenS5{F0: 1, F1: v1, F2: 201}
    for fi48 := 0; fi48 < 6; fi48++ {
        v49 := GenS6{F0: "a"}
        for fi50 := 0; fi50 < 18; fi50++ {
            var v51 int = 2147483647
            _ = fi50
            _ = v51
        }
        v6 = ((!("" >= "\n\t")) || v6)
        _ = fi48
        _ = v49
    }
    v52 := GenS2{F0: ("\n\t" + "a"), F1: 4294967295, F2: (v13 ^ (^(v13)))}
    for fi53 := 0; fi53 < 9; fi53++ {
        if true {
            v3 = 2147483647
            v54 := GenS7{F0: v3, F1: (v46 - (v46 - uint8(v3)))}
            _ = v54
        } else {
            switch (fi53 / (((fi53 | 2147483647)) | 1)) {
            case 0:
                v7 = -46
            }
            v6 = v6
            for fi55 := 0; fi55 < 17; fi55++ {
                v56 := GenS6{F0: "\n\t"}
                v2 = v2
                _ = fi55
                _ = v56
            }
        }
        v5 -= (v5 | v5)
        var v57 uint = 42
        _ = fi53
        _ = v57
    }
    if v6 {
        var v58 int16 = 42
        v58 = v58
        for fi59 := 0; fi59 < 15; fi59++ {
            for fi60 := 0; fi60 < 9; fi60++ {
                var v61 uint8 = v46
                _ = fi60
                _ = v61
            }
            v62 := GenS5{F0: (v58 / (((v58 | v58)) | 1)), F1: v6, F2: 440}
            _ = fi59
            _ = v62
        }
        _ = v58
    }
    var v63 float64 = 3.14159
    var v64 uint64 = 689
    fmt.Println(v4, v7, v64)
    v65 := []int{-48, 127}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v52
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v7
}
