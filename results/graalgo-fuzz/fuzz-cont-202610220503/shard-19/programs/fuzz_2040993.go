package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int32; F2 int8 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 bool; F2 uint }
type GenS6 struct { F0 int16; F1 int64; F2 bool }
type GenS7 struct { F0 uint; F1 int64; F2 string }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 int16; F1 uint; F2 int }

func main() {
    var v1 int16 = 32767
    var v2 uint8 = 42
    var v3 uint32 = 2
    var v4 bool = true
    var v5 string = "a"
    v6 := []int{268, 32767, 2147483647, int(v1)}
    if v4 {
        var v7 uint8 = v2
        var v8 string = "hello"
        _ = v7
        _ = v8
    }
    v6 = append(v6, 595)
    if (!v4) {
        v2 = v2
        if v4 {
            var v9 string = v5
            if len(v6) > 0 {
                v6[0] = -9223372036854775808
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                if v4 {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if len(v6) > 0 {
                            v6[0] = -2147483648
                        }
                        _ = fi11
                    }
                } else {
                    var v12 int8 = -69
                    _ = v12
                }
                switch "a" {
                case "a":
                    if (2147483647 == 1) {
                        if ((v4 && false) && (((true && false) && (false && ((int64(v1) >= 279) || v4))) && v4)) {
                            var v13 int32 = 865
                            v14 := GenS1{F0: (^(v2)), F1: (v3 >> ((4294967295) & 63))}
                            _ = v13
                            _ = v14
                        } else {
                            for fi15 := 0; fi15 < 2; fi15++ {
                                if v4 {
                                    if ((-525 >= -578) || false) {
                                        if (v2 == v2) {
                                            if (1 > v2) {
                                                if v4 {
                                                    _ = len(v6)
                                                    v16 := GenS1{F0: v2, F1: (v3 >> ((18446744073709551615) & 63))}
                                                    _ = v16
                                                }
                                                v17 := GenS1{F0: 191, F1: (^(v3))}
                                                _ = v17
                                            } else {
                                                fmt.Println(v5)
                                                var v18 int32 = 906
                                                fmt.Println(v5, fi15)
                                                _ = v18
                                            }
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = (^(fi15))
                                            }
                                            v19 := GenS1{F0: (v2 >> ((65535) & 63)), F1: v3}
                                            if v4 {
                                                switch "hello" {
                                                case "hello":
                                                    v20 := GenS2{F0: 2, F1: 0, F2: -88}
                                                    _ = v20
                                                case "":
                                                    fmt.Println(v9)
                                                case "foo bar":
                                                    v21 := GenS3{F0: (v1 / ((v1) | 1)), F1: (v3 & v3), F2: (^(v2))}
                                                    _ = v21
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = (fi15 << ((42) & 63))
                                                }
                                            }
                                            _ = v19
                                        }
                                        if ((-1 < 1) || (!((^(fi15)) == -78))) {
                                            var v22 int64 = 9223372036854775807
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v22
                                        } else {
                                            v3 = (v3 >> ((18446744073709551615) & 63))
                                            fmt.Println(fi15, v5, v5)
                                            if len(v6) > 0 {
                                                v6[0] = fi10
                                            }
                                            v23 := GenS3{F0: -240, F1: 42, F2: 115}
                                            var v24 float32 = (1e10 * 1.5)
                                            var v25 uint = 1
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        v26 := GenS4{F0: v5}
                                        _ = v26
                                    } else {
                                        v27 := GenS3{F0: (v1 * 1), F1: (v3 | v3), F2: 255}
                                        v28 := GenS5{F0: (!v4), F1: (v5 != v9), F2: 65535}
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 string = v5
                                    _ = v29
                                }
                                if (v4 || (2147483647 > -2147483648)) {
                                    if true {
                                        switch "hello" {
                                        case "hello":
                                            v2 = 0
                                            v2 -= (v2 / ((153) | 1))
                                            var v30 int = 4294967295
                                            _ = v30
                                        case "":
                                            v1 = 732
                                        case "\n\t":
                                            if (false && v4) {
                                                fmt.Println(v2)
                                            }
                                        }
                                        v31 := GenS5{F0: false, F1: (v1 > int16(v2)), F2: 18446744073709551615}
                                        if (-1 > -1) {
                                            var v32 float32 = (0.1 + float32(0.1))
                                            v5 = (v9 + v9)
                                            var v33 string = (((v9 + ((v9 + "\n\t") + "foo bar")) + v5) + "a")
                                            _ = v32
                                            _ = v33
                                        } else {
                                            var v34 float32 = 1.5
                                            v1 += v1
                                            if (477 > 4294967295) {
                                                fmt.Println(fi10, v4, v1)
                                                fmt.Println(fi10, fi15, v2)
                                            }
                                            _ = v34
                                        }
                                        _ = v31
                                    } else {
                                        var v35 uint32 = v3
                                        var v36 bool = (-2.25 > 0.1)
                                        if false {
                                            v2 -= 165
                                            switch (fi15 & 1) {
                                            case 2147483647:
                                                if true {
                                                    v1 -= int16(fi10)
                                                    var v37 uint64 = 42
                                                    _ = v37
                                                }
                                            case 73:
                                                if (3.14159 == (3.14159 / 1.5)) {
                                                    v3 = v3
                                                    v3 = v35
                                                } else {
                                                    if v36 {
                                                        var v38 int = 4294967295
                                                        var v39 int32 = 0
                                                        v4 = (0 <= 1)
                                                        if v4 {
                                                            switch (fi15 % (((fi10 | (fi15 - (fi10 | (v38 ^ (fi10 / ((fi10) | 1))))))) | 1)) {
                                                            case 2147483647:
                                                                if v4 {
                                                                    if ((v2 << ((4294967295) & 63)) >= v2) {
                                                                        _ = len(v6)
                                                                    }
                                                                    v39 = (-(v39))
                                                                }
                                                            case 524:
                                                                var v40 uint = 296
                                                                _ = v40
                                                            default:
                                                                v9 = "a"
                                                            }
                                                            var v41 uint8 = 128
                                                            _ = v41
                                                        } else {
                                                            v42 := GenS1{F0: v2, F1: v3}
                                                            v4 = v4
                                                            v39 += v39
                                                            _ = v42
                                                        }
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        var v43 float64 = 0.0
                                                        var v44 uint16 = 32767
                                                        v2 -= (v2 + 127)
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    var v45 int32 = -1
                                                    fmt.Println(v45, v1, v5)
                                                    _ = v45
                                                }
                                            }
                                            v46 := GenS2{F0: 65535, F1: 1, F2: -1}
                                            _ = v46
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v5, v5)
                                            case "":
                                                var v47 uint32 = v35
                                                _ = v47
                                            }
                                        }
                                        v48 := GenS6{F0: (v1 << ((255) & 63)), F1: 823, F2: ((fi10 | fi10) > 128)}
                                        _ = v35
                                        _ = v36
                                        _ = v48
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = fi15
                                    }
                                    v2 += v2
                                    var v49 uint = 2
                                    v4 = ((v3 << ((520) & 63)) != v3)
                                    _ = v49
                                }
                                _ = fi15
                            }
                        }
                        v1 += -32768
                    }
                case "\n\t":
                    if (!(((v4 && ((!(2 >= 297)) || v4)) || (100 > 0)) && v4)) {
                        v4 = (!v4)
                    } else {
                        var v50 uint32 = (v3 % (((v3 ^ v3)) | 1))
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 int8 = -95
                            if (v4 || v4) {
                                var v53 float32 = float32(((0.1 - -2.25) + float64(0.1)))
                                var v54 int32 = -1
                                if v4 {
                                    switch "\n\t" {
                                    case "hello":
                                        if (!(false || v4)) {
                                            fmt.Println(v9, fi51, v3)
                                            v5 = ("" + (v5 + v9))
                                        } else {
                                            var v55 int8 = (v52 * -68)
                                            _ = v55
                                        }
                                    default:
                                        fmt.Println(fi51, v4)
                                    }
                                }
                                _ = v53
                                _ = v54
                            } else {
                                v3 -= (v50 | (v3 % ((v50) | 1)))
                                v1 -= (v1 - 0)
                                v9 = "foo bar"
                            }
                            v56 := GenS7{F0: 18446744073709551615, F1: 42, F2: (v5 + v5)}
                            v57 := GenS8{F0: v52, F1: 63}
                            _ = fi51
                            _ = v52
                            _ = v56
                            _ = v57
                        }
                        var v58 uint8 = 127
                        _ = v50
                        _ = v58
                    }
                case "foo bar":
                    if (v9 < "foo bar") {
                        var v59 uint = 300
                        var v60 float64 = 0.1
                        if (v60 >= ((v60 - v60) * 1.5)) {
                            v1 += v1
                            var v61 string = v9
                            _ = v61
                        }
                        _ = v59
                        _ = v60
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if len(v6) > 0 {
                            v6[0] = fi10
                        }
                        v4 = (v1 <= v1)
                    }
                }
                if (!(0 <= 11)) {
                    v62 := GenS9{F0: v1, F1: 42, F2: fi10}
                    _ = v62
                }
                _ = fi10
            }
            _ = v9
        }
    }
    var v63 uint = 2
    v3 -= (v3 - v3)
    v3 -= v3
    var v64 int = 4294967295
    fmt.Println(v1, v63)
    fmt.Println(v3)
    if len(v6) > 0 {
        v6[0] = v64
    }
    v1 = v1
    v3 = (v3 >> (((v63 / ((2) | 1))) & 63))
    var v65 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
}
