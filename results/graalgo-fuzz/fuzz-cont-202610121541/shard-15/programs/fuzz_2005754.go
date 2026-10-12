package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 string; F2 uint }
type GenS6 struct { F0 int32; F1 uint; F2 uint64 }
type GenS7 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 int = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int16 = 1000
    v5 := []uint32{1, 468, 1, 643}
    fmt.Println(v1)
    switch ("" + "hello") {
    case "":
        switch -357 {
        case -815:
            fmt.Println(v1)
        case 128:
            v4 -= -518
        case -9223372036854775808:
            v1 = (v1 ^ v1)
        default:
            var v6 int64 = 1
            _ = v6
        }
    case "a":
        switch "foo bar" {
        case "\n\t":
            var v7 bool = true
            _ = v7
        case "foo bar":
            for fi8 := 0; fi8 < 12; fi8++ {
                v4 = v4
                _ = fi8
            }
        case "hello":
            v4 -= (^(v4))
        default:
            for fi9 := 0; fi9 < 11; fi9++ {
                switch "" {
                case "foo bar":
                    fmt.Println(v4, fi9)
                default:
                    v4 = v4
                }
                switch (v1 * (v1 & v1)) {
                case -260:
                    v4 += v4
                case 0:
                    if len(v5) > 0 {
                        v5[0] = 2147483648
                    }
                case -2147483648:
                    v1 -= (v1 - (fi9 | v1))
                }
                v3[862] = "a"
                _ = fi9
            }
        }
    default:
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS1{F0: "", F1: 0}
            v5 = append(v5, 1)
            var v12 bool = true
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    v1 = 1
    var v13 uint32 = 2
    for fi14 := 0; fi14 < 20; fi14++ {
        v4 = -504
        var v15 int32 = 1
        var v16 int = (v1 + (v1 * v1))
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: "foo bar", F1: -1}
    for fi18 := 0; fi18 < 1; fi18++ {
        switch (^(fi18)) {
        case 128:
            v2[(v13 << ((2) & 63))] = 1
        default:
            if true {
                v4 = v4
            } else {
                fmt.Println(v13)
                switch fi18 {
                case 9223372036854775807:
                    v13 = v13
                default:
                    v4 += v4
                }
                v19 := GenS1{F0: "hello", F1: -2147483648}
                _ = v19
            }
        }
        v2[785] = v13
        _ = fi18
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS2{F0: (-(v1)), F1: true}
        _ = fi20
        _ = v21
    }
    switch v1 {
    case -126:
        for fi22 := 0; fi22 < 4; fi22++ {
            if true {
                v2[(v13 + v13)] = (^(v13))
                if (1e10 >= 3.14159) {
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if (0 != 890) {
                            var v24 uint = 65535
                            if (false && true) {
                                v25 := GenS3{F0: 264}
                                _ = len(v2)
                                v13 += uint32(v24)
                                _ = v25
                            }
                            var v26 int16 = -32768
                            v27 := GenS4{F0: 360}
                            var v28 uint16 = 0
                            _ = v24
                            _ = v26
                            _ = v27
                            _ = v28
                        } else {
                            switch fi22 {
                            case 577:
                                var v29 uint64 = 18446744073709551615
                                _ = v29
                            case 0:
                                v30 := GenS4{F0: 616}
                                v13 -= v13
                                _ = v30
                            default:
                                v31 := GenS3{F0: 42}
                                var v32 uint32 = (v13 << ((2) & 63))
                                _ = v31
                                _ = v32
                            }
                        }
                        if false {
                            v33 := GenS1{F0: "\n\t", F1: -1}
                            v13 += 382
                            if true {
                                v34 := GenS5{F0: v4, F1: "hello", F2: 65535}
                                _ = len(v3)
                                var v35 int32 = -2147483648
                                _ = v34
                                _ = v35
                            } else {
                                _ = len(v3)
                                var v36 int64 = -9223372036854775808
                                var v37 uint8 = 0
                                _ = v36
                                _ = v37
                            }
                            _ = v33
                        } else {
                            v13 = v13
                            var v38 uint16 = 1
                            _ = v38
                        }
                        v39 := GenS2{F0: v1, F1: true}
                        var v40 int64 = 0
                        delete(v2, (v13 / (((v13 << ((2) & 63))) | 1)))
                        _ = fi23
                        _ = v39
                        _ = v40
                    }
                    if true {
                        v5 = append(v5, v13)
                    }
                    var v41 int8 = 0
                    _ = v41
                }
            }
            delete(v3, v13)
            _ = fi22
        }
    }
    v42 := GenS6{F0: 0, F1: 2, F2: 233}
    for fi43 := 0; fi43 < 3; fi43++ {
        v44 := v2[v13]
        var v45 bool = (!true)
        fmt.Println(v45)
        _ = fi43
        _ = v44
        _ = v45
    }
    v13 = v13
    v46 := GenS5{F0: (v4 / (((v4 % ((v4) | 1))) | 1)), F1: "\n\t", F2: 0}
    for fi47 := 0; fi47 < 14; fi47++ {
        if (642 <= 42) {
            var v48 int32 = 2147483647
            v13 = 347
            _ = v48
        }
        for fi49 := 0; fi49 < 10; fi49++ {
            v50 := GenS7{F0: (fi47 ^ (v1 << ((95) & 63))), F1: "\n\t", F2: 843}
            _ = fi49
            _ = v50
        }
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v5
}
