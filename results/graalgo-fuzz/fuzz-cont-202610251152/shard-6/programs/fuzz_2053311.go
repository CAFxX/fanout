package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 uint32; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 int8; F1 int8; F2 int32 }
func genFn0(p0 int32) uint8 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 float64 = (float64(1.5) * 0.1)
        var v9 int64 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = 32767
    _ = p0
    _ = v10
    return 0
}
func genFn1() bool {
    v11 := []int{2147483647, -905, -9223372036854775808}
    v12 := GenS1{F0: 0}
    v13 := GenS2{F0: (898 == -688)}
    _ = v11
    _ = v12
    _ = v13
    return (((!(42 == -9223372036854775808)) && (!(2147483647 > 1))) && (4294967295 <= 0))
}

func main() {
    v1 := []string{("\n\t" + (((("a" + "\n\t") + "") + "") + "foo bar")), "a", "hello", "foo bar"}
    var v2 int = 1
    var v3 string = "foo bar"
    v4 := make(map[int]string)
    var v5 int16 = 1
    var v6 int8 = 127
    var v14 int8 = 42
    if true {
        var v15 uint64 = 313
        v16 := v4[-1]
        if ((1e10 != 3.14159) || (1000 <= v5)) {
            for fi17 := 0; fi17 < 11; fi17++ {
                for fi18 := 0; fi18 < 18; fi18++ {
                    if true {
                        switch (fi17 << ((303) & 63)) {
                        case -1:
                            var v19 uint = 18446744073709551615
                            _ = v19
                        default:
                            v15 -= (v15 | v15)
                        }
                        v5 += v5
                    }
                    switch (v3 + ("a" + v3)) {
                    case "":
                        v16 = v16
                    default:
                        v20 := GenS1{F0: 255}
                        _ = v20
                    }
                    v3 = (v16 + ("foo bar" + v16))
                    _ = fi18
                }
                var v21 uint = 2
                _ = fi17
                _ = v21
            }
            v15 = v15
            for fi22 := 0; fi22 < 15; fi22++ {
                v15 -= 18446744073709551615
                switch (fi22 | (v2 | int(v5))) {
                case 32767:
                    switch (fi22 * fi22) {
                    case -1:
                        if (602 > -885) {
                            var v23 uint8 = 76
                            if true {
                                var v24 uint32 = 2147483647
                                _ = v24
                            } else {
                                v6 += (v14 << ((1) & 63))
                            }
                            _ = v23
                        }
                    case 128:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 1:
                        switch v16 {
                        case "":
                            var v25 bool = false
                            _ = v25
                        case "foo bar":
                            for fi26 := 0; fi26 < 18; fi26++ {
                                var v27 uint32 = 0
                                _ = fi26
                                _ = v27
                            }
                        case "a":
                            if true {
                                delete(v4, fi22)
                                v6 = -56
                            }
                        }
                    }
                default:
                    fmt.Println(v5, v15, v2)
                }
                _ = fi22
            }
        }
        _ = v15
        _ = v16
    } else {
        v28 := GenS1{F0: 42}
        v29 := v4[9223372036854775807]
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        fmt.Println(v14, v14, v14)
        fmt.Println(v3)
        delete(v4, v2)
        _ = fi30
    }
    if ((!(!(false && false))) && true) {
        v14 += -1
        v6 = v14
        fmt.Println(v14, v14)
    } else {
        for fi31 := 0; fi31 < 15; fi31++ {
            delete(v4, fi31)
            _ = fi31
        }
        if (false || (0.0 == (0.0 / 1e10))) {
            for fi32 := 0; fi32 < 6; fi32++ {
                _ = len(v1)
                _ = fi32
            }
            if (514 == 862) {
                var v33 uint64 = 826
                _ = v33
            } else {
                fmt.Println(v2, v2, v2)
                fmt.Println(v14, v14, v3)
            }
        }
        v34 := GenS3{F0: (v5 << ((0) & 63)), F1: 1, F2: 4294967295}
        _ = v34
    }
    v35 := GenS4{F0: 165}
    switch (v3 + v3) {
    case "foo bar":
        v6 = -18
    }
    for fi36 := 0; fi36 < 7; fi36++ {
        if false {
            _ = len(v4)
        } else {
            v37 := GenS5{F0: -1, F1: 1}
            _ = v37
        }
        v38 := GenS3{F0: (v5 - (v5 ^ 1)), F1: 94, F2: 469}
        _ = fi36
        _ = v38
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (42 == 255) {
        v14 = (v6 ^ v14)
        v5 -= v5
    } else {
        fmt.Println(v3, v3, v6)
    }
    v2 = (v2 / ((int(v5)) | 1))
    if false {
        v4[int(v5)] = ""
        var v39 uint64 = 867
        for fi40 := 0; fi40 < 2; fi40++ {
            v1 = append(v1, ((v3 + v3) + (v3 + v3)))
            switch v3 {
            case "foo bar":
                if (!true) {
                    switch ("foo bar" + v3) {
                    case "a":
                        if true {
                            if ((0.1 * (3.14159 - 0.0)) >= 3.14159) {
                                v2 = fi40
                                v41 := GenS1{F0: genFn0(-572)}
                                switch v3 {
                                case "hello":
                                    for fi42 := 0; fi42 < 20; fi42++ {
                                        switch int(v14) {
                                        case 127:
                                            var v43 uint = 0
                                            _ = v43
                                        default:
                                            v3 = v3
                                        }
                                        v44 := GenS5{F0: 1, F1: -572}
                                        if false {
                                            v45 := GenS6{F0: v14, F1: v6, F2: -641}
                                            _ = v45
                                        } else {
                                            v1 = append(v1, ("foo bar" + v3))
                                            v39 = (^(v39))
                                            v5 = (v5 / (((v5 >> ((65535) & 63))) | 1))
                                        }
                                        _ = fi42
                                        _ = v44
                                    }
                                case "":
                                    v46 := v4[v2]
                                    _ = v46
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = v41
                            } else {
                                var v47 int = v2
                                switch v2 {
                                case -2147483648:
                                    v48 := GenS4{F0: 255}
                                    _ = v48
                                default:
                                    v47 = v47
                                }
                                _ = v47
                            }
                            fmt.Println(fi40, v5)
                            v14 = (^(v6))
                        } else {
                            var v49 int16 = v5
                            _ = len(v1)
                            v3 = "foo bar"
                            _ = v49
                        }
                    case "hello":
                        var v50 int32 = 1
                        _ = v50
                    default:
                        var v51 int32 = -1
                        _ = v51
                    }
                    var v52 int64 = 0
                    var v53 uint64 = 563
                    _ = v52
                    _ = v53
                } else {
                    v6 = int8(v39)
                    v3 = v3
                }
            }
            _ = fi40
        }
        _ = v39
    } else {
        if true {
            var v54 uint16 = 625
            _ = v54
        }
        if ((v3 + (v3 + (((v3 + ((v3 + v3) + "")) + v3) + "\n\t"))) <= v3) {
            var v55 int16 = -32768
            var v56 int32 = -418
            v5 -= (v5 >> ((0) & 63))
            _ = v55
            _ = v56
        }
        if (!(98 > 2147483647)) {
            switch v3 {
            case "\n\t":
                fmt.Println(v6, v6)
            case "":
                v57 := GenS2{F0: true}
                _ = v57
            case "foo bar":
                var v58 string = ""
                _ = v58
            default:
                v5 = v5
            }
        } else {
            v5 = 42
            var v59 uint16 = 2
            _ = v59
        }
    }
    v60 := GenS3{F0: v5, F1: 1, F2: 18446744073709551615}
    var v61 uint64 = 736
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
