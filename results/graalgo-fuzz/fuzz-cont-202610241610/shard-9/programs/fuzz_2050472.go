package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint8; F1 string; F2 uint32 }

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 int8 = 40
    var v4 int16 = -669
    var v5 uint16 = 0
    var v6 uint32 = 0
    var v7 bool = (true && (!(1 != 128)))
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := GenS1{F0: (-(fi8))}
        _ = len(v1)
        _ = fi8
        _ = v9
    }
    if (((128 == 218) || false) && (v4 > -838)) {
        if (!false) {
            v10 := GenS2{F0: 184}
            _ = len(v2)
            var v11 uint8 = 68
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            if ((false && (65535 < v6)) && (!(((!(127 != -128)) && (((v3 * 36) >= 100) || (v7 && v7))) || (!(1 != 2))))) {
                if v7 {
                    v5 = 833
                    if (((3.14159 / (3.14159 - 0.1)) * 3.14159) <= (0.1 * float32((1.5 - (-2.25 / (1e10 + 0.0)))))) {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            v6 = 268
                            v14 := GenS1{F0: (fi12 & int(v5))}
                            if false {
                                switch fi12 {
                                case 32768:
                                    v6 += 606
                                case -897:
                                    v7 = true
                                case -1:
                                    v3 -= v3
                                }
                                v6 = v6
                            } else {
                                var v15 uint32 = (v6 * 2147483648)
                                v4 = 32767
                                if v7 {
                                    fmt.Println(v7, v6, v15)
                                    if (v5 < 694) {
                                        var v16 int8 = v3
                                        v4 = v4
                                        _ = v16
                                    }
                                    var v17 int64 = 9223372036854775807
                                    _ = v17
                                } else {
                                    if (26 < 0) {
                                        v18 := GenS2{F0: 128}
                                        v7 = v7
                                        switch "hello" {
                                        case "":
                                            _ = len(v2)
                                        case "hello":
                                            var v19 uint64 = 597
                                            _ = v19
                                        }
                                        _ = v18
                                    } else {
                                        v20 := GenS2{F0: 127}
                                        v7 = (3.14159 == float32(((float64(0.0) - 1e10) + 3.14159)))
                                        v21 := GenS1{F0: fi13}
                                        _ = v20
                                        _ = v21
                                    }
                                }
                                _ = v15
                            }
                            _ = fi13
                            _ = v14
                        }
                        var v22 float32 = 0.0
                        _ = v22
                    }
                }
            } else {
                if true {
                    var v23 int32 = 0
                    v24 := GenS1{F0: int(v5)}
                    _ = v23
                    _ = v24
                }
                v25 := GenS3{F0: 18446744073709551615, F1: "a"}
                fmt.Println(v5)
                _ = v25
            }
            v5 -= 255
            var v26 uint = 65535
            _ = fi12
            _ = v26
        }
        if v7 {
            if (0 > 128) {
                for fi27 := 0; fi27 < 14; fi27++ {
                    fmt.Println(v3, fi27, v3)
                    _ = len(v2)
                    var v28 uint64 = 0
                    _ = fi27
                    _ = v28
                }
                fmt.Println(v4, v7)
            }
            var v29 int16 = v4
            _ = v29
        } else {
            v6 -= (v6 ^ (v6 ^ 0))
            v30 := GenS2{F0: 255}
            _ = v30
        }
    } else {
        if v7 {
            _ = len(v2)
        } else {
            switch 32767 {
            case 32767:
                var v31 int64 = 0
                _ = v31
            case -227:
                var v32 float64 = -2.25
                _ = v32
            default:
                v33 := v2["hello"]
                _ = v33
            }
            fmt.Println(v6, v5, v5)
        }
        if (!v7) {
            var v34 int = 32768
            _ = v34
        } else {
            if v7 {
                v35 := GenS3{F0: 4294967295, F1: (("foo bar" + ("\n\t" + "hello")) + ("foo bar" + "hello"))}
                var v36 uint16 = 255
                for fi37 := 0; fi37 < 12; fi37++ {
                    fmt.Println(v4)
                    _ = fi37
                }
                _ = v35
                _ = v36
            }
            v38 := GenS4{F0: 255, F1: ("\n\t" + "hello"), F2: v6}
            _ = v38
        }
        var v39 int16 = (v4 + (v4 | 1))
        _ = v39
    }
    for fi40 := 0; fi40 < 19; fi40++ {
        v41 := GenS1{F0: (fi40 ^ (fi40 * fi40))}
        v42 := GenS2{F0: 79}
        for fi43 := 0; fi43 < 1; fi43++ {
            var v44 int64 = 42
            v7 = (42 == v5)
            var v45 int16 = 521
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi40
        _ = v41
        _ = v42
    }
    fmt.Println(v6, v5)
    v46 := GenS4{F0: 127, F1: ((("a" + "a") + "foo bar") + ("a" + "")), F2: v6}
    var v47 int8 = -65
    for fi48 := 0; fi48 < 14; fi48++ {
        var v49 int16 = 32767
        var v50 int8 = -1
        _ = fi48
        _ = v49
        _ = v50
    }
    v4 = 414
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
