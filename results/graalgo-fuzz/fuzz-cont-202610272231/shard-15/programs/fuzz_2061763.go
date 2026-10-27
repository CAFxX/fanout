package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 uint8; F1 uint32; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 int; F1 string; F2 uint }
type GenS10 struct { F0 string; F1 int64; F2 uint8 }

func main() {
    v1 := []string{"foo bar"}
    var v2 int32 = 42
    var v3 int32 = -1
    v2 += (-(v3))
    var v4 uint16 = 32767
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (!true) {
            fmt.Println(v3, v4)
            v5 := GenS1{F0: 0, F1: 115, F2: -9223372036854775808}
            _ = v5
        }
    } else {
        var v6 uint32 = 787
        _ = v6
    }
    var v7 uint8 = 234
    switch "" {
    case "\n\t":
        if true {
            _ = len(v1)
            v8 := GenS1{F0: 919, F1: (v4 << ((65535) & 63)), F2: -1}
            _ = v8
        }
    }
    v3 -= (v3 / ((v2) | 1))
    v2 = (v2 << ((0) & 63))
    if false {
        var v9 int64 = 1
        var v10 uint = 1
        if true {
            fmt.Println(v2, v10)
            v4 += (v4 % (((v4 >> ((v10) & 63))) | 1))
        } else {
            v11 := GenS1{F0: 140, F1: v4, F2: -2147483648}
            var v12 uint = v10
            _ = v11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    v4 = (v4 << ((4294967295) & 63))
    if (!false) {
        var v13 float32 = 1e10
        v13 -= v13
        if len(v1) > 0 {
            v1[0] = ""
        }
        _ = v13
    }
    v2 += (v3 - v2)
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            v15 := GenS1{F0: 656, F1: (v4 >> ((44) & 63)), F2: fi14}
            v3 += v3
            switch "a" {
            case "":
                if len(v1) > 0 {
                    v1[0] = ("a" + ("hello" + "\n\t"))
                }
            case "hello":
                var v16 uint32 = 2147483648
                _ = v16
            case "foo bar":
                if true {
                    var v17 uint16 = v4
                    for fi18 := 0; fi18 < 11; fi18++ {
                        v4 -= 1
                        _ = fi18
                    }
                    _ = v17
                }
            default:
                if (false && false) {
                    fmt.Println(fi14, fi14)
                    switch ("hello" + "") {
                    case "foo bar":
                        for fi19 := 0; fi19 < 4; fi19++ {
                            switch (^(fi19)) {
                            case -51:
                                var v20 float64 = float64(1.5)
                                _ = v20
                            case 42:
                                if ((3.14159 / (1.5 + 1.5)) > 1.5) {
                                    v1 = append(v1, "\n\t")
                                    v21 := GenS2{F0: 42, F1: 76}
                                    v22 := GenS3{F0: (v7 | (^(v7)))}
                                    _ = v21
                                    _ = v22
                                } else {
                                    v1 = append(v1, "a")
                                    var v23 bool = true
                                    _ = v23
                                }
                            case 2147483647:
                                var v24 uint64 = 42
                                _ = v24
                            }
                            v4 -= 981
                            _ = fi19
                        }
                    case "":
                        v25 := GenS4{F0: true, F1: 42}
                        _ = v25
                    default:
                        var v26 int64 = 447
                        _ = v26
                    }
                    v4 -= v4
                }
            }
            _ = v15
        } else {
            if (fi14 < fi14) {
                var v27 int = 42
                fmt.Println(fi14, v3)
                var v28 uint64 = 4294967295
                _ = v27
                _ = v28
            } else {
                if false {
                    switch 9223372036854775807 {
                    case 127:
                        v29 := GenS1{F0: 2, F1: v4, F2: (fi14 & (fi14 % ((fi14) | 1)))}
                        _ = v29
                    case 529:
                        fmt.Println(fi14)
                    case -299:
                        fmt.Println(v7, fi14, v7)
                    default:
                        if true {
                            v4 -= (^(v4))
                        }
                    }
                    if (42 > -9223372036854775808) {
                        if ((1000 > 42) && (2147483647 <= 297)) {
                            if (false || false) {
                                fmt.Println(v3)
                            } else {
                                var v30 int16 = 32767
                                var v31 int64 = int64(v2)
                                _ = v30
                                _ = v31
                            }
                            for fi32 := 0; fi32 < 19; fi32++ {
                                v33 := GenS5{F0: 42, F1: int32(v7)}
                                var v34 string = ("hello" + "foo bar")
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                    } else {
                        switch (-(fi14)) {
                        case 32767:
                            fmt.Println(v4, v4, v4)
                        default:
                            for fi35 := 0; fi35 < 4; fi35++ {
                                var v36 uint8 = 226
                                v4 += (v4 / (((^(v4))) | 1))
                                switch (fi14 ^ 1) {
                                case 0:
                                    v37 := GenS4{F0: false, F1: 1}
                                    _ = v37
                                }
                                _ = fi35
                                _ = v36
                            }
                        }
                        switch "" {
                        case "a":
                            v3 = -511
                        case "hello":
                            switch ("" + "a") {
                            case "\n\t":
                                for fi38 := 0; fi38 < 8; fi38++ {
                                    v39 := GenS1{F0: 2, F1: 2, F2: fi38}
                                    var v40 float32 = float32(3.14159)
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            case "foo bar":
                                v7 += (v7 ^ (v7 + (v7 / ((v7) | 1))))
                            default:
                                for fi41 := 0; fi41 < 9; fi41++ {
                                    var v42 uint = 1
                                    var v43 bool = ((2 <= 1) && false)
                                    _ = fi41
                                    _ = v42
                                    _ = v43
                                }
                            }
                        case "\n\t":
                            if true {
                                if (2 < 42) {
                                    switch 32768 {
                                    case 4294967295:
                                        if false {
                                            if true {
                                                _ = len(v1)
                                            } else {
                                                v3 = 2147483647
                                                fmt.Println(v4, v2)
                                            }
                                        } else {
                                            fmt.Println(fi14, v7)
                                            for fi44 := 0; fi44 < 16; fi44++ {
                                                if true {
                                                    v45 := GenS6{F0: (v7 ^ 255), F1: 42, F2: 42}
                                                    _ = v45
                                                } else {
                                                    fmt.Println(fi14)
                                                }
                                                _ = len(v1)
                                                _ = fi44
                                            }
                                            switch "a" {
                                            case "hello":
                                                var v46 uint64 = 4294967295
                                                _ = v46
                                            case "foo bar":
                                                _ = len(v1)
                                            case "\n\t":
                                                var v47 uint8 = v7
                                                _ = v47
                                            }
                                        }
                                    case -1:
                                        var v48 int8 = -128
                                        _ = v48
                                    default:
                                        fmt.Println(v2, v4)
                                    }
                                } else {
                                    switch ("" + "a") {
                                    case "a":
                                        switch fi14 {
                                        case 752:
                                            var v49 uint32 = 1
                                            _ = v49
                                        }
                                    case "\n\t":
                                        v50 := GenS4{F0: ((-(v3)) >= v3), F1: -9223372036854775808}
                                        _ = v50
                                    }
                                    switch ("foo bar" + "\n\t") {
                                    case "foo bar":
                                        switch (^(fi14)) {
                                        case 128:
                                            v51 := GenS7{F0: false}
                                            _ = v51
                                        case 127:
                                            fmt.Println(v2, v4, v4)
                                        case 42:
                                            if true {
                                                if false {
                                                    if true {
                                                        v52 := GenS8{F0: fi14, F1: 127}
                                                        v4 -= (v4 + (v4 ^ (v4 / (((v4 * (v4 + (v4 / ((v4) | 1))))) | 1))))
                                                        _ = v52
                                                    }
                                                }
                                            } else {
                                                v7 += (v7 >> ((4294967295) & 63))
                                                for fi53 := 0; fi53 < 8; fi53++ {
                                                    v7 += 1
                                                    _ = len(v1)
                                                    _ = fi53
                                                }
                                            }
                                        default:
                                            for fi54 := 0; fi54 < 1; fi54++ {
                                                v55 := GenS9{F0: -2147483648, F1: "", F2: 42}
                                                _ = fi54
                                                _ = v55
                                            }
                                        }
                                    default:
                                        switch 32768 {
                                        case -2147483648:
                                            fmt.Println(fi14, v3)
                                        case 32767:
                                            fmt.Println(v4, v3, v3)
                                        case 128:
                                            _ = len(v1)
                                        }
                                    }
                                    switch fi14 {
                                    case 1:
                                        fmt.Println(v7, v2)
                                    }
                                }
                            } else {
                                switch fi14 {
                                case -30:
                                    v2 = (^(v3))
                                case -854:
                                    v3 = (v2 << ((18446744073709551615) & 63))
                                }
                                v56 := GenS10{F0: "hello", F1: 42, F2: (v7 >> ((502) & 63))}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v56
                            }
                        default:
                            v2 -= v2
                        }
                        v7 -= 127
                    }
                }
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
    }
    _ = len(v1)
    v7 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
