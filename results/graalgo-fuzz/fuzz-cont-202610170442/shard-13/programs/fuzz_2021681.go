package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []uint32{2, 951, 2147483647}
    var v2 uint16 = 813
    v3 := []string{("a" + ("foo bar" + ("a" + "foo bar"))), ("a" + ("foo bar" + "hello")), (("foo bar" + "\n\t") + "foo bar"), "\n\t"}
    var v4 float64 = 1.5
    v5 := GenS1{F0: false, F1: 9223372036854775807, F2: 0}
    var v6 int = -2147483648
    if len(v1) > 0 {
        v1[0] = 127
    }
    var v7 string = "\n\t"
    if ((!false) && true) {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int = (^(v6))
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: (!false), F1: 322, F2: -1}
        _ = v10
    } else {
        fmt.Println(v7, v7, v4)
        fmt.Println(v4, v7)
    }
    switch ("a" + (v7 + ((v7 + v7) + v7))) {
    case "a":
        v2 += 32767
    default:
        var v11 int8 = 127
        _ = v11
    }
    switch v7 {
    case "":
        for fi12 := 0; fi12 < 4; fi12++ {
            fmt.Println(fi12, v2)
            switch fi12 {
            case -101:
                switch (fi12 / ((9223372036854775807) | 1)) {
                case 42:
                    if true {
                        fmt.Println(fi12)
                    }
                case 2147483647:
                    if true {
                        v6 = 1
                        fmt.Println(fi12, v2)
                        var v13 int32 = -1
                        _ = v13
                    } else {
                        v14 := GenS2{F0: 100, F1: 0, F2: -92}
                        v15 := GenS3{F0: (v6 + fi12)}
                        if (!(127 >= 255)) {
                            var v16 int64 = -9223372036854775808
                            _ = v16
                        }
                        _ = v14
                        _ = v15
                    }
                case -468:
                    v17 := GenS4{F0: -1}
                    _ = v17
                }
            case 509:
                var v18 int = -898
                _ = v18
            default:
                if (1 <= 311) {
                    fmt.Println(v7, v7)
                    v19 := GenS5{F0: false}
                    _ = v19
                } else {
                    if true {
                        v20 := GenS5{F0: (37 != -1)}
                        if (((0.0 * (float32(v4) / (-2.25 + 1e10))) / 1e10) >= (1e10 * -2.25)) {
                            v2 -= v2
                            v7 = ""
                        }
                        _ = v20
                    } else {
                        switch v6 {
                        case 4294967295:
                            v2 = 255
                        }
                    }
                    v6 -= 4294967295
                    _ = len(v1)
                }
            }
            fmt.Println(v4, v2)
            _ = fi12
        }
    case "\n\t":
        v4 += ((v4 * (3.14159 * -2.25)) + v4)
    case "a":
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 float32 = 1e10
            v23 := GenS4{F0: -32768}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi21
            _ = v22
            _ = v23
        }
    default:
        v24 := GenS1{F0: ((v7 + ((v7 + ("foo bar" + "")) + v7)) != v7), F1: (-(v6)), F2: 125}
        _ = v24
    }
    fmt.Println(v4)
    v2 = (v2 + v2)
    v25 := GenS4{F0: -802}
    v7 = v7
    switch v6 {
    case -316:
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int8 = 1
            v1 = append(v1, 2)
            switch (v7 + v7) {
            case "a":
                v27 += (v27 - v27)
            case "foo bar":
                if true {
                    var v28 uint64 = 25
                    switch v7 {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                    default:
                        v29 := GenS6{F0: 1000, F1: v7}
                        _ = v29
                    }
                    var v30 float64 = ((v4 + v4) / (v4 - v4))
                    _ = v28
                    _ = v30
                }
            case "hello":
                if false {
                    var v31 int64 = -857
                    v7 = "foo bar"
                    _ = v31
                } else {
                    if false {
                        for fi32 := 0; fi32 < 15; fi32++ {
                            switch v7 {
                            case "\n\t":
                                var v33 int64 = 0
                                _ = v33
                            case "hello":
                                v3 = append(v3, v7)
                                var v34 uint64 = 2
                                var v35 uint16 = 0
                                _ = v34
                                _ = v35
                            case "foo bar":
                                var v36 int16 = -108
                                _ = v36
                            default:
                                v37 := GenS1{F0: false, F1: fi32, F2: v27}
                                _ = v37
                            }
                            _ = fi32
                        }
                        for fi38 := 0; fi38 < 3; fi38++ {
                            var v39 int32 = 538
                            var v40 int64 = 9223372036854775807
                            v41 := GenS3{F0: (v6 << ((48) & 63))}
                            if len(v3) > 0 {
                                v3[0] = v7
                            }
                            v42 := GenS6{F0: 42, F1: v7}
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                    }
                }
            default:
                v43 := GenS7{F0: 215, F1: (((v7 + "") + (v7 + ("foo bar" + v7))) + v7)}
                _ = v43
            }
            _ = fi26
            _ = v27
        }
    default:
        v6 = (v6 + (v6 / ((42) | 1)))
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        if len(v1) > 0 {
            v1[0] = 1
        }
        fmt.Println(v2, v7)
        switch v7 {
        case "\n\t":
            for fi45 := 0; fi45 < 7; fi45++ {
                switch fi44 {
                case 32768:
                    v46 := GenS8{F0: 32767, F1: v2}
                    _ = v46
                case 393:
                    var v47 uint64 = 42
                    _ = v47
                case 203:
                    fmt.Println(v6, v7, v4)
                default:
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                }
                _ = fi45
            }
        case "foo bar":
            v4 = v4
        case "hello":
            v48 := GenS4{F0: -1}
            _ = v48
        }
        _ = fi44
    }
    if (!(false || true)) {
        v49 := GenS2{F0: -3, F1: 1000, F2: -1}
        if (4294967295 <= 2147483648) {
            if (true && true) {
                _ = len(v1)
            }
        }
        _ = v49
    } else {
        var v50 uint8 = 42
        _ = v50
    }
    var v51 int16 = 531
    var v52 int32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v7
}
