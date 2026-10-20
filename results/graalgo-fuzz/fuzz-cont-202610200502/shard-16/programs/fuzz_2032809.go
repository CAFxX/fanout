package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 bool; F1 int32; F2 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint }
func genFn0(p0 uint, p1 uint8) uint8 {
    for fi5 := 0; fi5 < 17; fi5++ {
        p1 = p1
        v6 := []int{fi5}
        if true {
            var v7 uint32 = 2
            _ = v7
        } else {
            fmt.Println(p1, fi5)
        }
        _ = fi5
        _ = v6
    }
    var v8 uint8 = (p1 ^ (p1 * p1))
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: 1000}
        switch "\n\t" {
        case "a":
            v11 := GenS1{F0: 0}
            _ = v11
        case "foo bar":
            for fi12 := 0; fi12 < 9; fi12++ {
                v8 -= (v8 + (v8 | (p1 - (^(p1)))))
                _ = fi12
            }
        case "":
            var v13 int64 = -9223372036854775808
            _ = v13
        default:
            p0 += p0
        }
        p0 = 4294967295
        _ = fi9
        _ = v10
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        p1 = (v8 / ((p1) | 1))
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v8
    return (p1 - p1)
}
func genFn1(p0 int) uint8 {
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch p0 {
            case 32767:
                v17 := GenS2{F0: (fi16 / (((fi15 & 9223372036854775807)) | 1))}
                _ = v17
            case 9223372036854775807:
                p0 = -9223372036854775808
            case 0:
                if true {
                    v18 := GenS3{F0: 0, F1: 4294967295}
                    switch "\n\t" {
                    case "\n\t":
                        p0 = fi15
                    case "a":
                        p0 += fi16
                    case "":
                        p0 = 127
                    default:
                        if (42 > 0) {
                            var v19 uint16 = 0
                            var v20 int64 = -446
                            _ = v19
                            _ = v20
                        } else {
                            p0 = fi16
                        }
                    }
                    _ = v18
                } else {
                    p0 -= (fi16 >> ((42) & 63))
                }
            default:
                var v21 uint64 = 4294967295
                _ = v21
            }
            v22 := GenS1{F0: 1000}
            _ = fi16
            _ = v22
        }
        _ = fi15
    }
    var v23 string = "a"
    _ = p0
    _ = v23
    return 0
}

