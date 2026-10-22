package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int32; F1 uint }
type GenS9 struct { F0 int32; F1 int32; F2 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int16; F1 string }
func genFn0(p0 uint16, p1 bool, p2 bool) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int64 = 531
        v7 := GenS1{F0: 128, F1: 42}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{2147483647}
    var v9 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p0 - 32767)
}
func genFn1(p0 int16, p1 uint32, p2 string) int64 {
    v10 := GenS1{F0: 42, F1: 20}
    var v11 int64 = -1
    fmt.Println(p0, p1)
    v12 := []int{758, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return v11
}

func main() {
    v1 := []string{"hello", "foo bar", "foo bar", ("a" + "foo bar")}
    var v2 bool = true
    var v3 int64 = -368
    v4 := make(map[int]string)
    if len(v1) > 0 {
        v1[0] = ("" + "\n\t")
    }
    var v13 int32 = -2147483648
    if (!(-128 == 44)) {
        v14 := GenS1{F0: 255, F1: -71}
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        var v15 float64 = -2.25
        _ = v14
        _ = v15
    } else {
        switch ("a" + "") {
        case "hello":
            var v16 int64 = (^(v3))
            _ = v16
        case "foo bar":
            v17 := GenS2{F0: 1, F1: 0}
            _ = v17
        case "\n\t":
            for fi18 := 0; fi18 < 6; fi18++ {
                v13 = v13
                fmt.Println(v3, v2, v2)
                _ = fi18
            }
        }
        fmt.Println(v13, v13, v2)
    }
    v3 = v3
    if true {
        var v19 uint = 218
        if v2 {
            var v20 int64 = -448
            switch 544 {
            case 128:
                if (42 > 320) {
                    fmt.Println(v13)
                    v21 := GenS3{F0: 42, F1: v2}
                    v4[-46] = ("hello" + "")
                    _ = v21
                } else {
                    delete(v4, 127)
                    for fi22 := 0; fi22 < 16; fi22++ {
                        var v23 int = (fi22 ^ fi22)
                        var v24 int16 = 1
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    v25 := GenS4{F0: 42}
                    _ = v25
                }
            case 127:
                for fi26 := 0; fi26 < 12; fi26++ {
                    var v27 uint8 = 129
                    v28 := GenS3{F0: 631, F1: (((("foo bar" + ("hello" + ("" + "foo bar"))) + "a") + "hello") > "hello")}
                    for fi29 := 0; fi29 < 3; fi29++ {
                        fmt.Println(v19, v20)
                        if (v2 && (v27 > (v27 + v27))) {
                            v2 = (v2 || ((v2 || (65535 != 2)) || v2))
                            v27 -= (v27 / ((57) | 1))
                            var v30 uint64 = 1
                            var v31 uint = (v19 / ((uint(v27)) | 1))
                            _ = v30
                            _ = v31
                        } else {
                            v13 = (v13 * (v13 ^ 705))
                            if v2 {
                                if true {
                                    var v32 int8 = 1
                                    v1 = append(v1, "hello")
                                    v27 = v27
                                    _ = v32
                                }
                                if v2 {
                                    if (v2 || (218 <= 65535)) {
                                        v13 += v13
                                        if (930 >= 2147483647) {
                                            var v33 int32 = v13
                                            _ = v33
                                        }
                                        if true {
                                            v19 += (v19 % (((^(v19))) | 1))
                                            if v2 {
                                                v19 += 783
                                            } else {
                                                v34 := GenS1{F0: (v27 | v27), F1: 1}
                                                _ = v34
                                            }
                                        }
                                        v35 := GenS5{F0: 0, F1: 0}
                                        _ = v35
                                    } else {
                                        var v36 bool = (((1e10 / ((((0.0 + (0.1 / (0.0 - float64((3.14159 - 1.5))))) - 1.5) + (1.5 / (((-2.25 - (((((1e10 + ((float64(0.0) / float64(0.1)) / 1e10)) / 1.0) - 0.0) - -2.25) / 0.1)) + -2.25) / 0.1))) * 1e10)) - 1.5) <= (3.14159 * -2.25))
                                        v3 -= genFn1(-32768, 2147483647, "a")
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v3 = (v20 ^ (v3 * (v3 - genFn1(209, 377, ("a" + "a")))))
                                        v13 = v13
                                        v37 := GenS6{F0: 128, F1: 245}
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 string = ("" + "a")
                                    fmt.Println(v3, v13)
                                    _ = v38
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    v19 -= v19
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    v2 = true
                                    if (0 != v3) {
                                        fmt.Println(v20, v19)
                                        v13 += v13
                                        switch ("" + "") {
                                        case "a":
                                            v39 := GenS7{F0: 959}
                                            _ = v39
                                        default:
                                            fmt.Println(v13, v3, v19)
                                        }
                                    } else {
                                        if v2 {
                                            var v40 float32 = 1e10
                                            switch 128 {
                                            case 0:
                                                delete(v4, 439)
                                            case 9223372036854775807:
                                                v41 := GenS2{F0: 98, F1: (fi29 >> ((696) & 63))}
                                                v1 = append(v1, "a")
                                                _ = v41
                                            case -541:
                                                var v42 int16 = 665
                                                _ = v42
                                            default:
                                                v2 = ((uint16(v27) <= 32767) && ((v20 - (v20 >> ((v19) & 63))) >= 100))
                                            }
                                            _ = v40
                                        }
                                    }
                                }
                            }
                        }
                        _ = fi29
                    }
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            default:
                v43 := GenS8{F0: (v13 & v13), F1: (^(v19))}
                _ = v43
            }
            v2 = (((v13 >> (((v19 & 2)) & 63)) >= v13) || v2)
            _ = v20
        } else {
            for fi44 := 0; fi44 < 19; fi44++ {
                v13 += v13
                for fi45 := 0; fi45 < 6; fi45++ {
                    v46 := GenS9{F0: -1, F1: (v13 ^ (v13 >> ((65535) & 63))), F2: v2}
                    _ = fi45
                    _ = v46
                }
                for fi47 := 0; fi47 < 19; fi47++ {
                    if (v2 && (true && v2)) {
                        var v48 int64 = int64(v13)
                        if v2 {
                            var v49 int64 = v3
                            var v50 int64 = v3
                            var v51 int = (^(fi44))
                            if len(v1) > 0 {
                                v1[0] = (("" + "foo bar") + "")
                            }
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        _ = v48
                    } else {
                        switch (fi47 + fi47) {
                        case 127:
                            var v52 string = ("foo bar" + "foo bar")
                            _ = v52
                        case 2147483647:
                            v53 := GenS10{F0: ("foo bar" + "hello")}
                            _ = v53
                        case 128:
                            var v54 uint = (v19 ^ v19)
                            _ = v54
                        default:
                            var v55 uint = v19
                            fmt.Println(fi44, fi44, v55)
                            fmt.Println(v19, fi47)
                            _ = v55
                        }
                        v2 = (18446744073709551615 <= 42)
                        _ = len(v4)
                        v19 -= (^(v19))
                        fmt.Println(fi47)
                    }
                    _ = fi47
                }
                _ = fi44
            }
            for fi56 := 0; fi56 < 4; fi56++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi57 := 0; fi57 < 8; fi57++ {
                    switch fi57 {
                    case 128:
                        v19 = 2
                    default:
                        var v58 float32 = 0.1
                        _ = v58
                    }
                    v59 := GenS9{F0: 42, F1: (v13 << ((v19) & 63)), F2: (2 <= (^(v19)))}
                    var v60 uint8 = 42
                    _ = fi57
                    _ = v59
                    _ = v60
                }
                _ = fi56
            }
        }
        _ = v19
    } else {
        var v61 uint16 = 255
        if (-2.25 > ((3.14159 * (float64(float32((((-2.25 / float64(-2.25)) / (3.14159 * (-2.25 / 1.5))) - 3.14159))) - float64((1.5 * 3.14159)))) * 3.14159)) {
            if v2 {
                var v62 int8 = 100
                _ = v62
            }
            var v63 int64 = -261
            var v64 uint16 = (^(v61))
            _ = v63
            _ = v64
        }
        _ = v61
    }
    if (!true) {
        var v65 string = "hello"
        v3 = v3
        for fi66 := 0; fi66 < 13; fi66++ {
            var v67 string = v65
            var v68 int16 = -32768
            _ = fi66
            _ = v67
            _ = v68
        }
        _ = v65
    }
    fmt.Println(v13, v2)
    for fi69 := 0; fi69 < 8; fi69++ {
        for fi70 := 0; fi70 < 3; fi70++ {
            _ = len(v4)
            _ = fi70
        }
        switch "hello" {
        case "hello":
            v3 = (v3 ^ genFn1(1, 632, ""))
        case "a":
            var v71 float64 = float64(1e10)
            _ = v71
        case "\n\t":
            switch (("foo bar" + "") + "hello") {
            case "":
                if false {
                    for fi72 := 0; fi72 < 12; fi72++ {
                        v2 = (4294967295 != 4294967295)
                        _ = fi72
                    }
                    for fi73 := 0; fi73 < 17; fi73++ {
                        v74 := GenS7{F0: 1}
                        _ = fi73
                        _ = v74
                    }
                } else {
                    var v75 int = 0
                    _ = v75
                }
            }
        }
        _ = fi69
    }
    var v76 uint8 = 255
    v13 += v13
    v77 := GenS3{F0: 873, F1: true}
    v3 -= 1
    var v78 int16 = 32767
    for fi79 := 0; fi79 < 5; fi79++ {
        var v80 uint32 = 2
        v80 = v80
        _ = fi79
        _ = v80
    }
    if false {
        if ("\n\t" <= ("" + "foo bar")) {
            for fi81 := 0; fi81 < 10; fi81++ {
                for fi82 := 0; fi82 < 11; fi82++ {
                    if v2 {
                        var v83 float32 = ((0.1 / 1.0) / 1.0)
                        v84 := GenS7{F0: 2147483648}
                        _ = v83
                        _ = v84
                    }
                    _ = fi82
                }
                for fi85 := 0; fi85 < 1; fi85++ {
                    v76 = 127
                    _ = fi85
                }
                v1 = append(v1, ("" + "hello"))
                _ = fi81
            }
            if v2 {
                v86 := GenS11{F0: -443, F1: ""}
                _ = v86
            }
        }
        var v87 int64 = int64(v13)
        _ = v87
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v76
    _ = v77
    _ = v78
}
