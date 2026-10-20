package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int8; F1 uint; F2 int64 }

func main() {
    var v1 uint = 2
    var v2 int = 131
    var v3 int64 = 0
    v4 := make(map[int]string)
    var v5 int16 = 0
    v6 := make(map[uint32]uint32)
    var v7 uint32 = 585
    v8 := []int{}
    var v9 uint64 = 1
    fmt.Println(v3)
    fmt.Println(v7, v5, v2)
    v10 := v6[v7]
    var v11 float64 = (1e10 + 1e10)
    v10 -= 65535
    v3 -= v3
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 float64 = 3.14159
        if true {
            v14 := GenS1{F0: v1}
            var v15 uint = (v1 & (^(v1)))
            for fi16 := 0; fi16 < 19; fi16++ {
                if false {
                    v7 -= v10
                    v9 -= v9
                    v17 := v6[v7]
                    fmt.Println(v13, v10, v7)
                    delete(v4, (fi12 & (v2 % ((int(v3)) | 1))))
                    _ = v17
                } else {
                    switch "" {
                    case "a":
                        v18 := GenS1{F0: 18446744073709551615}
                        _ = v18
                    default:
                        v19 := GenS1{F0: v15}
                        _ = v19
                    }
                    var v20 int32 = -2147483648
                    _ = v20
                }
                delete(v6, 832)
                v4[(^(fi16))] = (("foo bar" + "a") + (("a" + "hello") + "a"))
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            v21 := GenS1{F0: (v1 | 18446744073709551615)}
            _ = v21
        }
        _ = fi12
        _ = v13
    }
    if false {
        v5 += v5
        if false {
            if (v9 == (v9 + v9)) {
                fmt.Println(v2, v3, v11)
            }
        } else {
            v3 += 1
        }
        fmt.Println(v5)
    } else {
        var v22 string = "\n\t"
        fmt.Println(v2, v2)
        switch (v2 & (v2 & (v2 | (v2 << ((2) & 63))))) {
        case 42:
            v23 := GenS2{F0: (^(v9)), F1: 298}
            _ = v23
        case 32768:
            if true {
                var v24 string = v22
                for fi25 := 0; fi25 < 18; fi25++ {
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if true {
                            _ = len(v8)
                            v24 = v24
                            if (v24 > v24) {
                                v24 = v24
                                fmt.Println(fi25)
                            } else {
                                if true {
                                    var v27 uint = uint(v9)
                                    var v28 uint32 = v10
                                    delete(v6, 122)
                                    if ((fi26 << (((v1 - (v1 / ((v1) | 1)))) & 63)) > fi25) {
                                        if false {
                                            if true {
                                                v29 := GenS2{F0: 860, F1: 430}
                                                var v30 uint16 = 357
                                                var v31 int16 = v5
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                            }
                                            var v32 float64 = 0.0
                                            _ = v32
                                        }
                                        var v33 bool = true
                                        switch -610 {
                                        case 2147483647:
                                            _ = len(v8)
                                        case -9223372036854775808:
                                            var v34 int16 = (v5 ^ (^(v5)))
                                            _ = v34
                                        case -109:
                                            var v35 bool = (v33 || (v22 > v24))
                                            _ = v35
                                        }
                                        _ = v33
                                    } else {
                                        v4[9223372036854775807] = v24
                                    }
                                    var v36 float32 = 1.5
                                    _ = v27
                                    _ = v28
                                    _ = v36
                                } else {
                                    if true {
                                        _ = len(v8)
                                    } else {
                                        var v37 float32 = 1e10
                                        _ = v37
                                    }
                                    _ = len(v8)
                                }
                                v38 := GenS3{F0: 165}
                                switch ("\n\t" + ((v22 + v24) + (v24 + "foo bar"))) {
                                case "foo bar":
                                    switch v22 {
                                    case "":
                                        if false {
                                            v8 = append(v8, fi25)
                                        } else {
                                            v39 := GenS3{F0: 56}
                                            _ = v39
                                        }
                                        v40 := GenS4{F0: -114, F1: v3}
                                        fmt.Println(v11, v7)
                                        v10 = uint32(v1)
                                        _ = v40
                                    }
                                case "":
                                    var v41 int8 = 15
                                    _ = v41
                                default:
                                    var v42 int32 = -1
                                    _ = v42
                                }
                                _ = v38
                            }
                        }
                        if true {
                            v43 := GenS3{F0: uint(v9)}
                            _ = v43
                        } else {
                            if (42 == 181) {
                                fmt.Println(v11)
                            } else {
                                if ((v5 << (((v1 & (v1 - v1))) & 63)) < (v5 << ((v1) & 63))) {
                                    v8 = append(v8, v2)
                                    var v44 int8 = 0
                                    v45 := GenS5{F0: int8(v9), F1: (v1 - v1), F2: v3}
                                    _ = v44
                                    _ = v45
                                } else {
                                    fmt.Println(v7, v24, v2)
                                    v1 += 65535
                                }
                                if len(v8) > 0 {
                                    v8[0] = 0
                                }
                            }
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            var v46 uint16 = 42
                            _ = v46
                        }
                        _ = fi26
                    }
                    _ = fi25
                }
                _ = len(v8)
                _ = v24
            }
        case 32767:
            var v47 uint64 = 0
            _ = v47
        }
        _ = v22
    }
    var v48 bool = true
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    switch "" {
    case "a":
        _ = len(v4)
    case "\n\t":
        v48 = (!true)
    case "hello":
        if v48 {
            switch v2 {
            case 32767:
                var v49 uint32 = v10
                _ = v49
            }
        }
    default:
        var v50 bool = false
        _ = v50
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        var v52 float64 = float64(0.1)
        _ = fi51
        _ = v52
    }
    for fi53 := 0; fi53 < 12; fi53++ {
        var v54 float32 = float32(1e10)
        _ = len(v8)
        delete(v6, v7)
        _ = fi53
        _ = v54
    }
    switch (v2 & (v2 ^ (v2 % ((int(v9)) | 1)))) {
    case 2147483647:
        if (((int32(v1) < -1) || v48) && true) {
            var v55 int64 = -9223372036854775808
            fmt.Println(v55, v9)
            _ = v55
        } else {
            v10 += v10
            for fi56 := 0; fi56 < 18; fi56++ {
                if false {
                    switch ("" + "hello") {
                    case "\n\t":
                        for fi57 := 0; fi57 < 3; fi57++ {
                            var v58 int64 = (v3 << ((v1) & 63))
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            var v59 float64 = 0.0
                            v7 = v10
                            if true {
                                fmt.Println(v10, v48, v7)
                                switch fi56 {
                                case 4294967295:
                                    if len(v8) > 0 {
                                        v8[0] = (v2 / ((fi57) | 1))
                                    }
                                case -412:
                                    v9 = v9
                                case 32768:
                                    v60 := GenS1{F0: (v1 / (((v1 / ((94) | 1))) | 1))}
                                    _ = v60
                                default:
                                    var v61 int16 = v5
                                    _ = v61
                                }
                                switch v2 {
                                case -2147483648:
                                    var v62 uint8 = 42
                                    _ = v62
                                case 0:
                                    var v63 int32 = 1
                                    _ = v63
                                case 32768:
                                    _ = len(v8)
                                default:
                                    fmt.Println(fi56, v58, v10)
                                }
                            }
                            v64 := v4[(v2 >> ((uint(v9)) & 63))]
                            _ = fi57
                            _ = v58
                            _ = v59
                            _ = v64
                        }
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    if false {
                        if (v9 >= (^(v9))) {
                            if (uint64(v7) == (v9 * v9)) {
                                v1 = (v1 - (v1 & v1))
                            }
                            v11 += 1.5
                            v7 = (v7 + v7)
                        } else {
                            delete(v4, (^(fi56)))
                            if (!false) {
                                v4[(-(fi56))] = ""
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    fmt.Println(fi65, v3)
                                    _ = fi65
                                }
                            }
                        }
                    }
                } else {
                    fmt.Println(v2)
                    var v66 int64 = v3
                    _ = v66
                }
                _ = fi56
            }
        }
    case 859:
        if (42 < 803) {
            v48 = ((!(1 == 2147483647)) && (((^(v1)) >= v1) || ((1 == -1) && true)))
        }
    case 32767:
        switch ("" + "\n\t") {
        case "\n\t":
            v67 := GenS1{F0: 865}
            _ = v67
        case "":
            fmt.Println(v9, v9, v48)
        case "hello":
            var v68 uint = (v1 / ((1) | 1))
            _ = v68
        default:
            v48 = v48
        }
    default:
        v69 := GenS4{F0: -53, F1: (v3 * v3)}
        _ = v69
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
