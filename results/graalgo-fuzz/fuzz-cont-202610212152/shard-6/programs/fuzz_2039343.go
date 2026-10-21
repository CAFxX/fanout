package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int16; F1 bool; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 0.0
    v3 := make(map[string]bool)
    var v4 uint16 = 42
    v5 := GenS1{F0: 0, F1: (v4 + (v4 & (^(v4))))}
    if true {
        switch "" {
        case "foo bar":
            switch (("hello" + "foo bar") + "foo bar") {
            case "":
                if true {
                    switch 2147483647 {
                    case 9223372036854775807:
                        switch -9223372036854775808 {
                        case 0:
                            var v6 uint8 = 0
                            _ = v6
                        case -259:
                            v4 -= 65535
                        }
                    case 127:
                        _ = len(v3)
                    default:
                        if ((42 < 0) || (42 == uint8(v4))) {
                            for fi7 := 0; fi7 < 16; fi7++ {
                                if ("\n\t" <= ("\n\t" + "")) {
                                    switch fi7 {
                                    case 128:
                                        v8 := GenS1{F0: 1, F1: v4}
                                        _ = v8
                                    case 42:
                                        fmt.Println(v2, fi7, v4)
                                    case -9223372036854775808:
                                        v4 -= 700
                                    default:
                                        var v9 int64 = 396
                                        _ = v9
                                    }
                                }
                                v10 := []int{(fi7 + fi7), fi7}
                                _ = fi7
                                _ = v10
                            }
                        }
                    }
                } else {
                    v2 += v2
                    if false {
                        if true {
                            if (206 >= 255) {
                                var v11 uint = 255
                                var v12 float32 = (v2 / v2)
                                delete(v3, "\n\t")
                                _ = v11
                                _ = v12
                            }
                            var v13 uint64 = 1
                            v14 := GenS1{F0: 58, F1: (v4 | (^(v4)))}
                            _ = v13
                            _ = v14
                        }
                        v15 := []int{-91, -1, 4294967295, -9223372036854775808}
                        v4 = v4
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 18; fi16++ {
                            for fi17 := 0; fi17 < 19; fi17++ {
                                v4 = 2
                                var v18 float64 = 1.5
                                _ = fi17
                                _ = v18
                            }
                            _ = fi16
                        }
                        v4 = 500
                        var v19 uint8 = 55
                        _ = v19
                    }
                    v20 := GenS1{F0: 100, F1: 787}
                    _ = v20
                }
            default:
                if ("" == "a") {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if true {
                            if true {
                                var v22 uint16 = (v4 | v4)
                                v23 := []int{fi21, fi21, fi21, (fi21 + fi21)}
                                _ = v22
                                _ = v23
                            } else {
                                v24 := []int{fi21}
                                var v25 uint16 = 888
                                _ = len(v24)
                                _ = v24
                                _ = v25
                            }
                            v26 := []int{}
                            v4 = v4
                            _ = v26
                        } else {
                            v4 = v4
                            var v27 int8 = -128
                            var v28 uint16 = v4
                            _ = v27
                            _ = v28
                        }
                        if false {
                            for fi29 := 0; fi29 < 2; fi29++ {
                                switch "hello" {
                                case "a":
                                    v30 := GenS1{F0: 100, F1: v4}
                                    _ = v30
                                default:
                                    fmt.Println(fi21, v4, fi29)
                                }
                                v4 += 32767
                                _ = fi29
                            }
                            v4 -= (^(v4))
                        } else {
                            v31 := GenS2{F0: 2, F1: (-(fi21))}
                            v2 = ((v2 - (v2 + v2)) * ((v2 * v2) * (float32((1.5 / 3.14159)) * v2)))
                            for fi32 := 0; fi32 < 14; fi32++ {
                                v33 := v1[42]
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        }
                        _ = fi21
                    }
                    var v34 uint64 = 2
                    _ = v34
                } else {
                    var v35 int8 = -1
                    _ = v35
                }
            }
        case "\n\t":
            var v36 int32 = 42
            _ = v36
        case "hello":
            v2 += v2
        default:
            var v37 int8 = 0
            _ = v37
        }
    } else {
        for fi38 := 0; fi38 < 8; fi38++ {
            if (true || (0 != 2)) {
                var v39 int = 0
                var v40 bool = false
                v41 := GenS1{F0: 42, F1: 0}
                _ = v39
                _ = v40
                _ = v41
            } else {
                if false {
                    v4 = 754
                    fmt.Println(v4)
                    var v42 uint32 = 883
                    _ = v42
                } else {
                    var v43 int8 = 46
                    var v44 int32 = 2147483647
                    _ = v43
                    _ = v44
                }
            }
            for fi45 := 0; fi45 < 3; fi45++ {
                var v46 int = -794
                v2 = v2
                v47 := []int{}
                _ = fi45
                _ = v46
                _ = v47
            }
            _ = fi38
        }
    }
    v48 := GenS3{F0: 135, F1: false, F2: -32768}
    for fi49 := 0; fi49 < 8; fi49++ {
        var v50 string = ""
        v51 := GenS1{F0: 0, F1: v4}
        v52 := GenS3{F0: 32767, F1: (false || (!((!true) || true))), F2: 42}
        _ = fi49
        _ = v50
        _ = v51
        _ = v52
    }
    if true {
        for fi53 := 0; fi53 < 11; fi53++ {
            if (!(int16(v4) != 32767)) {
                if true {
                    var v54 uint16 = (v4 - (^(v4)))
                    _ = len(v1)
                    _ = v54
                } else {
                    var v55 bool = true
                    v55 = (((65535 >= 4294967295) && (0 <= 42)) && (v2 == (float32((float64((v2 / 1.0)) - 1e10)) - v2)))
                    _ = v55
                }
                v2 = (v2 / v2)
            } else {
                fmt.Println(v4)
                var v56 uint64 = 2
                var v57 float32 = 1.5
                _ = v56
                _ = v57
            }
            _ = fi53
        }
        if false {
            v58 := []int{}
            if true {
                v2 -= (((1e10 * v2) - float32(3.14159)) / 0.1)
            }
            _ = v58
        } else {
            v59 := []int{1, -274}
            _ = v59
        }
        for fi60 := 0; fi60 < 4; fi60++ {
            var v61 bool = true
            fmt.Println(v61, v61, v2)
            _ = fi60
            _ = v61
        }
    } else {
        v4 = v4
        v62 := GenS2{F0: 419, F1: 32767}
        var v63 int8 = 42
        _ = v62
        _ = v63
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        var v65 int32 = -2147483648
        _ = fi64
        _ = v65
    }
    if true {
        switch -9223372036854775808 {
        case 127:
            _ = len(v3)
        case 42:
            var v66 string = "a"
            _ = v66
        }
    }
    for fi67 := 0; fi67 < 12; fi67++ {
        v68 := GenS1{F0: -111, F1: v4}
        _ = fi67
        _ = v68
    }
    v2 += -2.25
    switch 4294967295 {
    case -320:
        v69 := v1[2147483648]
        _ = v69
    default:
        for fi70 := 0; fi70 < 7; fi70++ {
            v2 = v2
            _ = fi70
        }
    }
    v71 := []int{-1}
    for fi72 := 0; fi72 < 7; fi72++ {
        var v73 float64 = (float64(v2) * 0.0)
        var v74 int16 = 42
        _ = fi72
        _ = v73
        _ = v74
    }
    v4 += v4
    v4 += 1
    if true {
        fmt.Println(v4)
        switch ("foo bar" + (("" + "foo bar") + ("a" + "a"))) {
        case "":
            for fi75 := 0; fi75 < 19; fi75++ {
                switch (("a" + "foo bar") + "\n\t") {
                case "":
                    v2 -= ((v2 - (1e10 * v2)) - 1.5)
                }
                v2 += v2
                _ = fi75
            }
        default:
            if false {
                v4 = (v4 - 65535)
                switch "a" {
                case "hello":
                    v4 = (v4 % ((v4) | 1))
                }
            }
        }
        for fi76 := 0; fi76 < 5; fi76++ {
            for fi77 := 0; fi77 < 6; fi77++ {
                var v78 uint16 = (v4 / ((v4) | 1))
                _ = fi77
                _ = v78
            }
            _ = fi76
        }
    }
    v79 := GenS2{F0: 975, F1: -610}
    fmt.Println(v2, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v71:", len(v71), v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v71
    _ = v79
}
