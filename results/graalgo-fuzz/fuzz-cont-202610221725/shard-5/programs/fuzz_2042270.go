package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []int{4294967295, 0, 1, -1}
    var v2 float32 = -2.25
    var v3 uint8 = 2
    var v4 float64 = -2.25
    v5 := []uint32{5, 2, 2, 309}
    var v6 uint32 = 143
    for fi7 := 0; fi7 < 13; fi7++ {
        v3 += (v3 | 0)
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        v9 := GenS1{F0: 241, F1: (v6 << ((18446744073709551615) & 63))}
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        if false {
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 string = "a"
                fmt.Println(v12, v3, fi10)
                _ = fi11
                _ = v12
            }
            v2 = float32((((v4 * v4) / 1e10) * float64(v2)))
            for fi13 := 0; fi13 < 15; fi13++ {
                switch (^(fi10)) {
                case 1:
                    if false {
                        var v14 int8 = -85
                        _ = v14
                    }
                case 127:
                    if ((((3.14159 / v2) + (3.14159 + v2)) + v2) <= 3.14159) {
                        var v15 int32 = 1
                        if len(v5) > 0 {
                            v5[0] = (v6 % (((v6 % (((v6 + (^(v6)))) | 1))) | 1))
                        }
                        var v16 bool = false
                        _ = v15
                        _ = v16
                    }
                default:
                    _ = len(v5)
                }
                var v17 int64 = 0
                _ = fi13
                _ = v17
            }
        } else {
            var v18 uint16 = 255
            switch fi10 {
            case -2147483648:
                v4 = v4
            case 4294967295:
                var v19 int = (^(fi10))
                _ = v19
            case -1:
                v5 = append(v5, 1)
            default:
                fmt.Println(v18, v2, v3)
            }
            _ = v18
        }
        if (42 >= 42) {
            fmt.Println(v4)
        }
        var v20 float32 = float32(((-2.25 + 0.0) * v4))
        _ = fi10
        _ = v20
    }
    var v21 int = 32768
    if (-103 >= -128) {
        var v22 uint32 = v6
        for fi23 := 0; fi23 < 11; fi23++ {
            if (2 > 0) {
                v22 += v22
            } else {
                v22 += v6
            }
            fmt.Println(v6)
            _ = fi23
        }
        v2 = v2
        _ = v22
    }
    v1 = append(v1, (v21 >> ((65535) & 63)))
    v21 += (v21 >> ((211) & 63))
    if len(v5) > 0 {
        v5[0] = v6
    }
    if ((v3 == 156) || false) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi24 := 0; fi24 < 18; fi24++ {
            for fi25 := 0; fi25 < 11; fi25++ {
                var v26 string = ((("foo bar" + ("a" + "a")) + "a") + "hello")
                switch "" {
                case "":
                    var v27 uint64 = 118
                    _ = v27
                default:
                    var v28 float64 = float64(0.0)
                    _ = v28
                }
                v29 := GenS2{F0: (v6 * v6), F1: (fi25 * (-(v21)))}
                _ = fi25
                _ = v26
                _ = v29
            }
            _ = fi24
        }
    } else {
        _ = len(v1)
    }
    if (true || (v4 <= v4)) {
        if len(v5) > 0 {
            v5[0] = (v6 << ((18446744073709551615) & 63))
        }
        var v30 uint8 = 255
        for fi31 := 0; fi31 < 6; fi31++ {
            for fi32 := 0; fi32 < 6; fi32++ {
                switch (fi32 | (^(fi32))) {
                case -219:
                    _ = len(v1)
                case -704:
                    switch -717 {
                    case 2147483647:
                        switch "\n\t" {
                        case "hello":
                            var v33 float32 = 3.14159
                            _ = v33
                        case "":
                            var v34 int = (fi31 - -433)
                            _ = v34
                        default:
                            var v35 int8 = 100
                            _ = v35
                        }
                    case -701:
                        v3 -= v3
                        v6 = (v6 ^ v6)
                        if (71 != v6) {
                            var v36 int16 = 76
                            v5 = append(v5, (v6 << ((215) & 63)))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v36
                        }
                    }
                case -335:
                    switch "a" {
                    case "foo bar":
                        var v37 uint16 = 255
                        _ = v37
                    case "hello":
                        var v38 string = ("hello" + "")
                        var v39 int32 = int32(v6)
                        v40 := GenS3{F0: 42}
                        v39 += -1
                        var v41 int = fi32
                        var v42 float64 = -2.25
                        _ = v38
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    default:
                        if ((!false) && (0.1 != ((v4 - ((-2.25 * v4) + v4)) + float64(v2)))) {
                            fmt.Println(v2, v6)
                        }
                    }
                default:
                    v6 = 42
                    if len(v5) > 0 {
                        v5[0] = (^(v6))
                    }
                }
                _ = fi32
            }
            v1 = append(v1, 0)
            fmt.Println(v2, v3)
            _ = fi31
        }
        _ = v30
    }
    if (622 <= int16(v21)) {
        fmt.Println(v4, v21, v2)
    }
    if true {
        var v43 float32 = -2.25
        v44 := GenS3{F0: 0}
        _ = v43
        _ = v44
    } else {
        if (((!(53 == (v6 * (v6 | v6)))) && (550 == 1)) || (0.0 <= (v4 + 1.5))) {
            v1 = append(v1, -1)
            for fi45 := 0; fi45 < 7; fi45++ {
                if false {
                    if (18446744073709551615 != uint64(v6)) {
                        for fi46 := 0; fi46 < 7; fi46++ {
                            var v47 int8 = 1
                            _ = fi46
                            _ = v47
                        }
                        var v48 int = 32768
                        _ = v48
                    } else {
                        if len(v5) > 0 {
                            v5[0] = v6
                        }
                    }
                    for fi49 := 0; fi49 < 14; fi49++ {
                        var v50 string = "\n\t"
                        v3 += (v3 << ((0) & 63))
                        v51 := GenS3{F0: (v3 | (v3 >> ((2) & 63)))}
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    var v52 int8 = 69
                    _ = v52
                } else {
                    v53 := GenS2{F0: v6, F1: v21}
                    _ = v53
                }
                fmt.Println(v21, v4)
                _ = fi45
            }
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        }
        switch (v21 ^ (v21 ^ 0)) {
        case 9223372036854775807:
            if (0 != 32767) {
                var v54 string = ""
                var v55 int64 = 1
                v1 = append(v1, (v21 | v21))
                _ = v54
                _ = v55
            }
        case -1:
            var v56 uint32 = v6
            _ = v56
        default:
            for fi57 := 0; fi57 < 9; fi57++ {
                var v58 int16 = 32767
                v58 -= (-(v58))
                _ = fi57
                _ = v58
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
