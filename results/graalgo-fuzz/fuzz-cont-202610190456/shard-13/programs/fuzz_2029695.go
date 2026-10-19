package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int }
func genFn0() uint {
    var v7 int64 = -1
    if ((((v7 * -9223372036854775808) <= (v7 >> ((42) & 63))) || true) && (0 == 32767)) {
        v8 := GenS1{F0: -32768, F1: (false || true)}
        switch ("foo bar" + "\n\t") {
        case "a":
            v7 -= v7
        default:
            var v9 uint32 = 1
            _ = v9
        }
        _ = v8
    } else {
        v7 -= (v7 >> ((65535) & 63))
        var v10 bool = (false && false)
        _ = v10
    }
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 2; fi12++ {
        if (v7 > -1) {
            v7 -= (v7 | v7)
        }
        _ = fi12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float64 = 1e10
        v15 := GenS1{F0: 0, F1: false}
        _ = fi13
        _ = v14
        _ = v15
    }
    _ = v11
    _ = v7
    return 175
}
func genFn1(p0 int64, p1 int32, p2 uint16) int8 {
    v16 := []int{128, 42, -958}
    if true {
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 14; fi18++ {
                p0 = p0
                v19 := GenS1{F0: 617, F1: true}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return -128
}

func main() {
    var v1 int = -9223372036854775808
    var v2 int64 = -103
    v3 := []uint32{4294967295, 42, 2, 2}
    v4 := make(map[string]bool)
    var v5 float32 = ((((1e10 - (((1.5 - -2.25) + -2.25) / (1e10 - 3.14159))) * 0.0) + 3.14159) + 1.5)
    var v6 uint64 = 337
    switch "\n\t" {
    case "a":
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 int8 = 79
            var v22 int8 = genFn1(747, 543, 32767)
            _ = fi20
            _ = v21
            _ = v22
        }
    case "hello":
        v1 -= (-(v1))
    default:
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint32 = 4294967295
            var v25 int32 = -1
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v26 float32 = v5
    if false {
        if false {
            for fi27 := 0; fi27 < 8; fi27++ {
                fmt.Println(v1, fi27)
                fmt.Println(v2, v1)
                _ = fi27
            }
            v28 := GenS1{F0: 32767, F1: (959 >= 32767)}
            if true {
                if true {
                    v6 += v6
                    for fi29 := 0; fi29 < 6; fi29++ {
                        if false {
                            var v30 float64 = ((0.0 - float64(v26)) + 0.0)
                            _ = v30
                        } else {
                            v26 = float32(3.14159)
                            var v31 int16 = int16(v6)
                            _ = v31
                        }
                        _ = fi29
                    }
                    for fi32 := 0; fi32 < 8; fi32++ {
                        var v33 float32 = 1.5
                        _ = fi32
                        _ = v33
                    }
                } else {
                    var v34 bool = false
                    fmt.Println(v6)
                    _ = len(v3)
                    _ = v34
                }
            } else {
                v35 := GenS1{F0: 42, F1: (-1 < int8(v2))}
                var v36 float32 = ((1e10 / -2.25) + (0.0 + v26))
                _ = v35
                _ = v36
            }
            _ = v28
        } else {
            _ = len(v4)
            var v37 float32 = v5
            var v38 uint16 = 2
            _ = v37
            _ = v38
        }
    } else {
        delete(v4, ("\n\t" + "hello"))
        for fi39 := 0; fi39 < 14; fi39++ {
            v40 := GenS2{F0: 255, F1: genFn1((v2 + 0), -288, 2)}
            _ = fi39
            _ = v40
        }
    }
    switch v1 {
    case -9223372036854775808:
        fmt.Println(v2, v5)
    case -2147483648:
        for fi41 := 0; fi41 < 1; fi41++ {
            for fi42 := 0; fi42 < 13; fi42++ {
                v43 := GenS2{F0: 42, F1: 27}
                if true {
                    v4["hello"] = false
                    if (!(1 != 97)) {
                        v5 = v26
                        if false {
                            fmt.Println(v2)
                            fmt.Println(fi42, v5)
                        } else {
                            _ = len(v4)
                            var v44 bool = (!true)
                            switch (fi42 >> ((623) & 63)) {
                            case 9223372036854775807:
                                var v45 bool = false
                                _ = v45
                            default:
                                _ = len(v4)
                            }
                            if (genFn0() != 950) {
                                fmt.Println(v2, v44, v6)
                                fmt.Println(fi41, v5, fi41)
                                v46 := GenS1{F0: 364, F1: false}
                                v47 := v4[""]
                                _ = v46
                                _ = v47
                            }
                            fmt.Println(v6)
                            v1 = (fi42 % ((1) | 1))
                            var v48 string = ((("" + "foo bar") + "\n\t") + ("\n\t" + "\n\t"))
                            _ = v44
                            _ = v48
                        }
                        v26 -= v26
                    } else {
                        _ = len(v4)
                        delete(v4, "a")
                    }
                    var v49 uint16 = 32767
                    _ = v49
                } else {
                    delete(v4, "a")
                }
                v5 = v5
                if true {
                    var v50 int16 = -32768
                    v5 -= 3.14159
                    v51 := GenS2{F0: 42, F1: 48}
                    v50 = 0
                    switch fi42 {
                    case 32767:
                        v1 -= fi42
                        v1 = v1
                    default:
                        switch fi41 {
                        case 42:
                            if true {
                                v52 := v4["\n\t"]
                                _ = v52
                            }
                        case 0:
                            if ((3.14159 + 1e10) < 1e10) {
                                fmt.Println(v1, v2, v5)
                            } else {
                                var v53 bool = true
                                _ = v53
                            }
                        }
                    }
                    _ = v50
                    _ = v51
                } else {
                    switch (^(fi41)) {
                    case 773:
                        switch ("hello" + ("hello" + (("\n\t" + ("a" + "hello")) + "\n\t"))) {
                        case "a":
                            switch (^(fi42)) {
                            case -58:
                                v54 := GenS2{F0: 2, F1: 100}
                                _ = v54
                            case 4294967295:
                                v26 += 3.14159
                            }
                        case "":
                            var v55 int8 = 104
                            _ = v55
                        }
                    default:
                        if (false && true) {
                            var v56 int8 = 17
                            _ = v56
                        }
                    }
                    if false {
                        if true {
                            v57 := GenS3{F0: 496, F1: -2147483648, F2: "a"}
                            var v58 uint64 = (^(v6))
                            _ = v57
                            _ = v58
                        } else {
                            var v59 uint = 0
                            switch "hello" {
                            case "hello":
                                fmt.Println(v2)
                                switch "hello" {
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                    var v60 uint64 = v6
                                    _ = v60
                                case "foo bar":
                                    if false {
                                        v6 -= (v6 >> ((42) & 63))
                                        fmt.Println(v59)
                                        var v61 uint16 = 255
                                        _ = v61
                                    }
                                case "":
                                    v1 = (fi41 + (v1 & 2147483647))
                                default:
                                    var v62 int8 = 127
                                    _ = v62
                                }
                            case "a":
                                v59 -= 0
                            case "foo bar":
                                fmt.Println(v5, v6, v26)
                            default:
                                var v63 int = -9223372036854775808
                                delete(v4, ("foo bar" + (("foo bar" + "a") + "a")))
                                _ = len(v3)
                                v63 -= (fi41 * (fi41 - fi41))
                                _ = v63
                            }
                            var v64 uint16 = 255
                            _ = v59
                            _ = v64
                        }
                    } else {
                        v26 = v5
                        var v65 uint64 = (v6 * 724)
                        _ = v65
                    }
                    v2 += v2
                }
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
    default:
        var v66 uint8 = 42
        _ = v66
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    if (!((((((v5 * 1e10) + v5) / 1.5) * ((((v26 * float32(3.14159)) - 0.1) - v26) * float32((-2.25 + 1.5)))) * -2.25) != v5)) {
        delete(v4, ("" + ("hello" + "")))
        v6 = v6
        v67 := GenS4{F0: (v6 - 310), F1: 0}
        _ = v67
    } else {
        v1 += 128
        v68 := GenS3{F0: 1, F1: -699, F2: "foo bar"}
        for fi69 := 0; fi69 < 16; fi69++ {
            v1 = v1
            if (!(!true)) {
                v6 = (v6 / ((v6) | 1))
            } else {
                v70 := GenS5{F0: 123, F1: v1}
                if (!(!true)) {
                    var v71 float32 = (0.1 / v5)
                    v26 += (3.14159 / v26)
                    _ = v71
                } else {
                    var v72 int64 = v2
                    _ = v72
                }
                _ = v70
            }
            if false {
                var v73 float32 = ((-2.25 * (float32(3.14159) + 0.1)) * (v5 / v5))
                fmt.Println(v2, v2, v2)
                var v74 int64 = -1
                _ = v73
                _ = v74
            } else {
                var v75 int16 = 1000
                var v76 bool = (9223372036854775807 > (v1 + 127))
                _ = v75
                _ = v76
            }
            _ = fi69
        }
        _ = v68
    }
    for fi77 := 0; fi77 < 3; fi77++ {
        for fi78 := 0; fi78 < 18; fi78++ {
            var v79 uint = 4294967295
            var v80 string = "hello"
            _ = fi78
            _ = v79
            _ = v80
        }
        switch ("foo bar" + "a") {
        case "foo bar":
            fmt.Println(v1, v26)
        }
        for fi81 := 0; fi81 < 6; fi81++ {
            if (2 <= 42) {
                _ = len(v4)
                v82 := GenS4{F0: v6, F1: 255}
                _ = len(v4)
                _ = v82
            }
            _ = fi81
        }
        _ = fi77
    }
    v1 += (v1 | v1)
    v83 := GenS5{F0: 128, F1: v1}
    var v84 uint64 = (v6 / ((v6) | 1))
    v3 = append(v3, 42)
    for fi85 := 0; fi85 < 18; fi85++ {
        v86 := GenS3{F0: 827, F1: 2147483647, F2: "hello"}
        var v87 float32 = ((v5 - v26) * v26)
        for fi88 := 0; fi88 < 13; fi88++ {
            var v89 float32 = v87
            _ = fi88
            _ = v89
        }
        _ = fi85
        _ = v86
        _ = v87
    }
    if false {
        var v90 uint8 = 1
        _ = v90
    }
    var v91 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v83
    _ = v84
    _ = v91
}
