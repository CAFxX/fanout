package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 string }
type GenS4 struct { F0 uint8; F1 string; F2 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 uint64; F1 int32 }

func main() {
    var v1 int64 = 18
    var v2 uint32 = 254
    var v3 int32 = -479
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{fi4, fi4, (fi4 | 2147483647)}
        for fi6 := 0; fi6 < 15; fi6++ {
            _ = len(v5)
            if true {
                v1 += v1
            }
            _ = fi6
        }
        fmt.Println(fi4, v3)
        _ = fi4
        _ = v5
    }
    v7 := []int{}
    v8 := GenS1{F0: v1}
    if false {
        v2 = 65535
        _ = len(v7)
    }
    var v9 bool = (v3 != 42)
    for fi10 := 0; fi10 < 20; fi10++ {
        if (v9 || ((v2 >> ((56) & 63)) < v2)) {
            for fi11 := 0; fi11 < 9; fi11++ {
                if v9 {
                    if ((fi11 ^ (fi10 & fi10)) == 127) {
                        var v12 uint32 = (v2 - (^(v2)))
                        v12 += (v12 - v2)
                        var v13 int64 = (v1 % ((9223372036854775807) | 1))
                        var v14 uint8 = 102
                        v7 = append(v7, fi11)
                        fmt.Println(v14)
                        _ = v12
                        _ = v13
                        _ = v14
                    } else {
                        if (1000 < 1) {
                            if (-1 > (fi10 & fi10)) {
                                v2 = v2
                            }
                            switch ("a" + "foo bar") {
                            case "hello":
                                switch (fi10 ^ -2147483648) {
                                case -2147483648:
                                    var v15 uint = 65535
                                    _ = v15
                                case -1:
                                    if v9 {
                                        var v16 string = (("foo bar" + "hello") + "foo bar")
                                        if (!v9) {
                                            v17 := GenS1{F0: (^(v1))}
                                            _ = v17
                                        }
                                        _ = v16
                                    }
                                case -300:
                                    if v9 {
                                        switch (fi11 >> ((520) & 63)) {
                                        case 0:
                                            if (float64(0.1) > float64(1.5)) {
                                                if v9 {
                                                    var v18 uint32 = v2
                                                    _ = v18
                                                }
                                                v19 := GenS2{F0: 435, F1: 42, F2: 0}
                                                var v20 uint = 1
                                                _ = v19
                                                _ = v20
                                            }
                                        }
                                        var v21 uint32 = (v2 | (v2 % (((v2 - (v2 + (v2 & (v2 / ((v2) | 1)))))) | 1)))
                                        v22 := GenS3{F0: 127, F1: 100, F2: "\n\t"}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                            }
                            var v23 bool = v9
                            _ = v23
                        }
                        if v9 {
                            fmt.Println(fi11, v9, v2)
                        } else {
                            if (true && (-1 != 42)) {
                                fmt.Println(v9)
                                fmt.Println(v9)
                                switch "a" {
                                case "\n\t":
                                    switch "" {
                                    case "\n\t":
                                        v7 = append(v7, fi11)
                                    default:
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                    }
                                case "a":
                                    v9 = (0.0 <= 0.1)
                                }
                                var v24 int64 = (v1 * (v1 | 9223372036854775807))
                                switch 9223372036854775807 {
                                case -9223372036854775808:
                                    var v25 bool = v9
                                    _ = v25
                                case 32767:
                                    v9 = true
                                default:
                                    _ = len(v7)
                                }
                                _ = v24
                            } else {
                                if (!(v9 || (620 > v1))) {
                                    v26 := GenS4{F0: 158, F1: "", F2: v3}
                                    v27 := GenS1{F0: v1}
                                    fmt.Println(v9, v2)
                                    _ = v26
                                    _ = v27
                                } else {
                                    if v9 {
                                        v3 -= (^(v3))
                                        v28 := GenS5{F0: 17}
                                        _ = v28
                                    }
                                    v29 := GenS2{F0: 42, F1: 781, F2: 1}
                                    _ = v29
                                }
                            }
                            switch (^(fi11)) {
                            case -1:
                                switch "" {
                                case "foo bar":
                                    v30 := GenS5{F0: 18446744073709551615}
                                    _ = v30
                                case "a":
                                    if true {
                                        if true {
                                            if (v9 && false) {
                                                var v31 string = "\n\t"
                                                _ = v31
                                            } else {
                                                v1 += 42
                                            }
                                            v32 := GenS1{F0: -307}
                                            v2 = v2
                                            _ = v32
                                        }
                                        fmt.Println(v2)
                                    } else {
                                        v33 := GenS6{F0: (236 > 356)}
                                        if len(v7) > 0 {
                                            v7[0] = int(v2)
                                        }
                                        _ = v33
                                    }
                                case "\n\t":
                                    v34 := GenS2{F0: 18446744073709551615, F1: 1, F2: 1}
                                    _ = v34
                                }
                            case -386:
                                switch (fi10 + int(v1)) {
                                case 9223372036854775807:
                                    v7 = append(v7, 2147483647)
                                case 32767:
                                    if (!false) {
                                        if v9 {
                                            v2 -= v2
                                        }
                                    }
                                case -204:
                                    v35 := GenS7{F0: (65535 > 894), F1: (fi10 ^ (fi10 / ((fi10) | 1)))}
                                    _ = v35
                                }
                            case 954:
                                v1 = v1
                            }
                            switch 150 {
                            case 32767:
                                fmt.Println(v1, fi10)
                            case 42:
                                if false {
                                    var v36 int64 = 563
                                    var v37 uint = 65535
                                    if v9 {
                                        v7 = append(v7, (fi11 * (fi11 ^ fi11)))
                                        if (!((255 < 767) && v9)) {
                                            if v9 {
                                                v38 := GenS8{F0: v2, F1: 851}
                                                v3 = (v3 >> (((v37 * (^(v37)))) & 63))
                                                _ = v38
                                            } else {
                                                fmt.Println(v37, v2)
                                                switch "hello" {
                                                case "foo bar":
                                                    switch ("" + ("" + "\n\t")) {
                                                    case "\n\t":
                                                        v39 := GenS9{F0: v36, F1: v2}
                                                        var v40 float64 = 0.1
                                                        _ = v39
                                                        _ = v40
                                                    case "foo bar":
                                                        if v9 {
                                                            var v41 int8 = 1
                                                            switch "\n\t" {
                                                            case "hello":
                                                                if (!v9) {
                                                                    var v42 int32 = v3
                                                                    var v43 bool = v9
                                                                    switch (fi10 * 32767) {
                                                                    case 260:
                                                                        if false {
                                                                            v44 := GenS10{F0: uint16(v2)}
                                                                            if len(v7) >= 2 {
                                                                                v7 = v7[1:len(v7)]
                                                                            }
                                                                            _ = v44
                                                                        }
                                                                    case 661:
                                                                        v45 := GenS11{F0: v37}
                                                                        _ = v45
                                                                    case 484:
                                                                        var v46 float64 = float64(-2.25)
                                                                        _ = v46
                                                                    }
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            }
                                                            fmt.Println(v36, v9)
                                                            _ = v41
                                                        }
                                                        v37 += v37
                                                    }
                                                }
                                                var v47 float64 = (float64(1.5) / ((((((-2.25 + 1.5) - 1.5) / 3.14159) / 1.0) / (float64((0.0 * -2.25)) + ((1e10 / 1e10) / (float64(1e10) - 1.5)))) * 1e10))
                                                v2 = 2147483647
                                                _ = v47
                                            }
                                        }
                                        fmt.Println(v37)
                                    }
                                    switch "a" {
                                    case "foo bar":
                                        var v48 int32 = (v3 / (((v3 | (v3 % ((-1) | 1)))) | 1))
                                        _ = v48
                                    }
                                    _ = v36
                                    _ = v37
                                } else {
                                    if false {
                                        v3 = v3
                                        v3 = v3
                                        if (!((v2 % ((2) | 1)) != 2)) {
                                            v2 += v2
                                            fmt.Println(v2, v9)
                                        } else {
                                            switch "foo bar" {
                                            case "hello":
                                                var v49 float32 = 0.0
                                                _ = v49
                                            case "\n\t":
                                                switch 790 {
                                                case 2147483647:
                                                    var v50 float32 = 0.1
                                                    _ = v50
                                                case 832:
                                                    v3 = v3
                                                case -1:
                                                    fmt.Println(fi10, v2, fi10)
                                                default:
                                                    var v51 float64 = (-2.25 / 3.14159)
                                                    _ = v51
                                                }
                                            case "a":
                                                v3 = (v3 - (v3 & (v3 % ((v3) | 1))))
                                            default:
                                                v7 = append(v7, (fi11 + (fi10 & fi11)))
                                            }
                                        }
                                    } else {
                                        v7 = append(v7, 128)
                                        var v52 int32 = v3
                                        var v53 uint64 = 49
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 uint16 = 255
                                    _ = v54
                                }
                            default:
                                var v55 string = (("\n\t" + ("\n\t" + "\n\t")) + "\n\t")
                                _ = v55
                            }
                        }
                        var v56 float64 = -2.25
                        var v57 int16 = 827
                        v1 = (v1 / (((-(v1))) | 1))
                        _ = v56
                        _ = v57
                    }
                    var v58 uint32 = (v2 ^ v2)
                    v59 := GenS2{F0: 334, F1: 42, F2: 0}
                    _ = v58
                    _ = v59
                }
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                v3 += 2147483647
                _ = fi11
            }
        } else {
            var v60 float64 = (-2.25 - (1.5 + (0.1 / 1.0)))
            v3 = v3
            v61 := GenS12{F0: 42, F1: 82}
            _ = v60
            _ = v61
        }
        _ = fi10
    }
    v2 += v2
    v3 = v3
    v9 = (!false)
    v3 = v3
    v62 := GenS13{F0: 18446744073709551615, F1: 42}
    v1 = 1
    var v63 uint = 345
    v3 = (v3 * 0)
    v9 = (uint64(v63) <= 443)
    if len(v7) > 0 {
        v7[0] = 32768
    }
    v64 := GenS2{F0: 1, F1: 159, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v63
    _ = v64
    _ = v7
    _ = v8
    _ = v9
}
