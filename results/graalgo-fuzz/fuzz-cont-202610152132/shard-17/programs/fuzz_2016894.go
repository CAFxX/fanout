package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int64; F1 int8; F2 int8 }
type GenS3 struct { F0 int64; F1 int64; F2 uint }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int16; F2 uint64 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 float32 = 0.1
    var v5 int8 = 0
    var v6 uint8 = (v1 << ((4294967295) & 63))
    v1 = v1
    if (3.14159 <= 0.1) {
        if (-115 >= -9223372036854775808) {
            v3[""] = true
            v7 := GenS1{F0: (("" + ("hello" + "foo bar")) + ("foo bar" + (("foo bar" + ("a" + "hello")) + "hello"))), F1: ("" + "")}
            _ = v7
        }
    }
    v2[32768] = 4294967295
    var v8 string = "foo bar"
    v9 := []int{4294967295}
    if (42 == 514) {
        var v10 uint8 = 154
        if true {
            v10 += v10
            v6 = v1
        } else {
            var v11 int64 = 184
            _ = len(v9)
            var v12 int16 = 42
            _ = v11
            _ = v12
        }
        if false {
            v13 := GenS2{F0: 0, F1: (v5 ^ (v5 | (-(v5)))), F2: -15}
            var v14 int8 = v5
            _ = v13
            _ = v14
        }
        _ = v10
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            if (960 <= 255) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    v6 = v6
                    var v17 int = (fi15 * (fi15 % ((fi16) | 1)))
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 19; fi18++ {
                    switch 42 {
                    case 42:
                        v4 -= v4
                    default:
                        switch fi18 {
                        case 32768:
                            var v19 uint64 = 0
                            _ = v19
                        case -9223372036854775808:
                            switch (fi18 >> ((255) & 63)) {
                            case 42:
                                switch v8 {
                                case "\n\t":
                                    var v20 int32 = 42
                                    _ = v20
                                case "foo bar":
                                    v21 := GenS2{F0: 9223372036854775807, F1: 127, F2: v5}
                                    _ = v21
                                case "hello":
                                    v22 := GenS2{F0: 768, F1: v5, F2: (v5 * (v5 * (v5 >> ((255) & 63))))}
                                    _ = v22
                                default:
                                    fmt.Println(fi15, v4)
                                }
                            case -2147483648:
                                v4 = (1e10 - 1e10)
                            case 9223372036854775807:
                                v8 = (((v8 + ((v8 + v8) + (v8 + "a"))) + (("a" + (v8 + "a")) + v8)) + "a")
                            }
                            v6 = v1
                            if true {
                                v23 := GenS1{F0: "foo bar", F1: v8}
                                v8 = v8
                                _ = v23
                            } else {
                                var v24 int64 = 124
                                _ = v24
                            }
                            fmt.Println(v6, v6)
                            v4 -= 3.14159
                        case -773:
                            switch fi15 {
                            case 4294967295:
                                v9 = append(v9, -2147483648)
                            case 340:
                                v1 = 255
                            }
                        default:
                            if true {
                                v25 := GenS2{F0: 1, F1: 1, F2: (-(v5))}
                                _ = len(v9)
                                _ = v25
                            } else {
                                switch int(v6) {
                                case 2147483647:
                                    if true {
                                        var v26 int8 = v5
                                        v27 := GenS1{F0: "", F1: (v8 + v8)}
                                        _ = v26
                                        _ = v27
                                    }
                                case -979:
                                    v28 := GenS2{F0: 287, F1: -120, F2: 31}
                                    _ = v28
                                case 4294967295:
                                    v4 += ((3.14159 / -2.25) * 0.1)
                                default:
                                    v1 = v6
                                }
                                v6 = (v1 >> ((42) & 63))
                            }
                        }
                    }
                    if (false || true) {
                        if true {
                            var v29 uint16 = 317
                            switch fi15 {
                            case 83:
                                switch (fi18 + (^(fi15))) {
                                case 4294967295:
                                    var v30 int16 = 1
                                    _ = v30
                                case 42:
                                    fmt.Println(v29)
                                case -1:
                                    v31 := GenS3{F0: -1, F1: 1, F2: 4294967295}
                                    _ = v31
                                }
                            }
                            _ = v29
                        } else {
                            v32 := v2[fi18]
                            var v33 int64 = -1
                            _ = v32
                            _ = v33
                        }
                        if ((v6 * 255) != (v6 | v6)) {
                            v4 -= (v4 - v4)
                        } else {
                            switch ("" + ((v8 + v8) + v8)) {
                            case "\n\t":
                                if (1 >= 1) {
                                    var v34 int8 = (v5 / ((v5) | 1))
                                    var v35 int16 = 1000
                                    var v36 uint = 394
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                } else {
                                    v6 = v6
                                    v37 := GenS4{F0: 32767, F1: 127, F2: 0}
                                    _ = v37
                                }
                            case "hello":
                                v6 = v6
                            default:
                                switch -9223372036854775808 {
                                case 127:
                                    v6 = (v6 + 0)
                                case -9223372036854775808:
                                    var v38 bool = (2 > 0)
                                    _ = v38
                                default:
                                    switch 127 {
                                    case 32767:
                                        fmt.Println(v5, v4, v1)
                                    }
                                }
                            }
                            v5 = (v5 / ((v5) | 1))
                            delete(v2, (fi18 + fi15))
                        }
                        var v39 uint16 = 813
                        _ = v39
                    } else {
                        _ = len(v2)
                    }
                    _ = fi18
                }
            } else {
                v8 = ((((("foo bar" + ("hello" + v8)) + v8) + v8) + "hello") + (("a" + (v8 + "foo bar")) + "foo bar"))
                for fi40 := 0; fi40 < 6; fi40++ {
                    v41 := v2[(-(fi40))]
                    _ = fi40
                    _ = v41
                }
            }
            _ = fi15
        }
    }
    v8 = v8
    v9 = append(v9, 906)
    v42 := GenS2{F0: 1, F1: (v5 + (v5 % ((1) | 1))), F2: (v5 - v5)}
    _ = len(v9)
    if (false || true) {
        v43 := GenS2{F0: 1, F1: v5, F2: (v5 - v5)}
        _ = len(v2)
        _ = v43
    } else {
        v44 := GenS5{F0: "hello", F1: 42}
        _ = len(v9)
        if true {
            for fi45 := 0; fi45 < 5; fi45++ {
                fmt.Println(v8, v5, fi45)
                v46 := v3[v8]
                _ = fi45
                _ = v46
            }
            v8 = "hello"
            fmt.Println(v1, v1, v8)
        }
        _ = v44
    }
    if false {
        if false {
            if false {
                var v47 uint16 = 65535
                if true {
                    if false {
                        if true {
                            var v48 float64 = 0.0
                            if false {
                                v2[1] = 0
                                v49 := GenS3{F0: 9223372036854775807, F1: -9223372036854775808, F2: 18446744073709551615}
                                _ = v49
                            } else {
                                if false {
                                    v6 = v6
                                } else {
                                    fmt.Println(v1, v8, v47)
                                    v3[""] = true
                                    for fi50 := 0; fi50 < 1; fi50++ {
                                        delete(v3, v8)
                                        v2[int(v5)] = 2147483648
                                        _ = fi50
                                    }
                                }
                                var v51 int8 = -8
                                v52 := v2[-918]
                                _ = v51
                                _ = v52
                            }
                            _ = v48
                        } else {
                            v1 += v6
                            _ = len(v9)
                        }
                        v1 = v1
                        switch v8 {
                        case "":
                            if false {
                                switch 32767 {
                                case 2147483647:
                                    v6 = 1
                                }
                                var v53 int16 = 28
                                v54 := GenS6{F0: (2 == 42), F1: 82}
                                _ = v53
                                _ = v54
                            }
                        default:
                            var v55 uint8 = 255
                            _ = v55
                        }
                    } else {
                        v8 = "a"
                    }
                } else {
                    var v56 int8 = -1
                    _ = v56
                }
                var v57 int8 = (v5 ^ -128)
                _ = v47
                _ = v57
            }
            v4 -= v4
        }
        v1 -= (v1 + v1)
        if false {
            if (1 < 0) {
                fmt.Println(v6, v4)
                var v58 int64 = 0
                if (1e10 < float64((v4 - v4))) {
                    switch "hello" {
                    case "a":
                        var v59 int16 = 0
                        _ = v59
                    case "hello":
                        _ = len(v9)
                    default:
                        v1 += 0
                    }
                } else {
                    fmt.Println(v1)
                    fmt.Println(v8)
                }
                _ = v58
            }
            v60 := GenS7{F0: (v6 << ((255) & 63)), F1: -708, F2: 1}
            if len(v9) > 0 {
                v9[0] = 775
            }
            _ = v60
        } else {
            switch -2147483648 {
            case -947:
                for fi61 := 0; fi61 < 18; fi61++ {
                    var v62 float64 = 0.0
                    _ = fi61
                    _ = v62
                }
            case 2147483647:
                var v63 int32 = 42
                _ = v63
            case 0:
                v64 := GenS6{F0: true, F1: 270}
                _ = v64
            }
            v4 -= v4
            for fi65 := 0; fi65 < 19; fi65++ {
                if false {
                    v1 = (v1 - (^(v1)))
                    var v66 uint32 = 545
                    fmt.Println(v66, v5)
                    _ = v66
                }
                v67 := GenS3{F0: -1, F1: -345, F2: 0}
                _ = fi65
                _ = v67
            }
        }
    } else {
        v4 += 0.0
        v4 -= (0.0 / 1.5)
    }
    var v68 int64 = -9223372036854775808
    switch "hello" {
    case "\n\t":
        for fi69 := 0; fi69 < 10; fi69++ {
            fmt.Println(v4)
            v1 += v1
            switch v8 {
            case "foo bar":
                fmt.Println(fi69, v5, v5)
            case "":
                for fi70 := 0; fi70 < 11; fi70++ {
                    var v71 int16 = 42
                    _ = fi70
                    _ = v71
                }
            default:
                v6 = (v1 | (^(v1)))
            }
            _ = fi69
        }
    case "hello":
        switch "hello" {
        case "hello":
            var v72 int = -2147483648
            _ = v72
        case "":
            var v73 uint32 = 1
            _ = v73
        }
    default:
        var v74 uint8 = (v6 * (v1 % (((v6 + v1)) | 1)))
        _ = v74
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v68:", v68)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v68
    _ = v8
    _ = v9
}
