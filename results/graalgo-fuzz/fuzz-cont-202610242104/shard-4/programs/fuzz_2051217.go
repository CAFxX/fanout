package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }

func main() {
    var v1 int8 = -1
    var v2 string = ("hello" + "a")
    var v3 int = 1
    var v4 float32 = 1.5
    var v5 bool = false
    var v6 uint = 0
    var v7 float64 = -2.25
    var v8 uint16 = 42
    if v5 {
        fmt.Println(v3, v7, v1)
    } else {
        fmt.Println(v4)
        var v9 uint = 804
        if (v8 >= v8) {
            v10 := GenS1{F0: 0, F1: -766, F2: 2147483648}
            var v11 uint32 = 4294967295
            _ = v10
            _ = v11
        }
        _ = v9
    }
    v12 := GenS2{F0: "\n\t"}
    v13 := []int{v3, 1, (v3 << ((v6) & 63)), (v3 ^ (v3 ^ (v3 * (v3 & v3))))}
    if len(v13) > 0 {
        v13[0] = 1
    }
    switch v2 {
    case "a":
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 int16 = 0
            _ = fi14
            _ = v15
        }
    case "hello":
        v16 := GenS1{F0: int64(v3), F1: 1, F2: 1}
        _ = v16
    case "":
        var v17 uint = 0
        _ = v17
    default:
        switch v3 {
        case 42:
            switch v2 {
            case "hello":
                for fi18 := 0; fi18 < 10; fi18++ {
                    if len(v13) > 0 {
                        v13[0] = (fi18 & fi18)
                    }
                    if v5 {
                        if len(v13) > 0 {
                            v13[0] = (v3 ^ fi18)
                        }
                        var v19 uint8 = 42
                        if (18446744073709551615 < 685) {
                            var v20 uint64 = 474
                            v1 -= v1
                            _ = v20
                        }
                        _ = v19
                    }
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if len(v13) > 0 {
                            v13[0] = fi21
                        }
                        if ("\n\t" < ("foo bar" + "a")) {
                            if len(v13) > 0 {
                                v13[0] = fi21
                            }
                            switch v3 {
                            case -2147483648:
                                fmt.Println(v6, fi18, fi21)
                            }
                        } else {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                        if false {
                            v5 = (false && v5)
                        } else {
                            if false {
                                fmt.Println(v5, v1)
                                v7 += (v7 * (v7 * (float64(-2.25) * v7)))
                                switch v2 {
                                case "":
                                    v22 := GenS2{F0: (v2 + v2)}
                                    _ = v22
                                case "foo bar":
                                    if v5 {
                                        fmt.Println(v1, v2, fi21)
                                    } else {
                                        if v5 {
                                            v23 := GenS2{F0: v2}
                                            _ = v23
                                        } else {
                                            var v24 int32 = -1
                                            v25 := GenS3{F0: -808, F1: 1}
                                            switch v2 {
                                            case "hello":
                                                fmt.Println(v8)
                                            case "\n\t":
                                                v6 -= (v6 & v6)
                                            case "a":
                                                if (!(!(!v5))) {
                                                    var v26 int = 42
                                                    v26 = -365
                                                    var v27 uint = v6
                                                    switch ((v2 + v2) + v2) {
                                                    case "":
                                                        v28 := GenS4{F0: true, F1: v24}
                                                        _ = v28
                                                    case "foo bar":
                                                        var v29 int = (^(v3))
                                                        _ = v29
                                                    case "\n\t":
                                                        v30 := GenS4{F0: true, F1: v24}
                                                        _ = v30
                                                    }
                                                    _ = v26
                                                    _ = v27
                                                } else {
                                                    var v31 float64 = v7
                                                    var v32 int = 127
                                                    v33 := GenS5{F0: ((65535 >= 65535) && ("a" != "foo bar")), F1: 255}
                                                    switch v2 {
                                                    case "foo bar":
                                                        if false {
                                                            if (4294967295 < v6) {
                                                                switch "\n\t" {
                                                                case "\n\t":
                                                                    fmt.Println(v2, v32, v2)
                                                                case "hello":
                                                                    v6 = (v6 + (^(v6)))
                                                                }
                                                                v34 := GenS6{F0: 42}
                                                                v4 = v4
                                                                _ = v34
                                                            }
                                                            var v35 uint16 = v8
                                                            v8 = v8
                                                            var v36 float64 = ((v7 * v31) * v31)
                                                            var v37 uint32 = 552
                                                            _ = v35
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        var v38 uint16 = (v8 % ((uint16(v6)) | 1))
                                                        v13 = append(v13, (-(fi18)))
                                                        _ = v38
                                                    case "a":
                                                        switch v2 {
                                                        case "\n\t":
                                                            v31 = (3.14159 + float64(float32((v31 - float64(v4)))))
                                                        default:
                                                            v7 += v7
                                                        }
                                                    case "\n\t":
                                                        var v39 int16 = 219
                                                        _ = v39
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                }
                                            }
                                            _ = v24
                                            _ = v25
                                        }
                                        switch 606 {
                                        case -2147483648:
                                            var v40 float32 = v4
                                            _ = v40
                                        case 32768:
                                            fmt.Println(v6, v7)
                                        case 127:
                                            var v41 int16 = -961
                                            _ = v41
                                        default:
                                            if false {
                                                v8 = (v8 ^ v8)
                                            }
                                        }
                                        var v42 int8 = (v1 * (v1 >> ((v6) & 63)))
                                        _ = v42
                                    }
                                case "hello":
                                    switch (v3 / (((fi21 | v3)) | 1)) {
                                    case 4294967295:
                                        v8 = (v8 + v8)
                                    case 32768:
                                        switch v2 {
                                        case "\n\t":
                                            v43 := GenS5{F0: (v6 != uint(v8)), F1: (v8 & v8)}
                                            _ = v43
                                        case "foo bar":
                                            switch "hello" {
                                            case "foo bar":
                                                var v44 uint = (^(v6))
                                                _ = v44
                                            case "hello":
                                                var v45 int8 = 1
                                                _ = v45
                                            case "\n\t":
                                                v7 = 0.0
                                            default:
                                                if len(v13) > 0 {
                                                    v13[0] = -298
                                                }
                                            }
                                        }
                                    case 454:
                                        if v5 {
                                            v3 += 2147483647
                                            if len(v13) > 0 {
                                                v13[0] = (fi21 / ((42) | 1))
                                            }
                                            var v46 float32 = 0.1
                                            var v47 uint = (v6 ^ v6)
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                }
                            }
                            v48 := GenS6{F0: 4294967295}
                            if v5 {
                                fmt.Println(v6, v6)
                                var v49 uint32 = 267
                                _ = v49
                            } else {
                                switch fi21 {
                                case 2147483647:
                                    var v50 uint = (v6 / ((v6) | 1))
                                    _ = v50
                                }
                            }
                            if len(v13) > 0 {
                                v13[0] = -2147483648
                            }
                            _ = v48
                        }
                        _ = fi21
                    }
                    _ = fi18
                }
            case "\n\t":
                var v51 int16 = -827
                _ = v51
            case "foo bar":
                var v52 uint32 = 334
                _ = v52
            default:
                switch 660 {
                case 42:
                    for fi53 := 0; fi53 < 5; fi53++ {
                        for fi54 := 0; fi54 < 8; fi54++ {
                            var v55 uint8 = 127
                            v1 += 127
                            _ = fi54
                            _ = v55
                        }
                        _ = fi53
                    }
                case 845:
                    for fi56 := 0; fi56 < 17; fi56++ {
                        var v57 int8 = (v1 * -28)
                        _ = fi56
                        _ = v57
                    }
                case 127:
                    v58 := GenS6{F0: 4294967295}
                    _ = v58
                }
            }
        case -963:
            var v59 uint32 = 4294967295
            _ = v59
        case -9223372036854775808:
            for fi60 := 0; fi60 < 13; fi60++ {
                v3 = (-(fi60))
                if (((^(v8)) < (v8 ^ v8)) || (!(v6 == (v6 / (((v6 >> ((v6) & 63))) | 1))))) {
                    fmt.Println(v2, v1, v2)
                    if (!v5) {
                        if v5 {
                            for fi61 := 0; fi61 < 19; fi61++ {
                                v4 = 0.0
                                var v62 int = int(v8)
                                v8 = (v8 / ((1) | 1))
                                fmt.Println(v1, v5, fi61)
                                _ = fi61
                                _ = v62
                            }
                        } else {
                            v63 := GenS1{F0: 0, F1: 0, F2: 2147483647}
                            _ = v63
                        }
                    }
                    switch fi60 {
                    case 752:
                        fmt.Println(v4)
                    default:
                        var v64 int64 = -819
                        _ = v64
                    }
                } else {
                    for fi65 := 0; fi65 < 18; fi65++ {
                        var v66 string = ((("\n\t" + ((v2 + ("\n\t" + (((v2 + "foo bar") + "foo bar") + (v2 + v2)))) + v2)) + ((v2 + "\n\t") + (v2 + ((v2 + ("foo bar" + v2)) + v2)))) + v2)
                        v8 = v8
                        _ = fi65
                        _ = v66
                    }
                    if v5 {
                        var v67 uint32 = 216
                        switch "foo bar" {
                        case "":
                            var v68 uint16 = 267
                            _ = v68
                        case "hello":
                            if len(v13) > 0 {
                                v13[0] = fi60
                            }
                        }
                        _ = v67
                    } else {
                        v4 += 0.1
                        v69 := GenS7{F0: v8}
                        v70 := GenS6{F0: 429}
                        _ = v69
                        _ = v70
                    }
                }
                _ = fi60
            }
        default:
            v6 -= 65535
        }
    }
    var v71 int32 = 0
    v1 += (v1 * 100)
    for fi72 := 0; fi72 < 16; fi72++ {
        var v73 int32 = v71
        v71 = (^(v71))
        var v74 int64 = -9223372036854775808
        _ = fi72
        _ = v73
        _ = v74
    }
    if v5 {
        var v75 uint64 = uint64(v8)
        _ = v75
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v71:", v71)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v71
    _ = v8
}
