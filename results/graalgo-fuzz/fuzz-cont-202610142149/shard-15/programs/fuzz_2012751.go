package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 int64 }

func main() {
    var v1 uint16 = 255
    var v2 float32 = 1.5
    var v3 int32 = -2147483648
    var v4 uint8 = 255
    if false {
        v5 := GenS1{F0: (v1 * v1), F1: uint64(v1), F2: 42}
        var v6 uint64 = 278
        v7 := []int{}
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 int8 = 1
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 bool = (((!(343 > 0)) && ("hello" != ("\n\t" + "\n\t"))) && (1 == (-(v3))))
        var v11 int8 = v8
        v3 += (-(v3))
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := []int{127, int(v1)}
    v3 = v3
    fmt.Println(v4, v8)
    fmt.Println(v2)
    fmt.Println(v4, v8, v1)
    if len(v12) > 0 {
        v12[0] = 2147483647
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "hello" {
        case "":
            v4 -= (v4 - v4)
        case "a":
            v14 := GenS1{F0: (v1 ^ (v1 << ((241) & 63))), F1: 18446744073709551615, F2: -772}
            _ = v14
        case "hello":
            for fi15 := 0; fi15 < 13; fi15++ {
                switch fi15 {
                case 325:
                    var v16 float64 = (0.1 - 3.14159)
                    _ = v16
                }
                if (1 > 1) {
                    v17 := GenS2{F0: v4, F1: 799}
                    var v18 int32 = int32(fi13)
                    if (!((fi15 & (^(fi13))) < 0)) {
                        if true {
                            v4 += v4
                            v12 = append(v12, -9223372036854775808)
                            if false {
                                v4 += (v4 % ((v4) | 1))
                                v19 := GenS2{F0: v4, F1: 1}
                                v3 = (v18 % (((v3 >> ((421) & 63))) | 1))
                                _ = v19
                            } else {
                                v8 = v8
                            }
                        } else {
                            v8 = 127
                            v20 := GenS2{F0: 1, F1: 2}
                            _ = v20
                        }
                        if false {
                            var v21 uint64 = 0
                            var v22 bool = ((v8 ^ 42) <= v8)
                            var v23 int8 = v8
                            switch -2147483648 {
                            case 1:
                                v24 := GenS3{F0: (v4 + uint8(v1)), F1: 0}
                                _ = v24
                            case 128:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            default:
                                v23 = 100
                            }
                            if (42 == 506) {
                                switch "foo bar" {
                                case "hello":
                                    var v25 int16 = 754
                                    _ = v25
                                case "foo bar":
                                    var v26 int64 = 0
                                    _ = v26
                                default:
                                    var v27 int = fi15
                                    _ = v27
                                }
                                var v28 uint16 = v1
                                if v22 {
                                    if ((v21 + 0) != 18446744073709551615) {
                                        if v22 {
                                            if (!v22) {
                                                v29 := GenS4{F0: int16(v4), F1: 0}
                                                _ = len(v12)
                                                _ = v29
                                            }
                                            var v30 string = "a"
                                            _ = v30
                                        } else {
                                            switch "" {
                                            case "\n\t":
                                                var v31 int32 = v18
                                                _ = v31
                                            case "foo bar":
                                                v21 += v21
                                            }
                                            if v22 {
                                                if ((^(v4)) > (v4 ^ v4)) {
                                                    _ = len(v12)
                                                    fmt.Println(fi15, v22)
                                                }
                                                if (!(127 > 32768)) {
                                                    var v32 int16 = 1
                                                    _ = v32
                                                } else {
                                                    v33 := GenS3{F0: 127, F1: 480}
                                                    _ = v33
                                                }
                                                v21 = (v21 ^ 4294967295)
                                            } else {
                                                if (v22 || (!(752 > 534))) {
                                                    fmt.Println(v1)
                                                }
                                            }
                                            if (!true) {
                                                v8 += (v23 | (v8 ^ v8))
                                                var v34 uint = 18446744073709551615
                                                v23 = v8
                                                _ = v34
                                            }
                                        }
                                        v12 = append(v12, fi15)
                                        v12 = append(v12, int(v21))
                                    }
                                    var v35 uint = 1
                                    v12 = append(v12, 127)
                                    _ = v35
                                }
                                _ = v28
                            }
                            v23 = (v23 / (((v23 * -1)) | 1))
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            var v36 int8 = (v8 & 127)
                            _ = v36
                        }
                    }
                    _ = v17
                    _ = v18
                }
                _ = fi15
            }
        default:
            var v37 int64 = -9223372036854775808
            _ = v37
        }
        if false {
            switch ("a" + ("" + "hello")) {
            case "hello":
                fmt.Println(v4)
            case "\n\t":
                var v38 float32 = (v2 * float32(1.5))
                _ = v38
            case "foo bar":
                for fi39 := 0; fi39 < 8; fi39++ {
                    v12 = append(v12, fi39)
                    v3 += v3
                    switch (^(fi13)) {
                    case -778:
                        if true {
                            var v40 uint16 = 65535
                            var v41 int64 = -403
                            _ = v40
                            _ = v41
                        }
                    case 4294967295:
                        var v42 float64 = (3.14159 / 0.1)
                        _ = v42
                    default:
                        var v43 int8 = v8
                        _ = v43
                    }
                    var v44 float64 = 0.0
                    fmt.Println(fi13, fi39, v44)
                    _ = fi39
                    _ = v44
                }
            }
            if true {
                v45 := GenS5{F0: 0}
                var v46 uint8 = (v4 * (^(v4)))
                _ = v45
                _ = v46
            }
            if true {
                for fi47 := 0; fi47 < 1; fi47++ {
                    var v48 uint32 = 65535
                    var v49 int32 = (-(v3))
                    var v50 uint16 = (v1 << ((255) & 63))
                    _ = fi47
                    _ = v48
                    _ = v49
                    _ = v50
                }
                if len(v12) > 0 {
                    v12[0] = (fi13 << ((640) & 63))
                }
            }
        }
        _ = fi13
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        var v52 uint32 = uint32(v4)
        if (196 >= v4) {
            var v53 int8 = (v8 * -38)
            _ = v53
        }
        _ = fi51
        _ = v52
    }
    if len(v12) > 0 {
        v12[0] = 127
    }
    if (v2 == v2) {
        fmt.Println(v8, v1)
        if len(v12) > 0 {
            v12[0] = int(v3)
        }
    } else {
        for fi54 := 0; fi54 < 12; fi54++ {
            v3 += v3
            switch "" {
            case "":
                var v55 uint16 = v1
                _ = v55
            }
            if (614 > -32768) {
                v56 := GenS6{F0: 4294967295}
                _ = v56
            }
            _ = fi54
        }
        v57 := GenS7{F0: -9223372036854775808}
        fmt.Println(v4, v3, v8)
        _ = v57
    }
    for fi58 := 0; fi58 < 16; fi58++ {
        v59 := GenS2{F0: v4, F1: 42}
        _ = fi58
        _ = v59
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v60 := GenS8{F0: ("" + "a"), F1: 73}
    var v61 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v60
    _ = v61
    _ = v8
}
