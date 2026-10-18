package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint; F1 uint64; F2 int64 }

func main() {
    v1 := []uint32{269, 0, 4294967295}
    v2 := []uint32{}
    var v3 int16 = 1
    var v4 uint = 275
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    if (1 >= 634) {
        switch -293 {
        case 128:
            switch 2147483647 {
            case -121:
                var v7 int = 0
                _ = v7
            case 876:
                for fi8 := 0; fi8 < 11; fi8++ {
                    var v9 int32 = 1
                    _ = len(v6)
                    _ = fi8
                    _ = v9
                }
            case 2147483647:
                if true {
                    for fi10 := 0; fi10 < 13; fi10++ {
                        v3 -= v3
                        _ = fi10
                    }
                    v11 := GenS1{F0: ("foo bar" + "a")}
                    _ = v11
                }
            default:
                switch "" {
                case "\n\t":
                    if (231 >= 35) {
                        var v12 bool = (0.1 >= float64(0.0))
                        _ = v12
                    } else {
                        if (!false) {
                            v6[127] = "\n\t"
                            v6[1] = ("a" + "\n\t")
                        } else {
                            if false {
                                v4 = v4
                                v4 -= (^(v4))
                            }
                            v13 := v5[0]
                            delete(v5, 1)
                            _ = v13
                        }
                        var v14 uint8 = 244
                        if (true || (4294967295 <= 641)) {
                            var v15 int16 = (v3 / (((v3 * v3)) | 1))
                            _ = v15
                        } else {
                            v4 += 2
                            if len(v2) > 0 {
                                v2[0] = 315
                            }
                            for fi16 := 0; fi16 < 19; fi16++ {
                                var v17 int8 = -1
                                v4 = 0
                                _ = fi16
                                _ = v17
                            }
                        }
                        _ = v14
                    }
                default:
                    v18 := GenS2{F0: (776 != -954), F1: v3, F2: 127}
                    _ = v18
                }
            }
        case 1:
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 int8 = 1
                v5[513] = ((v20 & (-(v20))) < v20)
                if true {
                    v20 += (v20 >> ((v4) & 63))
                    v21 := GenS3{F0: 255, F1: 0}
                    _ = v21
                } else {
                    if (false && false) {
                        switch (^(fi19)) {
                        case -539:
                            v4 += 341
                        }
                        v22 := v6[(^(fi19))]
                        _ = v22
                    }
                    fmt.Println(v4)
                }
                _ = fi19
                _ = v20
            }
        case 42:
            if (6 > 4294967295) {
                if (1 > 2147483647) {
                    if (62 < 211) {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        for fi23 := 0; fi23 < 3; fi23++ {
                            v4 = v4
                            _ = fi23
                        }
                        fmt.Println(v3, v4)
                    }
                    for fi24 := 0; fi24 < 3; fi24++ {
                        if (!(true || ((!false) || false))) {
                            switch "" {
                            case "hello":
                                v25 := GenS4{F0: 9223372036854775807, F1: false, F2: 561}
                                _ = v25
                            }
                        }
                        if (v3 <= 0) {
                            var v26 float64 = 1e10
                            var v27 uint = 0
                            _ = v26
                            _ = v27
                        }
                        v3 -= v3
                        _ = fi24
                    }
                } else {
                    v4 += (v4 & v4)
                }
                v3 -= v3
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            }
        default:
            var v28 uint = v4
            _ = v28
        }
        fmt.Println(v3)
    }
    var v29 int16 = (^(v3))
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 3; fi31++ {
            v32 := GenS1{F0: ("foo bar" + "\n\t")}
            v33 := GenS1{F0: "\n\t"}
            _ = fi31
            _ = v32
            _ = v33
        }
        if true {
            v3 -= (v3 >> ((v4) & 63))
            var v34 int32 = int32(v3)
            _ = len(v2)
            _ = v34
        } else {
            var v35 int32 = 2147483647
            if (!(777 != 65535)) {
                if false {
                    var v36 uint64 = 241
                    v37 := GenS5{F0: 42}
                    var v38 float64 = 3.14159
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            fmt.Println(fi30)
            _ = v35
        }
        _ = fi30
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v4 = v4
        v4 = (^(v4))
        v4 = v4
        _ = fi39
    }
    v29 = v29
    v40 := GenS6{F0: 684}
    fmt.Println(v29)
    fmt.Println(v3)
    switch "\n\t" {
    case "a":
        for fi41 := 0; fi41 < 18; fi41++ {
            v29 = 42
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 uint16 = 2
                var v44 int16 = v29
                v2 = append(v2, 2)
                fmt.Println(v3, v44)
                v45 := GenS7{F0: (true || (!(false && false)))}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            if (v29 == -1) {
                var v46 uint8 = 39
                _ = v46
            } else {
                var v47 uint16 = 42
                _ = v47
            }
            _ = fi41
        }
    default:
        var v48 bool = (734 == -1)
        _ = v48
    }
    if (!true) {
        var v49 bool = false
        var v50 int32 = 525
        switch "\n\t" {
        case "\n\t":
            for fi51 := 0; fi51 < 2; fi51++ {
                switch "" {
                case "\n\t":
                    _ = len(v2)
                case "hello":
                    delete(v6, fi51)
                }
                v2 = append(v2, 1)
                _ = fi51
            }
        }
        _ = v49
        _ = v50
    } else {
        v29 = v29
        fmt.Println(v4)
        v52 := GenS8{F0: (v4 & 42), F1: 42, F2: 1}
        _ = v52
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        switch "\n\t" {
        case "hello":
            v54 := GenS3{F0: 1, F1: 633}
            _ = v54
        case "":
            var v55 float64 = 3.14159
            _ = v55
        case "a":
            _ = len(v5)
        }
        v4 += (v4 ^ v4)
        var v56 int8 = 100
        _ = fi53
        _ = v56
    }
    v3 = -1
    for fi57 := 0; fi57 < 3; fi57++ {
        for fi58 := 0; fi58 < 12; fi58++ {
            _ = len(v1)
            _ = fi58
        }
        _ = fi57
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        for fi60 := 0; fi60 < 18; fi60++ {
            if (fi59 <= fi59) {
                if (233 >= 9223372036854775807) {
                    var v61 float32 = 0.0
                    var v62 int64 = 1
                    if false {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        fmt.Println(v29, v29, v29)
                    }
                    _ = v61
                    _ = v62
                }
            }
            switch (fi59 ^ 32767) {
            case 1:
                if len(v1) > 0 {
                    v1[0] = 0
                }
            }
            _ = fi60
        }
        _ = fi59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
