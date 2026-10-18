package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 uint32; F1 string; F2 string }
type GenS6 struct { F0 uint32; F1 int8; F2 string }
type GenS7 struct { F0 uint16 }
func genFn0(p0 string) uint32 {
    fmt.Println(p0)
    v6 := []int{128, 42}
    fmt.Println(p0)
    v7 := GenS1{F0: -1, F1: -1}
    p0 = (p0 + p0)
    _ = p0
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 uint) string {
    v8 := GenS1{F0: 4294967295, F1: 122}
    if (false || true) {
        var v9 int16 = 32767
        p1 += (p1 & 153)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 uint16 = 2
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            p0 = (p0 + ("" + "\n\t"))
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 42
    var v3 uint16 = 0
    v4 := []int{298, 32768, 1, 32767}
    var v5 float32 = 3.14159
    var v13 uint64 = 412
    var v14 int64 = 859
    var v15 uint = 65535
    fmt.Println(v14, v13, v13)
    for fi16 := 0; fi16 < 17; fi16++ {
        fmt.Println(v14)
        for fi17 := 0; fi17 < 7; fi17++ {
            v13 -= v13
            v3 -= (^(v3))
            _ = fi17
        }
        var v18 float32 = 0.0
        _ = fi16
        _ = v18
    }
    v2 += 255
    if (-828 <= 32768) {
        var v19 int32 = 703
        _ = v19
    }
    v14 += 1
    if (((v14 % (((v14 - v14)) | 1)) > v14) && (2147483647 != genFn0(("a" + (("\n\t" + "hello") + "\n\t"))))) {
        var v20 bool = true
        v21 := GenS1{F0: 1, F1: (^(v14))}
        if v20 {
            for fi22 := 0; fi22 < 5; fi22++ {
                switch (fi22 & fi22) {
                case -939:
                    var v23 float32 = v5
                    _ = v23
                default:
                    if v20 {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            fmt.Println(fi24)
                            v15 = uint(v3)
                            v25 := GenS1{F0: -225, F1: -1}
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        v26 := GenS2{F0: 42}
                        var v27 int64 = 0
                        _ = v26
                        _ = v27
                    }
                }
                fmt.Println(v5, v15)
                for fi28 := 0; fi28 < 20; fi28++ {
                    v29 := GenS3{F0: 780, F1: (-(v14)), F2: v13}
                    fmt.Println(fi22, v2)
                    v13 = (^(v13))
                    var v30 uint64 = 2
                    fmt.Println(v3, v13, fi22)
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi22
            }
            v31 := GenS3{F0: (v13 >> ((v15) & 63)), F1: -1, F2: (v13 % (((^(v13))) | 1))}
            _ = v31
        }
        _ = v20
        _ = v21
    }
    v2 = (v2 * (v3 * v2))
    v32 := GenS2{F0: 0}
    if (-2.25 > (-2.25 / ((((float64(v5) * 0.0) / 3.14159) * 0.0) - (-2.25 + 1.5)))) {
        for fi33 := 0; fi33 < 11; fi33++ {
            for fi34 := 0; fi34 < 9; fi34++ {
                var v35 float64 = 3.14159
                if (-1 > 1) {
                    v14 = (v14 | (-(v14)))
                    fmt.Println(v2, fi34, fi34)
                    var v36 int32 = -1
                    _ = v36
                } else {
                    fmt.Println(fi34)
                    if ((!(false && ((-1 >= -1) || true))) || true) {
                        switch "hello" {
                        case "hello":
                            fmt.Println(fi34, v14, v5)
                            v35 += 3.14159
                            if false {
                                v35 = v35
                                v14 -= v14
                                if true {
                                    v37 := GenS1{F0: fi33, F1: (v14 << (((v15 << (((v15 ^ v15)) & 63))) & 63))}
                                    v5 += -2.25
                                    fmt.Println(fi34, v35, v2)
                                    if false {
                                        fmt.Println(v5)
                                        v38 := GenS4{F0: v3, F1: 2}
                                        var v39 uint64 = (v13 | v13)
                                        _ = len(v1)
                                        _ = v38
                                        _ = v39
                                    } else {
                                        fmt.Println(fi34)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v1[42] = fi33
                                    }
                                    fmt.Println(v14)
                                    _ = v37
                                } else {
                                    var v40 uint = v15
                                    if true {
                                        if false {
                                            if (!((fi34 + fi33) == fi34)) {
                                                var v41 uint32 = 0
                                                v1[v41] = fi34
                                                v3 -= 686
                                                v42 := GenS5{F0: v41, F1: "hello", F2: "foo bar"}
                                                v13 -= (v13 + uint64(v40))
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 int8 = -104
                                                v43 = (^(v43))
                                                v13 = (v13 % ((1) | 1))
                                                _ = v43
                                            }
                                            delete(v1, 42)
                                            v13 = 2
                                            _ = len(v4)
                                            var v44 int64 = 720
                                            _ = v44
                                        }
                                    }
                                    _ = v40
                                }
                            }
                            fmt.Println(v13, fi34)
                            if (("foo bar" + "a") >= "") {
                                if false {
                                    switch ("a" + "foo bar") {
                                    case "":
                                        fmt.Println(v3)
                                    default:
                                        if (true && true) {
                                            v13 -= 0
                                            var v45 int64 = v14
                                            _ = v45
                                        }
                                    }
                                    v14 += (v14 % ((v14) | 1))
                                    var v46 int64 = (v14 ^ (-(v14)))
                                    if (0 > -382) {
                                        fmt.Println(v46, v13, fi33)
                                        v1[1] = fi34
                                        if false {
                                            if (1 <= 0) {
                                                if false {
                                                    switch (("a" + "foo bar") + "hello") {
                                                    case "a":
                                                        var v47 int16 = -944
                                                        _ = v47
                                                    default:
                                                        var v48 uint32 = 42
                                                        _ = v48
                                                    }
                                                }
                                                var v49 uint = v15
                                                _ = v49
                                            }
                                        } else {
                                            var v50 int32 = 42
                                            _ = v50
                                        }
                                    } else {
                                        v35 = float64((v5 - v5))
                                        fmt.Println(v14, fi33)
                                    }
                                    if true {
                                        if false {
                                            v14 -= v46
                                            var v51 float32 = (float32((0.0 / (v35 * v35))) - 1.5)
                                            var v52 uint32 = 4294967295
                                            var v53 bool = (true || false)
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        } else {
                                            var v54 uint32 = 0
                                            v2 += (^(v3))
                                            var v55 uint8 = 127
                                            _ = v54
                                            _ = v55
                                        }
                                    } else {
                                        var v56 uint8 = 190
                                        _ = v56
                                    }
                                    _ = v46
                                } else {
                                    if false {
                                        if (!false) {
                                            fmt.Println(fi34)
                                            switch "" {
                                            case "a":
                                                var v57 int16 = 32767
                                                _ = v57
                                            case "\n\t":
                                                v4 = append(v4, fi33)
                                                var v58 uint32 = 851
                                                _ = v58
                                            case "foo bar":
                                                var v59 string = ("" + ("\n\t" + "foo bar"))
                                                _ = v59
                                            default:
                                                v60 := GenS6{F0: 610, F1: -128, F2: "a"}
                                                _ = v60
                                            }
                                        }
                                    } else {
                                        var v61 int32 = -1
                                        fmt.Println(v5, fi33)
                                        var v62 uint8 = 255
                                        _ = v61
                                        _ = v62
                                    }
                                }
                            } else {
                                var v63 int64 = (v14 << ((v15) & 63))
                                _ = v63
                            }
                        }
                        v35 -= float64(1.5)
                        v64 := GenS1{F0: (fi33 - (fi33 ^ (fi33 + (fi33 % ((fi34) | 1))))), F1: v14}
                        _ = v64
                    } else {
                        var v65 string = "foo bar"
                        v14 = (v14 / (((v14 << ((213) & 63))) | 1))
                        var v66 uint32 = 1
                        var v67 bool = (v15 <= 0)
                        if true {
                            v68 := GenS7{F0: (^(v2))}
                            _ = v68
                        } else {
                            var v69 bool = v67
                            _ = v69
                        }
                        var v70 uint16 = v2
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v65
                        _ = v66
                        _ = v67
                        _ = v70
                    }
                }
                if len(v4) > 0 {
                    v4[0] = 0
                }
                var v71 uint16 = (v2 | 42)
                var v72 float32 = v5
                _ = fi34
                _ = v35
                _ = v71
                _ = v72
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v73 int16 = 1
            _ = fi33
            _ = v73
        }
    }
    v3 = v2
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
