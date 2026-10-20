package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int64 }
func genFn0() uint {
    var v5 uint32 = 42
    v6 := GenS1{F0: -32768, F1: 848, F2: 42}
    if true {
        v5 = (^(v5))
    }
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int16 = 0
    var v2 int32 = -67
    var v3 uint64 = 0
    v4 := make(map[int]string)
    if false {
        fmt.Println(v2, v3, v3)
        switch 175 {
        case -9223372036854775808:
            switch "" {
            case "":
                v3 = v3
            case "\n\t":
                for fi7 := 0; fi7 < 10; fi7++ {
                    v2 += (v2 & -195)
                    if false {
                        v4[(fi7 | (fi7 + (fi7 % ((42) | 1))))] = "hello"
                        switch fi7 {
                        case -159:
                            v1 = v1
                        case -9223372036854775808:
                            v8 := v4[fi7]
                            _ = v8
                        case -402:
                            if true {
                                v9 := GenS2{F0: int32(v1)}
                                var v10 int = -9223372036854775808
                                _ = v10
                                _ = v9
                            }
                        default:
                            v11 := []int{fi7, fi7, (fi7 ^ 128)}
                            _ = v11
                        }
                        var v12 int64 = 9223372036854775807
                        _ = v12
                    } else {
                        var v13 uint32 = 2147483648
                        _ = v13
                    }
                    delete(v4, 1)
                    _ = fi7
                }
            case "foo bar":
                v4[434] = ""
            }
        default:
            switch -553 {
            case -368:
                v14 := GenS2{F0: v2}
                _ = v14
            }
        }
    }
    if (((^(v1)) != v1) || false) {
        for fi15 := 0; fi15 < 12; fi15++ {
            v4[0] = ("foo bar" + "\n\t")
            _ = fi15
        }
    }
    v1 += (v1 ^ v1)
    v1 = (v1 ^ (v1 + v1))
    var v16 int32 = -2147483648
    v17 := []int{-449, 32767, 42}
    for fi18 := 0; fi18 < 4; fi18++ {
        if len(v17) > 0 {
            v17[0] = fi18
        }
        if true {
            var v19 uint = 289
            v4[(fi18 ^ (fi18 >> ((65535) & 63)))] = (((("hello" + "\n\t") + ("hello" + "foo bar")) + ((("" + (("a" + ("a" + "a")) + ("\n\t" + ""))) + "a") + "\n\t")) + "")
            _ = v19
        } else {
            v2 = v16
            switch "foo bar" {
            case "\n\t":
                var v20 uint = 1
                _ = v20
            case "":
                var v21 float32 = (float32(1e10) - 1.5)
                _ = v21
            case "a":
                for fi22 := 0; fi22 < 15; fi22++ {
                    switch 4294967295 {
                    case 2147483647:
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    case 4294967295:
                        var v24 string = "a"
                        _ = v24
                    case 0:
                        delete(v4, (fi22 / (((-(fi22))) | 1)))
                    }
                    v16 -= (v16 * v16)
                    v2 = (v2 & (v2 + (^(v2))))
                    _ = fi22
                }
            default:
                var v25 bool = true
                _ = v25
            }
            v1 = (^(v1))
        }
        _ = fi18
    }
    fmt.Println(v3)
    fmt.Println(v2)
    switch "\n\t" {
    case "\n\t":
        if true {
            v1 = (v1 | -32768)
            switch "foo bar" {
            case "hello":
                var v26 uint8 = 162
                _ = v26
            case "a":
                v2 -= v16
            case "foo bar":
                fmt.Println(v1, v1, v16)
            default:
                var v27 uint16 = 723
                _ = v27
            }
            fmt.Println(v3)
        } else {
            v28 := GenS2{F0: v2}
            v2 = (v16 / ((v2) | 1))
            v2 = (v2 - v2)
            _ = v28
        }
    case "hello":
        var v29 uint = genFn0()
        _ = v29
    }
    v30 := GenS1{F0: 42, F1: 0, F2: v16}
    if (!(false && true)) {
        var v31 int = -592
        _ = v31
    }
    v32 := v4[505]
    v33 := GenS3{F0: (v3 >> ((genFn0()) & 63)), F1: -237}
    if true {
        var v34 uint8 = 255
        var v35 uint8 = 1
        _ = len(v17)
        _ = v34
        _ = v35
    }
    v36 := GenS2{F0: -2147483648}
    v1 = (v1 ^ 860)
    for fi37 := 0; fi37 < 12; fi37++ {
        fmt.Println(v2)
        for fi38 := 0; fi38 < 12; fi38++ {
            switch (fi37 * fi38) {
            case 2147483647:
                v1 = (v1 & (v1 << ((0) & 63)))
            case -9223372036854775808:
                if false {
                    if true {
                        var v39 int64 = 42
                        if true {
                            switch (fi38 << ((982) & 63)) {
                            case 4294967295:
                                v40 := GenS3{F0: (v3 / (((v3 ^ 311)) | 1)), F1: 1}
                                _ = v40
                            }
                            var v41 uint8 = 42
                            if (((v41 * (v41 >> ((18446744073709551615) & 63))) == v41) && (int16(v3) != v1)) {
                                if true {
                                    v42 := GenS1{F0: v1, F1: (fi37 % ((fi38) | 1)), F2: v2}
                                    v1 = (^(v1))
                                    _ = v42
                                } else {
                                    switch v32 {
                                    case "\n\t":
                                        var v43 bool = true
                                        _ = v43
                                    case "a":
                                        var v44 float64 = ((float64((3.14159 * 3.14159)) - 3.14159) * ((-2.25 * 1e10) - 1.5))
                                        _ = v44
                                    case "":
                                        var v45 uint = 255
                                        switch "foo bar" {
                                        case "a":
                                            delete(v4, (fi37 % ((fi38) | 1)))
                                        case "foo bar":
                                            v32 = ("foo bar" + (v32 + v32))
                                        }
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                        delete(v4, (fi38 & (-(fi38))))
                                        if (v41 > 127) {
                                            var v46 float32 = (0.1 - 0.0)
                                            if len(v17) >= 2 {
                                                v17 = v17[1:len(v17)]
                                            }
                                            fmt.Println(v41)
                                            _ = v46
                                        }
                                        var v47 bool = true
                                        _ = v45
                                        _ = v47
                                    default:
                                        var v48 uint = 255
                                        _ = v48
                                    }
                                }
                                var v49 int8 = 127
                                _ = v49
                            } else {
                                var v50 int32 = (-(v2))
                                v51 := GenS2{F0: v16}
                                v41 = v41
                                _ = v50
                                _ = v51
                            }
                            _ = v41
                        }
                        _ = v39
                    }
                } else {
                    switch ((v32 + (("foo bar" + (((v32 + (v32 + v32)) + "foo bar") + (v32 + (v32 + ("hello" + v32))))) + ("hello" + v32))) + v32) {
                    case "a":
                        v16 -= v16
                    case "foo bar":
                        v52 := v4[(fi37 | (fi38 / ((fi37) | 1)))]
                        if (-816 == (v16 & v16)) {
                            v2 -= (-(v2))
                            var v53 uint32 = 42
                            _ = v53
                        }
                        _ = v52
                    default:
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                    if (v3 < 682) {
                        var v54 int8 = 16
                        v55 := GenS1{F0: (v1 >> ((533) & 63)), F1: (fi37 << ((601) & 63)), F2: -2147483648}
                        _ = v54
                        _ = v55
                    }
                }
            default:
                v56 := GenS3{F0: v3, F1: 0}
                v4[fi38] = (v32 + "hello")
                var v57 uint = 460
                _ = v56
                _ = v57
            }
            var v58 int16 = 32767
            _ = fi38
            _ = v58
        }
        v3 -= 2
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v36
    _ = v4
}
