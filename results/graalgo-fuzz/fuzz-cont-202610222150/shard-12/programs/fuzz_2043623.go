package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
type GenS5 struct { F0 int64; F1 uint32; F2 string }
type GenS6 struct { F0 int16; F1 int; F2 uint64 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint = 286
    var v2 uint = (^(v1))
    v3 := []string{}
    var v4 float64 = float64((3.14159 * 3.14159))
    v5 := []string{("" + ("\n\t" + "")), "", "a"}
    v1 = v2
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 42
        for fi8 := 0; fi8 < 2; fi8++ {
            v7 += 0
            _ = len(v3)
            _ = fi8
        }
        var v9 uint32 = 434
        _ = fi6
        _ = v7
        _ = v9
    }
    _ = len(v5)
    if false {
        for fi10 := 0; fi10 < 5; fi10++ {
            if (1 < -1) {
                if ((!false) || (true || (!true))) {
                    fmt.Println(v2, fi10)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (0.0 < (0.1 + (0.0 / float32(v4)))) {
                        var v11 float64 = (((v4 + v4) / ((v4 + v4) + float64((0.0 * 1e10)))) / (v4 + (float64(3.14159) + ((v4 - v4) / v4))))
                        _ = v11
                    }
                }
                switch "a" {
                case "foo bar":
                    v12 := GenS1{F0: 0, F1: (("foo bar" + "a") + "foo bar"), F2: 0}
                    _ = v12
                }
            } else {
                v2 += (^(v2))
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := GenS1{F0: 2, F1: ("hello" + ("" + "hello")), F2: 245}
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = fi10
        }
        v15 := GenS2{F0: -17}
        if false {
            var v16 float64 = v4
            _ = v16
        } else {
            v2 -= (^(v1))
            for fi17 := 0; fi17 < 4; fi17++ {
                for fi18 := 0; fi18 < 20; fi18++ {
                    if (1 > 1) {
                        v4 = v4
                        v1 -= v1
                        var v19 uint16 = 0
                        var v20 float64 = (v4 / v4)
                        _ = v19
                        _ = v20
                    }
                    switch 32 {
                    case -31:
                        v3 = append(v3, ((("hello" + "") + "foo bar") + (("a" + ("a" + "a")) + (("" + "hello") + ("foo bar" + (("foo bar" + "foo bar") + ""))))))
                    case 127:
                        v4 = 1.5
                    default:
                        v1 = (v2 | v2)
                    }
                    v4 -= (float64(-2.25) + (v4 / v4))
                    v21 := GenS3{F0: "", F1: 32767}
                    _ = fi18
                    _ = v21
                }
                _ = fi17
            }
        }
        _ = v15
    } else {
        fmt.Println(v2, v2, v2)
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v1, v4)
            if (596 == 673) {
                fmt.Println(v2, v1, v1)
                if (false && false) {
                    var v22 uint64 = 42
                    var v23 uint32 = 2147483648
                    var v24 uint64 = v22
                    _ = v22
                    _ = v23
                    _ = v24
                }
            } else {
                _ = len(v3)
            }
        }
    }
    for fi25 := 0; fi25 < 4; fi25++ {
        switch int(v1) {
        case 32767:
            v3 = append(v3, (("foo bar" + "\n\t") + "\n\t"))
        case -9223372036854775808:
            var v26 int8 = int8(fi25)
            _ = v26
        default:
            v4 += v4
        }
        v4 = 3.14159
        _ = fi25
    }
    v2 = v2
    v4 = v4
    v1 += (v2 % ((v2) | 1))
    v27 := GenS2{F0: 794}
    fmt.Println(v1, v1)
    var v28 float32 = 0.1
    if false {
        var v29 uint32 = 65
        v30 := GenS4{F0: -839, F1: v2, F2: 2}
        var v31 float64 = v4
        _ = v29
        _ = v30
        _ = v31
    }
    var v32 float32 = 0.0
    for fi33 := 0; fi33 < 15; fi33++ {
        fmt.Println(v1)
        v34 := GenS5{F0: 0, F1: 54, F2: ("" + "hello")}
        if ((v1 + (v1 & (v2 + v2))) <= (^(v2))) {
            var v35 bool = true
            _ = v35
        } else {
            if (42 < 1) {
                if false {
                    v1 = v1
                    for fi36 := 0; fi36 < 13; fi36++ {
                        v37 := GenS6{F0: -1, F1: fi36, F2: 2}
                        v38 := GenS5{F0: 9223372036854775807, F1: 4294967295, F2: ""}
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                }
            } else {
                for fi39 := 0; fi39 < 13; fi39++ {
                    v40 := GenS7{F0: (fi39 & fi33)}
                    v4 += (v4 - 1.5)
                    _ = fi39
                    _ = v40
                }
                switch ("\n\t" + ("\n\t" + (("\n\t" + "hello") + ""))) {
                case "foo bar":
                    if true {
                        v4 = float64(v28)
                        var v41 int = -86
                        v1 += (v1 & 512)
                        _ = v41
                    } else {
                        v32 = 1e10
                    }
                case "":
                    if false {
                        for fi42 := 0; fi42 < 18; fi42++ {
                            if len(v5) > 0 {
                                v5[0] = "\n\t"
                            }
                            v43 := GenS4{F0: 174, F1: v1, F2: 640}
                            _ = fi42
                            _ = v43
                        }
                        _ = len(v3)
                    }
                }
                v2 = v1
            }
            v32 -= v28
            if false {
                switch (fi33 * 260) {
                case -2147483648:
                    fmt.Println(v2)
                case -1:
                    if false {
                        for fi44 := 0; fi44 < 20; fi44++ {
                            var v45 int16 = 1
                            switch (fi33 % ((fi33) | 1)) {
                            case -9223372036854775808:
                                if false {
                                    switch "foo bar" {
                                    case "\n\t":
                                        if ((^(v1)) > v2) {
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                v45 += 32767
                                            case "hello":
                                                if len(v5) > 0 {
                                                    v5[0] = "a"
                                                }
                                            }
                                            switch 32768 {
                                            case 4294967295:
                                                v32 = (((v28 / v32) - (v32 * ((1e10 - v28) - (v32 / (1.5 / ((v28 * v32) - v28)))))) - v28)
                                            case 165:
                                                fmt.Println(v2, v2)
                                            }
                                        }
                                    case "foo bar":
                                        var v46 string = "hello"
                                        _ = v46
                                    case "hello":
                                        if len(v3) > 0 {
                                            v3[0] = (("foo bar" + "hello") + ("" + (("" + "hello") + "a")))
                                        }
                                    default:
                                        fmt.Println(fi33)
                                        var v47 int8 = 61
                                        _ = v47
                                    }
                                    fmt.Println(v1, v1, v2)
                                }
                            case 4294967295:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case -624:
                                v45 = 32767
                            }
                            switch 32767 {
                            case 32767:
                                var v48 int16 = -32768
                                _ = v48
                            case 2147483647:
                                v1 = v1
                            case 4294967295:
                                v1 = (v2 ^ v1)
                            default:
                                var v49 bool = (false && true)
                                _ = v49
                            }
                            _ = fi44
                            _ = v45
                        }
                    }
                case 9223372036854775807:
                    for fi50 := 0; fi50 < 6; fi50++ {
                        var v51 uint16 = 606
                        if (42 <= v51) {
                            if false {
                                _ = len(v3)
                                v2 = (v1 << (((v1 % ((v1) | 1))) & 63))
                                _ = len(v3)
                            } else {
                                if false {
                                    var v52 int8 = -1
                                    v28 -= ((v28 * float32((1.5 / 3.14159))) / v28)
                                    _ = v52
                                }
                                v28 = 0.0
                            }
                            v1 = v1
                        } else {
                            var v53 uint32 = 42
                            var v54 uint8 = 255
                            if true {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if false {
                                    v55 := GenS6{F0: -1, F1: (fi50 * 1), F2: 2}
                                    var v56 uint8 = (v54 & v54)
                                    var v57 int16 = 1000
                                    if (!true) {
                                        if (42 < 0) {
                                            v5 = append(v5, "foo bar")
                                        } else {
                                            if (!false) {
                                                if true {
                                                    var v58 int64 = 42
                                                    _ = v58
                                                } else {
                                                    var v59 int64 = 1
                                                    if true {
                                                        fmt.Println(fi33, fi50)
                                                        v32 -= 0.0
                                                        v51 -= 2
                                                        var v60 int32 = 522
                                                        _ = v60
                                                    }
                                                    _ = v59
                                                }
                                                v53 += (v53 - 65535)
                                            }
                                        }
                                        v61 := GenS3{F0: ("" + ("\n\t" + ("" + (("\n\t" + "foo bar") + "\n\t")))), F1: v51}
                                        _ = v61
                                    }
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                } else {
                                    _ = len(v3)
                                }
                            }
                            _ = v53
                            _ = v54
                        }
                        switch ((("" + ("" + (("\n\t" + ("\n\t" + "\n\t")) + "\n\t"))) + "\n\t") + "hello") {
                        case "\n\t":
                            if len(v5) > 0 {
                                v5[0] = ("" + ("a" + "foo bar"))
                            }
                        case "":
                            v32 += 3.14159
                        case "hello":
                            fmt.Println(v28, v2, fi33)
                        default:
                            v32 -= float32(1.5)
                            fmt.Println(v2, v2, v2)
                            var v62 uint64 = 0
                            switch "\n\t" {
                            case "a":
                                var v63 uint8 = 51
                                _ = v63
                            case "\n\t":
                                var v64 int32 = 1
                                _ = v64
                            }
                            if (!(!(253 >= 0))) {
                                var v65 string = "hello"
                                _ = v65
                            } else {
                                var v66 uint8 = 127
                                _ = v66
                            }
                            _ = v62
                        }
                        _ = fi50
                        _ = v51
                    }
                default:
                    v4 = -2.25
                }
            }
        }
        _ = fi33
        _ = v34
    }
    switch "" {
    case "foo bar":
        var v67 float32 = v32
        _ = v67
    default:
        for fi68 := 0; fi68 < 9; fi68++ {
            if (255 > v1) {
                var v69 int64 = -9223372036854775808
                var v70 int16 = 1
                _ = v69
                _ = v70
            }
            var v71 uint8 = 127
            v32 += ((((v32 - 3.14159) * (v28 / 3.14159)) + v32) + v28)
            _ = fi68
            _ = v71
        }
    }
    var v72 uint16 = 790
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v72
}
