package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool; F1 int64; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 int; F2 int }
type GenS9 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS10 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0(p0 uint8) int16 {
    var v6 int = 1
    v7 := GenS1{F0: (v6 ^ (v6 ^ -1))}
    _ = p0
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]uint32)
    var v3 bool = (670 <= 0)
    var v4 uint64 = 1
    var v5 string = "\n\t"
    v1 = v5
    v8 := []int{1, 708}
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 18; fi10++ {
            switch v5 {
            case "\n\t":
                _ = len(v2)
            case "a":
                v2[42] = 871
            case "":
                switch 1 {
                case 2147483647:
                    var v11 uint32 = 2147483647
                    _ = v11
                case -259:
                    switch fi9 {
                    case 9223372036854775807:
                        if (!((!v3) && true)) {
                            _ = len(v2)
                        }
                    }
                case -2147483648:
                    var v12 int64 = 9223372036854775807
                    _ = v12
                default:
                    var v13 int64 = 494
                    _ = v13
                }
            default:
                switch (fi10 - fi9) {
                case -350:
                    if true {
                        v14 := GenS2{F0: 0, F1: -86}
                        var v15 uint64 = 899
                        _ = len(v2)
                        _ = v14
                        _ = v15
                    }
                default:
                    _ = len(v2)
                }
            }
            if true {
                v2[42] = 469
                var v16 int32 = 0
                if (0 != 42) {
                    var v17 float32 = 1e10
                    v3 = v3
                    switch fi10 {
                    case 213:
                        var v18 uint64 = 42
                        var v19 float64 = -2.25
                        _ = v18
                        _ = v19
                    case 32768:
                        if (!v3) {
                            switch ("" + ("hello" + "")) {
                            case "\n\t":
                                if ((v4 / ((v4) | 1)) <= (v4 | 690)) {
                                    var v20 float64 = 1e10
                                    _ = v20
                                }
                            }
                        }
                    case 42:
                        if (!(v3 && false)) {
                            v8 = append(v8, (-(fi9)))
                            fmt.Println(v5, v1, v1)
                            switch ((v5 + "foo bar") + v1) {
                            case "":
                                fmt.Println(v4, v5)
                            default:
                                var v21 int8 = 127
                                _ = v21
                            }
                        }
                    }
                    _ = v17
                }
                _ = v16
            } else {
                v22 := GenS3{F0: 1000}
                _ = v22
            }
            if v3 {
                v23 := GenS1{F0: -2147483648}
                _ = v23
            }
            _ = fi10
        }
        _ = fi9
    }
    switch v5 {
    case "hello":
        v24 := GenS4{F0: -128, F1: -498, F2: v4}
        _ = v24
    case "foo bar":
        v25 := GenS5{F0: 814}
        _ = v25
    default:
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := GenS6{F0: v3, F1: 1, F2: v4}
            if (!v3) {
                if (793 <= 683) {
                    switch 916 {
                    case 42:
                        delete(v2, 2)
                    default:
                        if false {
                            for fi28 := 0; fi28 < 3; fi28++ {
                                var v29 int64 = 66
                                switch fi28 {
                                case 9223372036854775807:
                                    switch (v1 + (v1 + ("" + v1))) {
                                    case "foo bar":
                                        var v30 int = 128
                                        _ = v30
                                    case "a":
                                        if (false || v3) {
                                            var v31 int64 = (v29 & v29)
                                            v2[13] = 2147483648
                                            var v32 int = -9223372036854775808
                                            _ = v31
                                            _ = v32
                                        }
                                    case "\n\t":
                                        v3 = (v3 && (v3 || (!v3)))
                                    default:
                                        v3 = true
                                    }
                                }
                                var v33 bool = (1 < 128)
                                v34 := GenS7{F0: v5}
                                _ = fi28
                                _ = v29
                                _ = v33
                                _ = v34
                            }
                            v5 = ((v1 + ("foo bar" + v5)) + v5)
                        }
                    }
                    _ = len(v2)
                } else {
                    if (127 >= 100) {
                        v5 = (v5 + "hello")
                        switch (("" + v1) + "\n\t") {
                        case "hello":
                            switch "foo bar" {
                            case "hello":
                                v35 := GenS8{F0: 42, F1: fi26, F2: (fi26 | (fi26 ^ (fi26 % (((fi26 + fi26)) | 1))))}
                                _ = v35
                            case "a":
                                switch fi26 {
                                case -2147483648:
                                    var v36 float64 = 0.0
                                    _ = v36
                                case 128:
                                    v4 = 0
                                case 32768:
                                    _ = len(v8)
                                }
                            case "\n\t":
                                switch fi26 {
                                case 32768:
                                    for fi37 := 0; fi37 < 19; fi37++ {
                                        var v38 float32 = 1e10
                                        v39 := GenS6{F0: (((0.1 * -2.25) - (-2.25 - 1e10)) < (3.14159 - float64(0.0))), F1: 1, F2: v4}
                                        v40 := v2[13]
                                        _ = fi37
                                        _ = v38
                                        _ = v39
                                        _ = v40
                                    }
                                case 32767:
                                    if (!false) {
                                        v1 = (v1 + v5)
                                        v5 = (v1 + v1)
                                        if (42 <= 2) {
                                            var v41 float32 = 3.14159
                                            _ = v41
                                        } else {
                                            var v42 int8 = -122
                                            _ = v42
                                        }
                                    } else {
                                        delete(v2, 4294967295)
                                        switch (fi26 * -2147483648) {
                                        case 50:
                                            v43 := GenS9{F0: -1, F1: 2147483648, F2: (^(v4))}
                                            _ = v43
                                        }
                                        var v44 float32 = (-2.25 / (float32(float64(0.0)) - 3.14159))
                                        _ = v44
                                    }
                                }
                            }
                        case "\n\t":
                            v45 := GenS7{F0: (v1 + "a")}
                            _ = v45
                        case "a":
                            v3 = true
                        }
                    }
                    v46 := v2[1]
                    _ = v46
                }
                switch "" {
                case "hello":
                    if (!((100 < -1) && v3)) {
                        var v47 int8 = 32
                        _ = v47
                    }
                }
            }
            var v48 uint64 = v4
            _ = fi26
            _ = v27
            _ = v48
        }
    }
    fmt.Println(v5, v3)
    v1 = (v5 + (v1 + ""))
    if true {
        for fi49 := 0; fi49 < 14; fi49++ {
            if true {
                for fi50 := 0; fi50 < 17; fi50++ {
                    fmt.Println(v3, fi49)
                    _ = fi50
                }
                v51 := GenS6{F0: false, F1: -9223372036854775808, F2: v4}
                _ = v51
            } else {
                v5 = ("foo bar" + "foo bar")
            }
            v1 = "\n\t"
            switch v1 {
            case "\n\t":
                v5 = ((v1 + "hello") + v1)
            case "a":
                if len(v8) > 0 {
                    v8[0] = 127
                }
            default:
                if false {
                    v1 = v5
                }
            }
            _ = fi49
        }
    }
    var v52 int8 = 42
    if (2 >= 173) {
        v53 := GenS10{F0: (v52 & (v52 % ((0) | 1))), F1: -552, F2: 1}
        var v54 float64 = float64(3.14159)
        v55 := GenS5{F0: 65535}
        _ = v53
        _ = v54
        _ = v55
    } else {
        var v56 uint = 65535
        _ = v56
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        var v58 uint64 = (v4 * 0)
        v1 = (v1 + ("" + v1))
        v59 := GenS5{F0: 18446744073709551615}
        _ = fi57
        _ = v58
        _ = v59
    }
    v60 := GenS5{F0: 4294967295}
    var v61 int16 = -836
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v60
    _ = v61
    _ = v8
}
