package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int8; F1 int32; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint16; F2 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 198
    var v3 int8 = 42
    var v4 int16 = 32767
    v5 := GenS1{F0: "hello"}
    var v6 bool = (!true)
    v7 := GenS1{F0: (("\n\t" + "\n\t") + ("a" + "\n\t"))}
    v8 := []int{-129, 737}
    fmt.Println(v6)
    v6 = false
    if (!true) {
        v9 := v1[2]
        _ = v9
    } else {
        var v10 string = "hello"
        var v11 uint16 = 452
        _ = v10
        _ = v11
    }
    v1[v2] = v6
    if (!false) {
        var v12 int32 = -1
        var v13 uint8 = 42
        var v14 uint16 = 229
        _ = v12
        _ = v13
        _ = v14
    }
    if ((0.1 * ((1.5 + -2.25) - float32(1e10))) < 1e10) {
        if (-668 != 1) {
            for fi15 := 0; fi15 < 12; fi15++ {
                if (!v6) {
                    v6 = (v4 == (-(v4)))
                } else {
                    v16 := GenS1{F0: ("" + ("a" + (("foo bar" + "") + "hello")))}
                    _ = v16
                }
                v17 := v1[v2]
                _ = fi15
                _ = v17
            }
            var v18 uint64 = 2
            _ = v18
        }
    }
    var v19 uint16 = 148
    switch 32767 {
    case 127:
        v3 += (v3 | (v3 & v3))
    case 32767:
        _ = len(v1)
    }
    if (128 > 117) {
        delete(v1, uint32(v19))
        v6 = false
        v2 -= (^(v2))
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        if v6 {
            var v21 uint8 = 1
            v1[(v2 / (((v2 * 494)) | 1))] = (!(181 > 255))
            _ = v21
        } else {
            for fi22 := 0; fi22 < 5; fi22++ {
                if (1 < 1) {
                    switch int(v19) {
                    case 127:
                        v3 = (v3 & int8(fi20))
                    default:
                        v3 = (v3 - (v3 * (^(v3))))
                    }
                    v6 = v6
                }
                if v6 {
                    if (!(!true)) {
                        if false {
                            v4 += (v4 | (v4 | (-(v4))))
                            v23 := GenS1{F0: "foo bar"}
                            _ = v23
                        } else {
                            fmt.Println(fi20)
                        }
                        v6 = ((v3 << ((255) & 63)) > v3)
                        _ = len(v8)
                    } else {
                        v24 := GenS1{F0: ("foo bar" + "hello")}
                        _ = v24
                    }
                }
                if v6 {
                    var v25 uint8 = 128
                    switch fi20 {
                    case 2147483647:
                        v25 = uint8(v2)
                    default:
                        v26 := v1[v2]
                        _ = v26
                    }
                    _ = v25
                } else {
                    var v27 uint64 = 1
                    var v28 uint16 = (v19 << ((196) & 63))
                    if v6 {
                        v28 = (v19 + 65535)
                        switch (fi22 << ((2) & 63)) {
                        case 9223372036854775807:
                            v6 = v6
                        default:
                            var v29 int32 = -137
                            _ = v29
                        }
                        if len(v8) > 0 {
                            v8[0] = 2147483647
                        }
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi22
            }
            if (v4 >= v4) {
                fmt.Println(fi20, v4, v4)
                switch (fi20 + fi20) {
                case -9223372036854775808:
                    var v30 uint64 = 0
                    _ = v30
                case -261:
                    var v31 int64 = 9223372036854775807
                    _ = v31
                case 524:
                    if (true && false) {
                        v4 -= (^(v4))
                        if v6 {
                            v4 = v4
                        } else {
                            for fi32 := 0; fi32 < 6; fi32++ {
                                fmt.Println(v4, v3, v6)
                                var v33 uint32 = v2
                                _ = fi32
                                _ = v33
                            }
                        }
                    } else {
                        var v34 uint = 18446744073709551615
                        _ = v34
                    }
                }
            } else {
                if v6 {
                    var v35 uint = uint(v19)
                    if ((v2 / (((^(v2))) | 1)) == 223) {
                        switch "\n\t" {
                        case "a":
                            if ((0.1 - 1e10) >= 0.1) {
                                switch fi20 {
                                case 2147483647:
                                    if v6 {
                                        v4 = (v4 << ((v35) & 63))
                                        v1[2147483648] = (!v6)
                                        _ = len(v1)
                                    } else {
                                        for fi36 := 0; fi36 < 19; fi36++ {
                                            v35 = v35
                                            v6 = (!v6)
                                            v6 = (186 < 97)
                                            if v6 {
                                                if v6 {
                                                    fmt.Println(v19)
                                                }
                                                v37 := GenS2{F0: v3, F1: 1, F2: -151}
                                                _ = v37
                                            } else {
                                                var v38 string = "\n\t"
                                                fmt.Println(v2, v4)
                                                _ = v38
                                            }
                                            var v39 bool = (((0.1 == float64(1.5)) && v6) || (3.14159 >= 0.1))
                                            if (fi20 >= 4294967295) {
                                                fmt.Println(v3, fi36)
                                                var v40 uint64 = 71
                                                if (-339 == -356) {
                                                    if (1 <= 42) {
                                                        var v41 int64 = -9223372036854775808
                                                        fmt.Println(v2)
                                                        _ = v41
                                                    }
                                                    if false {
                                                        v40 = 4294967295
                                                    } else {
                                                        fmt.Println(v2, fi20)
                                                    }
                                                    fmt.Println(fi36)
                                                    if v6 {
                                                        fmt.Println(fi36, v19, v39)
                                                    }
                                                    v4 += 1
                                                }
                                                _ = v40
                                            }
                                            var v42 uint16 = v19
                                            _ = fi36
                                            _ = v39
                                            _ = v42
                                        }
                                        if (((42 >= 2) || v6) && v6) {
                                            switch ("foo bar" + ("" + "\n\t")) {
                                            case "\n\t":
                                                var v43 int16 = 1000
                                                _ = v43
                                            case "a":
                                                if len(v8) > 0 {
                                                    v8[0] = (fi20 ^ (^(fi20)))
                                                }
                                            case "hello":
                                                v8 = append(v8, 96)
                                            }
                                        }
                                    }
                                default:
                                    _ = len(v1)
                                }
                                v6 = false
                                if len(v8) >= 2 {
                                    v8 = v8[1:len(v8)]
                                }
                            }
                        case "hello":
                            var v44 uint8 = 1
                            _ = v44
                        }
                        if (!("" != ("a" + "foo bar"))) {
                            v3 -= v3
                            if (0.0 > 3.14159) {
                                v45 := GenS3{F0: (v19 | 57)}
                                var v46 string = "a"
                                delete(v1, v2)
                                _ = v45
                                _ = v46
                            }
                        } else {
                            v47 := GenS4{F0: -1, F1: 42, F2: 0}
                            _ = v47
                        }
                        v4 = (^(v4))
                    } else {
                        switch fi20 {
                        case -9223372036854775808:
                            v6 = v6
                        case 315:
                            v48 := GenS3{F0: (v19 >> ((671) & 63))}
                            _ = v48
                        case 128:
                            fmt.Println(v35, v2, v19)
                        }
                    }
                    switch int(v19) {
                    case 42:
                        v49 := GenS4{F0: v3, F1: 42, F2: 42}
                        _ = v49
                    case 1:
                        v2 = (v2 % ((v2) | 1))
                    case -702:
                        v2 = v2
                    default:
                        v50 := GenS5{F0: (v4 / ((-32768) | 1))}
                        _ = v50
                    }
                    _ = v35
                } else {
                    v3 = -77
                    switch 9223372036854775807 {
                    case -2147483648:
                        if false {
                            fmt.Println(v4, v2)
                            if len(v8) > 0 {
                                v8[0] = 127
                            }
                            switch "foo bar" {
                            case "hello":
                                var v51 float64 = 1.5
                                _ = v51
                            }
                        } else {
                            var v52 uint = 2
                            v3 = (v3 / ((v3) | 1))
                            _ = v52
                        }
                    case -9223372036854775808:
                        if v6 {
                            v3 -= (v3 << ((uint(v19)) & 63))
                            v53 := GenS1{F0: ""}
                            _ = v53
                        }
                    default:
                        v19 += (v19 ^ 32767)
                    }
                }
                v54 := GenS6{F0: (v3 + v3), F1: 1, F2: ""}
                _ = v54
            }
        }
        v4 = (^(v4))
        v4 = v4
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
