package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 int; F1 uint8; F2 int }
type GenS5 struct { F0 uint64 }
func genFn0(p0 string, p1 int8, p2 uint64) int16 {
    if (4294967295 != 4294967295) {
        for fi6 := 0; fi6 < 13; fi6++ {
            if false {
                p0 = p0
            } else {
                switch "foo bar" {
                case "\n\t":
                    var v7 int8 = p1
                    _ = v7
                case "":
                    var v8 int64 = 1
                    _ = v8
                case "a":
                    p1 += (p1 * 100)
                default:
                    var v9 float32 = (3.14159 + float32(1e10))
                    _ = v9
                }
            }
            if true {
                fmt.Println(p2)
            } else {
                v10 := []int{}
                p1 = 42
                var v11 float64 = 1.5
                _ = v10
                _ = v11
            }
            _ = fi6
        }
    }
    var v12 float32 = -2.25
    var v13 string = p0
    v13 = ((p0 + p0) + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}
func genFn1() uint16 {
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 string = ""
        v15 = v15
        v16 := []int{(fi14 & (fi14 % ((128) | 1))), -2147483648}
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 838, F1: 0, F2: -175}
    var v18 int64 = -533
    if true {
        v19 := GenS2{F0: v18}
        var v20 string = "\n\t"
        _ = v19
        _ = v20
    }
    if false {
        var v21 int8 = 100
        if (32768 != int(v18)) {
            if false {
                v21 = (v21 >> ((42) & 63))
                for fi22 := 0; fi22 < 3; fi22++ {
                    if (!(338 > genFn0("hello", (v21 & (v21 + v21)), 674))) {
                        for fi23 := 0; fi23 < 15; fi23++ {
                            if false {
                                v24 := []int{}
                                var v25 uint32 = 42
                                switch fi22 {
                                case 32768:
                                    switch ("\n\t" + "") {
                                    case "foo bar":
                                        v26 := GenS2{F0: v18}
                                        _ = v26
                                    }
                                case -9223372036854775808:
                                    fmt.Println(fi22, fi23)
                                }
                                _ = v24
                                _ = v25
                            }
                            var v27 int32 = 0
                            fmt.Println(fi23, v21)
                            _ = fi23
                            _ = v27
                        }
                        v21 = v21
                    } else {
                        v21 += int8(fi22)
                    }
                    var v28 int32 = 865
                    _ = fi22
                    _ = v28
                }
                if true {
                    var v29 int16 = 700
                    _ = v29
                }
            }
            v18 = (v18 + (v18 | (v18 ^ v18)))
        }
        fmt.Println(v18)
        _ = v21
    }
    _ = v17
    _ = v18
    return 2
}

func main() {
    v1 := []string{("\n\t" + "hello"), "", "hello", ("hello" + ("hello" + "a"))}
    var v2 string = ("" + ("" + "a"))
    v3 := make(map[uint32]int)
    var v4 uint64 = 282
    var v5 int16 = 32767
    if false {
        switch 4294967295 {
        case -9223372036854775808:
            for fi30 := 0; fi30 < 11; fi30++ {
                v2 = v2
                _ = fi30
            }
        case 9223372036854775807:
            v4 -= (^(v4))
        }
        fmt.Println(v2, v5, v4)
        var v31 bool = (v5 != (v5 / (((v5 ^ genFn0("foo bar", -1, v4))) | 1)))
        _ = v31
    } else {
        if (3.14159 > (1.5 / float32(1e10))) {
            switch "foo bar" {
            case "hello":
                var v32 int16 = (v5 & -876)
                _ = v32
            case "\n\t":
                switch 1 {
                case -814:
                    if len(v1) > 0 {
                        v1[0] = v2
                    }
                case 42:
                    switch v2 {
                    case "":
                        v4 += (v4 & (^(v4)))
                    }
                case 32768:
                    for fi33 := 0; fi33 < 11; fi33++ {
                        _ = len(v3)
                        v34 := GenS3{F0: "", F1: true, F2: 65535}
                        _ = fi33
                        _ = v34
                    }
                default:
                    if (true || true) {
                        var v35 string = ""
                        var v36 uint8 = 1
                        _ = v35
                        _ = v36
                    } else {
                        _ = len(v3)
                    }
                }
            default:
                if (0.1 < 0.1) {
                    v5 -= genFn0(v2, -117, (^(v4)))
                    var v37 int64 = 1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v37
                } else {
                    var v38 int = 932
                    v39 := GenS4{F0: v38, F1: 127, F2: (v38 ^ v38)}
                    _ = v38
                    _ = v39
                }
            }
        }
        v4 = 186
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v41 := v3[2]
        fmt.Println(v2)
        v5 = (v5 | (v5 + (v5 & genFn0(v2, 100, 1))))
        _ = fi40
        _ = v41
    }
    var v42 float64 = (0.0 * ((float64(1e10) * (1e10 * 1.5)) * (0.1 - float64(0.1))))
    var v43 int64 = -410
    fmt.Println(v4, v43, v4)
    v44 := GenS4{F0: 127, F1: 209, F2: 448}
    v4 -= v4
    v5 = (-(v5))
    var v45 uint32 = 0
    fmt.Println(v45, v5, v45)
    fmt.Println(v5, v5)
    switch 333 {
    case 9223372036854775807:
        fmt.Println(v5, v42)
    case -9223372036854775808:
        var v46 int16 = (v5 & v5)
        _ = v46
    case 32767:
        if (!(v4 > (v4 >> ((2) & 63)))) {
            switch 529 {
            case -2147483648:
                var v47 string = v2
                _ = v47
            case -911:
                for fi48 := 0; fi48 < 8; fi48++ {
                    var v49 bool = (!false)
                    var v50 bool = (v49 && ((-(v5)) >= v5))
                    if (2 < 128) {
                        v51 := GenS1{F0: 603, F1: (^(v5)), F2: -274}
                        v52 := GenS4{F0: (fi48 << ((953) & 63)), F1: 2, F2: (fi48 / (((fi48 ^ fi48)) | 1))}
                        v53 := GenS5{F0: (v4 | v4)}
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    _ = fi48
                    _ = v49
                    _ = v50
                }
            case 943:
                var v54 uint8 = 214
                _ = v54
            default:
                v4 = uint64(v45)
            }
            delete(v3, (^(v45)))
            delete(v3, (v45 - v45))
        }
    }
    var v55 int64 = int64(v45)
    v5 -= (v5 ^ -1)
    var v56 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v55
    _ = v56
}
