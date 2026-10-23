package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int32 }
type GenS6 struct { F0 uint64; F1 int; F2 int16 }
type GenS7 struct { F0 int; F1 int32; F2 bool }
type GenS8 struct { F0 uint8 }

func main() {
    var v1 int8 = 127
    var v2 uint32 = 356
    var v3 int64 = 9223372036854775807
    v4 := []uint32{v2, v2, v2, (v2 % (((v2 >> ((255) & 63))) | 1))}
    var v5 uint32 = 4294967295
    var v6 uint = 0
    v6 -= uint(v5)
    v7 := GenS1{F0: v6, F1: v3}
    if true {
        if (v3 <= (v3 + (v3 * v3))) {
            v2 = (v2 & v2)
            for fi8 := 0; fi8 < 6; fi8++ {
                for fi9 := 0; fi9 < 10; fi9++ {
                    var v10 int16 = -32768
                    switch (fi9 << ((v6) & 63)) {
                    case 121:
                        if true {
                            v6 -= v6
                            v11 := GenS1{F0: 9, F1: v3}
                            v12 := GenS2{F0: "hello"}
                            _ = v11
                            _ = v12
                        } else {
                            switch 1 {
                            case 4294967295:
                                var v13 int16 = (^(v10))
                                _ = v13
                            default:
                                var v14 bool = true
                                _ = v14
                            }
                            var v15 int16 = v10
                            _ = v15
                        }
                    case 9223372036854775807:
                        _ = len(v4)
                    default:
                        if (false || (566 <= -2147483648)) {
                            var v16 uint16 = 32767
                            v17 := GenS3{F0: uint8(v16)}
                            var v18 int8 = v1
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    }
                    switch fi8 {
                    case -674:
                        var v19 int16 = v10
                        _ = v19
                    case 202:
                        fmt.Println(v5)
                    case 127:
                        var v20 uint8 = 2
                        _ = v20
                    default:
                        var v21 int16 = v10
                        _ = v21
                    }
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            v2 = 1
        } else {
            v5 -= (v2 | v5)
        }
        v6 -= v6
        var v22 int = -280
        _ = v22
    } else {
        v23 := GenS3{F0: 61}
        _ = v23
    }
    var v24 float32 = 3.14159
    fmt.Println(v2, v6)
    v1 = (v1 % (((v1 * (v1 ^ 0))) | 1))
    switch 127 {
    case 4294967295:
        fmt.Println(v5, v1, v3)
    case -978:
        switch 4294967295 {
        case 422:
            _ = len(v4)
        case 797:
            v5 += 1
        case -62:
            v6 = (v6 >> ((65535) & 63))
        default:
            switch ("\n\t" + "foo bar") {
            case "hello":
                switch "foo bar" {
                case "foo bar":
                    _ = len(v4)
                default:
                    var v25 int = -396
                    _ = v25
                }
            default:
                var v26 bool = false
                _ = v26
            }
        }
    case -2147483648:
        switch 32768 {
        case 9223372036854775807:
            var v27 int = 128
            _ = v27
        }
    default:
        v28 := GenS4{F0: 77, F1: (v2 + (v2 << (((v6 / ((370) | 1))) & 63))), F2: 18446744073709551615}
        _ = v28
    }
    var v29 int64 = (v3 & v3)
    switch 42 {
    case 0:
        var v30 int32 = 42
        _ = v30
    case 4294967295:
        fmt.Println(v29)
    case 1:
        v6 += 255
    default:
        for fi31 := 0; fi31 < 2; fi31++ {
            for fi32 := 0; fi32 < 10; fi32++ {
                var v33 int = fi32
                v33 -= v33
                _ = len(v4)
                _ = fi32
                _ = v33
            }
            if true {
                if true {
                    if false {
                        v5 -= (^(v5))
                        v34 := GenS3{F0: 2}
                        _ = v34
                    }
                    v3 = v29
                    v29 -= int64(v1)
                } else {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        var v36 uint = v6
                        _ = len(v4)
                        if true {
                            v5 = v2
                            if (false || (true || true)) {
                                switch "" {
                                case "hello":
                                    var v37 bool = true
                                    _ = v37
                                case "\n\t":
                                    if true {
                                        if len(v4) > 0 {
                                            v4[0] = (v5 / (((v5 << (((v6 << (((v6 * (v6 % ((v6) | 1)))) & 63))) & 63))) | 1))
                                        }
                                        v38 := GenS5{F0: 128, F1: -1}
                                        _ = v38
                                    }
                                case "foo bar":
                                    v6 = v36
                                default:
                                    v4 = append(v4, v2)
                                }
                                _ = len(v4)
                                v36 = 755
                            } else {
                                switch fi35 {
                                case 2147483647:
                                    var v39 bool = (!(213 <= 255))
                                    _ = v39
                                default:
                                    v36 = (v36 & (v36 - 230))
                                }
                            }
                            var v40 uint = (^(v36))
                            v41 := GenS3{F0: 164}
                            if false {
                                var v42 int = fi31
                                v36 = 679
                                if true {
                                    var v43 uint32 = (v5 + (v2 % ((v2) | 1)))
                                    switch (v42 - (v42 % ((32767) | 1))) {
                                    case -1:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if false {
                                            var v44 int8 = 100
                                            _ = v44
                                        } else {
                                            v2 = (^(v5))
                                        }
                                        v5 += v43
                                    case -2147483648:
                                        v3 = v3
                                    case -271:
                                        if len(v4) > 0 {
                                            v4[0] = v5
                                        }
                                    }
                                    fmt.Println(v40, v5)
                                    _ = v43
                                } else {
                                    var v45 uint8 = 2
                                    _ = v45
                                }
                                _ = v42
                            }
                            _ = v40
                            _ = v41
                        }
                        _ = fi35
                        _ = v36
                    }
                    for fi46 := 0; fi46 < 10; fi46++ {
                        fmt.Println(v5)
                        _ = fi46
                    }
                }
                v47 := GenS6{F0: 290, F1: -2147483648, F2: 42}
                _ = v47
            } else {
                if false {
                    fmt.Println(v29, v5)
                    if (false || true) {
                        if (true && (v5 <= v2)) {
                            v5 += (v5 | (v2 << ((v6) & 63)))
                            for fi48 := 0; fi48 < 1; fi48++ {
                                v29 += -174
                                var v49 uint64 = 42
                                _ = fi48
                                _ = v49
                            }
                            v1 += v1
                        } else {
                            fmt.Println(v1, v1)
                            var v50 float32 = (v24 / (v24 - (0.0 - v24)))
                            _ = v50
                        }
                    } else {
                        var v51 uint16 = 42
                        _ = v51
                    }
                    fmt.Println(v1, v1)
                }
                v3 += v29
                v3 -= 1
            }
            _ = fi31
        }
    }
    v52 := GenS7{F0: 2147483647, F1: -1, F2: (false && (1000 < 30))}
    for fi53 := 0; fi53 < 14; fi53++ {
        var v54 int8 = v1
        fmt.Println(fi53)
        v6 += (v6 % ((v6) | 1))
        _ = fi53
        _ = v54
    }
    v2 -= v5
    switch 42 {
    case -190:
        fmt.Println(v6, v2)
    case -9223372036854775808:
        var v55 int8 = v1
        _ = v55
    case 4294967295:
        v4 = append(v4, v5)
    }
    v56 := GenS8{F0: 128}
    var v57 uint = v6
    v6 = v57
    v2 = 0
    switch 32768 {
    case 1:
        var v58 int8 = -1
        _ = v58
    case 0:
        fmt.Println(v2, v24)
    case 42:
        for fi59 := 0; fi59 < 16; fi59++ {
            var v60 uint32 = (v2 / (((v2 << (((v57 & (v57 & v57))) & 63))) | 1))
            if (!false) {
                v1 = (v1 << ((v57) & 63))
                _ = len(v4)
            }
            _ = fi59
            _ = v60
        }
    default:
        v6 = (v6 - v57)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v56
    _ = v57
    _ = v6
    _ = v7
}
