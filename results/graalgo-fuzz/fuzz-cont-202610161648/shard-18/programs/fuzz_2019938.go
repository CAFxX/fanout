package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 string; F1 int32; F2 string }
type GenS6 struct { F0 int; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint32; F1 uint64; F2 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 bool; F2 int8 }
type GenS10 struct { F0 int; F1 uint16; F2 string }
type GenS11 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32, p1 uint64, p2 string) string {
    switch p2 {
    case "a":
        p1 = p1
    case "foo bar":
        v7 := []int{15, -983, 42}
        _ = v7
    default:
        v8 := []int{-709}
        _ = v8
    }
    var v9 int16 = 32767
    fmt.Println(v9, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return "a"
}

func main() {
    var v1 int = 571
    var v2 float32 = ((1.5 - float32(float64(0.1))) * (float32(0.1) + float32(0.1)))
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{65535}
    var v6 int16 = -1
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 float64 = (1.5 / float64(v2))
        _ = fi10
        _ = v11
    }
    fmt.Println(v6, v2)
    if true {
        var v12 string = "foo bar"
        var v13 int64 = -969
        v14 := GenS1{F0: "hello"}
        _ = v12
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: "foo bar"}
            if (((v2 + 3.14159) >= (((v2 + v2) + (1.5 + v2)) + 1e10)) || false) {
                var v17 int8 = 1
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v2)
    if (4294967295 == 4294967295) {
        fmt.Println(v1)
        var v18 uint = 255
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS2{F0: 50}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v1 = (-(fi20))
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v2 += v2
                v23 := GenS2{F0: 82}
                v1 += v1
                _ = fi22
                _ = v23
            }
            if true {
                var v24 uint16 = 42
                v1 = (v1 << ((v18) & 63))
                _ = v24
            } else {
                v2 -= v2
            }
            _ = fi19
        }
        _ = v18
    }
    if (v1 != (-(v1))) {
        var v25 uint = 0
        _ = v25
    } else {
        v26 := GenS2{F0: 42}
        v2 = ((v2 - 0.0) - (3.14159 / v2))
        _ = v26
    }
    if (!(!(-2147483648 >= 1))) {
        v27 := GenS3{F0: 1, F1: v1}
        _ = v27
    }
    if (127 < -1) {
        switch v1 {
        case 127:
            for fi28 := 0; fi28 < 6; fi28++ {
                var v29 uint = 255
                _ = fi28
                _ = v29
            }
        case 4294967295:
            fmt.Println(v1, v6)
        default:
            switch (v1 | (v1 + v1)) {
            case -680:
                switch 886 {
                case 32768:
                    fmt.Println(v1)
                case 4294967295:
                    if true {
                        if false {
                            if false {
                                v6 = v6
                                switch "\n\t" {
                                case "foo bar":
                                    for fi30 := 0; fi30 < 13; fi30++ {
                                        var v31 float32 = v2
                                        switch v1 {
                                        case -961:
                                            v6 -= -1
                                        case -2147483648:
                                            v32 := GenS4{F0: 0}
                                            _ = v32
                                        case -317:
                                            _ = len(v3)
                                        default:
                                            for fi33 := 0; fi33 < 13; fi33++ {
                                                v4[337] = 2147483648
                                                v6 -= (^(v6))
                                                _ = fi33
                                            }
                                        }
                                        v5 = append(v5, 59)
                                        _ = fi30
                                        _ = v31
                                    }
                                case "":
                                    fmt.Println(v1, v6, v6)
                                case "a":
                                    v34 := GenS5{F0: (("" + "hello") + (("hello" + "hello") + ("" + "hello"))), F1: 1, F2: (("foo bar" + "hello") + "foo bar")}
                                    _ = v34
                                default:
                                    switch (((("a" + ("" + "")) + (("" + "\n\t") + ("hello" + ("hello" + ("" + (("a" + ("foo bar" + "hello")) + "")))))) + ("hello" + "hello")) + ("\n\t" + ((("a" + ("a" + ("" + ("hello" + "")))) + "foo bar") + "foo bar"))) {
                                    case "\n\t":
                                        for fi35 := 0; fi35 < 8; fi35++ {
                                            v36 := GenS5{F0: "hello", F1: 42, F2: "\n\t"}
                                            _ = fi35
                                            _ = v36
                                        }
                                    case "":
                                        var v37 uint32 = 4294967295
                                        _ = v37
                                    case "foo bar":
                                        v1 += -571
                                    }
                                }
                                v2 = float32(3.14159)
                            } else {
                                v38 := GenS6{F0: (-(v1)), F1: 18446744073709551615, F2: 0}
                                _ = v38
                            }
                            if (((float32(-2.25) - (1.5 * v2)) <= 0.1) && true) {
                                for fi39 := 0; fi39 < 12; fi39++ {
                                    var v40 int = v1
                                    _ = fi39
                                    _ = v40
                                }
                            }
                            var v41 uint32 = 65535
                            _ = v41
                        } else {
                            v1 += v1
                            v42 := GenS7{F0: 290, F1: 4294967295, F2: (false && true)}
                            _ = v42
                        }
                        var v43 bool = true
                        _ = v43
                    } else {
                        v44 := GenS6{F0: -1, F1: 0, F2: 468}
                        v2 = (float32(-2.25) * 3.14159)
                        _ = v44
                    }
                case 0:
                    switch (v1 % ((int(v6)) | 1)) {
                    case 4294967295:
                        if true {
                            v6 = (v6 + (v6 >> ((0) & 63)))
                            var v45 uint32 = 0
                            _ = v45
                        } else {
                            v1 -= 0
                            fmt.Println(v2)
                            switch -651 {
                            case 4294967295:
                                var v46 uint32 = 2
                                _ = v46
                            case 127:
                                v2 = ((v2 * 0.0) / v2)
                            case 32768:
                                if false {
                                    var v47 float64 = float64((v2 / 1.0))
                                    fmt.Println(v1)
                                    _ = v47
                                }
                            default:
                                if (!(true && (-2.25 == 0.1))) {
                                    if true {
                                        switch (v1 + v1) {
                                        case 42:
                                            v5 = append(v5, 603)
                                        case -9223372036854775808:
                                            var v48 int32 = 0
                                            _ = v48
                                        case 1:
                                            v1 -= -2147483648
                                        }
                                        for fi49 := 0; fi49 < 3; fi49++ {
                                            var v50 string = "foo bar"
                                            _ = fi49
                                            _ = v50
                                        }
                                    } else {
                                        if true {
                                            v6 -= v6
                                        } else {
                                            fmt.Println(v1, v6)
                                            delete(v4, 1)
                                        }
                                        switch 32767 {
                                        case -898:
                                            v2 = (v2 * v2)
                                        }
                                        for fi51 := 0; fi51 < 13; fi51++ {
                                            v52 := v4[1]
                                            v53 := GenS8{F0: 762}
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                }
                            }
                        }
                    case 612:
                        for fi54 := 0; fi54 < 4; fi54++ {
                            var v55 int = (-(v1))
                            var v56 uint8 = 1
                            _ = fi54
                            _ = v55
                            _ = v56
                        }
                    case 0:
                        var v57 uint8 = 84
                        _ = v57
                    default:
                        var v58 float64 = 1e10
                        _ = v58
                    }
                }
            case -671:
                v1 = v1
            case 4294967295:
                for fi59 := 0; fi59 < 20; fi59++ {
                    if true {
                        var v60 string = "hello"
                        v2 += float32(float64((0.0 + -2.25)))
                        _ = v60
                    } else {
                        v61 := GenS9{F0: -839, F1: false, F2: 79}
                        v62 := v4[2147483647]
                        for fi63 := 0; fi63 < 9; fi63++ {
                            switch "" {
                            case "a":
                                if true {
                                    var v64 int16 = -239
                                    _ = len(v5)
                                    var v65 int = 4294967295
                                    _ = v64
                                    _ = v65
                                }
                            case "foo bar":
                                v5 = append(v5, v62)
                                v62 = 959
                                fmt.Println(fi63, v62, v1)
                            default:
                                switch ("" + ((("" + (("foo bar" + ("foo bar" + "hello")) + "\n\t")) + "hello") + "\n\t")) {
                                case "":
                                    v2 += (v2 / -2.25)
                                case "hello":
                                    v66 := GenS10{F0: fi63, F1: 2, F2: ((("" + "foo bar") + "") + "hello")}
                                    var v67 int8 = 7
                                    fmt.Println(v1)
                                    _ = v66
                                    _ = v67
                                }
                            }
                            fmt.Println(v2, fi59, v6)
                            var v68 int64 = -357
                            _ = fi63
                            _ = v68
                        }
                        _ = v61
                        _ = v62
                    }
                    if ((^(v6)) >= 1) {
                        var v69 int = 720
                        switch 32768 {
                        case 755:
                            switch -9223372036854775808 {
                            case 42:
                                v69 = (v1 - fi59)
                            default:
                                v2 -= float32(float64(v2))
                            }
                        }
                        _ = len(v5)
                        _ = v69
                    }
                    var v70 int64 = 42
                    _ = fi59
                    _ = v70
                }
            default:
                for fi71 := 0; fi71 < 14; fi71++ {
                    var v72 int32 = 0
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v73 := GenS7{F0: 42, F1: 1, F2: (-108 != 127)}
                    _ = fi71
                    _ = v72
                    _ = v73
                }
            }
        }
        v1 = v1
    } else {
        for fi74 := 0; fi74 < 10; fi74++ {
            v5 = append(v5, 65535)
            _ = fi74
        }
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        fmt.Println(v2, v1)
        _ = fi75
    }
    fmt.Println(v6, v1)
    var v76 int16 = v6
    for fi77 := 0; fi77 < 15; fi77++ {
        v78 := GenS11{F0: -1, F1: 248}
        _ = fi77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
}
