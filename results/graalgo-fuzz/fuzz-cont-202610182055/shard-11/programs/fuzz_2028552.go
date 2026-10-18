package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS7 struct { F0 string; F1 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int8; F1 int }
type GenS10 struct { F0 int8; F1 uint64; F2 string }
func genFn0() uint64 {
    var v6 uint64 = 42
    var v7 int8 = 127
    if false {
        if false {
            fmt.Println(v6, v6)
            fmt.Println(v7)
            for fi8 := 0; fi8 < 7; fi8++ {
                v7 = (v7 / ((v7) | 1))
                var v9 int = (^(fi8))
                v10 := []int{32768, (fi8 & fi8), (v9 % ((v9) | 1))}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            v11 := GenS1{F0: -19}
            _ = v11
        }
        var v12 int32 = -2147483648
        var v13 uint = 0
        _ = v12
        _ = v13
    } else {
        if true {
            v14 := GenS1{F0: -794}
            _ = v14
        } else {
            if (((3.14159 + 1e10) / float64((0.1 + 0.1))) != (0.1 - (0.1 * 1.5))) {
                fmt.Println(v7, v7)
                var v15 bool = false
                v16 := GenS1{F0: 2147483647}
                _ = v15
                _ = v16
            }
            var v17 int8 = v7
            _ = v17
        }
        if ((1.5 + (1e10 - 0.0)) < float64((0.1 * -2.25))) {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint = 2
                var v20 int8 = 0
                v7 += v7
                _ = fi18
                _ = v19
                _ = v20
            }
        }
    }
    _ = v6
    _ = v7
    return (v6 * v6)
}
func genFn1(p0 uint16, p1 string, p2 uint16) uint32 {
    fmt.Println(p0)
    var v21 float32 = 1.5
    var v22 uint = 590
    v23 := GenS2{F0: -1}
    v24 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 2
}

