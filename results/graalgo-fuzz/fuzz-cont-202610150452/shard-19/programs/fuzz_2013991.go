package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint32; F2 int16 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 bool; F1 int64; F2 int64 }
func genFn0() uint64 {
    if (!(2 == 911)) {
        v5 := []int{0, 144, 128, 128}
        var v6 int64 = -1
        v6 = (v6 | (-(v6)))
        _ = v5
        _ = v6
    }
    var v7 int32 = -1
    var v8 int32 = v7
    v9 := []int{89, 4294967295}
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}
func genFn1() string {
    var v10 uint16 = 502
    var v11 string = "hello"
    v12 := []int{-9223372036854775808}
    var v13 uint = 42
    v13 += v13
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return "hello"
}

func main() {
    var v1 int16 = 1
    var v2 uint = 255
    var v3 string = "\n\t"
    v4 := make(map[string]string)
    _ = len(v4)
    for fi14 := 0; fi14 < 16; fi14++ {
        switch v3 {
        case "hello":
            v2 -= (v2 % (((v2 % ((255) | 1))) | 1))
        case "\n\t":
            v15 := GenS1{F0: -293}
            _ = v15
        case "foo bar":
            v3 = ((v3 + v3) + "a")
        default:
            if false {
                var v16 int64 = -473
                var v17 int64 = 0
                _ = v16
                _ = v17
            } else {
                v2 -= v2
                switch v3 {
                case "a":
                    var v18 bool = true
                    _ = v18
                default:
                    switch v3 {
                    case "foo bar":
                        for fi19 := 0; fi19 < 14; fi19++ {
                            switch (v3 + v3) {
                            case "foo bar":
                                v3 = "foo bar"
                            case "a":
                                v20 := GenS2{F0: true, F1: 42, F2: (v1 + v1)}
                                _ = v20
                            case "":
                                if false {
                                    v3 = "a"
                                    _ = len(v4)
                                    v2 -= 2
                                }
                            default:
                                if false {
                                    switch fi14 {
                                    case 488:
                                        v3 = "\n\t"
                                    case 0:
                                        switch fi14 {
                                        case -9223372036854775808:
                                            if true {
                                                v3 = ("hello" + v3)
                                                var v21 float64 = -2.25
                                                _ = v21
                                            } else {
                                                var v22 int8 = 1
                                                _ = v22
                                            }
                                        default:
                                            var v23 float32 = 0.1
                                            _ = v23
                                        }
                                    case 42:
                                        v24 := GenS1{F0: 970}
                                        _ = v24
                                    default:
                                        if false {
                                            var v25 uint8 = 54
                                            fmt.Println(v25, v3, v2)
                                            _ = v25
                                        } else {
                                            v3 = ((v3 + v3) + v3)
                                            switch (v3 + ("a" + v3)) {
                                            case "a":
                                                fmt.Println(fi14, v2, v2)
                                            case "\n\t":
                                                v1 = 247
                                            case "foo bar":
                                                v1 += 42
                                                v1 -= 1000
                                                v2 += v2
                                            default:
                                                v26 := GenS1{F0: 1}
                                                _ = v26
                                            }
                                        }
                                    }
                                    var v27 uint32 = 1
                                    if ((true || (v1 > v1)) && true) {
                                        v4["\n\t"] = ((v3 + ((v3 + "hello") + ("foo bar" + "a"))) + "")
                                        v3 = v3
                                        v1 = (v1 * v1)
                                    } else {
                                        switch ("\n\t" + (v3 + (((((v3 + (v3 + v3)) + ("a" + v3)) + v3) + "") + ("\n\t" + (v3 + v3))))) {
                                        case "hello":
                                            v28 := GenS3{F0: -95, F1: fi19, F2: (v27 >> ((720) & 63))}
                                            _ = v28
                                        case "a":
                                            var v29 bool = (-1 < 1)
                                            _ = v29
                                        }
                                        var v30 float32 = 3.14159
                                        _ = v30
                                    }
                                    _ = v27
                                }
                            }
                            fmt.Println(fi14, v3)
                            v3 = v3
                            _ = fi19
                        }
                    }
                }
                v2 += 4294967295
            }
        }
        v4[v3] = (v3 + v3)
        for fi31 := 0; fi31 < 17; fi31++ {
            fmt.Println(fi31)
            var v32 bool = (100 != 1)
            switch v3 {
            case "\n\t":
                v2 = v2
            default:
                v33 := []int{fi31, (fi31 | fi14), (fi14 ^ (fi31 << (((^(v2))) & 63))), -1}
                _ = v33
            }
            var v34 int32 = 0
            _ = fi31
            _ = v32
            _ = v34
        }
        _ = fi14
    }
    var v35 int8 = 42
    v36 := GenS4{F0: 32767, F1: 1}
    v37 := []int{-2147483648, 2147483647}
    switch "hello" {
    case "a":
        for fi38 := 0; fi38 < 20; fi38++ {
            var v39 uint16 = 705
            v40 := GenS1{F0: -9223372036854775808}
            var v41 uint8 = 1
            _ = fi38
            _ = v39
            _ = v40
            _ = v41
        }
    case "foo bar":
        for fi42 := 0; fi42 < 11; fi42++ {
            v43 := GenS4{F0: 1, F1: 2}
            _ = fi42
            _ = v43
        }
    }
    if false {
        if false {
            switch v3 {
            case "hello":
                var v44 bool = true
                _ = v44
            case "foo bar":
                if (false || (0 == 65535)) {
                    fmt.Println(v1, v35)
                    var v45 uint8 = 244
                    _ = v45
                }
            case "a":
                switch 920 {
                case 9223372036854775807:
                    var v46 float64 = 1e10
                    _ = v46
                case 2147483647:
                    v47 := GenS1{F0: -573}
                    _ = v47
                case 32768:
                    fmt.Println(v1, v3, v35)
                }
            }
            switch -2147483648 {
            case 128:
                v48 := GenS2{F0: ((!(!false)) || false), F1: 178, F2: 42}
                _ = v48
            }
        } else {
            v49 := GenS3{F0: (v35 % ((-128) | 1)), F1: -1, F2: 2147483647}
            _ = v49
        }
        v50 := GenS2{F0: true, F1: 754, F2: v1}
        fmt.Println(v2)
        _ = v50
    } else {
        v51 := GenS5{F0: (!true), F1: 9223372036854775807, F2: 0}
        v2 += (v2 / ((v2) | 1))
        switch (((((v3 + "") + "hello") + "a") + ("a" + (("" + (("hello" + ((v3 + (("hello" + v3) + v3)) + ("" + v3))) + v3)) + ("foo bar" + "foo bar")))) + (v3 + "")) {
        case "":
            switch v3 {
            case "\n\t":
                var v52 int = -458
                _ = v52
            default:
                _ = len(v4)
            }
        case "\n\t":
            var v53 bool = false
            _ = v53
        }
        _ = v51
    }
    var v54 bool = (2 == genFn0())
    v55 := v4[v3]
    if v54 {
        var v56 uint16 = uint16(v2)
        if (831 < 29) {
            if (413 != 65535) {
                v1 = v1
                fmt.Println(v3)
            }
            for fi57 := 0; fi57 < 5; fi57++ {
                v54 = ((v56 / (((v56 ^ 42)) | 1)) < (v56 + v56))
                _ = fi57
            }
        }
        v2 += v2
        _ = v56
    }
    var v58 uint = v2
    v58 -= 65535
    var v59 uint16 = 1
    v58 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v54
    _ = v55
    _ = v58
    _ = v59
}
