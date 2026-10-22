package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 uint8; F1 uint }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int; F1 uint; F2 uint16 }
type GenS10 struct { F0 int; F1 uint32; F2 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string }
func genFn0() int64 {
    v5 := GenS1{F0: 18446744073709551615, F1: 100}
    for fi6 := 0; fi6 < 20; fi6++ {
        if (!(21 <= 47)) {
            var v7 int16 = 32767
            if (!true) {
                switch "hello" {
                case "a":
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 bool = true
                        v10 := GenS1{F0: 886, F1: 0}
                        var v11 int = (^(fi6))
                        _ = fi8
                        _ = v10
                        _ = v11
                        _ = v9
                    }
                case "\n\t":
                    v12 := []int{fi6, (fi6 % ((fi6) | 1)), (fi6 / ((-219) | 1)), fi6}
                    _ = v12
                case "foo bar":
                    fmt.Println(fi6, v7)
                }
            } else {
                v7 = v7
                switch "hello" {
                case "\n\t":
                    var v13 float32 = float32(3.14159)
                    _ = v13
                case "":
                    v14 := []int{4294967295}
                    _ = v14
                case "hello":
                    fmt.Println(v7)
                default:
                    v15 := []int{}
                    _ = v15
                }
                var v16 int = fi6
                _ = v16
            }
            var v17 int16 = v7
            _ = v17
            _ = v7
        } else {
            v18 := GenS1{F0: 18446744073709551615, F1: -1}
            switch "foo bar" {
            case "foo bar":
                var v19 uint16 = 42
                _ = v19
            case "":
                for fi20 := 0; fi20 < 4; fi20++ {
                    v21 := GenS2{F0: 42, F1: 550}
                    switch (fi20 >> ((1) & 63)) {
                    case 201:
                        var v22 uint16 = 0
                        _ = v22
                    default:
                        if true {
                            var v23 int16 = -775
                            _ = v23
                        } else {
                            var v24 int8 = 47
                            _ = v24
                        }
                    }
                    switch (fi20 / (((fi6 >> ((65535) & 63))) | 1)) {
                    case 32767:
                        var v25 int32 = -1
                        _ = v25
                    default:
                        v26 := []int{(fi6 >> ((255) & 63)), fi6}
                        _ = v26
                    }
                    switch 0 {
                    case 32768:
                        var v27 string = ("hello" + "\n\t")
                        _ = v27
                    case 9223372036854775807:
                        var v28 int32 = -297
                        _ = v28
                    }
                    _ = fi20
                    _ = v21
                }
            case "hello":
                var v29 int8 = 1
                _ = v29
            }
            _ = v18
        }
        _ = fi6
    }
    _ = v5
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 2147483648
    v2 := make(map[uint32]bool)
    var v3 uint8 = 128
    var v4 uint16 = 726
    var v30 uint32 = 365
    v2[uint32(v4)] = true
    for fi31 := 0; fi31 < 12; fi31++ {
        if (!(34 < 65535)) {
            fmt.Println(v3, v1)
            if false {
                switch "\n\t" {
                case "":
                    v32 := []int{(fi31 - fi31), 0, 1}
                    _ = v32
                default:
                    fmt.Println(v1)
                }
                switch (fi31 & 4294967295) {
                case 42:
                    if false {
                        fmt.Println(v3, v1)
                        var v33 int64 = -409
                        var v34 int8 = 123
                        _ = v33
                        _ = v34
                    }
                case -728:
                    v30 += v1
                case 32768:
                    if ((5 <= 1) && (892 <= 32767)) {
                        v35 := GenS1{F0: 766, F1: 127}
                        _ = v35
                    } else {
                        switch fi31 {
                        case 128:
                            v36 := []int{(fi31 * fi31), fi31, (fi31 % (((fi31 << ((0) & 63))) | 1)), fi31}
                            _ = v36
                        case -2147483648:
                            for fi37 := 0; fi37 < 15; fi37++ {
                                if (true && true) {
                                    v38 := GenS2{F0: 0, F1: -32768}
                                    if ((fi31 | (fi37 | (fi31 << ((4294967295) & 63)))) > fi37) {
                                        fmt.Println(fi31, fi37)
                                        var v39 uint = 4294967295
                                        v1 -= v1
                                        _ = v39
                                    } else {
                                        v40 := GenS3{F0: v30}
                                        v41 := []int{}
                                        var v42 uint64 = uint64(v4)
                                        _ = v40
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = v38
                                } else {
                                    v43 := GenS4{F0: v3, F1: (^(v1))}
                                    v30 = (v1 | v30)
                                    if (213 <= 107) {
                                        v1 = 870
                                    }
                                    _ = v43
                                }
                                v44 := []int{280, (fi31 - fi37), (fi37 - fi37)}
                                v4 = (^(v4))
                                _ = fi37
                                _ = v44
                            }
                        default:
                            v45 := GenS1{F0: 990, F1: 42}
                            _ = v45
                        }
                        fmt.Println(v4, v4, v3)
                    }
                }
                if (-149 != 9223372036854775807) {
                    var v46 uint16 = (v4 >> ((255) & 63))
                    var v47 bool = (3.14159 >= float32((-2.25 * 0.0)))
                    if v47 {
                        v48 := GenS5{F0: v3, F1: 255}
                        for fi49 := 0; fi49 < 5; fi49++ {
                            v50 := GenS5{F0: (^(v3)), F1: 1}
                            var v51 uint = 255
                            v30 = v30
                            fmt.Println(v30, v46, v47)
                            fmt.Println(fi49, v47, v47)
                            fmt.Println(v3)
                            var v52 float32 = 1e10
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                        v53 := GenS6{F0: 86, F1: (fi31 ^ (fi31 | fi31))}
                        _ = v48
                        _ = v53
                    } else {
                        if true {
                            v54 := GenS5{F0: v3, F1: 43}
                            if true {
                                fmt.Println(v4, v4)
                                v30 = (^(v1))
                            }
                            _ = v54
                        } else {
                            switch ("\n\t" + ("foo bar" + "")) {
                            case "foo bar":
                                switch "hello" {
                                case "foo bar":
                                    switch ("foo bar" + "hello") {
                                    case "\n\t":
                                        v30 += uint32(v3)
                                    case "":
                                        switch "\n\t" {
                                        case "a":
                                            for fi55 := 0; fi55 < 8; fi55++ {
                                                v30 = (v1 | v1)
                                                _ = fi55
                                            }
                                        case "foo bar":
                                            for fi56 := 0; fi56 < 8; fi56++ {
                                                v57 := GenS7{F0: 0}
                                                _ = fi56
                                                _ = v57
                                            }
                                        case "":
                                            var v58 int64 = -100
                                            _ = v58
                                        }
                                    default:
                                        if (4294967295 == 0) {
                                            var v59 uint32 = v30
                                            var v60 int = 42
                                            _ = v59
                                            _ = v60
                                        } else {
                                            switch (fi31 - (fi31 / ((fi31) | 1))) {
                                            case 127:
                                                v3 = (^(v3))
                                            default:
                                                var v61 int16 = 537
                                                _ = v61
                                            }
                                        }
                                    }
                                }
                            case "":
                                for fi62 := 0; fi62 < 1; fi62++ {
                                    if false {
                                        fmt.Println(v30)
                                        v63 := GenS8{F0: 42, F1: 5, F2: (v30 | 2147483647)}
                                        _ = v63
                                    } else {
                                        switch (("" + "hello") + "\n\t") {
                                        case "a":
                                            v64 := GenS6{F0: -765, F1: fi31}
                                            _ = v64
                                        case "":
                                            var v65 float32 = 0.0
                                            _ = v65
                                        }
                                        if false {
                                            v3 -= (v3 + (^(v3)))
                                        } else {
                                            if false {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    fmt.Println(v30, fi31, v46)
                                                default:
                                                    fmt.Println(v4, fi62, v46)
                                                }
                                            }
                                        }
                                    }
                                    v30 = (^(v1))
                                    var v66 int16 = -32768
                                    _ = fi62
                                    _ = v66
                                }
                            case "hello":
                                fmt.Println(v4)
                            }
                            if v47 {
                                fmt.Println(v4, fi31, fi31)
                                var v67 int32 = 0
                                v67 = v67
                                _ = v67
                            } else {
                                v68 := GenS2{F0: 1, F1: int16(v3)}
                                _ = v68
                            }
                        }
                    }
                    _ = v46
                    _ = v47
                }
            }
        }
        _ = fi31
    }
    var v69 int16 = 1
    var v70 uint8 = 128
    var v71 int16 = 1000
    if false {
        v70 -= (v3 + v3)
    } else {
        switch 32767 {
        case 9223372036854775807:
            var v72 int = 723
            _ = v72
        default:
            fmt.Println(v69, v4)
        }
        v73 := []int{42, 128, -647}
        if false {
            for fi74 := 0; fi74 < 15; fi74++ {
                v70 -= v3
                _ = fi74
            }
        }
        _ = v73
    }
    var v75 float64 = (3.14159 * 1e10)
    _ = len(v2)
    var v76 int16 = (v71 + (v69 | v71))
    v77 := []int{690, -9223372036854775808, 696, -943}
    fmt.Println(v30)
    if true {
        var v78 int16 = (v71 << ((0) & 63))
        var v79 uint16 = (v4 | (v4 / (((v4 * v4)) | 1)))
        var v80 int64 = 0
        _ = v78
        _ = v79
        _ = v80
    } else {
        if (719 < 4294967295) {
            var v81 uint32 = (v30 + (v30 >> ((42) & 63)))
            v71 += (v76 << ((734) & 63))
            _ = v81
        } else {
            v82 := GenS3{F0: 2147483647}
            v77 = append(v77, 1)
            fmt.Println(v69, v76)
            _ = v82
        }
        for fi83 := 0; fi83 < 17; fi83++ {
            switch fi83 {
            case 4294967295:
                if (9223372036854775807 == 9223372036854775807) {
                    if ((v70 != (v70 * uint8(v4))) || true) {
                        switch fi83 {
                        case 4294967295:
                            v3 -= (v70 ^ v3)
                        default:
                            var v84 uint = 0
                            _ = v84
                        }
                        if (1.5 < -2.25) {
                            _ = len(v2)
                            var v85 uint64 = 51
                            v30 += 42
                            _ = v85
                        } else {
                            v4 = 1
                        }
                        var v86 float64 = v75
                        _ = v86
                    } else {
                        var v87 int32 = -484
                        _ = v87
                    }
                    switch (fi83 ^ 32768) {
                    case 1:
                        v88 := GenS7{F0: 42}
                        _ = v88
                    case 128:
                        v89 := GenS9{F0: fi83, F1: 255, F2: 0}
                        _ = v89
                    default:
                        delete(v2, v30)
                    }
                    fmt.Println(fi83, v71)
                }
            case 230:
                var v90 uint = 18446744073709551615
                _ = v90
            default:
                for fi91 := 0; fi91 < 5; fi91++ {
                    if false {
                        v77 = append(v77, fi91)
                    } else {
                        _ = len(v77)
                        var v92 uint16 = (v4 * (v4 ^ 255))
                        _ = v92
                    }
                    if len(v77) >= 2 {
                        v77 = v77[1:len(v77)]
                    }
                    v71 = 32767
                    _ = fi91
                }
            }
            v93 := GenS10{F0: -537, F1: (^(v30)), F2: 65535}
            for fi94 := 0; fi94 < 8; fi94++ {
                v70 -= 127
                _ = fi94
            }
            _ = fi83
            _ = v93
        }
        v95 := GenS11{F0: -9223372036854775808}
        _ = v95
    }
    v96 := GenS12{F0: "\n\t"}
    v4 = 815
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("slicelen v77:", len(v77), v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v69
    _ = v70
    _ = v71
    _ = v75
    _ = v76
    _ = v77
    _ = v96
}
