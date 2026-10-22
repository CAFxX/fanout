package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 uint32; F1 uint64 }
type GenS9 struct { F0 int32; F1 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int64; F1 uint16 }

func main() {
    var v1 string = "a"
    var v2 int16 = 1
    var v3 uint16 = 42
    v4 := []string{}
    var v5 uint32 = 2147483647
    switch 909 {
    case -744:
        if (false || true) {
            switch v1 {
            case "":
                fmt.Println(v2, v3)
            }
        } else {
            var v6 bool = false
            switch "" {
            case "\n\t":
                var v7 int = -9223372036854775808
                _ = v7
            case "hello":
                var v8 uint32 = v5
                _ = v8
            }
            fmt.Println(v6, v2)
            _ = v6
        }
    case 9223372036854775807:
        fmt.Println(v3, v2, v5)
    }
    v3 += (v3 & 32767)
    if (-2.25 > 1.5) {
        v5 = 0
        switch "\n\t" {
        case "":
            for fi9 := 0; fi9 < 9; fi9++ {
                if (!true) {
                    v3 = 255
                } else {
                    var v10 uint16 = (v3 - v3)
                    _ = v10
                }
                if (0.0 <= 0.1) {
                    v11 := GenS1{F0: v5}
                    if (false && (681 < (v5 % (((v5 >> ((2) & 63))) | 1)))) {
                        _ = len(v4)
                        v5 = 1
                    }
                    for fi12 := 0; fi12 < 17; fi12++ {
                        var v13 uint64 = 477
                        v14 := GenS2{F0: 4294967295}
                        v15 := GenS1{F0: (v5 + v5)}
                        _ = fi12
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = v11
                } else {
                    v16 := GenS3{F0: 459, F1: (v2 + 42)}
                    var v17 uint8 = 89
                    fmt.Println(v3, v3, v2)
                    _ = v16
                    _ = v17
                }
                _ = fi9
            }
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        if false {
            var v19 uint16 = v3
            if false {
                fmt.Println(v5, v19)
                for fi20 := 0; fi20 < 16; fi20++ {
                    if true {
                        if false {
                            v21 := GenS2{F0: fi20}
                            _ = v21
                        }
                        switch v1 {
                        case "foo bar":
                            var v22 uint8 = 185
                            _ = v22
                        case "hello":
                            fmt.Println(v2)
                        default:
                            switch ((v1 + "foo bar") + (v1 + v1)) {
                            case "hello":
                                v2 -= (^(v2))
                            }
                        }
                    } else {
                        fmt.Println(fi20, v3, v5)
                    }
                    _ = fi20
                }
                fmt.Println(v5, v5, v19)
            }
            switch "a" {
            case "":
                v23 := GenS2{F0: fi18}
                _ = v23
            case "\n\t":
                if true {
                    fmt.Println(v19, v19)
                    if true {
                        fmt.Println(v19, fi18, v1)
                        switch (^(fi18)) {
                        case -2147483648:
                            v1 = v1
                        case -9223372036854775808:
                            v24 := GenS1{F0: uint32(v19)}
                            _ = v24
                        case 159:
                            if true {
                                var v25 int = (fi18 >> ((2) & 63))
                                switch 4294967295 {
                                case 128:
                                    var v26 uint = 2
                                    _ = v26
                                default:
                                    v27 := GenS4{F0: (v19 | (v3 | (v19 ^ 664)))}
                                    _ = v27
                                }
                                fmt.Println(v25, v3, fi18)
                                _ = v25
                            } else {
                                var v28 int32 = -2147483648
                                fmt.Println(v5, v19)
                                _ = v28
                            }
                        }
                        fmt.Println(v19, v2, v19)
                    }
                    if true {
                        switch ((v1 + ("\n\t" + (((v1 + v1) + (("a" + ("\n\t" + ((("foo bar" + v1) + v1) + ("hello" + v1)))) + ((v1 + "foo bar") + "a"))) + ((v1 + ("hello" + v1)) + (v1 + ("foo bar" + (("\n\t" + v1) + v1))))))) + "a") {
                        case "hello":
                            if true {
                                var v29 int = int(v19)
                                var v30 int = v29
                                _ = v29
                                _ = v30
                            }
                        case "foo bar":
                            switch v1 {
                            case "\n\t":
                                v31 := GenS3{F0: 1, F1: v2}
                                _ = v31
                            }
                        case "\n\t":
                            v5 -= 185
                        }
                        v32 := GenS4{F0: (v3 - (v19 % ((v3) | 1)))}
                        var v33 uint8 = 23
                        _ = v32
                        _ = v33
                    }
                }
            }
            _ = v19
        } else {
            v5 += v5
            switch ("hello" + v1) {
            case "a":
                switch v1 {
                case "a":
                    v34 := GenS5{F0: (fi18 - fi18)}
                    _ = v34
                }
            default:
                v4 = append(v4, v1)
            }
            if (false && true) {
                var v35 uint = 236
                switch (fi18 * (fi18 ^ (fi18 | 0))) {
                case 1:
                    for fi36 := 0; fi36 < 5; fi36++ {
                        switch "foo bar" {
                        case "":
                            v4 = append(v4, "hello")
                            var v37 uint32 = 389
                            v38 := GenS1{F0: (v5 << ((327) & 63))}
                            _ = v37
                            _ = v38
                        case "a":
                            v3 += 0
                        case "foo bar":
                            if false {
                                v39 := GenS6{F0: v2, F1: 855, F2: 0}
                                _ = v39
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = v1
                                }
                            }
                        default:
                            v1 = v1
                        }
                        _ = fi36
                    }
                case 128:
                    if false {
                        v35 -= (v35 >> ((v35) & 63))
                        if len(v4) > 0 {
                            v4[0] = (v1 + v1)
                        }
                        v35 -= v35
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v2 += (v2 * -926)
                    }
                }
                _ = v35
            } else {
                v40 := GenS7{F0: (921 < 0), F1: false}
                var v41 int32 = int32(v2)
                _ = v40
                _ = v41
            }
        }
        switch "foo bar" {
        case "a":
            switch (("" + "a") + (v1 + "\n\t")) {
            case "":
                var v42 float32 = 3.14159
                _ = v42
            case "\n\t":
                switch -9223372036854775808 {
                case -1:
                    for fi43 := 0; fi43 < 4; fi43++ {
                        fmt.Println(fi18, v2)
                        if len(v4) > 0 {
                            v4[0] = (v1 + "\n\t")
                        }
                        _ = fi43
                    }
                case 4294967295:
                    var v44 int16 = (v2 - v2)
                    _ = v44
                case -857:
                    if true {
                        for fi45 := 0; fi45 < 11; fi45++ {
                            switch (-(fi18)) {
                            case 42:
                                if true {
                                    v2 = 32767
                                    var v46 string = (v1 + "foo bar")
                                    _ = v46
                                }
                            default:
                                var v47 uint8 = 128
                                _ = v47
                            }
                            if false {
                                fmt.Println(v1)
                            }
                            var v48 string = "foo bar"
                            _ = fi45
                            _ = v48
                        }
                    }
                default:
                    var v49 int8 = -119
                    _ = v49
                }
            case "foo bar":
                if false {
                    var v50 uint = 883
                    var v51 string = v1
                    _ = v50
                    _ = v51
                }
            }
        }
        if (-2.25 != 3.14159) {
            if true {
                v3 = v3
            } else {
                v52 := GenS1{F0: v5}
                _ = v52
            }
            v53 := GenS8{F0: v5, F1: 42}
            _ = len(v4)
            _ = v53
        } else {
            v5 = v5
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    v3 -= 684
                } else {
                    v54 := GenS4{F0: (^(v3))}
                    _ = v54
                }
                v55 := GenS9{F0: 773, F1: 2147483647}
                v2 += v2
                _ = v55
            }
        }
        _ = fi18
    }
    var v56 float64 = (0.1 + 1e10)
    _ = len(v4)
    for fi57 := 0; fi57 < 10; fi57++ {
        var v58 int8 = -1
        v4 = append(v4, ((((v1 + v1) + v1) + ("foo bar" + ("hello" + ("hello" + (v1 + "a"))))) + v1))
        var v59 float32 = 0.0
        _ = fi57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 13; fi60++ {
        if (!true) {
            fmt.Println(fi60, fi60, v5)
            for fi61 := 0; fi61 < 12; fi61++ {
                var v62 int = 579
                switch "foo bar" {
                case "":
                    switch "hello" {
                    case "a":
                        var v63 float64 = ((3.14159 * 1e10) + v56)
                        _ = v63
                    }
                case "hello":
                    var v64 int64 = 153
                    _ = v64
                default:
                    fmt.Println(fi61)
                }
                v62 += v62
                _ = fi61
                _ = v62
            }
            v3 += (v3 * v3)
        }
        _ = fi60
    }
    v56 = v56
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, v1)
    v65 := GenS7{F0: (false && (v3 >= 180)), F1: ((42 > v5) || false)}
    for fi66 := 0; fi66 < 6; fi66++ {
        var v67 uint8 = 2
        _ = fi66
        _ = v67
    }
    v68 := GenS10{F0: 127}
    v69 := GenS11{F0: 1, F1: (v3 - v3)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v65
    _ = v68
    _ = v69
}
