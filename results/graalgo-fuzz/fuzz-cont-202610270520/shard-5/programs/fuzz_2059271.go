package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 string; F1 uint8; F2 int8 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int; F1 bool }

func main() {
    var v1 float32 = 0.1
    var v2 float32 = (((0.1 * (1.5 - 0.0)) / v1) / (v1 - v1))
    v3 := make(map[int]string)
    var v4 uint = 2
    v5 := make(map[uint32]string)
    var v6 int64 = 42
    var v7 int = 4294967295
    var v8 string = "\n\t"
    if true {
        fmt.Println(v8)
    } else {
        v9 := []int{v7, v7, -9223372036854775808, (v7 ^ 32767)}
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                v2 += 3.14159
                _ = fi10
            }
            var v11 int32 = 1
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint16 = 808
                delete(v3, v7)
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            for fi14 := 0; fi14 < 7; fi14++ {
                if true {
                    var v15 uint = (v4 + 132)
                    fmt.Println(v15, v8, v8)
                    fmt.Println(v7, fi14, v8)
                    _ = v15
                } else {
                    var v16 int16 = 32767
                    for fi17 := 0; fi17 < 4; fi17++ {
                        var v18 float64 = 0.0
                        _ = fi17
                        _ = v18
                    }
                    if ((1.5 - -2.25) < 1.5) {
                        v5[155] = ""
                        if false {
                            switch v7 {
                            case 1:
                                for fi19 := 0; fi19 < 1; fi19++ {
                                    fmt.Println(v1)
                                    _ = fi19
                                }
                            case 414:
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    v9 = append(v9, 9223372036854775807)
                                    fmt.Println(v4)
                                    _ = fi20
                                }
                            case 2147483647:
                                fmt.Println(v4, v4)
                            }
                            if (false && true) {
                                if true {
                                    switch ((v8 + (v8 + (((v8 + v8) + ("a" + ("foo bar" + (v8 + v8)))) + v8))) + "a") {
                                    case "\n\t":
                                        if true {
                                            if false {
                                                v2 = float32(-2.25)
                                                var v21 int = v7
                                                for fi22 := 0; fi22 < 19; fi22++ {
                                                    var v23 uint8 = 57
                                                    if (false || true) {
                                                        if false {
                                                            fmt.Println(v7, v7, v16)
                                                        } else {
                                                            if (16 >= 100) {
                                                                v21 = 2147483647
                                                                v4 = (v4 * 1)
                                                                var v24 uint32 = 42
                                                                v16 = (v16 % (((v16 & v16)) | 1))
                                                                _ = v24
                                                            } else {
                                                                v1 += v1
                                                                switch (("" + "hello") + v8) {
                                                                case "\n\t":
                                                                    var v25 int8 = -1
                                                                    _ = v25
                                                                case "hello":
                                                                    v6 += v6
                                                                case "foo bar":
                                                                    v26 := GenS1{F0: 2147483647, F1: true}
                                                                    v7 += v21
                                                                    _ = v26
                                                                }
                                                            }
                                                            v27 := GenS2{F0: ("foo bar" + v8), F1: 241, F2: -1}
                                                            if true {
                                                                v23 = 255
                                                            }
                                                            _ = v27
                                                        }
                                                    } else {
                                                        v28 := GenS2{F0: v8, F1: uint8(v4), F2: 0}
                                                        v9 = append(v9, (fi22 & 9223372036854775807))
                                                        var v29 float64 = float64(float32(float64((v1 - float32(3.14159)))))
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                    _ = fi22
                                                    _ = v23
                                                }
                                                _ = v21
                                            } else {
                                                fmt.Println(v1, v16)
                                            }
                                        } else {
                                            if true {
                                                if false {
                                                    if (2147483647 >= 2147483647) {
                                                        if (true && (65535 == 4294967295)) {
                                                            v30 := GenS1{F0: 120, F1: ((1e10 == (0.1 * -2.25)) && true)}
                                                            _ = v30
                                                        }
                                                    }
                                                    _ = len(v9)
                                                }
                                                fmt.Println(v1, v7)
                                                v31 := GenS1{F0: 0, F1: (int(v16) == v7)}
                                                _ = v31
                                            } else {
                                                for fi32 := 0; fi32 < 2; fi32++ {
                                                    v9 = append(v9, fi32)
                                                    v3[v7] = (((("a" + ("\n\t" + (v8 + ("a" + "hello")))) + ((v8 + "a") + "\n\t")) + v8) + ("foo bar" + v8))
                                                    var v33 bool = true
                                                    _ = fi32
                                                    _ = v33
                                                }
                                            }
                                            var v34 uint8 = 0
                                            if true {
                                                var v35 int16 = v16
                                                switch (-(v7)) {
                                                case 1:
                                                    if (!((v6 >> (((v4 << (((v4 * v4)) & 63))) & 63)) != v6)) {
                                                        if len(v9) > 0 {
                                                            v9[0] = 837
                                                        }
                                                        switch (v7 | (v7 - int(v6))) {
                                                        case 972:
                                                            v36 := GenS3{F0: 1, F1: (v4 << ((0) & 63))}
                                                            _ = v36
                                                        }
                                                        for fi37 := 0; fi37 < 6; fi37++ {
                                                            v6 += v6
                                                            v38 := GenS4{F0: -17}
                                                            v9 = append(v9, 9223372036854775807)
                                                            _ = fi37
                                                            _ = v38
                                                        }
                                                    } else {
                                                        var v39 uint16 = 2
                                                        v8 = (("\n\t" + v8) + "hello")
                                                        for fi40 := 0; fi40 < 17; fi40++ {
                                                            var v41 float64 = (float64(v2) + -2.25)
                                                            fmt.Println(v16)
                                                            var v42 uint32 = 503
                                                            _ = fi40
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        _ = v39
                                                    }
                                                default:
                                                    switch v7 {
                                                    case 1:
                                                        v43 := GenS1{F0: 2147483647, F1: ((!false) && (v6 <= (v6 - -909)))}
                                                        _ = v43
                                                    case -384:
                                                        fmt.Println(v6, v1)
                                                    default:
                                                        if len(v9) > 0 {
                                                            v9[0] = fi14
                                                        }
                                                    }
                                                }
                                                var v44 uint = v4
                                                _ = v35
                                                _ = v44
                                            }
                                            _ = v34
                                        }
                                    case "foo bar":
                                        for fi45 := 0; fi45 < 8; fi45++ {
                                            var v46 int64 = (^(v6))
                                            v47 := GenS2{F0: (v8 + (v8 + v8)), F1: 1, F2: -47}
                                            _ = fi45
                                            _ = v46
                                            _ = v47
                                        }
                                    case "":
                                        v2 += float32((3.14159 - (((0.1 / 1e10) + (-2.25 + 1.5)) * 0.1)))
                                    }
                                    fmt.Println(v4, v16)
                                    if (fi14 > 28) {
                                        var v48 string = "foo bar"
                                        v49 := GenS5{F0: 0, F1: (v7 & (v7 / ((v7) | 1)))}
                                        _ = len(v9)
                                        _ = v48
                                        _ = v49
                                    } else {
                                        v8 = ((((("foo bar" + v8) + (v8 + "")) + v8) + (v8 + v8)) + "\n\t")
                                        v50 := GenS2{F0: "hello", F1: 42, F2: 127}
                                        _ = v50
                                    }
                                }
                                v16 += 0
                            }
                        }
                    } else {
                        if len(v9) > 0 {
                            v9[0] = fi14
                        }
                    }
                    _ = v16
                }
                v7 += (v7 + (fi14 & (fi14 ^ (v7 ^ (fi14 % ((fi14) | 1))))))
                var v51 string = ((v8 + "hello") + (((v8 + "") + "\n\t") + v8))
                _ = fi14
                _ = v51
            }
            fmt.Println(v4, v4, v1)
        }
        v52 := GenS6{F0: 0, F1: 1}
        _ = v52
        _ = v9
    }
    delete(v3, v7)
    for fi53 := 0; fi53 < 8; fi53++ {
        v54 := []int{467, (-(fi53)), fi53}
        for fi55 := 0; fi55 < 19; fi55++ {
            v6 = 1
            v56 := GenS7{F0: 2147483648, F1: true}
            _ = fi55
            _ = v56
        }
        switch (v8 + v8) {
        case "hello":
            for fi57 := 0; fi57 < 11; fi57++ {
                if false {
                    v54 = append(v54, -1)
                    v8 = "a"
                } else {
                    if (0 <= 489) {
                        if len(v54) >= 2 {
                            v54 = v54[1:len(v54)]
                        }
                        v58 := v5[42]
                        v58 = "hello"
                        fmt.Println(v8)
                        fmt.Println(v1, fi57, fi57)
                        _ = v58
                    }
                    var v59 string = "a"
                    _ = v59
                }
                v60 := GenS8{F0: (v7 | (^(fi57))), F1: false}
                var v61 int64 = v6
                if false {
                    var v62 uint32 = 2
                    if true {
                        v2 = ((v2 / (v1 + ((v2 + v2) / v2))) - 3.14159)
                        v2 -= (((1.5 - 1e10) * v2) / (3.14159 - 1.5))
                        var v63 uint64 = 0
                        _ = v63
                    } else {
                        var v64 int = (-(fi53))
                        var v65 int8 = 1
                        v54 = append(v54, int(v4))
                        fmt.Println(v64, v1, v65)
                        _ = v64
                        _ = v65
                    }
                    _ = v62
                } else {
                    var v66 int64 = (v61 & (v6 >> ((v4) & 63)))
                    if len(v54) >= 2 {
                        v54 = v54[1:len(v54)]
                    }
                    v61 = (v61 << (((^(v4))) & 63))
                    _ = v66
                }
                _ = fi57
                _ = v60
                _ = v61
            }
        default:
            v8 = v8
        }
        _ = fi53
        _ = v54
    }
    _ = len(v3)
    v6 = 1
    v7 += 2147483647
    var v67 float32 = (v1 - 1.5)
    v1 = 0.0
    v68 := GenS4{F0: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
    _ = v68
    _ = v7
    _ = v8
}
