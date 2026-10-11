package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 uint8; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS8 struct { F0 int8 }
func genFn0(p0 bool, p1 int64, p2 int32) int8 {
    var v5 string = ("foo bar" + (("hello" + ((("foo bar" + "foo bar") + "hello") + "foo bar")) + "a"))
    var v6 int32 = (-(p2))
    for fi7 := 0; fi7 < 7; fi7++ {
        p1 -= p1
        v8 := GenS1{F0: p2, F1: 32767}
        var v9 float32 = (1.5 + ((-2.25 + -2.25) / 1e10))
        _ = fi7
        _ = v8
        _ = v9
    }
    if false {
        fmt.Println(v5, p1)
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := []int{(fi10 - (^(fi10))), fi10}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -55
}
func genFn1() int8 {
    if ((!false) && (!true)) {
        var v12 int32 = 1
        var v13 int = -2147483648
        v14 := GenS1{F0: v12, F1: -32768}
        _ = v12
        _ = v13
        _ = v14
    } else {
        v15 := []int{2147483647, 4294967295}
        _ = v15
    }
    if false {
        switch "a" {
        case "a":
            for fi16 := 0; fi16 < 18; fi16++ {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := []int{(fi17 - 128), -514, fi16, fi16}
                    switch "foo bar" {
                    case "hello":
                        v19 := GenS2{F0: 0}
                        _ = v19
                    default:
                        if (65535 < 912) {
                            var v20 int = 42
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            v21 := GenS1{F0: 991, F1: 0}
                            _ = v20
                            _ = v21
                        }
                    }
                    var v22 float64 = (3.14159 * -2.25)
                    _ = fi17
                    _ = v18
                    _ = v22
                }
                v23 := []int{fi16, (-(fi16)), fi16}
                var v24 uint = 65535
                _ = fi16
                _ = v23
                _ = v24
            }
        case "hello":
            var v25 int64 = 634
            _ = v25
        case "":
            v26 := GenS1{F0: 42, F1: -32768}
            _ = v26
        default:
            var v27 uint64 = 4294967295
            _ = v27
        }
        if true {
            var v28 uint64 = 4294967295
            switch -200 {
            case 42:
                v29 := GenS1{F0: -853, F1: -1}
                _ = v29
            case 310:
                v30 := []int{2147483647, 42, 127}
                _ = v30
            default:
                var v31 int8 = 100
                _ = v31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                if (!false) {
                    if (-719 == 35) {
                        v33 := GenS1{F0: 375, F1: -1}
                        _ = v33
                    }
                    switch "a" {
                    case "hello":
                        for fi34 := 0; fi34 < 19; fi34++ {
                            switch fi32 {
                            case 9223372036854775807:
                                var v35 float32 = 3.14159
                                _ = v35
                            case 42:
                                v36 := GenS1{F0: -180, F1: 1}
                                _ = v36
                            }
                            switch "" {
                            case "":
                                v37 := GenS3{F0: -331, F1: v28, F2: (fi32 * (fi32 * fi32))}
                                v38 := []int{(fi34 / ((4294967295) | 1)), (^(fi34)), 565, (fi34 << ((0) & 63))}
                                v39 := GenS2{F0: -229}
                                _ = v37
                                _ = v38
                                _ = v39
                            case "foo bar":
                                var v40 string = "hello"
                                _ = v40
                            case "\n\t":
                                fmt.Println(fi34)
                            default:
                                if false {
                                    v41 := GenS4{F0: (32767 <= 0)}
                                    switch (^(fi32)) {
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "hello":
                                            v28 += 2
                                        case "a":
                                            v42 := GenS3{F0: 42, F1: 687, F2: (fi34 - -2147483648)}
                                            _ = v42
                                        case "\n\t":
                                            switch ("hello" + "") {
                                            case "foo bar":
                                                fmt.Println(fi32, fi32)
                                            case "":
                                                v43 := GenS3{F0: -1, F1: v28, F2: fi34}
                                                _ = v43
                                            default:
                                                switch "foo bar" {
                                                case "hello":
                                                    v44 := GenS5{F0: false, F1: 252, F2: 853}
                                                    _ = v44
                                                }
                                            }
                                        }
                                    }
                                    var v45 int16 = 1
                                    _ = v41
                                    _ = v45
                                }
                            }
                            _ = fi34
                        }
                    case "foo bar":
                        if false {
                            var v46 int32 = 42
                            var v47 int = fi32
                            var v48 int64 = 42
                            _ = v46
                            _ = v47
                            _ = v48
                        } else {
                            v28 = (v28 - (v28 ^ 139))
                            fmt.Println(fi32, fi32)
                        }
                    case "":
                        v49 := []int{128, 2147483647, (^(fi32)), (fi32 % (((fi32 / (((fi32 - (fi32 & -1))) | 1))) | 1))}
                        _ = v49
                    default:
                        fmt.Println(v28)
                    }
                    var v50 uint32 = 2
                    _ = v50
                } else {
                    var v51 uint = 18446744073709551615
                    switch -312 {
                    case 4294967295:
                        for fi52 := 0; fi52 < 16; fi52++ {
                            fmt.Println(v28)
                            fmt.Println(fi52, v28)
                            v28 = 1
                            _ = fi52
                        }
                    }
                    _ = v51
                }
                _ = fi32
            }
            _ = v28
        } else {
            var v53 float64 = 1e10
            v54 := GenS6{F0: 42}
            if (2147483648 < 2147483647) {
                v53 += v53
                for fi55 := 0; fi55 < 2; fi55++ {
                    var v56 int64 = 9223372036854775807
                    if true {
                        v57 := GenS5{F0: false, F1: 42, F2: 18446744073709551615}
                        _ = v57
                    } else {
                        fmt.Println(fi55, v56)
                    }
                    _ = fi55
                    _ = v56
                }
            }
            _ = v53
            _ = v54
        }
    } else {
        var v58 int = 127
        if (1.5 != 3.14159) {
            for fi59 := 0; fi59 < 16; fi59++ {
                if (((667 > 1) || (("hello" >= "") && true)) || true) {
                    v60 := []int{(v58 | (fi59 | 2147483647)), (fi59 / ((-175) | 1))}
                    if len(v60) > 0 {
                        v60[0] = v58
                    }
                    var v61 uint32 = 1
                    _ = v60
                    _ = v61
                }
                if true {
                    fmt.Println(v58, fi59)
                }
                var v62 uint = 910
                _ = fi59
                _ = v62
            }
        } else {
            v63 := GenS4{F0: (false && (42 >= 0))}
            _ = v63
        }
        fmt.Println(v58)
        _ = v58
    }
    v64 := []int{0, -214, 2147483647, 32768}
    var v65 int8 = 49
    _ = v64
    _ = v65
    return 0
}

