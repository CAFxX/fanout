package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint16; F2 bool }
type GenS6 struct { F0 uint16; F1 uint64; F2 int32 }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]int)
    v3 := []int{114, 32767}
    var v4 int32 = 0
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6, v1, v1)
            if (42 < 32767) {
                var v7 uint64 = 18446744073709551615
                v8 := GenS1{F0: 42, F1: "", F2: 2}
                _ = v7
                _ = v8
            }
            _ = fi6
        }
        v4 = (v4 + (v4 % ((v4) | 1)))
        var v9 int8 = -46
        _ = fi5
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        switch fi10 {
        case 127:
            for fi11 := 0; fi11 < 3; fi11++ {
                if true {
                    v4 = v4
                    v1 = -475
                }
                v12 := GenS1{F0: 2, F1: "hello", F2: 18446744073709551615}
                var v13 uint = 255
                v14 := GenS2{F0: (fi10 & fi10), F1: "foo bar", F2: 667}
                v15 := GenS2{F0: fi10, F1: "hello", F2: 1}
                v4 -= 0
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
        }
        v4 = -864
        _ = fi10
    }
    if (1 > 42) {
        var v16 uint64 = 2
        if (-2.25 < (1.5 - -2.25)) {
            v17 := v2[("" + "\n\t")]
            _ = v17
        } else {
            switch -1 {
            case 9223372036854775807:
                v18 := GenS2{F0: 32768, F1: ("foo bar" + ("foo bar" + ("\n\t" + "hello"))), F2: 2147483648}
                _ = v18
            }
            v16 += v16
        }
        v19 := GenS2{F0: 9223372036854775807, F1: "a", F2: 298}
        _ = v16
        _ = v19
    } else {
        var v20 uint16 = 142
        var v21 uint8 = 26
        for fi22 := 0; fi22 < 8; fi22++ {
            delete(v2, "foo bar")
            var v23 float64 = 1.5
            v23 -= v23
            _ = fi22
            _ = v23
        }
        _ = v20
        _ = v21
    }
    v24 := GenS2{F0: 2147483647, F1: "a", F2: 44}
    for fi25 := 0; fi25 < 12; fi25++ {
        var v26 int32 = v4
        _ = len(v2)
        var v27 bool = true
        _ = fi25
        _ = v26
        _ = v27
    }
    v4 = (v4 + v4)
    v28 := GenS3{F0: 0}
    var v29 uint = 0
    for fi30 := 0; fi30 < 5; fi30++ {
        for fi31 := 0; fi31 < 16; fi31++ {
            v32 := GenS1{F0: 18446744073709551615, F1: ("" + "hello"), F2: (v29 * (v29 % ((v29) | 1)))}
            fmt.Println(fi31, fi31, fi31)
            delete(v2, ("" + ("foo bar" + ("hello" + "hello"))))
            _ = fi31
            _ = v32
        }
        var v33 int64 = v1
        v34 := GenS4{F0: 106, F1: v33}
        _ = fi30
        _ = v33
        _ = v34
    }
    var v35 int8 = -118
    var v36 uint32 = 65535
    for fi37 := 0; fi37 < 8; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            if ((2 < 42) || false) {
                if false {
                    var v39 int = (fi37 / (((fi38 % ((fi38) | 1))) | 1))
                    _ = v39
                } else {
                    var v40 int16 = int16(fi38)
                    var v41 int16 = (v40 & (v40 & -1))
                    _ = v40
                    _ = v41
                }
                var v42 float32 = (float32((1.5 - float64((((1.5 + 0.1) / 1.5) / 1.0)))) - 0.1)
                _ = v42
            } else {
                var v43 int = 9223372036854775807
                _ = v43
            }
            if true {
                fmt.Println(v1, v4, v1)
                delete(v2, ("hello" + "a"))
            }
            v35 = (v35 % (((v35 & (v35 ^ (v35 | v35)))) | 1))
            _ = fi38
        }
        _ = fi37
    }
    if ((v36 % (((v36 >> ((255) & 63))) | 1)) != 2147483648) {
        for fi44 := 0; fi44 < 7; fi44++ {
            v45 := GenS2{F0: 9223372036854775807, F1: "hello", F2: (v36 + (v36 & (v36 / ((v36) | 1))))}
            if (true || (!false)) {
                var v46 int32 = v4
                v47 := GenS2{F0: fi44, F1: ("" + "hello"), F2: v36}
                for fi48 := 0; fi48 < 19; fi48++ {
                    switch fi48 {
                    case -130:
                        v49 := GenS1{F0: 876, F1: ("\n\t" + (("\n\t" + "\n\t") + ("\n\t" + "a"))), F2: v29}
                        _ = v49
                    default:
                        v50 := GenS4{F0: 2147483648, F1: (v1 % (((v1 | (v1 << (((v29 / ((v29) | 1))) & 63)))) | 1))}
                        _ = v50
                    }
                    if true {
                        var v51 uint64 = 18446744073709551615
                        fmt.Println(v51, v51, fi48)
                        v3 = append(v3, fi44)
                        _ = v51
                    } else {
                        var v52 uint64 = 1
                        var v53 int64 = v1
                        v3 = append(v3, 42)
                        var v54 int32 = v46
                        if (v35 < 1) {
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                        } else {
                            if false {
                                var v55 uint8 = 127
                                v36 -= 0
                                v36 = 283
                                _ = v55
                            }
                        }
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    if ((!false) || (int8(fi44) > 127)) {
                        fmt.Println(v1, v36, v29)
                        _ = len(v3)
                        v3 = append(v3, fi44)
                    } else {
                        v46 += -2147483648
                        v56 := GenS5{F0: 1, F1: 255, F2: true}
                        _ = v56
                    }
                    _ = fi48
                }
                _ = v46
                _ = v47
            } else {
                fmt.Println(v36, v36, v35)
                switch "a" {
                case "a":
                    if true {
                        if false {
                            var v57 uint = v29
                            for fi58 := 0; fi58 < 17; fi58++ {
                                v36 = (v36 / (((v36 - v36)) | 1))
                                var v59 int = fi58
                                _ = len(v2)
                                var v60 string = "a"
                                _ = fi58
                                _ = v59
                                _ = v60
                            }
                            _ = v57
                        }
                        var v61 int16 = 0
                        v62 := GenS4{F0: (v36 * v36), F1: (v1 % ((v1) | 1))}
                        _ = v61
                        _ = v62
                    }
                default:
                    var v63 uint16 = 255
                    _ = v63
                }
                v29 = (v29 + uint(v36))
            }
            _ = fi44
            _ = v45
        }
        if true {
            if true {
                var v64 bool = ((uint8(v36) == 42) && ((v1 << ((1) & 63)) > v1))
                _ = v64
            }
            for fi65 := 0; fi65 < 11; fi65++ {
                v35 = v35
                if true {
                    _ = len(v3)
                    v66 := GenS1{F0: 42, F1: ("" + ("\n\t" + "\n\t")), F2: v29}
                    fmt.Println(v29, fi65)
                    _ = v66
                }
                _ = fi65
            }
        }
    } else {
        var v67 bool = (uint8(v29) != 6)
        _ = v67
    }
    v68 := GenS6{F0: 615, F1: 4294967295, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v68
}
