package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint; F2 uint }
func genFn0(p0 bool, p1 int64, p2 int8) int64 {
    p0 = p0
    v6 := GenS1{F0: -2147483648}
    v7 := GenS1{F0: 301}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return int64(p2)
}

func main() {
    var v1 int = 128
    var v2 string = (((("" + "") + "a") + "foo bar") + ("hello" + "a"))
    var v3 uint32 = 175
    v4 := make(map[string]uint32)
    var v5 uint16 = 65535
    v8 := v4[(v2 + v2)]
    var v9 uint = 2
    v3 += 2147483648
    fmt.Println(v1, v5, v2)
    if (true || false) {
        var v10 uint16 = 264
        var v11 uint32 = (v3 % ((v3) | 1))
        _ = v10
        _ = v11
    }
    v2 = ("a" + v2)
    var v12 string = v2
    var v13 uint16 = v5
    if false {
        var v14 uint = (v9 / (((v9 % (((v9 % (((^(v9))) | 1))) | 1))) | 1))
        _ = v14
    } else {
        if true {
            for fi15 := 0; fi15 < 1; fi15++ {
                delete(v4, ("\n\t" + ""))
                fmt.Println(v9)
                var v16 int64 = 9223372036854775807
                _ = fi15
                _ = v16
            }
            if (465 == 42) {
                v3 = (v3 - (v3 % ((42) | 1)))
            }
        } else {
            _ = len(v4)
        }
        var v17 float64 = 1.5
        var v18 int16 = -32768
        _ = v17
        _ = v18
    }
    v19 := []int{v1, (-(v1)), (v1 << ((255) & 63)), (v1 << ((v9) & 63))}
    v20 := GenS1{F0: -428}
    if (!(42 != 95)) {
        v4[("a" + v2)] = v8
        var v21 int16 = 1
        if ((!((false && (((v2 >= v2) || (0.1 != (((1.5 + -2.25) - 3.14159) * 1.5))) || false)) && false)) && ((!false) && false)) {
            delete(v4, (("" + "hello") + "foo bar"))
            var v22 int16 = v21
            if (false && (0 >= 88)) {
                switch v2 {
                case "\n\t":
                    if (!true) {
                        _ = len(v4)
                    } else {
                        if (true && (1 != 9223372036854775807)) {
                            v19 = append(v19, 2147483647)
                            v21 = v21
                        } else {
                            var v23 int16 = (v22 >> ((42) & 63))
                            v12 = v12
                            _ = len(v19)
                            _ = v23
                        }
                        v24 := GenS2{F0: v8, F1: v22}
                        if (708 > (v13 & (v13 >> (((v9 - 4294967295)) & 63)))) {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                v26 := GenS1{F0: 0}
                                if false {
                                    switch "foo bar" {
                                    case "":
                                        for fi27 := 0; fi27 < 2; fi27++ {
                                            v19 = append(v19, fi27)
                                            var v28 int16 = -1
                                            var v29 uint = v9
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                    default:
                                        v5 -= v5
                                    }
                                    fmt.Println(v21)
                                } else {
                                    for fi30 := 0; fi30 < 12; fi30++ {
                                        if (3.14159 < (3.14159 + 3.14159)) {
                                            v31 := GenS2{F0: (v3 ^ (v3 << ((v9) & 63))), F1: v22}
                                            switch v1 {
                                            case 32768:
                                                switch "hello" {
                                                case "\n\t":
                                                    v8 = v3
                                                case "a":
                                                    fmt.Println(v8, fi30)
                                                default:
                                                    v4["foo bar"] = v3
                                                }
                                            case 441:
                                                var v32 uint = uint(v3)
                                                _ = v32
                                            default:
                                                if (!(true && false)) {
                                                    _ = len(v19)
                                                    if (3.14159 == 0.0) {
                                                        v4[(("" + "\n\t") + "hello")] = (v8 % (((^(v8))) | 1))
                                                        var v33 string = "a"
                                                        _ = v33
                                                    } else {
                                                        fmt.Println(v13)
                                                        var v34 int32 = 0
                                                        v3 = v8
                                                        _ = v34
                                                    }
                                                } else {
                                                    switch 0 {
                                                    case 0:
                                                        v5 = v5
                                                        if ((18446744073709551615 <= 0) && true) {
                                                            v35 := GenS1{F0: -21}
                                                            v36 := GenS2{F0: v8, F1: v21}
                                                            if ((42 > 1) || (genFn0(true, -1, -1) == genFn0((false && ((v8 > uint32(v5)) || true)), -1, -128))) {
                                                                fmt.Println(v12, v5)
                                                            } else {
                                                                v37 := v4["a"]
                                                                switch ((v12 + v2) + "foo bar") {
                                                                case "":
                                                                    var v38 string = "a"
                                                                    _ = v38
                                                                }
                                                                var v39 uint = v9
                                                                _ = v37
                                                                _ = v39
                                                            }
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        var v40 float64 = (1.5 / 1.5)
                                                        if true {
                                                            v12 = v12
                                                            _ = len(v4)
                                                            fmt.Println(v2, v22, v21)
                                                            var v41 float64 = 0.1
                                                            var v42 uint = 255
                                                            v1 = (-(fi30))
                                                            v43 := GenS1{F0: -2147483648}
                                                            _ = v41
                                                            _ = v42
                                                            _ = v43
                                                        } else {
                                                            v40 = -2.25
                                                        }
                                                        _ = v40
                                                    case 2147483647:
                                                        v13 = v5
                                                    case -718:
                                                        if (255 == 231) {
                                                            fmt.Println(v2, v9, v9)
                                                            v44 := v4[(v2 + v12)]
                                                            _ = v44
                                                        }
                                                    default:
                                                        v2 = (v12 + "a")
                                                    }
                                                }
                                            }
                                            _ = v31
                                        }
                                        if len(v19) > 0 {
                                            v19[0] = fi30
                                        }
                                        _ = fi30
                                    }
                                    var v45 bool = true
                                    _ = v45
                                }
                                _ = fi25
                                _ = v26
                            }
                            var v46 string = v12
                            v47 := GenS2{F0: 308, F1: 32767}
                            _ = v46
                            _ = v47
                        } else {
                            var v48 int32 = -271
                            var v49 int64 = -1
                            _ = v48
                            _ = v49
                        }
                        _ = v24
                    }
                }
                v12 = ((((v12 + ("a" + v2)) + ((v12 + v2) + "\n\t")) + v12) + (v12 + ("\n\t" + v2)))
                for fi50 := 0; fi50 < 5; fi50++ {
                    for fi51 := 0; fi51 < 7; fi51++ {
                        v52 := GenS3{F0: -330, F1: v9}
                        _ = len(v19)
                        _ = fi51
                        _ = v52
                    }
                    _ = fi50
                }
            }
            _ = v22
        } else {
            if true {
                v12 = (v12 + "\n\t")
                if (-1 < -76) {
                    v53 := GenS1{F0: 0}
                    v9 = (v9 | v9)
                    if false {
                        if ((^(v21)) >= (v21 * (v21 >> ((v9) & 63)))) {
                            v5 -= (v13 / ((uint16(v9)) | 1))
                            var v54 uint32 = (v3 >> (((v9 << ((912) & 63))) & 63))
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = v54
                        }
                        v21 = v21
                        for fi55 := 0; fi55 < 15; fi55++ {
                            switch (v12 + (v2 + (v2 + v2))) {
                            case "":
                                for fi56 := 0; fi56 < 18; fi56++ {
                                    var v57 uint8 = 255
                                    _ = fi56
                                    _ = v57
                                }
                            case "foo bar":
                                fmt.Println(v9, v3)
                            case "\n\t":
                                fmt.Println(v2, v8, v9)
                            }
                            switch "" {
                            case "a":
                                v1 -= fi55
                            case "hello":
                                for fi58 := 0; fi58 < 13; fi58++ {
                                    switch 127 {
                                    case -648:
                                        v21 = (^(v21))
                                    case 128:
                                        if (!false) {
                                            v59 := GenS2{F0: v8, F1: v21}
                                            if (false || (false && true)) {
                                                if (v8 < (v8 ^ v8)) {
                                                    v8 -= (^(v8))
                                                    v2 = v12
                                                    v60 := GenS2{F0: 42, F1: -1}
                                                    _ = v60
                                                }
                                                var v61 uint = 0
                                                fmt.Println(v2)
                                                var v62 int16 = 0
                                                v63 := GenS1{F0: 150}
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            } else {
                                                var v64 float32 = -2.25
                                                _ = v64
                                            }
                                            _ = v59
                                        }
                                    default:
                                        v65 := GenS4{F0: 1, F1: v8}
                                        var v66 int8 = 100
                                        fmt.Println(v66)
                                        _ = v65
                                        _ = v66
                                    }
                                    if false {
                                        var v67 uint64 = 953
                                        var v68 string = "hello"
                                        switch fi58 {
                                        case 128:
                                            var v69 int32 = 749
                                            _ = v69
                                        case 0:
                                            v70 := GenS3{F0: genFn0(true, -1, 127), F1: 2}
                                            _ = v70
                                        }
                                        _ = v67
                                        _ = v68
                                    } else {
                                        switch v12 {
                                        case "a":
                                            if (!true) {
                                                if (!true) {
                                                    var v71 bool = ((v2 + ((v2 + "") + ("" + "\n\t"))) > v2)
                                                    var v72 int32 = -954
                                                    _ = v71
                                                    _ = v72
                                                }
                                                var v73 int32 = -993
                                                fmt.Println(fi55, v21)
                                                v74 := GenS5{F0: (^(v3)), F1: 0, F2: 2}
                                                delete(v4, "a")
                                                fmt.Println(v2, v5)
                                                v9 = (v9 / (((v9 * v9)) | 1))
                                                _ = v73
                                                _ = v74
                                            }
                                        case "":
                                            var v75 int16 = (v21 + (v21 ^ int16(fi58)))
                                            _ = v75
                                        case "foo bar":
                                            fmt.Println(v12)
                                        default:
                                            var v76 int8 = -1
                                            _ = v76
                                        }
                                        _ = len(v4)
                                    }
                                    var v77 int32 = -2147483648
                                    _ = fi58
                                    _ = v77
                                }
                            case "foo bar":
                                delete(v4, v2)
                            }
                            var v78 uint32 = v8
                            _ = fi55
                            _ = v78
                        }
                    }
                    _ = v53
                } else {
                    v21 = 1
                    var v79 float32 = ((1e10 + (float32(((0.0 - 3.14159) - 0.0)) / 1.5)) + ((0.1 - -2.25) * 3.14159))
                    var v80 uint32 = v3
                    _ = v79
                    _ = v80
                }
            } else {
                var v81 int8 = -85
                v3 += v3
                _ = v81
            }
            var v82 uint = 822
            _ = v82
        }
        _ = v21
    }
    _ = len(v19)
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
