package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS5 struct { F0 int8; F1 bool }
func genFn0() int16 {
    var v7 string = "hello"
    v8 := []int{}
    v8 = append(v8, 42)
    v9 := GenS1{F0: (v7 + ("a" + "\n\t"))}
    _ = v7
    _ = v8
    _ = v9
    return -231
}

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    var v4 uint32 = 474
    v5 := []int{-1, -912, -2147483648, 1}
    v6 := make(map[string]string)
    var v10 uint = 255
    var v11 uint64 = 844
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint64 = (^(v11))
        for fi14 := 0; fi14 < 19; fi14++ {
            switch (fi12 % ((fi14) | 1)) {
            case 42:
                var v15 int = fi12
                _ = v15
            case 19:
                v16 := GenS1{F0: ""}
                _ = v16
            }
            v17 := GenS1{F0: ("foo bar" + "foo bar")}
            var v18 uint8 = 0
            switch (fi14 * (fi12 - -1)) {
            case 106:
                if false {
                    var v19 float32 = (3.14159 * (1.5 + (0.0 + float32(0.0))))
                    switch fi12 {
                    case 868:
                        var v20 float64 = (1.5 / 1.0)
                        _ = v20
                    }
                    _ = v19
                } else {
                    switch (("foo bar" + ("foo bar" + "\n\t")) + "a") {
                    case "foo bar":
                        var v21 bool = true
                        _ = v21
                    case "a":
                        var v22 uint8 = (v18 / ((2) | 1))
                        _ = v22
                    case "":
                        var v23 uint32 = (v4 ^ v4)
                        _ = v23
                    default:
                        var v24 float32 = 0.1
                        _ = v24
                    }
                }
            default:
                var v25 uint64 = v13
                _ = v25
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi14
            _ = v17
            _ = v18
        }
        var v26 int32 = int32(v10)
        _ = fi12
        _ = v13
        _ = v26
    }
    _ = len(v2)
    fmt.Println(v11, v10, v10)
    delete(v2, v4)
    var v27 bool = (!false)
    var v28 int = 128
    if (255 < v10) {
        if v27 {
            var v29 uint16 = 255
            delete(v2, (v4 - (v4 / ((321) | 1))))
            _ = v29
        }
        switch v28 {
        case 225:
            if false {
                var v30 bool = v27
                for fi31 := 0; fi31 < 10; fi31++ {
                    if (2147483647 <= (v4 & (v4 / ((2) | 1)))) {
                        var v32 int8 = (v1 >> ((v10) & 63))
                        var v33 int32 = 0
                        var v34 float64 = 1.5
                        _ = v32
                        _ = v33
                        _ = v34
                    } else {
                        fmt.Println(v4, v11, fi31)
                        _ = len(v5)
                        v28 += 1
                    }
                    var v35 uint32 = (^(v4))
                    for fi36 := 0; fi36 < 6; fi36++ {
                        v4 = v4
                        _ = fi36
                    }
                    _ = fi31
                    _ = v35
                }
                _ = v30
            } else {
                v37 := v6[(("foo bar" + "") + "hello")]
                switch "\n\t" {
                case "foo bar":
                    switch v37 {
                    case "hello":
                        if (0.1 != 0.1) {
                            v38 := GenS1{F0: v37}
                            fmt.Println(v1)
                            for fi39 := 0; fi39 < 2; fi39++ {
                                var v40 int32 = 42
                                fmt.Println(v27, fi39, v4)
                                _ = fi39
                                _ = v40
                            }
                            _ = v38
                        } else {
                            for fi41 := 0; fi41 < 20; fi41++ {
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v27 = v27
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = fi42
                                }
                                for fi43 := 0; fi43 < 19; fi43++ {
                                    v44 := v2[v4]
                                    fmt.Println(fi43, v4)
                                    _ = fi43
                                    _ = v44
                                }
                                var v45 float64 = 1e10
                                _ = fi41
                                _ = v45
                            }
                        }
                    case "foo bar":
                        v10 = v10
                    default:
                        for fi46 := 0; fi46 < 17; fi46++ {
                            var v47 int32 = 2147483647
                            _ = fi46
                            _ = v47
                        }
                    }
                case "hello":
                    var v48 float32 = 0.1
                    _ = v48
                default:
                    if (((!(v27 || (!(!v27)))) && (!v27)) || (v11 <= v11)) {
                        v4 = 827
                    } else {
                        _ = len(v6)
                    }
                }
                v10 -= (v10 * (^(v10)))
                _ = v37
            }
        case -872:
            v49 := GenS2{F0: 422}
            _ = v49
        case 580:
            v11 = (v11 | uint64(v10))
        default:
            v50 := GenS1{F0: "\n\t"}
            _ = v50
        }
    }
    var v51 int = (^(v28))
    v28 = v51
    var v52 uint16 = 228
    switch "foo bar" {
    case "":
        if (-1 < -2147483648) {
            v4 = v4
            if v27 {
                var v53 int32 = 1
                _ = v53
            } else {
                switch v28 {
                case 128:
                    fmt.Println(v52, v4)
                case 2147483647:
                    if (((0.0 / 1e10) - 1.5) != 1e10) {
                        v4 = v4
                        v54 := GenS2{F0: 42}
                        for fi55 := 0; fi55 < 18; fi55++ {
                            var v56 float64 = 1e10
                            delete(v2, v4)
                            fmt.Println(fi55, v10)
                            _ = fi55
                            _ = v56
                        }
                        _ = v54
                    } else {
                        v5 = append(v5, (v51 + v51))
                    }
                }
                for fi57 := 0; fi57 < 11; fi57++ {
                    var v58 string = ("hello" + ("\n\t" + "a"))
                    switch (v58 + "foo bar") {
                    case "":
                        switch v58 {
                        case "":
                            for fi59 := 0; fi59 < 10; fi59++ {
                                v60 := GenS2{F0: 9223372036854775807}
                                switch v28 {
                                case 2147483647:
                                    v52 -= 238
                                default:
                                    if (!(v58 > v58)) {
                                        v11 -= (v11 & v11)
                                        v4 += (^(v4))
                                        delete(v2, (v4 % (((v4 | 0)) | 1)))
                                        v61 := v3[v58]
                                        _ = v61
                                    } else {
                                        if (!((v52 % ((v52) | 1)) == (v52 - v52))) {
                                            fmt.Println(v58, v4, v10)
                                            var v62 int = -365
                                            var v63 uint32 = uint32(v52)
                                            var v64 int64 = -1
                                            v6[v58] = v58
                                            var v65 int16 = genFn0()
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                        var v66 int16 = 42
                                        v67 := GenS2{F0: -713}
                                        _ = v66
                                        _ = v67
                                    }
                                }
                                var v68 bool = false
                                switch v28 {
                                case 2147483647:
                                    v58 = v58
                                }
                                _ = fi59
                                _ = v60
                                _ = v68
                            }
                        }
                    case "\n\t":
                        var v69 bool = ((v1 % ((v1) | 1)) != v1)
                        _ = v69
                    case "a":
                        switch ("foo bar" + "") {
                        case "\n\t":
                            for fi70 := 0; fi70 < 8; fi70++ {
                                fmt.Println(v1, v58, fi57)
                                var v71 float64 = (0.1 / 1.5)
                                _ = fi70
                                _ = v71
                            }
                        case "":
                            for fi72 := 0; fi72 < 12; fi72++ {
                                switch "a" {
                                case "foo bar":
                                    v5 = append(v5, 774)
                                case "":
                                    fmt.Println(v11, v27, v11)
                                case "a":
                                    v51 -= (v28 * fi72)
                                default:
                                    v51 -= (v28 << (((v10 << (((v10 & v10)) & 63))) & 63))
                                }
                                v73 := GenS3{F0: v1, F1: v1, F2: 831}
                                fmt.Println(v1, v58)
                                var v74 bool = v27
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                        case "foo bar":
                            if (((420 > -9223372036854775808) && ((float64(0.1) * 1e10) < 1.5)) && true) {
                                v1 += (-(v1))
                            }
                        }
                    }
                    for fi75 := 0; fi75 < 19; fi75++ {
                        if false {
                            v10 = v10
                            var v76 uint64 = (v11 * v11)
                            _ = v76
                        }
                        var v77 uint64 = (v11 % (((v11 | 777)) | 1))
                        v4 = (^(v4))
                        _ = fi75
                        _ = v77
                    }
                    _ = fi57
                    _ = v58
                }
            }
            v78 := GenS2{F0: int64(v1)}
            _ = v78
        }
    default:
        if v27 {
            for fi79 := 0; fi79 < 10; fi79++ {
                v10 = v10
                _ = fi79
            }
            switch v51 {
            case 9223372036854775807:
                var v80 int16 = 1
                _ = v80
            default:
                v81 := v6[""]
                _ = v81
            }
        } else {
            v82 := GenS4{F0: (v52 / ((v52) | 1)), F1: v4, F2: v1}
            fmt.Println(v1, v1, v51)
            _ = v82
        }
    }
    v83 := GenS5{F0: (^(v1)), F1: v27}
    v51 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
    _ = v83
}
