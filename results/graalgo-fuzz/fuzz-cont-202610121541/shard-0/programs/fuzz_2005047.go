package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint64; F1 string; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int8; F1 int16 }
func genFn0(p0 int) int {
    v5 := GenS1{F0: 1, F1: 42, F2: 42}
    if ((!(float32((1e10 - (((3.14159 - 1.5) * (3.14159 * 0.0)) / ((-2.25 * 1e10) / ((((1e10 + 0.1) - 1.5) - 0.1) + float64(-2.25)))))) == (float32(3.14159) * 0.0))) && (false || (89 < 127))) {
        p0 = p0
    } else {
        var v6 uint64 = 2
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 += fi7
        p0 -= fi7
        _ = fi7
    }
    var v8 int = 42
    var v9 int8 = -128
    _ = p0
    _ = v5
    _ = v8
    _ = v9
    return (p0 + -9223372036854775808)
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int32 = -2147483648
    v3 := make(map[string]bool)
    v4 := []uint32{615, 42}
    var v10 int64 = v1
    var v11 int16 = -1
    if (v2 > (v2 - 2147483647)) {
        var v12 uint = 0
        _ = v12
    } else {
        if (genFn0(431) != -470) {
            var v13 int64 = (^(v10))
            var v14 uint64 = 136
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 3; fi15++ {
            if (float32((1.5 - 1.5)) > (0.1 * 3.14159)) {
                if (v11 < 1) {
                    delete(v3, "")
                    if true {
                        for fi16 := 0; fi16 < 17; fi16++ {
                            var v17 int32 = 666
                            if true {
                                var v18 int16 = -32768
                                var v19 int = fi15
                                _ = len(v4)
                                _ = v18
                                _ = v19
                            } else {
                                if true {
                                    var v20 uint8 = 1
                                    v2 -= (v17 / ((v2) | 1))
                                    _ = v20
                                }
                                _ = len(v3)
                            }
                            delete(v3, "hello")
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        var v21 int16 = v11
                        _ = v21
                    }
                }
                v1 -= 42
            } else {
                v22 := GenS1{F0: 0, F1: 703, F2: 234}
                _ = v22
            }
            _ = fi15
        }
        v23 := GenS1{F0: 2, F1: 42, F2: 42}
        _ = v23
    }
    v10 = (v1 ^ (v1 / (((v10 & -9223372036854775808)) | 1)))
    v11 += (v11 - (v11 | v11))
    _ = len(v4)
    var v24 int16 = (v11 | v11)
    v25 := GenS1{F0: 860, F1: 61, F2: 255}
    v1 -= -131
    v3[("hello" + "\n\t")] = true
    for fi26 := 0; fi26 < 18; fi26++ {
        if true {
            var v27 int64 = (v10 << ((42) & 63))
            _ = v27
        } else {
            if (false && true) {
                for fi28 := 0; fi28 < 15; fi28++ {
                    var v29 uint32 = 0
                    _ = fi28
                    _ = v29
                }
                for fi30 := 0; fi30 < 19; fi30++ {
                    v10 = (v10 % (((-(v10))) | 1))
                    v31 := GenS2{F0: -9223372036854775808}
                    v2 = v2
                    _ = fi30
                    _ = v31
                }
            } else {
                v32 := GenS3{F0: 789, F1: (v11 + v24)}
                _ = v32
            }
            v24 = v24
        }
        _ = fi26
    }
    fmt.Println(v11, v2, v11)
    v2 = v2
    fmt.Println(v1, v24)
    _ = len(v4)
    if ((false && false) || ((!false) && false)) {
        var v33 uint32 = 2
        for fi34 := 0; fi34 < 6; fi34++ {
            var v35 uint = 255
            for fi36 := 0; fi36 < 6; fi36++ {
                v33 -= (^(v33))
                v24 = (^(v11))
                v37 := GenS3{F0: v35, F1: 42}
                switch 1 {
                case 562:
                    _ = len(v3)
                case 10:
                    switch "" {
                    case "\n\t":
                        v24 += v24
                    case "a":
                        var v38 uint = 772
                        _ = v38
                    case "foo bar":
                        if ((true && false) || (0 <= 2)) {
                            var v39 uint = 432
                            v2 -= (v2 | v2)
                            fmt.Println(v33, v24)
                            v1 = int64(v33)
                            switch (fi34 + fi34) {
                            case 32767:
                                v1 = 48
                            case 490:
                                fmt.Println(v39, v35, fi36)
                            }
                            _ = v39
                        } else {
                            if ("hello" <= "a") {
                                v40 := GenS4{F0: (v11 + (v11 & v11)), F1: v33}
                                var v41 uint = (v35 * 1)
                                switch "" {
                                case "a":
                                    var v42 int32 = v2
                                    _ = v42
                                }
                                var v43 uint8 = 42
                                _ = v40
                                _ = v41
                                _ = v43
                            } else {
                                v44 := GenS2{F0: (fi36 % ((27) | 1))}
                                var v45 int32 = int32(v24)
                                _ = v44
                                _ = v45
                            }
                            var v46 int = (fi34 - (fi36 - (fi36 | 127)))
                            v46 -= fi34
                            v47 := GenS5{F0: 18446744073709551615, F1: ("a" + "foo bar"), F2: v46}
                            _ = v46
                            _ = v47
                        }
                    default:
                        fmt.Println(v35, v2, v35)
                    }
                    if true {
                        var v48 uint8 = 255
                        var v49 uint32 = v33
                        _ = v48
                        _ = v49
                    } else {
                        if ((42 > -128) && (!true)) {
                            var v50 int8 = 0
                            _ = v50
                        } else {
                            switch (((("" + ("a" + "foo bar")) + ("foo bar" + ("hello" + ("a" + "a")))) + ("a" + "")) + "") {
                            case "foo bar":
                                v51 := GenS6{F0: v33}
                                v11 = 1000
                                _ = v51
                            case "a":
                                var v52 uint8 = 255
                                _ = v52
                            case "\n\t":
                                if true {
                                    if true {
                                        fmt.Println(v11, v33)
                                        v53 := GenS1{F0: 1, F1: 0, F2: 127}
                                        _ = v53
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = v33
                                        }
                                        v3[(("hello" + "\n\t") + ("" + "\n\t"))] = false
                                        v54 := GenS7{F0: 2147483647}
                                        _ = v54
                                    }
                                    v55 := GenS8{F0: -128, F1: v24}
                                    if false {
                                        fmt.Println(v24, v2, fi36)
                                    } else {
                                        var v56 uint64 = 586
                                        v3[("foo bar" + "\n\t")] = true
                                        v57 := GenS8{F0: 127, F1: v11}
                                        _ = v56
                                        _ = v57
                                    }
                                    _ = v55
                                } else {
                                    var v58 float32 = 1.5
                                    switch (fi36 >> ((255) & 63)) {
                                    case 0:
                                        v35 += v35
                                    case 32767:
                                        switch genFn0(1) {
                                        case 128:
                                            if ((127 != 2) && true) {
                                                fmt.Println(v33)
                                                _ = len(v3)
                                            } else {
                                                v33 = v33
                                                var v59 bool = false
                                                _ = v59
                                            }
                                        case -2147483648:
                                            if true {
                                                var v60 string = ""
                                                v61 := GenS6{F0: 0}
                                                v11 += 42
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v62 := GenS3{F0: v35, F1: v24}
                                                fmt.Println(v10, v24, v10)
                                                _ = v62
                                            }
                                        case 199:
                                            v58 = v58
                                        }
                                    default:
                                        var v63 uint64 = 0
                                        _ = v63
                                    }
                                    switch ("a" + "hello") {
                                    case "\n\t":
                                        v64 := GenS3{F0: v35, F1: v24}
                                        _ = v64
                                    case "":
                                        v10 += v10
                                    default:
                                        var v65 int64 = 780
                                        _ = v65
                                    }
                                    _ = v58
                                }
                            }
                            _ = len(v4)
                        }
                        v66 := v3["\n\t"]
                        _ = v66
                    }
                }
                _ = fi36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        v11 = v24
        _ = v33
    }
    var v67 uint8 = 0
    v67 = v67
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v67
}
