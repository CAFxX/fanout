package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64) uint8 {
    p1 -= (-(p1))
    p0 = 251
    _ = p0
    _ = p1
    return p0
}
func genFn1() int16 {
    v4 := GenS1{F0: 4294967295, F1: "\n\t"}
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5)
        for fi6 := 0; fi6 < 2; fi6++ {
            if true {
                var v7 uint16 = 42
                var v8 string = "a"
                _ = v7
                _ = v8
            } else {
                var v9 int32 = -625
                v9 -= -220
                switch "" {
                case "foo bar":
                    var v10 uint16 = 0
                    _ = v10
                case "":
                    v9 -= 42
                default:
                    switch (fi5 % ((128) | 1)) {
                    case 32767:
                        v9 = (v9 * v9)
                        var v11 int64 = -1
                        _ = v11
                    default:
                        if (!true) {
                            var v12 int = (-(fi5))
                            var v13 uint32 = 42
                            switch fi5 {
                            case 128:
                                v14 := []int{}
                                _ = v14
                            case 42:
                                v13 -= 4294967295
                            case 9223372036854775807:
                                v15 := []int{128, (^(fi6)), (fi5 / (((-(fi5))) | 1)), fi5}
                                _ = v15
                            default:
                                fmt.Println(v12, v13, v9)
                            }
                            _ = v12
                            _ = v13
                        } else {
                            v16 := GenS2{F0: 829}
                            v17 := []int{331, fi5, 9223372036854775807, fi5}
                            switch "a" {
                            case "foo bar":
                                v18 := GenS1{F0: 18446744073709551615, F1: "hello"}
                                _ = v18
                            case "":
                                var v19 float64 = 3.14159
                                _ = v19
                            case "hello":
                                var v20 int16 = -541
                                _ = v20
                            default:
                                fmt.Println(fi5, fi6)
                            }
                            _ = v16
                            _ = v17
                        }
                    }
                }
                _ = v9
            }
            v21 := []int{}
            _ = fi6
            _ = v21
        }
        v22 := []int{}
        _ = fi5
        _ = v22
    }
    if true {
        if false {
            var v23 uint = 0
            v23 = v23
            _ = v23
        }
        for fi24 := 0; fi24 < 19; fi24++ {
            switch (fi24 % (((fi24 + fi24)) | 1)) {
            case 128:
                var v25 float64 = 1e10
                _ = v25
            }
            _ = fi24
        }
    } else {
        switch 9223372036854775807 {
        case 127:
            if false {
                var v26 int16 = 42
                v26 = v26
                var v27 uint = 1
                _ = v26
                _ = v27
            }
        default:
            var v28 int8 = 1
            _ = v28
        }
        if (!(42 <= -101)) {
            v29 := GenS1{F0: 42, F1: "foo bar"}
            var v30 int16 = 42
            v30 = v30
            _ = v29
            _ = v30
        } else {
            switch "hello" {
            case "hello":
                if true {
                    switch ("hello" + ("hello" + "a")) {
                    case "":
                        if true {
                            var v31 int64 = 9223372036854775807
                            for fi32 := 0; fi32 < 2; fi32++ {
                                v31 -= (v31 - (v31 - (v31 + v31)))
                                var v33 uint16 = 42
                                _ = fi32
                                _ = v33
                            }
                            _ = v31
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if (!false) {
                                    var v34 int8 = 1
                                    _ = v34
                                }
                            default:
                                var v35 float32 = 0.1
                                _ = v35
                            }
                        }
                    case "a":
                        v36 := []int{32767}
                        _ = v36
                    case "hello":
                        var v37 int = -1
                        _ = v37
                    default:
                        switch -1 {
                        case 128:
                            v38 := []int{-1, 9223372036854775807, 2147483647}
                            _ = v38
                        default:
                            var v39 int8 = 42
                            _ = v39
                        }
                    }
                    var v40 int64 = -1
                    _ = v40
                }
            }
            var v41 uint = 1
            _ = v41
        }
    }
    _ = v4
    return 674
}

