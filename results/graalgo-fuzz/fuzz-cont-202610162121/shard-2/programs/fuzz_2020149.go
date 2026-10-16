package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint; F2 int }
type GenS5 struct { F0 string; F1 uint32; F2 bool }
type GenS6 struct { F0 bool; F1 int8; F2 int32 }
type GenS7 struct { F0 int64; F1 uint64; F2 uint }
type GenS8 struct { F0 string; F1 uint }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 241
    var v3 uint8 = (v1 >> ((4294967295) & 63))
    var v4 uint16 = 42
    var v5 int32 = -2147483648
    var v6 int8 = 46
    v6 = (v6 ^ (v6 | 100))
    var v7 string = ""
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: v4}
        _ = fi8
        _ = v9
    }
    if (!(v7 >= v7)) {
        if (32767 >= 409) {
            var v10 uint64 = 4294967295
            var v11 string = "\n\t"
            var v12 int = 42
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if (0 <= 42) {
        var v13 float32 = ((1e10 - 3.14159) / 3.14159)
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        v15 := []int{(fi14 + fi14), fi14}
        for fi16 := 0; fi16 < 10; fi16++ {
            v7 = v7
            var v17 int64 = int64(v2)
            _ = fi16
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    v1 -= 128
    var v18 int = int(v4)
    var v19 int64 = 9223372036854775807
    v20 := []int{(-(v18)), v18}
    switch v18 {
    case -981:
        v21 := GenS1{F0: (v4 & (v4 / ((830) | 1)))}
        _ = v21
    default:
        v19 = v19
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        if ((-(v5)) > 42) {
            if true {
                v6 = (^(v6))
            }
        } else {
            switch "hello" {
            case "foo bar":
                v23 := GenS1{F0: 0}
                _ = v23
            case "hello":
                var v24 int64 = v19
                _ = v24
            case "\n\t":
                v25 := GenS2{F0: 2, F1: 0}
                _ = v25
            }
            var v26 int = 0
            _ = v26
        }
        var v27 float64 = 0.0
        if false {
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 uint64 = 0
                if (true && (v19 == 9223372036854775807)) {
                    fmt.Println(fi28, v4, fi22)
                    v30 := GenS1{F0: v4}
                    _ = v30
                }
                fmt.Println(fi28, fi22)
                _ = fi28
                _ = v29
            }
            v20 = append(v20, v18)
        } else {
            switch 4294967295 {
            case -672:
                fmt.Println(v6)
            case 32768:
                var v31 string = ((v7 + v7) + (("hello" + v7) + v7))
                _ = v31
            case 128:
                v32 := GenS1{F0: uint16(v1)}
                _ = v32
            default:
                var v33 int16 = 0
                _ = v33
            }
        }
        _ = fi22
        _ = v27
    }
    v34 := GenS2{F0: 4294967295, F1: uint32(v1)}
    for fi35 := 0; fi35 < 12; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS1{F0: v4}
            fmt.Println(v3, v6, fi35)
            _ = fi36
            _ = v37
        }
        var v38 float64 = (1e10 + (1.5 + (((1e10 / 1.0) / (float64(float32(3.14159)) + 1e10)) * -2.25)))
        v39 := GenS1{F0: v4}
        _ = fi35
        _ = v38
        _ = v39
    }
    v40 := GenS3{F0: (v4 & (v4 % ((v4) | 1))), F1: 2147483647}
    if (uint64(v1) <= v2) {
        switch "hello" {
        case "foo bar":
            if (1e10 <= 3.14159) {
                for fi41 := 0; fi41 < 9; fi41++ {
                    fmt.Println(v1, v7, fi41)
                    _ = fi41
                }
                var v42 bool = true
                if ((v19 / ((42) | 1)) >= int64(v4)) {
                    v43 := GenS1{F0: v4}
                    _ = v43
                } else {
                    if len(v20) > 0 {
                        v20[0] = (^(v18))
                    }
                    if v42 {
                        var v44 uint = 18446744073709551615
                        v6 += (v6 << (((v44 | 42)) & 63))
                        var v45 int64 = (v19 << ((2) & 63))
                        _ = v44
                        _ = v45
                    }
                }
                _ = v42
            } else {
                v46 := GenS3{F0: (v4 % ((806) | 1)), F1: 0}
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                _ = v46
            }
        case "\n\t":
            v47 := GenS4{F0: 74, F1: 766, F2: 32767}
            _ = v47
        case "a":
            var v48 int8 = v6
            _ = v48
        }
        v20 = append(v20, int(v4))
    } else {
        v1 -= v1
        v20 = append(v20, -656)
        if true {
            for fi49 := 0; fi49 < 12; fi49++ {
                if false {
                    switch (v18 ^ -9223372036854775808) {
                    case 4294967295:
                        if len(v20) > 0 {
                            v20[0] = 1000
                        }
                    case 9223372036854775807:
                        if (((0.0 - (((3.14159 / -2.25) * (1.5 / 1e10)) + float32(-2.25))) + -2.25) < ((3.14159 / 1.0) - (0.0 - (1.5 - 1.5)))) {
                            v1 = (v3 + 1)
                            switch ((("\n\t" + v7) + v7) + v7) {
                            case "foo bar":
                                var v50 uint32 = 2147483648
                                _ = v50
                            case "":
                                var v51 float32 = -2.25
                                _ = v51
                            default:
                                switch ("hello" + "a") {
                                case "a":
                                    v7 = v7
                                }
                            }
                        }
                    }
                    var v52 int16 = 1
                    _ = v52
                } else {
                    fmt.Println(v2, fi49, v18)
                }
                _ = fi49
            }
            for fi53 := 0; fi53 < 13; fi53++ {
                for fi54 := 0; fi54 < 15; fi54++ {
                    if false {
                        v55 := GenS1{F0: 150}
                        _ = v55
                    }
                    v18 = 127
                    v56 := GenS5{F0: v7, F1: 1, F2: false}
                    _ = fi54
                    _ = v56
                }
                _ = fi53
            }
            switch "" {
            case "":
                for fi57 := 0; fi57 < 11; fi57++ {
                    v19 = v19
                    _ = fi57
                }
            case "foo bar":
                var v58 uint64 = 0
                _ = v58
            }
        } else {
            v5 -= 1
            v19 = (v19 << ((65535) & 63))
            fmt.Println(v7, v4)
        }
    }
    if ((1e10 / 1.0) <= -2.25) {
        switch v7 {
        case "foo bar":
            v2 = 142
        case "a":
            if (!true) {
                for fi59 := 0; fi59 < 15; fi59++ {
                    if (int16(v6) < 141) {
                        v60 := GenS6{F0: true, F1: v6, F2: (-(v5))}
                        _ = v60
                    } else {
                        v61 := GenS1{F0: v4}
                        switch v18 {
                        case 4294967295:
                            v18 = v18
                        case 32767:
                            switch fi59 {
                            case 4294967295:
                                if ((v18 <= (fi59 - 128)) || (!("" == "a"))) {
                                    v62 := GenS1{F0: v4}
                                    var v63 int16 = -32768
                                    for fi64 := 0; fi64 < 15; fi64++ {
                                        fmt.Println(v7)
                                        v63 = 0
                                        _ = fi64
                                    }
                                    _ = v62
                                    _ = v63
                                } else {
                                    switch ("\n\t" + (v7 + v7)) {
                                    case "foo bar":
                                        var v65 int64 = (v19 - v19)
                                        _ = v65
                                    case "hello":
                                        if false {
                                            v66 := GenS5{F0: v7, F1: 655, F2: false}
                                            for fi67 := 0; fi67 < 4; fi67++ {
                                                if false {
                                                    var v68 float64 = ((0.0 / 1.0) + (0.1 * 3.14159))
                                                    _ = v68
                                                }
                                                if (!(!(false && (642 >= 384)))) {
                                                    var v69 float64 = (-2.25 - 1.5)
                                                    _ = v69
                                                }
                                                _ = fi67
                                            }
                                            _ = v66
                                        } else {
                                            if ((42 == 4294967295) && true) {
                                                _ = len(v20)
                                                var v70 uint16 = (v4 & v4)
                                                _ = v70
                                            } else {
                                                v71 := GenS7{F0: v19, F1: 2, F2: 18446744073709551615}
                                                _ = v71
                                            }
                                            v72 := GenS3{F0: 0, F1: 399}
                                            _ = len(v20)
                                            _ = v72
                                        }
                                    case "a":
                                        if false {
                                            var v73 uint32 = 42
                                            for fi74 := 0; fi74 < 2; fi74++ {
                                                if ((!(false && (v4 <= v4))) && true) {
                                                    switch v7 {
                                                    case "hello":
                                                        var v75 string = ("foo bar" + v7)
                                                        _ = v75
                                                    case "":
                                                        v73 -= uint32(v1)
                                                    }
                                                } else {
                                                    _ = len(v20)
                                                    v2 = v2
                                                }
                                                fmt.Println(v2, v4, v2)
                                                var v76 float32 = (1.5 * (-2.25 - 3.14159))
                                                if true {
                                                    fmt.Println(v76)
                                                }
                                                _ = fi74
                                                _ = v76
                                            }
                                            var v77 int32 = v5
                                            _ = v73
                                            _ = v77
                                        }
                                    }
                                    var v78 uint = 4294967295
                                    _ = v78
                                }
                            default:
                                fmt.Println(v3)
                            }
                        case -722:
                            fmt.Println(v3, v7)
                        }
                        _ = v61
                    }
                    v2 += 1
                    _ = fi59
                }
                _ = len(v20)
            }
        }
        if false {
            v79 := GenS6{F0: true, F1: v6, F2: v5}
            v80 := GenS8{F0: v7, F1: 42}
            _ = v79
            _ = v80
        } else {
            var v81 int64 = v19
            fmt.Println(v1)
            _ = v81
        }
        var v82 bool = (!true)
        _ = v82
    } else {
        var v83 int8 = v6
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