func main() {
    var v1 bool = ((0 != 1) && (1.5 != 3.14159))
    var v2 uint = 800
    var v3 bool = (true && true)
    var v4 int = 2147483647
    for fi66 := 0; fi66 < 5; fi66++ {
        var v67 string = ("a" + "hello")
        if false {
            fmt.Println(v3, v4)
            var v68 uint = (^(v2))
            _ = v68
        } else {
            v67 = (v67 + v67)
        }
        _ = fi66
        _ = v67
    }
    var v69 uint32 = 2
    fmt.Println(v3)
    for fi70 := 0; fi70 < 19; fi70++ {
        v69 = v69
        _ = fi70
    }
    v71 := []int{v4, v4}
    var v72 int = (v4 >> ((v2) & 63))
    if v1 {
        var v73 uint32 = v69
        if false {
            var v74 int = v4
            _ = v74
        }
        if (v73 > v69) {
            if v3 {
                v73 += v73
            } else {
                v69 = 4294967295
                var v75 int16 = -992
                _ = v75
            }
        }
        _ = v73
    }
    var v76 int = (v72 ^ v72)
    v77 := GenS7{F0: v69, F1: 65535, F2: 168}
    var v78 float64 = 1.5
    v79 := GenS8{F0: 92}
    fmt.Println(v69, v2)
    for fi80 := 0; fi80 < 11; fi80++ {
        var v81 uint8 = 1
        _ = fi80
        _ = v81
    }
    fmt.Println(v78, v78)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("slicelen v71:", len(v71), v71)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v71
    _ = v72
    _ = v76
    _ = v77
    _ = v78
    _ = v79
}
