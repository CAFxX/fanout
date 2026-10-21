package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 int32; F2 string }
type GenS5 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS6 struct { F0 uint16; F1 bool; F2 uint }
type GenS7 struct { F0 string; F1 int; F2 uint8 }
func genFn0(p0 int) uint64 {
    var v7 int8 = 26
    if ((1 > (v7 / ((v7) | 1))) && false) {
        var v8 uint32 = 2147483648
        var v9 uint = 65535
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint = 737
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v4 := []int{32768, 0}
    v5 := make(map[int]int)
    var v6 int = 127
    v10 := GenS1{F0: 978}
    var v11 uint32 = 964
    var v12 uint8 = 237
    v11 -= v11
    for fi13 := 0; fi13 < 14; fi13++ {
        v6 -= (v6 | fi13)
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: int64(v1)}
            _ = fi14
            _ = v15
        }
        switch v6 {
        case 727:
            var v16 int32 = -2147483648
            _ = v16
        case -444:
            switch fi13 {
            case -2147483648:
                if true {
                    v12 = 128
                    fmt.Println(v6)
                    switch "foo bar" {
                    case "hello":
                        v1 = 176
                    case "\n\t":
                        var v17 uint32 = 65535
                        _ = v17
                    default:
                        v1 = v1
                    }
                }
            case -1:
                v18 := GenS1{F0: -31}
                _ = v18
            default:
                if true {
                    switch v6 {
                    case -2147483648:
                        if (("" + "\n\t") >= "hello") {
                            var v19 int8 = -3
                            _ = v19
                        } else {
                            fmt.Println(v12)
                        }
                    case 2147483647:
                        if len(v4) > 0 {
                            v4[0] = v6
                        }
                    case 32768:
                        var v20 uint = 1
                        _ = v20
                    default:
                        v6 = fi13
                    }
                    var v21 uint64 = genFn0((fi13 % (((v6 * v6)) | 1)))
                    _ = v21
                }
            }
        case 77:
            if ((^(v11)) >= 1) {
                v22 := GenS1{F0: -9223372036854775808}
                for fi23 := 0; fi23 < 12; fi23++ {
                    if ((255 == v12) || true) {
                        v24 := GenS2{F0: -1}
                        var v25 float64 = float64(1e10)
                        _ = v24
                        _ = v25
                    }
                    if len(v4) > 0 {
                        v4[0] = (-(fi23))
                    }
                    if (!true) {
                        v12 = (v12 / ((uint8(v11)) | 1))
                    } else {
                        var v26 uint16 = 32767
                        if true {
                            if ((533 == 1000) || false) {
                                if ((v12 & (^(v12))) > 255) {
                                    v6 = (fi13 * 32767)
                                } else {
                                    v26 = (v26 / ((42) | 1))
                                    fmt.Println(v26)
                                    v27 := GenS3{F0: 0}
                                    _ = v27
                                }
                                if false {
                                    fmt.Println(fi23, v12, v6)
                                    v11 = v11
                                } else {
                                    switch fi23 {
                                    case -2147483648:
                                        v6 = (fi23 | v6)
                                        v28 := GenS2{F0: 0}
                                        var v29 int16 = 1
                                        _ = v28
                                        _ = v29
                                    case 69:
                                        var v30 string = "foo bar"
                                        v1 = (v1 & 0)
                                        _ = v30
                                    default:
                                        fmt.Println(v26)
                                    }
                                    fmt.Println(v12, v12, fi23)
                                    switch "\n\t" {
                                    case "hello":
                                        v31 := GenS4{F0: -474, F1: -413, F2: ("a" + ("" + "a"))}
                                        _ = v31
                                    case "\n\t":
                                        if false {
                                            fmt.Println(v6, v1)
                                        } else {
                                            v11 = (v11 * v11)
                                        }
                                    case "a":
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            fmt.Println(fi13, v11, v11)
                                        case "a":
                                            var v32 uint32 = (v11 % ((v11) | 1))
                                            _ = v32
                                        }
                                    default:
                                        if ((float64(1e10) - 0.0) >= (3.14159 * (float64(1e10) - float64(float32(float64(0.0)))))) {
                                            v33 := GenS4{F0: -9223372036854775808, F1: -1, F2: "hello"}
                                            _ = v33
                                        }
                                    }
                                    if ((v26 < (v26 / ((v26) | 1))) && true) {
                                        v12 = 255
                                        v34 := GenS3{F0: (^(v1))}
                                        fmt.Println(v12)
                                        _ = v34
                                    } else {
                                        if true {
                                            if true {
                                                if true {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v35 uint64 = uint64(v1)
                                                        var v36 float64 = 1.5
                                                        switch "" {
                                                        case "":
                                                            v37 := v5[v6]
                                                            fmt.Println(v1, v35, v6)
                                                            var v38 float64 = v36
                                                            _ = v37
                                                            _ = v38
                                                        default:
                                                            fmt.Println(fi13)
                                                        }
                                                        _ = v35
                                                        _ = v36
                                                    case "a":
                                                        v11 = (v11 ^ (v11 >> (((v1 >> ((v1) & 63))) & 63)))
                                                    default:
                                                        if len(v4) > 0 {
                                                            v4[0] = fi13
                                                        }
                                                        var v39 int8 = 100
                                                        _ = v39
                                                    }
                                                    v12 = v12
                                                }
                                            } else {
                                                var v40 uint16 = (^(v26))
                                                fmt.Println(v11, v1, v26)
                                                if false {
                                                    v6 = fi13
                                                    v6 = v6
                                                    v41 := GenS5{F0: (v12 % ((v12) | 1)), F1: 270, F2: 655}
                                                    _ = v41
                                                } else {
                                                    if ("" < "\n\t") {
                                                        v26 += (v26 % (((v26 << ((4294967295) & 63))) | 1))
                                                    }
                                                }
                                                _ = v40
                                            }
                                            var v42 uint16 = (v26 ^ v26)
                                            _ = v42
                                        }
                                        v43 := GenS6{F0: (v26 << (((^(v1))) & 63)), F1: (float32(float64(-2.25)) >= 1e10), F2: 804}
                                        _ = v43
                                    }
                                }
                                v44 := GenS6{F0: (v26 & uint16(v1)), F1: (false || true), F2: (v1 | 160)}
                                _ = v44
                            }
                            v45 := GenS4{F0: 92, F1: 645, F2: "foo bar"}
                            _ = v45
                        } else {
                            var v46 bool = (true && (false && (!(0.0 > 3.14159))))
                            _ = v46
                        }
                        _ = v26
                    }
                    var v47 uint = 817
                    _ = fi23
                    _ = v47
                }
                switch fi13 {
                case 0:
                    v1 = v1
                case 216:
                    var v48 float64 = -2.25
                    _ = v48
                case 1:
                    var v49 int = v6
                    _ = v49
                default:
                    if false {
                        v11 += (v11 | (v11 - v11))
                        if true {
                            v11 -= v11
                            switch 42 {
                            case 107:
                                fmt.Println(fi13, fi13)
                            case 127:
                                var v50 int16 = 1000
                                _ = v50
                            }
                            v51 := GenS7{F0: "a", F1: (fi13 << ((v1) & 63)), F2: 127}
                            _ = v51
                        }
                    }
                }
                _ = v22
            } else {
                var v52 uint = v1
                for fi53 := 0; fi53 < 13; fi53++ {
                    v6 += -230
                    _ = fi53
                }
                _ = v52
            }
        default:
            for fi54 := 0; fi54 < 18; fi54++ {
                fmt.Println(fi13, fi54)
                v3[fi54] = "\n\t"
                v55 := GenS7{F0: ((("a" + "\n\t") + ("" + "")) + "foo bar"), F1: 127, F2: (v12 | 109)}
                _ = fi54
                _ = v55
            }
        }
        _ = fi13
    }
    for fi56 := 0; fi56 < 15; fi56++ {
        var v57 string = "a"
        v11 += 726
        if (!true) {
            v12 -= v12
            switch 32768 {
            case -1:
                v1 -= uint(v12)
            case 128:
                v12 -= v12
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v58 uint8 = 0
            if false {
                var v59 uint = (v1 - (v1 - (v1 >> ((v1) & 63))))
                _ = v59
            } else {
                if len(v4) > 0 {
                    v4[0] = -897
                }
            }
            _ = v58
        }
        _ = fi56
        _ = v57
    }
    var v60 uint = (v1 | v1)
    switch v6 {
    case 0:
        var v61 bool = (2 < genFn0(v6))
        _ = v61
    case 42:
        if (-42 >= 0) {
            delete(v5, -1)
            v62 := v2[v11]
            v63 := GenS2{F0: 482}
            _ = v62
            _ = v63
        } else {
            fmt.Println(v60)
            var v64 string = "a"
            v65 := GenS1{F0: 42}
            _ = v64
            _ = v65
        }
    case 32767:
        var v66 uint16 = 42
        _ = v66
    }
    if (!(float64(3.14159) != 0.0)) {
        var v67 int8 = 127
        var v68 uint64 = 2
        _ = v67
        _ = v68
    } else {
        switch -371 {
        case 127:
            v12 += (v12 | v12)
        case 9223372036854775807:
            var v69 uint64 = 4294967295
            _ = v69
        default:
            if false {
                fmt.Println(v11)
                var v70 bool = true
                _ = v70
            }
        }
    }
    v71 := v5[-581]
    var v72 float64 = ((-2.25 + 0.0) + (0.0 / 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v71
    _ = v72
}
