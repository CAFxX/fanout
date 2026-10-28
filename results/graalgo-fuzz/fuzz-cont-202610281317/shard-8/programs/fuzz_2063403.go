package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint8; F1 bool; F2 int16 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int8 }
func genFn0(p0 bool, p1 int16) uint16 {
    p0 = ((p0 && ((-135 >= 900) || p0)) && p0)
    p1 += p1
    _ = p0
    _ = p1
    return 65535
}
func genFn1() string {
    var v7 uint = 548
    v7 = v7
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            var v9 uint = v7
            _ = v9
        }
        var v10 string = (("hello" + (("\n\t" + "foo bar") + ("" + ""))) + ("a" + ""))
        _ = fi8
        _ = v10
    }
    v11 := GenS1{F0: 436, F1: "foo bar"}
    fmt.Println(v7)
    _ = v11
    _ = v7
    return ("\n\t" + ("" + "\n\t"))
}

func main() {
    v1 := []string{}
    var v2 float32 = (0.0 / 1.5)
    var v3 uint32 = 4294967295
    var v4 bool = (42 == 2147483647)
    v5 := make(map[string]uint32)
    var v6 string = "hello"
    for fi12 := 0; fi12 < 2; fi12++ {
        if (0 > 93) {
            var v13 bool = v4
            _ = v13
        } else {
            switch "foo bar" {
            case "foo bar":
                for fi14 := 0; fi14 < 10; fi14++ {
                    if len(v1) > 0 {
                        v1[0] = v6
                    }
                    v5["foo bar"] = v3
                    fmt.Println(v4)
                    _ = fi14
                }
            }
            if (v4 && (!(v4 || v4))) {
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint16 = 616
                    v17 := GenS1{F0: v16, F1: ("" + (v6 + (v6 + v6)))}
                    var v18 int8 = -1
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                var v19 bool = v4
                v2 += (v2 - 3.14159)
                _ = v19
            } else {
                switch fi12 {
                case 2147483647:
                    var v20 float32 = (float32((float64(v2) * 3.14159)) + v2)
                    _ = v20
                default:
                    var v21 uint64 = 506
                    _ = v21
                }
                if ((0 != 2) || v4) {
                    var v22 uint32 = (v3 >> ((1) & 63))
                    _ = v22
                } else {
                    v23 := GenS1{F0: genFn0(v4, -539), F1: ""}
                    _ = len(v1)
                    _ = v23
                }
            }
        }
        _ = fi12
    }
    if (1 != 1000) {
        for fi24 := 0; fi24 < 1; fi24++ {
            v1 = append(v1, v6)
            _ = fi24
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v4)
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v4 = false
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 string = v6
        _ = fi25
        _ = v26
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v27 := GenS1{F0: genFn0(v4, 1000), F1: v6}
        v28 := GenS2{F0: 255}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int = fi29
        v31 := GenS3{F0: 1, F1: 628}
        var v32 uint8 = 42
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if false {
            if v4 {
                fmt.Println(v4, v6, fi33)
                v34 := GenS2{F0: 937}
                var v35 int32 = -1
                _ = v34
                _ = v35
            }
            fmt.Println(v4, v3, v2)
            v2 -= (3.14159 / ((((v2 - 0.1) * v2) * (v2 - ((v2 * v2) / v2))) + ((v2 / v2) - 1e10)))
        }
        if true {
            var v36 uint = 0
            switch (((v6 + "a") + ("" + v6)) + ("\n\t" + ("\n\t" + (((v6 + v6) + "hello") + v6)))) {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            switch ("" + v6) {
            case "a":
                var v37 uint8 = 67
                _ = v37
            default:
                v38 := v5[v6]
                _ = v38
            }
            _ = v36
        } else {
            if (!(32767 < -610)) {
                v1 = append(v1, v6)
            }
            switch "" {
            case "":
                v3 -= (v3 - (v3 ^ v3))
            }
            v2 += v2
        }
        v5["foo bar"] = (v3 >> ((18446744073709551615) & 63))
        _ = fi33
    }
    var v39 int16 = 1000
    if v4 {
        var v40 int64 = -9223372036854775808
        _ = v40
    } else {
        switch 4294967295 {
        case 35:
            if false {
                if (!(!(true || v4))) {
                    v4 = ((!(!v4)) || v4)
                    if ((-128 >= -128) && v4) {
                        switch 9223372036854775807 {
                        case 32768:
                            fmt.Println(v6)
                        default:
                            fmt.Println(v3)
                        }
                        if false {
                            if (1e10 == (1.5 / (1.5 + (0.1 - float64((v2 + (float32((0.0 - 1e10)) - -2.25))))))) {
                                v41 := GenS4{F0: 4, F1: (v4 && ((false && (v4 && (42 <= (v39 | 1000)))) && (float64(v2) == 0.0))), F2: (v39 & (v39 + (v39 - 0)))}
                                _ = v41
                            }
                            v3 = (v3 % (((v3 % ((1) | 1))) | 1))
                        }
                        fmt.Println(v6, v3)
                    }
                    for fi42 := 0; fi42 < 5; fi42++ {
                        var v43 uint64 = 18446744073709551615
                        _ = fi42
                        _ = v43
                    }
                } else {
                    v44 := GenS5{F0: (v6 + v6), F1: "foo bar"}
                    fmt.Println(v3, v3)
                    if (42 < 1) {
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    } else {
                        for fi46 := 0; fi46 < 9; fi46++ {
                            for fi47 := 0; fi47 < 16; fi47++ {
                                var v48 uint64 = 499
                                v5[v6] = v3
                                switch "hello" {
                                case "foo bar":
                                    if false {
                                        if ((!true) || true) {
                                            switch -2147483648 {
                                            case 0:
                                                var v49 int32 = 2147483647
                                                _ = v49
                                            }
                                            if false {
                                                fmt.Println(v3, fi46)
                                                var v50 int64 = 0
                                                _ = v50
                                            }
                                        } else {
                                            var v51 uint32 = v3
                                            _ = v51
                                        }
                                        switch fi47 {
                                        case 1:
                                            v39 += -1
                                        case -787:
                                            v4 = v4
                                        case 4294967295:
                                            v52 := v5[(v6 + v6)]
                                            _ = v52
                                        default:
                                            _ = len(v1)
                                        }
                                    } else {
                                        v39 = (v39 % ((v39) | 1))
                                        if ((!((v39 ^ v39) >= v39)) || (-32768 <= (^(v39)))) {
                                            var v53 float64 = float64(v2)
                                            fmt.Println(v2, v3, v3)
                                            _ = len(v5)
                                            _ = v53
                                        } else {
                                            if v4 {
                                                var v54 int = 150
                                                _ = v54
                                            } else {
                                                var v55 uint32 = 1
                                                _ = v55
                                            }
                                        }
                                    }
                                case "a":
                                    fmt.Println(v39, v2)
                                    v56 := GenS1{F0: genFn0(((float32(-2.25) < v2) && (90 < -128)), v39), F1: "\n\t"}
                                    v48 -= 18446744073709551615
                                    _ = v56
                                }
                                _ = fi47
                                _ = v48
                            }
                            _ = fi46
                        }
                        v57 := GenS6{F0: 127}
                        _ = v57
                    }
                    _ = v44
                }
            }
        case -9223372036854775808:
            switch (((v6 + (v6 + v6)) + v6) + "a") {
            case "":
                if (!v4) {
                    v4 = (int(v3) > -9223372036854775808)
                    if (!(969 <= 1)) {
                        fmt.Println(v6)
                    } else {
                        fmt.Println(v3, v3, v39)
                        var v58 int = -378
                        _ = v58
                    }
                } else {
                    v6 = v6
                }
            case "hello":
                var v59 int16 = v39
                _ = v59
            default:
                v3 -= 65535
            }
        case 2147483647:
            v2 -= 0.1
        default:
            switch (v6 + (v6 + v6)) {
            case "a":
                for fi60 := 0; fi60 < 2; fi60++ {
                    v61 := GenS5{F0: (v6 + (v6 + "")), F1: "\n\t"}
                    v4 = v4
                    _ = fi60
                    _ = v61
                }
            }
        }
        v39 = -1
    }
    var v62 int8 = 1
    var v63 int8 = (v62 ^ v62)
    v1 = append(v1, v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
}
