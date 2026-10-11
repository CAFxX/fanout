package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 uint16 }
type GenS4 struct { F0 uint8; F1 string; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 uint16; F2 int8 }
type GenS10 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int8 = 1
    v4 := make(map[string]string)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    v3 = -1
    fmt.Println(v3)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 bool = (false || (!true))
        v3 = v3
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 128}
    switch "\n\t" {
    case "\n\t":
        var v10 float64 = 0.0
        _ = v10
    case "foo bar":
        if false {
            var v11 uint = 238
            v11 = v11
            _ = v11
        }
    case "":
        var v12 uint8 = 136
        _ = v12
    default:
        var v13 string = (("hello" + "hello") + (("\n\t" + "hello") + "\n\t"))
        _ = v13
    }
    v14 := GenS2{F0: 255}
    var v15 uint8 = 255
    v16 := GenS3{F0: 769, F1: 42}
    for fi17 := 0; fi17 < 15; fi17++ {
        delete(v4, "\n\t")
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS4{F0: v15, F1: "\n\t", F2: 0}
            v20 := GenS3{F0: 4294967295, F1: 0}
            fmt.Println(v15)
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v21 := GenS3{F0: 0, F1: 1}
    v22 := GenS2{F0: 2}
    _ = len(v4)
    switch -123 {
    case -334:
        if false {
            v15 = 42
            var v23 int8 = -128
            v3 += -118
            _ = v23
        }
    case 1:
        v24 := GenS2{F0: 255}
        _ = v24
    case 9223372036854775807:
        v25 := GenS2{F0: 32767}
        _ = v25
    default:
        for fi26 := 0; fi26 < 16; fi26++ {
            for fi27 := 0; fi27 < 1; fi27++ {
                var v28 uint16 = 42
                fmt.Println(v15)
                _ = fi27
                _ = v28
            }
            switch ("" + "hello") {
            case "":
                v3 = v3
            case "a":
                v3 = 42
            case "hello":
                fmt.Println(v15)
            default:
                var v29 uint64 = 253
                _ = v29
            }
            for fi30 := 0; fi30 < 17; fi30++ {
                switch ("foo bar" + (("\n\t" + ("" + "a")) + "\n\t")) {
                case "":
                    if false {
                        v31 := GenS3{F0: 595, F1: 400}
                        switch "a" {
                        case "hello":
                            if (fi26 > (fi26 << ((255) & 63))) {
                                v32 := GenS4{F0: v15, F1: "\n\t", F2: -1}
                                _ = v32
                            } else {
                                var v33 int16 = 1
                                if (((18446744073709551615 != 42) || true) || false) {
                                    var v34 uint16 = 365
                                    v35 := GenS5{F0: v15, F1: 704, F2: v34}
                                    _ = v34
                                    _ = v35
                                }
                                _ = v33
                            }
                        case "\n\t":
                            delete(v1, 42)
                        case "":
                            if (535 < 1) {
                                v3 = v3
                                switch (fi30 >> ((18446744073709551615) & 63)) {
                                case 743:
                                    switch "a" {
                                    case "\n\t":
                                        fmt.Println(v15)
                                    case "a":
                                        var v36 int64 = 0
                                        _ = v36
                                    case "":
                                        if (!false) {
                                            var v37 int64 = 619
                                            _ = v37
                                        } else {
                                            var v38 int = (fi26 - (fi30 & fi30))
                                            switch ("\n\t" + "") {
                                            case "a":
                                                v3 -= (v3 - 1)
                                            case "":
                                                var v39 int16 = 32767
                                                _ = v39
                                            }
                                            v15 -= v15
                                            v3 += v3
                                            var v40 uint8 = (v15 | v15)
                                            _ = v38
                                            _ = v40
                                        }
                                    }
                                    var v41 float32 = 0.0
                                    var v42 uint8 = (v15 >> ((583) & 63))
                                    _ = v41
                                    _ = v42
                                case -668:
                                    v43 := GenS5{F0: (v15 + v15), F1: 133, F2: 255}
                                    _ = v43
                                default:
                                    v3 = (v3 * v3)
                                }
                            }
                        }
                        _ = v31
                    }
                default:
                    v44 := []int{128, 32767, (^(fi26)), fi30}
                    _ = v44
                }
                if false {
                    var v45 uint32 = 2
                    _ = v45
                } else {
                    if true {
                        var v46 uint = 255
                        var v47 int32 = 2147483647
                        _ = v46
                        _ = v47
                    } else {
                        v48 := GenS6{F0: 18446744073709551615, F1: fi26}
                        fmt.Println(v3, fi30)
                        switch -358 {
                        case 333:
                            var v49 int16 = 1
                            _ = v49
                        default:
                            var v50 string = ("a" + "foo bar")
                            _ = v50
                        }
                        _ = v48
                    }
                    switch fi26 {
                    case -9223372036854775808:
                        if (42 == 42) {
                            v51 := []int{fi26, fi26}
                            v52 := GenS1{F0: (v15 / ((127) | 1))}
                            var v53 int32 = 2147483647
                            var v54 float32 = (((-2.25 / 1.0) - (float32((-2.25 - -2.25)) / (0.1 * 1e10))) * (0.0 + ((0.0 / -2.25) - -2.25)))
                            fmt.Println(fi30, v54, v15)
                            _ = v51
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    case 1:
                        delete(v1, 65535)
                    case -721:
                        if false {
                            if (42 <= -1) {
                                v55 := GenS4{F0: (v15 - v15), F1: "", F2: 32767}
                                _ = v55
                            }
                            v56 := GenS7{F0: fi26}
                            v57 := GenS5{F0: v15, F1: 0, F2: 225}
                            _ = v56
                            _ = v57
                        } else {
                            var v58 uint = 0
                            v59 := GenS2{F0: 0}
                            fmt.Println(v3, fi26)
                            v15 += v15
                            _ = v58
                            _ = v59
                        }
                    default:
                        if false {
                            var v60 uint64 = 2
                            if true {
                                fmt.Println(v60)
                                switch "\n\t" {
                                case "\n\t":
                                    v3 = (v3 - 100)
                                case "":
                                    v61 := GenS6{F0: v60, F1: (fi30 - 42)}
                                    _ = v61
                                case "a":
                                    var v62 bool = false
                                    _ = v62
                                }
                            }
                            _ = v60
                        } else {
                            v3 = -128
                        }
                        v63 := GenS8{F0: 18446744073709551615}
                        switch "" {
                        case "foo bar":
                            v64 := []int{fi26, 1}
                            _ = v64
                        case "":
                            v65 := GenS2{F0: 2}
                            _ = v65
                        case "hello":
                            v3 -= (v3 ^ v3)
                        default:
                            v66 := GenS9{F0: 1000, F1: 32767, F2: v3}
                            _ = v66
                        }
                        _ = v63
                    }
                    v15 = (^(v15))
                    _ = len(v2)
                    v15 = v15
                }
                v67 := GenS10{F0: 4294967295}
                _ = fi30
                _ = v67
            }
            _ = fi26
        }
    }
    var v68 uint32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v9
}
