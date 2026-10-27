package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint16; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 bool; F1 string; F2 int }

func main() {
    var v1 uint = 0
    var v2 int = -4
    var v3 bool = false
    var v4 int = v2
    var v5 float64 = 0.1
    v6 := GenS1{F0: 4294967295, F1: -601}
    v7 := []int{(v2 | (v4 >> ((v1) & 63)))}
    if len(v7) > 0 {
        v7[0] = (v2 / ((9223372036854775807) | 1))
    }
    var v8 float32 = -2.25
    var v9 float64 = 3.14159
    for fi10 := 0; fi10 < 20; fi10++ {
        if v3 {
            fmt.Println(v9, v3)
            switch (("a" + ("a" + (((("a" + "hello") + "") + "") + "\n\t"))) + "") {
            case "hello":
                v7 = append(v7, v4)
            case "":
                switch (v4 + v2) {
                case 1:
                    var v11 int8 = 127
                    _ = v11
                }
            default:
                v4 = fi10
            }
        }
        v3 = (2 != v1)
        _ = fi10
    }
    v8 -= 1.5
    if v3 {
        v9 -= v5
        fmt.Println(v2, v5, v1)
    } else {
        var v12 int32 = -205
        var v13 string = "a"
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                fmt.Println(v9, v3)
                v8 = 0.1
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
            var v15 uint8 = 1
            _ = fi14
            _ = v15
        }
        _ = v12
        _ = v13
    }
    if v3 {
        for fi16 := 0; fi16 < 10; fi16++ {
            switch v2 {
            case 419:
                for fi17 := 0; fi17 < 5; fi17++ {
                    if v3 {
                        if v3 {
                            if len(v7) > 0 {
                                v7[0] = (fi17 + fi17)
                            }
                            v2 -= v2
                        }
                    } else {
                        if (!v3) {
                            if v3 {
                                v8 = v8
                                v18 := GenS1{F0: 726, F1: -262}
                                _ = v18
                            }
                            v9 -= v5
                        } else {
                            v3 = (118 != 2)
                            v19 := GenS1{F0: 0, F1: -991}
                            fmt.Println(v3)
                            v4 = (v2 + (-(fi17)))
                            v8 = (v8 - 0.1)
                            _ = v19
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v20 float32 = v8
                        _ = v20
                    }
                    var v21 int32 = 683
                    v2 -= (-(v2))
                    var v22 int64 = 9223372036854775807
                    fmt.Println(v9)
                    _ = fi17
                    _ = v21
                    _ = v22
                }
            case 42:
                var v23 string = ((("\n\t" + "\n\t") + "") + "foo bar")
                _ = v23
            case 32767:
                v8 -= v8
            }
            _ = fi16
        }
        if v3 {
            fmt.Println(v9)
            v24 := GenS2{F0: 1, F1: 1, F2: false}
            if v3 {
                v8 = (v8 - (((-2.25 / ((((float32(v5) + 0.0) * v8) / v8) + (0.0 - v8))) - float32(v9)) + v8))
                if (!true) {
                    fmt.Println(v2, v2, v1)
                } else {
                    var v25 uint32 = 886
                    if len(v7) > 0 {
                        v7[0] = v2
                    }
                    _ = v25
                }
                v26 := GenS2{F0: 65535, F1: 1, F2: v3}
                _ = v26
            } else {
                fmt.Println(v5, v5, v4)
                switch 121 {
                case 0:
                    var v27 uint64 = 42
                    _ = v27
                case 218:
                    v8 = ((v8 * 1.5) / 1.0)
                case 127:
                    if false {
                        var v28 string = "hello"
                        for fi29 := 0; fi29 < 7; fi29++ {
                            for fi30 := 0; fi30 < 15; fi30++ {
                                var v31 float32 = v8
                                switch (v28 + (("hello" + (v28 + ((v28 + v28) + "hello"))) + ((v28 + v28) + v28))) {
                                case "hello":
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                case "a":
                                    v32 := GenS1{F0: 0, F1: 32767}
                                    _ = v32
                                }
                                _ = fi30
                                _ = v31
                            }
                            _ = fi29
                        }
                        if (((!v3) || true) && (!(!(v3 && (-2.25 < v8))))) {
                            if v3 {
                                var v33 uint32 = 65535
                                var v34 float64 = (float64(v8) - (v5 - v9))
                                _ = v33
                                _ = v34
                            }
                            var v35 uint64 = 613
                            _ = v35
                        } else {
                            var v36 uint32 = 2147483648
                            var v37 bool = ((((v5 * v9) + v5) / v9) >= v9)
                            _ = v36
                            _ = v37
                        }
                        _ = v28
                    }
                default:
                    v38 := GenS1{F0: 4294967295, F1: 1}
                    _ = v38
                }
                v3 = ((("" + ("foo bar" + "hello")) == "foo bar") && false)
            }
            _ = v24
        }
    }
    v39 := GenS1{F0: 18446744073709551615, F1: -32768}
    var v40 float32 = 1e10
    if len(v7) > 0 {
        v7[0] = (v2 + -9223372036854775808)
    }
    switch "" {
    case "foo bar":
        v9 = v9
    }
    if (!false) {
        fmt.Println(v1)
    } else {
        if v3 {
            for fi41 := 0; fi41 < 20; fi41++ {
                for fi42 := 0; fi42 < 11; fi42++ {
                    v5 += float64(3.14159)
                    _ = fi42
                }
                fmt.Println(v1)
                v3 = (!v3)
                _ = fi41
            }
            if (1 > 0) {
                if (18446744073709551615 <= 2) {
                    switch (v4 + v2) {
                    case 9223372036854775807:
                        for fi43 := 0; fi43 < 6; fi43++ {
                            v44 := GenS2{F0: uint32(v1), F1: 128, F2: ((v1 % ((v1) | 1)) < 18446744073709551615)}
                            _ = fi43
                            _ = v44
                        }
                    case 32768:
                        var v45 uint = v1
                        _ = v45
                    case 127:
                        v46 := GenS2{F0: 1, F1: 1, F2: false}
                        _ = v46
                    }
                    var v47 uint64 = 0
                    var v48 float64 = float64((v40 / v40))
                    _ = v47
                    _ = v48
                }
                fmt.Println(v3, v3, v8)
            }
            _ = len(v7)
        } else {
            if v3 {
                if (!((v1 / ((65535) | 1)) >= v1)) {
                    if (((v4 ^ 1) >= (v2 << (((v1 * v1)) & 63))) || ((1e10 / 1.5) <= 1e10)) {
                        for fi49 := 0; fi49 < 1; fi49++ {
                            v50 := GenS2{F0: 2147483648, F1: 65535, F2: (32767 == (-(fi49)))}
                            for fi51 := 0; fi51 < 12; fi51++ {
                                if false {
                                    if true {
                                        if (int(v1) < (fi49 | 9223372036854775807)) {
                                            if (((-32768 != -32768) || (4294967295 < (-(fi51)))) || false) {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                            }
                                            if len(v7) > 0 {
                                                v7[0] = v2
                                            }
                                            var v52 string = "\n\t"
                                            _ = v52
                                        }
                                        v53 := GenS1{F0: 939, F1: -1}
                                        _ = v53
                                    } else {
                                        var v54 int32 = 1
                                        var v55 float32 = v40
                                        _ = v54
                                        _ = v55
                                    }
                                } else {
                                    if len(v7) > 0 {
                                        v7[0] = fi51
                                    }
                                    v3 = (32767 != 1)
                                    v9 += v5
                                    if ((65535 < 229) || true) {
                                        if v3 {
                                            if true {
                                                if v3 {
                                                    if (42 >= fi51) {
                                                        fmt.Println(v1, v1, v40)
                                                        var v56 int16 = 304
                                                        _ = v56
                                                    }
                                                    v5 = v5
                                                    _ = len(v7)
                                                } else {
                                                    v9 += ((v5 - float64(v40)) / 1e10)
                                                }
                                                switch (-(fi51)) {
                                                case 9223372036854775807:
                                                    var v57 int16 = 0
                                                    _ = v57
                                                case 2147483647:
                                                    v40 += v8
                                                case -1:
                                                    switch ("foo bar" + ((("a" + "") + ("hello" + "\n\t")) + ("foo bar" + ("" + "foo bar")))) {
                                                    case "":
                                                        fmt.Println(fi51)
                                                    }
                                                default:
                                                    v58 := GenS2{F0: 2147483648, F1: 2, F2: (false || v3)}
                                                    _ = v58
                                                }
                                            }
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = (v4 - fi51)
                                        }
                                    } else {
                                        v9 = -2.25
                                        switch ("\n\t" + ("" + ("a" + "a"))) {
                                        case "a":
                                            switch (^(v4)) {
                                            case 9223372036854775807:
                                                v59 := GenS3{F0: 42}
                                                _ = v59
                                            default:
                                                var v60 int64 = -252
                                                _ = v60
                                            }
                                        case "foo bar":
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                        default:
                                            var v61 string = "a"
                                            _ = v61
                                        }
                                        var v62 uint64 = 18446744073709551615
                                        _ = v62
                                    }
                                    v63 := GenS4{F0: 42, F1: 32767}
                                    _ = v63
                                }
                                var v64 float64 = v5
                                if v3 {
                                    if len(v7) > 0 {
                                        v7[0] = 128
                                    }
                                    var v65 bool = v3
                                    v66 := GenS5{F0: 4294967295}
                                    _ = v65
                                    _ = v66
                                } else {
                                    var v67 int8 = -128
                                    _ = v67
                                }
                                _ = fi51
                                _ = v64
                            }
                            fmt.Println(v4)
                            _ = fi49
                            _ = v50
                        }
                        switch v2 {
                        case 4294967295:
                            fmt.Println(v40, v40, v8)
                        default:
                            v3 = ((v40 < v8) && v3)
                        }
                    }
                } else {
                    if (v3 && (!((!v3) || true))) {
                        v68 := GenS6{F0: ("foo bar" + ""), F1: 42}
                        _ = v68
                    } else {
                        var v69 float32 = 1.5
                        if len(v7) > 0 {
                            v7[0] = v4
                        }
                        _ = v69
                    }
                    var v70 string = (("foo bar" + "a") + ("a" + "hello"))
                    _ = v70
                }
                fmt.Println(v1, v8)
                v71 := GenS6{F0: "", F1: 297}
                _ = v71
            } else {
                v72 := GenS7{F0: true, F1: (("" + ("a" + "")) + "hello"), F2: (v2 / ((-327) | 1))}
                var v73 int16 = -809
                _ = v72
                _ = v73
            }
        }
    }
    v9 -= v9
    switch 128 {
    case 4294967295:
        v7 = append(v7, v2)
    case 42:
        v8 = float32(1.5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
