package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS6 struct { F0 uint; F1 string }
type GenS7 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    v1 := []int{127, -376, 127}
    var v2 int8 = 106
    v3 := []string{"hello"}
    v4 := make(map[uint32]uint32)
    var v5 int8 = (v2 + v2)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    v2 = -37
    var v6 uint8 = 2
    switch -1 {
    case 873:
        switch "hello" {
        case "hello":
            if (916 != 1) {
                v7 := v4[2147483648]
                for fi8 := 0; fi8 < 10; fi8++ {
                    for fi9 := 0; fi9 < 4; fi9++ {
                        var v10 uint = 181
                        _ = fi9
                        _ = v10
                    }
                    _ = fi8
                }
                _ = v7
            }
        case "a":
            fmt.Println(v6, v5, v5)
        case "foo bar":
            v11 := GenS1{F0: -2147483648, F1: 2}
            _ = v11
        default:
            var v12 uint32 = 117
            _ = v12
        }
    case 4294967295:
        var v13 uint16 = 32767
        _ = v13
    case 42:
        for fi14 := 0; fi14 < 7; fi14++ {
            v15 := GenS1{F0: 1, F1: 2}
            _ = fi14
            _ = v15
        }
    default:
        v16 := GenS2{F0: 1000, F1: 251, F2: 42}
        _ = v16
    }
    if true {
        var v17 uint8 = 127
        _ = len(v4)
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 14; fi19++ {
                if true {
                    var v20 int = fi19
                    if false {
                        v21 := GenS2{F0: -636, F1: 78, F2: 1}
                        _ = v21
                    } else {
                        if false {
                            v22 := GenS3{F0: 0, F1: (42 >= 431), F2: (((!(v17 != 42)) || (false && ((!(42 != 128)) && true))) && (!false))}
                            _ = v22
                        } else {
                            switch (fi18 + -1) {
                            case 4294967295:
                                fmt.Println(v5, v2, v20)
                            case -1:
                                if false {
                                    v2 = v5
                                }
                            default:
                                if (303 == 65535) {
                                    v6 += 1
                                    fmt.Println(v20)
                                    var v23 int8 = (v2 | (^(v2)))
                                    _ = v23
                                }
                            }
                        }
                    }
                    if true {
                        v17 = (v17 % (((^(v6))) | 1))
                    } else {
                        if len(v1) > 0 {
                            v1[0] = fi19
                        }
                        var v24 uint16 = 986
                        _ = v24
                    }
                    _ = v20
                }
                v25 := GenS4{F0: 42, F1: 2}
                fmt.Println(v17)
                var v26 uint32 = 2147483647
                switch fi18 {
                case -804:
                    v2 = -128
                }
                var v27 uint = 42
                var v28 int8 = v5
                switch "foo bar" {
                case "":
                    switch fi19 {
                    case -433:
                        v27 = (v27 * v27)
                    case 9223372036854775807:
                        v27 -= (v27 << ((v27) & 63))
                    }
                case "hello":
                    v17 += v17
                }
                switch "" {
                case "foo bar":
                    v29 := GenS3{F0: 0, F1: true, F2: (!true)}
                    _ = v29
                case "":
                    _ = len(v4)
                default:
                    v30 := GenS2{F0: 985, F1: v6, F2: int16(v27)}
                    _ = v30
                }
                _ = fi19
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
            _ = fi18
        }
        _ = v17
    }
    fmt.Println(v6, v6, v5)
    v31 := v4[401]
    if (true || true) {
        var v32 uint16 = 255
        var v33 int32 = -1
        _ = v32
        _ = v33
    } else {
        v34 := GenS5{F0: 9223372036854775807, F1: 1, F2: 42}
        v2 -= v2
        _ = v34
    }
    for fi35 := 0; fi35 < 11; fi35++ {
        var v36 uint16 = 229
        if false {
            for fi37 := 0; fi37 < 15; fi37++ {
                v6 = (v6 & 128)
                delete(v4, (v31 << ((838) & 63)))
                var v38 string = "foo bar"
                _ = fi37
                _ = v38
            }
            var v39 uint8 = (v6 | (v6 << ((138) & 63)))
            _ = v39
        } else {
            v40 := GenS6{F0: 255, F1: "foo bar"}
            var v41 int16 = -221
            for fi42 := 0; fi42 < 2; fi42++ {
                v2 -= (v5 / ((v2) | 1))
                fmt.Println(v36)
                _ = fi42
            }
            _ = v40
            _ = v41
        }
        _ = fi35
        _ = v36
    }
    v2 = 36
    v2 = (v5 - (v2 * (v2 & -73)))
    for fi43 := 0; fi43 < 17; fi43++ {
        v44 := v4[421]
        for fi45 := 0; fi45 < 6; fi45++ {
            v46 := GenS7{F0: (^(v6)), F1: int32(fi43), F2: 2}
            var v47 uint8 = v6
            _ = fi45
            _ = v46
            _ = v47
        }
        for fi48 := 0; fi48 < 1; fi48++ {
            var v49 float32 = 1e10
            fmt.Println(v2, fi43, v5)
            _ = fi48
            _ = v49
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
