package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 42
    v3 := make(map[uint32]int)
    var v4 int8 = 127
    var v5 uint = 740
    fmt.Println(v2)
    var v6 uint32 = 214
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int64 = -1
        switch "foo bar" {
        case "a":
            var v9 uint8 = 176
            _ = v9
        default:
            switch ("foo bar" + "\n\t") {
            case "":
                fmt.Println(v5, v2)
            case "hello":
                var v10 uint16 = 32767
                _ = v10
            case "foo bar":
                if ((42 <= 18446744073709551615) || false) {
                    if false {
                        switch (fi7 - fi7) {
                        case 4294967295:
                            switch ("hello" + (("hello" + "hello") + "foo bar")) {
                            case "a":
                                v11 := GenS1{F0: 65535, F1: 1, F2: 2}
                                _ = v11
                            case "hello":
                                var v12 int16 = 42
                                _ = v12
                            default:
                                var v13 uint = (^(v5))
                                _ = v13
                            }
                        case 1:
                            for fi14 := 0; fi14 < 1; fi14++ {
                                delete(v3, (^(v6)))
                                if false {
                                    v2 -= v8
                                    switch "foo bar" {
                                    case "\n\t":
                                        v15 := GenS1{F0: 2147483648, F1: -961, F2: 2}
                                        _ = v15
                                    case "a":
                                        v6 -= v6
                                    case "":
                                        v8 += v8
                                    default:
                                        fmt.Println(fi7, v2)
                                    }
                                    v16 := []int{2147483647, fi14}
                                    _ = v16
                                }
                                switch "" {
                                case "":
                                    v17 := GenS2{F0: 1, F1: (fi7 | (fi14 * 32767)), F2: 32767}
                                    var v18 int8 = (v4 >> (((v5 ^ v5)) & 63))
                                    var v19 int16 = 617
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                default:
                                    fmt.Println(v6, v5, fi14)
                                }
                                v3[v6] = fi7
                                _ = fi14
                            }
                        case 2147483647:
                            v20 := GenS3{F0: (v5 - (v5 % ((v5) | 1))), F1: (v5 & (v5 << ((v5) & 63))), F2: 2147483648}
                            _ = v20
                        default:
                            switch (fi7 - fi7) {
                            case -717:
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    v6 = v6
                                    if false {
                                        v5 = (v5 | v5)
                                        v22 := v3[(v6 | v6)]
                                        if (v2 == 9223372036854775807) {
                                            if false {
                                                v23 := GenS3{F0: 631, F1: v5, F2: (^(v6))}
                                                _ = v23
                                            } else {
                                                v24 := GenS3{F0: (v5 % (((v5 * v5)) | 1)), F1: 1, F2: (v6 & (v6 ^ v6))}
                                                v5 += (v5 >> ((2) & 63))
                                                _ = v24
                                            }
                                            var v25 string = ("" + ("foo bar" + "\n\t"))
                                            var v26 uint16 = 1
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = v22
                                    }
                                    if false {
                                        v4 = v4
                                        switch -1 {
                                        case 473:
                                            v8 = v8
                                        case -2147483648:
                                            v27 := GenS4{F0: 255, F1: -11}
                                            _ = v27
                                        case 668:
                                            v8 = v2
                                        }
                                    } else {
                                        v28 := []int{fi7, 128, (^(fi7))}
                                        v29 := GenS3{F0: 2, F1: (v5 * v5), F2: (v6 / ((42) | 1))}
                                        _ = v28
                                        _ = v29
                                    }
                                    if false {
                                        if false {
                                            var v30 float64 = 1.5
                                            _ = v30
                                        } else {
                                            v6 = (v6 << (((v5 % (((v5 % ((v5) | 1))) | 1))) & 63))
                                            v8 -= (v8 + (v2 >> (((v5 + 65535)) & 63)))
                                        }
                                        v31 := GenS5{F0: (^(v5))}
                                        v32 := GenS3{F0: 694, F1: v5, F2: (v6 ^ 41)}
                                        _ = v31
                                        _ = v32
                                    }
                                    _ = fi21
                                }
                            case 9223372036854775807:
                                v1[(((("hello" + (("\n\t" + ("a" + "\n\t")) + ("foo bar" + (("" + ((("foo bar" + "\n\t") + "foo bar") + "a")) + ("\n\t" + "hello"))))) + "\n\t") + "\n\t") + "foo bar")] = true
                            }
                        }
                    } else {
                        v33 := GenS3{F0: v5, F1: (v5 + (v5 + (v5 ^ v5))), F2: 65535}
                        _ = v33
                    }
                } else {
                    var v34 float64 = 0.0
                    var v35 uint32 = 1
                    _ = len(v1)
                    _ = v34
                    _ = v35
                }
            }
        }
        _ = len(v3)
        _ = fi7
        _ = v8
    }
    switch 542 {
    case -2147483648:
        var v36 bool = true
        _ = v36
    case 1:
        v37 := GenS5{F0: v5}
        _ = v37
    case 42:
        for fi38 := 0; fi38 < 18; fi38++ {
            var v39 int64 = -1
            for fi40 := 0; fi40 < 19; fi40++ {
                if ((-(fi40)) >= (fi40 % ((fi38) | 1))) {
                    v41 := GenS6{F0: 255}
                    var v42 uint8 = 147
                    v43 := GenS5{F0: 65535}
                    _ = v41
                    _ = v42
                    _ = v43
                } else {
                    v1[("\n\t" + "foo bar")] = false
                    v2 = (v2 + (v2 << (((v5 % ((v5) | 1))) & 63)))
                    fmt.Println(v5, fi38, v2)
                    v6 -= (v6 - 2147483647)
                    fmt.Println(fi40)
                }
                v39 += (^(v39))
                var v44 string = "foo bar"
                _ = fi40
                _ = v44
            }
            v45 := []int{(-(fi38)), 32768, fi38}
            _ = fi38
            _ = v39
            _ = v45
        }
    }
    switch 127 {
    case 4294967295:
        if false {
            switch ("a" + "a") {
            case "hello":
                if false {
                    switch "" {
                    case "foo bar":
                        switch ("" + "hello") {
                        case "a":
                            v2 += v2
                        case "\n\t":
                            v4 = (^(v4))
                        case "hello":
                            var v46 uint64 = 1
                            _ = v46
                        }
                    default:
                        v6 = (v6 / (((v6 / ((v6) | 1))) | 1))
                    }
                    v47 := GenS7{F0: -1}
                    v48 := []int{int(v2), 32768, 42, 32768}
                    _ = v47
                    _ = v48
                } else {
                    switch "" {
                    case "":
                        var v49 uint64 = 1
                        _ = v49
                    case "a":
                        v5 = (v5 / (((v5 % (((v5 << ((v5) & 63))) | 1))) | 1))
                    }
                    var v50 uint8 = 222
                    v51 := []int{1, -2147483648, 503}
                    _ = v50
                    _ = v51
                }
            case "\n\t":
                v2 -= v2
            default:
                if (("" + ("hello" + "")) > ("foo bar" + "")) {
                    fmt.Println(v2, v6, v5)
                    if false {
                        v2 = v2
                        v1[""] = true
                        _ = len(v3)
                    } else {
                        fmt.Println(v2)
                        if false {
                            for fi52 := 0; fi52 < 12; fi52++ {
                                v5 = (v5 ^ v5)
                                v53 := GenS7{F0: 1}
                                v54 := []int{}
                                _ = fi52
                                _ = v53
                                _ = v54
                            }
                            v4 = v4
                        }
                        if (2 == v5) {
                            switch "foo bar" {
                            case "a":
                                switch -352 {
                                case -2147483648:
                                    v55 := GenS6{F0: 1}
                                    _ = v55
                                case 2147483647:
                                    v6 -= 631
                                default:
                                    for fi56 := 0; fi56 < 3; fi56++ {
                                        for fi57 := 0; fi57 < 20; fi57++ {
                                            var v58 int8 = (v4 >> (((v5 | 18446744073709551615)) & 63))
                                            _ = fi57
                                            _ = v58
                                        }
                                        _ = fi56
                                    }
                                }
                            }
                        }
                    }
                    v4 -= v4
                }
            }
            fmt.Println(v4, v2)
        }
    case 32767:
        switch 32767 {
        case -1:
            for fi59 := 0; fi59 < 5; fi59++ {
                if true {
                    for fi60 := 0; fi60 < 4; fi60++ {
                        v61 := GenS8{F0: ((357 < 42) && (-2.25 == (1.5 - 0.1)))}
                        var v62 uint8 = 76
                        if false {
                            fmt.Println(v5)
                            switch int(v2) {
                            case 32767:
                                v2 += v2
                            }
                        }
                        _ = fi60
                        _ = v61
                        _ = v62
                    }
                    v6 += (v6 / ((907) | 1))
                }
                v5 -= 255
                v5 += 1
                _ = fi59
            }
        default:
            var v63 uint32 = v6
            _ = v63
        }
    case 2147483647:
        var v64 uint64 = 531
        _ = v64
    default:
        v65 := []int{-939, 970, 127, 4294967295}
        _ = v65
    }
    delete(v3, (v6 / (((v6 << ((uint(v6)) & 63))) | 1)))
    _ = len(v3)
    v4 -= -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
