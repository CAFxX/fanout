package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64; F2 string }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32; F1 string; F2 int }
type GenS5 struct { F0 int64; F1 int32; F2 uint64 }
func genFn0(p0 bool) uint8 {
    switch "foo bar" {
    case "\n\t":
        var v4 int = 0
        _ = v4
    case "":
        if (0 == 2147483647) {
            for fi5 := 0; fi5 < 2; fi5++ {
                var v6 int16 = 32767
                v7 := GenS1{F0: 4}
                _ = fi5
                _ = v6
                _ = v7
            }
            p0 = (39 != 128)
        }
    }
    fmt.Println(p0)
    for fi8 := 0; fi8 < 3; fi8++ {
        p0 = (((!(42 <= 42)) && true) || p0)
        v9 := []int{(fi8 + 42)}
        _ = fi8
        _ = v9
    }
    _ = p0
    return 1
}

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483647, 147, 819}
    var v3 int8 = 100
    _ = len(v2)
    var v10 int8 = 100
    if true {
        for fi11 := 0; fi11 < 15; fi11++ {
            if false {
                v12 := GenS1{F0: 72}
                _ = v12
            }
            v13 := GenS1{F0: 0}
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint64 = 2
                var v16 uint8 = 2
                v17 := GenS2{F0: 0, F1: int64(v15), F2: "\n\t"}
                fmt.Println(fi14, fi14)
                v10 += -128
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi11
            _ = v13
        }
        v18 := GenS3{F0: 1, F1: 867}
        for fi19 := 0; fi19 < 19; fi19++ {
            for fi20 := 0; fi20 < 11; fi20++ {
                var v21 bool = (!false)
                var v22 uint8 = 2
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi19
        }
        _ = v18
    }
    switch (v1 + (v1 + v1)) {
    case "":
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                switch fi23 {
                case 2147483647:
                    for fi24 := 0; fi24 < 5; fi24++ {
                        var v25 int64 = -964
                        _ = fi24
                        _ = v25
                    }
                case 32768:
                    if (!false) {
                        var v26 uint = 4294967295
                        _ = v26
                    } else {
                        fmt.Println(v10, fi23, v10)
                    }
                default:
                    var v27 float64 = 0.0
                    _ = v27
                }
                v28 := GenS4{F0: 2147483647, F1: v1, F2: (fi23 << ((639) & 63))}
                _ = fi23
                _ = v28
            }
            var v29 uint16 = 906
            for fi30 := 0; fi30 < 3; fi30++ {
                if (!true) {
                    var v31 uint = 1
                    for fi32 := 0; fi32 < 16; fi32++ {
                        switch fi32 {
                        case 42:
                            switch "a" {
                            case "hello":
                                if ((^(v3)) > 100) {
                                    var v33 uint64 = 501
                                    var v34 uint8 = uint8(v33)
                                    _ = v33
                                    _ = v34
                                }
                            }
                            var v35 float32 = -2.25
                            fmt.Println(v3, v3, v35)
                            _ = v35
                        case -9223372036854775808:
                            if false {
                                fmt.Println(fi32, v1)
                                v36 := GenS5{F0: -1, F1: 253, F2: 0}
                                v2 = append(v2, 2)
                                _ = v36
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v3, fi32)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v37 uint64 = 1
                                _ = v37
                            }
                        case -977:
                            v1 = v1
                        default:
                            if (((v1 != v1) && (933 != -1)) && (!(!true))) {
                                v10 = v10
                            }
                        }
                        _ = fi32
                    }
                    _ = v31
                }
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                _ = fi30
            }
            _ = v29
        } else {
            v38 := GenS2{F0: 42, F1: 0, F2: v1}
            switch v1 {
            case "hello":
                var v39 int64 = -150
                _ = v39
            case "a":
                var v40 uint = 4294967295
                _ = v40
            case "foo bar":
                v41 := GenS4{F0: 2147483647, F1: (v1 + v1), F2: 9223372036854775807}
                _ = v41
            default:
                v10 = (-(v10))
            }
            _ = v38
        }
    case "a":
        v3 = (v3 >> ((255) & 63))
    case "hello":
        fmt.Println(v1, v10)
    }
    switch v1 {
    case "foo bar":
        var v42 int = 500
        _ = v42
    }
    var v43 uint = 28
    if (-1 > -2147483648) {
        for fi44 := 0; fi44 < 2; fi44++ {
            fmt.Println(v3, v1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi44
        }
        v45 := GenS1{F0: 59}
        _ = v45
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v46 int32 = 42
    for fi47 := 0; fi47 < 12; fi47++ {
        var v48 float64 = 1.5
        _ = fi47
        _ = v48
    }
    if false {
        fmt.Println(v10, v1)
        v49 := GenS5{F0: -9223372036854775808, F1: (v46 & v46), F2: 0}
        switch ("\n\t" + ("hello" + "\n\t")) {
        case "a":
            switch 42 {
            case 1:
                v43 -= 630
            }
        case "foo bar":
            if false {
                fmt.Println(v43)
                switch 0 {
                case -1:
                    var v50 int64 = 867
                    _ = v50
                case 584:
                    var v51 float64 = float64((0.1 - (0.0 - float32(1e10))))
                    _ = v51
                case 9223372036854775807:
                    v10 = 127
                default:
                    _ = len(v2)
                }
            }
        case "\n\t":
            var v52 int8 = (-(v10))
            _ = v52
        default:
            if ((42 != 2) && (!false)) {
                v53 := GenS3{F0: 2, F1: (v46 * -1)}
                switch v1 {
                case "":
                    if (1 == 42) {
                        v43 = (v43 << ((0) & 63))
                    }
                default:
                    fmt.Println(v3, v46)
                }
                _ = v53
            } else {
                v43 -= (^(v43))
                var v54 string = v1
                switch v54 {
                case "a":
                    if false {
                        var v55 uint8 = 42
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        _ = v55
                    }
                case "hello":
                    for fi56 := 0; fi56 < 9; fi56++ {
                        v54 = (v1 + ((v54 + "") + "\n\t"))
                        fmt.Println(v43, v54)
                        v54 = v1
                        _ = fi56
                    }
                default:
                    v3 -= v10
                }
                _ = v54
            }
        }
        _ = v49
    }
    v3 -= v10
    v3 = -128
    if (false || false) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    switch v1 {
    case "hello":
        fmt.Println(v46, v10, v3)
    }
    for fi57 := 0; fi57 < 6; fi57++ {
        v43 -= 255
        fmt.Println(fi57, v10)
        _ = fi57
    }
    fmt.Println(v1, v46, v46)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v43
    _ = v46
}
