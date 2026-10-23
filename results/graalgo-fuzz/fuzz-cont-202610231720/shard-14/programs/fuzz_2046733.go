package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint; F1 uint; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint64 }

func main() {
    var v1 int8 = 32
    var v2 uint8 = 0
    var v3 float64 = 0.1
    fmt.Println(v3)
    v4 := []int{959, 9223372036854775807}
    _ = len(v4)
    var v5 int64 = 1
    switch -1 {
    case 4294967295:
        v5 = v5
    default:
        if false {
            if true {
                var v6 uint64 = 405
                v7 := GenS1{F0: -2147483648, F1: 69}
                v8 := GenS2{F0: (v2 % ((uint8(v6)) | 1)), F1: 2, F2: v6}
                _ = v6
                _ = v7
                _ = v8
            } else {
                _ = len(v4)
                v3 = v3
                if false {
                    if true {
                        v4 = append(v4, 4294967295)
                        v3 = ((1.5 / -2.25) + float64((3.14159 * 3.14159)))
                        fmt.Println(v2, v2)
                    } else {
                        v3 -= (0.0 * v3)
                        var v9 uint64 = 0
                        _ = v9
                    }
                    _ = len(v4)
                } else {
                    v2 += (v2 / ((193) | 1))
                    if (!(!(false || (42 != 186)))) {
                        if false {
                            v1 -= (v1 | v1)
                            var v10 uint64 = 0
                            v11 := GenS2{F0: v2, F1: 32767, F2: v10}
                            _ = v10
                            _ = v11
                        }
                    } else {
                        var v12 uint32 = 2
                        if true {
                            for fi13 := 0; fi13 < 5; fi13++ {
                                v14 := GenS3{F0: (("hello" + ("\n\t" + ("foo bar" + "a"))) + ("hello" + "hello")), F1: (v5 | 951)}
                                v15 := GenS3{F0: "", F1: -513}
                                var v16 uint8 = 2
                                _ = fi13
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            v17 := GenS4{F0: 267, F1: 59, F2: uint8(v12)}
                            _ = v17
                        }
                        _ = v12
                    }
                    v18 := GenS1{F0: 127, F1: v2}
                    _ = v18
                }
            }
            if false {
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 int8 = v1
                    fmt.Println(v3, v3, v3)
                    v2 += (v2 << ((1) & 63))
                    _ = fi19
                    _ = v20
                }
                if true {
                    var v21 uint = 18446744073709551615
                    _ = v21
                }
            } else {
                v22 := GenS3{F0: "", F1: (v5 * (-(v5)))}
                _ = v22
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            v1 += 1
        }
    }
    v1 = (v1 + 23)
    for fi23 := 0; fi23 < 6; fi23++ {
        v1 -= -99
        v24 := GenS1{F0: -2147483648, F1: v2}
        _ = fi23
        _ = v24
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "a":
        v3 -= v3
    case "hello":
        fmt.Println(v3, v2)
    case "":
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 float64 = ((3.14159 * 0.1) * (-2.25 - v3))
            switch 4294967295 {
            case 4294967295:
                switch -922 {
                case -140:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        fmt.Println(fi27, v3)
                        _ = fi27
                    }
                case 346:
                    if (true || false) {
                        v26 = v26
                    } else {
                        var v28 uint8 = v2
                        _ = v28
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = fi25
                    }
                }
            case -9223372036854775808:
                switch (("\n\t" + (("hello" + ("" + "")) + (("foo bar" + "") + ""))) + "foo bar") {
                case "foo bar":
                    v5 = -1
                case "":
                    v29 := GenS5{F0: (-2.25 <= (((3.14159 + 0.1) / (float32(float64(1e10)) / (1e10 * (3.14159 + 1e10)))) / (0.0 + 1.5)))}
                    _ = v29
                case "\n\t":
                    v30 := GenS3{F0: "foo bar", F1: v5}
                    _ = v30
                }
            case 42:
                fmt.Println(v26, v26)
            default:
                _ = len(v4)
            }
            fmt.Println(v26, v3, fi25)
            _ = fi25
            _ = v26
        }
    default:
        var v31 int32 = 1
        _ = v31
    }
    if false {
        if true {
            var v32 int16 = -32768
            var v33 float64 = ((((0.1 - ((3.14159 - v3) + float64((1e10 / 1.5)))) + 1e10) * v3) + v3)
            v34 := GenS4{F0: 18446744073709551615, F1: 0, F2: v2}
            _ = v32
            _ = v33
            _ = v34
        } else {
            if (((("foo bar" + (("\n\t" + "") + "a")) + ("" + "hello")) + "hello") == (("a" + ("foo bar" + "")) + "hello")) {
                v5 -= v5
                var v35 uint = 923
                fmt.Println(v3)
                _ = v35
            } else {
                v3 += -2.25
                var v36 int8 = -88
                if (4294967295 >= 42) {
                    var v37 int64 = v5
                    var v38 uint16 = 316
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    _ = v37
                    _ = v38
                }
                _ = v36
            }
        }
        if true {
            _ = len(v4)
        }
        var v39 string = "foo bar"
        _ = v39
    } else {
        if true {
            for fi40 := 0; fi40 < 17; fi40++ {
                fmt.Println(fi40)
                if true {
                    for fi41 := 0; fi41 < 18; fi41++ {
                        v4 = append(v4, fi41)
                        var v42 float64 = 3.14159
                        if len(v4) > 0 {
                            v4[0] = (fi41 / ((497) | 1))
                        }
                        _ = fi41
                        _ = v42
                    }
                    if false {
                        v3 -= 0.0
                        switch ("foo bar" + "") {
                        case "foo bar":
                            var v43 uint16 = 32767
                            _ = v43
                        case "hello":
                            fmt.Println(fi40, v2)
                        default:
                            var v44 string = ((((("\n\t" + "\n\t") + "a") + ((("\n\t" + ((("\n\t" + ("hello" + (("hello" + "\n\t") + (("a" + "a") + (("hello" + "foo bar") + ""))))) + (("foo bar" + "a") + ("hello" + "foo bar"))) + "\n\t")) + "\n\t") + ("hello" + "hello"))) + "foo bar") + "\n\t")
                            _ = v44
                        }
                        fmt.Println(v5, v5)
                    } else {
                        for fi45 := 0; fi45 < 6; fi45++ {
                            var v46 int8 = v1
                            _ = fi45
                            _ = v46
                        }
                        var v47 uint16 = 0
                        var v48 float32 = float32(v3)
                        _ = v47
                        _ = v48
                    }
                    v1 -= v1
                }
                _ = fi40
            }
            for fi49 := 0; fi49 < 14; fi49++ {
                for fi50 := 0; fi50 < 1; fi50++ {
                    var v51 float32 = (1e10 - ((float32(v3) - 1e10) + (-2.25 * (1.5 - float32(v3)))))
                    _ = fi50
                    _ = v51
                }
                for fi52 := 0; fi52 < 3; fi52++ {
                    if len(v4) > 0 {
                        v4[0] = (fi49 - 785)
                    }
                    v53 := GenS6{F0: (v1 ^ v1)}
                    if (!(v2 != 133)) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v54 bool = true
                        if ((!((^(v5)) > (v5 - v5))) || false) {
                            var v55 int32 = 527
                            var v56 uint = 65535
                            var v57 int8 = (v1 + v1)
                            if v54 {
                                if (v55 < v55) {
                                    if len(v4) > 0 {
                                        v4[0] = (fi49 % ((fi52) | 1))
                                    }
                                } else {
                                    var v58 int8 = -128
                                    _ = v58
                                }
                            } else {
                                _ = len(v4)
                                v54 = (v55 != -338)
                            }
                            v1 = (v57 + (v57 & v1))
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            v59 := GenS7{F0: 1, F1: 2, F2: 4294967295}
                            var v60 bool = v54
                            _ = v59
                            _ = v60
                        }
                        _ = v54
                    }
                    v3 += (0.0 * v3)
                    _ = fi52
                    _ = v53
                }
                fmt.Println(v1, v2)
                _ = fi49
            }
            for fi61 := 0; fi61 < 6; fi61++ {
                fmt.Println(v1, v2)
                if (false || true) {
                    v5 = -773
                    v2 -= v2
                    for fi62 := 0; fi62 < 13; fi62++ {
                        v63 := GenS6{F0: v1}
                        var v64 bool = (((fi61 * 128) > fi62) && (((float32(v3) >= -2.25) && true) && true))
                        v64 = (!(v1 <= (^(v1))))
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                } else {
                    v2 -= v2
                    fmt.Println(v3)
                    v4 = append(v4, (fi61 - 2147483647))
                }
                var v65 uint32 = 0
                _ = fi61
                _ = v65
            }
        } else {
            v66 := GenS8{F0: (v5 << ((735) & 63))}
            var v67 float64 = (v3 * (v3 - 0.1))
            _ = v66
            _ = v67
        }
    }
    fmt.Println(v2, v5)
    var v68 float64 = ((v3 + v3) * v3)
    var v69 uint64 = 945
    if (-1 < (v5 + v5)) {
        for fi70 := 0; fi70 < 13; fi70++ {
            v69 += (^(v69))
            v69 = (^(v69))
            _ = fi70
        }
        var v71 int32 = 1
        _ = v71
    } else {
        if true {
            v72 := GenS1{F0: 4294967295, F1: v2}
            _ = v72
        }
        v4 = append(v4, 128)
    }
    v73 := GenS9{F0: v69, F1: v69, F2: v69}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v68
    _ = v69
    _ = v73
}
