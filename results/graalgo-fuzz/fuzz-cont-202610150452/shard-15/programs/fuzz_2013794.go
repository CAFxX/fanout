package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int8 }
type GenS2 struct { F0 uint; F1 string; F2 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint }
func genFn0(p0 int64) uint {
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            p0 -= 0
            v9 := GenS1{F0: 65535, F1: -1, F2: 127}
            if (0 >= 711) {
                if (42 < 2) {
                    v10 := []int{fi8}
                    p0 -= 1
                    _ = v10
                }
                v11 := []int{(fi8 | 442), fi7}
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                if true {
                    fmt.Println(fi8, p0)
                    var v12 uint64 = 0
                    switch ("" + (("hello" + ("" + "\n\t")) + "foo bar")) {
                    case "hello":
                        p0 = int64(v12)
                    }
                    p0 -= -343
                    _ = v12
                }
                var v13 int8 = 42
                if ((("" + "hello") + "a") >= "foo bar") {
                    v14 := GenS1{F0: 2147483648, F1: v13, F2: v13}
                    v15 := GenS1{F0: 0, F1: (v13 >> ((255) & 63)), F2: v13}
                    v13 = v13
                    _ = v14
                    _ = v15
                }
                _ = v11
                _ = v13
            } else {
                v16 := []int{}
                _ = v16
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v17 int8 = 42
    switch ("hello" + ("\n\t" + "foo bar")) {
    case "foo bar":
        if true {
            for fi18 := 0; fi18 < 15; fi18++ {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v20 := []int{}
                    p0 -= -1
                    v20 = append(v20, (fi19 | (fi18 * (-(fi19)))))
                    _ = fi19
                    _ = v20
                }
                if (!((2 < 269) || ((1e10 / float32(1e10)) != 1.5))) {
                    v17 = (v17 | 100)
                    var v21 float32 = 0.0
                    for fi22 := 0; fi22 < 9; fi22++ {
                        var v23 int64 = p0
                        v24 := []int{fi18, fi18, 0}
                        var v25 int = 4294967295
                        _ = fi22
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    _ = v21
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        var v26 uint8 = 42
                        _ = v26
                    default:
                        fmt.Println(fi18)
                    }
                    if (((!false) || true) || (3.14159 == (0.0 - float32(1e10)))) {
                        fmt.Println(v17)
                    } else {
                        v27 := GenS1{F0: 2147483647, F1: v17, F2: (-(v17))}
                        v28 := []int{}
                        v29 := GenS1{F0: 42, F1: (v17 % (((v17 - (v17 + v17))) | 1)), F2: (v17 ^ v17)}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint = 0
                    _ = v30
                }
                _ = fi18
            }
        }
    case "":
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 int = 335
            switch (fi31 | (v32 / ((fi31) | 1))) {
            case 475:
                v32 -= v32
            case 9223372036854775807:
                var v33 uint8 = 2
                _ = v33
            case 1:
                v34 := GenS2{F0: 255, F1: (("" + "foo bar") + "foo bar"), F2: 42}
                _ = v34
            }
            v35 := GenS3{F0: (42 < 0)}
            _ = fi31
            _ = v32
            _ = v35
        }
    }
    var v36 float32 = (1.5 * 1e10)
    _ = p0
    _ = v17
    _ = v36
    return 638
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 766
    var v3 int64 = 1
    v4 := []string{"hello", (("foo bar" + "") + "hello")}
    v5 := []string{("" + "a"), (("hello" + ("a" + ("\n\t" + "foo bar"))) + ((((("" + "hello") + (("" + "hello") + "a")) + (("" + "\n\t") + (("hello" + "") + ("hello" + "a")))) + "") + ("foo bar" + "hello")))}
    v6 := make(map[uint32]int)
    var v37 uint8 = 1
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v3, v2)
    v3 = int64(v37)
    if true {
        switch "hello" {
        case "":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "\n\t":
            _ = len(v5)
        case "hello":
            if false {
                fmt.Println(v3)
            }
        default:
            for fi38 := 0; fi38 < 14; fi38++ {
                switch ("\n\t" + "a") {
                case "a":
                    fmt.Println(fi38, v1, v37)
                case "foo bar":
                    v6[717] = 42
                default:
                    switch "a" {
                    case "foo bar":
                        switch "" {
                        case "foo bar":
                            for fi39 := 0; fi39 < 4; fi39++ {
                                v40 := GenS1{F0: v2, F1: 0, F2: 100}
                                var v41 uint16 = 65535
                                v42 := v6[v2]
                                v1 -= v1
                                _ = fi39
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                        case "a":
                            v5 = append(v5, "foo bar")
                        case "hello":
                            fmt.Println(v37, v3, fi38)
                        default:
                            v2 = (v2 | (v2 ^ v2))
                        }
                    case "a":
                        switch "" {
                        case "foo bar":
                            v1 = (v1 - 1.5)
                        default:
                            v43 := GenS4{F0: fi38}
                            _ = v43
                        }
                    case "hello":
                        delete(v6, v2)
                    }
                }
                var v44 string = "a"
                switch fi38 {
                case 316:
                    if true {
                        var v45 float64 = (((1e10 + 3.14159) / (v1 * v1)) + 0.0)
                        _ = v45
                    } else {
                        v44 = v44
                    }
                case 9223372036854775807:
                    if false {
                        if (255 == uint16(v2)) {
                            if (!false) {
                                fmt.Println(v1, v1)
                                v2 = (v2 << ((2) & 63))
                                for fi46 := 0; fi46 < 4; fi46++ {
                                    fmt.Println(v1, fi38)
                                    switch (v44 + v44) {
                                    case "hello":
                                        v3 -= v3
                                    default:
                                        if false {
                                            switch "\n\t" {
                                            case "\n\t":
                                                fmt.Println(fi46, v2)
                                            case "foo bar":
                                                v47 := GenS5{F0: 1, F1: (v37 / ((51) | 1))}
                                                _ = v47
                                            default:
                                                var v48 int8 = -128
                                                _ = v48
                                            }
                                        } else {
                                            fmt.Println(v44, v1)
                                        }
                                    }
                                    _ = fi46
                                }
                            }
                            if true {
                                for fi49 := 0; fi49 < 13; fi49++ {
                                    var v50 int32 = 0
                                    var v51 uint32 = v2
                                    if true {
                                        var v52 uint = 933
                                        _ = v52
                                    }
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                            }
                            var v53 uint32 = v2
                            _ = v53
                        } else {
                            v2 = 1
                        }
                    } else {
                        _ = len(v5)
                        v54 := GenS5{F0: 42, F1: 128}
                        _ = v54
                    }
                default:
                    fmt.Println(v1)
                }
                _ = fi38
                _ = v44
            }
        }
        if len(v4) > 0 {
            v4[0] = ("hello" + "")
        }
    } else {
        if true {
            v55 := GenS6{F0: 42}
            if (175 <= 1) {
                if false {
                    var v56 uint32 = v2
                    var v57 uint32 = (v56 << ((255) & 63))
                    _ = v56
                    _ = v57
                } else {
                    if true {
                        var v58 float64 = (v1 + float64(float32(v1)))
                        _ = v58
                    }
                    if true {
                        if ((42 <= -1) && ((!(!true)) || false)) {
                            switch 910 {
                            case 42:
                                var v59 string = "\n\t"
                                _ = v59
                            case 15:
                                var v60 bool = ((3.14159 - 3.14159) <= v1)
                                _ = v60
                            case 127:
                                fmt.Println(v3, v1)
                            default:
                                switch ((("hello" + "hello") + "\n\t") + "") {
                                case "foo bar":
                                    if ((32767 == -577) && true) {
                                        fmt.Println(v1, v1)
                                        fmt.Println(v1, v3)
                                        var v61 float64 = 1e10
                                        _ = v61
                                    } else {
                                        v6[uint32(v37)] = 32767
                                    }
                                default:
                                    switch ("a" + "") {
                                    case "a":
                                        if true {
                                            if false {
                                                var v62 bool = false
                                                var v63 uint = genFn0((v3 % ((v3) | 1)))
                                                if v62 {
                                                    var v64 bool = true
                                                    _ = v64
                                                } else {
                                                    v37 = (v37 * v37)
                                                    if false {
                                                        v65 := GenS1{F0: (v2 / (((v2 + (^(v2)))) | 1)), F1: 12, F2: 100}
                                                        for fi66 := 0; fi66 < 17; fi66++ {
                                                            fmt.Println(v3, v63, v37)
                                                            if (0 < 894) {
                                                                var v67 uint16 = 65535
                                                                v68 := GenS1{F0: (v2 - 2), F1: -1, F2: 0}
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS1{F0: (v2 | 1), F1: 100, F2: 127}
                                                                _ = v69
                                                            }
                                                            fmt.Println(v3)
                                                            _ = fi66
                                                        }
                                                        v5 = append(v5, "foo bar")
                                                        _ = v65
                                                    } else {
                                                        v62 = ("a" == "")
                                                        var v70 bool = ((!((^(v2)) == 141)) && false)
                                                        fmt.Println(v70, v37)
                                                        _ = v70
                                                    }
                                                    v37 -= v37
                                                }
                                                _ = v62
                                                _ = v63
                                            }
                                            if false {
                                                var v71 float32 = 3.14159
                                                _ = v71
                                            }
                                        } else {
                                            v72 := GenS3{F0: (!false)}
                                            v1 += v1
                                            _ = v72
                                        }
                                    case "\n\t":
                                        switch "hello" {
                                        case "hello":
                                            v73 := GenS3{F0: false}
                                            _ = v73
                                        case "":
                                            switch "a" {
                                            case "a":
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case "":
                                                if len(v5) > 0 {
                                                    v5[0] = "foo bar"
                                                }
                                            }
                                        case "foo bar":
                                            var v74 uint32 = 777
                                            _ = v74
                                        default:
                                            v5 = append(v5, "\n\t")
                                        }
                                    }
                                }
                            }
                            v75 := GenS7{F0: (false && false)}
                            if (-1 < 32767) {
                                v5 = append(v5, ("foo bar" + "hello"))
                            } else {
                                var v76 uint32 = v2
                                v2 = v76
                                v37 -= 56
                                _ = v76
                            }
                            _ = v75
                        }
                        if false {
                            v1 = (((0.0 + 1.5) + ((v1 / v1) + v1)) * v1)
                        }
                    }
                }
                switch -9223372036854775808 {
                case 994:
                    _ = len(v6)
                case -9223372036854775808:
                    for fi77 := 0; fi77 < 19; fi77++ {
                        v78 := GenS6{F0: 305}
                        v79 := GenS8{F0: v2, F1: 42}
                        _ = len(v6)
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                case 0:
                    var v80 uint32 = (v2 << ((2) & 63))
                    _ = v80
                default:
                    fmt.Println(v1, v2)
                }
                v81 := GenS6{F0: 2}
                _ = v81
            }
            var v82 int = 4294967295
            _ = v55
            _ = v82
        }
        v2 -= v2
    }
    v83 := GenS6{F0: 18446744073709551615}
    v84 := GenS7{F0: false}
    for fi85 := 0; fi85 < 17; fi85++ {
        fmt.Println(v3)
        if true {
            v37 = v37
        } else {
            if len(v5) > 0 {
                v5[0] = ("\n\t" + "\n\t")
            }
            fmt.Println(v3, v37)
        }
        _ = fi85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
}
