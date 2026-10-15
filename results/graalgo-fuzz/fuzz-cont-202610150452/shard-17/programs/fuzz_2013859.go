package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int8; F1 int; F2 bool }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 string; F2 int32 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0() int64 {
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint = 1
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 612}
    _ = v6
    _ = v9
    return 993
}
func genFn1(p0 string) uint16 {
    fmt.Println(p0)
    var v10 string = p0
    fmt.Println(v10, v10)
    if (!true) {
        p0 = ((v10 + p0) + "hello")
        v11 := []int{32767}
        switch 2147483647 {
        case -2147483648:
            v12 := GenS1{F0: 1, F1: 2147483647}
            _ = v12
        case -783:
            if len(v11) > 0 {
                v11[0] = -1
            }
        }
        _ = v11
    } else {
        v13 := []int{}
        _ = v13
    }
    if true {
        var v14 bool = false
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS1{F0: 335, F1: fi15}
            _ = fi15
            _ = v16
        }
        v14 = (((1.5 * (-2.25 + 1.5)) - 0.0) == (0.1 * 0.1))
        _ = v14
    }
    _ = p0
    _ = v10
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[string]string)
    var v3 int = 32767
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int16 = 802
        _ = fi17
        _ = v18
    }
    var v19 string = "\n\t"
    var v20 int64 = 588
    fmt.Println(v19, v3)
    for fi21 := 0; fi21 < 7; fi21++ {
        v20 = genFn0()
        for fi22 := 0; fi22 < 1; fi22++ {
            if false {
                var v23 uint8 = 128
                _ = v23
            } else {
                v19 = (((("a" + v19) + (v19 + ((v19 + ((v19 + "\n\t") + v19)) + v19))) + "") + (("foo bar" + v19) + v19))
                v24 := GenS1{F0: 641, F1: 4294967295}
                _ = v24
            }
            fmt.Println(v19)
            _ = fi22
        }
        v25 := GenS1{F0: 1, F1: -9223372036854775808}
        _ = fi21
        _ = v25
    }
    _ = len(v4)
    var v26 uint32 = 42
    switch v3 {
    case 0:
        for fi27 := 0; fi27 < 10; fi27++ {
            v19 = v19
            _ = fi27
        }
    case 1:
        v20 += 1
    case -817:
        v20 -= genFn0()
    default:
        var v28 float64 = ((0.0 * v1) + (float64(0.1) * (float64(((3.14159 * (3.14159 * 3.14159)) / 3.14159)) - v1)))
        _ = v28
    }
    fmt.Println(v19, v3, v20)
    var v29 float32 = (((0.1 + 1e10) / -2.25) / 1e10)
    var v30 uint32 = 2147483648
    v3 = -184
    var v31 uint16 = 77
    var v32 float32 = v29
    var v33 int32 = 2147483647
    if (v3 < (v3 << ((2) & 63))) {
        v34 := GenS2{F0: 126, F1: int(v26), F2: ("a" < "foo bar")}
        v35 := []int{}
        _ = v34
        _ = v35
    } else {
        var v36 int8 = 1
        switch (v3 ^ (-(v3))) {
        case 818:
            for fi37 := 0; fi37 < 7; fi37++ {
                fmt.Println(v31, v19)
                _ = fi37
            }
        case -2147483648:
            var v38 int8 = (^(v36))
            _ = v38
        case -1:
            for fi39 := 0; fi39 < 15; fi39++ {
                for fi40 := 0; fi40 < 12; fi40++ {
                    switch (v3 / ((32768) | 1)) {
                    case 42:
                        v19 = v19
                        if true {
                            v41 := []int{}
                            v3 = (fi40 >> ((4294967295) & 63))
                            _ = v41
                        } else {
                            var v42 uint32 = (^(v30))
                            _ = v42
                        }
                        v30 = v26
                        if (!(253 > 1)) {
                            v43 := GenS3{F0: 4294967295, F1: uint64(v31)}
                            if ((v31 + (v31 & 255)) != (v31 + (v31 * v31))) {
                                fmt.Println(v32, v29)
                                var v44 int64 = (v20 >> ((1) & 63))
                                _ = v44
                            }
                            v19 = ""
                            _ = v43
                        }
                        fmt.Println(v30, v29, v31)
                    case 364:
                        var v45 int32 = v33
                        _ = v45
                    }
                    var v46 int = (-(fi39))
                    if true {
                        var v47 int = 2147483647
                        v26 = (^(v30))
                        if (0 < 255) {
                            v48 := GenS3{F0: 18446744073709551615, F1: 373}
                            if (0.1 == (3.14159 * v1)) {
                                var v49 int = fi39
                                v36 = (v36 >> ((343) & 63))
                                _ = v49
                            } else {
                                v32 -= ((-2.25 - v29) + (v29 / (v32 - -2.25)))
                            }
                            _ = v48
                        }
                        _ = v47
                    } else {
                        var v50 uint16 = (v31 - (v31 >> ((18446744073709551615) & 63)))
                        v26 = (v26 / ((494) | 1))
                        switch v19 {
                        case "hello":
                            v51 := GenS4{F0: v33}
                            _ = v51
                        case "":
                            v52 := []int{191, 882}
                            _ = v52
                        case "\n\t":
                            fmt.Println(v32)
                        default:
                            switch v3 {
                            case -624:
                                v53 := GenS2{F0: v36, F1: (v3 >> ((18446744073709551615) & 63)), F2: (false || (!((^(v26)) != (v30 - uint32(v31)))))}
                                v54 := GenS5{F0: 32767, F1: ((v19 + v19) + ""), F2: v33}
                                v19 = "foo bar"
                                if false {
                                    switch v19 {
                                    case "foo bar":
                                        if (int(v20) <= (fi40 << ((42) & 63))) {
                                            v55 := v4["foo bar"]
                                            v32 -= 0.0
                                            _ = v55
                                        } else {
                                            v31 = 0
                                            delete(v5, ((v19 + v19) + ""))
                                            _ = len(v2)
                                        }
                                    case "a":
                                        switch v19 {
                                        case "hello":
                                            var v56 int16 = 0
                                            _ = v56
                                        case "":
                                            fmt.Println(v31, v30)
                                        case "foo bar":
                                            fmt.Println(v36, v29)
                                        default:
                                            fmt.Println(v3)
                                        }
                                    case "hello":
                                        switch 885 {
                                        case -238:
                                            v57 := v2[v19]
                                            _ = v57
                                        case 4294967295:
                                            if false {
                                                var v58 float32 = -2.25
                                                v59 := GenS6{F0: 2, F1: 2, F2: v26}
                                                fmt.Println(fi40, v46)
                                                _ = v58
                                                _ = v59
                                            } else {
                                                switch v19 {
                                                case "foo bar":
                                                    v60 := GenS3{F0: 2, F1: 633}
                                                    _ = v60
                                                case "a":
                                                    if false {
                                                        fmt.Println(v36, v29)
                                                        v61 := v4[(((v19 + v19) + "hello") + "foo bar")]
                                                        v62 := GenS7{F0: (-(v20)), F1: (v30 ^ (v26 | v30))}
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                case "\n\t":
                                                    v63 := GenS1{F0: 1, F1: fi39}
                                                    _ = v63
                                                }
                                                var v64 float32 = 3.14159
                                                _ = v64
                                            }
                                        case 128:
                                            var v65 uint64 = 2
                                            _ = v65
                                        }
                                        v66 := GenS4{F0: int32(v26)}
                                        var v67 int = (fi40 / (((^(v3))) | 1))
                                        fmt.Println(v30)
                                        switch (-(fi40)) {
                                        case -2147483648:
                                            var v68 int64 = v20
                                            v50 = v50
                                            _ = v68
                                        case 1:
                                            if (v19 >= "a") {
                                                var v69 int = 481
                                                v70 := []int{v69, fi40, (v3 >> ((918) & 63))}
                                                _ = v69
                                                _ = v70
                                            }
                                            if (true || true) {
                                                v46 -= (v3 >> ((uint(v31)) & 63))
                                            }
                                        case 32767:
                                            v71 := GenS3{F0: 42, F1: 19}
                                            _ = v71
                                        }
                                        var v72 string = "hello"
                                        _ = v66
                                        _ = v67
                                        _ = v72
                                    }
                                    var v73 bool = true
                                    _ = v73
                                } else {
                                    _ = len(v5)
                                    v74 := GenS8{F0: 0, F1: 32767, F2: v30}
                                    v1 += v1
                                    _ = v74
                                }
                                _ = v53
                                _ = v54
                            }
                        }
                        _ = v50
                    }
                    _ = fi40
                    _ = v46
                }
                _ = fi39
            }
        }
        _ = v36
    }
    var v75 uint16 = 336
    for fi76 := 0; fi76 < 12; fi76++ {
        var v77 bool = true
        _ = fi76
        _ = v77
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v75
}
