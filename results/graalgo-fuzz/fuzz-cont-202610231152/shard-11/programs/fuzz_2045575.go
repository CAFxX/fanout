package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 int64 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 uint }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 int; F1 int; F2 uint32 }
func genFn0() int32 {
    var v5 int32 = -2147483648
    switch ("" + "hello") {
    case "a":
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 int32 = 0
                _ = fi7
                _ = v8
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint8 = 128
                switch "foo bar" {
                case "":
                    v5 = v5
                    var v11 uint16 = 33
                    fmt.Println(v5, fi6, v10)
                    _ = v11
                }
                v12 := GenS1{F0: 412, F1: 0}
                _ = fi9
                _ = v10
                _ = v12
            }
            v13 := []int{fi6}
            _ = fi6
            _ = v13
        }
    default:
        if (!((100 <= 1) && (2 < 82))) {
            for fi14 := 0; fi14 < 10; fi14++ {
                v15 := GenS2{F0: -574, F1: (-(fi14))}
                var v16 uint16 = 159
                v16 += v16
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int16 = 546
            _ = v17
        } else {
            if (false || false) {
                var v18 float64 = 0.1
                v18 += v18
                _ = v18
            } else {
                for fi19 := 0; fi19 < 4; fi19++ {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        if ((65535 <= 2) && (4294967295 <= 981)) {
                            switch (fi20 / (((fi19 % (((fi20 * (fi19 + (fi19 | -761)))) | 1))) | 1)) {
                            case 127:
                                v5 = v5
                            }
                        }
                        v21 := GenS3{F0: -20}
                        _ = fi20
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 3; fi22++ {
                        var v23 int64 = 0
                        v23 -= v23
                        _ = fi22
                        _ = v23
                    }
                    v5 = v5
                    _ = fi19
                }
                var v24 int8 = -61
                _ = v24
            }
            if true {
                for fi25 := 0; fi25 < 20; fi25++ {
                    var v26 float32 = 0.0
                    v27 := GenS4{F0: 65535, F1: -197, F2: -9223372036854775808}
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v28 := GenS1{F0: 1, F1: 32767}
                _ = v28
            } else {
                var v29 int16 = 1
                var v30 uint = 65535
                var v31 uint16 = 712
                _ = v29
                _ = v30
                _ = v31
            }
        }
    }
    _ = v5
    return (v5 | 781)
}
func genFn1(p0 int32, p1 uint16, p2 int) string {
    v32 := []int{p2, p2, p2, 41}
    if (!((!true) && false)) {
        if (1e10 < 1.5) {
            fmt.Println(p0, p0, p1)
        } else {
            fmt.Println(p2, p1)
            var v33 uint8 = 128
            for fi34 := 0; fi34 < 4; fi34++ {
                var v35 uint8 = v33
                var v36 float64 = (((3.14159 + 0.1) * 1e10) * 1e10)
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v33
        }
    } else {
        p0 += p0
    }
    p1 += p1
    var v37 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v37
    return ("a" + "foo bar")
}

func main() {
    var v1 float64 = -2.25
    v2 := []int{}
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    for fi38 := 0; fi38 < 3; fi38++ {
        if (-9223372036854775808 <= 127) {
            fmt.Println(fi38, fi38)
            if (757 != 18446744073709551615) {
                v1 += 3.14159
                if (-32768 <= -919) {
                    v1 = v1
                    for fi39 := 0; fi39 < 6; fi39++ {
                        var v40 int64 = 887
                        switch (fi39 * 32768) {
                        case 32768:
                            v3[fi39] = false
                        case 128:
                            switch "foo bar" {
                            case "foo bar":
                                v1 = v1
                            case "a":
                                v40 += (v40 - v40)
                            case "hello":
                                v41 := GenS5{F0: 0, F1: 51}
                                _ = v41
                            default:
                                fmt.Println(v40, fi39, v1)
                            }
                        case 902:
                            delete(v3, (fi39 * fi38))
                        default:
                            fmt.Println(fi39)
                        }
                        _ = fi39
                        _ = v40
                    }
                    var v42 uint64 = 0
                    _ = v42
                }
                v1 = v1
            } else {
                v43 := GenS5{F0: 978, F1: 0}
                var v44 float64 = v1
                var v45 int32 = genFn0()
                _ = v43
                _ = v44
                _ = v45
            }
            var v46 string = ("hello" + ("hello" + "hello"))
            _ = v46
        } else {
            var v47 int = fi38
            for fi48 := 0; fi48 < 15; fi48++ {
                var v49 uint = 42
                v47 += (^(v47))
                _ = fi48
                _ = v49
            }
            _ = v47
        }
        _ = fi38
    }
    if (((false || true) && false) && false) {
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            delete(v4, "")
        case "\n\t":
            v1 = (3.14159 * (0.1 - (v1 / (1.5 * float64(1e10)))))
        default:
            switch 9223372036854775807 {
            case 2147483647:
                v50 := GenS3{F0: 42}
                _ = v50
            case 9223372036854775807:
                if true {
                    _ = len(v2)
                }
            }
        }
    } else {
        if (1 <= 100) {
            fmt.Println(v1)
            var v51 string = ""
            _ = v51
        } else {
            v1 = v1
            if (!(false || ("hello" > "a"))) {
                fmt.Println(v1)
                var v52 string = "a"
                var v53 float32 = 3.14159
                _ = v52
                _ = v53
            } else {
                var v54 uint64 = 4294967295
                _ = v54
            }
            var v55 int16 = -32768
            _ = v55
        }
        if (4294967295 > 4294967295) {
            v1 += 1e10
        } else {
            var v56 float64 = v1
            if (69 >= -700) {
                for fi57 := 0; fi57 < 5; fi57++ {
                    var v58 int16 = 41
                    if false {
                        if (18446744073709551615 >= 42) {
                            if false {
                                var v59 int64 = -1
                                v60 := GenS5{F0: 2147483647, F1: 232}
                                _ = v59
                                _ = v60
                            }
                            for fi61 := 0; fi61 < 20; fi61++ {
                                fmt.Println(v56)
                                switch ("hello" + "\n\t") {
                                case "\n\t":
                                    var v62 string = ("foo bar" + "foo bar")
                                    _ = v62
                                case "a":
                                    if ("foo bar" != ("a" + "\n\t")) {
                                        fmt.Println(fi61)
                                    }
                                case "":
                                    var v63 bool = false
                                    _ = v63
                                }
                                _ = fi61
                            }
                            v64 := GenS5{F0: genFn0(), F1: 127}
                            _ = v64
                        }
                        v65 := GenS6{F0: "a", F1: "\n\t"}
                        fmt.Println(v56)
                        _ = v65
                    }
                    v66 := GenS7{F0: 42}
                    _ = fi57
                    _ = v58
                    _ = v66
                }
                if true {
                    v1 += v56
                    var v67 int = 2147483647
                    _ = v67
                }
                for fi68 := 0; fi68 < 14; fi68++ {
                    fmt.Println(v1, v1, v56)
                    v1 += float64(-2.25)
                    v69 := GenS8{F0: 536, F1: 712}
                    _ = fi68
                    _ = v69
                }
            }
            _ = v56
        }
        v1 -= v1
    }
    v1 = v1
    for fi70 := 0; fi70 < 6; fi70++ {
        v71 := GenS5{F0: 0, F1: 93}
        v72 := GenS9{F0: 177, F1: 255}
        var v73 float32 = 0.0
        _ = fi70
        _ = v71
        _ = v72
        _ = v73
    }
    switch "foo bar" {
    case "hello":
        v1 = v1
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v74 uint = 971
    if (!(255 == 159)) {
        for fi75 := 0; fi75 < 2; fi75++ {
            v74 += v74
            for fi76 := 0; fi76 < 17; fi76++ {
                if false {
                    v1 = v1
                    v77 := GenS10{F0: 684, F1: (fi76 | fi75), F2: 89}
                    fmt.Println(fi75, fi76, fi76)
                    _ = v77
                } else {
                    fmt.Println(v1, fi76)
                }
                _ = len(v2)
                _ = fi76
            }
            _ = fi75
        }
    }
    if len(v2) > 0 {
        v2[0] = 733
    }
    if (-70 == 0) {
        delete(v3, 0)
    }
    if ((209 > 1) && (1e10 == 3.14159)) {
        var v78 uint = 176
        if false {
            var v79 uint8 = 7
            v1 = v1
            _ = v79
        }
        _ = v78
    } else {
        fmt.Println(v1, v74)
        var v80 int16 = -623
        _ = v80
    }
    var v81 float32 = 1.5
    v1 = v1
    delete(v3, 42)
    for fi82 := 0; fi82 < 16; fi82++ {
        fmt.Println(v1, fi82)
        _ = fi82
    }
    if false {
        var v83 int16 = 1
        v74 = v74
        v84 := v3[128]
        _ = v83
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v74:", v74)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v74
    _ = v81
}
