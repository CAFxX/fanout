package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8; F1 int; F2 uint32 }
type GenS5 struct { F0 int16; F1 bool }

func main() {
    var v1 uint8 = 128
    var v2 uint8 = (v1 ^ (^(v1)))
    var v3 uint16 = 32767
    var v4 uint32 = 386
    var v5 int16 = -481
    var v6 bool = (!true)
    switch "hello" {
    case "hello":
        v5 = v5
    case "":
        v7 := GenS1{F0: (798 < 42), F1: int64(v1)}
        _ = v7
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            v1 = v1
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 uint = 18446744073709551615
                v1 = v1
                if v6 {
                    v11 := GenS1{F0: true, F1: 9223372036854775807}
                    var v12 uint64 = 2
                    v2 += v1
                    _ = v11
                    _ = v12
                } else {
                    var v13 int8 = 88
                    v6 = (((true || v6) && (true || v6)) || v6)
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 int8 = -128
                v4 += 4294967295
                _ = fi14
                _ = v15
            }
            _ = fi8
        }
    default:
        if ((v2 / ((v2) | 1)) >= 1) {
            fmt.Println(v3)
            if v6 {
                for fi16 := 0; fi16 < 8; fi16++ {
                    if ((-2.25 - (3.14159 / -2.25)) <= -2.25) {
                        v17 := []int{(^(fi16)), fi16}
                        fmt.Println(v2)
                        var v18 bool = v6
                        _ = v17
                        _ = v18
                    } else {
                        if (true && ((42 != 1) && v6)) {
                            if (true || (!v6)) {
                                v19 := []int{}
                                v5 -= v5
                                var v20 string = "hello"
                                _ = v19
                                _ = v20
                            } else {
                                var v21 float32 = (-2.25 - float32(((0.0 + -2.25) + float64(1.5))))
                                var v22 float64 = (0.1 / 1e10)
                                _ = v21
                                _ = v22
                            }
                        } else {
                            if v6 {
                                var v23 uint8 = v2
                                switch (^(fi16)) {
                                case 9223372036854775807:
                                    v6 = v6
                                case -461:
                                    v24 := GenS1{F0: (-493 == 2147483647), F1: 42}
                                    _ = v24
                                default:
                                    var v25 int32 = 0
                                    _ = v25
                                }
                                _ = v23
                            } else {
                                for fi26 := 0; fi26 < 4; fi26++ {
                                    v27 := GenS1{F0: v6, F1: 1}
                                    switch "hello" {
                                    case "a":
                                        v3 -= (v3 % (((v3 << ((255) & 63))) | 1))
                                        v28 := GenS2{F0: ((("\n\t" + "foo bar") + ("foo bar" + "")) + ""), F1: (!true), F2: v3}
                                        var v29 uint32 = 65535
                                        _ = v28
                                        _ = v29
                                    case "hello":
                                        v4 -= v4
                                    default:
                                        var v30 uint32 = 42
                                        _ = v30
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            fmt.Println(fi16)
                            fmt.Println(v5, v2)
                        }
                        v31 := GenS3{F0: 42, F1: (fi16 ^ fi16)}
                        _ = v31
                    }
                    _ = fi16
                }
            } else {
                var v32 float64 = -2.25
                _ = v32
            }
        }
    }
    v33 := []int{}
    for fi34 := 0; fi34 < 3; fi34++ {
        if false {
            if v6 {
                var v35 uint16 = uint16(v4)
                _ = v35
            } else {
                switch ("\n\t" + "hello") {
                case "":
                    if false {
                        var v36 int16 = -297
                        if v6 {
                            var v37 int = 32768
                            v33 = append(v33, fi34)
                            v36 += -32768
                            _ = v37
                        } else {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1, v5, v5)
                            default:
                                var v38 uint64 = 4294967295
                                _ = v38
                            }
                            if (!v6) {
                                fmt.Println(v3, v36, v5)
                            } else {
                                switch 487 {
                                case 1:
                                    if v6 {
                                        v39 := GenS4{F0: v2, F1: (fi34 / (((fi34 + -1)) | 1)), F2: uint32(v2)}
                                        fmt.Println(fi34, v4, v36)
                                        _ = v39
                                    }
                                case 9223372036854775807:
                                    v1 -= 127
                                case 4294967295:
                                    fmt.Println(v5)
                                }
                                var v40 int = -774
                                _ = v40
                            }
                            var v41 string = (("hello" + "foo bar") + ("\n\t" + "foo bar"))
                            _ = v41
                        }
                        _ = v36
                    }
                case "a":
                    v1 = (v2 * (v2 + v1))
                default:
                    v3 -= v3
                }
                v33 = append(v33, fi34)
            }
            switch fi34 {
            case -2147483648:
                v1 = (v1 % ((v2) | 1))
            case 0:
                v1 = v2
            case 9223372036854775807:
                v2 -= v1
            }
            v4 -= v4
        } else {
            _ = len(v33)
            if (false || v6) {
                var v42 string = "\n\t"
                var v43 float64 = ((1e10 + -2.25) / 3.14159)
                _ = v42
                _ = v43
            }
            v2 = 1
        }
        if (42 == 18446744073709551615) {
            fmt.Println(v2, v5, v6)
            fmt.Println(v4)
            switch fi34 {
            case 0:
                if len(v33) > 0 {
                    v33[0] = int(v2)
                }
            default:
                v5 += (v5 * v5)
            }
        } else {
            v4 += (v4 - 4294967295)
        }
        _ = fi34
    }
    if ((v6 || v6) && (-9223372036854775808 > 540)) {
        v44 := GenS5{F0: (v5 | (v5 - v5)), F1: true}
        if false {
            fmt.Println(v6, v6)
            var v45 uint16 = v3
            var v46 uint8 = v1
            _ = v45
            _ = v46
        } else {
            var v47 int8 = 71
            _ = v47
        }
        _ = v44
    } else {
        var v48 int = -1
        _ = v48
    }
    if (!true) {
        v49 := GenS4{F0: (v1 / (((v1 / ((v1) | 1))) | 1)), F1: 17, F2: v4}
        fmt.Println(v5, v3)
        if v6 {
            var v50 int64 = -542
            v50 -= (v50 / ((v50) | 1))
            _ = v50
        }
        _ = v49
    }
    v33 = append(v33, 4294967295)
    var v51 uint = 65535
    if v6 {
        fmt.Println(v51)
        fmt.Println(v2, v5, v3)
        switch ("foo bar" + ("a" + (("\n\t" + "foo bar") + "a"))) {
        case "":
            switch 1 {
            case 4294967295:
                v33 = append(v33, 2147483647)
            default:
                var v52 float64 = 0.0
                _ = v52
            }
        case "foo bar":
            v6 = ((v3 ^ (^(v3))) == v3)
        default:
            v53 := GenS3{F0: (v2 / (((v1 + v1)) | 1)), F1: -9223372036854775808}
            _ = v53
        }
    } else {
        if v6 {
            var v54 uint8 = v2
            _ = v54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