func main() {
    v1 := []string{"a", "\n\t", "a", ""}
    v2 := []string{}
    var v3 int16 = -32768
    _ = len(v2)
    if (genFn0(127, 42) == 0) {
        for fi42 := 0; fi42 < 5; fi42++ {
            switch ((("hello" + ("\n\t" + ("hello" + "foo bar"))) + "a") + ("foo bar" + ("" + ("hello" + (("hello" + "foo bar") + ("foo bar" + "foo bar")))))) {
            case "":
                v3 = genFn1()
            default:
                var v43 int = 32768
                _ = v43
            }
            v44 := GenS1{F0: 71, F1: "hello"}
            switch "" {
            case "":
                v45 := GenS2{F0: 65535}
                _ = v45
            case "a":
                var v46 int8 = 127
                _ = v46
            case "foo bar":
                for fi47 := 0; fi47 < 15; fi47++ {
                    fmt.Println(fi42, v3)
                    _ = fi47
                }
            default:
                _ = len(v2)
            }
            _ = fi42
            _ = v44
        }
        v3 = v3
        var v48 float64 = 3.14159
        _ = v48
    }
    if len(v1) > 0 {
        v1[0] = ("a" + "a")
    }
    v3 = (v3 / (((v3 << ((255) & 63))) | 1))
    if ((0.0 * (float32(1e10) / 0.1)) >= 0.0) {
        if ((!false) && (!true)) {
            v3 += v3
        } else {
            var v49 int32 = 63
            _ = v49
        }
        v3 += (^(v3))
    }
    if false {
        v3 += (v3 * v3)
    } else {
        if (-128 != 24) {
            switch "foo bar" {
            case "hello":
                v3 = v3
            case "":
                var v50 uint8 = 255
                _ = v50
            case "a":
                var v51 string = "\n\t"
                _ = v51
            default:
                if true {
                    switch -27 {
                    case 4294967295:
                        if (true || (0 == 1)) {
                            fmt.Println(v3)
                            v3 = v3
                            if true {
                                var v52 uint = 65535
                                _ = v52
                            } else {
                                if ((float32(1.5) <= 0.0) || (255 >= 718)) {
                                    var v53 uint = 834
                                    v54 := GenS3{F0: (v53 << ((v53) & 63)), F1: 127, F2: 127}
                                    _ = v53
                                    _ = v54
                                }
                                var v55 int = 1
                                _ = v55
                            }
                        } else {
                            var v56 int8 = 0
                            var v57 uint64 = 4294967295
                            _ = v56
                            _ = v57
                        }
                    case 128:
                        for fi58 := 0; fi58 < 13; fi58++ {
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            _ = fi58
                        }
                    case 257:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        if (9223372036854775807 == 1) {
                            var v59 uint8 = genFn0(1, -9223372036854775808)
                            v2 = append(v2, (("hello" + "\n\t") + ""))
                            for fi60 := 0; fi60 < 1; fi60++ {
                                if false {
                                    v59 -= (v59 / (((v59 * v59)) | 1))
                                } else {
                                    var v61 int8 = 127
                                    _ = v61
                                }
                                var v62 uint = 4294967295
                                fmt.Println(v3)
                                _ = fi60
                                _ = v62
                            }
                            _ = v59
                        } else {
                            if true {
                                v3 += (v3 | v3)
                                var v63 uint64 = 902
                                v3 = -1
                                _ = v63
                            } else {
                                switch "\n\t" {
                                case "":
                                    switch -2147483648 {
                                    case 42:
                                        v64 := GenS4{F0: -554, F1: true, F2: 382}
                                        _ = v64
                                    default:
                                        v65 := GenS5{F0: 65535, F1: 115, F2: 728}
                                        _ = v65
                                    }
                                case "foo bar":
                                    v3 = (^(v3))
                                }
                                v3 += 1
                            }
                            v66 := GenS6{F0: -90}
                            _ = v66
                        }
                    }
                } else {
                    v3 = (v3 - (v3 & -1))
                }
            }
            switch 0 {
            case 1:
                for fi67 := 0; fi67 < 10; fi67++ {
                    var v68 float32 = 0.0
                    _ = fi67
                    _ = v68
                }
            default:
                v69 := GenS7{F0: (v3 / ((v3) | 1))}
                _ = v69
            }
        }
        for fi70 := 0; fi70 < 6; fi70++ {
            var v71 string = "a"
            fmt.Println(v71)
            for fi72 := 0; fi72 < 15; fi72++ {
                fmt.Println(fi72, v71)
                v71 = (v71 + (v71 + "\n\t"))
                var v73 uint64 = 42
                _ = fi72
                _ = v73
            }
            _ = fi70
            _ = v71
        }
    }
    v3 = v3
    v3 = v3
    v3 -= v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