func main() {
    var v1 int = 471
    var v2 float64 = ((((0.0 / 1.5) / 0.1) + -2.25) - (-2.25 / 1e10))
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    v1 = (^(v1))
    v24 := GenS4{F0: ("hello" + "a"), F1: 7}
    _ = len(v4)
    var v25 uint64 = 0
    v2 = v2
    var v26 uint8 = 2
    v26 = (v26 % ((genFn0(284, genFn0(880, 62))) | 1))
    if false {
        v27 := []int{v1, v1, (^(v1)), 525}
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 int32 = -1
            _ = fi28
            _ = v29
        }
        if len(v27) > 0 {
            v27[0] = v1
        }
        _ = v27
    } else {
        var v30 float32 = (-2.25 + (1e10 + float32(v2)))
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 int32 = -1
            var v33 uint = 18446744073709551615
            switch v1 {
            case 651:
                for fi34 := 0; fi34 < 9; fi34++ {
                    if false {
                        _ = len(v4)
                    }
                    delete(v3, ("" + "hello"))
                    _ = fi34
                }
            case -710:
                for fi35 := 0; fi35 < 14; fi35++ {
                    if ((v33 ^ (v33 + v33)) <= v33) {
                        fmt.Println(v26, fi31)
                    }
                    _ = fi35
                }
            case -776:
                if false {
                    var v36 float32 = (v30 - float32(0.1))
                    var v37 bool = false
                    switch "" {
                    case "a":
                        var v38 uint16 = 1
                        _ = v38
                    case "":
                        v39 := []int{v1, fi31, 4294967295, v1}
                        _ = v39
                    case "\n\t":
                        v40 := GenS2{F0: (v1 * -212)}
                        _ = v40
                    default:
                        if (v37 && ((2147483647 == 65535) || v37)) {
                            if v37 {
                                fmt.Println(v30)
                                v33 = v33
                            }
                            v26 -= genFn1((fi31 - (-(v1))))
                            v41 := []int{32768, fi31}
                            _ = v41
                        }
                    }
                    _ = v36
                    _ = v37
                } else {
                    fmt.Println(v26)
                    for fi42 := 0; fi42 < 18; fi42++ {
                        switch fi31 {
                        case 620:
                            var v43 float64 = (3.14159 + v2)
                            _ = v43
                        case 4294967295:
                            v44 := GenS5{F0: true, F1: (v32 >> (((v33 / ((62) | 1))) & 63)), F2: fi42}
                            _ = v44
                        case 887:
                            if (((true || false) || (uint64(v26) <= (v25 + (v25 / ((v25) | 1))))) && true) {
                                var v45 string = "hello"
                                fmt.Println(v45)
                                _ = v45
                            }
                        default:
                            if (0.0 > (((v2 + v2) * (-2.25 + -2.25)) + ((v2 - (0.0 * v2)) * v2))) {
                                if (217 <= v25) {
                                    v26 = 42
                                    if ((v25 | v25) < (v25 % ((v25) | 1))) {
                                        v33 = v33
                                        v2 = float64(v30)
                                    } else {
                                        if false {
                                            var v46 int16 = -32768
                                            if (true && true) {
                                                fmt.Println(v32, v26, v2)
                                                v3[""] = false
                                            }
                                            _ = v46
                                        }
                                        if false {
                                            switch "hello" {
                                            case "hello":
                                                switch (fi31 ^ fi42) {
                                                case 32768:
                                                    switch fi42 {
                                                    case 9223372036854775807:
                                                        fmt.Println(v32, v26, v1)
                                                    }
                                                case -2147483648:
                                                    v47 := GenS6{F0: 1}
                                                    _ = v47
                                                default:
                                                    var v48 int32 = -2147483648
                                                    _ = v48
                                                }
                                            case "":
                                                var v49 uint = 4294967295
                                                _ = v49
                                            }
                                        } else {
                                            switch (fi42 << ((v33) & 63)) {
                                            case -9223372036854775808:
                                                fmt.Println(v30)
                                            case 128:
                                                if true {
                                                    if true {
                                                        if true {
                                                            var v50 uint32 = 495
                                                            v32 += (v32 * (v32 | (v32 & v32)))
                                                            _ = v50
                                                        } else {
                                                            v33 = (v33 * (v33 % (((v33 ^ v33)) | 1)))
                                                            switch "hello" {
                                                            case "foo bar":
                                                                switch fi31 {
                                                                case -1:
                                                                    fmt.Println(v32, v2, v30)
                                                                case 1:
                                                                    v51 := GenS7{F0: 2, F1: 2}
                                                                    _ = v51
                                                                }
                                                            default:
                                                                fmt.Println(v32, v25)
                                                            }
                                                        }
                                                    } else {
                                                        if false {
                                                            delete(v4, 127)
                                                            var v52 int8 = -122
                                                            _ = v52
                                                        }
                                                        v1 = (v1 - int(v33))
                                                    }
                                                    delete(v3, ("hello" + "hello"))
                                                } else {
                                                    switch fi42 {
                                                    case 128:
                                                        v1 = 128
                                                    case 127:
                                                        v53 := GenS6{F0: 42}
                                                        _ = v53
                                                    case 32767:
                                                        var v54 uint16 = 580
                                                        _ = v54
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                switch (((("foo bar" + "") + "") + "a") + "a") {
                                case "\n\t":
                                    v55 := GenS3{F0: 42, F1: 4294967295}
                                    _ = v55
                                case "foo bar":
                                    var v56 string = ""
                                    _ = v56
                                case "":
                                    v57 := GenS8{F0: 0}
                                    _ = v57
                                }
                                v33 = (v33 % (((^(v33))) | 1))
                            } else {
                                var v58 bool = (!("hello" < "hello"))
                                _ = v58
                            }
                        }
                        _ = fi42
                    }
                    var v59 int8 = -1
                    _ = v59
                }
            default:
                var v60 int8 = 1
                _ = v60
            }
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = v30
    }
    v4[584] = (-(v1))
    switch v1 {
    case 2147483647:
        v61 := GenS9{F0: 974}
        _ = v61
    case 9223372036854775807:
        delete(v4, (v1 - v1))
    default:
        v62 := v3[((("\n\t" + ((("foo bar" + ("a" + (("\n\t" + "hello") + "foo bar"))) + "hello") + "a")) + "foo bar") + "\n\t")]
        _ = v62
    }
    fmt.Println(v2, v1)
    var v63 float32 = -2.25
    v64 := GenS5{F0: false, F1: -2147483648, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v63
    _ = v64
}
