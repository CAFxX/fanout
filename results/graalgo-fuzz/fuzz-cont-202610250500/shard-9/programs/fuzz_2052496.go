package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int; F2 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int32; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint8) bool {
    v7 := []int{42, 469}
    var v8 float32 = -2.25
    if (!(!(-9223372036854775808 != 9223372036854775807))) {
        for fi9 := 0; fi9 < 1; fi9++ {
            if len(v7) > 0 {
                v7[0] = 32768
            }
            var v10 bool = false
            _ = fi9
            _ = v10
        }
    } else {
        v8 -= v8
        var v11 uint32 = 745
        if true {
            var v12 uint64 = (p0 << ((2) & 63))
            p2 = (p2 % ((p2) | 1))
            _ = v12
        }
        _ = v11
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if ((255 < p1) || true) {
        fmt.Println(p2)
        p0 -= p0
        var v13 bool = true
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return false
}
func genFn1() int32 {
    var v14 int64 = 1
    if false {
        if (true || (!true)) {
            switch "a" {
            case "foo bar":
                v14 = 0
            case "\n\t":
                var v15 int64 = v14
                _ = v15
            case "hello":
                fmt.Println(v14)
            }
            var v16 uint16 = 0
            v17 := GenS1{F0: 250, F1: 29}
            _ = v16
            _ = v17
        } else {
            if true {
                v14 += (v14 + (^(v14)))
            } else {
                switch "" {
                case "foo bar":
                    v14 = -479
                }
            }
            fmt.Println(v14)
            var v18 uint16 = 329
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            if (32767 > 42) {
                var v20 int = (fi19 ^ (^(fi19)))
                _ = v20
            } else {
                v14 = v14
            }
            _ = fi19
        }
    }
    _ = v14
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 string = "hello"
    var v3 int64 = 0
    v4 := []uint32{2}
    var v5 float32 = 1e10
    var v6 uint64 = 332
    var v21 uint16 = 710
    if (((v3 >> ((18446744073709551615) & 63)) <= 1) || (((^(v6)) == v6) && (1 == 255))) {
        v22 := GenS1{F0: 42, F1: 42}
        for fi23 := 0; fi23 < 8; fi23++ {
            v24 := GenS1{F0: 42, F1: 50}
            if (((v2 > (v2 + v2)) || true) || (!true)) {
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                for fi25 := 0; fi25 < 5; fi25++ {
                    v2 = v2
                    if (v3 >= v3) {
                        switch (v2 + ((v2 + "a") + v2)) {
                        case "a":
                            v26 := GenS2{F0: (fi25 & fi23)}
                            fmt.Println(v5, v5, fi25)
                            _ = v26
                        default:
                            v21 += (^(v21))
                        }
                        if (2147483647 < 197) {
                            var v27 float64 = 0.1
                            _ = v27
                        }
                    }
                    switch "" {
                    case "hello":
                        var v28 int32 = 967
                        _ = v28
                    }
                    _ = fi25
                }
                var v29 int64 = 0
                _ = v29
            } else {
                var v30 float32 = (((v5 + (((-2.25 / ((((v5 - v5) + v5) - 0.1) - ((0.0 * v5) * v5))) - v5) - 0.1)) - v5) + 0.1)
                _ = v30
            }
            v6 = v6
            _ = fi23
            _ = v24
        }
        switch "a" {
        case "hello":
            var v31 bool = false
            _ = v31
        case "":
            var v32 int32 = 2147483647
            _ = v32
        case "\n\t":
            switch (v2 + v2) {
            case "\n\t":
                for fi33 := 0; fi33 < 9; fi33++ {
                    v34 := GenS3{F0: (true && (!(1.5 <= ((-2.25 * (((0.0 * (0.1 * 1e10)) * (1.5 / 0.1)) + (0.0 * 1e10))) * 3.14159)))), F1: int(v21), F2: 318}
                    if false {
                        if (fi33 >= (fi33 % (((fi33 << ((4294967295) & 63))) | 1))) {
                            _ = len(v4)
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch fi33 {
                                case -2147483648:
                                    v35 := GenS4{F0: 42}
                                    _ = v35
                                default:
                                    switch v2 {
                                    case "\n\t":
                                        if true {
                                            var v36 uint16 = (v21 & (^(v21)))
                                            _ = v36
                                        }
                                    case "":
                                        fmt.Println(v3)
                                    }
                                }
                            } else {
                                _ = len(v4)
                                v5 = 1e10
                            }
                            v37 := GenS4{F0: int32(v21)}
                            _ = v37
                        } else {
                            if false {
                                v38 := GenS2{F0: (fi33 + (fi33 >> ((69) & 63)))}
                                _ = v38
                            } else {
                                fmt.Println(v6, v3)
                                switch (fi33 + -2147483648) {
                                case -2147483648:
                                    var v39 uint16 = 255
                                    _ = v39
                                case -353:
                                    for fi40 := 0; fi40 < 10; fi40++ {
                                        switch ("hello" + "a") {
                                        case "\n\t":
                                            var v41 float64 = -2.25
                                            _ = v41
                                        }
                                        _ = fi40
                                    }
                                default:
                                    v2 = v2
                                }
                            }
                            var v42 float32 = v5
                            var v43 int32 = 1
                            _ = v42
                            _ = v43
                        }
                    }
                    var v44 string = (("a" + "a") + ("\n\t" + v2))
                    _ = fi33
                    _ = v34
                    _ = v44
                }
            case "foo bar":
                v45 := GenS5{F0: 1, F1: 71, F2: -125}
                _ = v45
            case "a":
                var v46 uint8 = 255
                _ = v46
            default:
                v2 = v2
            }
        }
        _ = v22
    }
    v2 = (v2 + v2)
    var v47 int8 = -67
    if true {
        var v48 uint = 369
        v4 = append(v4, 2)
        _ = v48
    }
    v49 := GenS5{F0: v3, F1: -2147483648, F2: -1}
    v6 -= 976
    v50 := GenS6{F0: 65535, F1: 0, F2: 883}
    for fi51 := 0; fi51 < 9; fi51++ {
        v6 += v6
        v5 += 0.0
        _ = fi51
    }
    var v52 float64 = ((-2.25 / 3.14159) / (1e10 * ((1e10 + 0.1) / -2.25)))
    v53 := v1[((((("foo bar" + v2) + v2) + v2) + v2) + (v2 + v2))]
    for fi54 := 0; fi54 < 15; fi54++ {
        if ((v47 ^ 100) > (v47 / (((v47 / ((v47) | 1))) | 1))) {
            v55 := GenS4{F0: -2147483648}
            var v56 int = fi54
            var v57 uint8 = 204
            _ = v55
            _ = v56
            _ = v57
        }
        fmt.Println(fi54, v6, v5)
        _ = fi54
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v49
    _ = v5
    _ = v50
    _ = v52
    _ = v53
    _ = v6
}
