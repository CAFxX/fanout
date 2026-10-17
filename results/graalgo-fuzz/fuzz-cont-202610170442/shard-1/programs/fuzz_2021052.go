package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int8; F1 int64; F2 int64 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint64; F1 bool; F2 int }
type GenS7 struct { F0 string; F1 string; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 128
    v3 := make(map[int]int)
    v4 := []uint32{859, 535, 2}
    v5 := []uint32{2147483647, 42}
    switch "" {
    case "hello":
        fmt.Println(v2)
    case "a":
        for fi6 := 0; fi6 < 18; fi6++ {
            for fi7 := 0; fi7 < 2; fi7++ {
                v2 += 127
                fmt.Println(fi6, fi7, fi7)
                switch -1 {
                case 2147483647:
                    v2 = v2
                case 630:
                    _ = len(v3)
                }
                _ = len(v5)
                _ = fi7
            }
            for fi8 := 0; fi8 < 18; fi8++ {
                fmt.Println(fi6)
                _ = fi8
            }
            _ = fi6
        }
    }
    fmt.Println(v2)
    if false {
        v2 = 255
    } else {
        v9 := GenS1{F0: 4294967295}
        fmt.Println(v2)
        _ = v9
    }
    var v10 float64 = ((0.0 - 3.14159) - 1e10)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 string = "foo bar"
        var v13 bool = (v10 <= float64((1.5 + 3.14159)))
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v10, v2)
    fmt.Println(v2, v10)
    v14 := GenS1{F0: 156}
    if (-380 < 4294967295) {
        v15 := GenS1{F0: 524}
        v5 = append(v5, 762)
        if (2 >= 255) {
            for fi16 := 0; fi16 < 10; fi16++ {
                var v17 float64 = 0.0
                _ = fi16
                _ = v17
            }
        } else {
            v18 := GenS2{F0: 100}
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 float64 = v10
                fmt.Println(v20)
                for fi21 := 0; fi21 < 19; fi21++ {
                    v2 += v2
                    v22 := v1[2147483648]
                    v10 += v20
                    if (42 >= -32768) {
                        if v22 {
                            switch (fi19 - (fi21 / (((fi19 | 127)) | 1))) {
                            case -356:
                                v23 := GenS1{F0: uint64(v2)}
                                _ = v23
                            case -434:
                                v2 -= (v2 - (^(v2)))
                            case -204:
                                var v24 int64 = 42
                                _ = v24
                            default:
                                v25 := GenS3{F0: (2147483647 >= 0), F1: -1}
                                _ = v25
                            }
                            var v26 bool = true
                            v2 += v2
                            _ = v26
                        } else {
                            if false {
                                v10 = v10
                                fmt.Println(v2)
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v27 := GenS3{F0: v22, F1: 9223372036854775807}
                        var v28 string = "a"
                        _ = len(v5)
                        switch int(v2) {
                        case 127:
                            fmt.Println(v20)
                        case 32767:
                            if false {
                                var v29 bool = (true && false)
                                _ = v29
                            }
                        default:
                            fmt.Println(fi21)
                        }
                        fmt.Println(v10)
                        _ = v27
                        _ = v28
                    }
                    _ = fi21
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            if ((3.14159 <= ((0.0 / (0.0 + (float32(((v10 / v10) * (v10 + 3.14159))) + 0.0))) * (float32((-2.25 + float64(0.0))) - 1e10))) || true) {
                var v30 int64 = 0
                fmt.Println(v30)
                _ = v30
            }
            _ = v18
        }
        _ = v15
    } else {
        v31 := GenS4{F0: 1, F1: -24, F2: -395}
        _ = v31
    }
    if (683 >= 42) {
        var v32 int8 = -1
        if false {
            for fi33 := 0; fi33 < 18; fi33++ {
                fmt.Println(fi33, v10)
                if true {
                    v10 += v10
                }
                _ = fi33
            }
        } else {
            v32 = (v32 % (((v32 % ((47) | 1))) | 1))
        }
        v5 = append(v5, 2147483648)
        _ = v32
    } else {
        switch 4294967295 {
        case 2147483647:
            v34 := v1[809]
            _ = v34
        case 9223372036854775807:
            var v35 uint64 = uint64(v2)
            _ = v35
        case -2147483648:
            if false {
                if true {
                    if false {
                        fmt.Println(v2, v2)
                        var v36 float32 = (float32(v10) - (((((1e10 / (1.5 * 1.5)) + 1e10) / float32(v10)) / 1.5) - 0.0))
                        _ = v36
                    } else {
                        var v37 int8 = 0
                        v2 = v2
                        v38 := GenS5{F0: 535, F1: 5}
                        _ = v37
                        _ = v38
                    }
                }
                if false {
                    var v39 uint = 65535
                    v2 += v2
                    _ = v39
                } else {
                    if true {
                        v5 = append(v5, 0)
                        var v40 uint16 = 42
                        var v41 int64 = -9223372036854775808
                        _ = v40
                        _ = v41
                    }
                    v42 := GenS6{F0: 938, F1: true, F2: -286}
                    _ = v42
                }
            } else {
                var v43 int8 = 1
                var v44 bool = ((false && (18446744073709551615 == 65535)) || ((^(v2)) > v2))
                for fi45 := 0; fi45 < 10; fi45++ {
                    var v46 string = ("foo bar" + "a")
                    v1[301] = false
                    var v47 int16 = 361
                    _ = fi45
                    _ = v46
                    _ = v47
                }
                _ = v43
                _ = v44
            }
        }
        var v48 uint32 = 1
        var v49 uint32 = v48
        _ = v48
        _ = v49
    }
    fmt.Println(v10, v10)
    switch -1 {
    case 42:
        var v50 int16 = 42
        _ = v50
    case 9223372036854775807:
        v51 := GenS7{F0: (("a" + "foo bar") + "a"), F1: ("a" + ("\n\t" + "foo bar")), F2: (!(-460 <= 0))}
        _ = v51
    case -9223372036854775808:
        if true {
            v2 = v2
            delete(v3, 32768)
        } else {
            switch ("a" + "\n\t") {
            case "\n\t":
                if false {
                    var v52 int16 = 1
                    _ = v52
                }
            }
            fmt.Println(v10, v2)
        }
    default:
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    v10 -= v10
    v5 = append(v5, 0)
    v53 := GenS5{F0: 42, F1: 0}
    v10 -= v10
    v10 -= (3.14159 + v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
