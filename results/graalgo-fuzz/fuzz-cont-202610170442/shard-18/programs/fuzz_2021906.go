package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 int16; F2 string }
type GenS5 struct { F0 int8; F1 int16; F2 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64 }
func genFn0() uint {
    switch "a" {
    case "foo bar":
        var v5 int = 0
        _ = v5
    default:
        v6 := GenS1{F0: 42, F1: 285, F2: 1}
        _ = v6
    }
    var v7 int8 = -1
    fmt.Println(v7)
    _ = v7
    return 65535
}

func main() {
    var v1 float64 = (float64(1e10) / 1.5)
    v2 := []uint32{}
    var v3 string = "\n\t"
    var v4 string = v3
    if (false && false) {
        var v8 int = 9223372036854775807
        _ = v8
    } else {
        switch "\n\t" {
        case "":
            for fi9 := 0; fi9 < 9; fi9++ {
                if (!true) {
                    v3 = ((v4 + "") + ((v3 + v4) + v3))
                } else {
                    v1 -= (v1 + float64(((0.1 / ((float32((0.1 * (0.1 - ((((v1 * 3.14159) * 0.1) + float64(3.14159)) * 1e10)))) - (1.5 / ((float32(v1) / (float32(v1) * 0.1)) * float32((v1 * ((v1 * ((v1 - float64(((1.5 + 1e10) * 1.5))) + v1)) + ((0.0 / (v1 * 1.5)) * 1.5))))))) - (float32(0.0) + float32((1e10 - (((1.5 - v1) * (v1 * v1)) / v1)))))) / 1.0)))
                }
                v1 = (v1 / 1.5)
                _ = fi9
            }
        case "a":
            if (!(true || true)) {
                for fi10 := 0; fi10 < 17; fi10++ {
                    switch fi10 {
                    case 32767:
                        v3 = v4
                    default:
                        if false {
                            if false {
                                if true {
                                    v1 = v1
                                }
                                for fi11 := 0; fi11 < 14; fi11++ {
                                    v12 := GenS2{F0: 0, F1: (255 > 4294967295), F2: 255}
                                    _ = fi11
                                    _ = v12
                                }
                            } else {
                                v1 = v1
                                switch (("foo bar" + "foo bar") + v4) {
                                case "foo bar":
                                    if true {
                                        v13 := GenS2{F0: 42, F1: (44 == 42), F2: 1}
                                        _ = v13
                                    } else {
                                        var v14 uint32 = 244
                                        _ = v14
                                    }
                                case "\n\t":
                                    fmt.Println(v4)
                                default:
                                    v4 = "a"
                                }
                                fmt.Println(v4, v4)
                            }
                            v4 = "\n\t"
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 2
                            }
                            v3 = ("foo bar" + "")
                        }
                    }
                    v3 = ""
                    if (42 >= (fi10 - fi10)) {
                        v15 := GenS1{F0: 37, F1: -32768, F2: 1}
                        fmt.Println(v1)
                        switch (v4 + (("a" + v4) + ((v4 + "hello") + v3))) {
                        case "foo bar":
                            v16 := GenS3{F0: 18446744073709551615}
                            _ = v16
                        case "hello":
                            if (-2147483648 < 42) {
                                var v17 int8 = 42
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 17; fi18++ {
                                    if (255 <= 2) {
                                        v19 := GenS4{F0: fi18, F1: 0, F2: v3}
                                        var v20 int8 = 103
                                        _ = v19
                                        _ = v20
                                    }
                                    v4 = (("hello" + v4) + (v3 + ("" + "hello")))
                                    fmt.Println(v3)
                                    _ = fi18
                                }
                                if ((float64(((0.1 * (1.5 * 1e10)) - (0.1 + 1e10))) - v1) < float64(0.1)) {
                                    v21 := GenS3{F0: 4294967295}
                                    _ = v21
                                } else {
                                    var v22 uint8 = 128
                                    switch (fi10 * -353) {
                                    case -1:
                                        var v23 float64 = v1
                                        _ = v23
                                    case 396:
                                        v2 = append(v2, 2147483647)
                                    case -695:
                                        v1 += v1
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v22
                                }
                            }
                        }
                        _ = v15
                    }
                    _ = fi10
                }
                var v24 uint8 = 255
                if true {
                    if ((438 >= -429) && true) {
                        v25 := GenS3{F0: 2}
                        switch ((v3 + (v3 + "\n\t")) + v3) {
                        case "":
                            v4 = "\n\t"
                        }
                        var v26 uint8 = (v24 * 2)
                        _ = v25
                        _ = v26
                    }
                }
                _ = v24
            } else {
                _ = len(v2)
                v2 = append(v2, 65535)
                var v27 uint8 = 1
                _ = v27
            }
        case "\n\t":
            var v28 int = 73
            _ = v28
        }
        v4 = v3
    }
    switch v3 {
    case "\n\t":
        v4 = v4
    case "":
        v2 = append(v2, 4294967295)
    case "hello":
        fmt.Println(v3, v4, v4)
    }
    var v29 uint8 = 211
    var v30 int32 = 312
    var v31 int32 = 2147483647
    v32 := GenS3{F0: 4294967295}
    fmt.Println(v4, v29)
    v33 := GenS1{F0: (^(v29)), F1: 42, F2: -32768}
    if (788 < -1) {
        switch "\n\t" {
        case "":
            fmt.Println(v3, v30)
        case "\n\t":
            fmt.Println(v3)
        default:
            if false {
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := GenS1{F0: v29, F1: 1000, F2: -1}
                    if false {
                        for fi36 := 0; fi36 < 18; fi36++ {
                            fmt.Println(v1, v31, v31)
                            var v37 string = ("a" + "a")
                            v3 = ""
                            _ = fi36
                            _ = v37
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        switch ((v4 + v4) + ("foo bar" + v3)) {
                        case "foo bar":
                            if false {
                                _ = len(v2)
                                fmt.Println(fi34, v31)
                                v38 := GenS1{F0: 42, F1: 42, F2: 1000}
                                _ = v38
                            } else {
                                if true {
                                    v39 := GenS5{F0: -128, F1: 42, F2: 0}
                                    for fi40 := 0; fi40 < 12; fi40++ {
                                        v4 = v3
                                        var v41 int64 = 1
                                        switch "a" {
                                        case "hello":
                                            v31 -= v31
                                        case "a":
                                            var v42 uint8 = v29
                                            _ = v42
                                        default:
                                            var v43 uint = 506
                                            _ = v43
                                        }
                                        _ = fi40
                                        _ = v41
                                    }
                                    _ = v39
                                }
                                fmt.Println(v30, v1)
                                var v44 float64 = v1
                                _ = v44
                            }
                        default:
                            for fi45 := 0; fi45 < 4; fi45++ {
                                v46 := GenS4{F0: (^(fi34)), F1: 42, F2: ((v3 + "") + "\n\t")}
                                _ = fi45
                                _ = v46
                            }
                        }
                    }
                    var v47 uint32 = 608
                    _ = fi34
                    _ = v35
                    _ = v47
                }
                if (!true) {
                    var v48 uint = 1
                    v49 := GenS1{F0: (v29 % (((v29 & v29)) | 1)), F1: 848, F2: 42}
                    var v50 int = -6
                    _ = v48
                    _ = v49
                    _ = v50
                } else {
                    fmt.Println(v29, v3, v31)
                }
                switch -385 {
                case 127:
                    _ = len(v2)
                }
            }
        }
        for fi51 := 0; fi51 < 15; fi51++ {
            for fi52 := 0; fi52 < 4; fi52++ {
                v31 -= 1
                _ = len(v2)
                switch (((v4 + v3) + ((("foo bar" + (("a" + v4) + v3)) + v4) + ((v3 + "") + ("a" + v3)))) + (v4 + ("foo bar" + ((v3 + v3) + (((v3 + (("a" + (("\n\t" + (v4 + ("foo bar" + "a"))) + (((v3 + (v3 + v4)) + ((v3 + v4) + (v3 + v3))) + v3))) + "")) + "\n\t") + (((v4 + "") + (v4 + ("\n\t" + (("a" + (v3 + "")) + v3)))) + ("a" + ("\n\t" + ("hello" + ("hello" + "")))))))))) {
                case "hello":
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                case "\n\t":
                    if (-72 > 79) {
                        fmt.Println(v4)
                    } else {
                        switch "\n\t" {
                        case "a":
                            var v53 int64 = 0
                            _ = v53
                        case "hello":
                            var v54 int32 = (v30 << ((255) & 63))
                            _ = v54
                        }
                        if (!(100 > 101)) {
                            v4 = (v4 + ("hello" + (v3 + (("" + "\n\t") + (v3 + "")))))
                            v55 := GenS6{F0: 9223372036854775807}
                            _ = v55
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                case "foo bar":
                    if (!(v29 > 255)) {
                        if true {
                            var v56 float64 = 1.5
                            _ = v56
                        } else {
                            var v57 float32 = 1e10
                            v58 := GenS2{F0: int16(fi51), F1: false, F2: 2}
                            _ = v57
                            _ = v58
                        }
                        v59 := GenS7{F0: v31}
                        var v60 float64 = v1
                        if len(v2) > 0 {
                            v2[0] = 65535
                        }
                        _ = v59
                        _ = v60
                    }
                default:
                    if (18446744073709551615 > 4294967295) {
                        v3 = (v4 + ("foo bar" + "a"))
                    } else {
                        v61 := GenS8{F0: 42}
                        v62 := GenS1{F0: 2, F1: 1, F2: 1000}
                        _ = v61
                        _ = v62
                    }
                }
                fmt.Println(v30, v4, v29)
                _ = fi52
            }
            _ = fi51
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
