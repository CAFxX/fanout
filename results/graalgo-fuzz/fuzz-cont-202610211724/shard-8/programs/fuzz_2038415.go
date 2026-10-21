package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 int16) int {
    var v7 int = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        p0 = "hello"
        _ = fi8
    }
    fmt.Println(p1)
    v7 -= v7
    fmt.Println(v7, p0)
    _ = p0
    _ = p1
    _ = v7
    return v7
}

func main() {
    var v1 int64 = 42
    v2 := []uint32{2, 2147483647, 2147483648}
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v5 := []uint32{}
    v6 := []uint32{}
    var v9 int64 = (v1 * (^(v1)))
    var v10 uint8 = 0
    if (!false) {
        var v11 int64 = v9
        _ = v11
    }
    fmt.Println(v10)
    v10 = (v10 - (v10 % ((v10) | 1)))
    v9 += 1
    v1 += v9
    if true {
        switch genFn0("", 1) {
        case 719:
            v10 = v10
        }
        switch 628 {
        case 937:
            v12 := GenS1{F0: 2147483648, F1: (v9 * v9)}
            _ = v12
        case 0:
            if len(v2) > 0 {
                v2[0] = 0
            }
        default:
            var v13 int16 = 338
            _ = v13
        }
    }
    switch -9223372036854775808 {
    case 32767:
        v9 += v1
    }
    var v14 uint8 = v10
    if false {
        for fi15 := 0; fi15 < 14; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                var v17 int = fi16
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        v18 := GenS1{F0: 381, F1: v9}
        switch genFn0(("" + "foo bar"), 257) {
        case -777:
            fmt.Println(v14, v10, v1)
        default:
            fmt.Println(v1, v14, v1)
        }
        _ = v18
    }
    v14 -= 42
    var v19 float64 = ((((1.5 - 1e10) / (float64(-2.25) * 0.1)) - 0.0) / -2.25)
    switch ("" + ("" + ("hello" + "a"))) {
    case "\n\t":
        var v20 int32 = 477
        _ = v20
    case "a":
        for fi21 := 0; fi21 < 4; fi21++ {
            for fi22 := 0; fi22 < 7; fi22++ {
                if false {
                    v3["foo bar"] = (((("a" + ("hello" + ("" + "a"))) + ("foo bar" + "hello")) + "") + "a")
                    v1 = -380
                    v23 := GenS2{F0: 184, F1: -43}
                    _ = v23
                }
                var v24 int8 = -95
                _ = len(v5)
                _ = fi22
                _ = v24
            }
            for fi25 := 0; fi25 < 4; fi25++ {
                if true {
                    var v26 int8 = 1
                    _ = v26
                }
                var v27 int = fi21
                var v28 int = (fi21 - fi21)
                _ = fi25
                _ = v27
                _ = v28
            }
            if false {
                v4[713] = 540
                _ = len(v5)
                v29 := GenS1{F0: 2147483648, F1: v9}
                _ = v29
            }
            _ = fi21
        }
    case "foo bar":
        for fi30 := 0; fi30 < 14; fi30++ {
            if true {
                var v31 uint = 2
                _ = v31
            } else {
                for fi32 := 0; fi32 < 6; fi32++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    v4[2147483648] = 2147483647
                    if false {
                        v33 := GenS1{F0: 2, F1: v1}
                        v14 -= v10
                        fmt.Println(v14, v10)
                        switch ("a" + "hello") {
                        case "\n\t":
                            v9 -= (v9 << ((2) & 63))
                        case "a":
                            v34 := GenS2{F0: 2, F1: -8}
                            _ = v34
                        case "hello":
                            v19 -= (v19 * v19)
                            fmt.Println(v9)
                        default:
                            switch "foo bar" {
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "":
                                if false {
                                    v14 = v10
                                } else {
                                    v9 = v1
                                }
                            default:
                                var v35 string = "hello"
                                _ = v35
                            }
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v33
                    }
                    _ = fi32
                }
            }
            _ = fi30
        }
    }
    if false {
        v1 = 0
        v1 = 814
        delete(v3, "")
    }
    v14 += v14
    if false {
        switch -2147483648 {
        case -1:
            for fi36 := 0; fi36 < 13; fi36++ {
                var v37 uint16 = 65535
                v19 -= float64(0.1)
                for fi38 := 0; fi38 < 2; fi38++ {
                    v39 := GenS3{F0: v14}
                    v40 := GenS1{F0: uint32(v10), F1: int64(v37)}
                    v10 -= (v10 + v10)
                    _ = len(v5)
                    var v41 uint32 = 917
                    v42 := GenS4{F0: 0, F1: uint16(v10), F2: 1}
                    v1 += 0
                    v43 := GenS5{F0: 0}
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                _ = fi36
                _ = v37
            }
        case 4294967295:
            v1 = 703
        case 574:
            v9 = v1
        }
        for fi44 := 0; fi44 < 13; fi44++ {
            for fi45 := 0; fi45 < 9; fi45++ {
                var v46 int64 = -245
                _ = fi45
                _ = v46
            }
            if false {
                for fi47 := 0; fi47 < 11; fi47++ {
                    var v48 uint = 255
                    var v49 float32 = 1e10
                    var v50 int8 = -125
                    v48 = v48
                    _ = fi47
                    _ = v48
                    _ = v49
                    _ = v50
                }
            } else {
                var v51 bool = true
                _ = v51
            }
            switch "a" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "a":
                v1 = v1
            }
            _ = fi44
        }
        var v52 uint8 = 87
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
