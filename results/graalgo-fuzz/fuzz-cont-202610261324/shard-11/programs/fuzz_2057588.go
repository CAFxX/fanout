package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
type GenS3 struct { F0 int32; F1 int64; F2 uint64 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool; F1 int16; F2 int64 }
type GenS10 struct { F0 int16; F1 int8; F2 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 uint = 18446744073709551615
    var v3 string = ""
    v4 := make(map[uint32]string)
    var v5 uint32 = 2147483648
    var v6 uint8 = v1
    fmt.Println(v3)
    var v7 uint64 = 1
    switch v3 {
    case "hello":
        if (v1 > v6) {
            v3 = "foo bar"
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{2147483647, (fi8 % (((fi8 * fi8)) | 1)), (fi8 + fi8), fi8}
                v10 := GenS1{F0: 0, F1: v5, F2: (v1 & v6)}
                fmt.Println(fi8, fi8, v5)
                _ = fi8
                _ = v10
                _ = v9
            }
            if false {
                for fi11 := 0; fi11 < 12; fi11++ {
                    v4[636] = ((v3 + v3) + v3)
                    v12 := GenS2{F0: -1, F1: 9223372036854775807}
                    _ = fi11
                    _ = v12
                }
                var v13 string = "a"
                _ = v13
            }
        }
    case "\n\t":
        if true {
            if (925 != 905) {
                if (!false) {
                    var v14 int32 = 929
                    _ = v14
                } else {
                    fmt.Println(v3, v6, v3)
                    switch (("" + (v3 + ("foo bar" + "a"))) + v3) {
                    case "":
                        v15 := GenS2{F0: 100, F1: 532}
                        _ = v15
                    default:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            for fi17 := 0; fi17 < 4; fi17++ {
                                if (v3 < v3) {
                                    _ = len(v4)
                                    if true {
                                        v18 := GenS3{F0: 0, F1: -343, F2: 618}
                                        if true {
                                            var v19 float32 = float32(3.14159)
                                            _ = v19
                                        } else {
                                            var v20 int8 = 100
                                            var v21 uint32 = 2147483647
                                            v6 = (v6 | uint8(v7))
                                            _ = v20
                                            _ = v21
                                        }
                                        _ = v18
                                    } else {
                                        v7 += (v7 + (v7 & v7))
                                        v22 := GenS4{F0: (fi16 - fi17), F1: 0}
                                        _ = v22
                                    }
                                    var v23 int8 = 127
                                    _ = v23
                                }
                                var v24 int64 = -354
                                var v25 int8 = 127
                                _ = fi17
                                _ = v24
                                _ = v25
                            }
                            switch (v3 + "foo bar") {
                            case "\n\t":
                                switch v3 {
                                case "\n\t":
                                    v26 := GenS5{F0: 255}
                                    _ = v26
                                case "foo bar":
                                    fmt.Println(fi16, v7)
                                case "a":
                                    v27 := []int{(^(fi16)), fi16, (^(fi16)), 32767}
                                    _ = v27
                                }
                            case "a":
                                var v28 float64 = (float64(3.14159) * 0.0)
                                _ = v28
                            case "":
                                v29 := []int{fi16, fi16, (-(fi16))}
                                _ = v29
                            default:
                                if (!true) {
                                    _ = len(v4)
                                    fmt.Println(fi16, fi16, v3)
                                } else {
                                    v6 += 2
                                    v4[4294967295] = ((("hello" + v3) + (v3 + "hello")) + (v3 + v3))
                                    v30 := GenS6{F0: uint(v6), F1: v3}
                                    _ = v30
                                }
                            }
                            switch "" {
                            case "hello":
                                v31 := GenS7{F0: (v5 + v5)}
                                _ = v31
                            case "":
                                var v32 int16 = -752
                                _ = v32
                            }
                            _ = fi16
                        }
                    }
                }
                switch 9223372036854775807 {
                case 2147483647:
                    var v33 float32 = (float32(1.5) * 1e10)
                    _ = v33
                case 797:
                    switch 2147483647 {
                    case 1:
                        var v34 int8 = -128
                        _ = v34
                    }
                }
            }
            if (v3 <= v3) {
                var v35 bool = true
                _ = v35
            }
        }
    case "a":
        var v36 int = -778
        _ = v36
    }
    v1 = (^(v6))
    for fi37 := 0; fi37 < 5; fi37++ {
        if (float64((-2.25 / (-2.25 * 1.5))) < ((((3.14159 / 1.0) / -2.25) + 0.0) / float64(1.5))) {
            fmt.Println(v6)
            v5 = 141
        }
        v38 := GenS5{F0: (v1 | v6)}
        switch "" {
        case "foo bar":
            v3 = (v3 + "a")
        case "a":
            v4[0] = "a"
        default:
            v7 += (v7 | (v7 / ((501) | 1)))
        }
        _ = fi37
        _ = v38
    }
    v1 -= (v6 | v1)
    if true {
        for fi39 := 0; fi39 < 11; fi39++ {
            fmt.Println(v2, fi39)
            v5 = (v5 / (((^(v5))) | 1))
            _ = fi39
        }
        for fi40 := 0; fi40 < 11; fi40++ {
            var v41 float64 = 0.0
            _ = fi40
            _ = v41
        }
        v42 := GenS1{F0: 233, F1: (v5 ^ (v5 % ((v5) | 1))), F2: (v1 * 255)}
        _ = v42
    } else {
        v6 += v6
        if true {
            for fi43 := 0; fi43 < 16; fi43++ {
                if false {
                    v44 := []int{}
                    v45 := GenS8{F0: (v7 ^ 18446744073709551615), F1: false}
                    _ = v44
                    _ = v45
                } else {
                    switch v3 {
                    case "foo bar":
                        for fi46 := 0; fi46 < 20; fi46++ {
                            v47 := GenS5{F0: v6}
                            _ = fi46
                            _ = v47
                        }
                    default:
                        v48 := GenS7{F0: uint32(v7)}
                        _ = v48
                    }
                    v2 = v2
                }
                v2 = v2
                _ = fi43
            }
            var v49 uint32 = (^(v5))
            _ = v49
        } else {
            v7 += (^(v7))
            var v50 int64 = -762
            switch v3 {
            case "foo bar":
                if false {
                    if (2147483647 != 42) {
                        if (906 < 42) {
                            var v51 uint = 42
                            _ = v51
                        }
                    }
                    switch v3 {
                    case "foo bar":
                        if ("foo bar" <= (v3 + "")) {
                            var v52 string = ((v3 + (v3 + ((v3 + v3) + "hello"))) + v3)
                            v53 := GenS2{F0: 42, F1: (-(v50))}
                            if true {
                                fmt.Println(v3)
                                v54 := GenS9{F0: true, F1: -951, F2: 9223372036854775807}
                                _ = v54
                            }
                            _ = v52
                            _ = v53
                        }
                    case "hello":
                        if true {
                            v1 = v6
                            switch "hello" {
                            case "hello":
                                if false {
                                    v55 := GenS8{F0: 1, F1: (1.5 >= (0.1 - 1e10))}
                                    _ = v55
                                } else {
                                    v5 += v5
                                }
                            case "a":
                                var v56 int = 9223372036854775807
                                _ = v56
                            default:
                                if true {
                                    if false {
                                        delete(v4, (v5 + (v5 ^ v5)))
                                    }
                                    fmt.Println(v50)
                                }
                            }
                            v6 -= v1
                        } else {
                            for fi57 := 0; fi57 < 20; fi57++ {
                                var v58 int8 = 66
                                var v59 string = "hello"
                                _ = fi57
                                _ = v58
                                _ = v59
                            }
                            v60 := GenS10{F0: -1, F1: 127, F2: 4294967295}
                            for fi61 := 0; fi61 < 13; fi61++ {
                                v7 -= v7
                                switch fi61 {
                                case -491:
                                    fmt.Println(v6, v50, v5)
                                case -642:
                                    v1 = (v6 / ((128) | 1))
                                case 42:
                                    var v62 int8 = 127
                                    _ = v62
                                }
                                v3 = v3
                                _ = fi61
                            }
                            _ = v60
                        }
                    case "\n\t":
                        var v63 bool = (65535 <= v2)
                        _ = v63
                    }
                } else {
                    if ((v50 / (((v50 + (v50 - (v50 / ((1) | 1))))) | 1)) >= (v50 | v50)) {
                        var v64 int8 = 77
                        var v65 bool = (!(32768 != -9223372036854775808))
                        _ = v64
                        _ = v65
                    }
                    for fi66 := 0; fi66 < 2; fi66++ {
                        var v67 int16 = 1000
                        fmt.Println(v67, v50)
                        _ = fi66
                        _ = v67
                    }
                }
            case "a":
                var v68 int = 128
                _ = v68
            case "\n\t":
                for fi69 := 0; fi69 < 18; fi69++ {
                    var v70 int8 = 32
                    for fi71 := 0; fi71 < 13; fi71++ {
                        v72 := GenS5{F0: v6}
                        v73 := []int{127, fi71}
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                    _ = fi69
                    _ = v70
                }
            default:
                v74 := GenS7{F0: (v5 % (((v5 % ((v5) | 1))) | 1))}
                _ = v74
            }
            _ = v50
        }
    }
    var v75 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v75
}
