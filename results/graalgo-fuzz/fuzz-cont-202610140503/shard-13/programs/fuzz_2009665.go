package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 string }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 int32; F2 uint8 }
type GenS9 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint8; F1 int16; F2 int16 }
func genFn0(p0 int16, p1 int64) int8 {
    v7 := []int{-9223372036854775808}
    if false {
        var v8 uint64 = 1
        p0 = 42
        _ = v8
    } else {
        v9 := GenS1{F0: 255, F1: 42, F2: "\n\t"}
        var v10 int16 = (^(p0))
        p0 -= p0
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 0
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]bool)
    var v3 uint8 = 255
    v4 := []int{327}
    v5 := make(map[int]string)
    var v6 uint8 = 41
    fmt.Println(v1, v1, v6)
    switch v1 {
    case "":
        v1 = (("foo bar" + v1) + v1)
    case "foo bar":
        for fi11 := 0; fi11 < 8; fi11++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi11
        }
    case "a":
        var v12 int = 4294967295
        _ = v12
    default:
        if (("hello" != v1) || (42 < 2147483648)) {
            fmt.Println(v3)
            if false {
                v1 = v1
            } else {
                v13 := GenS2{F0: -80, F1: 42, F2: -9223372036854775808}
                _ = v13
            }
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS3{F0: (fi14 ^ fi14), F1: -1}
                var v16 int32 = 1
                v17 := GenS4{F0: fi14, F1: 191, F2: 4294967295}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(fi18, v1, v6)
                v1 = (((v1 + "") + (("a" + ("\n\t" + "foo bar")) + v1)) + (("foo bar" + (v1 + (v1 + "foo bar"))) + (("hello" + v1) + (("hello" + (v1 + ((v1 + v1) + "\n\t"))) + v1))))
                if len(v4) > 0 {
                    v4[0] = 160
                }
                _ = fi18
            }
            switch 127 {
            case -1:
                fmt.Println(v1, v3)
            case 127:
                v19 := GenS5{F0: 1}
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 8; fi20++ {
        if (2147483647 > -685) {
            switch ("a" + v1) {
            case "foo bar":
                v1 = ((((v1 + v1) + ((v1 + v1) + ((v1 + v1) + v1))) + (v1 + "foo bar")) + "a")
            default:
                v21 := GenS6{F0: 9223372036854775807, F1: 65535}
                _ = v21
            }
            for fi22 := 0; fi22 < 7; fi22++ {
                switch fi20 {
                case -108:
                    v1 = v1
                case 289:
                    v6 += v6
                case 42:
                    var v23 int16 = -32768
                    v3 = (v6 + 0)
                    _ = v23
                default:
                    v3 -= 2
                    if len(v4) > 0 {
                        v4[0] = (fi22 % (((fi22 << ((65535) & 63))) | 1))
                    }
                    if (true && true) {
                        switch (v1 + v1) {
                        case "hello":
                            var v24 int32 = -1
                            _ = v24
                        case "a":
                            fmt.Println(v1)
                        case "":
                            var v25 uint = 56
                            if ((!(("hello" + v1) < "hello")) && true) {
                                v26 := GenS2{F0: -2147483648, F1: 42, F2: 1}
                                _ = v26
                            } else {
                                v25 = v25
                                _ = len(v5)
                                if ((true && false) || (((1.5 + 0.1) > (0.0 + (-2.25 * 0.1))) && (!false))) {
                                    switch (fi22 % ((fi22) | 1)) {
                                    case 0:
                                        var v27 int8 = 1
                                        _ = v27
                                    case -174:
                                        if false {
                                            switch (v1 + v1) {
                                            case "hello":
                                                fmt.Println(fi22, fi20)
                                            case "foo bar":
                                                v28 := GenS3{F0: fi20, F1: 1}
                                                _ = v28
                                            default:
                                                var v29 int16 = 1
                                                _ = v29
                                            }
                                            var v30 uint = 747
                                            _ = v30
                                        } else {
                                            v25 = (v25 / (((^(v25))) | 1))
                                            v31 := GenS6{F0: fi20, F1: (v25 >> ((163) & 63))}
                                            _ = v31
                                        }
                                        if false {
                                            var v32 float32 = -2.25
                                            _ = v32
                                        } else {
                                            if (true || true) {
                                                if false {
                                                    v33 := GenS7{F0: true}
                                                    v1 = "a"
                                                    _ = v33
                                                } else {
                                                    v1 = v1
                                                }
                                                var v34 uint32 = 1
                                                _ = v34
                                            } else {
                                                v3 = 127
                                                v35 := GenS1{F0: 255, F1: fi22, F2: ""}
                                                _ = v35
                                            }
                                            var v36 int32 = -171
                                            v37 := GenS5{F0: 18446744073709551615}
                                            _ = v36
                                            _ = v37
                                        }
                                        var v38 float64 = (3.14159 * (float64(float32(1.5)) + (0.1 / (1e10 * 0.1))))
                                        _ = v38
                                    }
                                    switch 533 {
                                    case -1:
                                        v6 = v3
                                    case -687:
                                        if false {
                                            v4 = append(v4, fi22)
                                            v39 := GenS1{F0: 173, F1: fi20, F2: "hello"}
                                            _ = v39
                                        } else {
                                            v40 := GenS8{F0: (2 < 800), F1: 42, F2: (^(v6))}
                                            if ((v25 % ((v25) | 1)) != (v25 >> ((65535) & 63))) {
                                                v41 := v2[v1]
                                                if ((v6 == v3) && v41) {
                                                    v4 = append(v4, fi20)
                                                    switch -895 {
                                                    case 128:
                                                        v42 := GenS2{F0: 1, F1: 1, F2: fi22}
                                                        _ = v42
                                                    case 9223372036854775807:
                                                        v1 = v1
                                                    case 1:
                                                        v6 += v3
                                                    default:
                                                        fmt.Println(fi22, v6)
                                                    }
                                                }
                                                v43 := v2[((v1 + ("" + (((((("hello" + ((v1 + v1) + v1)) + v1) + "a") + ("\n\t" + v1)) + (v1 + ((("\n\t" + "") + v1) + ((v1 + v1) + (("" + (("hello" + ((("foo bar" + (((v1 + v1) + (v1 + "\n\t")) + v1)) + v1) + "a")) + ("a" + "hello"))) + v1))))) + (v1 + v1)))) + "hello")]
                                                _ = v41
                                                _ = v43
                                            }
                                            _ = v40
                                        }
                                    case 688:
                                        if ((true || ((^(v25)) < 1)) && ((fi20 / (((fi20 / ((fi20) | 1))) | 1)) != (fi22 ^ fi22))) {
                                            _ = len(v4)
                                            var v44 uint64 = 18446744073709551615
                                            fmt.Println(v3, v6)
                                            _ = v44
                                        }
                                    default:
                                        v45 := GenS9{F0: 95, F1: v6, F2: 127}
                                        _ = v45
                                    }
                                    v3 = (v6 << ((v25) & 63))
                                    v6 = 42
                                }
                                v3 = 170
                            }
                            _ = v25
                        }
                        fmt.Println(fi20, fi20)
                        if ((true && (2 >= 612)) || false) {
                            var v46 int = (fi22 & -2147483648)
                            var v47 int64 = -129
                            var v48 int32 = -2147483648
                            _ = v46
                            _ = v47
                            _ = v48
                        } else {
                            v3 -= (v3 & 255)
                            v5[-1] = v1
                            v49 := GenS3{F0: 1, F1: -1}
                            var v50 bool = true
                            _ = v49
                            _ = v50
                        }
                    } else {
                        switch fi22 {
                        case -935:
                            fmt.Println(fi22, v1, v6)
                        case -303:
                            v51 := GenS9{F0: v3, F1: (v6 | v6), F2: 30}
                            _ = v51
                        default:
                            switch fi22 {
                            case -356:
                                var v52 int = (fi20 + (fi20 & (fi20 | (fi22 / ((fi22) | 1)))))
                                _ = v52
                            }
                        }
                    }
                }
                switch 9223372036854775807 {
                case 9223372036854775807:
                    var v53 int8 = 98
                    _ = v53
                default:
                    if false {
                        v3 -= (v3 >> ((65535) & 63))
                    }
                }
                v3 += 105
                v54 := GenS4{F0: fi20, F1: 2, F2: 2}
                if (-2147483648 != 0) {
                    if true {
                        v4 = append(v4, 32767)
                    }
                } else {
                    v5[1] = v1
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v3 = v6
                }
                _ = fi22
                _ = v54
            }
            var v55 int32 = 2147483647
            _ = v55
        }
        _ = fi20
    }
    var v56 int16 = -930
    var v57 uint64 = 549
    var v58 float32 = 0.0
    v59 := GenS10{F0: v6, F1: 449, F2: -32768}
    if true {
        if false {
            fmt.Println(v1)
        } else {
            fmt.Println(v1, v1, v58)
            v60 := GenS4{F0: -2147483648, F1: 327, F2: 70}
            v61 := v5[32768]
            _ = v60
            _ = v61
        }
    }
    v6 = 1
    v62 := GenS6{F0: -2147483648, F1: 813}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v62
}
