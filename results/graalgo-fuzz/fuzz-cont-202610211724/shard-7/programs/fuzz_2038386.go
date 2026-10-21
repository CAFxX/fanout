package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint; F1 int64; F2 int }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 int8; F1 int; F2 uint8 }
func genFn0() int8 {
    var v4 uint = 0
    fmt.Println(v4)
    v4 += (v4 | 787)
    v5 := []int{32768, 1, -394, -2147483648}
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 float64 = 0.0
    v2 := []string{"\n\t", (("foo bar" + "hello") + "hello"), "\n\t", "a"}
    var v3 uint64 = 0
    if true {
        var v6 float32 = float32((3.14159 / float64(1e10)))
        _ = v6
    }
    if true {
        v3 += v3
        switch "" {
        case "foo bar":
            if (0.0 > ((0.0 * 0.1) + 3.14159)) {
                for fi7 := 0; fi7 < 15; fi7++ {
                    v8 := GenS1{F0: 37, F1: 0, F2: 187}
                    _ = fi7
                    _ = v8
                }
                if (97 >= genFn0()) {
                    for fi9 := 0; fi9 < 16; fi9++ {
                        switch "\n\t" {
                        case "":
                            for fi10 := 0; fi10 < 6; fi10++ {
                                fmt.Println(v1, v1)
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                _ = fi10
                            }
                        }
                        _ = fi9
                    }
                    if ((false || ((!(!true)) || ((!true) || ((true && (!false)) && ((false && true) || ((v1 * v1) <= v1)))))) && false) {
                        if len(v2) > 0 {
                            v2[0] = ((("\n\t" + ("hello" + "")) + "") + "hello")
                        }
                        for fi11 := 0; fi11 < 10; fi11++ {
                            _ = len(v2)
                            v12 := GenS1{F0: 2147483647, F1: 2147483648, F2: -428}
                            fmt.Println(v1, fi11, v1)
                            _ = fi11
                            _ = v12
                        }
                        v1 += (v1 * (1.5 / 1.0))
                    } else {
                        v13 := GenS2{F0: 954}
                        if (1e10 < float32(1e10)) {
                            var v14 uint64 = v3
                            fmt.Println(v1)
                            fmt.Println(v3, v14, v14)
                            _ = v14
                        }
                        _ = v13
                    }
                    v1 = v1
                } else {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v3 = (v3 / (((^(v3))) | 1))
                        v1 -= v1
                        _ = fi15
                    }
                    v1 -= 1.5
                }
            } else {
                v3 = (v3 + v3)
            }
        case "hello":
            switch "" {
            case "\n\t":
                v1 = float64(-2.25)
            case "a":
                var v16 int64 = 1
                _ = v16
            case "hello":
                var v17 uint = 807
                _ = v17
            default:
                if true {
                    v1 -= v1
                    v18 := GenS3{F0: "", F1: (("" + "a") + (("" + ("a" + ("hello" + ("\n\t" + "foo bar")))) + "hello")), F2: 2147483648}
                    v3 -= v3
                    _ = v18
                }
            }
        default:
            var v19 int = int(v3)
            _ = v19
        }
    } else {
        v20 := GenS4{F0: 42, F1: 2}
        for fi21 := 0; fi21 < 15; fi21++ {
            switch "" {
            case "\n\t":
                if true {
                    var v22 bool = (940 < 341)
                    _ = v22
                } else {
                    var v23 int32 = -903
                    var v24 int64 = 1
                    switch 32768 {
                    case 128:
                        var v25 float64 = 0.1
                        _ = v25
                    case 42:
                        var v26 int32 = v23
                        _ = v26
                    default:
                        v23 = int32(fi21)
                    }
                    _ = v23
                    _ = v24
                }
            default:
                v3 -= (v3 - (^(v3)))
            }
            v2 = append(v2, "\n\t")
            _ = fi21
        }
        v27 := GenS2{F0: 440}
        _ = v20
        _ = v27
    }
    var v28 uint32 = 163
    v1 = float64((1e10 * -2.25))
    var v29 uint8 = uint8(v28)
    switch "foo bar" {
    case "\n\t":
        for fi30 := 0; fi30 < 1; fi30++ {
            fmt.Println(v29, v28)
            switch "" {
            case "a":
                v31 := GenS2{F0: -201}
                _ = v31
            case "foo bar":
                fmt.Println(v28)
            }
            var v32 int64 = 0
            _ = fi30
            _ = v32
        }
    case "a":
        var v33 uint = 4294967295
        _ = v33
    default:
        var v34 uint = 255
        _ = v34
    }
    if (!true) {
        var v35 int8 = 42
        if true {
            if true {
                v3 += 18446744073709551615
                v3 = v3
            } else {
                switch "a" {
                case "":
                    var v36 int = -477
                    _ = v36
                case "a":
                    fmt.Println(v1, v29)
                case "hello":
                    _ = len(v2)
                default:
                    v35 = (v35 | v35)
                }
            }
            _ = len(v2)
        } else {
            fmt.Println(v28, v29)
            v37 := GenS5{F0: (^(v29)), F1: v29}
            _ = v37
        }
        _ = v35
    }
    var v38 uint64 = (v3 & v3)
    if (v29 <= v29) {
        v1 = (v1 / (0.1 + v1))
        if ("foo bar" < "\n\t") {
            var v39 float64 = v1
            _ = v39
        } else {
            var v40 uint32 = (^(v28))
            v41 := GenS3{F0: ("foo bar" + ""), F1: "a", F2: 725}
            _ = v40
            _ = v41
        }
    }
    var v42 uint = 2
    v38 -= 2
    var v43 int16 = 32767
    switch (("hello" + "a") + "hello") {
    case "hello":
        if false {
            v28 -= (v28 << (((v42 >> ((2) & 63))) & 63))
            switch 32767 {
            case 384:
                v1 = (1e10 / ((((float64(3.14159) * 1e10) / (v1 - v1)) / (v1 + v1)) - v1))
            case 2147483647:
                fmt.Println(v1, v42, v38)
            }
            switch "foo bar" {
            case "foo bar":
                var v44 int64 = 340
                _ = v44
            case "hello":
                v28 = 65535
            default:
                if true {
                    fmt.Println(v28, v29, v38)
                    var v45 uint8 = (v29 >> ((v42) & 63))
                    fmt.Println(v28, v28)
                    _ = v45
                } else {
                    switch -1 {
                    case 128:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -2147483648:
                        if ((v3 % ((uint64(v29)) | 1)) >= v38) {
                            if (float64(1e10) == v1) {
                                v43 = v43
                                if (1 >= 100) {
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = ""
                                        }
                                        for fi46 := 0; fi46 < 15; fi46++ {
                                            fmt.Println(v3, v28)
                                            var v47 uint8 = v29
                                            v48 := GenS2{F0: -992}
                                            _ = fi46
                                            _ = v47
                                            _ = v48
                                        }
                                        var v49 string = "a"
                                        _ = v49
                                    } else {
                                        switch 334 {
                                        case -2147483648:
                                            v29 = v29
                                        case -9223372036854775808:
                                            v1 = (0.0 * v1)
                                        default:
                                            for fi50 := 0; fi50 < 1; fi50++ {
                                                v51 := GenS6{F0: (v42 / ((v42) | 1)), F1: 42, F2: fi50}
                                                _ = len(v2)
                                                var v52 string = "a"
                                                _ = fi50
                                                _ = v51
                                                _ = v52
                                            }
                                        }
                                        v43 -= (v43 & 32767)
                                    }
                                } else {
                                    var v53 float32 = 1.5
                                    v3 = v3
                                    _ = v53
                                }
                            }
                        } else {
                            switch ("hello" + (("" + "hello") + "hello")) {
                            case "a":
                                var v54 uint32 = 378
                                _ = v54
                            case "":
                                _ = len(v2)
                            default:
                                v28 += (v28 / (((v28 % ((v28) | 1))) | 1))
                            }
                            for fi55 := 0; fi55 < 1; fi55++ {
                                if (!(!false)) {
                                    if (!(!true)) {
                                        v56 := GenS7{F0: 516, F1: 887}
                                        fmt.Println(v38)
                                        if (false && false) {
                                            _ = len(v2)
                                            v57 := GenS8{F0: 127, F1: fi55, F2: v29}
                                            _ = v57
                                        }
                                        _ = v56
                                    }
                                }
                                _ = fi55
                            }
                        }
                    default:
                        for fi58 := 0; fi58 < 13; fi58++ {
                            v2 = append(v2, "\n\t")
                            var v59 int8 = 127
                            var v60 string = ("foo bar" + "hello")
                            _ = fi58
                            _ = v59
                            _ = v60
                        }
                    }
                    if false {
                        var v61 int64 = -9223372036854775808
                        _ = v61
                    }
                    v28 += v28
                }
            }
        } else {
            v62 := GenS4{F0: 2147483647, F1: 1}
            var v63 int8 = -1
            _ = v62
            _ = v63
        }
    }
    var v64 int32 = 2147483647
    var v65 float64 = -2.25
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v42
    _ = v43
    _ = v64
    _ = v65
}
