package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 int64; F2 int64 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 bool; F1 bool; F2 int8 }
type GenS11 struct { F0 int32; F1 uint8; F2 bool }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string }
func genFn0(p0 int32, p1 uint16) uint32 {
    switch 1 {
    case 457:
        v6 := []int{0, 0, -2147483648}
        _ = v6
    default:
        var v7 uint = 1
        _ = v7
    }
    v8 := GenS1{F0: -70, F1: 32767}
    var v9 uint8 = 127
    v10 := GenS2{F0: 1}
    v11 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 2147483647
}
func genFn1(p0 uint8, p1 int32) uint8 {
    fmt.Println(p1, p0)
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 string = ("foo bar" + "hello")
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    return (p0 * 1)
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{}
    v3 := []uint32{}
    var v4 int64 = -951
    var v5 int = -1
    switch "\n\t" {
    case "":
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v3 = append(v3, genFn0(532, 65535))
                switch "a" {
                case "hello":
                    switch fi14 {
                    case 1:
                        _ = len(v1)
                    case -821:
                        v4 = (v4 % ((v4) | 1))
                    }
                case "foo bar":
                    if (18446744073709551615 < 0) {
                        var v16 int8 = 5
                        switch (v5 + -670) {
                        case 191:
                            var v17 string = ((("foo bar" + "") + "foo bar") + (("" + ("" + "foo bar")) + ("foo bar" + ("foo bar" + (("foo bar" + "a") + "a")))))
                            _ = v17
                        }
                        _ = v16
                    }
                }
                if false {
                    var v18 float32 = 3.14159
                    _ = v18
                } else {
                    v5 = (fi15 | (fi14 % ((295) | 1)))
                    v5 = fi15
                    v3 = append(v3, 2147483648)
                }
                _ = fi15
            }
            for fi19 := 0; fi19 < 19; fi19++ {
                if false {
                    v3 = append(v3, genFn0(0, 255))
                    if true {
                        if (genFn0(-478, 838) <= 0) {
                            fmt.Println(v5, v4)
                            var v20 uint = 0
                            if true {
                                var v21 uint64 = 1
                                v4 = (v4 >> ((610) & 63))
                                _ = v21
                            } else {
                                var v22 uint32 = 4294967295
                                fmt.Println(fi19, v5)
                                _ = v22
                            }
                            _ = v20
                        }
                        var v23 float64 = float64(3.14159)
                        _ = v23
                    }
                }
                _ = fi19
            }
            v4 = 42
            _ = fi14
        }
    case "a":
        switch "foo bar" {
        case "\n\t":
            for fi24 := 0; fi24 < 2; fi24++ {
                v5 = (fi24 ^ (fi24 << ((4294967295) & 63)))
                _ = fi24
            }
        case "a":
            for fi25 := 0; fi25 < 13; fi25++ {
                if (!((true || true) && (-2.25 > (3.14159 * (float64(3.14159) / 1.5))))) {
                    fmt.Println(fi25, fi25)
                    delete(v1, "hello")
                }
                _ = fi25
            }
        case "":
            switch "" {
            case "a":
                var v26 uint16 = 0
                _ = v26
            case "foo bar":
                for fi27 := 0; fi27 < 13; fi27++ {
                    fmt.Println(v5, v5)
                    v28 := GenS1{F0: 42, F1: 513}
                    v5 += fi27
                    _ = fi27
                    _ = v28
                }
            default:
                for fi29 := 0; fi29 < 17; fi29++ {
                    var v30 int = v5
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi29
                    _ = v30
                }
            }
        default:
            var v31 int = -43
            _ = v31
        }
    case "hello":
        v32 := GenS3{F0: 18446744073709551615, F1: 1}
        _ = v32
    }
    var v33 float64 = -2.25
    var v34 bool = (true || false)
    v35 := GenS2{F0: (v5 * (v5 - (v5 << ((1) & 63))))}
    var v36 int64 = v4
    var v37 uint16 = 281
    v33 = -2.25
    v38 := GenS2{F0: v5}
    v36 += 61
    switch ("\n\t" + "\n\t") {
    case "hello":
        fmt.Println(v34, v5, v4)
    case "":
        v36 = (^(v36))
    case "a":
        for fi39 := 0; fi39 < 2; fi39++ {
            if ((fi39 << ((18446744073709551615) & 63)) <= 42) {
                v40 := v1["foo bar"]
                v1[v40] = v40
                var v41 int16 = -1
                _ = v40
                _ = v41
            } else {
                v42 := v1[("hello" + "")]
                var v43 int16 = -32768
                for fi44 := 0; fi44 < 11; fi44++ {
                    fmt.Println(fi44, v37, v37)
                    if v34 {
                        fmt.Println(v4, v33)
                        var v45 uint8 = 127
                        switch (v5 << ((0) & 63)) {
                        case 128:
                            delete(v1, v42)
                        case -917:
                            var v46 int8 = 127
                            _ = v46
                        }
                        v47 := GenS3{F0: 4, F1: v37}
                        var v48 string = "hello"
                        _ = v45
                        _ = v47
                        _ = v48
                    } else {
                        v43 -= int16(v37)
                    }
                    if ((fi39 ^ (^(fi39))) == -171) {
                        if true {
                            v49 := GenS3{F0: 42, F1: v37}
                            if true {
                                v1["hello"] = "a"
                                v42 = "a"
                                switch ("a" + (((v42 + v42) + v42) + "\n\t")) {
                                case "foo bar":
                                    var v50 int64 = (^(v4))
                                    _ = v50
                                case "hello":
                                    var v51 string = "hello"
                                    _ = v51
                                }
                            } else {
                                if v34 {
                                    if false {
                                        v33 += (v33 / ((v33 * v33) * float64((-2.25 + (0.0 - 0.1)))))
                                    }
                                    if false {
                                        var v52 int16 = (v43 - v43)
                                        _ = v52
                                    }
                                    v53 := GenS4{F0: ("\n\t" + "\n\t"), F1: v43}
                                    if v34 {
                                        var v54 int32 = 695
                                        fmt.Println(v5, v42)
                                        var v55 uint = 255
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v53
                                }
                                v56 := GenS5{F0: false}
                                _ = v56
                            }
                            switch "" {
                            case "foo bar":
                                v57 := GenS5{F0: ((-128 == 100) && (!(v34 && v34)))}
                                delete(v1, ("" + (((v42 + (v42 + v42)) + "\n\t") + v42)))
                                _ = v57
                            case "a":
                                v58 := GenS4{F0: v42, F1: v43}
                                _ = v58
                            }
                            _ = v49
                        }
                        if v34 {
                            _ = len(v1)
                            fmt.Println(fi39, fi39)
                        } else {
                            v59 := GenS6{F0: ("a" + ((v42 + ("foo bar" + v42)) + v42))}
                            _ = v59
                        }
                        v60 := GenS7{F0: -1}
                        _ = v60
                    }
                    _ = fi44
                }
                _ = v42
                _ = v43
            }
            _ = fi39
        }
    default:
        for fi61 := 0; fi61 < 18; fi61++ {
            for fi62 := 0; fi62 < 10; fi62++ {
                v63 := GenS2{F0: fi62}
                _ = fi62
                _ = v63
            }
            for fi64 := 0; fi64 < 20; fi64++ {
                v5 = (fi64 * (fi61 | v5))
                switch ("" + "\n\t") {
                case "\n\t":
                    if (127 <= 0) {
                        var v65 uint16 = v37
                        var v66 int16 = 1
                        if true {
                            v33 += ((1e10 / 1e10) + (v33 / v33))
                            v67 := GenS8{F0: 51, F1: (^(v4)), F2: (v36 | v36)}
                            if v34 {
                                v3 = append(v3, 1)
                                fmt.Println(v4, v36)
                                var v68 uint16 = v37
                                _ = v68
                            }
                            _ = v67
                        } else {
                            v37 = 0
                            var v69 string = "a"
                            _ = v69
                        }
                        _ = v65
                        _ = v66
                    }
                case "a":
                    var v70 int = (fi61 / ((fi64) | 1))
                    v71 := GenS2{F0: (v70 % ((418) | 1))}
                    var v72 uint16 = (^(v37))
                    v37 += 1
                    _ = v70
                    _ = v71
                    _ = v72
                case "":
                    if (!(!((v33 == 0.1) || false))) {
                        if (!true) {
                            if (v34 || (v4 <= (^(v4)))) {
                                v36 += 0
                                v73 := GenS8{F0: 971, F1: (v4 + v36), F2: (v36 ^ 42)}
                                _ = v73
                            }
                            if (1e10 <= 1e10) {
                                v4 = (v4 % (((v36 | (v4 ^ -960))) | 1))
                                v34 = true
                                v74 := GenS3{F0: 18446744073709551615, F1: v37}
                                _ = v74
                            } else {
                                v75 := GenS8{F0: 42, F1: v36, F2: v4}
                                _ = len(v1)
                                _ = v75
                            }
                            fmt.Println(v36, v34)
                        }
                        switch "a" {
                        case "\n\t":
                            _ = len(v2)
                        case "a":
                            var v76 int = (fi64 & fi61)
                            _ = v76
                        case "foo bar":
                            v36 = (^(v36))
                        default:
                            v77 := GenS9{F0: 1}
                            _ = v77
                        }
                    } else {
                        v33 += v33
                    }
                default:
                    fmt.Println(fi64)
                }
                switch fi64 {
                case 351:
                    v78 := GenS2{F0: 933}
                    _ = v78
                case -986:
                    v1[("" + "\n\t")] = "\n\t"
                default:
                    fmt.Println(v33, fi61, v36)
                }
                switch "foo bar" {
                case "\n\t":
                    var v79 int = -893
                    _ = v79
                case "":
                    v80 := GenS10{F0: v34, F1: v34, F2: 100}
                    _ = v80
                case "a":
                    var v81 int64 = int64(fi64)
                    _ = v81
                default:
                    v33 = v33
                }
                v82 := GenS2{F0: int(v36)}
                _ = fi64
                _ = v82
            }
            _ = fi61
        }
    }
    v83 := GenS7{F0: 957}
    if v34 {
        delete(v1, (("a" + ((("foo bar" + "foo bar") + "hello") + "foo bar")) + "a"))
        v34 = v34
        if false {
            if true {
                var v84 float32 = (3.14159 * float32(float64(0.1)))
                _ = v84
            }
            v36 -= v4
            v4 = (v36 * (v4 % (((v4 / (((v4 * v36)) | 1))) | 1)))
        }
    } else {
        var v85 bool = true
        for fi86 := 0; fi86 < 5; fi86++ {
            if (!true) {
                v87 := GenS10{F0: true, F1: v34, F2: -65}
                v88 := GenS11{F0: 0, F1: 184, F2: v85}
                _ = v87
                _ = v88
            }
            v36 = 482
            if (500 <= (v4 - (-(v4)))) {
                fmt.Println(v5, fi86, v5)
            }
            _ = fi86
        }
        if len(v3) > 0 {
            v3[0] = 65
        }
        _ = v85
    }
    v89 := GenS12{F0: 886}
    for fi90 := 0; fi90 < 7; fi90++ {
        v91 := GenS13{F0: ("\n\t" + ("hello" + "foo bar"))}
        v33 = float64((1.5 + -2.25))
        _ = fi90
        _ = v91
    }
    var v92 uint32 = 65535
    fmt.Println(v4, v5, v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v83
    _ = v89
    _ = v92
}
