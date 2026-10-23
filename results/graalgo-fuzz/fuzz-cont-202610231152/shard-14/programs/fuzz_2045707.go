package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 uint64 }

func main() {
    v1 := []uint32{1, 2147483648, 895}
    var v2 int16 = 1
    var v3 string = ("" + "a")
    var v4 uint = 0
    v4 -= (v4 + (v4 | v4))
    v5 := GenS1{F0: 2147483647}
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            v3 = (((((v3 + v3) + v3) + (v3 + v3)) + v3) + v3)
            _ = fi7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            v4 = 922
            _ = fi8
        }
        _ = fi6
    }
    v4 = (^(v4))
    var v9 uint32 = 4294967295
    v3 = v3
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 uint8 = 41
        if ("\n\t" > "foo bar") {
            var v12 float32 = -2.25
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 bool = ((v3 + v3) < (v3 + v3))
                v15 := GenS1{F0: (^(v9))}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        } else {
            if len(v1) > 0 {
                v1[0] = v9
            }
            for fi16 := 0; fi16 < 20; fi16++ {
                v3 = (v3 + v3)
                _ = fi16
            }
            if (!(((-185 <= -795) || (1.5 == ((-2.25 / (1e10 + (3.14159 * 3.14159))) - 0.1))) && false)) {
                v17 := GenS2{F0: v9}
                _ = v17
            }
        }
        if (!false) {
            for fi18 := 0; fi18 < 10; fi18++ {
                if true {
                    _ = len(v1)
                    v2 = (v2 | v2)
                    switch v3 {
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = v9
                        }
                    case "\n\t":
                        fmt.Println(v2, v9)
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v3 = v3
                    }
                }
                switch (fi10 >> ((v4) & 63)) {
                case 32768:
                    v3 = ""
                }
                _ = fi18
            }
        } else {
            fmt.Println(v9)
            for fi19 := 0; fi19 < 1; fi19++ {
                switch fi10 {
                case 9223372036854775807:
                    var v20 string = "a"
                    var v21 int32 = 724
                    if ((v9 + v9) > 2147483648) {
                        var v22 int32 = (v21 & (v21 / (((v21 / (((v21 | (v21 ^ (v21 << ((v4) & 63))))) | 1))) | 1)))
                        v23 := GenS2{F0: (^(v9))}
                        _ = len(v1)
                        _ = v22
                        _ = v23
                    }
                    _ = v20
                    _ = v21
                case 2147483647:
                    v2 = (v2 / ((-42) | 1))
                case 32768:
                    var v24 uint32 = (^(v9))
                    _ = v24
                default:
                    if true {
                        if len(v1) > 0 {
                            v1[0] = v9
                        }
                        if (v3 > "hello") {
                            v25 := GenS2{F0: 2147483647}
                            if (3.14159 < float32(float64((((1.5 * (((1.5 + float32(1e10)) + 0.0) * (float32(0.0) * (-2.25 / 1.0)))) + 1.5) * 0.1)))) {
                                v26 := GenS2{F0: 2147483648}
                                _ = v26
                            } else {
                                v27 := GenS3{F0: (v4 & v4)}
                                v4 = 460
                                v11 = v11
                                _ = v27
                            }
                            if len(v1) > 0 {
                                v1[0] = (v9 * (v9 | 2))
                            }
                            _ = v25
                        }
                        var v28 int8 = 92
                        _ = v28
                    }
                }
                _ = fi19
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                if ((65535 <= (^(v9))) && true) {
                    var v30 int8 = -25
                    v2 = v2
                    _ = len(v1)
                    _ = v30
                } else {
                    if len(v1) > 0 {
                        v1[0] = v9
                    }
                    var v31 int16 = (v2 >> ((v4) & 63))
                    var v32 uint = v4
                    _ = v31
                    _ = v32
                }
                _ = fi29
            }
        }
        _ = fi10
        _ = v11
    }
    var v33 bool = (0.0 < 0.1)
    var v34 int = -824
    if true {
        if true {
            switch v34 {
            case 2147483647:
                var v35 int16 = v2
                _ = v35
            case 42:
                switch v34 {
                case 127:
                    switch v34 {
                    case -6:
                        for fi36 := 0; fi36 < 19; fi36++ {
                            var v37 int8 = 1
                            v34 = fi36
                            _ = fi36
                            _ = v37
                        }
                    case -9223372036854775808:
                        v1 = append(v1, v9)
                    case 557:
                        var v38 bool = (!(v33 || v33))
                        _ = v38
                    }
                case 1:
                    fmt.Println(v33, v4)
                }
            }
            switch (v3 + v3) {
            case "hello":
                var v39 uint16 = 255
                _ = v39
            default:
                var v40 uint64 = 4294967295
                _ = v40
            }
        } else {
            if v33 {
                var v41 int64 = int64(v34)
                _ = v41
            }
            var v42 int = 0
            var v43 int32 = -591
            _ = v42
            _ = v43
        }
    }
    switch v3 {
    case "\n\t":
        var v44 uint64 = 461
        _ = v44
    case "hello":
        var v45 float64 = 3.14159
        _ = v45
    case "":
        v2 = v2
    default:
        v46 := GenS4{F0: 0, F1: 0, F2: -62}
        _ = v46
    }
    switch v34 {
    case -638:
        v9 += 2
    case 0:
        for fi47 := 0; fi47 < 20; fi47++ {
            if (!v33) {
                if (361 > 4294967295) {
                    var v48 uint8 = 128
                    _ = v48
                }
            } else {
                switch (v34 >> ((v4) & 63)) {
                case -2147483648:
                    for fi49 := 0; fi49 < 3; fi49++ {
                        var v50 int32 = -676
                        _ = fi49
                        _ = v50
                    }
                }
                var v51 uint8 = 2
                if (fi47 != (fi47 * v34)) {
                    v52 := GenS5{F0: 4294967295, F1: 605, F2: uint64(v9)}
                    for fi53 := 0; fi53 < 9; fi53++ {
                        v51 = v51
                        _ = fi53
                    }
                    var v54 uint = (v4 % ((v4) | 1))
                    _ = v52
                    _ = v54
                }
                _ = v51
            }
            v4 = (^(v4))
            switch 32 {
            case 128:
                v55 := GenS3{F0: v4}
                _ = v55
            case -592:
                v1 = append(v1, v9)
            }
            _ = fi47
        }
    default:
        var v56 uint64 = 462
        _ = v56
    }
    v1 = append(v1, v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
