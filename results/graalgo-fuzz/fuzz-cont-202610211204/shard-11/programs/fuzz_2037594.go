package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint16; F2 int32 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int32; F2 int8 }

func main() {
    var v1 int16 = -998
    var v2 int32 = 1
    var v3 int64 = 9223372036854775807
    var v4 string = ""
    v5 := []uint32{0, 2147483647}
    var v6 int8 = 118
    var v7 int64 = -772
    var v8 uint64 = 4294967295
    switch 2147483647 {
    case 42:
        var v9 int16 = (v1 ^ v1)
        _ = v9
    case -1:
        fmt.Println(v1, v4)
    case -440:
        var v10 uint64 = v8
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 int16 = v1
        fmt.Println(v2)
        _ = fi11
        _ = v12
    }
    v6 = (v6 << ((42) & 63))
    var v13 int32 = v2
    switch v4 {
    case "hello":
        switch 42 {
        case 1:
            var v14 uint8 = 1
            _ = v14
        }
    case "foo bar":
        var v15 uint16 = 255
        _ = v15
    }
    var v16 int64 = -1
    if false {
        if true {
            var v17 int8 = (v6 << ((65535) & 63))
            fmt.Println(v13, v7)
            _ = v17
        }
        var v18 int8 = (-(v6))
        if true {
            v19 := GenS1{F0: 0, F1: 42}
            if false {
                var v20 float32 = (-2.25 + 0.1)
                v21 := GenS1{F0: 2, F1: 332}
                _ = v20
                _ = v21
            }
            switch 295 {
            case 1:
                var v22 uint64 = (v8 / ((v8) | 1))
                _ = v22
            case 592:
                for fi23 := 0; fi23 < 3; fi23++ {
                    if true {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            var v25 bool = true
                            if true {
                                if (fi24 < (fi23 ^ (fi23 * fi24))) {
                                    switch ((v4 + "a") + "\n\t") {
                                    case "foo bar":
                                        if (1 < 255) {
                                            fmt.Println(v7, v25, v4)
                                            v26 := GenS2{F0: uint(v8), F1: (v4 + (v4 + "foo bar"))}
                                            _ = v26
                                        } else {
                                            v1 = (v1 << ((4294967295) & 63))
                                            var v27 bool = (153 < (fi24 % ((fi23) | 1)))
                                            switch v4 {
                                            case "a":
                                                v6 += v18
                                            }
                                            _ = v27
                                        }
                                    case "":
                                        switch v4 {
                                        case "a":
                                            fmt.Println(v2)
                                            v25 = (-1 > (fi23 - fi24))
                                            switch fi24 {
                                            case -943:
                                                switch v4 {
                                                case "hello":
                                                    v28 := GenS1{F0: 255, F1: 32767}
                                                    _ = v28
                                                case "foo bar":
                                                    switch ((((((v4 + ("\n\t" + "a")) + v4) + v4) + (("\n\t" + v4) + ("\n\t" + v4))) + v4) + "\n\t") {
                                                    case "\n\t":
                                                        var v29 int64 = v3
                                                        _ = v29
                                                    default:
                                                        switch fi24 {
                                                        case 1:
                                                            fmt.Println(fi23, v8)
                                                        default:
                                                            if false {
                                                                fmt.Println(v13, v6, v4)
                                                            }
                                                        }
                                                    }
                                                    var v30 uint16 = 303
                                                    _ = v30
                                                case "a":
                                                    switch fi23 {
                                                    case 9223372036854775807:
                                                        v3 = 512
                                                    default:
                                                        var v31 int = int(v1)
                                                        _ = v31
                                                    }
                                                    v1 = v1
                                                    v4 = (v4 + "a")
                                                    fmt.Println(v7, v16, v1)
                                                default:
                                                    v32 := GenS1{F0: 1, F1: 42}
                                                    _ = v32
                                                }
                                            case 829:
                                                v16 -= v3
                                            case 32767:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                }
                                            default:
                                                v5 = append(v5, 42)
                                            }
                                            v6 = 100
                                        default:
                                            if v25 {
                                                v13 = (v13 - -240)
                                                v3 = v16
                                                v33 := GenS3{F0: 160, F1: 42, F2: 262}
                                                fmt.Println(v25, v16, fi23)
                                                switch (fi24 / ((-155) | 1)) {
                                                case 1:
                                                    var v34 int16 = (v1 + v1)
                                                    _ = v34
                                                case -390:
                                                    var v35 string = "\n\t"
                                                    _ = v35
                                                }
                                                v16 = 683
                                                var v36 float32 = ((1.5 - 3.14159) - 1e10)
                                                var v37 int8 = v18
                                                var v38 string = ("foo bar" + "")
                                                v39 := GenS4{F0: (v13 >> ((255) & 63)), F1: 1, F2: 1}
                                                fmt.Println(v37, v1)
                                                _ = v33
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                    case "a":
                                        v13 += (v2 - (v2 - -2147483648))
                                    }
                                } else {
                                    v3 += v16
                                    switch 32767 {
                                    case 2147483647:
                                        v40 := GenS4{F0: 2147483647, F1: 85, F2: v13}
                                        _ = v40
                                    case -462:
                                        switch v4 {
                                        case "hello":
                                            v7 = (v7 | (v7 ^ v3))
                                        case "":
                                            fmt.Println(v7, v13)
                                        default:
                                            v41 := GenS2{F0: 275, F1: "a"}
                                            _ = v41
                                        }
                                    case -265:
                                        fmt.Println(fi24, v8)
                                        v7 -= v7
                                        if false {
                                            v6 = -87
                                        } else {
                                            _ = len(v5)
                                        }
                                    default:
                                        fmt.Println(v16, fi24)
                                    }
                                }
                                v2 -= 2147483647
                            } else {
                                v42 := GenS4{F0: (v13 % ((v13) | 1)), F1: 32767, F2: 1}
                                switch (fi24 - fi24) {
                                case 2147483647:
                                    v5 = append(v5, 0)
                                case -59:
                                    var v43 float32 = (1.5 * 0.0)
                                    _ = v43
                                case 104:
                                    var v44 uint32 = 2147483647
                                    _ = v44
                                default:
                                    v13 = 571
                                }
                                _ = v42
                            }
                            _ = fi24
                            _ = v25
                        }
                    }
                    _ = fi23
                }
            default:
                for fi45 := 0; fi45 < 19; fi45++ {
                    var v46 uint = 0
                    fmt.Println(fi45)
                    var v47 int16 = (v1 + v1)
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            _ = v19
        }
        _ = v18
    }
    v8 = v8
    if false {
        v8 -= (v8 >> ((176) & 63))
    } else {
        v48 := GenS5{F0: (-(v13)), F1: true}
        var v49 int8 = 100
        _ = v48
        _ = v49
    }
    v13 = v13
    if (int32(v3) > (v2 << ((2) & 63))) {
        if (!true) {
            var v50 float32 = (1e10 - (((0.1 - 0.0) + 3.14159) + (1e10 / 1.0)))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v50
        } else {
            v3 = (v3 / ((v16) | 1))
        }
        v6 = v6
    } else {
        if true {
            var v51 int16 = -463
            _ = v51
        }
    }
    for fi52 := 0; fi52 < 13; fi52++ {
        v1 = v1
        _ = fi52
    }
    if false {
        var v53 float32 = 1e10
        _ = v53
    } else {
        var v54 int8 = v6
        var v55 int16 = v1
        _ = v54
        _ = v55
    }
    v56 := GenS5{F0: -970, F1: false}
    for fi57 := 0; fi57 < 14; fi57++ {
        fmt.Println(v1)
        v2 = 0
        for fi58 := 0; fi58 < 2; fi58++ {
            v1 += (v1 - 0)
            if false {
                if true {
                    if (true || (v8 > (v8 - (v8 / (((v8 & 1)) | 1))))) {
                        fmt.Println(v4, v1, fi57)
                        if len(v5) > 0 {
                            v5[0] = 382
                        }
                    }
                    var v59 float32 = (0.1 / 1.0)
                    _ = v59
                } else {
                    if ((^(v8)) > 2) {
                        var v60 int64 = v16
                        _ = v60
                    } else {
                        var v61 uint64 = (v8 + 480)
                        _ = v61
                    }
                    v5 = append(v5, 4294967295)
                    v7 -= 479
                    v16 = (v3 | -1)
                    v62 := GenS2{F0: 454, F1: ""}
                    _ = v62
                }
                fmt.Println(v7)
            } else {
                fmt.Println(v16, v4, v2)
                var v63 int16 = int16(v16)
                var v64 uint = 0
                _ = v63
                _ = v64
            }
            var v65 int = fi58
            v6 = v6
            if ((!true) || false) {
                v7 = (v16 << ((2) & 63))
                switch "\n\t" {
                case "\n\t":
                    v16 -= -9223372036854775808
                case "a":
                    var v66 bool = (!(false && false))
                    v3 = v3
                    _ = v66
                case "hello":
                    v67 := GenS6{F0: (-(v1))}
                    _ = v67
                }
                switch v4 {
                case "a":
                    var v68 uint = 2
                    var v69 float64 = (3.14159 / 1.0)
                    var v70 int8 = 0
                    _ = v68
                    _ = v69
                    _ = v70
                case "\n\t":
                    var v71 uint8 = 127
                    v72 := GenS7{F0: (v6 >> ((2) & 63)), F1: 179, F2: (v6 | (-(v6)))}
                    _ = v71
                    _ = v72
                case "foo bar":
                    v1 = 0
                default:
                    fmt.Println(v6, v2, v1)
                }
                fmt.Println(v13)
                v73 := GenS6{F0: v1}
                var v74 float32 = 0.1
                _ = v73
                _ = v74
            }
            _ = fi58
            _ = v65
        }
        _ = fi57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v7
    _ = v8
}
