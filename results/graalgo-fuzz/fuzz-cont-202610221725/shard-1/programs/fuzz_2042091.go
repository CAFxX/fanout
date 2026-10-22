package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 bool; F1 int16 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 string }
type GenS8 struct { F0 bool; F1 string; F2 int32 }
type GenS9 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint8; F1 int64; F2 int32 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint8) uint16 {
    p1 = 229
    v4 := GenS1{F0: true, F1: -1, F2: 255}
    _ = p0
    _ = p1
    _ = v4
    return p0
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (0.0 * 0.0)
    v3 := []int{1, -348}
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 int16 = 1000
            var v7 uint64 = 0
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int16 = 746
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
        _ = v8
    } else {
        v9 := GenS2{F0: 32767, F1: -1, F2: -68}
        var v10 uint8 = 255
        _ = v10
        _ = v9
    }
    fmt.Println(v2)
    var v11 bool = (!false)
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 = v11
        if v11 {
            for fi13 := 0; fi13 < 2; fi13++ {
                if v11 {
                    v1["foo bar"] = "a"
                } else {
                    var v14 bool = (65535 >= 1)
                    if v14 {
                        v11 = (3.14159 != 1.5)
                    }
                    _ = v14
                }
                var v15 uint32 = 4294967295
                v1[""] = "hello"
                _ = fi13
                _ = v15
            }
            v16 := v1[("\n\t" + "foo bar")]
            var v17 float64 = (0.0 - (0.1 + 1e10))
            _ = v16
            _ = v17
        }
        _ = fi12
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v2 = v2
        v1[("foo bar" + "")] = ("foo bar" + "foo bar")
        v19 := GenS2{F0: 130, F1: -1, F2: 42}
        _ = fi18
        _ = v19
    }
    if (v11 || (662 != 18446744073709551615)) {
        var v20 int16 = 42
        var v21 float32 = 0.0
        _ = v20
        _ = v21
    } else {
        if v11 {
            switch "a" {
            case "hello":
                for fi22 := 0; fi22 < 12; fi22++ {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        if v11 {
                            if (true || (!(-1 == -252))) {
                                v24 := GenS3{F0: 65535, F1: 153}
                                fmt.Println(fi22, fi22)
                                _ = v24
                            } else {
                                if v11 {
                                    var v25 uint = 255
                                    var v26 int16 = 42
                                    _ = v25
                                    _ = v26
                                } else {
                                    if v11 {
                                        v27 := GenS1{F0: (true && (!(!((false || (!v11)) || false)))), F1: -9223372036854775808, F2: 4294967295}
                                        var v28 uint = 4294967295
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v11 = v11
                                        var v29 int = (fi22 % (((fi23 % (((fi22 & 9223372036854775807)) | 1))) | 1))
                                        _ = v29
                                    }
                                    if (true || v11) {
                                        v2 = ((v2 / (v2 + 0.1)) + 0.1)
                                        v2 = (0.1 * v2)
                                    } else {
                                        fmt.Println(fi22)
                                        var v30 uint8 = 130
                                        _ = len(v3)
                                        var v31 float32 = 0.0
                                        delete(v1, "a")
                                        _ = v30
                                        _ = v31
                                    }
                                }
                                if v11 {
                                    fmt.Println(v2, fi22)
                                    v11 = (255 < 65535)
                                    var v32 int32 = 2147483647
                                    _ = v32
                                } else {
                                    if v11 {
                                        fmt.Println(v11)
                                        var v33 uint64 = 1
                                        _ = v33
                                    } else {
                                        switch fi22 {
                                        case -37:
                                            v11 = (39 <= 1)
                                        case 735:
                                            v2 -= (float64(float32((float64(-2.25) * (1.5 - (1e10 * 3.14159))))) * (((((v2 * v2) * v2) * (float64(0.0) / 1.5)) / 0.1) + 1e10))
                                        case -2147483648:
                                            var v34 uint8 = 1
                                            v11 = true
                                            v2 += float64((-2.25 * (float32((float64((0.1 * (0.0 + (0.1 + -2.25)))) + float64(float32(((float64(0.1) * (0.1 - ((v2 - (-2.25 * float64((0.1 - 0.0)))) + v2))) / (v2 - (v2 * (v2 + v2)))))))) / -2.25)))
                                            var v35 float64 = v2
                                            v3 = append(v3, 128)
                                            _ = v34
                                            _ = v35
                                        default:
                                            fmt.Println(fi22, fi22)
                                        }
                                    }
                                    v2 = v2
                                    fmt.Println(fi22, fi22)
                                }
                            }
                            v2 -= float64(-2.25)
                            v11 = (2147483647 == 97)
                        } else {
                            v36 := GenS4{F0: (!v11), F1: 32767}
                            if v11 {
                                v37 := GenS5{F0: -32768, F1: 1}
                                _ = v37
                            } else {
                                switch -2147483648 {
                                case 4294967295:
                                    if v11 {
                                        v38 := GenS6{F0: 1}
                                        _ = v38
                                    }
                                case 959:
                                    v11 = false
                                default:
                                    var v39 string = ("foo bar" + ((("foo bar" + "") + "\n\t") + ""))
                                    _ = v39
                                }
                                v40 := GenS7{F0: -258, F1: "a"}
                                switch fi23 {
                                case 0:
                                    fmt.Println(v11, fi22, v11)
                                case -1:
                                    v2 += v2
                                case 128:
                                    _ = len(v3)
                                }
                                _ = v40
                            }
                            _ = v36
                        }
                        if v11 {
                            v11 = (-1 <= 737)
                        } else {
                            switch fi22 {
                            case 32767:
                                v2 -= float64(1e10)
                            case -53:
                                var v41 uint64 = 4294967295
                                _ = v41
                            case 127:
                                var v42 float64 = (v2 * v2)
                                _ = v42
                            default:
                                v11 = v11
                            }
                        }
                        _ = fi23
                    }
                    _ = fi22
                }
            }
        } else {
            var v43 int = 9223372036854775807
            delete(v1, "foo bar")
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 12; fi44++ {
        switch "a" {
        case "":
            var v45 bool = v11
            _ = v45
        default:
            var v46 uint8 = 0
            _ = v46
        }
        _ = fi44
    }
    v2 -= -2.25
    if (0 > 2147483647) {
        v47 := v1["a"]
        for fi48 := 0; fi48 < 16; fi48++ {
            fmt.Println(v47, v47)
            v49 := GenS8{F0: v11, F1: v47, F2: 1}
            v11 = v11
            _ = fi48
            _ = v49
        }
        _ = v47
    } else {
        switch "\n\t" {
        case "\n\t":
            v2 = (v2 * v2)
        case "hello":
            for fi50 := 0; fi50 < 17; fi50++ {
                for fi51 := 0; fi51 < 20; fi51++ {
                    if v11 {
                        if len(v3) > 0 {
                            v3[0] = (fi50 | fi51)
                        }
                        v52 := GenS6{F0: 67}
                        var v53 uint64 = 896
                        _ = v52
                        _ = v53
                    } else {
                        if ((!true) && (0.1 < 0.0)) {
                            v2 = v2
                            v2 += (((v2 * v2) - -2.25) / (v2 + float64(1e10)))
                        }
                    }
                    _ = fi51
                }
                v11 = (42 == 149)
                _ = len(v1)
                _ = fi50
            }
        case "":
            for fi54 := 0; fi54 < 6; fi54++ {
                for fi55 := 0; fi55 < 19; fi55++ {
                    switch -9223372036854775808 {
                    case -607:
                        var v56 uint = 1
                        _ = v56
                    }
                    var v57 bool = (!((42 > -816) && false))
                    if (1000 < -32768) {
                        switch (fi55 / ((fi54) | 1)) {
                        case 32768:
                            _ = len(v3)
                        case 128:
                            fmt.Println(fi55, v11)
                        case 374:
                            fmt.Println(v57, fi55, v57)
                        default:
                            v2 = (3.14159 / v2)
                        }
                        var v58 string = "a"
                        v1["foo bar"] = "\n\t"
                        _ = v58
                    } else {
                        if (!false) {
                            switch (fi54 << ((899) & 63)) {
                            case 32767:
                                fmt.Println(v57, v2)
                            case 0:
                                v59 := v1["hello"]
                                _ = v59
                            case -9223372036854775808:
                                var v60 bool = v57
                                _ = v60
                            }
                            var v61 int8 = -40
                            if v57 {
                                if (((v61 / ((-1) | 1)) > (v61 << ((461) & 63))) && true) {
                                    v11 = true
                                    var v62 int8 = (v61 & -121)
                                    switch 823 {
                                    case -1:
                                        fmt.Println(fi55, v2, v57)
                                    }
                                    _ = v62
                                }
                                v63 := GenS9{F0: 854, F1: 83, F2: 42}
                                if false {
                                    v1[""] = "foo bar"
                                    v61 += (-(v61))
                                    if v57 {
                                        v2 += 0.0
                                    }
                                } else {
                                    v64 := GenS10{F0: 12, F1: 0, F2: -308}
                                    _ = v64
                                }
                                _ = v63
                            } else {
                                v65 := GenS10{F0: 219, F1: 9223372036854775807, F2: 475}
                                if (!(1 != 255)) {
                                    if false {
                                        switch "hello" {
                                        case "a":
                                            v66 := GenS4{F0: (true || v11), F1: 352}
                                            _ = v66
                                        case "hello":
                                            v67 := GenS8{F0: (65535 != 42), F1: "\n\t", F2: 2147483647}
                                            _ = v67
                                        }
                                        switch 562 {
                                        case 127:
                                            v57 = v57
                                        case -440:
                                            v11 = ((fi54 / ((fi55) | 1)) > (fi55 ^ (-(fi55))))
                                        }
                                        v68 := GenS7{F0: 452, F1: "hello"}
                                        _ = v68
                                    } else {
                                        if false {
                                            switch fi55 {
                                            case -2147483648:
                                                v69 := GenS11{F0: (!(false || (0 < 4294967295)))}
                                                _ = v69
                                            case 127:
                                                if (42 < 2147483647) {
                                                    if (v2 == v2) {
                                                        if v57 {
                                                            v70 := GenS12{F0: 2}
                                                            var v71 int32 = -2147483648
                                                            var v72 bool = (607 <= 4294967295)
                                                            _ = v70
                                                            _ = v71
                                                            _ = v72
                                                        }
                                                    } else {
                                                        var v73 uint8 = 255
                                                        v11 = v11
                                                        v74 := GenS11{F0: (42 >= 1)}
                                                        _ = v73
                                                        _ = v74
                                                    }
                                                    fmt.Println(fi55, v61)
                                                    switch (fi55 | fi54) {
                                                    case 860:
                                                        fmt.Println(v2, v57, v11)
                                                    default:
                                                        switch ("" + "\n\t") {
                                                        case "\n\t":
                                                            fmt.Println(v11, fi55)
                                                        case "foo bar":
                                                            v61 -= (^(v61))
                                                        case "":
                                                            if (v11 || (v57 || v11)) {
                                                                v2 += 0.0
                                                                v2 = (0.1 / (v2 - 0.1))
                                                                fmt.Println(v2, v57, v61)
                                                            }
                                                        default:
                                                            v75 := GenS7{F0: 0, F1: "hello"}
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                            default:
                                                v76 := GenS6{F0: (v61 % ((91) | 1))}
                                                _ = v76
                                            }
                                            v2 -= 0.0
                                        } else {
                                            fmt.Println(v2)
                                            v57 = v57
                                            fmt.Println(v11, v61, fi55)
                                        }
                                        if (!true) {
                                            var v77 int32 = -2147483648
                                            fmt.Println(v77)
                                            switch ("a" + "hello") {
                                            case "a":
                                                v11 = (v61 <= (v61 - (-(v61))))
                                            }
                                            _ = v77
                                        }
                                    }
                                }
                                _ = v65
                            }
                            _ = v61
                        }
                    }
                    switch 294 {
                    case 0:
                        v11 = v11
                    case 32768:
                        v78 := GenS4{F0: (-714 != 981), F1: 622}
                        _ = v78
                    default:
                        v57 = (3.14159 <= v2)
                    }
                    _ = fi55
                    _ = v57
                }
                _ = fi54
            }
        }
    }
    if (65535 >= 2147483648) {
        delete(v1, "a")
        _ = len(v1)
    } else {
        if ((3.14159 - 1.5) == 1e10) {
            _ = len(v1)
            v11 = ((("foo bar" + "hello") + (("\n\t" + "hello") + "hello")) <= "a")
            fmt.Println(v2, v2)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
