package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 uint; F2 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int64; F2 uint16 }

func main() {
    var v1 float32 = -2.25
    v2 := []uint32{1, 179}
    var v3 uint = 2
    var v4 int32 = 42
    var v5 int32 = v4
    v4 = 309
    for fi6 := 0; fi6 < 19; fi6++ {
        v4 = 42
        _ = fi6
    }
    var v7 int64 = 17
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch ("" + (("a" + "hello") + ("" + "foo bar"))) {
    case "":
        v5 -= (v4 % ((-1) | 1))
    case "foo bar":
        for fi8 := 0; fi8 < 10; fi8++ {
            fmt.Println(v3, v5, v4)
            _ = fi8
        }
    case "\n\t":
        var v9 uint16 = 976
        _ = v9
    }
    v4 += (v5 - (v5 | v5))
    var v10 float32 = v1
    fmt.Println(v3)
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            v5 += v5
            for fi12 := 0; fi12 < 16; fi12++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v13 string = "\n\t"
                if (244 < 25) {
                    var v14 uint16 = 42
                    _ = v14
                } else {
                    v3 = v3
                }
                _ = fi12
                _ = v13
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 float32 = v10
                v10 += 0.1
                _ = fi15
                _ = v16
            }
            _ = fi11
        }
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            switch -1 {
            case 644:
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v7, fi17)
                    v19 := GenS1{F0: 191}
                    v3 += (v3 * 0)
                    _ = fi18
                    _ = v19
                }
            case 4294967295:
                if (-475 >= -1) {
                    if (65535 <= 511) {
                        var v20 uint16 = 1
                        var v21 int16 = 128
                        _ = v20
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v2 = append(v2, 0)
                            _ = fi22
                        }
                        for fi23 := 0; fi23 < 6; fi23++ {
                            v7 = 9223372036854775807
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi23
                        }
                        if false {
                            var v24 uint = 1
                            v25 := GenS1{F0: 560}
                            _ = v24
                            _ = v25
                        } else {
                            switch "hello" {
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v26 uint = 18446744073709551615
                                _ = v26
                            }
                            if (((v1 + v10) * (1.5 - v10)) != (0.0 * (v10 / (-2.25 * 1.5)))) {
                                var v27 bool = ((0.0 >= -2.25) || false)
                                _ = v27
                            }
                            var v28 int8 = 42
                            _ = v28
                        }
                    }
                    v5 += (v5 % ((-582) | 1))
                    v29 := GenS2{F0: 2, F1: 1, F2: fi17}
                    _ = v29
                } else {
                    var v30 float32 = (v1 + float32((-2.25 + (3.14159 - (1.5 + 1.5)))))
                    for fi31 := 0; fi31 < 7; fi31++ {
                        var v32 string = "hello"
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                }
            case 42:
                var v33 int16 = -1
                _ = v33
            }
            _ = fi17
        }
    }
    if false {
        var v34 uint16 = uint16(v3)
        for fi35 := 0; fi35 < 2; fi35++ {
            switch "a" {
            case "":
                v36 := GenS3{F0: -1, F1: (v3 & 643), F2: 463}
                _ = v36
            case "hello":
                switch fi35 {
                case 4294967295:
                    v37 := GenS2{F0: 1, F1: 0, F2: (fi35 - fi35)}
                    _ = v37
                default:
                    var v38 uint16 = 0
                    _ = v38
                }
            case "foo bar":
                switch (fi35 * fi35) {
                case 4294967295:
                    var v39 int = fi35
                    _ = v39
                case 373:
                    var v40 string = "\n\t"
                    _ = v40
                case 761:
                    if (0 == -1) {
                        var v41 float64 = 0.1
                        switch "" {
                        case "hello":
                            var v42 uint = (v3 - v3)
                            _ = v42
                        case "a":
                            var v43 string = ("a" + "hello")
                            _ = v43
                        case "":
                            if false {
                                v4 = (^(v4))
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    var v45 string = (("foo bar" + "") + "")
                                    fmt.Println(v1, v4)
                                    _ = fi44
                                    _ = v45
                                }
                            } else {
                                switch ("foo bar" + ("" + ("\n\t" + "hello"))) {
                                case "a":
                                    v4 = (v5 << ((206) & 63))
                                }
                                if len(v2) > 0 {
                                    v2[0] = 223
                                }
                                var v46 uint16 = (^(v34))
                                _ = v46
                            }
                        default:
                            for fi47 := 0; fi47 < 14; fi47++ {
                                var v48 float32 = -2.25
                                v3 -= v3
                                _ = fi47
                                _ = v48
                            }
                        }
                        for fi49 := 0; fi49 < 17; fi49++ {
                            if true {
                                v50 := GenS4{F0: v7, F1: 42, F2: "foo bar"}
                                _ = v50
                            } else {
                                if true {
                                    switch (((("\n\t" + "") + ((("a" + "\n\t") + (("a" + "a") + "foo bar")) + "hello")) + ("foo bar" + "a")) + "\n\t") {
                                    case "a":
                                        var v51 uint = (v3 * 65535)
                                        _ = v51
                                    case "foo bar":
                                        v5 = -1
                                    default:
                                        var v52 int8 = 0
                                        _ = v52
                                    }
                                } else {
                                    var v53 int16 = 32767
                                    fmt.Println(v1, v34)
                                    if ("\n\t" <= "\n\t") {
                                        var v54 uint16 = (^(v34))
                                        v55 := GenS3{F0: v53, F1: (v3 + 18446744073709551615), F2: 301}
                                        var v56 string = ("hello" + "\n\t")
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v53
                                }
                                v57 := GenS5{F0: int32(v3)}
                                var v58 uint = v3
                                _ = v57
                                _ = v58
                            }
                            v59 := GenS6{F0: 226, F1: v34}
                            _ = fi49
                            _ = v59
                        }
                        _ = v41
                    }
                }
            }
            if len(v2) > 0 {
                v2[0] = uint32(v3)
            }
            v60 := GenS7{F0: v7}
            _ = fi35
            _ = v60
        }
        _ = len(v2)
        _ = v34
    } else {
        v5 = v4
    }
    fmt.Println(v5)
    for fi61 := 0; fi61 < 8; fi61++ {
        _ = len(v2)
        if (true || (false || false)) {
            v62 := GenS8{F0: 65535, F1: v7, F2: 92}
            _ = v62
        }
        _ = fi61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
