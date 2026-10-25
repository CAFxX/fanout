package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
func genFn0() int {
    v7 := GenS1{F0: 2}
    switch 0 {
    case 0:
        var v8 bool = true
        _ = v8
    case -9223372036854775808:
        var v9 uint64 = 42
        _ = v9
    case -1:
        v10 := []int{42, -341, -2147483648}
        _ = v10
    }
    if (-1 <= -128) {
        var v11 string = "a"
        _ = v11
    } else {
        switch -643 {
        case 953:
            v12 := GenS1{F0: 2147483648}
            _ = v12
        case -136:
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := GenS1{F0: 63}
                _ = fi13
                _ = v14
            }
        case -9223372036854775808:
            var v15 uint64 = 0
            _ = v15
        }
        var v16 int64 = -332
        _ = v16
    }
    switch "\n\t" {
    case "":
        switch -9223372036854775808 {
        case 309:
            v17 := GenS2{F0: ((("a" + "") + "\n\t") + ("foo bar" + ""))}
            _ = v17
        }
    case "hello":
        v18 := []int{128, 9223372036854775807, -397}
        _ = v18
    default:
        var v19 uint16 = 1
        _ = v19
    }
    if ("foo bar" >= "") {
        var v20 uint8 = 42
        var v21 uint8 = 1
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := []int{(fi22 >> ((1) & 63)), (fi22 & 127), (fi22 >> ((18446744073709551615) & 63))}
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    _ = v7
    return 4294967295
}

func main() {
    var v1 int64 = -648
    var v2 int = 128
    var v3 string = "\n\t"
    v4 := []uint32{}
    v5 := []uint32{830, 1, 4294967295}
    v6 := make(map[int]uint32)
    v24 := GenS3{F0: 42, F1: 1}
    var v25 uint = 479
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS4{F0: -725}
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: 0, F1: 824}
    var v29 int16 = -32768
    fmt.Println(v3, v1, v3)
    var v30 uint = v25
    if ((v29 >> (((v25 * (v30 << ((42) & 63)))) & 63)) == (v29 | int16(v1))) {
        switch (((v3 + v3) + v3) + (v3 + "foo bar")) {
        case "a":
            v1 -= (^(v1))
        case "hello":
            if true {
                var v31 int32 = 0
                fmt.Println(v25, v25, v30)
                _ = v31
            } else {
                if false {
                    if true {
                        var v32 uint8 = 136
                        switch "a" {
                        case "\n\t":
                            v33 := GenS3{F0: -826, F1: -2147483648}
                            _ = v33
                        case "hello":
                            switch v2 {
                            case -2147483648:
                                var v34 int8 = 1
                                _ = v34
                            case -560:
                                var v35 int16 = (v29 % ((v29) | 1))
                                _ = v35
                            case 9223372036854775807:
                                switch genFn0() {
                                case 32768:
                                    var v36 uint16 = 255
                                    _ = v36
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "hello":
                                        var v37 uint16 = 32767
                                        _ = v37
                                    case "":
                                        var v38 int32 = 994
                                        _ = v38
                                    case "foo bar":
                                        for fi39 := 0; fi39 < 6; fi39++ {
                                            switch (v2 | 2147483647) {
                                            case 127:
                                                for fi40 := 0; fi40 < 8; fi40++ {
                                                    var v41 int16 = v29
                                                    v42 := GenS1{F0: 2147483647}
                                                    var v43 string = v3
                                                    _ = fi40
                                                    _ = v41
                                                    _ = v42
                                                    _ = v43
                                                }
                                            case 42:
                                                var v44 int8 = 1
                                                _ = v44
                                            }
                                            fmt.Println(v29)
                                            _ = fi39
                                        }
                                    default:
                                        v29 = 0
                                    }
                                }
                            default:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        case "foo bar":
                            switch -353 {
                            case 128:
                                if (-128 == 1) {
                                    v3 = "hello"
                                }
                            default:
                                var v45 uint16 = uint16(v32)
                                _ = v45
                            }
                        default:
                            var v46 uint64 = 4294967295
                            _ = v46
                        }
                        _ = v32
                    }
                }
                v6[(v2 ^ (v2 - v2))] = 4294967295
            }
        case "":
            switch 128 {
            case 32768:
                v1 = (v1 | (^(v1)))
            }
        }
        v1 = (^(v1))
    } else {
        v47 := GenS5{F0: 2147483647}
        _ = v47
    }
    v1 += (^(v1))
    if ("hello" != v3) {
        var v48 int64 = (v1 ^ -9223372036854775808)
        v3 = v3
        _ = v48
    } else {
        for fi49 := 0; fi49 < 20; fi49++ {
            if false {
                _ = len(v6)
                if (-124 <= 0) {
                    v50 := GenS5{F0: 754}
                    _ = v50
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                var v51 uint8 = 0
                var v52 int64 = -83
                _ = v51
                _ = v52
            }
            v1 = int64(v2)
            _ = fi49
        }
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        v54 := GenS5{F0: 2147483648}
        if ((v29 >> ((255) & 63)) <= (v29 - (v29 << ((132) & 63)))) {
            v55 := v6[(fi53 >> (((v30 % ((v25) | 1))) & 63))]
            _ = v55
        } else {
            v56 := GenS4{F0: (-(v1))}
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    fmt.Println(v3, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
