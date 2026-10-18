package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int16; F1 int16; F2 string }
type GenS4 struct { F0 bool; F1 uint8; F2 uint }
type GenS5 struct { F0 uint32; F1 uint32; F2 int }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 uint; F1 uint16; F2 uint8 }
type GenS8 struct { F0 string }

func main() {
    var v1 bool = false
    var v2 bool = v1
    var v3 int16 = 1
    var v4 uint64 = 719
    v5 := []string{("a" + "hello"), "\n\t", "hello"}
    var v6 int64 = 1
    if (!(true || (-1 <= 23))) {
        var v7 uint64 = (v4 ^ (v4 >> ((42) & 63)))
        _ = v7
    } else {
        v6 = (v6 & -9223372036854775808)
        if len(v5) > 0 {
            v5[0] = ""
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v2 = (v3 > v3)
    fmt.Println(v1, v6)
    var v8 uint32 = 447
    var v9 string = "\n\t"
    if (v9 != (((v9 + "") + ("foo bar" + (v9 + v9))) + "")) {
        if v1 {
            v10 := GenS1{F0: 42}
            v2 = v1
            _ = v10
        }
        v8 = 42
        if v2 {
            v2 = (!v1)
            v2 = (1.5 <= -2.25)
        }
    } else {
        v5 = append(v5, (v9 + "foo bar"))
        v2 = false
    }
    if v2 {
        var v11 uint64 = v4
        if v2 {
            v3 -= 32767
            var v12 bool = v2
            _ = len(v5)
            _ = v12
        } else {
            if len(v5) > 0 {
                v5[0] = v9
            }
            if (237 < 65535) {
                v3 -= (v3 - (v3 % ((v3) | 1)))
                v13 := GenS1{F0: 4294967295}
                v14 := GenS1{F0: 4294967295}
                _ = v13
                _ = v14
            }
        }
        _ = v11
    }
    switch v9 {
    case "foo bar":
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(v9)
            var v16 uint = 42
            _ = fi15
            _ = v16
        }
    case "\n\t":
        if true {
            fmt.Println(v2, v2)
        } else {
            if v2 {
                if (1 > -1) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v8 = (v8 | v8)
                    v1 = (v1 && (238 < 469))
                } else {
                    if v2 {
                        v3 += v3
                        v17 := GenS1{F0: 18446744073709551615}
                        _ = v17
                    }
                    v18 := GenS1{F0: 0}
                    _ = v18
                }
            } else {
                v19 := GenS1{F0: uint(v4)}
                _ = v19
            }
            switch v9 {
            case "hello":
                if (255 < uint(v4)) {
                    if (!v2) {
                        fmt.Println(v2)
                        v2 = (1000 < (v3 >> ((418) & 63)))
                        switch (("" + "a") + "") {
                        case "":
                            if v1 {
                                v2 = (!(-32768 == -922))
                                fmt.Println(v1, v4)
                            } else {
                                fmt.Println(v8, v8, v2)
                                var v20 int8 = 1
                                if ((313 < 42) && (!v1)) {
                                    v21 := GenS2{F0: v3, F1: 19}
                                    switch 467 {
                                    case 9223372036854775807:
                                        fmt.Println(v2)
                                    case 32768:
                                        for fi22 := 0; fi22 < 1; fi22++ {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            v9 = v9
                                            var v23 int = (fi22 & (fi22 - (fi22 % ((fi22) | 1))))
                                            _ = fi22
                                            _ = v23
                                        }
                                    case -1:
                                        var v24 uint64 = (^(v4))
                                        _ = v24
                                    default:
                                        var v25 int64 = v6
                                        _ = v25
                                    }
                                    var v26 int16 = 1000
                                    _ = v21
                                    _ = v26
                                } else {
                                    var v27 bool = v1
                                    for fi28 := 0; fi28 < 15; fi28++ {
                                        fmt.Println(v2, v27, v9)
                                        v1 = (!(0 <= 397))
                                        if len(v5) > 0 {
                                            v5[0] = v9
                                        }
                                        _ = fi28
                                    }
                                    var v29 int16 = -666
                                    _ = v27
                                    _ = v29
                                }
                                _ = v20
                            }
                        case "\n\t":
                            v30 := GenS3{F0: (v3 | 0), F1: 32767, F2: ("\n\t" + v9)}
                            _ = v30
                        case "a":
                            var v31 int16 = 1
                            _ = v31
                        }
                    } else {
                        for fi32 := 0; fi32 < 17; fi32++ {
                            var v33 int16 = -32768
                            _ = fi32
                            _ = v33
                        }
                        v3 -= v3
                        if (v2 || v1) {
                            _ = len(v5)
                            var v34 uint16 = 696
                            _ = v34
                        } else {
                            var v35 uint16 = 32767
                            for fi36 := 0; fi36 < 10; fi36++ {
                                v1 = false
                                if len(v5) > 0 {
                                    v5[0] = v9
                                }
                                if v1 {
                                    switch "" {
                                    case "a":
                                        switch (v9 + (v9 + (v9 + v9))) {
                                        case "\n\t":
                                            v35 = (v35 - v35)
                                        case "foo bar":
                                            var v37 int16 = v3
                                            _ = v37
                                        case "":
                                            v6 = (-(v6))
                                        }
                                    case "foo bar":
                                        v5 = append(v5, v9)
                                    case "\n\t":
                                        var v38 int8 = int8(fi36)
                                        _ = v38
                                    default:
                                        for fi39 := 0; fi39 < 8; fi39++ {
                                            fmt.Println(v3, v8, v2)
                                            _ = fi39
                                        }
                                    }
                                    if true {
                                        if (false || v2) {
                                            if v2 {
                                                if (false || (255 >= 646)) {
                                                    var v40 uint16 = v35
                                                    fmt.Println(v35, v6, v3)
                                                    if v2 {
                                                        v35 = (v40 / ((uint16(v8)) | 1))
                                                        v2 = (fi36 >= fi36)
                                                        _ = len(v5)
                                                    }
                                                    _ = v40
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            }
                                            v4 = 709
                                        }
                                        if v2 {
                                            v41 := GenS1{F0: 1}
                                            if v1 {
                                                var v42 int64 = v6
                                                _ = len(v5)
                                                _ = v42
                                            } else {
                                                v43 := GenS4{F0: (v35 > v35), F1: 42, F2: 42}
                                                _ = v43
                                            }
                                            _ = v41
                                        } else {
                                            for fi44 := 0; fi44 < 10; fi44++ {
                                                v45 := GenS3{F0: (v3 + (v3 % ((v3) | 1))), F1: (v3 + (^(v3))), F2: ""}
                                                v35 -= 0
                                                _ = fi44
                                                _ = v45
                                            }
                                            var v46 int8 = -90
                                            switch fi36 {
                                            case 42:
                                                var v47 bool = ((v4 * 18446744073709551615) >= 42)
                                                _ = v47
                                            case 32768:
                                                var v48 int32 = 1
                                                _ = v48
                                            case 0:
                                                if true {
                                                    for fi49 := 0; fi49 < 10; fi49++ {
                                                        var v50 uint8 = 255
                                                        fmt.Println(v6, v4)
                                                        _ = fi49
                                                        _ = v50
                                                    }
                                                }
                                            }
                                            _ = v46
                                        }
                                    }
                                }
                                _ = fi36
                            }
                            v51 := GenS5{F0: v8, F1: v8, F2: 32767}
                            _ = v35
                            _ = v51
                        }
                    }
                    v8 = v8
                    switch -9223372036854775808 {
                    case 0:
                        if true {
                            var v52 int8 = -1
                            var v53 uint64 = (v4 % ((v4) | 1))
                            if (v8 >= v8) {
                                if false {
                                    if (v2 || (255 >= 42)) {
                                        v54 := GenS6{F0: v52, F1: 32767}
                                        fmt.Println(v8, v9)
                                        v5 = append(v5, v9)
                                        _ = v54
                                    } else {
                                        var v55 uint32 = (v8 + 2147483647)
                                        v56 := GenS7{F0: 0, F1: 65535, F2: 27}
                                        var v57 int16 = (v3 << ((458) & 63))
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                    }
                                    switch v9 {
                                    case "hello":
                                        v4 = (^(v4))
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                                v4 = (v4 + v53)
                            } else {
                                v58 := GenS1{F0: 715}
                                v6 = v6
                                _ = v58
                            }
                            _ = v52
                            _ = v53
                        } else {
                            v8 = v8
                            switch -533 {
                            case -1:
                                var v59 float32 = 0.1
                                _ = v59
                            default:
                                if false {
                                    v4 -= v4
                                } else {
                                    v60 := GenS5{F0: 4294967295, F1: 2147483648, F2: 4294967295}
                                    fmt.Println(v4, v8)
                                    _ = v60
                                }
                            }
                        }
                    case -35:
                        for fi61 := 0; fi61 < 5; fi61++ {
                            v3 -= (v3 + v3)
                            _ = fi61
                        }
                    case -9223372036854775808:
                        v5 = append(v5, (("" + v9) + ((("\n\t" + v9) + v9) + "a")))
                    default:
                        v2 = ((!(4294967295 >= uint64(v8))) && (true || (!v2)))
                    }
                }
            case "":
                v62 := GenS8{F0: v9}
                _ = v62
            }
        }
    }
    _ = len(v5)
    if v1 {
        v1 = (-1 > v6)
        for fi63 := 0; fi63 < 12; fi63++ {
            var v64 int = fi63
            _ = fi63
            _ = v64
        }
        var v65 uint32 = (v8 + v8)
        _ = v65
    } else {
        if (2 <= 1) {
            if v2 {
                for fi66 := 0; fi66 < 13; fi66++ {
                    switch v9 {
                    case "":
                        v9 = "foo bar"
                    case "\n\t":
                        v67 := GenS1{F0: 65535}
                        _ = v67
                    }
                    switch v9 {
                    case "\n\t":
                        v68 := GenS8{F0: ""}
                        _ = v68
                    case "hello":
                        var v69 uint16 = 255
                        _ = v69
                    }
                    fmt.Println(v6, v9)
                    _ = fi66
                }
                v4 -= (v4 << ((4294967295) & 63))
                if (-478 == (v3 - 0)) {
                    v3 -= (v3 | v3)
                    if (590 < (v4 % ((v4) | 1))) {
                        v9 = "\n\t"
                        var v70 float64 = -2.25
                        _ = v70
                    }
                } else {
                    v2 = v1
                    var v71 int = -9223372036854775808
                    _ = v71
                }
            } else {
                _ = len(v5)
                v2 = (18446744073709551615 != 876)
            }
            fmt.Println(v3)
        }
        if (250 >= 167) {
            v9 = v9
        } else {
            var v72 bool = v1
            var v73 int32 = 1
            if v72 {
                var v74 uint = 4294967295
                switch (v9 + "") {
                case "hello":
                    v8 -= v8
                }
                _ = v74
            } else {
                var v75 uint32 = (v8 % ((0) | 1))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = v75
            }
            _ = v72
            _ = v73
        }
        for fi76 := 0; fi76 < 17; fi76++ {
            var v77 int8 = 1
            _ = fi76
            _ = v77
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
