package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint }

func main() {
    var v1 bool = (false && false)
    v2 := []string{}
    v3 := make(map[int]int)
    v4 := []int{-1, 2147483647}
    switch 1 {
    case 1:
        fmt.Println(v1)
    case 128:
        fmt.Println(v1)
    case -2147483648:
        v5 := GenS1{F0: 42, F1: 42}
        _ = v5
    }
    var v6 int32 = -1
    if ((18446744073709551615 == 2) || (0 < 2)) {
        v1 = true
        if v1 {
            if v1 {
                for fi7 := 0; fi7 < 15; fi7++ {
                    delete(v3, (fi7 % ((fi7) | 1)))
                    fmt.Println(v1)
                    var v8 float32 = (-2.25 + (1.5 / 1.0))
                    _ = fi7
                    _ = v8
                }
                v9 := GenS2{F0: "hello"}
                _ = v9
            } else {
                var v10 uint8 = 127
                _ = v10
            }
            delete(v3, 795)
        } else {
            var v11 int64 = 0
            _ = v11
        }
        _ = len(v3)
    } else {
        var v12 float32 = 1.5
        _ = v12
    }
    switch 42 {
    case -485:
        switch 127 {
        case 2147483647:
            if false {
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(v1, v1)
                    if v1 {
                        v6 = (v6 * v6)
                    } else {
                        fmt.Println(fi13, v1)
                        switch -399 {
                        case -20:
                            v6 = v6
                        case -445:
                            var v14 bool = v1
                            _ = v14
                        case -962:
                            for fi15 := 0; fi15 < 11; fi15++ {
                                if (229 < 0) {
                                    fmt.Println(fi15, fi13, fi15)
                                }
                                fmt.Println(v1, fi15, v1)
                                v16 := v3[(fi13 & 128)]
                                _ = fi15
                                _ = v16
                            }
                        default:
                            var v17 int32 = 0
                            _ = v17
                        }
                    }
                    _ = fi13
                }
                switch "\n\t" {
                case "hello":
                    var v18 uint32 = 1
                    _ = v18
                case "\n\t":
                    var v19 int64 = 42
                    _ = v19
                case "a":
                    switch 559 {
                    case 128:
                        v6 -= v6
                    case 185:
                        switch -9223372036854775808 {
                        case 32768:
                            for fi20 := 0; fi20 < 16; fi20++ {
                                fmt.Println(v1, v1, fi20)
                                _ = fi20
                            }
                        }
                    case 1:
                        if len(v4) > 0 {
                            v4[0] = 687
                        }
                    }
                }
                switch 128 {
                case -289:
                    var v21 uint64 = 26
                    _ = v21
                case 128:
                    switch -9223372036854775808 {
                    case 312:
                        var v22 string = ("a" + "foo bar")
                        _ = v22
                    case 0:
                        switch ("hello" + "") {
                        case "a":
                            v6 = (v6 ^ v6)
                        case "foo bar":
                            var v23 uint32 = 1
                            _ = v23
                        default:
                            if ((true && true) || v1) {
                                for fi24 := 0; fi24 < 8; fi24++ {
                                    var v25 uint = 0
                                    if (1 == 1000) {
                                        v6 += v6
                                        var v26 int64 = 9223372036854775807
                                        _ = v26
                                    } else {
                                        var v27 uint16 = 1
                                        v28 := GenS3{F0: (v27 + 42), F1: v1}
                                        switch (fi24 << ((v25) & 63)) {
                                        case 32767:
                                            var v29 int64 = -9223372036854775808
                                            _ = v29
                                        default:
                                            var v30 int = 32768
                                            _ = v30
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                    for fi31 := 0; fi31 < 3; fi31++ {
                                        v32 := v3[(fi24 / (((-(fi24))) | 1))]
                                        _ = fi31
                                        _ = v32
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                                if (((255 != 201) || false) && ((-30 != 1) && ((v1 || v1) || (232 == 2)))) {
                                    v1 = (!(680 == 42))
                                } else {
                                    var v33 uint32 = 2
                                    _ = v33
                                }
                            } else {
                                v34 := GenS2{F0: "foo bar"}
                                _ = v34
                            }
                        }
                    }
                case -2147483648:
                    var v35 uint16 = 0
                    _ = v35
                }
            } else {
                var v36 int16 = 1000
                _ = v36
            }
        case -9223372036854775808:
            switch -946 {
            case 42:
                v1 = true
            case 9223372036854775807:
                v1 = ((2147483647 <= 2147483648) && (65535 > 2))
            case 17:
                v6 -= 1
            }
        case 127:
            var v37 uint = 42
            _ = v37
        default:
            if len(v4) > 0 {
                v4[0] = -2147483648
            }
        }
    case -326:
        switch 9223372036854775807 {
        case 32767:
            for fi38 := 0; fi38 < 1; fi38++ {
                _ = len(v4)
                switch 138 {
                case 127:
                    v39 := GenS2{F0: "\n\t"}
                    _ = v39
                }
                _ = fi38
            }
        case -9223372036854775808:
            v6 = (v6 >> ((65535) & 63))
        default:
            switch -1 {
            case 9223372036854775807:
                for fi40 := 0; fi40 < 6; fi40++ {
                    fmt.Println(v1)
                    _ = fi40
                }
            case 128:
                var v41 int64 = 42
                _ = v41
            }
        }
    case 2147483647:
        for fi42 := 0; fi42 < 1; fi42++ {
            var v43 float64 = 3.14159
            v44 := GenS4{F0: 42, F1: 127}
            _ = fi42
            _ = v43
            _ = v44
        }
    default:
        var v45 int8 = 0
        _ = v45
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 32767 {
    case 0:
        if v1 {
            v1 = (false || (2147483647 <= 4294967295))
            switch -39 {
            case -2147483648:
                _ = len(v3)
            case 32767:
                if false {
                    fmt.Println(v6, v6)
                }
            }
        } else {
            var v46 int32 = (v6 - v6)
            switch 2147483647 {
            case 42:
                fmt.Println(v6, v1, v1)
            case 0:
                if v1 {
                    switch 128 {
                    case -9223372036854775808:
                        if len(v4) > 0 {
                            v4[0] = 614
                        }
                    default:
                        for fi47 := 0; fi47 < 6; fi47++ {
                            switch "foo bar" {
                            case "foo bar":
                                v6 += (^(v46))
                            case "":
                                var v48 int16 = -32768
                                _ = v48
                            case "a":
                                _ = len(v3)
                            }
                            _ = fi47
                        }
                    }
                    for fi49 := 0; fi49 < 16; fi49++ {
                        var v50 int64 = -1
                        v50 += v50
                        var v51 float64 = (float64(1e10) * (1.5 + 0.0))
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                }
            case 224:
                if v1 {
                    var v52 string = "foo bar"
                    var v53 uint64 = 42
                    var v54 uint = 385
                    _ = v52
                    _ = v53
                    _ = v54
                } else {
                    v55 := GenS4{F0: 2, F1: -1}
                    if v1 {
                        if (-1 > -128) {
                            var v56 int = 205
                            if v1 {
                                switch 9223372036854775807 {
                                case -889:
                                    v57 := v3[v56]
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 17; fi58++ {
                                    v6 = (v46 << ((255) & 63))
                                    v59 := GenS5{F0: v56}
                                    v2 = append(v2, "foo bar")
                                    _ = fi58
                                    _ = v59
                                }
                            }
                            _ = v56
                        }
                    }
                    _ = v55
                }
            default:
                var v60 uint16 = 42
                _ = v60
            }
            var v61 int64 = 1
            _ = v46
            _ = v61
        }
    case -9223372036854775808:
        v62 := GenS6{F0: 883}
        _ = v62
    default:
        v6 -= (v6 & v6)
    }
    v1 = (true || v1)
    v6 -= v6
    v63 := GenS4{F0: 65535, F1: -83}
    var v64 int32 = 1
    var v65 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v63
    _ = v64
    _ = v65
}
