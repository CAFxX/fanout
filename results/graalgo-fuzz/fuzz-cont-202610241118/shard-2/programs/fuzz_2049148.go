package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 0
    var v3 int32 = -272
    var v4 uint = 65535
    v5 := GenS1{F0: v3, F1: 0}
    if (0 > 42) {
        if (42 < 243) {
            var v6 string = "a"
            v3 = (v3 & v3)
            fmt.Println(v3, v1, v3)
            _ = v6
        } else {
            v4 = v4
            var v7 string = ("foo bar" + "\n\t")
            _ = v7
        }
    } else {
        switch 649 {
        case 481:
            v8 := []int{9223372036854775807, 4294967295, 32767, 196}
            _ = v8
        case 838:
            v9 := []int{-757, 1}
            _ = v9
        default:
            var v10 int = -9223372036854775808
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint32 = 627
        for fi13 := 0; fi13 < 16; fi13++ {
            v3 += 0
            _ = fi13
        }
        v14 := GenS1{F0: (^(v3)), F1: (v3 * (v3 - 54))}
        _ = fi11
        _ = v12
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        v1 = v1
        var v16 int8 = v2
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 807, F1: 197, F2: ("\n\t" + "foo bar")}
    v18 := []int{1, 42, 4294967295}
    v19 := GenS3{F0: v2}
    var v20 uint64 = 414
    if (-965 > int16(v1)) {
        var v21 int64 = 718
        var v22 bool = false
        switch "\n\t" {
        case "\n\t":
            if v22 {
                if (v21 > (v21 + (v21 ^ (^(v21))))) {
                    var v23 bool = false
                    var v24 uint32 = 2147483647
                    switch "a" {
                    case "a":
                        if (!(((!v22) || v23) || (v1 < 42))) {
                            var v25 uint8 = 84
                            switch 127 {
                            case 4294967295:
                                switch 42 {
                                case -774:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                case 128:
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                }
                            case 759:
                                for fi26 := 0; fi26 < 9; fi26++ {
                                    var v27 int16 = 253
                                    _ = fi26
                                    _ = v27
                                }
                            case 42:
                                var v28 int = -351
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 11; fi29++ {
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    var v31 int16 = 0
                                    v18 = append(v18, (fi29 * int(v25)))
                                    _ = fi30
                                    _ = v31
                                }
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v32 float32 = 0.1
                                _ = fi29
                                _ = v32
                            }
                            _ = v25
                        } else {
                            v33 := GenS3{F0: (v2 / (((v2 >> (((v4 * v4)) & 63))) | 1))}
                            _ = len(v18)
                            v34 := GenS3{F0: -128}
                            _ = v33
                            _ = v34
                        }
                    case "":
                        var v35 string = "a"
                        _ = v35
                    case "hello":
                        var v36 int16 = 32767
                        _ = v36
                    }
                    _ = v23
                    _ = v24
                }
            } else {
                switch (("hello" + "foo bar") + "a") {
                case "":
                    if v22 {
                        fmt.Println(v1, v21, v3)
                    }
                case "a":
                    var v37 float64 = -2.25
                    _ = v37
                default:
                    var v38 int16 = 42
                    _ = v38
                }
                if v22 {
                    var v39 uint = (v4 % (((v4 * (v4 ^ v4))) | 1))
                    _ = v39
                } else {
                    switch 32768 {
                    case -63:
                        switch "a" {
                        case "":
                            if false {
                                v40 := GenS1{F0: v3, F1: (v3 >> (((^(v4))) & 63))}
                                _ = v40
                            } else {
                                if false {
                                    for fi41 := 0; fi41 < 3; fi41++ {
                                        switch "\n\t" {
                                        case "":
                                            for fi42 := 0; fi42 < 1; fi42++ {
                                                fmt.Println(v2, v20)
                                                switch fi42 {
                                                case -2147483648:
                                                    fmt.Println(v2, v20)
                                                default:
                                                    if v22 {
                                                        if true {
                                                            var v43 int8 = (v2 >> ((0) & 63))
                                                            v44 := GenS2{F0: 0, F1: 65535, F2: ("hello" + "")}
                                                            v43 -= int8(fi41)
                                                            var v45 int64 = 1
                                                            v46 := GenS1{F0: (v3 | 42), F1: 1}
                                                            var v47 uint64 = (v20 >> ((v4) & 63))
                                                            v48 := GenS4{F0: (v4 - (v4 * v4))}
                                                            var v49 uint16 = (^(v1))
                                                            v21 += (v21 * v21)
                                                            var v50 int64 = (v21 ^ (v45 % ((int64(v20)) | 1)))
                                                            if v22 {
                                                                v49 = v1
                                                            }
                                                            var v51 int32 = v3
                                                            v52 := GenS4{F0: uint(v20)}
                                                            _ = v43
                                                            _ = v44
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        fmt.Println(v4, v2, v4)
                                                    }
                                                }
                                                var v53 uint16 = v1
                                                _ = fi42
                                                _ = v53
                                            }
                                        }
                                        switch ("hello" + "") {
                                        case "hello":
                                            v20 -= (v20 % ((v20) | 1))
                                        case "\n\t":
                                            v20 -= (v20 >> ((255) & 63))
                                        default:
                                            fmt.Println(v22, v21)
                                        }
                                        _ = fi41
                                    }
                                    v18 = append(v18, 617)
                                    var v54 float32 = ((3.14159 + 0.1) / 1.0)
                                    _ = v54
                                }
                                v55 := GenS2{F0: 1000, F1: 2147483648, F2: "hello"}
                                v56 := GenS1{F0: (v3 ^ (v3 ^ 42)), F1: v3}
                                _ = v55
                                _ = v56
                            }
                        default:
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                        }
                    default:
                        for fi57 := 0; fi57 < 6; fi57++ {
                            for fi58 := 0; fi58 < 9; fi58++ {
                                var v59 int64 = (-(v21))
                                v60 := GenS4{F0: v4}
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            fmt.Println(v2)
                            _ = fi57
                        }
                    }
                    for fi61 := 0; fi61 < 6; fi61++ {
                        v2 += v2
                        var v62 float64 = 1e10
                        _ = fi61
                        _ = v62
                    }
                    var v63 int32 = (v3 >> ((675) & 63))
                    _ = v63
                }
            }
        case "a":
            v64 := GenS2{F0: 32767, F1: 2147483648, F2: "\n\t"}
            _ = v64
        case "foo bar":
            var v65 uint16 = v1
            _ = v65
        }
        _ = v21
        _ = v22
    }
    switch -2147483648 {
    case 2147483647:
        v4 -= v4
    case 1:
        _ = len(v18)
    default:
        var v66 uint64 = v20
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
