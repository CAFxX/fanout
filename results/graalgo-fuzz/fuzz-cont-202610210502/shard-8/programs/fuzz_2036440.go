package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int64; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint32; F1 uint; F2 int32 }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 int16 = 545
    var v2 int8 = -128
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 2; fi4++ {
        v5 := GenS1{F0: 185}
        if (-14 > v2) {
            v6 := GenS2{F0: (v1 % ((-32768) | 1)), F1: 107, F2: 128}
            _ = v6
        } else {
            v2 -= (v2 | (^(v2)))
            v2 = (v2 << ((18446744073709551615) & 63))
        }
        for fi7 := 0; fi7 < 11; fi7++ {
            v3[("" + "hello")] = ("\n\t" + "a")
            _ = fi7
        }
        _ = fi4
        _ = v5
    }
    switch "\n\t" {
    case "\n\t":
        v2 = -116
    case "hello":
        if false {
            v1 = 42
            v2 = v2
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                if (!true) {
                    var v9 float64 = 1.5
                    for fi10 := 0; fi10 < 18; fi10++ {
                        switch "\n\t" {
                        case "hello":
                            _ = len(v3)
                        case "foo bar":
                            v11 := []int{}
                            _ = v11
                        }
                        if (861 <= -2147483648) {
                            switch (^(fi8)) {
                            case -200:
                                var v12 string = "hello"
                                _ = v12
                            default:
                                fmt.Println(fi10, v1, fi10)
                            }
                        } else {
                            v1 -= (v1 + v1)
                            v1 += (^(v1))
                        }
                        _ = fi10
                    }
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v2 = v2
                        v14 := GenS2{F0: 118, F1: 938, F2: (fi13 % ((127) | 1))}
                        v2 = -93
                        var v15 uint = 18446744073709551615
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    _ = v9
                }
                fmt.Println(v2, fi8)
                if (181 > 400) {
                    fmt.Println(fi8, v2)
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := []int{}
                            var v18 uint16 = 1
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                } else {
                    var v19 float64 = -2.25
                    _ = v19
                }
                _ = fi8
            }
        }
    default:
        if (v1 == v1) {
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 uint32 = 122
                var v22 uint = 255
                if false {
                    var v23 bool = (false || true)
                    for fi24 := 0; fi24 < 7; fi24++ {
                        v25 := GenS1{F0: 42}
                        v21 = 0
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                } else {
                    v21 -= v21
                    var v26 float32 = float32(1.5)
                    fmt.Println(v26, v21, v26)
                    _ = v26
                }
                _ = fi20
                _ = v21
                _ = v22
            }
            if false {
                v2 += -30
            }
        } else {
            fmt.Println(v1, v1)
            var v27 int16 = 42
            _ = v27
        }
    }
    v28 := []int{}
    for fi29 := 0; fi29 < 1; fi29++ {
        if true {
            v2 -= (v2 << ((2) & 63))
            v30 := GenS2{F0: v1, F1: 42, F2: 127}
            if false {
                v2 -= 1
                v31 := GenS2{F0: v1, F1: 175, F2: fi29}
                if (!false) {
                    if false {
                        fmt.Println(v2)
                        if (42 >= 4294967295) {
                            v2 = v2
                            var v32 int8 = v2
                            var v33 int = -179
                            _ = v32
                            _ = v33
                        }
                        v2 -= (v2 | (v2 & v2))
                    } else {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            switch fi29 {
                            case 32768:
                                delete(v3, "\n\t")
                            case 42:
                                if (9223372036854775807 == 747) {
                                    v2 = (v2 ^ (v2 << ((62) & 63)))
                                    v2 = v2
                                }
                            default:
                                var v35 uint = 118
                                _ = v35
                            }
                            _ = fi34
                        }
                        var v36 int32 = -793
                        var v37 int64 = 1
                        _ = v36
                        _ = v37
                    }
                    delete(v3, "\n\t")
                    v38 := GenS3{F0: 167, F1: v2}
                    _ = v38
                }
                _ = v31
            }
            _ = v30
        } else {
            v2 = 127
        }
        var v39 string = "foo bar"
        _ = fi29
        _ = v39
    }
    var v40 float64 = 1.5
    switch 1 {
    case 4294967295:
        for fi41 := 0; fi41 < 2; fi41++ {
            v42 := GenS2{F0: v1, F1: 697, F2: fi41}
            var v43 int = -367
            v40 -= v40
            _ = fi41
            _ = v42
            _ = v43
        }
    default:
        if (-2147483648 < -2147483648) {
            if true {
                fmt.Println(v2)
            }
            var v44 uint64 = 4294967295
            fmt.Println(v40, v2, v1)
            _ = v44
        } else {
            var v45 int = 926
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 17; fi46++ {
        var v47 uint = 255
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 8; fi48++ {
        for fi49 := 0; fi49 < 4; fi49++ {
            var v50 uint16 = 255
            fmt.Println(fi48, v50, v40)
            _ = fi49
            _ = v50
        }
        for fi51 := 0; fi51 < 12; fi51++ {
            var v52 int = (fi51 + -2147483648)
            if true {
                if (!false) {
                    v40 -= v40
                    var v53 int32 = -892
                    _ = v53
                }
            }
            _ = len(v28)
            switch "" {
            case "foo bar":
                var v54 int = fi48
                _ = v54
            case "a":
                if (1 != 32767) {
                    v40 = float64((float32(float64(((((float32(0.0) / (1e10 + 3.14159)) + 1.5) * float32(v40)) / 1.0))) / 1.0))
                    v52 = 673
                    v55 := GenS4{F0: 2147483647}
                    _ = v55
                } else {
                    if false {
                        v52 += (fi48 / ((fi48) | 1))
                        if (((42 > 2147483648) || ((true && false) || ((65535 > 122) || ((-(v1)) <= v1)))) && true) {
                            if len(v28) > 0 {
                                v28[0] = (-(fi51))
                            }
                            fmt.Println(v1)
                            _ = len(v28)
                        } else {
                            v56 := GenS5{F0: -1, F1: 255}
                            switch 1 {
                            case 371:
                                _ = len(v3)
                            default:
                                v57 := GenS6{F0: 1, F1: v52}
                                _ = v57
                            }
                            var v58 int8 = v2
                            v59 := GenS7{F0: -216, F1: 1, F2: (^(v1))}
                            var v60 uint16 = 11
                            _ = v56
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                        fmt.Println(v52, fi51)
                        fmt.Println(v2)
                        switch v52 {
                        case -1:
                            switch ("" + ("\n\t" + "a")) {
                            case "a":
                                v61 := v3["foo bar"]
                                _ = v61
                            case "foo bar":
                                fmt.Println(v40)
                                v62 := GenS8{F0: 2147483648, F1: 1, F2: 536}
                                _ = len(v3)
                                var v63 float64 = 3.14159
                                v52 += (fi48 << ((2) & 63))
                                _ = v62
                                _ = v63
                            case "":
                                v1 = -32768
                            }
                        case 32767:
                            var v64 string = "a"
                            if false {
                                fmt.Println(v52)
                                _ = len(v28)
                            }
                            _ = v64
                        }
                        var v65 int32 = 42
                        var v66 bool = true
                        _ = v65
                        _ = v66
                    }
                    var v67 uint = 18446744073709551615
                    var v68 float32 = (1.5 * -2.25)
                    _ = v67
                    _ = v68
                }
            case "":
                _ = len(v3)
            default:
                fmt.Println(v52, fi51)
            }
            switch "foo bar" {
            case "foo bar":
                v1 = (^(v1))
            default:
                var v69 uint8 = 50
                _ = v69
            }
            _ = len(v28)
            var v70 int = v52
            fmt.Println(v52)
            if true {
                var v71 int16 = (v1 * -281)
                var v72 float64 = 0.0
                v73 := GenS8{F0: 1, F1: 339, F2: 2147483647}
                _ = v71
                _ = v72
                _ = v73
            }
            fmt.Println(v52, v2)
            _ = fi51
            _ = v52
            _ = v70
        }
        _ = fi48
    }
    switch (("hello" + ("\n\t" + "\n\t")) + "") {
    case "a":
        v74 := GenS4{F0: 633}
        _ = v74
    case "foo bar":
        switch 128 {
        case 32768:
            v2 = v2
        case 4294967295:
            v40 = 3.14159
        default:
            if (1 >= 1) {
                switch 4294967295 {
                case 4294967295:
                    var v75 uint = 643
                    _ = v75
                case 128:
                    v2 += (v2 >> ((323) & 63))
                }
            }
        }
    case "\n\t":
        v76 := GenS1{F0: 1}
        _ = v76
    }
    switch 1 {
    case 9223372036854775807:
        for fi77 := 0; fi77 < 20; fi77++ {
            v2 -= (-(v2))
            if true {
                v78 := GenS1{F0: 1}
                v1 = 42
                v1 += v1
                _ = v78
            }
            v79 := GenS8{F0: 2, F1: 934, F2: 0}
            _ = fi77
            _ = v79
        }
    case -47:
        v80 := GenS9{F0: 158, F1: 26}
        _ = v80
    case -631:
        var v81 float32 = (1e10 / ((float32(v40) + (-2.25 - 1e10)) + ((3.14159 / 1.5) - 1e10)))
        _ = v81
    }
    var v82 uint = 65535
    fmt.Println(v2)
    v82 -= 4294967295
    v1 -= 1
    v83 := GenS10{F0: uint64(v82)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v40
    _ = v82
    _ = v83
}
