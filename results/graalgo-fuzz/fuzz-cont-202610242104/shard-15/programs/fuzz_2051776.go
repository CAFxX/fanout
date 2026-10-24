package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 bool }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 uint; F2 int16 }
type GenS5 struct { F0 uint }
func genFn0() uint32 {
    var v7 int32 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{(fi8 << ((1) & 63)), fi8}
        v10 := GenS1{F0: 42, F1: false, F2: true}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int16 = 0
    v11 += (v11 - v11)
    var v12 float32 = (0.0 - (3.14159 * 0.1))
    _ = v11
    _ = v12
    _ = v7
    return 810
}

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 32767
    var v3 float64 = ((0.1 / (float64(v1) - float64(1e10))) / 1.0)
    var v4 int16 = 1
    v5 := []string{"\n\t", "hello"}
    v6 := []int{42}
    if (v2 > v2) {
        var v13 int32 = -2147483648
        if (false || false) {
            v6 = append(v6, 4294967295)
        }
        _ = v13
    }
    switch 0 {
    case -331:
        for fi14 := 0; fi14 < 12; fi14++ {
            v3 -= 0.0
            _ = fi14
        }
    }
    v2 = (^(v2))
    if (!false) {
        fmt.Println(v3)
    }
    if (-2147483648 < 42) {
        if len(v6) > 0 {
            v6[0] = 990
        }
    } else {
        fmt.Println(v4, v2)
        var v15 uint32 = 19
        if len(v6) > 0 {
            v6[0] = -214
        }
        _ = v15
    }
    v4 += 32767
    v16 := GenS2{F0: -1, F1: 635, F2: 128}
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 uint = 42
        v1 += 1e10
        _ = fi17
        _ = v18
    }
    if ((v1 / v1) == v1) {
        switch "a" {
        case "hello":
            v1 = (v1 - 3.14159)
        case "":
            _ = len(v5)
        default:
            v2 += (v2 | (v2 ^ 809))
        }
    } else {
        var v19 int = 42
        for fi20 := 0; fi20 < 18; fi20++ {
            if (!false) {
                var v21 float32 = -2.25
                if len(v5) > 0 {
                    v5[0] = ("" + "a")
                }
                var v22 string = "hello"
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 4; fi23++ {
                v24 := GenS2{F0: v19, F1: 0, F2: 1}
                switch ("hello" + "hello") {
                case "":
                    var v25 int16 = v4
                    _ = v25
                case "a":
                    _ = len(v5)
                default:
                    var v26 bool = false
                    _ = v26
                }
                switch (v19 << ((223) & 63)) {
                case 32767:
                    var v27 float32 = 1e10
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
            for fi28 := 0; fi28 < 11; fi28++ {
                v29 := GenS2{F0: fi28, F1: 784, F2: 255}
                var v30 uint64 = 2
                switch "foo bar" {
                case "hello":
                    var v31 int = (v19 * fi28)
                    _ = v31
                case "\n\t":
                    if false {
                        var v32 int32 = -262
                        var v33 int64 = int64(v2)
                        _ = v32
                        _ = v33
                    } else {
                        var v34 int16 = 369
                        _ = v34
                    }
                case "a":
                    switch 501 {
                    case 32767:
                        fmt.Println(v3, v3, fi28)
                    case 9223372036854775807:
                        if (((v4 - int16(v30)) >= (v4 % ((-32768) | 1))) || false) {
                            if ((v4 != (v4 + 830)) || false) {
                                var v35 uint16 = (v2 + 65535)
                                _ = v35
                            }
                            fmt.Println(v1)
                            if (v30 < v30) {
                                if true {
                                    v19 -= -2147483648
                                    v1 = v1
                                } else {
                                    if (true || (0.0 >= ((v1 / v1) / 1.0))) {
                                        v36 := GenS1{F0: -431, F1: true, F2: false}
                                        _ = v36
                                    }
                                    var v37 int = (^(fi20))
                                    _ = v37
                                }
                                _ = len(v6)
                                var v38 float64 = 0.0
                                _ = v38
                            } else {
                                if true {
                                    var v39 float64 = float64(3.14159)
                                    fmt.Println(v19)
                                    if (v4 < v4) {
                                        v40 := GenS3{F0: (1 > 1)}
                                        v41 := GenS2{F0: (fi28 ^ v19), F1: (v30 + uint64(v2)), F2: 127}
                                        var v42 bool = ("a" == "")
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 uint = 255
                                    var v44 uint16 = uint16(v43)
                                    _ = v39
                                    _ = v43
                                    _ = v44
                                }
                            }
                        } else {
                            var v45 bool = (9223372036854775807 < 1)
                            v46 := GenS2{F0: (-(v19)), F1: (^(v30)), F2: 127}
                            _ = v45
                            _ = v46
                        }
                    default:
                        var v47 int16 = (^(v4))
                        _ = v47
                    }
                default:
                    v48 := GenS4{F0: -1, F1: 361, F2: (v4 ^ (v4 % ((-1) | 1)))}
                    var v49 uint = 42
                    if false {
                        if false {
                            fmt.Println(v19, fi20, v3)
                            var v50 uint32 = 338
                            _ = v50
                        }
                        var v51 int32 = 685
                        if ((1e10 != 3.14159) && (1.5 < 0.0)) {
                            fmt.Println(v4, fi20, v4)
                            _ = len(v6)
                            v1 -= (v1 + v1)
                        } else {
                            v4 = (v4 & 1000)
                            v19 += fi20
                            if ((v2 >> ((657) & 63)) < v2) {
                                fmt.Println(fi20, v1)
                                v52 := GenS3{F0: false}
                                v3 -= v3
                                v53 := GenS2{F0: (fi28 + (fi20 >> ((65535) & 63))), F1: 18446744073709551615, F2: 148}
                                _ = v52
                                _ = v53
                            }
                        }
                        _ = v51
                    } else {
                        _ = len(v6)
                        if (true || ((0 != genFn0()) || true)) {
                            if (!true) {
                                var v54 uint32 = 42
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v54
                            } else {
                                v55 := GenS1{F0: -1, F1: true, F2: false}
                                var v56 float32 = float32(v3)
                                var v57 float64 = v3
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            fmt.Println(fi20)
                            switch fi28 {
                            case -582:
                                v58 := GenS5{F0: 65535}
                                _ = v58
                            }
                        } else {
                            var v59 uint = v49
                            var v60 int64 = 9223372036854775807
                            var v61 int8 = 100
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                        v19 += 363
                    }
                    _ = len(v6)
                    _ = v48
                    _ = v49
                }
                var v62 uint16 = 32767
                v30 = v30
                _ = fi28
                _ = v29
                _ = v30
                _ = v62
            }
            _ = fi20
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
