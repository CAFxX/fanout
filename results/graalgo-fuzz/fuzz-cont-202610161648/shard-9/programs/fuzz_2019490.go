package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 string; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint8; F1 int32; F2 int64 }
type GenS9 struct { F0 string; F1 string }
func genFn0() uint16 {
    v5 := GenS1{F0: "\n\t", F1: -32768}
    v6 := []int{-1, -2147483648}
    var v7 int8 = 127
    _ = v5
    _ = v6
    _ = v7
    return 157
}

func main() {
    var v1 uint64 = 952
    var v2 bool = false
    v3 := []string{("a" + ("a" + ("foo bar" + "\n\t"))), ("hello" + "foo bar")}
    var v4 int32 = 1
    if v2 {
        v8 := GenS2{F0: 2147483647}
        if true {
            var v9 int32 = (v4 + 2147483647)
            _ = v9
        }
        switch 32767 {
        case -139:
            for fi10 := 0; fi10 < 6; fi10++ {
                v1 = 42
                _ = fi10
            }
        case 4294967295:
            var v11 int32 = -1
            _ = v11
        }
        _ = v8
    }
    _ = len(v3)
    v12 := GenS1{F0: ((("" + ("hello" + "a")) + "") + "foo bar"), F1: 343}
    fmt.Println(v1, v4)
    v13 := GenS1{F0: "a", F1: 32767}
    v14 := GenS1{F0: ("\n\t" + "foo bar"), F1: 777}
    v2 = v2
    v4 = (v4 & v4)
    _ = len(v3)
    v4 += v4
    for fi15 := 0; fi15 < 18; fi15++ {
        switch (("hello" + "a") + ((("hello" + "") + "hello") + ("" + ("\n\t" + "foo bar")))) {
        case "foo bar":
            switch (-(fi15)) {
            case 535:
                if v2 {
                    switch fi15 {
                    case 42:
                        if (1000 >= 42) {
                            if (127 <= 42) {
                                for fi16 := 0; fi16 < 1; fi16++ {
                                    if (0 <= 2) {
                                        v1 += (v1 ^ v1)
                                        var v17 int8 = 100
                                        if len(v3) > 0 {
                                            v3[0] = "\n\t"
                                        }
                                        _ = v17
                                    } else {
                                        if true {
                                            v18 := GenS1{F0: "hello", F1: -32768}
                                            var v19 uint = 345
                                            switch fi15 {
                                            case 405:
                                                _ = len(v3)
                                            case 128:
                                                _ = len(v3)
                                            default:
                                                var v20 uint = (v19 ^ (v19 << (((v19 ^ uint(v1))) & 63)))
                                                _ = v20
                                            }
                                            v4 = v4
                                            if true {
                                                var v21 bool = ((fi15 & fi16) == 9223372036854775807)
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if true {
                                                        switch ("" + "") {
                                                        case "":
                                                            var v22 bool = (v21 && v21)
                                                            _ = v22
                                                        case "foo bar":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            v23 := GenS2{F0: 135}
                                                            _ = v23
                                                        }
                                                        v19 = (^(v19))
                                                    } else {
                                                        fmt.Println(v1)
                                                    }
                                                default:
                                                    v19 = v19
                                                }
                                                v3 = append(v3, ("hello" + ("foo bar" + "a")))
                                                var v24 uint = (v19 | 191)
                                                switch fi16 {
                                                case 9223372036854775807:
                                                    v24 -= v24
                                                case 469:
                                                    var v25 uint16 = 1
                                                    v21 = v21
                                                    var v26 float32 = -2.25
                                                    _ = v25
                                                    _ = v26
                                                default:
                                                    if true {
                                                        _ = len(v3)
                                                        v21 = (127 != 158)
                                                        v2 = ((0.1 - (1e10 - ((0.1 * 3.14159) * float32(-2.25)))) > ((3.14159 / (3.14159 + float32(-2.25))) - 0.1))
                                                        var v27 int32 = (v4 << ((v19) & 63))
                                                        _ = v27
                                                    }
                                                }
                                                _ = v21
                                                _ = v24
                                            } else {
                                                v28 := GenS2{F0: 42}
                                                var v29 int = fi15
                                                _ = v28
                                                _ = v29
                                            }
                                            _ = v18
                                            _ = v19
                                        } else {
                                            if v2 {
                                                v4 = v4
                                                v30 := GenS2{F0: 2147483647}
                                                v2 = false
                                                var v31 int = fi15
                                                _ = v30
                                                _ = v31
                                            }
                                            if (!(-2147483648 >= int32(fi16))) {
                                                var v32 uint8 = 128
                                                switch "hello" {
                                                case "hello":
                                                    v4 -= v4
                                                default:
                                                    var v33 int8 = 1
                                                    _ = v33
                                                }
                                                v34 := GenS2{F0: 0}
                                                _ = v32
                                                _ = v34
                                            } else {
                                                _ = len(v3)
                                                v35 := GenS3{F0: 4294967295}
                                                _ = v35
                                            }
                                        }
                                    }
                                    var v36 int32 = (v4 | int32(fi15))
                                    _ = fi16
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 20; fi37++ {
                                    if true {
                                        if (!(100 > 1)) {
                                            v1 = v1
                                            switch "" {
                                            case "hello":
                                                var v38 string = ("" + "")
                                                _ = v38
                                            case "":
                                                var v39 float64 = 0.0
                                                fmt.Println(fi15, fi37, fi15)
                                                fmt.Println(v39, v39, v4)
                                                _ = v39
                                            }
                                            v2 = (-2.25 <= 0.0)
                                            if v2 {
                                                var v40 int64 = -1
                                                if v2 {
                                                    v3 = append(v3, (("foo bar" + "\n\t") + "foo bar"))
                                                }
                                                fmt.Println(fi15, v40, fi37)
                                                _ = v40
                                            } else {
                                                v3 = append(v3, ("\n\t" + "a"))
                                            }
                                            v2 = v2
                                            switch 32768 {
                                            case 182:
                                                var v41 bool = false
                                                _ = v41
                                            case 127:
                                                v42 := GenS2{F0: 2147483648}
                                                _ = v42
                                            }
                                        }
                                        _ = len(v3)
                                        var v43 int64 = 9223372036854775807
                                        _ = v43
                                    } else {
                                        v3 = append(v3, ("hello" + "\n\t"))
                                    }
                                    _ = fi37
                                }
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi44
                                }
                            } else {
                                var v45 uint = 885
                                v1 += 617
                                _ = v45
                            }
                        } else {
                            v4 = (v4 % ((v4) | 1))
                            v1 += 64
                            v2 = ((("hello" + "foo bar") != (("foo bar" + "") + "foo bar")) && true)
                        }
                    case 1:
                        var v46 uint = 4294967295
                        _ = v46
                    case -623:
                        switch "foo bar" {
                        case "hello":
                            var v47 uint = 2
                            _ = v47
                        case "foo bar":
                            v2 = v2
                        }
                    default:
                        var v48 uint8 = 66
                        _ = v48
                    }
                }
            case -592:
                switch fi15 {
                case -964:
                    if true {
                        var v49 int64 = -9223372036854775808
                        _ = v49
                    } else {
                        v50 := GenS1{F0: "a", F1: 1}
                        v1 += 2
                        _ = v50
                    }
                }
            }
        case "hello":
            fmt.Println(v1)
        default:
            fmt.Println(fi15)
        }
        var v51 uint8 = 0
        v52 := GenS4{F0: -620, F1: 794, F2: 42}
        _ = fi15
        _ = v51
        _ = v52
    }
    var v53 float64 = 1.5
    if ((v53 + (((float64(-2.25) / ((v53 - (1e10 + 1.5)) - (float64((float32(((v53 + float64(1e10)) * v53)) / (1e10 - 1.5))) + v53))) * v53) + (v53 / 1.5))) != -2.25) {
        v54 := GenS5{F0: 18446744073709551615}
        for fi55 := 0; fi55 < 1; fi55++ {
            v56 := GenS6{F0: (v1 % ((18446744073709551615) | 1)), F1: "a", F2: 0}
            _ = fi55
            _ = v56
        }
        v53 = (((-2.25 + v53) / float64(float32((v53 + (v53 + 1e10))))) - float64(3.14159))
        _ = v54
    } else {
        if (v53 == (((float64(0.1) - (float64((0.0 / 0.1)) - 0.1)) * ((v53 / v53) + float64(0.1))) - v53)) {
            if (v2 && v2) {
                fmt.Println(v2, v53)
                if (!(746 > -89)) {
                    v57 := GenS6{F0: (v1 >> ((65535) & 63)), F1: "", F2: 499}
                    for fi58 := 0; fi58 < 15; fi58++ {
                        if ((!((v1 % ((v1) | 1)) < 18446744073709551615)) && (v2 && (127 < 19))) {
                            v2 = false
                            if (-32768 < -1) {
                                v59 := GenS7{F0: 128}
                                switch ((("a" + "") + (((("a" + "foo bar") + (("a" + ("" + ("a" + ""))) + (("foo bar" + "foo bar") + "foo bar"))) + "hello") + ("hello" + "a"))) + "hello") {
                                case "foo bar":
                                    for fi60 := 0; fi60 < 20; fi60++ {
                                        var v61 uint32 = 2
                                        v2 = true
                                        if (-32768 != 42) {
                                            v62 := GenS1{F0: ("hello" + "foo bar"), F1: -32768}
                                            if (v2 && ((v61 - (^(v61))) == (v61 - v61))) {
                                                var v63 bool = (1 <= -32768)
                                                _ = v63
                                            } else {
                                                var v64 int32 = v4
                                                fmt.Println(v1)
                                                v53 = float64(((1.5 / 0.1) * 1.5))
                                                v1 = (v1 << ((uint(v1)) & 63))
                                                v65 := GenS8{F0: uint8(v1), F1: 42, F2: 42}
                                                v66 := GenS2{F0: (v61 * (v61 - v61))}
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                            _ = v62
                                        } else {
                                            var v67 int32 = v4
                                            v68 := GenS4{F0: 611, F1: 0, F2: 2}
                                            _ = v67
                                            _ = v68
                                        }
                                        _ = fi60
                                        _ = v61
                                    }
                                case "hello":
                                    if v2 {
                                        var v69 uint = 65535
                                        _ = v69
                                    }
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                switch fi58 {
                                case 4294967295:
                                    var v70 int64 = 42
                                    _ = v70
                                }
                                _ = v59
                            } else {
                                v71 := GenS7{F0: 255}
                                v1 -= 18446744073709551615
                                if (true && v2) {
                                    v2 = v2
                                }
                                _ = v71
                            }
                            v72 := GenS9{F0: (("hello" + (("\n\t" + "foo bar") + ("" + "foo bar"))) + "hello"), F1: "foo bar"}
                            _ = v72
                        } else {
                            var v73 string = "a"
                            _ = v73
                        }
                        v53 = v53
                        var v74 int16 = 1000
                        _ = fi58
                        _ = v74
                    }
                    _ = v57
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v53
}
