package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 int8; F1 string; F2 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 bool; F1 int64 }
func genFn0() uint64 {
    v5 := []int{-2147483648}
    v6 := GenS1{F0: 217}
    _ = len(v5)
    var v7 uint8 = 255
    _ = v5
    _ = v6
    _ = v7
    return 695
}
func genFn1() int {
    for fi8 := 0; fi8 < 1; fi8++ {
        switch ("hello" + "") {
        case "hello":
            v9 := []int{}
            _ = v9
        }
        _ = fi8
    }
    var v10 uint32 = 273
    v10 += v10
    if true {
        v10 = v10
        v11 := []int{9223372036854775807, 771, 2147483647}
        v12 := GenS1{F0: 1}
        _ = v11
        _ = v12
    }
    if (-236 < -262) {
        v10 += 42
        var v13 float32 = (-2.25 * (0.1 / -2.25))
        var v14 bool = (false && ((0.0 + 1e10) < float64((v13 / v13))))
        _ = v13
        _ = v14
    }
    _ = v10
    return 32767
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = ("" != "hello")
    var v3 float64 = 1.5
    var v4 uint8 = 0
    v15 := GenS2{F0: 42, F1: 65535}
    if v2 {
        _ = len(v1)
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint16 = 0
                v3 -= v3
                _ = fi17
                _ = v18
            }
            _ = len(v1)
            var v19 uint16 = 280
            _ = fi16
            _ = v19
        }
        if (353 < -9223372036854775808) {
            fmt.Println(v3, v2)
        } else {
            v20 := []int{128, 128, -809}
            _ = len(v20)
            _ = v20
        }
    }
    fmt.Println(v2)
    v3 = (v3 - v3)
    if false {
        var v21 float32 = (1.5 - 0.1)
        v21 -= (v21 + v21)
        _ = v21
    } else {
        switch ("" + "") {
        case "a":
            v4 = 74
        default:
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 int16 = 1
                _ = fi22
                _ = v23
            }
        }
    }
    if v2 {
        for fi24 := 0; fi24 < 9; fi24++ {
            if v2 {
                for fi25 := 0; fi25 < 13; fi25++ {
                    v26 := GenS3{F0: 374, F1: (v4 % ((42) | 1)), F2: 343}
                    _ = fi25
                    _ = v26
                }
                var v27 float64 = ((0.1 * 0.0) - float64(0.1))
                _ = v27
            } else {
                if false {
                    v4 += (v4 + v4)
                    for fi28 := 0; fi28 < 3; fi28++ {
                        v29 := GenS3{F0: 42, F1: v4, F2: 296}
                        switch ((("\n\t" + ((("a" + "\n\t") + "a") + "a")) + ("hello" + "a")) + ("" + "")) {
                        case "":
                            fmt.Println(v2, fi28, fi28)
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            }
            v2 = (0 <= 482)
            _ = fi24
        }
        var v30 float32 = (0.1 - 1e10)
        _ = v30
    }
    if true {
        v31 := GenS3{F0: -9223372036854775808, F1: v4, F2: -439}
        _ = v31
    } else {
        var v32 uint64 = 4294967295
        for fi33 := 0; fi33 < 18; fi33++ {
            if (v2 || (v2 || v2)) {
                if false {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        fmt.Println(fi33, v32, fi33)
                        v4 = v4
                        _ = fi34
                    }
                    switch (-(fi33)) {
                    case 1:
                        if (255 >= 409) {
                            var v35 int16 = 899
                            v36 := v1[(fi33 / ((genFn1()) | 1))]
                            _ = v35
                            _ = v36
                        }
                    case -2147483648:
                        for fi37 := 0; fi37 < 5; fi37++ {
                            if (genFn1() <= (fi37 / ((4294967295) | 1))) {
                                var v38 uint32 = 42
                                if v2 {
                                    _ = len(v1)
                                    if v2 {
                                        if (v38 >= v38) {
                                            v39 := GenS4{F0: 42, F1: ""}
                                            _ = v39
                                        } else {
                                            var v40 float32 = float32((v3 / v3))
                                            switch (("hello" + "a") + "hello") {
                                            case "\n\t":
                                                v3 = v3
                                            }
                                            _ = v40
                                        }
                                        v41 := GenS5{F0: true, F1: 42}
                                        _ = v41
                                    } else {
                                        if true {
                                            var v42 uint16 = 42
                                            v43 := []int{fi33}
                                            _ = v42
                                            _ = v43
                                        }
                                    }
                                }
                                var v44 int16 = 1000
                                var v45 string = "hello"
                                delete(v1, (fi37 % (((fi33 * 183)) | 1)))
                                _ = v38
                                _ = v44
                                _ = v45
                            } else {
                                fmt.Println(v2, v2)
                                switch fi33 {
                                case 128:
                                    v2 = (v2 || v2)
                                case 2147483647:
                                    var v46 int16 = -32768
                                    _ = v46
                                case -693:
                                    v47 := GenS1{F0: 65535}
                                    _ = v47
                                default:
                                    if (!v2) {
                                        v48 := GenS1{F0: 961}
                                        _ = v48
                                    }
                                }
                            }
                            v49 := GenS2{F0: 32767, F1: 255}
                            _ = fi37
                            _ = v49
                        }
                    case -953:
                        if v2 {
                            var v50 float32 = 0.1
                            _ = v50
                        } else {
                            var v51 uint = 0
                            if true {
                                var v52 uint32 = 2147483648
                                if (v3 >= (0.1 - float64((0.0 * ((1e10 / 3.14159) - ((1.5 + -2.25) * 1e10)))))) {
                                    v3 = ((((v3 + (0.0 + (1.5 * v3))) - float64(0.1)) / ((float64((1.5 * float32(1.5))) * v3) + v3)) * 0.1)
                                    v32 = 42
                                    switch "a" {
                                    case "a":
                                        v2 = true
                                    default:
                                        var v53 float64 = (-2.25 - 1e10)
                                        _ = v53
                                    }
                                }
                                _ = v52
                            }
                            _ = v51
                        }
                    default:
                        fmt.Println(v4, fi33, v32)
                    }
                }
                v32 += v32
                v54 := GenS6{F0: -79, F1: "", F2: v4}
                _ = v54
            } else {
                v55 := GenS7{F0: "\n\t"}
                v3 = v3
                _ = v55
            }
            v3 -= v3
            v4 += v4
            _ = fi33
        }
        var v56 int32 = 0
        _ = v32
        _ = v56
    }
    v2 = v2
    if false {
        if ((v4 * (v4 & 0)) == 0) {
            v57 := GenS2{F0: 42, F1: 1}
            var v58 string = ""
            v59 := []int{-9223372036854775808, 32767}
            _ = v57
            _ = v58
            _ = v59
        }
    } else {
        v1[-902] = v2
        switch "a" {
        case "":
            v4 -= 255
        case "foo bar":
            for fi60 := 0; fi60 < 11; fi60++ {
                v61 := v1[(-(fi60))]
                switch "" {
                case "hello":
                    if true {
                        v2 = false
                        v1[genFn1()] = true
                    }
                case "\n\t":
                    v1[543] = true
                case "a":
                    if v61 {
                        for fi62 := 0; fi62 < 17; fi62++ {
                            if ((fi62 | -1) != genFn1()) {
                                v63 := GenS8{F0: v2, F1: -1}
                                switch fi62 {
                                case 1:
                                    var v64 float32 = 3.14159
                                    v4 -= (^(v4))
                                    _ = v64
                                case 42:
                                    var v65 int32 = -910
                                    _ = v65
                                default:
                                    if (v61 || v2) {
                                        fmt.Println(fi62, v61)
                                        v66 := []int{}
                                        if (v4 < v4) {
                                            if len(v66) >= 2 {
                                                v66 = v66[1:len(v66)]
                                            }
                                            fmt.Println(fi60)
                                            v4 = v4
                                        }
                                        _ = v66
                                    }
                                }
                                _ = v63
                            }
                            _ = fi62
                        }
                        fmt.Println(v4, v61)
                    }
                default:
                    for fi67 := 0; fi67 < 15; fi67++ {
                        v4 += v4
                        v68 := GenS9{F0: (42 == 42)}
                        var v69 int64 = 981
                        switch fi67 {
                        case 2147483647:
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                v2 = (42 < int16(v4))
                            }
                        case 0:
                            var v70 int = (fi60 - genFn1())
                            _ = v70
                        case -273:
                            v61 = (65535 < 531)
                        default:
                            v71 := GenS10{F0: 4294967295, F1: 42}
                            _ = v71
                        }
                        switch fi60 {
                        case 1:
                            v61 = v2
                        case 986:
                            if (2 <= (v4 << ((0) & 63))) {
                                if (v2 && false) {
                                    switch "\n\t" {
                                    case "a":
                                        v2 = (1000 >= 32767)
                                    case "\n\t":
                                        v3 += 0.0
                                    }
                                    _ = len(v1)
                                }
                                v61 = v61
                            }
                        case -515:
                            var v72 uint64 = 42
                            _ = v72
                        default:
                            var v73 int8 = 42
                            v69 += (-(v69))
                            _ = v73
                        }
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                }
                v74 := []int{}
                _ = fi60
                _ = v61
                _ = v74
            }
        default:
            var v75 int16 = 1
            _ = v75
        }
    }
    if false {
        v76 := []int{}
        v3 += v3
        _ = v76
    }
    fmt.Println(v3)
    if v2 {
        var v77 uint32 = 65535
        var v78 int64 = -812
        v4 = 127
        _ = v77
        _ = v78
    } else {
        for fi79 := 0; fi79 < 10; fi79++ {
            v4 = v4
            for fi80 := 0; fi80 < 11; fi80++ {
                var v81 uint32 = 1
                if false {
                    v82 := GenS11{F0: ((float64(3.14159) / v3) < v3), F1: 442}
                    if (!(943 < 2)) {
                        if v2 {
                            v83 := GenS11{F0: (!true), F1: -1}
                            v84 := GenS2{F0: 1, F1: 2}
                            v85 := []int{}
                            _ = v83
                            _ = v84
                            _ = v85
                        } else {
                            v81 = (v81 * (^(v81)))
                        }
                        v3 = v3
                        v81 += v81
                    }
                    v2 = (1 > 0)
                    _ = v82
                } else {
                    switch fi79 {
                    case -909:
                        var v86 int32 = -1
                        _ = v86
                    case 32767:
                        var v87 float64 = v3
                        _ = v87
                    }
                    v3 += v3
                    var v88 uint = 86
                    _ = v88
                }
                _ = fi80
                _ = v81
            }
            _ = fi79
        }
        var v89 int16 = 1000
        v4 = 128
        _ = v89
    }
    switch 2147483647 {
    case -2147483648:
        v90 := GenS6{F0: 0, F1: "a", F2: 127}
        _ = v90
    }
    fmt.Println(v3, v2)
    var v91 uint32 = 2147483647
    var v92 int32 = 877
    var v93 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v91:", v91)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v91
    _ = v92
    _ = v93
}
