package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 string; F1 int8 }
type GenS8 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 bool) int {
    var v5 uint32 = 52
    fmt.Println(v5, v5)
    p0 = p0
    if (-32768 < 32767) {
        v5 = 1
        if (!p0) {
            v5 = v5
            v6 := GenS1{F0: (^(v5))}
            v7 := []int{32767, 127, 127}
            _ = v6
            _ = v7
        } else {
            v8 := []int{32768}
            if (!(-899 != 569)) {
                p0 = p0
            } else {
                var v9 uint16 = 255
                _ = v9
            }
            switch 1 {
            case 572:
                p0 = p0
            case 0:
                if len(v8) > 0 {
                    v8[0] = 9223372036854775807
                }
            case 32768:
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 int64 = 0
                    v12 := GenS1{F0: (v5 / ((v5) | 1))}
                    var v13 string = (((("foo bar" + "hello") + "") + "hello") + "\n\t")
                    _ = fi10
                    _ = v11
                    _ = v12
                    _ = v13
                }
            default:
                v5 += (v5 * (^(v5)))
            }
            _ = v8
        }
        var v14 float32 = (0.0 - (-2.25 - 1.5))
        _ = v14
    }
    var v15 uint32 = (v5 << ((439) & 63))
    _ = p0
    _ = v15
    _ = v5
    return 128
}
func genFn1(p0 uint) int {
    v16 := GenS2{F0: 255, F1: "a", F2: 4294967295}
    p0 -= (p0 << (((p0 * (p0 << ((p0) & 63)))) & 63))
    var v17 int8 = 127
    v18 := GenS2{F0: 255, F1: "hello", F2: 2}
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return genFn0(false)
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    var v4 int64 = -890
    if true {
        var v19 int64 = v4
        v20 := GenS1{F0: 0}
        v4 = 153
        _ = v19
        _ = v20
    } else {
        v21 := GenS1{F0: 202}
        _ = v21
    }
    var v22 int = genFn0((128 > 2147483647))
    fmt.Println(v4)
    if false {
        if (2147483648 <= 808) {
            var v23 float64 = (((-2.25 / 0.1) + -2.25) / 1.5)
            if true {
                v2[""] = "\n\t"
            }
            fmt.Println(v23, v4, v1)
            _ = v23
        } else {
            v24 := GenS3{F0: 42}
            _ = v24
        }
    } else {
        var v25 float32 = ((-2.25 / 1e10) + 0.0)
        if (-1 > 0) {
            var v26 int8 = 127
            v22 = v22
            fmt.Println(v22, v22, v1)
            _ = v26
        }
        var v27 int64 = 959
        _ = v25
        _ = v27
    }
    if (2147483648 <= 2) {
        if true {
            if (!(false || true)) {
                v4 -= (v4 - v4)
                var v28 uint16 = 32767
                switch "\n\t" {
                case "a":
                    switch "a" {
                    case "hello":
                        v29 := []int{v22, (v22 >> ((18446744073709551615) & 63)), 32768, (v22 / ((v22) | 1))}
                        _ = v29
                    case "":
                        for fi30 := 0; fi30 < 20; fi30++ {
                            v31 := []int{}
                            for fi32 := 0; fi32 < 11; fi32++ {
                                v33 := GenS4{F0: 50}
                                if (false || (false && (((v4 << ((1) & 63)) < (v4 | v4)) || true))) {
                                    var v34 int64 = (-(v4))
                                    var v35 string = "a"
                                    delete(v3, (fi32 | (fi30 * (fi30 | (fi30 >> ((611) & 63))))))
                                    _ = v34
                                    _ = v35
                                } else {
                                    if false {
                                        v4 += 1
                                        v1 = v28
                                        var v36 int32 = -2147483648
                                        _ = v36
                                    }
                                }
                                switch v22 {
                                case 32768:
                                    v28 = 0
                                case 0:
                                    var v37 int64 = (v4 * (^(v4)))
                                    fmt.Println(fi32)
                                    var v38 uint = 142
                                    _ = v37
                                    _ = v38
                                case -2147483648:
                                    if true {
                                        var v39 float64 = float64(0.1)
                                        if false {
                                            var v40 uint16 = 336
                                            var v41 uint = 1
                                            var v42 float64 = 0.0
                                            v43 := v3[(fi32 ^ v22)]
                                            _ = v40
                                            _ = v41
                                            _ = v42
                                            _ = v43
                                        }
                                        if false {
                                            v44 := GenS2{F0: 255, F1: "", F2: 42}
                                            if len(v31) > 0 {
                                                v31[0] = v22
                                            }
                                            switch fi30 {
                                            case -122:
                                                if (64 > 2) {
                                                    var v45 uint = 0
                                                    v46 := GenS4{F0: 1}
                                                    _ = v45
                                                    _ = v46
                                                }
                                            case 1:
                                                var v47 uint32 = 65535
                                                _ = v47
                                            default:
                                                if true {
                                                    delete(v2, (((("foo bar" + "hello") + "") + "foo bar") + ""))
                                                    if (!true) {
                                                        v48 := GenS1{F0: 2147483648}
                                                        _ = v48
                                                    } else {
                                                        if ("foo bar" < "hello") {
                                                            v22 = (fi32 >> ((4294967295) & 63))
                                                        }
                                                    }
                                                }
                                            }
                                            if ((false && (!false)) && true) {
                                                v49 := GenS1{F0: 2147483647}
                                                v22 += genFn1(65535)
                                                if true {
                                                    if true {
                                                        var v50 float64 = -2.25
                                                        if len(v31) >= 2 {
                                                            v31 = v31[1:len(v31)]
                                                        }
                                                        v39 += v39
                                                        _ = v50
                                                    }
                                                    fmt.Println(v22, v4)
                                                }
                                                _ = v49
                                            } else {
                                                fmt.Println(v22)
                                            }
                                            v39 += (((1e10 - (-2.25 - (v39 + v39))) / v39) + 0.0)
                                            _ = v44
                                        }
                                        v22 = (v22 / ((genFn1(4294967295)) | 1))
                                        _ = v39
                                    }
                                }
                                v51 := GenS5{F0: 122}
                                _ = fi32
                                _ = v33
                                _ = v51
                            }
                            switch (-(v22)) {
                            case 1:
                                delete(v3, (v22 ^ fi30))
                            }
                            _ = fi30
                            _ = v31
                        }
                    default:
                        for fi52 := 0; fi52 < 14; fi52++ {
                            v2["foo bar"] = "a"
                            var v53 uint16 = v28
                            for fi54 := 0; fi54 < 4; fi54++ {
                                v28 = v28
                                _ = fi54
                            }
                            _ = fi52
                            _ = v53
                        }
                    }
                case "\n\t":
                    v55 := GenS6{F0: (v4 << ((42) & 63)), F1: (-1 != 42)}
                    _ = v55
                }
                _ = v28
            }
            v56 := []int{32768}
            _ = v56
        } else {
            v1 = (v1 ^ v1)
            for fi57 := 0; fi57 < 18; fi57++ {
                switch "\n\t" {
                case "\n\t":
                    v58 := GenS1{F0: 42}
                    _ = v58
                default:
                    if true {
                        v59 := GenS1{F0: 4294967295}
                        v60 := GenS6{F0: v4, F1: true}
                        for fi61 := 0; fi61 < 13; fi61++ {
                            v62 := GenS7{F0: "hello", F1: -1}
                            var v63 float64 = 1.5
                            switch (-(fi57)) {
                            case -109:
                                v64 := GenS7{F0: "\n\t", F1: int8(v1)}
                                _ = v64
                            case 1:
                                v65 := v3[genFn0((!(2 <= 255)))]
                                _ = v65
                            case -1:
                                _ = len(v2)
                            }
                            _ = fi61
                            _ = v62
                            _ = v63
                        }
                        _ = v59
                        _ = v60
                    }
                }
                _ = fi57
            }
        }
        v22 -= -9223372036854775808
        _ = len(v2)
    }
    fmt.Println(v1)
    _ = len(v3)
    v3[-1] = 959
    var v66 uint = 42
    switch genFn0(((v66 | (v66 ^ 0)) == v66)) {
    case 404:
        var v67 float64 = 0.0
        _ = v67
    default:
        v68 := GenS8{F0: true, F1: 1, F2: (-2.25 > ((0.1 / 1.0) + ((((0.0 * (1.5 * 1e10)) * 1e10) + (-2.25 / 1e10)) * (((1.5 - 0.0) - 0.1) - (0.0 / float64(1.5))))))}
        _ = v68
    }
    v66 = v66
    var v69 uint8 = uint8(v1)
    for fi70 := 0; fi70 < 7; fi70++ {
        v71 := []int{v22, (^(v22)), (v22 & 907)}
        _ = fi70
        _ = v71
    }
    v72 := []int{(v22 >> ((v66) & 63)), 32767, v22, v22}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v66
    _ = v69
    _ = v72
}
