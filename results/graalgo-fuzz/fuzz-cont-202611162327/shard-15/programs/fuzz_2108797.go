package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 int8; F1 string; F2 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int64; F2 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 bool; F2 uint16 }
type GenS8 struct { F0 string; F1 uint; F2 uint16 }
type GenS9 struct { F0 int; F1 string }
func genFn0() int64 {
    v7 := GenS1{F0: 1, F1: -45, F2: 1}
    v8 := GenS2{F0: 44, F1: (("foo bar" + ("\n\t" + "hello")) + (("a" + ("hello" + "a")) + "a")), F2: false}
    _ = v7
    _ = v8
    return 9223372036854775807
}
func genFn1() int32 {
    if true {
        var v9 bool = false
        v9 = ((float32(float64((0.1 * 0.0))) - 3.14159) <= float32(-2.25))
        for fi10 := 0; fi10 < 8; fi10++ {
            v9 = (!(v9 && v9))
            v9 = (((598 == 793) && true) && (2 == 2))
            var v11 float32 = 3.14159
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        var v12 uint8 = 45
        var v13 string = ""
        fmt.Println(v13)
        _ = v12
        _ = v13
    }
    v14 := GenS3{F0: ""}
    var v15 float32 = 0.0
    v16 := []int{-5, 9223372036854775807}
    _ = v14
    _ = v15
    _ = v16
    return 2147483647
}

func main() {
    v1 := []int{-2147483648, 127, 32767, 127}
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    var v4 string = "\n\t"
    v5 := make(map[int]int)
    var v6 uint = 237
    if false {
        v17 := GenS1{F0: 32767, F1: -79, F2: 100}
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                if true {
                    if (false || ((fi18 << (((v6 << (((^(v6))) & 63))) & 63)) <= (fi19 * 128))) {
                        switch fi18 {
                        case 2147483647:
                            if ((v6 * v6) > 1) {
                                switch v4 {
                                case "a":
                                    v4 = (((v4 + ((v4 + v4) + v4)) + (v4 + v4)) + v4)
                                case "":
                                    fmt.Println(fi18, v6, fi19)
                                    switch (((("a" + (v4 + "foo bar")) + (v4 + v4)) + "\n\t") + v4) {
                                    case "foo bar":
                                        if true {
                                            var v20 uint8 = 128
                                            _ = v20
                                        } else {
                                            var v21 int8 = -61
                                            if false {
                                                v22 := v2["a"]
                                                _ = v22
                                            } else {
                                                v23 := GenS4{F0: 1}
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v24 string = v4
                                                var v25 int = fi19
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            var v26 uint = v6
                                            _ = v21
                                            _ = v26
                                        }
                                    case "":
                                        if false {
                                            fmt.Println(v4, v6)
                                        }
                                    case "\n\t":
                                        fmt.Println(fi19, v4, fi18)
                                    }
                                    if false {
                                        if (255 != 2) {
                                            v4 = ((v4 + v4) + "\n\t")
                                            _ = len(v5)
                                            var v27 int16 = 1000
                                            fmt.Println(fi19, v6, fi18)
                                            _ = v27
                                        }
                                        if ((((0.1 + 1e10) * float32(1e10)) == (3.14159 / float32(1e10))) || (4294967295 != 2147483647)) {
                                            switch "a" {
                                            case "foo bar":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        var v28 bool = ((0.0 > 1e10) && true)
                                                        _ = v28
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    var v29 uint16 = 0
                                                    _ = v29
                                                }
                                            case "a":
                                                v4 = v4
                                            case "":
                                                v4 = v4
                                            default:
                                                var v30 uint16 = 65535
                                                _ = v30
                                            }
                                        } else {
                                            var v31 int8 = 42
                                            v32 := GenS5{F0: 65535, F1: -9223372036854775808, F2: (^(v31))}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        var v33 uint32 = 2147483648
                                        fmt.Println(v6, v6, v33)
                                        delete(v5, (fi19 & fi19))
                                        _ = v33
                                    } else {
                                        var v34 int32 = -492
                                        var v35 uint16 = 255
                                        _ = v34
                                        _ = v35
                                    }
                                    v6 = 740
                                    v36 := GenS1{F0: -32768, F1: 1, F2: 31}
                                    v6 -= 2
                                    var v37 uint16 = 2
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v38 float64 = ((0.0 - 0.1) + 1.5)
                                        _ = v38
                                    }
                                    _ = v36
                                    _ = v37
                                case "\n\t":
                                    switch v4 {
                                    case "foo bar":
                                        var v39 bool = true
                                        _ = v39
                                    }
                                default:
                                    var v40 int64 = -9223372036854775808
                                    _ = v40
                                }
                            } else {
                                var v41 int16 = 192
                                _ = v41
                            }
                        }
                    } else {
                        v4 = "\n\t"
                        v4 = v4
                        v1 = append(v1, (fi19 / ((fi18) | 1)))
                        v1 = append(v1, int(v6))
                    }
                } else {
                    v2[(v4 + v4)] = 2147483648
                    var v42 int64 = 0
                    _ = v42
                }
                _ = fi19
            }
            _ = len(v1)
            _ = fi18
        }
        v6 = v6
        _ = v17
    }
    fmt.Println(v4, v6)
    v1 = append(v1, 902)
    v4 = "a"
    for fi43 := 0; fi43 < 15; fi43++ {
        v4 = "foo bar"
        v4 = ""
        _ = fi43
    }
    switch (v4 + ("\n\t" + "")) {
    case "hello":
        v44 := GenS6{F0: (1.5 != (1e10 + (1.5 + -2.25)))}
        _ = v44
    case "":
        if false {
            v1 = append(v1, -9223372036854775808)
            v4 = v4
            var v45 int = 128
            _ = v45
        }
    case "\n\t":
        v1 = append(v1, 0)
    default:
        for fi46 := 0; fi46 < 13; fi46++ {
            v47 := GenS3{F0: ((v4 + v4) + "a")}
            _ = fi46
            _ = v47
        }
    }
    v48 := GenS7{F0: false, F1: true, F2: 0}
    v6 = v6
    var v49 int8 = 100
    v50 := GenS8{F0: "hello", F1: (^(v6)), F2: 938}
    if false {
        var v51 int8 = (v49 | (v49 << (((v6 - (v6 << (((v6 & v6)) & 63)))) & 63)))
        var v52 uint8 = 127
        if (true || true) {
            if true {
                if (v6 != 186) {
                    if false {
                        v53 := GenS5{F0: 901, F1: genFn0(), F2: 127}
                        var v54 uint16 = 77
                        _ = v53
                        _ = v54
                    } else {
                        switch v4 {
                        case "\n\t":
                            var v55 int = 32768
                            _ = v55
                        case "foo bar":
                            switch (((v4 + ("\n\t" + "foo bar")) + (v4 + (((v4 + "foo bar") + ((v4 + "hello") + (((("" + ("hello" + "")) + v4) + "\n\t") + (v4 + (v4 + v4))))) + (("" + "\n\t") + "")))) + "hello") {
                            case "hello":
                                v52 = 127
                            }
                        default:
                            if ((1 < 2) && (0 < 0)) {
                                fmt.Println(v49, v4, v51)
                            }
                        }
                    }
                    v51 += -128
                    fmt.Println(v52, v49)
                } else {
                    v52 += v52
                    switch v4 {
                    case "a":
                        if false {
                            v52 = (v52 ^ (^(v52)))
                            delete(v3, 2)
                            v52 -= v52
                        } else {
                            v6 -= (^(v6))
                            for fi56 := 0; fi56 < 6; fi56++ {
                                switch (fi56 >> (((v6 - (v6 >> ((187) & 63)))) & 63)) {
                                case -1:
                                    var v57 int = fi56
                                    _ = v57
                                case -973:
                                    v58 := GenS9{F0: (fi56 ^ (-(fi56))), F1: v4}
                                    _ = v58
                                case 128:
                                    v59 := GenS7{F0: (int8(v52) < v49), F1: ((684 < 4294967295) || (255 >= 32767)), F2: 0}
                                    _ = v59
                                }
                                if true {
                                    if false {
                                        if false {
                                            var v60 uint64 = 4294967295
                                            for fi61 := 0; fi61 < 17; fi61++ {
                                                switch v4 {
                                                case "a":
                                                    switch (fi56 << ((42) & 63)) {
                                                    case -584:
                                                        var v62 uint16 = 745
                                                        _ = v62
                                                    case 32767:
                                                        var v63 bool = false
                                                        _ = v63
                                                    case 1:
                                                        var v64 uint = (^(v6))
                                                        _ = v64
                                                    }
                                                case "hello":
                                                    if ((v6 ^ 897) != v6) {
                                                        var v65 uint16 = 42
                                                        var v66 float64 = float64(-2.25)
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    fmt.Println(v49)
                                                }
                                                _ = fi61
                                            }
                                            _ = v60
                                        } else {
                                            if (665 != 9223372036854775807) {
                                                for fi67 := 0; fi67 < 8; fi67++ {
                                                    var v68 int8 = -33
                                                    v1 = append(v1, -719)
                                                    var v69 uint8 = (v52 | (^(v52)))
                                                    fmt.Println(fi56, v49)
                                                    _ = fi67
                                                    _ = v68
                                                    _ = v69
                                                }
                                                var v70 int8 = -1
                                                _ = v70
                                            }
                                            fmt.Println(v52, v4)
                                        }
                                        v52 = (^(v52))
                                    } else {
                                        v1 = append(v1, (fi56 % ((42) | 1)))
                                        var v71 string = (v4 + "a")
                                        _ = v71
                                    }
                                    fmt.Println(v49, fi56, v49)
                                    if true {
                                        v49 = 0
                                    } else {
                                        fmt.Println(v6)
                                    }
                                }
                                var v72 int32 = 42
                                _ = fi56
                                _ = v72
                            }
                        }
                    case "hello":
                        v73 := GenS2{F0: -92, F1: v4, F2: (true || (!(v52 < (v52 / (((v52 & 128)) | 1)))))}
                        _ = v73
                    case "":
                        if (true && true) {
                            v74 := GenS6{F0: (false && false)}
                            _ = v74
                        } else {
                            if (0 < genFn0()) {
                                v3[65535] = 230
                                for fi75 := 0; fi75 < 11; fi75++ {
                                    var v76 int32 = 1
                                    var v77 string = v4
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                                v78 := GenS7{F0: true, F1: ((true && (!true)) || false), F2: 65535}
                                _ = v78
                            }
                            var v79 string = v4
                            _ = v79
                        }
                    }
                }
            }
        }
        _ = v51
        _ = v52
    } else {
        fmt.Println(v49, v4)
        v80 := GenS1{F0: 1, F1: -128, F2: 6}
        _ = v80
    }
    v49 = (v49 << ((v6) & 63))
    v6 -= 522
    if (((float32(0.0) * 0.1) - (-2.25 + 0.0)) > (float32((3.14159 + (1e10 - float64(float32(((3.14159 * 1.5) - 1e10)))))) + 1e10)) {
        var v81 bool = true
        v82 := GenS1{F0: -625, F1: v49, F2: -9}
        v49 += 100
        _ = v81
        _ = v82
    }
    if len(v1) > 0 {
        v1[0] = 32767
    }
    fmt.Println(v6, v4, v4)
    delete(v5, 127)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
