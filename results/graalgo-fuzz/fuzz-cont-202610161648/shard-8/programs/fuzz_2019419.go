package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint8; F1 int32 }

func main() {
    var v1 uint8 = 83
    v2 := make(map[uint32]uint32)
    var v3 int32 = -931
    v4 := make(map[string]uint32)
    var v5 uint = 4294967295
    v6 := make(map[string]bool)
    v1 -= (v1 << ((v5) & 63))
    var v7 uint16 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        fmt.Println(v7, v1, v3)
        v7 += (v7 * (v7 ^ 1))
        _ = fi8
    }
    v9 := GenS1{F0: 32767}
    if true {
        for fi10 := 0; fi10 < 14; fi10++ {
            v3 = 2147483647
            fmt.Println(v7, v7, v7)
            _ = fi10
        }
        v11 := []int{}
        _ = v11
    } else {
        v12 := GenS1{F0: -982}
        _ = v12
    }
    v1 -= (v1 & (v1 << (((v5 >> ((v5) & 63))) & 63)))
    if false {
        switch 178 {
        case -2147483648:
            v13 := GenS2{F0: v3, F1: 307}
            _ = v13
        default:
            _ = len(v4)
        }
        var v14 int8 = 0
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                v7 = v7
            } else {
                v4["hello"] = 608
                for fi16 := 0; fi16 < 2; fi16++ {
                    var v17 float64 = 1e10
                    var v18 float32 = ((0.0 + 1.5) + float32((0.1 - (1.5 / (v17 * ((1.5 / (v17 / v17)) / (0.0 + v17)))))))
                    switch 9223372036854775807 {
                    case 2147483647:
                        var v19 int32 = (v3 ^ v3)
                        _ = v19
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
            var v20 uint64 = 568
            _ = fi15
            _ = v20
        }
        _ = v14
    } else {
        switch 9223372036854775807 {
        case -52:
            fmt.Println(v1, v3, v1)
        }
        var v21 uint16 = v7
        _ = v21
    }
    v22 := GenS1{F0: 9223372036854775807}
    var v23 int16 = 32767
    if ((v5 + 0) < (v5 % (((v5 & 4294967295)) | 1))) {
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 uint64 = 1
            v26 := GenS1{F0: 4294967295}
            _ = fi24
            _ = v25
            _ = v26
        }
    } else {
        switch int(v7) {
        case 4294967295:
            var v27 bool = ((^(v23)) > v23)
            _ = v27
        case -9223372036854775808:
            v28 := []int{32767}
            _ = v28
        case 688:
            _ = len(v4)
        }
    }
    var v29 uint32 = 658
    v30 := GenS2{F0: int32(v7), F1: 9223372036854775807}
    v31 := []int{32767, 127, 1, 2147483647}
    switch 880 {
    case 4294967295:
        switch "hello" {
        case "a":
            var v32 uint8 = (v1 / ((224) | 1))
            _ = v32
        case "foo bar":
            v33 := GenS2{F0: (v3 + 2147483647), F1: -9223372036854775808}
            _ = v33
        case "hello":
            var v34 uint32 = 777
            _ = v34
        default:
            if true {
                var v35 int16 = v23
                var v36 int = 32767
                _ = v35
                _ = v36
            } else {
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 int32 = v3
                    delete(v4, "hello")
                    _ = fi37
                    _ = v38
                }
            }
        }
    case 128:
        switch "" {
        case "hello":
            for fi39 := 0; fi39 < 12; fi39++ {
                v40 := GenS3{F0: v3}
                v1 -= 0
                _ = fi39
                _ = v40
            }
        default:
            var v41 bool = false
            _ = v41
        }
    }
    if false {
        if ((v23 >> (((v5 | (v5 & uint(v29)))) & 63)) >= int16(v7)) {
            v31 = append(v31, 32768)
            if false {
                for fi42 := 0; fi42 < 2; fi42++ {
                    if true {
                        var v43 float64 = 0.0
                        _ = v43
                    }
                    switch ((((("a" + "hello") + ("a" + "foo bar")) + "a") + "foo bar") + "\n\t") {
                    case "a":
                        fmt.Println(v23, fi42)
                    case "\n\t":
                        var v44 uint32 = 4294967295
                        _ = v44
                    case "hello":
                        v5 -= 4294967295
                    default:
                        _ = len(v2)
                    }
                    if (((((0.1 + -2.25) / (1.5 - float32(((((float64((((float32((((((1.5 * 3.14159) / (1e10 + 1e10)) + 0.0) / (0.0 - (1.5 / (((0.1 * 1e10) - 0.1) * 3.14159)))) + float64((3.14159 / -2.25)))) / (0.1 + 3.14159)) + 0.1) + 1e10)) + (3.14159 / 1.0)) * ((1.5 * ((-2.25 / 1.5) / 0.1)) + 1.5)) + 0.0) * 3.14159)))) <= -2.25) && ((!(((v3 * (v3 % ((v3) | 1))) >= v3) && false)) && (v29 != (^(v29))))) || (42 < 0)) {
                        v45 := GenS4{F0: -9223372036854775808, F1: v5, F2: 4294967295}
                        var v46 int64 = 433
                        _ = v45
                        _ = v46
                    }
                    _ = fi42
                }
                for fi47 := 0; fi47 < 19; fi47++ {
                    fmt.Println(v23, v5, v5)
                    switch fi47 {
                    case -1:
                        _ = len(v6)
                    case 9223372036854775807:
                        v7 = v7
                    default:
                        v29 = (v29 * uint32(v7))
                    }
                    _ = fi47
                }
                switch "" {
                case "a":
                    for fi48 := 0; fi48 < 13; fi48++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v29)
                        case "a":
                            if false {
                                v49 := v2[v29]
                                _ = v49
                            }
                        case "":
                            fmt.Println(v3, v7, v5)
                        default:
                            var v50 int16 = -1
                            _ = v50
                        }
                        _ = fi48
                    }
                }
            }
            var v51 uint8 = (^(v1))
            _ = v51
        }
        for fi52 := 0; fi52 < 2; fi52++ {
            v53 := GenS2{F0: v3, F1: 9223372036854775807}
            _ = fi52
            _ = v53
        }
    } else {
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
        v54 := GenS5{F0: uint8(v7), F1: v3}
        var v55 uint16 = (v7 * 695)
        _ = v54
        _ = v55
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
