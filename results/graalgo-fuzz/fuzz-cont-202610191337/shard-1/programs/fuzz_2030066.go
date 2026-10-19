package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int32; F2 uint }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{""}
    var v3 uint32 = 2147483648
    v4 := make(map[int]bool)
    var v5 float32 = (0.0 / 1.5)
    v6 := make(map[string]bool)
    var v7 float64 = ((0.0 + -2.25) - 0.0)
    var v8 uint16 = 32767
    switch 2147483647 {
    case 42:
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v7)
            if ("hello" != "") {
                v5 += 3.14159
                for fi10 := 0; fi10 < 6; fi10++ {
                    if ("" >= "foo bar") {
                        fmt.Println(v3, v8)
                    } else {
                        v8 += 2
                        var v11 string = ""
                        v12 := GenS1{F0: v3, F1: true, F2: (v3 ^ v3)}
                        _ = v11
                        _ = v12
                    }
                    _ = fi10
                }
                var v13 int64 = -807
                _ = v13
            }
            _ = fi9
        }
    }
    if false {
        if true {
            switch "hello" {
            case "hello":
                for fi14 := 0; fi14 < 1; fi14++ {
                    fmt.Println(v3)
                    _ = fi14
                }
            }
            fmt.Println(v3)
        } else {
            switch "hello" {
            case "foo bar":
                for fi15 := 0; fi15 < 6; fi15++ {
                    for fi16 := 0; fi16 < 16; fi16++ {
                        if false {
                            delete(v1, 432)
                        } else {
                            if true {
                                if false {
                                    switch 42 {
                                    case -1:
                                        switch (fi16 ^ (fi15 ^ (fi16 | fi15))) {
                                        case 32767:
                                            if (!false) {
                                                var v17 uint32 = 143
                                                fmt.Println(v8, v5)
                                                _ = v17
                                            }
                                            v1[v3] = "\n\t"
                                        }
                                    case -9223372036854775808:
                                        v18 := GenS2{F0: 2}
                                        _ = v18
                                    case 42:
                                        fmt.Println(v7, v5)
                                    }
                                    var v19 int = (fi16 % ((int(v3)) | 1))
                                    var v20 uint = 2
                                    switch "hello" {
                                    case "\n\t":
                                        fmt.Println(v7, v19)
                                    case "a":
                                        delete(v6, ("" + ("" + "a")))
                                    default:
                                        v20 = uint(v8)
                                    }
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint64 = 242
                                    v4[(fi16 | (fi16 + -674))] = false
                                    switch fi16 {
                                    case -1:
                                        var v22 uint64 = (v21 / (((v21 << ((939) & 63))) | 1))
                                        _ = v22
                                    case -32:
                                        var v23 uint = 255
                                        _ = v23
                                    }
                                    _ = v21
                                }
                            } else {
                                v24 := GenS3{F0: 233, F1: 0, F2: 255}
                                _ = v24
                            }
                            if false {
                                switch ("hello" + (("\n\t" + "hello") + ("a" + "\n\t"))) {
                                case "\n\t":
                                    v8 += v8
                                    if true {
                                        if (!("\n\t" <= "a")) {
                                            v8 = (^(v8))
                                            fmt.Println(v7, fi16)
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = "foo bar"
                                            }
                                        }
                                        fmt.Println(v3, fi15, fi16)
                                    }
                                case "hello":
                                    var v25 int32 = 1
                                    _ = v25
                                case "":
                                    v6["foo bar"] = true
                                default:
                                    if false {
                                        v8 -= v8
                                    }
                                }
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                    v8 -= (v8 / ((v8) | 1))
                                    fmt.Println(v7)
                                    fmt.Println(v3, fi16, fi15)
                                } else {
                                    v5 += (1e10 / float32((((v7 + -2.25) + 1.5) + float64(v5))))
                                }
                                v26 := GenS3{F0: (v8 % (((v8 / (((^(v8))) | 1))) | 1)), F1: 2147483647, F2: 70}
                                if (((^(v3)) <= (^(v3))) || (true && ((^(v8)) <= 42))) {
                                    v7 += (float64(3.14159) + -2.25)
                                    var v27 int = (fi15 ^ (fi16 | (fi15 + (fi15 & (fi15 / (((-(fi16))) | 1))))))
                                    _ = v27
                                } else {
                                    if (false && false) {
                                        var v28 bool = true
                                        var v29 int64 = -998
                                        if false {
                                            if (0 != (fi16 << ((65535) & 63))) {
                                                var v30 float32 = float32((v7 * (v7 * float64(v5))))
                                                _ = v30
                                            } else {
                                                _ = len(v6)
                                                var v31 uint32 = v3
                                                if v28 {
                                                    var v32 uint = 24
                                                    _ = v32
                                                } else {
                                                    v33 := GenS3{F0: v8, F1: -463, F2: 0}
                                                    var v34 int32 = -932
                                                    _ = v33
                                                    _ = v34
                                                }
                                                _ = v31
                                            }
                                        }
                                        _ = v28
                                        _ = v29
                                    } else {
                                        var v35 int16 = -1
                                        _ = v35
                                    }
                                    v3 += v3
                                }
                                v3 += (v3 & v3)
                                _ = v26
                            } else {
                                if (42 > 42) {
                                    v3 += v3
                                    if true {
                                        if (32768 != fi16) {
                                            var v36 int16 = 1000
                                            var v37 int8 = -31
                                            _ = v36
                                            _ = v37
                                        } else {
                                            if false {
                                                v38 := GenS4{F0: 255}
                                                var v39 int = 32767
                                                _ = v38
                                                _ = v39
                                            }
                                            v40 := GenS4{F0: 34}
                                            var v41 uint32 = v3
                                            _ = v40
                                            _ = v41
                                        }
                                        v7 = v7
                                    }
                                    delete(v6, ((("foo bar" + ("\n\t" + ("foo bar" + ("a" + "a")))) + "foo bar") + "foo bar"))
                                }
                            }
                        }
                        if true {
                            switch ("\n\t" + ("foo bar" + ("\n\t" + (("" + "\n\t") + "foo bar")))) {
                            case "foo bar":
                                _ = len(v2)
                            case "a":
                                v5 -= v5
                            default:
                                switch fi16 {
                                case 2147483647:
                                    var v42 float32 = (v5 + (v5 + (((1e10 - 0.1) + 3.14159) / float32(0.1))))
                                    _ = v42
                                case 32767:
                                    v43 := GenS5{F0: fi16}
                                    _ = v43
                                case 617:
                                    fmt.Println(v7, v8)
                                }
                            }
                        }
                        _ = fi16
                    }
                    if (!false) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v44 := v6["hello"]
                        _ = v44
                    } else {
                        var v45 uint16 = v8
                        _ = v45
                    }
                    var v46 bool = false
                    _ = fi15
                    _ = v46
                }
            case "hello":
                for fi47 := 0; fi47 < 5; fi47++ {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v5, v7)
                    case "":
                        var v48 string = ((((("hello" + "foo bar") + ("" + "\n\t")) + "") + ("\n\t" + "\n\t")) + "\n\t")
                        _ = v48
                    default:
                        v6[""] = false
                    }
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = fi47
                }
            }
            fmt.Println(v8)
        }
        v49 := GenS6{F0: (v8 ^ (^(v8))), F1: 636, F2: 65535}
        v3 = v3
        _ = v49
    }
    v50 := GenS1{F0: 2147483647, F1: false, F2: (v3 + v3)}
    var v51 float64 = (float64(v5) + v7)
    var v52 string = ((("hello" + "\n\t") + (("hello" + "") + "foo bar")) + "\n\t")
    if false {
        for fi53 := 0; fi53 < 4; fi53++ {
            if false {
                for fi54 := 0; fi54 < 9; fi54++ {
                    v55 := GenS6{F0: (v8 ^ v8), F1: v3, F2: 462}
                    if false {
                        v5 = 3.14159
                        switch v52 {
                        case "a":
                            var v56 int32 = 1
                            _ = v56
                        }
                    } else {
                        if true {
                            if true {
                                v57 := GenS3{F0: v8, F1: -686, F2: 692}
                                _ = v57
                            } else {
                                fmt.Println(v3, fi53)
                            }
                        }
                    }
                    _ = fi54
                    _ = v55
                }
                v58 := GenS7{F0: (^(fi53))}
                if ((fi53 << ((936) & 63)) > fi53) {
                    fmt.Println(v51, v5, v3)
                    for fi59 := 0; fi59 < 13; fi59++ {
                        var v60 uint64 = 18446744073709551615
                        v61 := GenS6{F0: (v8 >> ((0) & 63)), F1: v3, F2: 2}
                        switch v52 {
                        case "\n\t":
                            var v62 float64 = v7
                            fmt.Println(v62, v8, v7)
                            v63 := GenS2{F0: 1}
                            _ = v62
                            _ = v63
                        case "hello":
                            v3 = (v3 & (v3 * (v3 << ((2) & 63))))
                        }
                        switch -9223372036854775808 {
                        case 336:
                            fmt.Println(fi53, v8)
                        case 414:
                            var v64 bool = true
                            _ = v64
                        case 42:
                            if len(v2) > 0 {
                                v2[0] = ("a" + "a")
                            }
                        default:
                            v2 = append(v2, "foo bar")
                        }
                        var v65 uint64 = uint64(v3)
                        _ = fi59
                        _ = v60
                        _ = v61
                        _ = v65
                    }
                } else {
                    v66 := GenS8{F0: 9223372036854775807, F1: 9223372036854775807}
                    _ = v66
                }
                _ = v58
            }
            if len(v2) > 0 {
                v2[0] = v52
            }
            _ = fi53
        }
        _ = len(v6)
    } else {
        v67 := GenS8{F0: -998, F1: 42}
        v68 := v6[("\n\t" + ((v52 + (v52 + v52)) + ("a" + v52)))]
        _ = v67
        _ = v68
    }
    var v69 uint64 = 930
    var v70 float32 = (v5 - (-2.25 + float32((v7 + (v51 * (v7 / v7))))))
    fmt.Println(v5, v7, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v69
    _ = v7
    _ = v70
    _ = v8
}
