package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint8) uint {
    var v5 int64 = -9223372036854775808
    v6 := GenS1{F0: 42, F1: 9223372036854775807, F2: "foo bar"}
    _ = p0
    _ = v5
    _ = v6
    return 595
}
func genFn1(p0 int8, p1 int16) uint {
    p1 = (p1 % ((p1) | 1))
    fmt.Println(p1, p1)
    var v7 int32 = 42
    v8 := []int{127}
    for fi9 := 0; fi9 < 13; fi9++ {
        if len(v8) > 0 {
            v8[0] = (fi9 * fi9)
        }
        var v10 int = fi9
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return genFn0(124)
}

func main() {
    v1 := []int{-2147483648, 127, -17, -306}
    var v2 int64 = 9223372036854775807
    var v3 bool = true
    var v4 bool = (!v3)
    _ = len(v1)
    v11 := GenS2{F0: 18446744073709551615}
    if v3 {
        v2 = -1
        v3 = false
        var v12 int16 = 42
        _ = v12
    } else {
        if len(v1) > 0 {
            v1[0] = -703
        }
        var v13 float64 = 0.1
        var v14 string = "foo bar"
        _ = v13
        _ = v14
    }
    var v15 int16 = -757
    switch 127 {
    case -646:
        if false {
            if true {
                v16 := GenS1{F0: 0, F1: 42, F2: "a"}
                switch ("" + "") {
                case "a":
                    v3 = ((v4 && true) && v4)
                case "":
                    switch ("\n\t" + "a") {
                    case "a":
                        if true {
                            var v17 string = ("\n\t" + "\n\t")
                            if len(v1) > 0 {
                                v1[0] = 127
                            }
                            if v4 {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    if v4 {
                                        fmt.Println(v17, v15, v2)
                                        if ((0.0 >= (3.14159 - 3.14159)) && false) {
                                            fmt.Println(v4)
                                            for fi19 := 0; fi19 < 12; fi19++ {
                                                if (fi18 < (fi19 + (fi18 | (fi18 * 128)))) {
                                                    var v20 int32 = 42
                                                    switch fi18 {
                                                    case 42:
                                                        fmt.Println(v2, v20, v17)
                                                    case 127:
                                                        var v21 float32 = 0.0
                                                        _ = v21
                                                    }
                                                    _ = v20
                                                }
                                                v22 := GenS1{F0: 18446744073709551615, F1: 550, F2: "\n\t"}
                                                _ = fi19
                                                _ = v22
                                            }
                                        } else {
                                            var v23 int = fi18
                                            _ = v23
                                        }
                                    }
                                    _ = fi18
                                }
                                v1 = append(v1, -310)
                            } else {
                                if (v2 < v2) {
                                    v3 = v3
                                    var v24 uint16 = 725
                                    var v25 int64 = (v2 >> ((591) & 63))
                                    _ = v24
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: 940}
                                    _ = v26
                                }
                                var v27 int64 = (v2 * (-(v2)))
                                _ = v27
                            }
                            _ = v17
                        } else {
                            var v28 float64 = 1.5
                            if (((1.5 - (float32(v28) + (0.1 + 1.5))) + 0.0) > 3.14159) {
                                var v29 uint32 = 42
                                for fi30 := 0; fi30 < 10; fi30++ {
                                    v1 = append(v1, fi30)
                                    for fi31 := 0; fi31 < 6; fi31++ {
                                        fmt.Println(fi30, fi30, fi31)
                                        _ = fi31
                                    }
                                    _ = fi30
                                }
                                var v32 uint = genFn0(2)
                                _ = v29
                                _ = v32
                            } else {
                                v4 = (!((979 != v2) || (-50 <= (v15 + v15))))
                            }
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    for fi33 := 0; fi33 < 6; fi33++ {
                                        var v34 int8 = 100
                                        if ((v15 & (v15 ^ -664)) >= (v15 >> ((116) & 63))) {
                                            for fi35 := 0; fi35 < 8; fi35++ {
                                                v15 -= (v15 - v15)
                                                _ = fi35
                                            }
                                            fmt.Println(v2, v15)
                                        } else {
                                            for fi36 := 0; fi36 < 4; fi36++ {
                                                var v37 bool = v3
                                                if ((-(v34)) > (v34 | (v34 >> ((18446744073709551615) & 63)))) {
                                                    v38 := GenS1{F0: 125, F1: 0, F2: "hello"}
                                                    v39 := GenS2{F0: 0}
                                                    _ = v38
                                                    _ = v39
                                                }
                                                v40 := GenS2{F0: 2}
                                                _ = fi36
                                                _ = v37
                                                _ = v40
                                            }
                                            if (14 > 42) {
                                                fmt.Println(v28, v34)
                                                for fi41 := 0; fi41 < 1; fi41++ {
                                                    var v42 uint64 = 42
                                                    _ = fi41
                                                    _ = v42
                                                }
                                            }
                                        }
                                        _ = fi33
                                        _ = v34
                                    }
                                    if v3 {
                                        v15 = 42
                                        fmt.Println(v28)
                                    }
                                }
                            case "foo bar":
                                v15 = (v15 - v15)
                            case "a":
                                var v43 uint = 275
                                _ = v43
                            default:
                                fmt.Println(v28, v4)
                            }
                            _ = v28
                        }
                    case "":
                        v1 = append(v1, -2147483648)
                    case "hello":
                        v44 := GenS2{F0: 65535}
                        _ = v44
                    }
                case "hello":
                    var v45 int16 = v15
                    _ = v45
                }
                _ = v16
            } else {
                v3 = v4
            }
            switch -811 {
            case 9223372036854775807:
                for fi46 := 0; fi46 < 17; fi46++ {
                    fmt.Println(fi46, v15, v2)
                    if false {
                        v1 = append(v1, fi46)
                        v47 := GenS3{F0: v15}
                        var v48 float64 = 3.14159
                        _ = v47
                        _ = v48
                    }
                    switch fi46 {
                    case 32768:
                        for fi49 := 0; fi49 < 18; fi49++ {
                            var v50 int = fi49
                            var v51 bool = false
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    case -169:
                        var v52 uint64 = 1
                        _ = v52
                    case 289:
                        v15 = -1
                    }
                    _ = fi46
                }
            case -847:
                var v53 bool = v4
                _ = v53
            case 32767:
                v15 += 0
            }
        }
    default:
        v1 = append(v1, 189)
    }
    v54 := GenS2{F0: 1}
    var v55 uint16 = 255
    for fi56 := 0; fi56 < 12; fi56++ {
        var v57 int32 = 1
        _ = fi56
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
