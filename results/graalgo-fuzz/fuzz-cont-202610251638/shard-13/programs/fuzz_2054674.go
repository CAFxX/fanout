package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint64; F1 int8; F2 uint8 }
func genFn0(p0 uint, p1 string, p2 uint) int8 {
    if false {
        var v5 int8 = 100
        _ = v5
    } else {
        if ((0.1 - 0.0) != (float32(3.14159) + (3.14159 * ((0.1 / 0.1) * float32(3.14159))))) {
            p0 = 1
        }
        v6 := []int{-9223372036854775808}
        _ = v6
    }
    var v7 float64 = 0.1
    for fi8 := 0; fi8 < 13; fi8++ {
        p1 = ("foo bar" + "\n\t")
        if (65535 != 4294967295) {
            if (42 != -1) {
                switch p1 {
                case "":
                    if true {
                        switch -597 {
                        case -1:
                            v9 := []int{-2147483648}
                            _ = v9
                        default:
                            var v10 int64 = 42
                            _ = v10
                        }
                    } else {
                        for fi11 := 0; fi11 < 3; fi11++ {
                            if (-110 > 18) {
                                var v12 int64 = 0
                                v13 := []int{0, fi8}
                                _ = v12
                                _ = v13
                            } else {
                                if (!(102 != 699)) {
                                    if (-32768 >= 0) {
                                        var v14 int32 = 2147483647
                                        v15 := []int{-1}
                                        _ = v14
                                        _ = v15
                                    }
                                    if true {
                                        p2 += 2
                                    }
                                }
                            }
                            p0 = 389
                            _ = fi11
                        }
                        if false {
                            var v16 int16 = -1
                            var v17 int8 = 7
                            if false {
                                v16 = (^(v16))
                                for fi18 := 0; fi18 < 1; fi18++ {
                                    p1 = (((p1 + "\n\t") + p1) + "a")
                                    v19 := []int{}
                                    _ = fi18
                                    _ = v19
                                }
                            } else {
                                var v20 uint32 = 544
                                _ = v20
                            }
                            _ = v16
                            _ = v17
                        }
                    }
                }
                p2 -= 42
                var v21 int = fi8
                _ = v21
            } else {
                v22 := []int{}
                _ = v22
            }
            v23 := GenS1{F0: 818, F1: fi8, F2: ((p1 + ("\n\t" + "a")) + "hello")}
            _ = v23
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 127
}

func main() {
    var v1 uint = 42
    var v2 uint8 = 59
    var v3 int16 = -823
    v4 := make(map[string]int)
    if (!(!true)) {
        switch ("a" + "a") {
        case "":
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 int64 = 0
                v4["foo bar"] = fi24
                if (!true) {
                    v1 = (v1 << ((v1) & 63))
                    v26 := GenS2{F0: 4294967295}
                    v27 := GenS2{F0: 2147483648}
                    _ = v26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        case "a":
            for fi28 := 0; fi28 < 2; fi28++ {
                v4["a"] = fi28
                v29 := GenS1{F0: (v1 & (^(v1))), F1: (fi28 - fi28), F2: "\n\t"}
                _ = fi28
                _ = v29
            }
        case "foo bar":
            v1 = 378
        default:
            v1 += 132
        }
        v30 := GenS3{F0: (v2 - (v2 & (^(v2))))}
        var v31 float32 = -2.25
        _ = v30
        _ = v31
    } else {
        fmt.Println(v3, v2, v3)
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint64 = uint64(v1)
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 uint32 = 65535
            if (!(v35 < (v35 + (^(v35))))) {
                switch -9223372036854775808 {
                case 2147483647:
                    v33 += (v33 / ((18446744073709551615) | 1))
                }
            }
            if (128 == v2) {
                v2 += uint8(v33)
            }
            _ = fi34
            _ = v35
        }
        switch fi32 {
        case -563:
            var v36 string = "a"
            _ = v36
        default:
            var v37 uint16 = 32767
            _ = v37
        }
        _ = fi32
        _ = v33
    }
    for fi38 := 0; fi38 < 12; fi38++ {
        v1 += (^(v1))
        v39 := v4[""]
        _ = fi38
        _ = v39
    }
    v1 += v1
    if ((float32(1.5) >= 1.5) && (0.1 != float32(float64(((1.5 / 1.0) / (0.1 / float32(3.14159))))))) {
        fmt.Println(v3, v2, v1)
    } else {
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            fmt.Println(v3, v1, v1)
        }
        for fi40 := 0; fi40 < 19; fi40++ {
            switch fi40 {
            case 1:
                v41 := GenS4{F0: 865, F1: 1}
                _ = v41
            case 259:
                v2 = (v2 + 255)
            case 128:
                v42 := GenS5{F0: "\n\t"}
                _ = v42
            }
            v3 += (v3 % ((805) | 1))
            _ = fi40
        }
        v43 := []int{128}
        _ = v43
    }
    var v44 int16 = v3
    if (991 < 42) {
        v45 := []int{-974, 584}
        _ = v45
    }
    switch "foo bar" {
    case "a":
        for fi46 := 0; fi46 < 19; fi46++ {
            v2 = v2
            _ = fi46
        }
    case "":
        var v47 string = "\n\t"
        _ = v47
    case "foo bar":
        for fi48 := 0; fi48 < 16; fi48++ {
            switch "\n\t" {
            case "":
                switch "foo bar" {
                case "":
                    for fi49 := 0; fi49 < 13; fi49++ {
                        v44 = (v44 & v3)
                        if (40 == 63) {
                            switch fi48 {
                            case 1:
                                var v50 int32 = 49
                                _ = v50
                            }
                            v51 := GenS2{F0: 2}
                            _ = v51
                        }
                        _ = fi49
                    }
                }
            case "a":
                fmt.Println(fi48, v44)
            case "\n\t":
                switch fi48 {
                case -864:
                    if false {
                        v52 := GenS6{F0: fi48, F1: genFn0(42, "a", 255)}
                        for fi53 := 0; fi53 < 11; fi53++ {
                            v54 := GenS3{F0: v2}
                            fmt.Println(v44, v3)
                            var v55 uint8 = v2
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        v1 = (v1 >> ((v1) & 63))
                        _ = v52
                    } else {
                        v1 -= v1
                        switch "\n\t" {
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                for fi56 := 0; fi56 < 19; fi56++ {
                                    v57 := GenS7{F0: 42, F1: -13, F2: (v2 + (^(v2)))}
                                    v58 := GenS5{F0: ""}
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                            case "a":
                                var v59 uint8 = v2
                                _ = v59
                            }
                        case "a":
                            for fi60 := 0; fi60 < 5; fi60++ {
                                v2 = v2
                                _ = fi60
                            }
                        case "foo bar":
                            for fi61 := 0; fi61 < 4; fi61++ {
                                if (v2 < 101) {
                                    v3 = v3
                                    v2 = 42
                                    v62 := []int{fi61}
                                    _ = v62
                                } else {
                                    fmt.Println(v44, fi48)
                                }
                                _ = len(v4)
                                _ = fi61
                            }
                        default:
                            for fi63 := 0; fi63 < 5; fi63++ {
                                if (false && true) {
                                    _ = len(v4)
                                } else {
                                    fmt.Println(fi48)
                                }
                                _ = fi63
                            }
                        }
                    }
                case 32768:
                    var v64 bool = (128 == uint8(v1))
                    _ = v64
                case 127:
                    fmt.Println(v3, v3)
                default:
                    var v65 int8 = 15
                    _ = v65
                }
            }
            _ = fi48
        }
    }
    var v66 float32 = 3.14159
    for fi67 := 0; fi67 < 19; fi67++ {
        for fi68 := 0; fi68 < 19; fi68++ {
            v69 := []int{fi68, 2147483647, 32767}
            _ = fi68
            _ = v69
        }
        fmt.Println(v1, fi67)
        v70 := GenS2{F0: 2147483648}
        _ = fi67
        _ = v70
    }
    v71 := v4["foo bar"]
    v72 := GenS3{F0: 252}
    var v73 int = v71
    v74 := GenS6{F0: v73, F1: -2}
    var v75 float32 = (v66 - 3.14159)
    var v76 uint32 = 947
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v66:", v66)
    fmt.Println("v71:", v71)
    fmt.Println("v73:", v73)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v66
    _ = v71
    _ = v72
    _ = v73
    _ = v74
    _ = v75
    _ = v76
}
