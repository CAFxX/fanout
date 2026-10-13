package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 uint32 }
type GenS4 struct { F0 int; F1 int16; F2 uint8 }
type GenS5 struct { F0 int8; F1 int16; F2 string }

func main() {
    var v1 uint32 = 697
    var v2 string = "foo bar"
    v3 := make(map[string]uint32)
    var v4 float32 = 1.5
    var v5 float64 = 1e10
    var v6 string = "hello"
    v4 += float32(v5)
    if false {
        var v7 int32 = 1
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        if ((true || false) || false) {
            if true {
                var v9 uint8 = 18
                for fi10 := 0; fi10 < 20; fi10++ {
                    v11 := GenS1{F0: -1}
                    v6 = "a"
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                var v12 float64 = 3.14159
                var v13 float32 = 0.0
                _ = v12
                _ = v13
            }
        } else {
            fmt.Println(v2, v4, v5)
            v14 := v3[v2]
            _ = v14
        }
        _ = fi8
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        switch (-(fi15)) {
        case 714:
            v6 = v6
        case -9223372036854775808:
            v1 = v1
        case -839:
            v16 := []int{(fi15 + fi15)}
            _ = v16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            if false {
                switch (fi15 & fi17) {
                case 42:
                    var v18 bool = (!false)
                    _ = v18
                default:
                    var v19 bool = (!true)
                    _ = v19
                }
                var v20 bool = (!false)
                v21 := GenS2{F0: -131, F1: 18, F2: 1}
                var v22 uint8 = 128
                v2 = v2
                _ = v20
                _ = v21
                _ = v22
            } else {
                fmt.Println(v1, fi15)
            }
            fmt.Println(fi15, v5)
            _ = fi17
        }
        v23 := v3[v6]
        _ = fi15
        _ = v23
    }
    delete(v3, "a")
    var v24 int = 1
    switch "a" {
    case "a":
        if false {
            _ = len(v3)
        } else {
            var v25 uint = 255
            _ = v25
        }
    case "foo bar":
        for fi26 := 0; fi26 < 11; fi26++ {
            for fi27 := 0; fi27 < 13; fi27++ {
                fmt.Println(v24)
                _ = fi27
            }
            _ = fi26
        }
    case "hello":
        for fi28 := 0; fi28 < 5; fi28++ {
            for fi29 := 0; fi29 < 14; fi29++ {
                switch v6 {
                case "":
                    var v30 int32 = 1
                    _ = v30
                default:
                    var v31 int32 = 0
                    _ = v31
                }
                _ = fi29
            }
            if ((4294967295 <= 42) || (false && (0.0 != float64((3.14159 / (((((float32((float64(v4) + float64(v4))) - 1e10) * v4) / v4) + v4) - v4)))))) {
                for fi32 := 0; fi32 < 18; fi32++ {
                    switch (v6 + ((v6 + "") + v2)) {
                    case "hello":
                        var v33 int8 = 127
                        _ = v33
                    case "a":
                        v5 -= ((v5 - v5) - v5)
                    }
                    v34 := GenS2{F0: 1000, F1: -128, F2: 42}
                    if (false || (!true)) {
                        var v35 int64 = 9223372036854775807
                        switch int(v1) {
                        case -840:
                            if (((0 >= 42) || (-9223372036854775808 == (fi32 % (((^(v24))) | 1)))) && true) {
                                v2 = (("\n\t" + (("" + v6) + "a")) + v2)
                                if (((v5 / (v5 - v5)) * v5) <= ((v5 - v5) + v5)) {
                                    v1 = (v1 << ((255) & 63))
                                    var v36 int8 = -122
                                    _ = v36
                                } else {
                                    var v37 uint16 = 65535
                                    var v38 int32 = 0
                                    var v39 uint64 = 42
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            }
                        case 32767:
                            v40 := []int{(fi32 + (v24 << ((384) & 63))), (fi28 | fi32), 9223372036854775807, (-(v24))}
                            _ = v40
                        case 753:
                            var v41 uint32 = v1
                            if true {
                                switch v24 {
                                case 2147483647:
                                    v4 -= 3.14159
                                case 356:
                                    if true {
                                        v35 += 0
                                        if (-1 < (v35 % (((v35 & v35)) | 1))) {
                                            switch 32768 {
                                            case 128:
                                                if ((true && false) && false) {
                                                    fmt.Println(v41)
                                                } else {
                                                    v41 = 2147483647
                                                    fmt.Println(v2, fi32)
                                                }
                                            case 895:
                                                fmt.Println(v5, v5)
                                            default:
                                                if false {
                                                    v4 = (v4 + (v4 * v4))
                                                    v24 = (fi28 % ((fi28) | 1))
                                                }
                                            }
                                            v41 -= v41
                                            v42 := v3[v2]
                                            _ = v42
                                        } else {
                                            v3[("a" + v6)] = v41
                                        }
                                    }
                                case 938:
                                    var v43 int = fi28
                                    _ = v43
                                }
                                fmt.Println(fi28, v1, fi32)
                                var v44 bool = ((3.14159 / v4) != float32((3.14159 * 0.0)))
                                _ = v44
                            }
                            fmt.Println(v6, v2, fi28)
                            _ = v41
                        }
                        _ = v35
                    } else {
                        var v45 string = v6
                        _ = v45
                    }
                    if false {
                        v46 := GenS3{F0: 173, F1: 0, F2: (v1 - v1)}
                        _ = v46
                    }
                    _ = fi32
                    _ = v34
                }
                v47 := GenS4{F0: (fi28 | 42), F1: 611, F2: 1}
                _ = v47
            } else {
                for fi48 := 0; fi48 < 2; fi48++ {
                    var v49 uint64 = 42
                    _ = fi48
                    _ = v49
                }
                for fi50 := 0; fi50 < 19; fi50++ {
                    if false {
                        v51 := GenS5{F0: -101, F1: 1, F2: v6}
                        _ = v51
                    } else {
                        var v52 float32 = 3.14159
                        _ = v52
                    }
                    v53 := GenS5{F0: -1, F1: 0, F2: ("hello" + v2)}
                    _ = fi50
                    _ = v53
                }
                fmt.Println(v1, v1, v24)
            }
            switch "hello" {
            case "hello":
                if (!(1.5 > 0.1)) {
                    v6 = v6
                    if false {
                        v54 := GenS3{F0: -910, F1: 9223372036854775807, F2: 2147483647}
                        v3[v6] = (^(v1))
                        _ = v54
                    }
                    var v55 string = ("" + (v2 + v6))
                    _ = v55
                }
            }
            _ = fi28
        }
    default:
        var v56 int8 = 127
        _ = v56
    }
    var v57 uint64 = 0
    var v58 uint32 = 2147483647
    switch (^(v24)) {
    case -1:
        for fi59 := 0; fi59 < 10; fi59++ {
            var v60 int = 0
            _ = fi59
            _ = v60
        }
    case 32768:
        _ = len(v3)
    case 32767:
        v58 += (v58 ^ v58)
    default:
        switch (^(v24)) {
        case 0:
            for fi61 := 0; fi61 < 7; fi61++ {
                switch fi61 {
                case -1:
                    delete(v3, ("a" + "a"))
                }
                _ = fi61
            }
        case 1:
            if true {
                v57 = 933
                _ = len(v3)
                switch v24 {
                case 2147483647:
                    for fi62 := 0; fi62 < 9; fi62++ {
                        var v63 int8 = 42
                        _ = fi62
                        _ = v63
                    }
                default:
                    var v64 uint16 = 1
                    _ = v64
                }
            } else {
                var v65 uint16 = 149
                switch -441 {
                case 9223372036854775807:
                    v1 = (v58 % ((v1) | 1))
                case -993:
                    v57 = v57
                case 127:
                    if (v1 < (v58 << ((uint(v57)) & 63))) {
                        if true {
                            v24 += (v24 & (^(v24)))
                        }
                        fmt.Println(v58, v57, v1)
                        var v66 int8 = -68
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 13; fi67++ {
                            var v68 int32 = -962
                            v6 = v2
                            fmt.Println(v2, v4)
                            _ = fi67
                            _ = v68
                        }
                        v69 := v3[("foo bar" + (v6 + v6))]
                        if (v6 < (v2 + "\n\t")) {
                            v58 += v69
                            v2 = (v2 + ("foo bar" + v2))
                            fmt.Println(v69)
                        } else {
                            v70 := []int{v24}
                            for fi71 := 0; fi71 < 12; fi71++ {
                                switch v24 {
                                case 1:
                                    var v72 bool = true
                                    _ = v72
                                }
                                if true {
                                    var v73 int8 = -117
                                    _ = v73
                                }
                                v6 = ("" + v6)
                                _ = fi71
                            }
                            if false {
                                var v74 int32 = 493
                                for fi75 := 0; fi75 < 14; fi75++ {
                                    v74 = (v74 * (v74 >> ((4294967295) & 63)))
                                    v2 = ("foo bar" + v2)
                                    fmt.Println(v1, v4, v4)
                                    _ = fi75
                                }
                                _ = v74
                            }
                            _ = v70
                        }
                        _ = v69
                    }
                default:
                    var v76 uint8 = 62
                    _ = v76
                }
                var v77 int = v24
                _ = v65
                _ = v77
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