func main() {
    var v1 string = "hello"
    var v2 float32 = 3.14159
    var v3 int = 127
    var v4 uint64 = 18446744073709551615
    v5 := make(map[uint32]string)
    if (("\n\t" + "\n\t") <= "\n\t") {
        if false {
            v3 += -528
            for fi25 := 0; fi25 < 6; fi25++ {
                if true {
                    var v26 uint16 = 0
                    var v27 int16 = -1
                    v28 := GenS2{F0: v27}
                    _ = v26
                    _ = v27
                    _ = v28
                }
                v3 = v3
                _ = fi25
            }
        } else {
            var v29 int16 = 149
            for fi30 := 0; fi30 < 16; fi30++ {
                v29 -= v29
                fmt.Println(v2, v29, v3)
                _ = fi30
            }
            v31 := GenS3{F0: 18446744073709551615}
            _ = v29
            _ = v31
        }
        v32 := GenS3{F0: v4}
        _ = v32
    }
    v33 := GenS4{F0: (v4 - 42), F1: ""}
    v34 := []int{}
    fmt.Println(v4, v1)
    switch v3 {
    case -624:
        var v35 string = v1
        _ = v35
    case -1:
        v1 = ("foo bar" + v1)
    case 128:
        var v36 int32 = -80
        _ = v36
    default:
        v37 := GenS5{F0: 888}
        _ = v37
    }
    if (!(!true)) {
        v38 := GenS6{F0: 4, F1: 619, F2: 0}
        switch v3 {
        case -9223372036854775808:
            fmt.Println(v1, v4, v2)
        case 9223372036854775807:
            if ((v4 << ((823) & 63)) > (v4 >> ((1) & 63))) {
                if true {
                    var v39 int16 = -1
                    var v40 float32 = ((v2 + 1.5) * v2)
                    _ = v39
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 9; fi41++ {
                        fmt.Println(v1, v4)
                        for fi42 := 0; fi42 < 3; fi42++ {
                            fmt.Println(fi41, v4, v4)
                            var v43 uint8 = 178
                            _ = fi42
                            _ = v43
                        }
                        _ = fi41
                    }
                    var v44 bool = (!false)
                    for fi45 := 0; fi45 < 6; fi45++ {
                        for fi46 := 0; fi46 < 19; fi46++ {
                            switch ("hello" + v1) {
                            case "hello":
                                v2 += ((1.5 + (float32(0.1) * 0.0)) - (1.5 + -2.25))
                            case "foo bar":
                                if (!(("" == v1) && true)) {
                                    var v47 int32 = -1
                                    v2 += (1.5 - v2)
                                    var v48 int = -445
                                    if (v4 <= v4) {
                                        v1 = (v1 + ("foo bar" + v1))
                                        if (("" + (v1 + v1)) < "\n\t") {
                                            v2 += 1.5
                                        }
                                        v4 = (v4 * 18446744073709551615)
                                    } else {
                                        if (!true) {
                                            if v44 {
                                                v34 = append(v34, 9223372036854775807)
                                            } else {
                                                v49 := v5[509]
                                                _ = v49
                                            }
                                            _ = len(v34)
                                        }
                                    }
                                    v34 = append(v34, (v48 + fi45))
                                    _ = v47
                                    _ = v48
                                }
                            case "\n\t":
                                v50 := GenS3{F0: v4}
                                _ = v50
                            }
                            var v51 uint16 = 2
                            if false {
                                var v52 bool = (!v44)
                                v1 = v1
                                if (v52 && ((v4 % ((18446744073709551615) | 1)) > 18446744073709551615)) {
                                    if (-1 < -64) {
                                        v53 := GenS4{F0: v4, F1: "hello"}
                                        _ = v53
                                    }
                                    fmt.Println(v4, fi46)
                                } else {
                                    delete(v5, uint32(v51))
                                    if len(v34) >= 2 {
                                        v34 = v34[1:len(v34)]
                                    }
                                }
                                _ = v52
                            } else {
                                _ = len(v34)
                            }
                            _ = fi46
                            _ = v51
                        }
                        var v54 bool = (!v44)
                        _ = fi45
                        _ = v54
                    }
                    _ = v44
                }
            }
        case -72:
            v2 += v2
        }
        _ = v38
    } else {
        fmt.Println(v3, v2, v3)
        if (!(false && (!false))) {
            var v55 uint16 = 185
            for fi56 := 0; fi56 < 9; fi56++ {
                v57 := GenS7{F0: ((v1 + v1) + v1), F1: 9223372036854775807}
                fmt.Println(v2)
                v58 := GenS6{F0: uint8(v55), F1: 42, F2: -9223372036854775808}
                _ = fi56
                _ = v57
                _ = v58
            }
            _ = v55
        }
        v59 := GenS8{F0: 2}
        _ = v59
    }
    if true {
        for fi60 := 0; fi60 < 20; fi60++ {
            fmt.Println(v3, v3, v2)
            _ = fi60
        }
    } else {
        fmt.Println(v4)
        v61 := v5[2]
        _ = v61
    }
    if ((v1 >= "a") && (0 == 0)) {
        v62 := GenS5{F0: -9223372036854775808}
        v63 := GenS6{F0: 33, F1: 173, F2: -870}
        if len(v34) > 0 {
            v34[0] = v3
        }
        _ = v62
        _ = v63
    } else {
        for fi64 := 0; fi64 < 20; fi64++ {
            var v65 float32 = 1e10
            for fi66 := 0; fi66 < 1; fi66++ {
                var v67 float64 = 1e10
                _ = fi66
                _ = v67
            }
            fmt.Println(fi64, v4)
            _ = fi64
            _ = v65
        }
    }
    v68 := GenS9{F0: 0, F1: (v3 & v3)}
    for fi69 := 0; fi69 < 4; fi69++ {
        fmt.Println(fi69, v1)
        for fi70 := 0; fi70 < 2; fi70++ {
            var v71 float32 = v2
            v72 := GenS2{F0: -1}
            v4 += (v4 * (^(v4)))
            v4 = 162
            _ = fi70
            _ = v71
            _ = v72
        }
        _ = fi69
    }
    switch v3 {
    case -1:
        var v73 int8 = 127
        _ = v73
    case 42:
        for fi74 := 0; fi74 < 10; fi74++ {
            var v75 uint32 = 283
            if true {
                var v76 int64 = 1
                v77 := GenS1{F0: v3}
                _ = v76
                _ = v77
            } else {
                v1 = ("" + v1)
                v75 = v75
                v1 = v1
            }
            _ = fi74
            _ = v75
        }
    case 9223372036854775807:
        var v78 uint32 = 563
        _ = v78
    }
    switch (v3 >> ((42) & 63)) {
    case 42:
        v1 = ("hello" + "")
    }
    if (genFn0() == (v4 & (v4 % ((18446744073709551615) | 1)))) {
        var v79 uint = 42
        _ = v79
    }
    var v80 int8 = -97
    var v81 uint16 = 42
    var v82 int32 = 916
    v83 := GenS8{F0: 42}
    v84 := GenS10{F0: (v80 - -34), F1: (v4 - (^(v4))), F2: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v68
    _ = v80
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
