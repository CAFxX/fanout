package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int }
func genFn0() int64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: false, F1: 100}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: (false || true), F1: -7}
    switch "hello" {
    case "\n\t":
        v10 := []int{416, 32768, 516, 2147483647}
        _ = v10
    case "":
        for fi11 := 0; fi11 < 4; fi11++ {
            var v12 float32 = 0.1
            v12 -= v12
            var v13 uint64 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
    case "a":
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := []int{-1, (fi14 / ((127) | 1)), (fi14 / (((fi14 | fi14)) | 1))}
            for fi16 := 0; fi16 < 16; fi16++ {
                switch fi16 {
                case -1:
                    switch fi14 {
                    case -1:
                        var v17 bool = false
                        _ = v17
                    case 752:
                        v18 := GenS1{F0: true, F1: 18}
                        _ = v18
                    default:
                        var v19 int16 = -32768
                        _ = v19
                    }
                case 32767:
                    fmt.Println(fi14, fi16)
                case -751:
                    if len(v15) > 0 {
                        v15[0] = fi14
                    }
                }
                v20 := GenS2{F0: true, F1: 118}
                _ = fi16
                _ = v20
            }
            if true {
                if (false || true) {
                    v21 := GenS1{F0: false, F1: 25}
                    _ = v21
                } else {
                    if ("hello" == "foo bar") {
                        if (((42 > 392) && true) || false) {
                            var v22 uint = 255
                            var v23 float32 = -2.25
                            _ = v22
                            _ = v23
                        } else {
                            fmt.Println(fi14)
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(fi14)
                            case "hello":
                                if true {
                                    var v24 uint16 = 1
                                    fmt.Println(v24, fi14)
                                    _ = v24
                                }
                            case "":
                                var v25 int8 = 127
                                _ = v25
                            }
                            var v26 bool = ((!((fi14 - (fi14 + fi14)) != -2147483648)) || (false || (!false)))
                            _ = v26
                        }
                        if (!true) {
                            var v27 int8 = -1
                            _ = len(v15)
                            _ = v27
                        }
                        var v28 float64 = ((0.0 / 1e10) / (1e10 / 0.1))
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 int32 = 2147483647
                            switch "" {
                            case "\n\t":
                                fmt.Println(v30, fi14, fi14)
                            default:
                                switch -15 {
                                case 128:
                                    fmt.Println(v30, fi14, fi14)
                                default:
                                    if len(v15) > 0 {
                                        v15[0] = (^(fi29))
                                    }
                                }
                            }
                            v31 := GenS3{F0: "a", F1: 0}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        fmt.Println(fi14)
                        if (42 > 650) {
                            if false {
                                for fi32 := 0; fi32 < 11; fi32++ {
                                    var v33 string = ("" + ("\n\t" + (("a" + "\n\t") + "")))
                                    if (18446744073709551615 >= 977) {
                                        if true {
                                            var v34 int16 = 32767
                                            v34 += -32768
                                            if false {
                                                _ = len(v15)
                                                switch fi32 {
                                                case 128:
                                                    var v35 int8 = int8(fi14)
                                                    _ = v35
                                                default:
                                                    if true {
                                                        v34 = v34
                                                        var v36 uint = 264
                                                        var v37 int64 = -9223372036854775808
                                                        var v38 int8 = 100
                                                        _ = v36
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                }
                                            } else {
                                                var v39 string = v33
                                                _ = v39
                                            }
                                            var v40 float64 = float64(1.5)
                                            _ = v34
                                            _ = v40
                                        } else {
                                            var v41 uint64 = 390
                                            if len(v15) > 0 {
                                                v15[0] = -169
                                            }
                                            var v42 float32 = 1e10
                                            _ = v41
                                            _ = v42
                                        }
                                        if false {
                                            v33 = (("hello" + (v33 + v33)) + "")
                                        } else {
                                            _ = len(v15)
                                            var v43 bool = false
                                            var v44 uint8 = 222
                                            _ = v43
                                            _ = v44
                                        }
                                        v33 = ((v33 + v33) + "foo bar")
                                        v15 = append(v15, fi32)
                                    } else {
                                        fmt.Println(fi14, fi32)
                                        var v45 int64 = 9223372036854775807
                                        _ = v45
                                    }
                                    fmt.Println(fi14, v33, fi32)
                                    var v46 uint = 255
                                    _ = fi32
                                    _ = v33
                                    _ = v46
                                }
                            } else {
                                var v47 int32 = 1
                                v48 := GenS4{F0: ("hello" + "foo bar"), F1: (v47 | (v47 & v47))}
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                    var v49 int64 = 0
                    _ = v49
                }
                var v50 string = ""
                if len(v15) >= 2 {
                    v15 = v15[1:len(v15)]
                }
                _ = v50
            } else {
                var v51 string = "a"
                _ = v51
            }
            _ = fi14
            _ = v15
        }
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        var v53 float64 = ((float64((1e10 + 0.0)) + 1.5) * (1.5 * ((1e10 - 0.0) + 0.1)))
        var v54 uint16 = 662
        _ = fi52
        _ = v53
        _ = v54
    }
    v55 := GenS2{F0: (2147483647 == -1), F1: 248}
    _ = v55
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 1
    var v2 int64 = -9223372036854775808
    var v3 float64 = 0.0
    var v4 uint32 = 0
    var v5 string = "foo bar"
    var v6 int8 = (v1 & (v1 % ((v1) | 1)))
    switch -902 {
    case 893:
        var v56 string = v5
        _ = v56
    default:
        v6 += (^(v6))
    }
    if (4294967295 != 0) {
        var v57 int = 128
        _ = v57
    }
    var v58 float64 = v3
    v59 := []int{32768, -227, -9223372036854775808, 127}
    switch (v5 + v5) {
    case "":
        for fi60 := 0; fi60 < 1; fi60++ {
            v59 = append(v59, fi60)
            _ = fi60
        }
    }
    v59 = append(v59, 4294967295)
    for fi61 := 0; fi61 < 1; fi61++ {
        switch fi61 {
        case -9223372036854775808:
            v62 := GenS4{F0: v5, F1: -2147483648}
            _ = v62
        case 0:
            v63 := GenS5{F0: 796}
            _ = v63
        }
        if false {
            fmt.Println(fi61)
        }
        _ = fi61
    }
    v58 -= 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
}
