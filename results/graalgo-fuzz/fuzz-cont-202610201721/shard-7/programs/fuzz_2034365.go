package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint32; F1 int32 }
type GenS5 struct { F0 uint; F1 uint64; F2 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool; F1 int16 }
func genFn0(p0 int64, p1 int8) uint16 {
    v4 := []int{1, 42}
    p0 = (p0 ^ p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int64 = (-(p0))
        p1 -= (^(p1))
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0)
    v7 := GenS1{F0: ""}
    _ = p0
    _ = p1
    _ = v4
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v8 uint16 = 993
    switch "hello" {
    case "hello":
        if true {
            v9 := GenS1{F0: ("a" + "")}
            var v10 uint8 = 6
            switch "\n\t" {
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "hello":
                var v11 int = 127
                _ = v11
            }
            _ = v10
            _ = v9
        } else {
            v8 -= (v8 >> ((775) & 63))
            if (247 == 2) {
                var v12 float64 = (1.5 + (1.5 * (float64(1.5) + -2.25)))
                _ = v12
            }
            v13 := GenS2{F0: 128, F1: 1, F2: false}
            _ = v13
        }
    case "foo bar":
        if ((false || false) || true) {
            if false {
                v8 += (v8 + (v8 ^ v8))
            }
        } else {
            v14 := GenS3{F0: 155, F1: 2147483647}
            switch 4294967295 {
            case 42:
                fmt.Println(v8)
            case -32:
                switch "foo bar" {
                case "foo bar":
                    if true {
                        v15 := GenS4{F0: 1, F1: -19}
                        v8 = 0
                        fmt.Println(v8)
                        _ = v15
                    }
                case "\n\t":
                    if true {
                        switch -9223372036854775808 {
                        case -1:
                            fmt.Println(v8)
                        }
                    }
                case "hello":
                    var v16 int16 = 42
                    _ = v16
                }
            case 127:
                for fi17 := 0; fi17 < 3; fi17++ {
                    v18 := GenS5{F0: 255, F1: 4294967295, F2: 0}
                    v8 += v8
                    v8 = (^(v8))
                    _ = fi17
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 16; fi19++ {
                    if true {
                        v2 = append(v2, (fi19 + fi19))
                        v8 = (v8 << ((258) & 63))
                        v8 = v8
                    }
                    fmt.Println(v8, v8)
                    _ = fi19
                }
            }
            _ = v14
        }
    case "a":
        v8 -= (v8 % ((genFn0(42, 1)) | 1))
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    if true {
        if false {
            var v20 string = ("hello" + "")
            _ = v20
        } else {
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint64 = 2
                _ = fi21
                _ = v22
            }
            var v23 uint = 1
            _ = v23
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 bool = true
            v26 := GenS6{F0: fi24}
            fmt.Println(v8)
            _ = fi24
            _ = v25
            _ = v26
        }
        for fi27 := 0; fi27 < 6; fi27++ {
            for fi28 := 0; fi28 < 7; fi28++ {
                v8 = v8
                _ = fi28
            }
            fmt.Println(v8, fi27)
            _ = fi27
        }
    } else {
        var v29 uint = 0
        v29 -= (^(v29))
        v8 += (v8 | v8)
        _ = v29
    }
    _ = len(v2)
    for fi30 := 0; fi30 < 16; fi30++ {
        switch 2147483647 {
        case 394:
            v31 := GenS3{F0: 0, F1: -106}
            _ = v31
        case 2147483647:
            v8 = (v8 * genFn0(9223372036854775807, 42))
        default:
            for fi32 := 0; fi32 < 6; fi32++ {
                var v33 int8 = 1
                switch -68 {
                case 32767:
                    v8 -= 255
                case 1:
                    if (!false) {
                        v34 := GenS2{F0: 42, F1: 743, F2: false}
                        fmt.Println(v8, v8)
                        v8 += v8
                        var v35 string = "a"
                        _ = v34
                        _ = v35
                    }
                default:
                    var v36 int16 = 1
                    _ = v36
                }
                _ = fi32
                _ = v33
            }
        }
        var v37 uint = 65535
        switch 128 {
        case 4294967295:
            v37 = v37
        default:
            v8 -= (^(v8))
        }
        _ = fi30
        _ = v37
    }
    for fi38 := 0; fi38 < 17; fi38++ {
        v8 = (v8 << ((844) & 63))
        var v39 uint32 = 0
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 int16 = 1
            _ = fi40
            _ = v41
        }
        _ = fi38
        _ = v39
    }
    switch 1 {
    case 0:
        v42 := GenS6{F0: 9223372036854775807}
        _ = v42
    case -241:
        v8 -= 255
    default:
        v43 := GenS6{F0: 128}
        _ = v43
    }
    var v44 int8 = 75
    v44 = -128
    fmt.Println(v44, v8)
    v8 += (v8 >> ((4294967295) & 63))
    v45 := GenS6{F0: 4294967295}
    if ((true && (4294967295 != 156)) && (4294967295 <= 0)) {
        v46 := GenS4{F0: 2147483648, F1: -889}
        if ((1e10 <= float64(1.5)) && (32767 == 1000)) {
            fmt.Println(v8)
            switch 42 {
            case 127:
                fmt.Println(v44, v44)
            case -2147483648:
                if len(v2) > 0 {
                    v2[0] = -997
                }
            case 9223372036854775807:
                v44 = (v44 - (v44 | 92))
            default:
                var v47 uint16 = (v8 % ((genFn0(9223372036854775807, (v44 - v44))) | 1))
                _ = v47
            }
        } else {
            var v48 uint16 = 255
            v49 := GenS4{F0: 2147483648, F1: -2147483648}
            v8 -= v8
            _ = v48
            _ = v49
        }
        fmt.Println(v44)
        _ = v46
    }
    v44 = int8(v8)
    v50 := GenS7{F0: (2 >= 180), F1: 1}
    for fi51 := 0; fi51 < 9; fi51++ {
        if (!false) {
            v8 -= (v8 / ((v8) | 1))
        }
        _ = fi51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v50
    _ = v8
}
