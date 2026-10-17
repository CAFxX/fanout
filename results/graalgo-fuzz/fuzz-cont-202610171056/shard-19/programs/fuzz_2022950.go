package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 int16; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int32; F2 int8 }

func main() {
    var v1 uint32 = 0
    var v2 int32 = 0
    var v3 int8 = 1
    var v4 uint16 = 255
    v5 := make(map[uint32]uint32)
    v6 := GenS1{F0: 0, F1: v2, F2: (v2 / ((v2) | 1))}
    v7 := []int{32767, -542, 561}
    v7 = append(v7, -2147483648)
    v2 -= 668
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int8 = (v3 & (v3 % ((57) | 1)))
        _ = fi8
        _ = v9
    }
    v1 += v1
    v2 = (v2 >> ((190) & 63))
    if false {
        var v10 float64 = 1e10
        var v11 bool = true
        _ = v10
        _ = v11
    } else {
        var v12 bool = (127 < 841)
        v13 := GenS2{F0: uint(v4)}
        for fi14 := 0; fi14 < 11; fi14++ {
            delete(v5, v1)
            fmt.Println(v1, v1, v4)
            _ = fi14
        }
        _ = v12
        _ = v13
    }
    v4 = v4
    if (1 <= 49) {
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 uint64 = 0
                if false {
                    v18 := GenS2{F0: 922}
                    if (0.1 <= -2.25) {
                        fmt.Println(v4)
                        v19 := GenS1{F0: -9223372036854775808, F1: -2147483648, F2: (v2 / ((int32(v1)) | 1))}
                        v3 -= v3
                        var v20 uint32 = (v1 / ((v1) | 1))
                        v21 := GenS2{F0: 0}
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                } else {
                    v4 -= (v4 - (v4 * (v4 & v4)))
                    v22 := GenS2{F0: 114}
                    var v23 int32 = v2
                    fmt.Println(v4)
                    _ = v22
                    _ = v23
                }
                switch ("hello" + "foo bar") {
                case "hello":
                    v17 += 0
                default:
                    var v24 int16 = int16(v1)
                    _ = v24
                }
                _ = fi16
                _ = v17
            }
            v2 = 42
            if ((!true) || (!true)) {
                fmt.Println(v3, v2)
            }
            _ = fi15
        }
        v3 = v3
    } else {
        if (560 >= -1) {
            v25 := GenS3{F0: 127}
            var v26 int64 = 955
            v3 += v3
            _ = v25
            _ = v26
        } else {
            if ((v2 ^ 42) <= v2) {
                v1 -= v1
                v7 = append(v7, 2147483647)
                for fi27 := 0; fi27 < 3; fi27++ {
                    var v28 float32 = (0.0 / 3.14159)
                    if (!(255 >= 867)) {
                        var v29 uint64 = 933
                        if true {
                            fmt.Println(v2, v1)
                            var v30 float64 = ((((((3.14159 * 0.1) * 3.14159) / ((0.1 / 1.0) - -2.25)) - (float64(v28) / 0.1)) - 1e10) + (1e10 * -2.25))
                            _ = v30
                        }
                        switch (fi27 * (fi27 & (fi27 | fi27))) {
                        case 0:
                            var v31 uint = 2
                            _ = v31
                        case 1:
                            for fi32 := 0; fi32 < 12; fi32++ {
                                var v33 uint16 = v4
                                _ = fi32
                                _ = v33
                            }
                        case -373:
                            for fi34 := 0; fi34 < 3; fi34++ {
                                v2 -= v2
                                fmt.Println(v29, fi34, v3)
                                _ = fi34
                            }
                        default:
                            var v35 int64 = 9223372036854775807
                            _ = v35
                        }
                        _ = v29
                    } else {
                        for fi36 := 0; fi36 < 19; fi36++ {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            if ((42 != 198) && false) {
                                var v37 uint8 = 0
                                v2 = (v2 + (v2 + -2147483648))
                                _ = v37
                            }
                            _ = fi36
                        }
                    }
                    _ = fi27
                    _ = v28
                }
            } else {
                if false {
                    if (255 > 2) {
                        var v38 float32 = ((1.5 + 0.1) / 1.0)
                        _ = v38
                    }
                    if (-9223372036854775808 <= 9223372036854775807) {
                        switch 32767 {
                        case 42:
                            v7 = append(v7, 1)
                        }
                        v1 += (v1 << ((42) & 63))
                        fmt.Println(v3, v3)
                    } else {
                        for fi39 := 0; fi39 < 11; fi39++ {
                            var v40 uint = 255
                            _ = fi39
                            _ = v40
                        }
                        switch "foo bar" {
                        case "\n\t":
                            for fi41 := 0; fi41 < 20; fi41++ {
                                v1 = 224
                                v42 := GenS4{F0: 42, F1: -9223372036854775808}
                                _ = fi41
                                _ = v42
                            }
                        case "hello":
                            var v43 int = -9223372036854775808
                            _ = v43
                        }
                    }
                }
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            for fi44 := 0; fi44 < 1; fi44++ {
                v3 = (v3 + (v3 >> ((4294967295) & 63)))
                if false {
                    v45 := GenS5{F0: 0, F1: 0}
                    if true {
                        if true {
                            var v46 int8 = (^(v3))
                            var v47 uint = 0
                            v1 += v1
                            _ = v46
                            _ = v47
                        }
                        var v48 uint8 = 43
                        _ = v48
                    }
                    _ = v45
                } else {
                    var v49 uint32 = (v1 + (v1 % ((v1) | 1)))
                    _ = v49
                }
                _ = fi44
            }
        }
        v50 := GenS6{F0: (v3 % (((v3 + v3)) | 1))}
        if (((false || (((uint(v1) <= 402) || ((903 <= 32767) || (0 == 42))) || true)) && (true && (32767 <= -927))) && ((("hello" + ((("\n\t" + ("" + ("hello" + "a"))) + "\n\t") + ("" + (("hello" + "foo bar") + "hello")))) + "hello") < "")) {
            var v51 uint64 = 177
            _ = v51
        }
        _ = v50
    }
    if (false && false) {
        fmt.Println(v1, v2)
    }
    if true {
        v52 := GenS7{F0: 18446744073709551615, F1: uint(v1)}
        _ = v52
    }
    _ = len(v7)
    if (!("hello" <= ("\n\t" + "a"))) {
        v53 := GenS8{F0: -32768, F1: 86}
        delete(v5, v1)
        v54 := GenS9{F0: 4294967295, F1: -2147483648, F2: v3}
        _ = v53
        _ = v54
    }
    switch "a" {
    case "a":
        var v55 float64 = ((float64(3.14159) - -2.25) * 0.0)
        _ = v55
    case "hello":
        v3 += (v3 * (^(v3)))
    case "\n\t":
        v3 = (v3 - v3)
    default:
        v56 := GenS2{F0: 42}
        _ = v56
    }
    v1 += (v1 | 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
