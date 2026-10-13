package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 uint64 }

func main() {
    var v1 float64 = 1.5
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]bool)
    var v4 int16 = -920
    v5 := []uint32{}
    var v6 int64 = (v2 | v2)
    _ = len(v3)
    for fi7 := 0; fi7 < 15; fi7++ {
        v6 = v2
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int32 = 42
            v2 = (v2 >> ((255) & 63))
            v10 := GenS1{F0: -128, F1: 42, F2: (^(fi8))}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v2 = -1
    v11 := GenS1{F0: -67, F1: 128, F2: -2147483648}
    v12 := GenS1{F0: 127, F1: 42, F2: 823}
    fmt.Println(v2)
    var v13 int32 = -2147483648
    var v14 uint32 = 568
    v4 -= (v4 << ((625) & 63))
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 int64 = (v2 ^ 867)
        var v17 int16 = v4
        switch (fi15 / ((fi15) | 1)) {
        case 760:
            switch fi15 {
            case 318:
                switch fi15 {
                case -818:
                    v4 = v4
                case 1:
                    v18 := GenS2{F0: false}
                    _ = v18
                case 625:
                    var v19 int16 = 867
                    _ = v19
                }
            case 2147483647:
                var v20 uint16 = 584
                _ = v20
            case -513:
                v16 = (v2 + v6)
            default:
                v21 := GenS2{F0: false}
                _ = v21
            }
        default:
            v14 += 1
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS2{F0: (true && true)}
        for fi24 := 0; fi24 < 9; fi24++ {
            v13 -= v13
            if (v4 < (v4 << ((704) & 63))) {
                _ = len(v5)
                if false {
                    if ("\n\t" == "\n\t") {
                        fmt.Println(v1, fi22)
                    } else {
                        var v25 string = "foo bar"
                        _ = v25
                    }
                    var v26 int64 = v6
                    if true {
                        _ = len(v3)
                        v27 := GenS1{F0: -128, F1: 103, F2: (fi22 & fi24)}
                        _ = v27
                    }
                    if false {
                        fmt.Println(v6)
                        if (v13 >= -2147483648) {
                            if (!true) {
                                v28 := GenS3{F0: false, F1: -1}
                                v29 := GenS3{F0: false, F1: int32(v14)}
                                _ = v28
                                _ = v29
                            }
                            var v30 uint64 = 18446744073709551615
                            v14 += 0
                            v14 = (v14 + (v14 / (((v14 / (((v14 + (v14 * v14))) | 1))) | 1)))
                            var v31 uint8 = 2
                            _ = v30
                            _ = v31
                        }
                        _ = len(v3)
                    }
                    if len(v5) > 0 {
                        v5[0] = v14
                    }
                    fmt.Println(v13)
                    var v32 int64 = (v2 & v2)
                    _ = v26
                    _ = v32
                } else {
                    var v33 int = fi24
                    _ = v33
                }
                v5 = append(v5, 1)
                if true {
                    _ = len(v5)
                    _ = len(v3)
                    var v34 string = (("" + "") + "a")
                    _ = v34
                } else {
                    var v35 uint32 = v14
                    if true {
                        var v36 string = ((("\n\t" + "foo bar") + "a") + "\n\t")
                        v37 := GenS1{F0: 0, F1: 121, F2: (fi22 + 128)}
                        v13 = v13
                        _ = v36
                        _ = v37
                    } else {
                        v13 += 42
                        if false {
                            v38 := GenS2{F0: (1 < 12)}
                            if false {
                                var v39 int16 = v4
                                v5 = append(v5, (^(v35)))
                                if (true || (((0 > v2) || true) && false)) {
                                    var v40 string = (("a" + ("\n\t" + ("foo bar" + "\n\t"))) + (("hello" + "\n\t") + ("foo bar" + "a")))
                                    _ = v40
                                }
                                _ = v39
                            } else {
                                switch (fi24 | fi24) {
                                case 2147483647:
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v41 uint32 = v35
                                        _ = v41
                                    case "foo bar":
                                        var v42 float64 = (1e10 * v1)
                                        _ = v42
                                    }
                                default:
                                    fmt.Println(v14, v2)
                                }
                                v35 -= (^(v35))
                            }
                            v43 := v3[42]
                            _ = v38
                            _ = v43
                        }
                    }
                    v6 = v2
                    _ = v35
                }
                v3[v14] = false
            }
            if (false || false) {
                var v44 float64 = (-2.25 + 1e10)
                v6 -= (v6 >> ((65535) & 63))
                v45 := GenS4{F0: (fi22 - (fi22 << ((65535) & 63))), F1: -34, F2: 32}
                _ = v44
                _ = v45
            }
            _ = fi24
        }
        v1 -= float64(float32(v1))
        _ = fi22
        _ = v23
    }
    var v46 int16 = 42
    for fi47 := 0; fi47 < 18; fi47++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        fmt.Println(v14)
        switch "" {
        case "foo bar":
            _ = len(v3)
        case "":
            fmt.Println(v14, v2)
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi47
    }
    if (v2 <= -9223372036854775808) {
        var v48 int16 = 1000
        _ = v48
    } else {
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 float32 = 1e10
            v51 := GenS5{F0: 18446744073709551615}
            var v52 bool = (true && (65535 < 42))
            _ = fi49
            _ = v50
            _ = v51
            _ = v52
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
