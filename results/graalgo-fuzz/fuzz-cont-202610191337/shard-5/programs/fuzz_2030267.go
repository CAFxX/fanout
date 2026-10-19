package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 int32 }
type GenS4 struct { F0 int64; F1 int64; F2 int16 }
type GenS5 struct { F0 int16; F1 uint64 }
func genFn0(p0 int64) int8 {
    var v7 uint8 = 152
    var v8 int16 = 32767
    _ = p0
    _ = v7
    _ = v8
    return 20
}
func genFn1(p0 uint, p1 uint8) int8 {
    var v9 int64 = 9223372036854775807
    var v10 uint32 = 42
    v11 := GenS1{F0: 1000}
    v12 := []int{1, 1, 128, 1}
    p0 = (p0 * (^(p0)))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return -62
}

func main() {
    var v1 uint = 0
    var v2 int32 = int32(v1)
    var v3 bool = true
    var v4 int32 = v2
    var v5 int32 = -1
    var v6 uint16 = 1
    v13 := []int{127, -1, 4294967295}
    var v14 uint8 = 128
    if len(v13) > 0 {
        v13[0] = -836
    }
    if ((!v3) || (v6 == (v6 - (^(v6))))) {
        v14 += 127
    } else {
        v15 := GenS1{F0: 32767}
        v4 -= v5
        if len(v13) > 0 {
            v13[0] = 1
        }
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if true {
            for fi17 := 0; fi17 < 5; fi17++ {
                if (v3 && true) {
                    v18 := GenS1{F0: 32767}
                    var v19 uint32 = 1
                    v19 += (^(v19))
                    _ = v18
                    _ = v19
                } else {
                    switch "\n\t" {
                    case "foo bar":
                        v13 = append(v13, (fi17 ^ fi16))
                    default:
                        var v20 int = (fi16 * (fi17 ^ fi17))
                        _ = v20
                    }
                    switch (fi17 & 1) {
                    case -115:
                        v1 += 65535
                    default:
                        var v21 float64 = 1e10
                        _ = v21
                    }
                }
                v13 = append(v13, (fi16 | 0))
                _ = fi17
            }
        }
        v14 += (v14 - v14)
        if ((true && v3) && v3) {
            var v22 uint32 = 2147483647
            v23 := GenS1{F0: 0}
            v14 = v14
            _ = v22
            _ = v23
        }
        _ = fi16
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v1 = v1
    if ((v14 + (v14 * 2)) > 128) {
        var v24 uint32 = 2147483647
        if len(v13) > 0 {
            v13[0] = -1
        }
        if (((0.1 - 0.0) - (((-2.25 - 1e10) + 1e10) * ((3.14159 / ((1.5 * (1.5 * 1.5)) - 0.1)) * 0.1))) >= 1.5) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        }
        _ = v24
    } else {
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                if ((!((0 != int16(v2)) && (!v3))) || ((v1 * v1) == uint(v6))) {
                    v27 := GenS1{F0: -32768}
                    v28 := GenS2{F0: 111, F1: 421}
                    v29 := GenS2{F0: 1, F1: 1}
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    v2 = 42
                    var v30 int32 = v2
                    _ = v30
                }
                _ = fi26
            }
            v4 = -2147483648
            _ = fi25
        }
    }
    if false {
        var v31 uint16 = v6
        if ((81 == (v6 << (((v1 << ((0) & 63))) & 63))) || ((42 != 44) && (286 == -2147483648))) {
            v2 -= 42
            v3 = (false && true)
        } else {
            v5 -= v4
            if v3 {
                v2 = (v5 ^ v4)
                if v3 {
                    for fi32 := 0; fi32 < 19; fi32++ {
                        for fi33 := 0; fi33 < 5; fi33++ {
                            switch ("foo bar" + "\n\t") {
                            case "\n\t":
                                if len(v13) > 0 {
                                    v13[0] = fi32
                                }
                            }
                            v34 := GenS2{F0: 1, F1: uint(v14)}
                            _ = fi33
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 18; fi35++ {
                            v31 = (v6 - (v31 << (((v1 >> (((v1 * (v1 % (((v1 / ((626) | 1))) | 1)))) & 63))) & 63)))
                            v1 = v1
                            v36 := GenS3{F0: -128, F1: v14, F2: v4}
                            _ = fi35
                            _ = v36
                        }
                        _ = fi32
                    }
                    v1 = 2
                    v31 = (v6 * 2)
                } else {
                    if false {
                        for fi37 := 0; fi37 < 13; fi37++ {
                            v38 := GenS1{F0: 42}
                            var v39 uint8 = (v14 | uint8(v1))
                            fmt.Println(v2, v2)
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        var v40 uint64 = 42
                        switch 2147483647 {
                        case 32768:
                            v13 = append(v13, 32767)
                        case 108:
                            v41 := GenS4{F0: 0, F1: 42, F2: -435}
                            _ = v41
                        default:
                            v31 += v6
                        }
                        _ = v40
                    }
                }
                var v42 float64 = 1.5
                _ = v42
            }
            fmt.Println(v4, v6)
        }
        for fi43 := 0; fi43 < 9; fi43++ {
            if true {
                switch ("a" + "") {
                case "a":
                    switch fi43 {
                    case -9223372036854775808:
                        var v44 uint64 = 670
                        _ = v44
                    case -584:
                        _ = len(v13)
                    default:
                        var v45 int32 = (v4 + int32(v6))
                        _ = v45
                    }
                case "":
                    for fi46 := 0; fi46 < 7; fi46++ {
                        v4 -= v4
                        switch fi43 {
                        case 9223372036854775807:
                            var v47 string = ("" + "\n\t")
                            _ = v47
                        default:
                            if v3 {
                                fmt.Println(v14, fi43, v6)
                                var v48 int64 = -1
                                _ = v48
                            } else {
                                if v3 {
                                    v49 := GenS1{F0: int16(v1)}
                                    v5 -= (v5 - int32(v31))
                                    _ = v49
                                } else {
                                    v50 := GenS3{F0: 1, F1: v14, F2: (v4 - (v2 & v2))}
                                    v51 := GenS5{F0: 0, F1: 18446744073709551615}
                                    _ = len(v13)
                                    _ = v50
                                    _ = v51
                                }
                                _ = len(v13)
                                v1 = v1
                            }
                        }
                        var v52 int16 = 42
                        _ = fi46
                        _ = v52
                    }
                case "hello":
                    if v3 {
                        var v53 int64 = -1
                        if v3 {
                            var v54 string = "foo bar"
                            _ = v54
                        }
                        _ = v53
                    } else {
                        v6 = v6
                        for fi55 := 0; fi55 < 7; fi55++ {
                            fmt.Println(v6, fi55, v31)
                            _ = fi55
                        }
                        for fi56 := 0; fi56 < 4; fi56++ {
                            v31 = v6
                            _ = fi56
                        }
                    }
                }
                switch (fi43 >> (((^(v1))) & 63)) {
                case -701:
                    var v57 uint64 = 1
                    _ = v57
                case 42:
                    v4 = (v4 << (((v1 >> (((v1 << ((v1) & 63))) & 63))) & 63))
                case 2147483647:
                    if len(v13) > 0 {
                        v13[0] = (fi43 * -1)
                    }
                }
            } else {
                v6 += (v6 & v31)
            }
            if false {
                fmt.Println(v5, v3, v5)
            }
            _ = fi43
        }
        _ = v31
    }
    var v58 uint64 = 0
    v6 = (^(v6))
    switch 32767 {
    case 1:
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    case 32767:
        for fi59 := 0; fi59 < 7; fi59++ {
            _ = len(v13)
            fmt.Println(v14, v5)
            var v60 uint16 = (v6 - v6)
            _ = fi59
            _ = v60
        }
    }
    var v61 int32 = (v5 * 2147483647)
    switch "\n\t" {
    case "":
        v6 = (v6 - (^(v6)))
    default:
        var v62 float32 = 1.5
        _ = v62
    }
    _ = len(v13)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v61
}
