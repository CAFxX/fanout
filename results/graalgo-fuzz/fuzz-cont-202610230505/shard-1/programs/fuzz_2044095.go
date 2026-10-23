package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int8; F1 int64 }

func main() {
    var v1 int16 = -114
    var v2 float32 = 0.1
    var v3 float64 = (1e10 - 0.1)
    v4 := []string{"\n\t", "hello", "", "hello"}
    v5 := []int{-316, 127, 4294967295, 32767}
    var v6 bool = false
    var v7 uint32 = 1
    var v8 float32 = 0.0
    for fi9 := 0; fi9 < 8; fi9++ {
        if (false && (((((("hello" + "\n\t") + "") + ("foo bar" + ("a" + ""))) + "hello") + "\n\t") > "a")) {
            var v10 int = fi9
            _ = v10
        } else {
            for fi11 := 0; fi11 < 14; fi11++ {
                if (0 > -128) {
                    fmt.Println(fi11, v6)
                    v7 -= (v7 * 567)
                    v5 = append(v5, fi9)
                }
                _ = fi11
            }
            fmt.Println(v7)
            _ = len(v5)
        }
        v1 = 1000
        _ = fi9
    }
    if (32768 >= 127) {
        switch "" {
        case "foo bar":
            v6 = (221 >= uint8(v7))
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "a":
            var v12 bool = (((v8 * v8) <= float32(((1.5 / (v3 - -2.25)) + 1e10))) || v6)
            _ = v12
        }
        switch -348 {
        case 32768:
            fmt.Println(v6, v8)
        default:
            switch "foo bar" {
            case "a":
                var v13 int = 32767
                _ = v13
            case "foo bar":
                for fi14 := 0; fi14 < 20; fi14++ {
                    _ = len(v5)
                    v7 = 2147483647
                    if v6 {
                        switch "foo bar" {
                        case "a":
                            for fi15 := 0; fi15 < 12; fi15++ {
                                fmt.Println(v6, v7)
                                _ = fi15
                            }
                        case "":
                            v1 -= (v1 & v1)
                        }
                    }
                    _ = fi14
                }
            default:
                v16 := GenS1{F0: ("a" + ""), F1: v1, F2: v7}
                _ = v16
            }
        }
        v17 := GenS1{F0: "\n\t", F1: (-(v1)), F2: (v7 >> ((4294967295) & 63))}
        _ = v17
    }
    if v6 {
        v18 := GenS1{F0: "hello", F1: 348, F2: v7}
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        if ((23 <= 100) || v6) {
            var v19 uint32 = (v7 + (^(v7)))
            _ = v19
        } else {
            fmt.Println(v8, v8, v7)
            v2 += 0.1
            switch "hello" {
            case "foo bar":
                fmt.Println(v8, v6, v2)
            case "hello":
                if (v6 && v6) {
                    switch "\n\t" {
                    case "foo bar":
                        v7 -= (v7 << ((435) & 63))
                    case "hello":
                        switch "foo bar" {
                        case "":
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                        case "\n\t":
                            v2 = (0.0 - v8)
                        }
                    case "":
                        for fi20 := 0; fi20 < 12; fi20++ {
                            if v6 {
                                var v21 uint16 = 171
                                if ((-9223372036854775808 != 0) && (true || (2147483647 < 2147483647))) {
                                    for fi22 := 0; fi22 < 13; fi22++ {
                                        fmt.Println(v21, v21, v7)
                                        _ = fi22
                                    }
                                }
                                v5 = append(v5, fi20)
                                _ = v21
                            } else {
                                var v23 float64 = v3
                                _ = v23
                            }
                            if len(v4) > 0 {
                                v4[0] = "hello"
                            }
                            fmt.Println(v2, v6)
                            _ = fi20
                        }
                    }
                    if ("\n\t" > ("a" + "\n\t")) {
                        v24 := GenS2{F0: (("hello" + "") + "\n\t")}
                        v25 := GenS1{F0: "a", F1: v1, F2: v7}
                        if (100 > 10) {
                            switch "\n\t" {
                            case "":
                                for fi26 := 0; fi26 < 18; fi26++ {
                                    v27 := GenS3{F0: (v7 * v7), F1: 2147483647, F2: v7}
                                    v1 -= 32767
                                    switch "\n\t" {
                                    case "":
                                        var v28 float32 = ((float32(v3) * v2) / v2)
                                        _ = v28
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                            case "\n\t":
                                if true {
                                    switch 1 {
                                    case 6:
                                        v29 := GenS1{F0: "", F1: (v1 / ((v1) | 1)), F2: v7}
                                        _ = v29
                                    case -9223372036854775808:
                                        if v6 {
                                            _ = len(v5)
                                        }
                                    default:
                                        if (!v6) {
                                            if (!false) {
                                                for fi30 := 0; fi30 < 6; fi30++ {
                                                    v31 := GenS1{F0: "foo bar", F1: v1, F2: 762}
                                                    for fi32 := 0; fi32 < 16; fi32++ {
                                                        v2 += ((v8 * 1e10) + (v2 - v8))
                                                        if v6 {
                                                            var v33 int = -2147483648
                                                            fmt.Println(fi32, fi32, fi32)
                                                            switch "a" {
                                                            case "a":
                                                                if v6 {
                                                                    v6 = (!((((((v3 / (0.0 + 3.14159)) + ((3.14159 * v3) * (0.1 / float64(((1.5 - v8) + ((v8 - v8) * v8)))))) <= 0.0) && (569 >= 931)) || (v6 && v6)) || (true && (!v6))))
                                                                    if len(v4) > 0 {
                                                                        v4[0] = "foo bar"
                                                                    }
                                                                    fmt.Println(v8, v1)
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v7, v6)
                                                            case "hello":
                                                                switch ("hello" + "foo bar") {
                                                                case "\n\t":
                                                                    v3 += 1.5
                                                                    var v34 int32 = -2147483648
                                                                    if (!v6) {
                                                                        if v6 {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (("\n\t" + "a") + "\n\t")
                                                                            }
                                                                            var v35 int = (fi32 & 42)
                                                                            _ = v35
                                                                        }
                                                                        v8 = (v8 / ((v8 - v8) - (v2 + ((0.0 * v8) * v8))))
                                                                        fmt.Println(v34, v3)
                                                                    } else {
                                                                        _ = len(v5)
                                                                        fmt.Println(fi32)
                                                                        switch (v33 * -2147483648) {
                                                                        case -1:
                                                                            v34 -= -1
                                                                        case -143:
                                                                            if false {
                                                                                switch ("\n\t" + "") {
                                                                                case "":
                                                                                    var v36 string = "a"
                                                                                    _ = v36
                                                                                case "a":
                                                                                    var v37 uint32 = 2
                                                                                    _ = v37
                                                                                default:
                                                                                    fmt.Println(v8, v33)
                                                                                }
                                                                                v38 := GenS4{F0: ("" + "")}
                                                                                _ = v38
                                                                            } else {
                                                                                v2 += v8
                                                                                var v39 int = (v33 | (fi32 | (fi30 & (v33 % ((fi32) | 1)))))
                                                                                _ = v39
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v34
                                                                default:
                                                                    v40 := GenS1{F0: "hello", F1: int16(fi30), F2: v7}
                                                                    _ = v40
                                                                }
                                                            }
                                                            _ = v33
                                                        }
                                                        if (!(100 >= 100)) {
                                                            if (v1 >= 375) {
                                                                if (777 == 1) {
                                                                    fmt.Println(v8, fi32, v1)
                                                                    _ = len(v5)
                                                                    var v41 int = 42
                                                                    _ = v41
                                                                }
                                                                if v6 {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = ("\n\t" + ((("a" + "a") + "a") + ""))
                                                                    }
                                                                } else {
                                                                    switch (fi32 - -1) {
                                                                    case 128:
                                                                        v7 = (v7 + (^(v7)))
                                                                    default:
                                                                        var v42 uint16 = 859
                                                                        _ = v42
                                                                    }
                                                                }
                                                                var v43 uint = 65535
                                                                _ = v43
                                                            }
                                                            var v44 uint8 = 1
                                                            _ = v44
                                                        }
                                                        if ((v7 & (v7 | (v7 % ((v7) | 1)))) <= 791) {
                                                            v45 := GenS4{F0: "hello"}
                                                            var v46 int8 = -85
                                                            v2 += v2
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = fi32
                                                    }
                                                    _ = fi30
                                                    _ = v31
                                                }
                                                v6 = false
                                            } else {
                                                if false {
                                                    v47 := GenS5{F0: 708, F1: 0}
                                                    var v48 float32 = 0.1
                                                    _ = v47
                                                    _ = v48
                                                }
                                                if ((v6 && (!(42 != 65535))) && v6) {
                                                    v3 -= ((float64((v8 * v2)) / -2.25) - v3)
                                                }
                                            }
                                            var v49 uint8 = 1
                                            v2 = (v2 - 1e10)
                                            _ = v49
                                        } else {
                                            for fi50 := 0; fi50 < 14; fi50++ {
                                                var v51 int32 = -908
                                                if len(v5) > 0 {
                                                    v5[0] = (fi50 << ((42) & 63))
                                                }
                                                _ = fi50
                                                _ = v51
                                            }
                                            if false {
                                                if false {
                                                    fmt.Println(v3, v1)
                                                }
                                            } else {
                                                v5 = append(v5, 1)
                                                if v6 {
                                                    v6 = v6
                                                } else {
                                                    var v52 uint16 = 65535
                                                    _ = v52
                                                }
                                                var v53 int16 = (v1 >> ((720) & 63))
                                                _ = v53
                                            }
                                        }
                                    }
                                    v7 = (v7 * v7)
                                }
                            case "a":
                                v1 = v1
                            default:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            v8 = ((v8 * (0.1 / (0.0 * v8))) / (v2 - v2))
                            for fi54 := 0; fi54 < 13; fi54++ {
                                switch 128 {
                                case 2147483647:
                                    if ((76 != -106) || true) {
                                        for fi55 := 0; fi55 < 1; fi55++ {
                                            v6 = (!(-32768 < v1))
                                            _ = len(v4)
                                            var v56 string = "foo bar"
                                            _ = fi55
                                            _ = v56
                                        }
                                        fmt.Println(v2, v3, v2)
                                        var v57 int32 = 1
                                        _ = v57
                                    } else {
                                        if v6 {
                                            v8 = float32((3.14159 / v3))
                                        }
                                        var v58 uint8 = 0
                                        fmt.Println(v58, v58, v58)
                                        _ = v58
                                    }
                                }
                                var v59 uint64 = 31
                                switch fi54 {
                                case 396:
                                    for fi60 := 0; fi60 < 5; fi60++ {
                                        v59 = 1
                                        if (int32(v1) <= -2147483648) {
                                            v61 := GenS6{F0: 671, F1: 127}
                                            _ = v61
                                        } else {
                                            var v62 float64 = ((((1e10 - float64(1.5)) - (1.5 * (v3 + (((0.1 * v3) / v3) * ((((float64(v8) * (v3 / 1e10)) + v3) / 0.1) - 0.0))))) - float64(v8)) * 3.14159)
                                            v8 = float32((3.14159 * v62))
                                            _ = v62
                                        }
                                        var v63 int = fi60
                                        _ = fi60
                                        _ = v63
                                    }
                                case -328:
                                    v2 = float32(v3)
                                default:
                                    fmt.Println(v8, v1)
                                }
                                _ = fi54
                                _ = v59
                            }
                        } else {
                            v6 = false
                            _ = len(v5)
                        }
                        _ = v24
                        _ = v25
                    } else {
                        var v64 int16 = v1
                        v65 := GenS4{F0: "a"}
                        _ = v64
                        _ = v65
                    }
                } else {
                    for fi66 := 0; fi66 < 11; fi66++ {
                        v67 := GenS5{F0: 4294967295, F1: 0}
                        switch 735 {
                        case 4294967295:
                            v6 = true
                        case 127:
                            if (!(!v6)) {
                                v1 = v1
                            } else {
                                fmt.Println(v2, v1, v2)
                                v68 := GenS6{F0: -769, F1: 1}
                                _ = v68
                            }
                        }
                        _ = fi66
                        _ = v67
                    }
                }
            case "\n\t":
                var v69 uint = uint(v7)
                _ = v69
            default:
                for fi70 := 0; fi70 < 7; fi70++ {
                    fmt.Println(fi70, fi70, v8)
                    _ = fi70
                }
            }
        }
        _ = v18
    }
    v71 := GenS7{F0: 1, F1: 238}
    var v72 uint16 = 0
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if v6 {
        v2 += float32(1.5)
        fmt.Println(v3, v2)
    } else {
        var v73 int32 = -128
        fmt.Println(v6)
        v6 = true
        _ = v73
    }
    var v74 int = -1
    var v75 int32 = 42
    fmt.Println(v1, v1, v2)
    v6 = (true && v6)
    var v76 int32 = -890
    var v77 int = int(v7)
    var v78 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v72
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
    _ = v8
}
