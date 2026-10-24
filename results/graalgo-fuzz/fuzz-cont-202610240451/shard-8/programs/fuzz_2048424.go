package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int32; F1 bool; F2 int }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 int) int32 {
    if false {
        var v5 uint16 = 2
        var v6 int = (-(p0))
        _ = v5
        _ = v6
    } else {
        var v7 uint = 2
        v7 = (v7 * v7)
        v7 -= v7
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "" {
        case "foo bar":
            fmt.Println(p0, p0)
        }
        _ = fi8
    }
    v9 := []int{2147483647}
    v9 = append(v9, (^(p0)))
    var v10 uint64 = 807
    _ = p0
    _ = v10
    _ = v9
    return 42
}

func main() {
    var v1 uint = 42
    v2 := []int{}
    var v3 bool = ("" >= ("\n\t" + "hello"))
    v4 := make(map[uint32]uint32)
    switch -1 {
    case -2147483648:
        v3 = (32768 < 32767)
    case -378:
        v1 -= (v1 >> (((v1 % ((v1) | 1))) & 63))
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if v3 {
            delete(v4, 2)
        }
        _ = fi11
    }
    _ = len(v4)
    if false {
        var v12 uint32 = 2
        if len(v2) > 0 {
            v2[0] = -672
        }
        _ = v12
    }
    fmt.Println(v1, v1)
    var v13 int32 = 968
    v14 := v4[0]
    _ = len(v2)
    var v15 int16 = 973
    var v16 string = ("hello" + ("hello" + "foo bar"))
    switch 42 {
    case -621:
        for fi17 := 0; fi17 < 9; fi17++ {
            if v3 {
                _ = len(v4)
                for fi18 := 0; fi18 < 19; fi18++ {
                    fmt.Println(v15, fi18)
                    v3 = (false || v3)
                    _ = fi18
                }
                fmt.Println(v3, v3)
            } else {
                for fi19 := 0; fi19 < 17; fi19++ {
                    switch v16 {
                    case "foo bar":
                        if true {
                            var v20 uint32 = v14
                            _ = v20
                        }
                    case "a":
                        v21 := GenS1{F0: -81, F1: -9223372036854775808, F2: (fi19 ^ (fi17 - fi17))}
                        _ = v21
                    default:
                        if ("a" != "\n\t") {
                            v3 = v3
                            v16 = (v16 + v16)
                            switch ((v16 + "\n\t") + v16) {
                            case "\n\t":
                                switch -182 {
                                case 469:
                                    var v22 bool = (!true)
                                    _ = v22
                                case 32767:
                                    switch ("\n\t" + v16) {
                                    case "a":
                                        var v23 int64 = 0
                                        v23 = (v23 + (-(v23)))
                                        _ = v23
                                    }
                                case 0:
                                    fmt.Println(v16, v13)
                                }
                            case "foo bar":
                                var v24 uint = (v1 ^ v1)
                                _ = v24
                            default:
                                if v3 {
                                    var v25 int32 = -2147483648
                                    if (false && (uint16(v1) <= 679)) {
                                        v13 += (v13 * -2147483648)
                                        fmt.Println(v3)
                                        v26 := GenS2{F0: (^(v13)), F1: v3}
                                        _ = v26
                                    }
                                    if v3 {
                                        if len(v2) > 0 {
                                            v2[0] = fi19
                                        }
                                        switch ((v16 + v16) + ((v16 + v16) + v16)) {
                                        case "a":
                                            if false {
                                                v13 += (-(v13))
                                            }
                                        case "\n\t":
                                            fmt.Println(v13)
                                        case "":
                                            if (!false) {
                                                v27 := GenS2{F0: (v13 >> (((v1 * (v1 & (v1 << (((v1 % ((v1) | 1))) & 63))))) & 63)), F1: (false || v3)}
                                                if (-1 == 9223372036854775807) {
                                                    _ = len(v2)
                                                } else {
                                                    fmt.Println(fi17)
                                                }
                                                var v28 int64 = int64(v14)
                                                v29 := GenS2{F0: v25, F1: v3}
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v1 -= (^(v1))
                                                switch (fi17 * -1) {
                                                case -9223372036854775808:
                                                    fmt.Println(v3)
                                                case 2147483647:
                                                    var v30 int64 = 42
                                                    _ = v30
                                                case 128:
                                                    switch ("" + ("hello" + v16)) {
                                                    case "\n\t":
                                                        fmt.Println(v1, fi17, fi19)
                                                    case "":
                                                        var v31 int32 = (^(v25))
                                                        _ = v31
                                                    case "foo bar":
                                                        var v32 int32 = (-(v25))
                                                        _ = v32
                                                    default:
                                                        switch fi17 {
                                                        case 32767:
                                                            v33 := GenS3{F0: (v25 >> ((1) & 63)), F1: v3, F2: (fi19 + -1)}
                                                            _ = v33
                                                        case 179:
                                                            v15 += (v15 & 511)
                                                        case -242:
                                                            var v34 uint32 = v14
                                                            _ = v34
                                                        default:
                                                            v3 = (!true)
                                                        }
                                                    }
                                                default:
                                                    var v35 uint = (v1 + (v1 - (v1 - 988)))
                                                    _ = v35
                                                }
                                                var v36 float32 = -2.25
                                                _ = v36
                                            }
                                        }
                                        v4[(^(v14))] = 2147483647
                                    }
                                    v37 := GenS4{F0: -1, F1: false, F2: (2 != 322)}
                                    var v38 float64 = 0.1
                                    switch fi17 {
                                    case 32767:
                                        if v3 {
                                            fmt.Println(v14, fi17)
                                        }
                                    }
                                    _ = v25
                                    _ = v37
                                    _ = v38
                                } else {
                                    switch (fi19 >> (((v1 >> (((v1 >> (((v1 ^ 65535)) & 63))) & 63))) & 63)) {
                                    case 1:
                                        if ((-(v13)) <= (^(v13))) {
                                            v16 = v16
                                        }
                                    case 42:
                                        v1 += (^(v1))
                                    default:
                                        if v3 {
                                            var v39 uint64 = 268
                                            var v40 int64 = 9223372036854775807
                                            v40 = int64(v1)
                                            _ = v39
                                            _ = v40
                                        } else {
                                            var v41 bool = (255 > 32767)
                                            fmt.Println(fi19)
                                            _ = v41
                                        }
                                    }
                                    var v42 bool = true
                                    v43 := GenS5{F0: 791, F1: 917}
                                    _ = v42
                                    _ = v43
                                }
                            }
                        }
                        if len(v2) > 0 {
                            v2[0] = (fi17 + (-(fi17)))
                        }
                    }
                    _ = fi19
                }
            }
            var v44 int = 522
            _ = fi17
            _ = v44
        }
    case 23:
        delete(v4, (^(v14)))
    case 0:
        for fi45 := 0; fi45 < 9; fi45++ {
            v1 = v1
            switch fi45 {
            case 1:
                for fi46 := 0; fi46 < 9; fi46++ {
                    var v47 string = "a"
                    _ = fi46
                    _ = v47
                }
            case 9223372036854775807:
                var v48 int32 = v13
                _ = v48
            case -890:
                var v49 uint8 = 128
                _ = v49
            }
            if v3 {
                v50 := GenS3{F0: genFn0(fi45), F1: true, F2: -182}
                _ = v50
            }
            _ = fi45
        }
    }
    v4[2147483647] = v14
    v51 := GenS6{F0: (v1 / (((v1 * (v1 >> (((v1 & v1)) & 63)))) | 1)), F1: 104, F2: 18446744073709551615}
    var v52 int = 1
    var v53 int8 = 1
    if (-672 > -1) {
        switch -555 {
        case -933:
            var v54 bool = v3
            _ = v54
        }
    }
    var v55 uint32 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
    _ = v55
}
