package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 string; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint64; F2 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 bool; F2 string }

func main() {
    v1 := []int{281, 1, 4294967295, 1}
    var v2 uint64 = 1
    var v3 float64 = (1e10 + (1.5 * (0.1 / float64((float32((((0.0 * (-2.25 - 0.1)) - 1e10) + (1.5 * 1.5))) * 3.14159)))))
    var v4 string = ("\n\t" + "foo bar")
    var v5 uint64 = 1
    v6 := GenS1{F0: 978, F1: (v5 | v2)}
    v7 := GenS2{F0: -382, F1: 2}
    var v8 uint8 = 255
    v5 += v5
    for fi9 := 0; fi9 < 4; fi9++ {
        if false {
            fmt.Println(v4, fi9)
        }
        if false {
            if false {
                var v10 int16 = -32768
                if false {
                    var v11 uint64 = (v2 - v2)
                    fmt.Println(v4)
                    _ = v11
                }
                _ = v10
            }
            v1 = append(v1, 42)
        } else {
            switch v4 {
            case "\n\t":
                v8 = (v8 % (((v8 << ((42) & 63))) | 1))
            }
            for fi12 := 0; fi12 < 3; fi12++ {
                fmt.Println(v5, v2, v8)
                v1 = append(v1, 4294967295)
                if len(v1) > 0 {
                    v1[0] = 0
                }
                if (1 > -32768) {
                    if (0 < 93) {
                        if len(v1) > 0 {
                            v1[0] = fi9
                        }
                        v4 = ((v4 + ((v4 + "foo bar") + v4)) + (v4 + (v4 + (v4 + (v4 + ("foo bar" + v4))))))
                        var v13 int = fi12
                        _ = v13
                    }
                    if false {
                        switch (fi9 >> ((255) & 63)) {
                        case 1:
                            v3 += 1e10
                        case -985:
                            switch fi9 {
                            case 9223372036854775807:
                                fmt.Println(fi9)
                            default:
                                v2 = (v5 | (v5 | (v2 % ((v2) | 1))))
                            }
                        default:
                            var v14 int = (fi12 << ((444) & 63))
                            _ = v14
                        }
                        v1 = append(v1, 128)
                        v5 = 2
                    }
                    switch fi12 {
                    case -1:
                        switch "hello" {
                        case "hello":
                            if false {
                                var v15 uint32 = 395
                                _ = len(v1)
                                var v16 int8 = 112
                                _ = v15
                                _ = v16
                            }
                        default:
                            switch ("" + v4) {
                            case "\n\t":
                                v4 = v4
                            case "hello":
                                switch fi9 {
                                case 584:
                                    if (true && (((2 < 255) && (255 >= 932)) && (true || false))) {
                                        if (912 != -2147483648) {
                                            _ = len(v1)
                                            if ((v2 != v5) || (v4 >= "a")) {
                                                var v17 string = v4
                                                var v18 string = v17
                                                _ = v17
                                                _ = v18
                                            } else {
                                                if (-1 >= 42) {
                                                    if false {
                                                        v4 = v4
                                                    } else {
                                                        if true {
                                                            if (1 > 327) {
                                                                v8 += (v8 / ((211) | 1))
                                                            } else {
                                                                fmt.Println(v2, fi12, v5)
                                                                fmt.Println(fi9)
                                                                _ = len(v1)
                                                            }
                                                        }
                                                        v19 := GenS1{F0: 0, F1: v2}
                                                        if (!(false || true)) {
                                                            switch v4 {
                                                            case "foo bar":
                                                                if len(v1) > 0 {
                                                                    v1[0] = 9223372036854775807
                                                                }
                                                            case "\n\t":
                                                                switch 970 {
                                                                case -2147483648:
                                                                    fmt.Println(v4, v2, v8)
                                                                case -957:
                                                                    var v20 uint8 = 127
                                                                    _ = v20
                                                                default:
                                                                    var v21 uint8 = (v8 & (v8 ^ v8))
                                                                    _ = v21
                                                                }
                                                                v22 := GenS3{F0: (v4 + (v4 + v4)), F1: 65535, F2: 2}
                                                                if len(v1) > 0 {
                                                                    v1[0] = 1
                                                                }
                                                                switch -2147483648 {
                                                                case 9223372036854775807:
                                                                    if false {
                                                                        if (!(1 <= 255)) {
                                                                            v2 = v2
                                                                            if ((42 != fi12) || (3.14159 >= 1.5)) {
                                                                                v23 := GenS4{F0: (!(false && ((!(!(0 < 9223372036854775807))) || (0.1 >= -2.25)))), F1: (v5 | (v5 >> ((173) & 63))), F2: -498}
                                                                                _ = v23
                                                                            }
                                                                        } else {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = fi9
                                                                            }
                                                                            if (false && false) {
                                                                                v24 := GenS3{F0: (v4 + "hello"), F1: 561, F2: 18446744073709551615}
                                                                                v4 = "hello"
                                                                                fmt.Println(v5, fi12)
                                                                                _ = v24
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 32767
                                                                            }
                                                                        }
                                                                        v5 = (v5 / ((v5) | 1))
                                                                    }
                                                                case -217:
                                                                    var v25 int8 = -87
                                                                    _ = v25
                                                                case 128:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = 930
                                                                    }
                                                                }
                                                                v4 = ""
                                                                v26 := GenS4{F0: true, F1: 0, F2: (fi9 & (fi9 + (fi12 % (((fi9 - int(v8))) | 1))))}
                                                                v2 += v5
                                                                _ = v22
                                                                _ = v26
                                                            }
                                                            if (730 == fi12) {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            } else {
                                                                if (-699 != 543) {
                                                                    if true {
                                                                        v27 := GenS3{F0: (v4 + "a"), F1: 65535, F2: (v5 - 42)}
                                                                        _ = v27
                                                                    } else {
                                                                        v4 = (v4 + v4)
                                                                        var v28 float32 = (0.0 + 1.5)
                                                                        var v29 uint16 = 39
                                                                        _ = v28
                                                                        _ = v29
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v2, v3)
                                                            if (!true) {
                                                                _ = len(v1)
                                                                if false {
                                                                    fmt.Println(v5)
                                                                } else {
                                                                    if false {
                                                                        v3 += 0.1
                                                                        var v30 uint8 = v8
                                                                        _ = v30
                                                                    }
                                                                }
                                                            } else {
                                                                if ("" != v4) {
                                                                    var v31 uint = 2
                                                                    v32 := GenS5{F0: fi9}
                                                                    v31 = (v31 | 42)
                                                                    var v33 uint8 = v8
                                                                    _ = v31
                                                                    _ = v32
                                                                    _ = v33
                                                                }
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = (fi12 * 127)
                                                            }
                                                        }
                                                        _ = v19
                                                    }
                                                    v34 := GenS6{F0: 18446744073709551615, F1: v8}
                                                    v5 = 500
                                                    _ = v34
                                                } else {
                                                    var v35 int = (fi12 * 4294967295)
                                                    if false {
                                                        switch v4 {
                                                        case "foo bar":
                                                            if (1000 != -1) {
                                                                v36 := GenS1{F0: 184, F1: (v2 >> ((18446744073709551615) & 63))}
                                                                var v37 uint16 = 65535
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                v4 = "\n\t"
                                                                v2 = (v5 & v2)
                                                            }
                                                        }
                                                    } else {
                                                        var v38 uint32 = 2147483647
                                                        v39 := GenS1{F0: -2147483648, F1: uint64(v38)}
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                    v5 -= (v5 | 4294967295)
                                                    _ = v35
                                                }
                                                _ = len(v1)
                                                switch fi9 {
                                                case -2147483648:
                                                    if len(v1) > 0 {
                                                        v1[0] = -616
                                                    }
                                                }
                                            }
                                            fmt.Println(fi9, v3, fi9)
                                            v40 := GenS6{F0: uint(v8), F1: 2}
                                            _ = v40
                                        }
                                        v41 := GenS4{F0: false, F1: (v5 & 57), F2: fi12}
                                        if true {
                                            switch "a" {
                                            case "hello":
                                                if false {
                                                    switch fi12 {
                                                    case -1:
                                                        switch (fi12 << ((4294967295) & 63)) {
                                                        case 42:
                                                            switch fi12 {
                                                            case 1:
                                                                var v42 int64 = 42
                                                                _ = v42
                                                            case 411:
                                                                var v43 int8 = -128
                                                                _ = v43
                                                            }
                                                        case -2147483648:
                                                            v2 = v2
                                                        }
                                                    case 42:
                                                        fmt.Println(v3)
                                                    default:
                                                        if true {
                                                            if ((false || (2 == 13)) && (-1 > -1)) {
                                                                var v44 int = fi12
                                                                v45 := GenS4{F0: false, F1: (v5 ^ (v5 - v2)), F2: fi9}
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        } else {
                                                            v4 = v4
                                                        }
                                                    }
                                                    switch "a" {
                                                    case "":
                                                        v46 := GenS2{F0: 1000, F1: (^(v8))}
                                                        _ = v46
                                                    }
                                                } else {
                                                    v8 += (v8 << ((255) & 63))
                                                    v1 = append(v1, (fi12 << ((0) & 63)))
                                                }
                                            case "a":
                                                fmt.Println(v2, v3)
                                            default:
                                                fmt.Println(fi12, fi12, v3)
                                            }
                                            v2 += (^(v2))
                                            var v47 uint8 = (^(v8))
                                            _ = v47
                                        }
                                        if true {
                                            var v48 int = (fi9 - (fi12 / (((fi12 / ((fi12) | 1))) | 1)))
                                            var v49 int32 = 0
                                            fmt.Println(v3)
                                            _ = v48
                                            _ = v49
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v41
                                    } else {
                                        fmt.Println(v2, v2, v2)
                                    }
                                case 32768:
                                    if false {
                                        var v50 uint32 = 42
                                        v4 = (((v4 + "") + "a") + "a")
                                        v2 += (v2 % ((42) | 1))
                                        v5 = v5
                                        _ = v50
                                    } else {
                                        fmt.Println(v3, v2, v3)
                                        v3 += v3
                                        if true {
                                            fmt.Println(v4)
                                            var v51 string = v4
                                            switch fi9 {
                                            case 949:
                                                switch -343 {
                                                case 996:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case 765:
                                                    _ = len(v1)
                                                case 0:
                                                    switch v51 {
                                                    case "":
                                                        v5 -= v2
                                                    case "\n\t":
                                                        var v52 bool = (518 != 0)
                                                        _ = v52
                                                    }
                                                default:
                                                    fmt.Println(v3)
                                                }
                                            case 2147483647:
                                                v4 = ("" + (v51 + (("\n\t" + v51) + "a")))
                                            case 127:
                                                v8 = (v8 - 42)
                                            default:
                                                _ = len(v1)
                                            }
                                            var v53 uint64 = 450
                                            v54 := GenS7{F0: (v4 + v51)}
                                            _ = v51
                                            _ = v53
                                            _ = v54
                                        } else {
                                            v55 := GenS8{F0: (v2 + 0), F1: false, F2: (v4 + (v4 + v4))}
                                            _ = v55
                                        }
                                    }
                                case -164:
                                    var v56 int = fi9
                                    _ = v56
                                default:
                                    var v57 int32 = -680
                                    _ = v57
                                }
                            default:
                                v2 -= v2
                            }
                        }
                    case 128:
                        v8 -= 127
                    case 0:
                        v3 = float64((0.0 * ((1.5 - 1e10) - 1.5)))
                    default:
                        var v58 string = (((((((v4 + (((v4 + v4) + v4) + ("" + v4))) + ("" + "")) + ("a" + ("\n\t" + v4))) + ("" + v4)) + (("" + "\n\t") + ("foo bar" + v4))) + (v4 + (((((v4 + v4) + ((v4 + "foo bar") + "a")) + v4) + v4) + v4))) + ("hello" + v4))
                        _ = v58
                    }
                }
                var v59 uint = 29
                _ = fi12
                _ = v59
            }
            for fi60 := 0; fi60 < 15; fi60++ {
                fmt.Println(v8, v5)
                v61 := GenS8{F0: 469, F1: false, F2: ("" + v4)}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v62 := GenS6{F0: 1, F1: (v8 / (((v8 >> ((65535) & 63))) | 1))}
                v4 = v4
                _ = fi60
                _ = v61
                _ = v62
            }
        }
        _ = fi9
    }
    v2 = (v2 + v2)
    var v63 int8 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v7
    _ = v8
}
