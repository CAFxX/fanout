package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int8; F1 int64; F2 int }

func main() {
    var v1 int64 = -993
    v2 := make(map[string]string)
    var v3 uint8 = 229
    var v4 int64 = v1
    v5 := make(map[uint32]int)
    var v6 int8 = 0
    fmt.Println(v1, v4, v4)
    v7 := GenS1{F0: 255}
    if (true && true) {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 65535}
            v10 := []int{(fi8 / ((fi8) | 1)), fi8}
            v11 := GenS1{F0: 2}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        v1 = v4
        for fi12 := 0; fi12 < 4; fi12++ {
            switch "foo bar" {
            case "a":
                v13 := []int{(fi12 >> ((2) & 63)), fi12}
                _ = v13
            case "":
                fmt.Println(fi12)
            default:
                v2[""] = "hello"
            }
            _ = fi12
        }
    } else {
        v1 = (v1 % ((9223372036854775807) | 1))
        v14 := GenS1{F0: 65535}
        _ = v14
    }
    fmt.Println(v1, v6)
    v15 := GenS1{F0: 65535}
    if ((false && (!true)) || false) {
        v3 += 127
    } else {
        var v16 int = 806
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS2{F0: v3, F1: uint32(v3)}
            if (!true) {
                switch (fi17 >> ((65535) & 63)) {
                case 890:
                    v16 = fi17
                case -888:
                    switch (v16 * v16) {
                    case 42:
                        v3 += (v3 << ((18446744073709551615) & 63))
                    default:
                        if (!(-2.25 >= 3.14159)) {
                            if true {
                                if true {
                                    var v19 uint16 = 408
                                    v3 += 2
                                    _ = v19
                                }
                                v20 := GenS2{F0: (v3 - 128), F1: uint32(v3)}
                                if (false && true) {
                                    switch ("\n\t" + "") {
                                    case "\n\t":
                                        fmt.Println(fi17, v16)
                                    case "":
                                        v21 := GenS3{F0: "foo bar"}
                                        _ = v21
                                    }
                                } else {
                                    v22 := GenS4{F0: v4}
                                    _ = v22
                                }
                                _ = v20
                            }
                            fmt.Println(v3, fi17, v4)
                            if (-152 != -1) {
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    v24 := []int{}
                                    var v25 float32 = 3.14159
                                    var v26 float32 = (v25 - float32(float64(float32((-2.25 * (-2.25 / float64(v25)))))))
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                for fi27 := 0; fi27 < 20; fi27++ {
                                    delete(v5, 4294967295)
                                    delete(v5, 2147483647)
                                    _ = fi27
                                }
                                fmt.Println(v6, v3, v3)
                            } else {
                                delete(v5, 0)
                                v5[1] = v16
                                switch "a" {
                                case "\n\t":
                                    switch (v16 % (((fi17 - (v16 << ((4294967295) & 63)))) | 1)) {
                                    case -184:
                                        delete(v2, "hello")
                                    case -1:
                                        var v28 uint64 = 18446744073709551615
                                        _ = v28
                                    case 174:
                                        var v29 float64 = -2.25
                                        _ = v29
                                    }
                                }
                            }
                        }
                    }
                case 911:
                    var v30 uint8 = v3
                    _ = v30
                default:
                    if ("foo bar" <= "\n\t") {
                        v3 = (v3 << ((18446744073709551615) & 63))
                        var v31 uint32 = 798
                        _ = v31
                    }
                }
                switch "\n\t" {
                case "a":
                    v32 := GenS3{F0: ""}
                    _ = v32
                case "\n\t":
                    var v33 string = ("foo bar" + ("" + ("foo bar" + "\n\t")))
                    _ = v33
                default:
                    if (65535 == 42) {
                        for fi34 := 0; fi34 < 1; fi34++ {
                            if (42 == 2147483647) {
                                v35 := []int{fi17, -311}
                                if (1.5 > float64(-2.25)) {
                                    v36 := GenS4{F0: v1}
                                    delete(v2, "a")
                                    if len(v35) > 0 {
                                        v35[0] = fi34
                                    }
                                    var v37 string = ("foo bar" + "")
                                    var v38 uint64 = 370
                                    _ = v36
                                    _ = v37
                                    _ = v38
                                } else {
                                    v2["foo bar"] = "a"
                                    if true {
                                        var v39 int = (-(v16))
                                        _ = v39
                                    } else {
                                        if true {
                                            var v40 uint64 = 472
                                            if true {
                                                fmt.Println(v3)
                                                var v41 int16 = 42
                                                var v42 uint = 2
                                                _ = v41
                                                _ = v42
                                            } else {
                                                var v43 int32 = 547
                                                var v44 bool = ((^(v6)) < v6)
                                                _ = len(v35)
                                                _ = v43
                                                _ = v44
                                            }
                                            _ = v40
                                        } else {
                                            v45 := v2["foo bar"]
                                            v46 := GenS1{F0: 65535}
                                            v5[42] = v16
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 int32 = 716
                                        var v48 int64 = v1
                                        fmt.Println(v48)
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                v49 := GenS4{F0: (-(v4))}
                                _ = v35
                                _ = v49
                            }
                            if false {
                                v3 = (v3 % ((42) | 1))
                                v50 := []int{}
                                v5[uint32(v3)] = (-(fi34))
                                _ = v50
                            }
                            _ = fi34
                        }
                        switch "\n\t" {
                        case "a":
                            v4 = v4
                        case "foo bar":
                            var v51 uint64 = 0
                            _ = v51
                        default:
                            for fi52 := 0; fi52 < 3; fi52++ {
                                v1 -= v1
                                _ = fi52
                            }
                        }
                        if false {
                            v53 := GenS1{F0: 522}
                            v6 -= (v6 % ((v6) | 1))
                            _ = v53
                        }
                    } else {
                        v54 := GenS2{F0: (v3 + v3), F1: 2147483648}
                        _ = v54
                    }
                }
            }
            fmt.Println(v1, v16, v6)
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    fmt.Println(v4, v3)
    v6 = (v6 % ((v6) | 1))
    if false {
        v1 += v1
        v3 = v3
    }
    v55 := GenS2{F0: v3, F1: 2147483647}
    v6 = v6
    v3 += (v3 & (v3 | v3))
    if (!false) {
        if false {
            var v56 uint = 631
            _ = v56
        } else {
            fmt.Println(v6, v6, v6)
            v6 = v6
            var v57 int64 = v4
            _ = v57
        }
        if (true && false) {
            switch -9223372036854775808 {
            case -9223372036854775808:
                v1 = v4
            }
            for fi58 := 0; fi58 < 17; fi58++ {
                v2["a"] = "foo bar"
                if true {
                    var v59 uint32 = 2
                    switch fi58 {
                    case -837:
                        if (((("foo bar" + ("" + "")) + (("hello" + ("foo bar" + ("hello" + "foo bar"))) + "foo bar")) + "a") > "a") {
                            for fi60 := 0; fi60 < 5; fi60++ {
                                fmt.Println(fi60, v3, v59)
                                fmt.Println(v3, v1)
                                _ = fi60
                            }
                            fmt.Println(v1, v4, v4)
                        }
                    case 4294967295:
                        v3 = 0
                    default:
                        v4 = (v1 & -9223372036854775808)
                    }
                    _ = v59
                }
                v1 = (v4 * 0)
                _ = fi58
            }
        } else {
            v6 = v6
            var v61 uint = 0
            v1 -= 42
            _ = v61
        }
        v62 := GenS5{F0: false, F1: 42}
        _ = v62
    } else {
        if (((3.14159 - -2.25) / 1.5) < float64(-2.25)) {
            var v63 string = "a"
            _ = v63
        } else {
            v4 += (^(v4))
            switch 32767 {
            case 505:
                switch -836 {
                case 0:
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v1 += v1
                        v6 = -60
                        _ = fi64
                    }
                case 32768:
                    v4 += v4
                default:
                    switch 2147483647 {
                    case 128:
                        var v65 float64 = 0.0
                        _ = v65
                    case -434:
                        v66 := []int{32768, 1, 884}
                        _ = v66
                    }
                }
            case -178:
                fmt.Println(v4, v6, v3)
            }
        }
        for fi67 := 0; fi67 < 4; fi67++ {
            fmt.Println(fi67)
            var v68 uint = 4294967295
            _ = fi67
            _ = v68
        }
        v1 = v1
    }
    v4 -= v1
    if (42 < 4294967295) {
        fmt.Println(v3, v6)
    } else {
        if ((0.1 + float64(3.14159)) >= ((1e10 - (1e10 - ((1e10 + float64((((1.5 + float32((3.14159 + -2.25))) * -2.25) * float32(float64((3.14159 - 0.1)))))) * (((0.1 + 0.0) * ((((-2.25 * 0.1) + 1e10) + 0.1) * -2.25)) + float64(0.0))))) * 0.0)) {
            v69 := GenS6{F0: 2, F1: ("foo bar" + ("a" + "a"))}
            v70 := []int{9223372036854775807}
            var v71 uint32 = 2
            _ = v69
            _ = v70
            _ = v71
        } else {
            if ((float64(-2.25) + 3.14159) != 0.0) {
                if true {
                    v72 := GenS1{F0: 65535}
                    v1 += (v4 | v1)
                    v73 := []int{416, 4294967295}
                    _ = v72
                    _ = v73
                } else {
                    var v74 bool = (false || (!false))
                    v75 := GenS4{F0: 57}
                    _ = v74
                    _ = v75
                }
                if false {
                    v76 := GenS2{F0: v3, F1: 4294967295}
                    v3 += 116
                    _ = v76
                } else {
                    for fi77 := 0; fi77 < 17; fi77++ {
                        v78 := GenS1{F0: 4294967295}
                        var v79 string = "\n\t"
                        v80 := GenS2{F0: (^(v3)), F1: 2}
                        _ = fi77
                        _ = v78
                        _ = v79
                        _ = v80
                    }
                    for fi81 := 0; fi81 < 9; fi81++ {
                        for fi82 := 0; fi82 < 17; fi82++ {
                            var v83 uint8 = v3
                            v84 := GenS7{F0: v6, F1: 255}
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                        var v85 int16 = 1000
                        var v86 uint8 = v3
                        _ = fi81
                        _ = v85
                        _ = v86
                    }
                    var v87 int64 = v1
                    _ = v87
                }
            }
        }
    }
    v88 := GenS8{F0: v6, F1: v1, F2: 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
    _ = v88
}
