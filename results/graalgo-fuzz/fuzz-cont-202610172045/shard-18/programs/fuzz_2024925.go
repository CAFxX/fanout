package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint; F1 uint64 }
type GenS9 struct { F0 int; F1 uint; F2 int64 }
type GenS10 struct { F0 int16; F1 uint16 }
func genFn0() uint {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: -850, F1: 0}
        var v6 uint16 = 32767
        _ = fi4
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS2{F0: 0, F1: 1}
        var v8 float64 = (1e10 * 1e10)
        for fi9 := 0; fi9 < 9; fi9++ {
            for fi10 := 0; fi10 < 14; fi10++ {
                if false {
                    switch (fi10 % ((fi9) | 1)) {
                    case 805:
                        v11 := []int{}
                        _ = v11
                    case 590:
                        if true {
                            var v12 float32 = 1.5
                            _ = v12
                        }
                    case -594:
                        v13 := GenS2{F0: 1, F1: 42}
                        _ = v13
                    default:
                        v8 = v8
                    }
                    v8 = (float64((0.1 / (1.5 / 1.5))) - v8)
                } else {
                    fmt.Println(fi10)
                    switch 127 {
                    case 24:
                        fmt.Println(fi9)
                    case -9223372036854775808:
                        v14 := []int{}
                        _ = v14
                    case 32768:
                        var v15 uint32 = 2147483647
                        _ = v15
                    }
                    v16 := GenS2{F0: 710, F1: 2147483647}
                    _ = v16
                }
                _ = fi10
            }
            var v17 float32 = float32(1e10)
            var v18 int8 = 90
            _ = fi9
            _ = v17
            _ = v18
        }
        _ = v7
        _ = v8
    } else {
        var v19 int = 127
        v20 := GenS3{F0: 750}
        var v21 string = "hello"
        _ = v19
        _ = v20
        _ = v21
    }
    var v22 uint8 = 1
    for fi23 := 0; fi23 < 19; fi23++ {
        v22 = 0
        if false {
            var v24 uint8 = v22
            switch "" {
            case "\n\t":
                v25 := []int{}
                _ = v25
            case "hello":
                fmt.Println(v24, v24, v22)
            case "":
                v24 += (v22 / ((v24) | 1))
            }
            v26 := []int{9223372036854775807, (fi23 << ((uint(v24)) & 63))}
            _ = v24
            _ = v26
        } else {
            v22 = v22
            fmt.Println(v22)
        }
        if false {
            v27 := GenS4{F0: -128}
            v28 := GenS1{F0: -1, F1: 474}
            _ = v27
            _ = v28
        }
        _ = fi23
    }
    _ = v22
    return 2
}
func genFn1(p0 int, p1 int8) int16 {
    v29 := []int{p0}
    v30 := GenS4{F0: p1}
    fmt.Println(p1, p0)
    p0 = p0
    _ = p0
    _ = p1
    _ = v29
    _ = v30
    return -101
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2
    v3 := make(map[string]int)
    if true {
        fmt.Println(v2)
        switch 127 {
        case -300:
            v2 += 2
        }
    }
    if false {
        for fi31 := 0; fi31 < 10; fi31++ {
            v32 := []int{}
            for fi33 := 0; fi33 < 20; fi33++ {
                v2 -= (^(v2))
                _ = fi33
            }
            _ = fi31
            _ = v32
        }
    } else {
        var v34 float32 = ((0.1 * -2.25) - 1e10)
        v2 = (^(v2))
        if false {
            if false {
                var v35 uint = 2
                v36 := GenS2{F0: 174, F1: 42}
                _ = v35
                _ = v36
            } else {
                switch 32767 {
                case 4294967295:
                    _ = len(v3)
                default:
                    if true {
                        v37 := []int{579, 32768, -2147483648}
                        v2 = (^(v2))
                        _ = v37
                    }
                }
            }
            v34 = float32(0.0)
            v38 := []int{}
            _ = v38
        } else {
            var v39 uint64 = 0
            fmt.Println(v39, v39, v39)
            _ = v39
        }
        _ = v34
    }
    var v40 string = ("\n\t" + ("hello" + ("hello" + ("" + "foo bar"))))
    fmt.Println(v40, v2)
    v41 := GenS5{F0: -432, F1: 4}
    v2 = 4294967295
    switch "\n\t" {
    case "hello":
        for fi42 := 0; fi42 < 13; fi42++ {
            var v43 int64 = -1
            v44 := []int{32768, fi42, (fi42 | fi42), int(v43)}
            _ = fi42
            _ = v43
            _ = v44
        }
    case "":
        var v45 float64 = -2.25
        _ = v45
    case "\n\t":
        var v46 uint8 = 2
        _ = v46
    }
    delete(v3, v40)
    for fi47 := 0; fi47 < 2; fi47++ {
        fmt.Println(v2, fi47)
        var v48 string = ("" + "foo bar")
        fmt.Println(v48, fi47)
        _ = fi47
        _ = v48
    }
    v2 = v2
    switch v40 {
    case "foo bar":
        for fi49 := 0; fi49 < 8; fi49++ {
            var v50 string = v40
            _ = fi49
            _ = v50
        }
    case "":
        if true {
            for fi51 := 0; fi51 < 1; fi51++ {
                for fi52 := 0; fi52 < 12; fi52++ {
                    v1[(fi51 * (fi52 + fi52))] = 4294967295
                    _ = fi52
                }
                fmt.Println(fi51, v2)
                for fi53 := 0; fi53 < 16; fi53++ {
                    if false {
                        switch ("foo bar" + (v40 + v40)) {
                        case "":
                            v54 := []int{2147483647, fi51}
                            _ = v54
                        case "\n\t":
                            var v55 int8 = 19
                            _ = v55
                        case "foo bar":
                            v40 = (v40 + (("\n\t" + "\n\t") + ((v40 + v40) + "\n\t")))
                        default:
                            v40 = v40
                        }
                        v40 = v40
                    } else {
                        if false {
                            switch fi51 {
                            case 49:
                                _ = len(v3)
                            case 9223372036854775807:
                                v56 := []int{fi51, -674, fi53}
                                _ = v56
                            }
                            switch (fi51 - (fi53 - fi51)) {
                            case -9223372036854775808:
                                _ = len(v1)
                            case 42:
                                var v57 uint = 1
                                v2 = 65535
                                _ = v57
                            case -398:
                                if (true || (1000 < 1)) {
                                    var v58 int32 = -1
                                    if true {
                                        v59 := GenS2{F0: 255, F1: -1}
                                        v2 = (v2 + 0)
                                        var v60 int64 = 388
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 float64 = 1e10
                                        _ = v61
                                    }
                                    v62 := GenS1{F0: 0, F1: 18446744073709551615}
                                    _ = v58
                                    _ = v62
                                } else {
                                    var v63 float32 = ((3.14159 * 0.1) * 3.14159)
                                    v40 = v40
                                    if true {
                                        v64 := GenS6{F0: -82, F1: 127}
                                        _ = v64
                                    }
                                    v63 = v63
                                    v63 -= v63
                                    v65 := v1[(^(fi51))]
                                    _ = v63
                                    _ = v65
                                }
                            }
                        }
                        fmt.Println(fi53, fi51, v2)
                        v66 := GenS5{F0: (fi51 ^ 32768), F1: 208}
                        _ = v66
                    }
                    _ = fi53
                }
                _ = fi51
            }
            if true {
                v2 = 2147483648
            } else {
                v67 := GenS3{F0: 9223372036854775807}
                for fi68 := 0; fi68 < 10; fi68++ {
                    v69 := GenS7{F0: -32768, F1: 255, F2: 255}
                    v70 := []int{fi68, fi68, 2147483647, -413}
                    _ = fi68
                    _ = v69
                    _ = v70
                }
                _ = v67
            }
        } else {
            for fi71 := 0; fi71 < 11; fi71++ {
                fmt.Println(fi71)
                _ = len(v1)
                switch 127 {
                case -152:
                    if ((false && true) && false) {
                        for fi72 := 0; fi72 < 14; fi72++ {
                            var v73 uint16 = 2
                            _ = fi72
                            _ = v73
                        }
                    }
                case -461:
                    if true {
                        v40 = v40
                        v40 = v40
                        for fi74 := 0; fi74 < 9; fi74++ {
                            v2 += (v2 ^ (v2 - (v2 / (((^(v2))) | 1))))
                            _ = fi74
                        }
                    } else {
                        v75 := GenS3{F0: 893}
                        _ = v75
                    }
                case 9223372036854775807:
                    v76 := GenS8{F0: 0, F1: 615}
                    _ = v76
                }
                _ = fi71
            }
            for fi77 := 0; fi77 < 19; fi77++ {
                var v78 int8 = -128
                v40 = ("hello" + v40)
                v1[(fi77 & fi77)] = fi77
                _ = fi77
                _ = v78
            }
        }
    default:
        switch 2147483647 {
        case 665:
            if true {
                v79 := []int{-67, 9223372036854775807, -189, 1}
                _ = v79
            } else {
                v80 := GenS9{F0: -9223372036854775808, F1: 1, F2: -9223372036854775808}
                _ = v80
            }
        case 0:
            var v81 uint = 489
            _ = v81
        }
    }
    v82 := GenS10{F0: 0, F1: 32767}
    v40 = v40
    if true {
        var v83 int8 = -128
        _ = v83
    }
    for fi84 := 0; fi84 < 15; fi84++ {
        v40 = (((v40 + v40) + v40) + (("\n\t" + (("\n\t" + (v40 + "foo bar")) + v40)) + (("" + v40) + v40)))
        _ = fi84
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v82
}
