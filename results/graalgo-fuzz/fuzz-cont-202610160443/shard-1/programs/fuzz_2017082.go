package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int32; F1 string; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint; F2 int16 }
type GenS10 struct { F0 int; F1 int; F2 uint8 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint64 }
func genFn0(p0 int8, p1 int32, p2 uint32) string {
    p0 -= p0
    fmt.Println(p2, p0)
    if false {
        fmt.Println(p1, p2)
    } else {
        p0 -= 9
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{}
    v3 := []int{-9223372036854775808, 9223372036854775807, 42, 128}
    var v4 bool = true
    var v5 int64 = -9223372036854775808
    v6 := GenS1{F0: 1, F1: 65535, F2: 772}
    switch "" {
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            if v4 {
                switch ("hello" + ("a" + "a")) {
                case "a":
                    v4 = (!false)
                }
                if v4 {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        switch "" {
                        case "\n\t":
                            if false {
                                if (((((1.5 + 1.5) + float32(-2.25)) - (float32((3.14159 / float64(-2.25))) + 3.14159)) / ((-2.25 - 3.14159) / 0.1)) < (((((0.0 + 0.1) * float32(1.5)) * 1.5) - (1e10 - 3.14159)) + 3.14159)) {
                                    delete(v1, 1)
                                } else {
                                    var v9 float32 = 0.1
                                    if (v4 && (-1 == 2147483647)) {
                                        switch (^(fi8)) {
                                        case 4294967295:
                                            var v10 bool = v4
                                            _ = v10
                                        case -166:
                                            var v11 int8 = 0
                                            _ = v11
                                        }
                                        switch (fi7 >> ((562) & 63)) {
                                        case 2147483647:
                                            v12 := GenS2{F0: -115, F1: 42, F2: v4}
                                            _ = v12
                                        case 481:
                                            delete(v1, 42)
                                        case 0:
                                            v13 := GenS3{F0: 0, F1: 32767, F2: 0}
                                            _ = v13
                                        }
                                    }
                                    _ = v9
                                }
                            } else {
                                v5 += v5
                            }
                        default:
                            v4 = v4
                        }
                        _ = len(v2)
                        if (!(4294967295 > 0)) {
                            if ((v5 >> ((4294967295) & 63)) < (v5 >> ((4294967295) & 63))) {
                                if (false && (52 >= 4294967295)) {
                                    fmt.Println(v4)
                                    v5 = v5
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    v14 := GenS4{F0: v4}
                                    _ = v14
                                case "\n\t":
                                    v15 := GenS3{F0: 1, F1: 32767, F2: 25}
                                    _ = v15
                                case "hello":
                                    var v16 int32 = -746
                                    _ = v16
                                }
                                if (9223372036854775807 == 971) {
                                    v4 = (0 != -721)
                                    switch "\n\t" {
                                    case "a":
                                        switch (-(fi7)) {
                                        case 606:
                                            var v17 int16 = -1
                                            _ = v17
                                        case 42:
                                            var v18 float32 = ((((-2.25 / (0.0 - -2.25)) * 0.0) * -2.25) + 0.0)
                                            _ = v18
                                        case 9223372036854775807:
                                            var v19 uint = 255
                                            _ = v19
                                        default:
                                            var v20 int8 = 42
                                            _ = v20
                                        }
                                    case "foo bar":
                                        if (v4 && (!(42 >= 2))) {
                                            v3 = append(v3, (fi8 | 199))
                                            if v4 {
                                                var v21 uint8 = 126
                                                _ = v21
                                            }
                                            v22 := GenS2{F0: 42, F1: 1000, F2: v4}
                                            _ = v22
                                        } else {
                                            if (v4 || true) {
                                                var v23 uint32 = 65535
                                                _ = v23
                                            } else {
                                                v24 := GenS5{F0: 0}
                                                v5 -= (v5 | -579)
                                                v25 := v1[2147483647]
                                                _ = v24
                                                _ = v25
                                            }
                                            v5 = (v5 >> ((255) & 63))
                                            var v26 int = fi7
                                            _ = v26
                                        }
                                    }
                                    var v27 float64 = 1e10
                                    _ = v27
                                } else {
                                    if (!true) {
                                        v28 := GenS6{F0: ("\n\t" + "")}
                                        _ = v28
                                    }
                                    if (!(!(((fi7 ^ fi8) >= 32767) && v4))) {
                                        fmt.Println(v4, fi8)
                                    } else {
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v5 = v5
                                            case "hello":
                                                var v29 int16 = 32767
                                                _ = v29
                                            default:
                                                v5 = v5
                                            }
                                            v30 := GenS2{F0: 32767, F1: -285, F2: ("" < "\n\t")}
                                            _ = v30
                                        }
                                    }
                                }
                                fmt.Println(v4, v4, fi8)
                                delete(v1, 42)
                                switch ("a" + "foo bar") {
                                case "":
                                    v4 = ((v5 | v5) == 0)
                                case "\n\t":
                                    var v31 int16 = 0
                                    _ = v31
                                case "a":
                                    if v4 {
                                        v32 := GenS7{F0: "hello", F1: (fi8 & (fi7 / ((fi8) | 1)))}
                                        var v33 int8 = 42
                                        v5 = (v5 - (v5 >> ((4294967295) & 63)))
                                        _ = v32
                                        _ = v33
                                    } else {
                                        var v34 bool = false
                                        _ = len(v1)
                                        if v34 {
                                            if (42 < 0) {
                                                _ = len(v1)
                                                v35 := GenS8{F0: 1, F1: "hello", F2: 127}
                                                _ = v35
                                            } else {
                                                v1[42] = v4
                                                if len(v3) > 0 {
                                                    v3[0] = int(v5)
                                                }
                                            }
                                        }
                                        if (v34 && (!(1 < 18446744073709551615))) {
                                            v34 = (!((((((0.0 + float32((((float64(0.0) * (3.14159 + 1.5)) * float64(-2.25)) - (3.14159 + 0.0)))) + 1e10) / (((float32(0.1) / (1.5 * 3.14159)) + 0.1) - 0.1)) / 1e10) * float32(0.0)) > 0.0))
                                            v36 := GenS5{F0: 2}
                                            if false {
                                                var v37 int64 = -1
                                                if (1 == -61) {
                                                    var v38 string = "hello"
                                                    fmt.Println(fi8, fi8)
                                                    var v39 string = "foo bar"
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    var v40 uint16 = 255
                                                    if v4 {
                                                        v3 = append(v3, fi7)
                                                    } else {
                                                        switch fi7 {
                                                        case 57:
                                                            var v41 uint64 = 18446744073709551615
                                                            _ = v41
                                                        case 2147483647:
                                                            var v42 uint16 = v40
                                                            _ = v42
                                                        case -2147483648:
                                                            v43 := GenS9{F0: 0, F1: 353, F2: -923}
                                                            _ = v43
                                                        }
                                                    }
                                                    var v44 int16 = -338
                                                    _ = v40
                                                    _ = v44
                                                }
                                                v45 := GenS9{F0: 615, F1: 4294967295, F2: 1}
                                                var v46 float64 = float64(0.1)
                                                fmt.Println(fi8)
                                                v47 := GenS6{F0: "a"}
                                                _ = v37
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                            _ = v36
                                        }
                                        _ = v34
                                    }
                                default:
                                    v4 = v4
                                }
                                if v4 {
                                    fmt.Println(v5, fi7, v5)
                                    fmt.Println(fi8, fi7)
                                } else {
                                    var v48 float64 = 1.5
                                    _ = v48
                                }
                                fmt.Println(fi8)
                                v5 -= v5
                                v4 = false
                                v4 = true
                                v49 := GenS10{F0: fi8, F1: -385, F2: 1}
                                _ = v49
                            } else {
                                v4 = (((2 != 32767) && (-2147483648 >= 1)) || v4)
                                switch "" {
                                case "":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v50 int32 = -1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if v4 {
                                        var v51 uint32 = 35
                                        _ = v51
                                    }
                                    _ = v50
                                }
                            }
                        } else {
                            v4 = v4
                            if (0 > 2) {
                                var v52 float32 = (1e10 - 3.14159)
                                _ = v52
                            } else {
                                v53 := v1[0]
                                if (((float64(((float32(((((-2.25 - float64(((3.14159 - (1e10 - float32(1.5))) - (0.1 * -2.25)))) / -2.25) / (((1e10 - 1.5) * ((float64(-2.25) - (0.0 * 3.14159)) / 1.5)) / 1.0)) / 3.14159)) * -2.25) - (((((3.14159 / 1.0) + 0.1) - 0.1) / 1e10) + 0.0))) - 1.5) - 1.5) <= 0.0) {
                                    switch (fi8 - 32767) {
                                    case 1:
                                        var v54 int64 = v5
                                        _ = v54
                                    case 42:
                                        var v55 float32 = 3.14159
                                        _ = v55
                                    default:
                                        v53 = v53
                                    }
                                    var v56 string = "a"
                                    _ = v56
                                }
                                var v57 uint8 = 2
                                _ = v53
                                _ = v57
                            }
                        }
                        _ = fi8
                    }
                }
            } else {
                var v58 int = fi7
                for fi59 := 0; fi59 < 8; fi59++ {
                    var v60 string = "foo bar"
                    var v61 bool = true
                    _ = fi59
                    _ = v60
                    _ = v61
                }
                var v62 float32 = 0.1
                _ = v58
                _ = v62
            }
            v63 := v1[2147483648]
            fmt.Println(v5)
            _ = fi7
            _ = v63
        }
    case "a":
        if ((!(42 < -265)) || ((!(!v4)) || (2147483647 >= 1))) {
            var v64 int32 = 0
            _ = v64
        }
    case "\n\t":
        var v65 float64 = 0.1
        _ = v65
    }
    if (690 >= 268) {
        for fi66 := 0; fi66 < 2; fi66++ {
            _ = len(v1)
            _ = fi66
        }
    } else {
        if ((false || v4) || (v5 > v5)) {
            fmt.Println(v5, v4)
            var v67 uint16 = 2
            _ = v67
        } else {
            var v68 float32 = float32(float64(float32(-2.25)))
            _ = v68
        }
    }
    v69 := GenS11{F0: -2147483648}
    v4 = true
    v4 = v4
    v5 = (^(v5))
    v70 := GenS7{F0: (("" + ("" + ("" + "hello"))) + ((("foo bar" + "") + "\n\t") + "a")), F1: 32767}
    v71 := GenS12{F0: 0}
    v72 := GenS7{F0: "", F1: -1}
    delete(v1, 530)
    v4 = (-2.25 != 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v70
    _ = v71
    _ = v72
}
