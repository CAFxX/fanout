package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int16 }

func main() {
    var v1 int32 = 0
    var v2 uint8 = 255
    var v3 int8 = -128
    var v4 bool = (!((4294967295 > 65535) || (881 == 2)))
    if (0.0 > 0.0) {
        switch -9223372036854775808 {
        case 621:
            var v5 int32 = v1
            _ = v5
        case 32768:
            switch ("a" + "hello") {
            case "a":
                var v6 int16 = 32767
                _ = v6
            case "\n\t":
                switch 42 {
                case 128:
                    v7 := GenS1{F0: 42}
                    _ = v7
                case -9223372036854775808:
                    if true {
                        for fi8 := 0; fi8 < 8; fi8++ {
                            if v4 {
                                v1 = (v1 >> ((447) & 63))
                            } else {
                                if (-128 > (v3 % ((v3) | 1))) {
                                    v2 = 42
                                    v3 = 42
                                }
                                for fi9 := 0; fi9 < 2; fi9++ {
                                    v10 := []int{}
                                    v3 = (v3 % (((-(v3))) | 1))
                                    _ = fi9
                                    _ = v10
                                }
                            }
                            if v4 {
                                switch "" {
                                case "foo bar":
                                    if (fi8 < fi8) {
                                        var v11 uint64 = 2
                                        _ = v11
                                    }
                                case "":
                                    v2 -= 255
                                }
                            } else {
                                for fi12 := 0; fi12 < 7; fi12++ {
                                    fmt.Println(fi8, v1)
                                    v13 := []int{290, (fi8 | 1), 2147483647}
                                    v14 := GenS1{F0: 42}
                                    v3 = v3
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                }
                            }
                            v15 := GenS1{F0: 0}
                            _ = fi8
                            _ = v15
                        }
                        v3 = (v3 >> ((65535) & 63))
                        v4 = v4
                    }
                }
            case "hello":
                v2 = 46
            }
        case 127:
            switch "" {
            case "\n\t":
                v2 += v2
            case "a":
                var v16 float64 = -2.25
                _ = v16
            case "":
                v1 = (v1 / ((v1) | 1))
            }
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            for fi18 := 0; fi18 < 18; fi18++ {
                v19 := GenS1{F0: 18446744073709551615}
                v20 := []int{fi18, fi17, 127}
                var v21 string = ((("foo bar" + "a") + "foo bar") + ("" + (("a" + ((("\n\t" + "a") + (("\n\t" + "foo bar") + ("hello" + "a"))) + "hello")) + ("hello" + "a"))))
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            if v4 {
                fmt.Println(v3)
                if ("" != ((("a" + "") + "") + "a")) {
                    v3 -= (v3 + 127)
                }
                v22 := []int{}
                _ = v22
            } else {
                switch "hello" {
                case "\n\t":
                    if (77 > 600) {
                        var v23 int32 = v1
                        fmt.Println(v1, v2)
                        _ = v23
                    }
                case "":
                    var v24 uint16 = 18
                    _ = v24
                case "a":
                    var v25 uint16 = 437
                    _ = v25
                }
            }
            var v26 int = fi17
            _ = fi17
            _ = v26
        }
        fmt.Println(v3, v1)
    }
    v27 := GenS1{F0: 18446744073709551615}
    var v28 float32 = 0.1
    v29 := GenS1{F0: 4294967295}
    for fi30 := 0; fi30 < 17; fi30++ {
        for fi31 := 0; fi31 < 14; fi31++ {
            var v32 bool = v4
            v33 := GenS1{F0: 18446744073709551615}
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    var v34 float64 = 3.14159
    for fi35 := 0; fi35 < 10; fi35++ {
        if (!(false && false)) {
            if v4 {
                var v36 int32 = (v1 % (((v1 + v1)) | 1))
                fmt.Println(fi35, v34)
                _ = v36
            } else {
                switch "hello" {
                case "foo bar":
                    fmt.Println(v2, v28, v34)
                case "a":
                    for fi37 := 0; fi37 < 8; fi37++ {
                        switch "hello" {
                        case "foo bar":
                            v3 = v3
                        case "":
                            if v4 {
                                var v38 uint32 = 1
                                _ = v38
                            } else {
                                v39 := GenS1{F0: 0}
                                fmt.Println(fi35)
                                var v40 uint32 = 491
                                _ = v39
                                _ = v40
                            }
                        case "\n\t":
                            v28 = float32(-2.25)
                        }
                        v34 = (3.14159 - v34)
                        _ = fi37
                    }
                }
            }
            var v41 uint = 4294967295
            if (0 == 697) {
                if (v4 && v4) {
                    v42 := GenS2{F0: -323, F1: 130}
                    fmt.Println(v2, v28, v28)
                    v43 := GenS1{F0: 0}
                    _ = v42
                    _ = v43
                }
                v44 := GenS3{F0: 1, F1: 4294967295, F2: 1000}
                _ = v44
            }
            _ = v41
        }
        switch "\n\t" {
        case "a":
            v3 += (v3 >> ((318) & 63))
        case "hello":
            for fi45 := 0; fi45 < 1; fi45++ {
                fmt.Println(v28, fi45, fi45)
                v46 := GenS1{F0: 1}
                v28 -= v28
                _ = fi45
                _ = v46
            }
        case "":
            switch "foo bar" {
            case "a":
                if v4 {
                    var v47 string = "\n\t"
                    _ = v47
                } else {
                    v48 := GenS4{F0: v3, F1: v3}
                    v49 := GenS4{F0: -27, F1: v3}
                    _ = v48
                    _ = v49
                }
            case "\n\t":
                if true {
                    v3 = v3
                    if (("a" + "\n\t") < (("" + "hello") + ("" + ("" + "hello")))) {
                        v50 := []int{(fi35 % (((fi35 - (fi35 - int(v2)))) | 1)), 4294967295, -9223372036854775808}
                        if false {
                            fmt.Println(v34)
                            var v51 int64 = -716
                            _ = v51
                        } else {
                            for fi52 := 0; fi52 < 7; fi52++ {
                                fmt.Println(v1, v4, v28)
                                _ = fi52
                            }
                        }
                        v28 -= v28
                        _ = v50
                    } else {
                        v53 := GenS2{F0: -1, F1: 42}
                        fmt.Println(v28, v4, v34)
                        if true {
                            v54 := []int{fi35, fi35, (fi35 + (fi35 | (fi35 % ((fi35) | 1)))), (fi35 % ((fi35) | 1))}
                            var v55 int32 = -891
                            var v56 uint32 = 42
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        _ = v53
                    }
                    var v57 int16 = 32767
                    _ = v57
                }
            case "":
                v58 := []int{(fi35 & fi35), fi35}
                _ = v58
            }
        }
        var v59 uint64 = 987
        _ = fi35
        _ = v59
    }
    fmt.Println(v3, v2, v4)
    v60 := GenS5{F0: (!(65535 > 1)), F1: ("" + "a")}
    for fi61 := 0; fi61 < 9; fi61++ {
        var v62 string = ""
        for fi63 := 0; fi63 < 9; fi63++ {
            v64 := []int{}
            if v4 {
                switch 4294967295 {
                case 32768:
                    fmt.Println(v1)
                }
                var v65 float32 = (0.0 / v28)
                v66 := GenS6{F0: 32767}
                v28 = v28
                _ = v65
                _ = v66
            } else {
                switch v62 {
                case "foo bar":
                    v4 = (260 != 5)
                }
                var v67 bool = (!(((864 == 893) && ((v1 & v1) > v1)) && (true && (0 != 244))))
                switch 32768 {
                case -9223372036854775808:
                    if true {
                        v4 = false
                    } else {
                        if len(v64) >= 2 {
                            v64 = v64[1:len(v64)]
                        }
                        v64 = append(v64, -82)
                    }
                case 67:
                    var v68 int8 = 127
                    _ = v68
                }
                _ = v67
            }
            _ = fi63
            _ = v64
        }
        _ = fi61
        _ = v62
    }
    v34 -= float64((v28 - v28))
    var v69 int16 = 32767
    fmt.Println(v69, v4)
    v70 := GenS3{F0: 110, F1: 2147483648, F2: (v69 * (^(v69)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v60
    _ = v69
    _ = v70
}
