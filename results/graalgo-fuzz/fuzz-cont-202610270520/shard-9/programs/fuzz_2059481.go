package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int64; F1 int8; F2 int64 }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int8; F1 int32; F2 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 317
    var v3 float32 = (0.0 + (((3.14159 + 1.5) / float32(1.5)) / (float32(1e10) / 3.14159)))
    v4 := make(map[string]uint32)
    v5 := []uint32{960}
    if (9223372036854775807 <= 4294967295) {
        var v6 int16 = 869
        _ = v6
    } else {
        var v7 uint32 = 508
        for fi8 := 0; fi8 < 20; fi8++ {
            v3 += float32((1e10 + (0.0 / 0.1)))
            switch ("\n\t" + "foo bar") {
            case "":
                v3 += -2.25
            case "foo bar":
                v3 = v3
            case "a":
                for fi9 := 0; fi9 < 11; fi9++ {
                    var v10 int64 = v2
                    _ = fi9
                    _ = v10
                }
            }
            for fi11 := 0; fi11 < 12; fi11++ {
                if true {
                    switch (fi8 << ((42) & 63)) {
                    case -909:
                        v12 := GenS1{F0: 42, F1: 40}
                        _ = v12
                    case 1:
                        if ((1 == 670) || false) {
                            v3 = v3
                            var v13 int8 = 127
                            _ = v13
                        }
                    default:
                        var v14 int16 = 42
                        _ = v14
                    }
                    v15 := GenS2{F0: 1}
                    _ = v15
                }
                _ = fi11
            }
            _ = fi8
        }
        switch (("hello" + "foo bar") + "\n\t") {
        case "hello":
            delete(v4, "foo bar")
        case "\n\t":
            v16 := GenS3{F0: v7}
            _ = v16
        case "a":
            v3 = float32((1e10 * (0.0 * float64(float32(0.0)))))
        default:
            switch "\n\t" {
            case "\n\t":
                v2 = (v2 ^ v2)
            case "hello":
                for fi17 := 0; fi17 < 1; fi17++ {
                    v1[(v7 & v7)] = "hello"
                    _ = fi17
                }
            case "":
                if (-9223372036854775808 <= -1) {
                    switch "a" {
                    case "hello":
                        v3 = v3
                    default:
                        if (42 > (^(v2))) {
                            for fi18 := 0; fi18 < 9; fi18++ {
                                var v19 bool = ((9223372036854775807 == (v2 | v2)) || (100 > 42))
                                fmt.Println(v2, fi18, v3)
                                _ = fi18
                                _ = v19
                            }
                            var v20 int32 = 2147483647
                            _ = v20
                        } else {
                            switch 2147483647 {
                            case 42:
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    switch -1 {
                                    case -2147483648:
                                        v22 := GenS3{F0: 374}
                                        _ = v22
                                    case 2147483647:
                                        var v23 string = "\n\t"
                                        _ = v23
                                    default:
                                        if (("" + "hello") >= ("hello" + "foo bar")) {
                                            _ = len(v5)
                                            if false {
                                                fmt.Println(v3, v3, v3)
                                                if ((false || (fi21 != fi21)) && ((fi21 % ((792) | 1)) >= (fi21 << ((1) & 63)))) {
                                                    if false {
                                                        for fi24 := 0; fi24 < 13; fi24++ {
                                                            fmt.Println(v3, v7)
                                                            if false {
                                                                v2 = (v2 | 42)
                                                            } else {
                                                                var v25 float64 = (((1.5 * 1e10) + (1e10 / 1.5)) * 0.1)
                                                                _ = v25
                                                            }
                                                            _ = fi24
                                                        }
                                                    } else {
                                                        v7 = 1
                                                    }
                                                    for fi26 := 0; fi26 < 20; fi26++ {
                                                        _ = len(v5)
                                                        if ((fi21 << ((1) & 63)) != fi26) {
                                                            v7 += v7
                                                        }
                                                        v27 := GenS4{F0: (v2 % ((v2) | 1)), F1: fi26}
                                                        _ = fi26
                                                        _ = v27
                                                    }
                                                } else {
                                                    var v28 uint = 4294967295
                                                    v7 -= 4294967295
                                                    var v29 uint64 = 161
                                                    _ = v28
                                                    _ = v29
                                                }
                                            }
                                        } else {
                                            for fi30 := 0; fi30 < 4; fi30++ {
                                                var v31 int64 = (v2 << ((842) & 63))
                                                _ = fi30
                                                _ = v31
                                            }
                                            v32 := GenS5{F0: -1, F1: 6, F2: 9223372036854775807}
                                            v5 = append(v5, (v7 ^ v7))
                                            _ = v32
                                        }
                                    }
                                    _ = fi21
                                }
                            case 127:
                                v2 = (v2 + (v2 ^ v2))
                            default:
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    v7 = v7
                                    _ = fi33
                                }
                            }
                            var v34 int8 = 42
                            _ = v34
                        }
                    }
                    var v35 uint64 = 2
                    fmt.Println(v2, v2, v7)
                    _ = v35
                }
            default:
                v36 := GenS1{F0: 245, F1: 1}
                _ = v36
            }
        }
        _ = v7
    }
    if false {
        v3 = v3
        for fi37 := 0; fi37 < 2; fi37++ {
            for fi38 := 0; fi38 < 8; fi38++ {
                if (!(float64(1.5) < 0.1)) {
                    switch (fi38 / ((32768) | 1)) {
                    case -720:
                        v2 = v2
                    default:
                        v2 = v2
                    }
                } else {
                    var v39 bool = (v2 < v2)
                    switch ("foo bar" + ("" + "foo bar")) {
                    case "a":
                        var v40 uint64 = 101
                        _ = v40
                    case "hello":
                        switch -927 {
                        case 4294967295:
                            v2 = v2
                        case 1:
                            v2 = 42
                        case 504:
                            v39 = v39
                        default:
                            var v41 int = fi38
                            _ = v41
                        }
                    default:
                        if (255 != 1) {
                            if (!true) {
                                v39 = (42 <= 255)
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                                v2 -= (^(v2))
                            }
                            var v42 uint = 0
                            _ = v42
                        } else {
                            v4["foo bar"] = 2147483648
                            var v43 int = (^(fi37))
                            _ = v43
                        }
                    }
                    _ = v39
                }
                switch (("hello" + ((("" + "foo bar") + "foo bar") + "hello")) + ("\n\t" + ("foo bar" + "a"))) {
                case "a":
                    v44 := GenS6{F0: 23, F1: 32767}
                    _ = v44
                }
                _ = fi38
            }
            var v45 uint16 = 100
            _ = fi37
            _ = v45
        }
    } else {
        var v46 uint8 = 127
        if (-958 <= -1) {
            v46 += v46
            var v47 int8 = -128
            switch "a" {
            case "foo bar":
                if (-2.25 > float64((0.0 - (v3 * v3)))) {
                    if false {
                        var v48 float64 = 3.14159
                        var v49 int8 = (v47 + -128)
                        if ((1.5 <= float32(v48)) && true) {
                            v46 = (v46 - v46)
                            for fi50 := 0; fi50 < 10; fi50++ {
                                fmt.Println(v3)
                                _ = fi50
                            }
                            v46 -= (v46 + (v46 / ((v46) | 1)))
                        } else {
                            fmt.Println(v3)
                        }
                        _ = v48
                        _ = v49
                    }
                } else {
                    fmt.Println(v47)
                    switch (("foo bar" + "") + ("a" + "\n\t")) {
                    case "foo bar":
                        var v51 uint64 = 42
                        _ = v51
                    default:
                        switch 32768 {
                        case 4294967295:
                            fmt.Println(v46)
                        default:
                            v1[2147483648] = ("" + "\n\t")
                        }
                    }
                    fmt.Println(v3, v2, v47)
                }
            default:
                v3 = -2.25
            }
            _ = v47
        } else {
            v3 -= 0.0
            var v52 uint32 = 1
            var v53 uint32 = (v52 & (^(v52)))
            _ = v52
            _ = v53
        }
        _ = v46
    }
    v4[("hello" + ("foo bar" + "\n\t"))] = 4294967295
    v1[202] = "hello"
    v3 -= ((v3 + (float32(-2.25) + 0.1)) / 0.1)
    if false {
        var v54 int8 = 1
        var v55 float64 = -2.25
        _ = len(v5)
        _ = v54
        _ = v55
    } else {
        if false {
            v56 := GenS3{F0: 42}
            _ = v56
        } else {
            for fi57 := 0; fi57 < 2; fi57++ {
                if true {
                    v58 := GenS7{F0: -16, F1: -1, F2: -2147483648}
                    _ = v58
                } else {
                    switch fi57 {
                    case 2147483647:
                        var v59 int64 = v2
                        _ = v59
                    case 42:
                        v3 += 0.1
                    }
                }
                var v60 float32 = v3
                _ = fi57
                _ = v60
            }
            for fi61 := 0; fi61 < 9; fi61++ {
                v2 = (v2 >> ((0) & 63))
                _ = fi61
            }
            v62 := GenS6{F0: 127, F1: 0}
            _ = v62
        }
        v2 = 42
    }
    var v63 string = "foo bar"
    switch 2147483647 {
    case 128:
        if false {
            if true {
                fmt.Println(v3, v3)
                if false {
                    v1[4294967295] = v63
                }
                for fi64 := 0; fi64 < 4; fi64++ {
                    v5 = append(v5, 2147483647)
                    _ = fi64
                }
            }
            var v65 uint32 = 0
            v3 -= 1e10
            _ = v65
        } else {
            v63 = v63
            if false {
                _ = len(v5)
            }
            v3 -= v3
        }
    default:
        fmt.Println(v2, v63, v63)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
