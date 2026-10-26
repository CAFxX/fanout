package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint16; F2 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool }
func genFn0(p0 int32, p1 uint8) int32 {
    var v5 uint8 = p1
    fmt.Println(p1, p1, v5)
    p0 += (p0 * (p0 + 42))
    _ = p0
    _ = p1
    _ = v5
    return 806
}

func main() {
    v1 := []string{}
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 float32 = ((1e10 / float32((1.5 - 0.1))) + float32(3.14159))
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 2, F1: -128, F2: 1}
        _ = fi6
        _ = v7
    }
    if v2 {
        switch ("foo bar" + "\n\t") {
        case "a":
            v4 = 3.14159
        case "foo bar":
            fmt.Println(v2, v4)
        case "":
            v8 := GenS1{F0: 2, F1: 71, F2: 0}
            _ = v8
        default:
            v3[-1] = ((("hello" + "a") + (("\n\t" + "a") + "foo bar")) + "foo bar")
        }
        _ = len(v3)
    } else {
        if (0.0 == float64(0.1)) {
            v4 += v4
            switch "hello" {
            case "hello":
                if ((!v2) || (v2 && (0.0 == 0.1))) {
                    var v9 string = (("hello" + "\n\t") + "foo bar")
                    _ = v9
                }
            }
        }
        if v2 {
            switch 32768 {
            case 42:
                var v10 int8 = 2
                _ = v10
            case 128:
                if (32767 != 255) {
                    v11 := GenS1{F0: 128, F1: 0, F2: -330}
                    _ = v11
                }
            default:
                for fi12 := 0; fi12 < 11; fi12++ {
                    var v13 int16 = 559
                    _ = fi12
                    _ = v13
                }
            }
            v14 := GenS2{F0: v2, F1: 1}
            v15 := GenS3{F0: (1 < 2)}
            _ = v14
            _ = v15
        } else {
            var v16 bool = (v2 && false)
            if ((65535 <= 71) || (true && false)) {
                if ((2147483647 <= 0) || v16) {
                    if false {
                        var v17 int8 = 127
                        v3[4294967295] = ("foo bar" + "hello")
                        _ = v17
                    } else {
                        v18 := GenS4{F0: 127}
                        if (("" <= "foo bar") && ((4294967295 <= 2147483647) && (88 <= -32768))) {
                            v2 = v16
                        } else {
                            v16 = v2
                            var v19 uint16 = 32767
                            switch -2147483648 {
                            case 427:
                                switch "" {
                                case "\n\t":
                                    var v20 float64 = (0.0 + ((float64((float32(1.5) / ((v4 / float32((0.1 / 1.0))) / v4))) - 3.14159) - 3.14159))
                                    _ = v20
                                case "foo bar":
                                    v2 = (v2 || (!(42 == 535)))
                                case "a":
                                    switch "a" {
                                    case "hello":
                                        switch 4294967295 {
                                        case 1:
                                            v19 -= (v19 | v19)
                                        default:
                                            for fi21 := 0; fi21 < 10; fi21++ {
                                                _ = len(v1)
                                                fmt.Println(fi21)
                                                v1 = append(v1, ("hello" + ("\n\t" + "hello")))
                                                _ = fi21
                                            }
                                        }
                                    }
                                }
                            case 0:
                                var v22 uint = 4294967295
                                _ = v22
                            default:
                                for fi23 := 0; fi23 < 3; fi23++ {
                                    v16 = (!(!(("" + "a") > "a")))
                                    if false {
                                        var v24 uint = 349
                                        _ = v24
                                    }
                                    _ = fi23
                                }
                            }
                            _ = v19
                        }
                        _ = v18
                    }
                    _ = len(v1)
                    switch "a" {
                    case "foo bar":
                        if v16 {
                            if (65535 > 2) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            for fi25 := 0; fi25 < 2; fi25++ {
                                fmt.Println(fi25, v16)
                                v26 := GenS5{F0: 180, F1: 31, F2: 2147483647}
                                switch "hello" {
                                case "foo bar":
                                    v27 := v3[933]
                                    _ = v27
                                case "hello":
                                    var v28 uint = 18446744073709551615
                                    _ = v28
                                case "":
                                    var v29 float64 = 3.14159
                                    _ = v29
                                }
                                _ = fi25
                                _ = v26
                            }
                            if true {
                                v2 = true
                            }
                        }
                    default:
                        var v30 uint16 = 478
                        _ = v30
                    }
                }
                fmt.Println(v4)
                v2 = v2
            } else {
                _ = len(v3)
                for fi31 := 0; fi31 < 13; fi31++ {
                    var v32 bool = (!(!v2))
                    for fi33 := 0; fi33 < 11; fi33++ {
                        v34 := GenS5{F0: 0, F1: 0, F2: 0}
                        _ = fi33
                        _ = v34
                    }
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    _ = fi31
                    _ = v32
                }
                for fi35 := 0; fi35 < 11; fi35++ {
                    for fi36 := 0; fi36 < 20; fi36++ {
                        v37 := GenS4{F0: -99}
                        if (1 <= 42) {
                            _ = len(v3)
                            fmt.Println(v4)
                            switch ("foo bar" + ("\n\t" + "\n\t")) {
                            case "hello":
                                v38 := GenS5{F0: 198, F1: 0, F2: 65535}
                                _ = v38
                            case "\n\t":
                                v39 := GenS4{F0: 80}
                                _ = v39
                            default:
                                v40 := GenS6{F0: -217, F1: 2, F2: fi36}
                                _ = v40
                            }
                        }
                        _ = fi36
                        _ = v37
                    }
                    var v41 uint = 621
                    var v42 float32 = (-2.25 / v4)
                    _ = fi35
                    _ = v41
                    _ = v42
                }
            }
            _ = v16
        }
    }
    var v43 bool = (true || v2)
    for fi44 := 0; fi44 < 12; fi44++ {
        var v45 float32 = (v4 * 1e10)
        var v46 int8 = -128
        for fi47 := 0; fi47 < 8; fi47++ {
            v45 = v45
            switch "foo bar" {
            case "\n\t":
                v48 := v3[(-(fi47))]
                if true {
                    v43 = (v2 || ((v48 + ((("foo bar" + "a") + (("foo bar" + v48) + (("hello" + ("foo bar" + "foo bar")) + v48))) + v48)) == "hello"))
                    var v49 int16 = 42
                    _ = v49
                } else {
                    v4 = (float32((3.14159 + (0.0 + 1e10))) * 1e10)
                }
                v50 := GenS3{F0: ((false || true) && v43)}
                _ = v48
                _ = v50
            case "foo bar":
                v4 += 0.0
            }
            _ = fi47
        }
        _ = fi44
        _ = v45
        _ = v46
    }
    for fi51 := 0; fi51 < 7; fi51++ {
        var v52 int8 = 0
        _ = fi51
        _ = v52
    }
    switch 128 {
    case -9223372036854775808:
        var v53 uint16 = 0
        _ = v53
    }
    var v54 int32 = genFn0(2147483647, 42)
    if (380 < 255) {
        switch "" {
        case "":
            v4 += ((1.5 - float32(((float64(v4) + 0.1) / 1e10))) / -2.25)
        case "foo bar":
            switch "hello" {
            case "a":
                v43 = v2
            case "foo bar":
                var v55 float32 = ((float32((0.1 + -2.25)) + (((0.0 - v4) + v4) * (v4 - (v4 - v4)))) * ((0.1 / v4) * v4))
                _ = v55
            case "":
                v2 = v2
            }
        default:
            var v56 uint8 = 165
            _ = v56
        }
    }
    var v57 uint32 = 65535
    for fi58 := 0; fi58 < 14; fi58++ {
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 int = fi58
            v54 -= (v54 * -148)
            if false {
                v61 := GenS7{F0: v54}
                fmt.Println(v57)
                _ = v61
            } else {
                v60 = int(v54)
                var v62 float32 = (0.1 - 0.0)
                _ = v62
            }
            fmt.Println(v57, fi58, v4)
            _ = fi59
            _ = v60
        }
        switch ("\n\t" + "a") {
        case "foo bar":
            for fi63 := 0; fi63 < 15; fi63++ {
                v43 = (!false)
                _ = fi63
            }
        case "":
            switch "hello" {
            case "hello":
                var v64 string = "hello"
                _ = v64
            case "foo bar":
                v65 := GenS8{F0: v43}
                _ = v65
            case "a":
                v66 := GenS6{F0: int16(v57), F1: 65535, F2: fi58}
                _ = v66
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        _ = fi58
    }
    var v67 int = 0
    var v68 uint8 = 187
    v67 = (-(v67))
    for fi69 := 0; fi69 < 6; fi69++ {
        v54 += -948
        _ = fi69
    }
    v2 = v43
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v54:", v54)
    fmt.Println("v57:", v57)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v54
    _ = v57
    _ = v67
    _ = v68
}
