package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 int8 }
func genFn0(p0 uint, p1 uint8, p2 uint64) int {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            if true {
                var v7 uint64 = uint64(p0)
                v8 := GenS1{F0: v7, F1: 859}
                _ = v7
                _ = v8
            }
            p0 += 5
            _ = fi6
        }
        var v9 uint64 = 1
        _ = fi5
        _ = v9
    }
    v10 := GenS2{F0: -1}
    switch ("a" + "foo bar") {
    case "a":
        if (300 <= 1) {
            var v11 int8 = 1
            var v12 int = 127
            _ = v11
            _ = v12
        }
    case "foo bar":
        var v13 int64 = 9223372036854775807
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(p0, fi14, p2)
            fmt.Println(p1, p1)
            _ = fi14
        }
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS3{F0: -1, F1: 2147483647, F2: 0}
            for fi17 := 0; fi17 < 14; fi17++ {
                var v18 int32 = 0
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
    }
    v19 := []int{174, 127, 42}
    for fi20 := 0; fi20 < 18; fi20++ {
        p1 = p1
        p2 -= p2
        p1 = (^(p1))
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v19
    return int(p1)
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = 128
    var v3 int8 = 100
    var v4 float64 = 1e10
    var v21 string = "foo bar"
    v2 -= v2
    v22 := []int{genFn0(1, 88, 2)}
    v21 = "foo bar"
    fmt.Println(v3, v1)
    for fi23 := 0; fi23 < 4; fi23++ {
        fmt.Println(v3)
        _ = fi23
    }
    if ("\n\t" > v21) {
        if true {
            if len(v22) > 0 {
                v22[0] = -2147483648
            }
            v22 = append(v22, (v2 + v2))
        }
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        switch ("foo bar" + v21) {
        case "hello":
            switch (v2 / (((-(v2))) | 1)) {
            case -763:
                switch v21 {
                case "\n\t":
                    v1 -= 664
                }
            default:
                if (0 != 0) {
                    v21 = ("foo bar" + v21)
                    v24 := GenS2{F0: 42}
                    v2 -= genFn0(811, 178, v1)
                    _ = v24
                }
            }
        }
    }
    var v25 bool = (42 <= v2)
    v26 := GenS3{F0: 1, F1: 736, F2: -1}
    v1 = v1
    v27 := GenS4{F0: v3, F1: -1}
    switch (v2 & v2) {
    case -204:
        var v28 int16 = 0
        _ = v28
    default:
        if len(v22) > 0 {
            v22[0] = 4294967295
        }
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        v30 := GenS2{F0: -218}
        _ = fi29
        _ = v30
    }
    if (1e10 > 1.5) {
        var v31 int = 32768
        _ = v31
    }
    if (-32768 != 42) {
        v4 += v4
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 1; fi33++ {
            v3 -= (v3 | v3)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            _ = fi33
        }
        var v34 string = ""
        if true {
            var v35 uint8 = 81
            switch v21 {
            case "":
                for fi36 := 0; fi36 < 6; fi36++ {
                    fmt.Println(v34, v34, v4)
                    if ((float32(0.0) * 1e10) >= -2.25) {
                        if false {
                            _ = len(v22)
                            var v37 bool = ((^(v1)) <= v1)
                            _ = v37
                        }
                        var v38 uint = 437
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        _ = v38
                    } else {
                        var v39 bool = (2 != 32767)
                        _ = v39
                    }
                    _ = fi36
                }
            case "hello":
                switch v34 {
                case "hello":
                    v1 -= (v1 ^ (v1 * v1))
                default:
                    switch (v34 + "hello") {
                    case "a":
                        fmt.Println(v25, v1)
                    case "foo bar":
                        if v25 {
                            var v40 int = genFn0(255, v35, 64)
                            switch fi32 {
                            case 32768:
                                var v41 uint8 = (v35 % ((255) | 1))
                                _ = v41
                            case 9223372036854775807:
                                for fi42 := 0; fi42 < 9; fi42++ {
                                    v40 = (v2 ^ (fi32 >> ((18446744073709551615) & 63)))
                                    var v43 int16 = 1
                                    v4 -= -2.25
                                    if v25 {
                                        var v44 int8 = 42
                                        if v25 {
                                            if len(v22) > 0 {
                                                v22[0] = 32767
                                            }
                                            v44 = -93
                                            fmt.Println(v25, v44)
                                        } else {
                                            v25 = (!(2 != 42))
                                        }
                                        _ = v44
                                    } else {
                                        switch v34 {
                                        case "a":
                                            var v45 int = v40
                                            _ = v45
                                        case "foo bar":
                                            if (float64(3.14159) <= 3.14159) {
                                                v25 = ("" == "foo bar")
                                            }
                                        }
                                        var v46 int32 = -1
                                        _ = v46
                                    }
                                    _ = len(v22)
                                    _ = fi42
                                    _ = v43
                                }
                            default:
                                fmt.Println(v4)
                            }
                            v35 -= (v35 % ((v35) | 1))
                            _ = v40
                        } else {
                            switch 32767 {
                            case 2147483647:
                                fmt.Println(v34, fi32)
                            default:
                                v21 = v34
                            }
                        }
                    case "\n\t":
                        fmt.Println(v35, v34)
                    }
                }
            }
            var v47 uint16 = 255
            _ = v35
            _ = v47
        }
        _ = fi32
        _ = v34
    }
    v3 -= (-(v3))
    switch v2 {
    case -9223372036854775808:
        var v48 int64 = -376
        _ = v48
    case 127:
        for fi49 := 0; fi49 < 2; fi49++ {
            v25 = true
            _ = fi49
        }
    default:
        v50 := GenS2{F0: int64(v1)}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
