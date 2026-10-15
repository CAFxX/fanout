package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int16 }
func genFn0(p0 int16, p1 int64) int32 {
    v4 := []int{128}
    p0 += (p0 ^ p0)
    _ = p0
    _ = p1
    _ = v4
    return 106
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]bool)
    v3 := make(map[uint32]int)
    v5 := []int{-2147483648, -849, 32768, 2147483647}
    var v6 int32 = 90
    var v7 uint16 = 255
    fmt.Println(v1, v7, v7)
    v8 := GenS1{F0: (!true)}
    v7 = 335
    var v9 int64 = -837
    v10 := GenS1{F0: (894 <= 42)}
    v11 := GenS1{F0: ((v7 - (v7 + (v7 ^ v7))) >= 0)}
    if (-1 < -1) {
        v12 := GenS2{F0: 4294967295, F1: 42, F2: v9}
        switch 0 {
        case -708:
            var v13 uint8 = uint8(v7)
            _ = v13
        case 32767:
            if false {
                if (567 >= 42) {
                    switch "a" {
                    case "":
                        switch "a" {
                        case "a":
                            var v14 uint32 = 0
                            _ = v14
                        case "hello":
                            for fi15 := 0; fi15 < 6; fi15++ {
                                var v16 uint16 = (v7 & (v7 & 42))
                                v6 -= (v6 >> ((65535) & 63))
                                _ = fi15
                                _ = v16
                            }
                        case "foo bar":
                            for fi17 := 0; fi17 < 12; fi17++ {
                                if (3.14159 < 1.5) {
                                    var v18 float64 = (0.1 * -2.25)
                                    _ = v18
                                }
                                _ = fi17
                            }
                        default:
                            for fi19 := 0; fi19 < 12; fi19++ {
                                _ = len(v2)
                                if len(v5) > 0 {
                                    v5[0] = fi19
                                }
                                _ = fi19
                            }
                        }
                    case "a":
                        var v20 string = "a"
                        _ = v20
                    case "foo bar":
                        fmt.Println(v7, v1)
                    }
                }
                for fi21 := 0; fi21 < 16; fi21++ {
                    var v22 int = fi21
                    fmt.Println(v1, v7)
                    switch -766 {
                    case -2147483648:
                        for fi23 := 0; fi23 < 16; fi23++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            var v24 bool = true
                            v25 := v2[(^(v22))]
                            switch "" {
                            case "a":
                                _ = len(v5)
                            case "\n\t":
                                if (0 <= 0) {
                                    fmt.Println(v7, v9)
                                }
                            case "foo bar":
                                v24 = ((v9 | (v9 ^ v9)) > (v9 & 1))
                            }
                            v26 := GenS3{F0: "foo bar"}
                            _ = fi23
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    default:
                        v27 := GenS2{F0: 995, F1: v7, F2: 1}
                        _ = v27
                    }
                    _ = fi21
                    _ = v22
                }
                switch "" {
                case "a":
                    v2[-9223372036854775808] = (true && (((!(0 <= 2147483647)) || ((v6 & v6) <= (-(v6)))) && true))
                }
            }
        case 42:
            if (1 != 1) {
                v2[775] = (!((^(v1)) != -32768))
            }
        default:
            v9 -= (v9 & v9)
        }
        _ = v12
    } else {
        var v28 int16 = -465
        fmt.Println(v6)
        fmt.Println(v6, v1, v6)
        _ = v28
    }
    v1 = 42
    v29 := GenS4{F0: (v9 << ((18446744073709551615) & 63)), F1: v1}
    if (!(!false)) {
        v1 = 815
    }
    var v30 int64 = 9223372036854775807
    if len(v5) > 0 {
        v5[0] = -201
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
