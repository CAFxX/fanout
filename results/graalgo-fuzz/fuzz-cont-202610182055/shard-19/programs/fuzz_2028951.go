package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint; F2 int64 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint; F1 int16; F2 bool }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 bool; F2 uint64 }

func main() {
    var v1 int8 = 100
    var v2 uint = 2
    v3 := make(map[int]int)
    var v4 int16 = 1000
    v5 := GenS1{F0: 2, F1: 239, F2: 34}
    if true {
        v6 := GenS1{F0: 0, F1: v2, F2: 22}
        fmt.Println(v4)
        _ = v6
    }
    var v7 float32 = (1e10 * 1.5)
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int = (fi8 >> (((v2 | 1)) & 63))
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := GenS1{F0: 255, F1: (^(v2)), F2: 42}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: 255, F1: v2, F2: 60}
        _ = fi12
        _ = v13
    }
    if (true || ((-1 > 42) && true)) {
        if (-1 < 177) {
            v14 := []int{4294967295, -2147483648, 0, 0}
            var v15 uint32 = 2
            var v16 uint = v2
            _ = v14
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 1; fi17++ {
                v7 = v7
                v18 := GenS2{F0: (v1 % ((127) | 1)), F1: 0, F2: 42}
                for fi19 := 0; fi19 < 4; fi19++ {
                    if (2147483648 < 691) {
                        switch fi17 {
                        case -9223372036854775808:
                            if true {
                                var v20 uint16 = 121
                                _ = v20
                            }
                        case 32767:
                            v2 = (v2 & v2)
                        default:
                            var v21 int = (fi17 & fi19)
                            _ = v21
                        }
                        var v22 uint16 = 42
                        fmt.Println(v1, v4, fi19)
                        v23 := GenS1{F0: 844, F1: v2, F2: uint8(v22)}
                        v1 += (v1 + 127)
                        _ = v22
                        _ = v23
                    } else {
                        v1 = 1
                        v24 := GenS1{F0: 255, F1: (v2 % ((4294967295) | 1)), F2: 255}
                        v1 = (v1 % ((v1) | 1))
                        _ = v24
                    }
                    var v25 int = fi17
                    var v26 uint8 = 0
                    _ = fi19
                    _ = v25
                    _ = v26
                }
                _ = fi17
                _ = v18
            }
            v27 := GenS1{F0: 2, F1: v2, F2: 161}
            if ((0 == 2147483647) || (394 > 0)) {
                fmt.Println(v4, v1, v1)
                switch 127 {
                case 4294967295:
                    if false {
                        for fi28 := 0; fi28 < 7; fi28++ {
                            for fi29 := 0; fi29 < 16; fi29++ {
                                switch (-(fi28)) {
                                case -1:
                                    fmt.Println(fi29, v4)
                                case -780:
                                    fmt.Println(fi28, fi28)
                                    var v30 uint16 = 940
                                    _ = v30
                                }
                                v31 := GenS3{F0: (v7 >= (3.14159 + (float32(1e10) + v7))), F1: (v2 / ((0) | 1)), F2: -9223372036854775808}
                                _ = fi29
                                _ = v31
                            }
                            _ = fi28
                        }
                        v32 := GenS4{F0: 376, F1: -469}
                        fmt.Println(v2)
                        _ = v32
                    }
                case -306:
                    var v33 int32 = -1
                    _ = v33
                case 321:
                    v34 := []int{765, 4294967295, 0}
                    _ = v34
                default:
                    var v35 float32 = v7
                    _ = v35
                }
                for fi36 := 0; fi36 < 16; fi36++ {
                    v4 = (v4 % ((v4) | 1))
                    if false {
                        v4 = (v4 % (((v4 << ((v2) & 63))) | 1))
                        v37 := GenS5{F0: (fi36 / ((125) | 1))}
                        for fi38 := 0; fi38 < 14; fi38++ {
                            v4 -= (v4 / (((v4 + v4)) | 1))
                            switch "\n\t" {
                            case "hello":
                                if false {
                                    fmt.Println(v2, v1)
                                    switch 32767 {
                                    case 140:
                                        if false {
                                            var v39 string = (((("\n\t" + ("a" + "hello")) + ((("hello" + ((("" + ("\n\t" + "")) + "hello") + (("a" + ("foo bar" + "\n\t")) + "\n\t"))) + "foo bar") + ("" + "hello"))) + ("\n\t" + "")) + "hello")
                                            v1 = v1
                                            v40 := GenS6{F0: -9223372036854775808, F1: 1}
                                            _ = len(v3)
                                            v41 := GenS7{F0: (v2 | 4294967295), F1: 1, F2: false}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                    case -9223372036854775808:
                                        if true {
                                            v42 := GenS1{F0: 32767, F1: (v2 % ((42) | 1)), F2: 0}
                                            if false {
                                                _ = len(v3)
                                                v43 := GenS3{F0: false, F1: 1, F2: 642}
                                                var v44 float32 = v7
                                                _ = v43
                                                _ = v44
                                            }
                                            v4 = v4
                                            _ = v42
                                        } else {
                                            v7 = ((((v7 + ((1e10 * v7) * v7)) * (v7 * (1e10 - 1.5))) * float32(-2.25)) + (float32((-2.25 - 1e10)) * v7))
                                            if (65535 > 661) {
                                                v45 := GenS8{F0: 0, F1: (fi36 / ((-945) | 1))}
                                                if false {
                                                    v4 = v4
                                                    var v46 float64 = (1.5 / ((0.0 + float64((float32(0.1) / v7))) * 0.0))
                                                    fmt.Println(v1)
                                                    _ = v46
                                                }
                                                _ = v45
                                            } else {
                                                if (1.5 >= 1e10) {
                                                    if (!(true && ((^(v4)) < (v4 >> (((^(v2))) & 63))))) {
                                                        v1 += v1
                                                    } else {
                                                        v3[(fi36 - (fi38 | fi36))] = fi38
                                                    }
                                                    switch fi38 {
                                                    case 32767:
                                                        if (1 == 255) {
                                                            var v47 int = fi36
                                                            var v48 float32 = (3.14159 * v7)
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                    case 128:
                                                        var v49 int16 = v4
                                                        fmt.Println(v49)
                                                        _ = v49
                                                    case -1:
                                                        if (0 < 0) {
                                                            fmt.Println(v2, v1, v7)
                                                        } else {
                                                            fmt.Println(fi38, v1)
                                                        }
                                                    }
                                                    var v50 uint8 = 42
                                                    _ = v50
                                                } else {
                                                    v51 := []int{(^(fi38)), 2147483647, -1, 9223372036854775807}
                                                    var v52 uint64 = 4294967295
                                                    fmt.Println(fi36, v4, v52)
                                                    _ = v51
                                                    _ = v52
                                                }
                                                v53 := []int{fi38, (fi38 % ((fi38) | 1))}
                                                _ = v53
                                            }
                                            v2 -= (v2 + 255)
                                        }
                                    }
                                    var v54 uint16 = 0
                                    v1 = (v1 + v1)
                                    _ = v54
                                } else {
                                    if true {
                                        if (v2 != 2) {
                                            var v55 int8 = (v1 + v1)
                                            fmt.Println(v2, v4)
                                            v2 -= (v2 + 135)
                                            _ = v55
                                        }
                                        v4 = v4
                                        v56 := GenS1{F0: 42, F1: (v2 / ((v2) | 1)), F2: 255}
                                        _ = v56
                                    }
                                }
                            case "\n\t":
                                if (!(-2147483648 > 1)) {
                                    v2 -= 720
                                } else {
                                    var v57 int64 = 0
                                    v58 := []int{32767, 1}
                                    _ = v57
                                    _ = v58
                                }
                            default:
                                switch "hello" {
                                case "a":
                                    var v59 float32 = 0.1
                                    _ = v59
                                case "hello":
                                    v3[-977] = (fi36 >> ((v2) & 63))
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        v60 := GenS2{F0: (v1 & v1), F1: 42, F2: 309}
                                        _ = v60
                                    case "hello":
                                        var v61 uint16 = 2
                                        _ = v61
                                    default:
                                        v2 += (^(v2))
                                    }
                                }
                            }
                            _ = fi38
                        }
                        _ = v37
                    } else {
                        var v62 uint = (v2 << ((v2) & 63))
                        var v63 uint8 = 88
                        if true {
                            v2 = v62
                            v1 = v1
                        } else {
                            for fi64 := 0; fi64 < 8; fi64++ {
                                if (591 >= 42) {
                                    var v65 uint16 = 2
                                    if false {
                                        v65 = v65
                                    }
                                    v1 += (v1 / (((v1 & (v1 + (v1 / ((int8(fi36)) | 1))))) | 1))
                                    _ = v65
                                } else {
                                    var v66 string = ""
                                    var v67 float64 = 1.5
                                    v7 += ((float32(v67) + (((1.5 / ((float32(1e10) * v7) / 0.1)) + v7) - v7)) * v7)
                                    _ = v66
                                    _ = v67
                                }
                                switch ("foo bar" + (("a" + "hello") + ("" + ("hello" + ("\n\t" + ""))))) {
                                case "foo bar":
                                    var v68 int32 = -1
                                    _ = v68
                                case "hello":
                                    v1 = (v1 + (^(v1)))
                                case "a":
                                    if (!true) {
                                        var v69 string = ("" + "foo bar")
                                        _ = v69
                                    }
                                }
                                var v70 bool = true
                                _ = fi64
                                _ = v70
                            }
                            var v71 float32 = (float32(1e10) / (v7 / float32(((1e10 + 0.0) - float64(3.14159)))))
                            _ = v71
                        }
                        _ = v62
                        _ = v63
                    }
                    _ = fi36
                }
            }
            _ = v27
        }
    }
    fmt.Println(v1)
    if false {
        v2 = (v2 | (v2 >> ((v2) & 63)))
        v72 := GenS1{F0: 32767, F1: v2, F2: 128}
        v73 := GenS6{F0: 262, F1: 42}
        _ = v72
        _ = v73
    } else {
        if ((v2 == (v2 + 255)) || true) {
            fmt.Println(v4, v2)
            switch "foo bar" {
            case "foo bar":
                v74 := GenS3{F0: (200 > 255), F1: v2, F2: 440}
                _ = v74
            default:
                v4 = v4
            }
        } else {
            v4 = -32768
        }
        var v75 uint64 = 0
        fmt.Println(v75, v2, v75)
        _ = v75
    }
    var v76 float64 = (0.0 * 1.5)
    if false {
        if (4294967295 < 127) {
            fmt.Println(v2, v76, v2)
            if true {
                if (!false) {
                    v1 = v1
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        for fi77 := 0; fi77 < 11; fi77++ {
                            var v78 string = "hello"
                            _ = fi77
                            _ = v78
                        }
                    }
                    if false {
                        v4 = (-(v4))
                    } else {
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            var v79 uint64 = 986
                            _ = v79
                        case "":
                            for fi80 := 0; fi80 < 11; fi80++ {
                                if (!(!(!false))) {
                                    if false {
                                        fmt.Println(v7)
                                    }
                                }
                                delete(v3, fi80)
                                _ = fi80
                            }
                        case "hello":
                            fmt.Println(v76, v2, v7)
                        default:
                            v1 -= 0
                        }
                        var v81 uint = v2
                        fmt.Println(v7, v2, v81)
                        _ = v81
                    }
                }
                var v82 int32 = 0
                _ = v82
            } else {
                v4 = 1
                var v83 float64 = 0.0
                _ = v83
            }
        } else {
            var v84 bool = true
            _ = v84
        }
    } else {
        v85 := GenS3{F0: (!(!(241 == 1))), F1: v2, F2: 9223372036854775807}
        _ = v85
    }
    switch 32768 {
    case 159:
        switch int(v4) {
        case 2147483647:
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "a":
                    for fi86 := 0; fi86 < 2; fi86++ {
                        if true {
                            var v87 int64 = 42
                            delete(v3, fi86)
                            _ = v87
                        }
                        _ = fi86
                    }
                }
            case "":
                if ((1 > int64(v4)) || true) {
                    for fi88 := 0; fi88 < 11; fi88++ {
                        var v89 uint64 = 42
                        v90 := GenS3{F0: false, F1: (v2 << ((uint(v89)) & 63)), F2: 0}
                        if true {
                            var v91 int32 = 48
                            delete(v3, 0)
                            _ = v91
                        } else {
                            for fi92 := 0; fi92 < 13; fi92++ {
                                if true {
                                    v1 = (v1 & v1)
                                    fmt.Println(v7, v89)
                                    fmt.Println(v7, v89)
                                } else {
                                    if (531 != 1) {
                                        if false {
                                            delete(v3, -727)
                                        } else {
                                            v93 := []int{4294967295}
                                            _ = v93
                                        }
                                        v3[(fi88 >> ((971) & 63))] = fi88
                                        v7 = (1.5 * float32(float64(float32(((v76 + ((float64(v7) / -2.25) + v76)) - 1e10)))))
                                        if false {
                                            v94 := GenS9{F0: 128}
                                            _ = v94
                                        }
                                    } else {
                                        v95 := v3[(fi88 | -329)]
                                        v96 := []int{(fi92 & int(v4)), (fi92 << ((v2) & 63)), fi88, v95}
                                        _ = v95
                                        _ = v96
                                    }
                                    v97 := GenS10{F0: 42, F1: false, F2: 0}
                                    var v98 uint = (v2 / ((v2) | 1))
                                    _ = v97
                                    _ = v98
                                }
                                v76 = (((v76 / ((1e10 / v76) * (float64(float32(3.14159)) + (v76 - v76)))) / (v76 - (v76 - v76))) / (float64((v7 - (-2.25 / (float32(v76) / ((1.5 / (v7 * 0.0)) - v7))))) * v76))
                                v4 += v4
                                _ = fi92
                            }
                            v99 := v3[420]
                            v1 = v1
                            _ = v99
                        }
                        _ = fi88
                        _ = v89
                        _ = v90
                    }
                }
            default:
                fmt.Println(v4, v1)
            }
        case -935:
            var v100 string = ("a" + (("a" + "foo bar") + (("\n\t" + "foo bar") + "\n\t")))
            _ = v100
        case -2147483648:
            var v101 float32 = ((v7 * (v7 - (float32(v76) - (0.1 / 1.0)))) + v7)
            _ = v101
        }
    case 0:
        fmt.Println(v2)
    case 1:
        var v102 uint32 = 697
        _ = v102
    }
    v103 := GenS5{F0: 42}
    fmt.Println(v4, v76)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v76
}
