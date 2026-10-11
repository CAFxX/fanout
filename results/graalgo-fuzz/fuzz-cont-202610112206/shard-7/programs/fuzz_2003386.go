package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint; F1 int64; F2 bool }
func genFn0(p0 int16, p1 int16, p2 uint64) uint32 {
    v7 := GenS1{F0: 9223372036854775807, F1: 360}
    if false {
        fmt.Println(p0, p1)
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(p0, p2)
            p0 = p0
            p0 -= p1
            _ = fi8
        }
        p2 -= p2
        var v9 uint16 = 759
        _ = v9
    }
    switch 127 {
    case 0:
        v10 := GenS2{F0: 255, F1: 198}
        _ = v10
    case -9223372036854775808:
        switch ((("foo bar" + "") + "\n\t") + "hello") {
        case "foo bar":
            for fi11 := 0; fi11 < 9; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    v13 := GenS2{F0: 21, F1: 42}
                    fmt.Println(p0, fi11)
                    _ = fi12
                    _ = v13
                }
                _ = fi11
            }
        case "a":
            p1 -= p0
        case "\n\t":
            var v14 uint64 = 2
            _ = v14
        }
    case -255:
        if false {
            fmt.Println(p2, p0)
        }
    }
    var v15 int8 = 109
    switch 42 {
    case 128:
        p0 = 32767
    case -585:
        if (-909 == 827) {
            if false {
                v16 := GenS1{F0: -1, F1: 2147483647}
                _ = v16
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                p2 = (p2 | 18446744073709551615)
                _ = fi17
            }
            if ((488 != 2) || ((1e10 * -2.25) < 3.14159)) {
                var v18 uint64 = 1
                switch 4294967295 {
                case 32767:
                    p1 = (p1 << ((255) & 63))
                case 521:
                    var v19 float32 = (-2.25 + -2.25)
                    _ = v19
                default:
                    p2 += 4294967295
                }
                if (2 <= v18) {
                    for fi20 := 0; fi20 < 20; fi20++ {
                        for fi21 := 0; fi21 < 18; fi21++ {
                            v22 := GenS3{F0: false, F1: -1}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                    p0 -= (^(p1))
                    p0 = p0
                }
                _ = v18
            }
        }
    default:
        p1 += (p0 + (^(p0)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    var v2 bool = true
    var v3 int = -681
    var v4 uint64 = 42
    var v5 int16 = v1
    v6 := []int{-304}
    _ = len(v6)
    var v23 uint8 = 60
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2, v3)
    v6 = append(v6, v3)
    var v24 uint8 = v23
    var v25 uint64 = uint64(v24)
    fmt.Println(v1)
    v4 = v25
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS2{F0: v23, F1: (v23 % (((v24 ^ v24)) | 1))}
        v6 = append(v6, fi26)
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        if v2 {
            v3 = v3
        } else {
            switch "\n\t" {
            case "hello":
                if ((v1 ^ v5) < v5) {
                    fmt.Println(v5, v24, v25)
                }
            default:
                switch "" {
                case "foo bar":
                    if (((0.1 / 1.5) / 0.1) > 1e10) {
                        if true {
                            if (0 < 0) {
                                var v29 uint8 = (v23 * 255)
                                _ = v29
                            }
                        } else {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                var v31 int8 = 1
                                v23 -= v23
                                v32 := GenS3{F0: v2, F1: v31}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        for fi33 := 0; fi33 < 1; fi33++ {
                            v25 = 758
                            v2 = false
                            v34 := GenS1{F0: 1, F1: 128}
                            _ = fi33
                            _ = v34
                        }
                        switch fi28 {
                        case -2147483648:
                            if true {
                                var v35 uint64 = 0
                                if v2 {
                                    var v36 float64 = float64(-2.25)
                                    _ = v36
                                } else {
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        v25 = 18446744073709551615
                                        _ = fi37
                                    }
                                    _ = len(v6)
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        fmt.Println(v24, fi28, fi38)
                                        fmt.Println(v35)
                                        _ = fi38
                                    }
                                }
                                var v39 uint8 = (v24 * (^(v23)))
                                _ = v35
                                _ = v39
                            }
                        default:
                            for fi40 := 0; fi40 < 12; fi40++ {
                                var v41 uint32 = 65535
                                fmt.Println(v25, fi28)
                                if (!v2) {
                                    if len(v6) > 0 {
                                        v6[0] = 1
                                    }
                                    var v42 string = ((("a" + "") + "hello") + "\n\t")
                                    _ = len(v6)
                                    v43 := GenS4{F0: 0}
                                    _ = v42
                                    _ = v43
                                }
                                _ = fi40
                                _ = v41
                            }
                        }
                    }
                case "hello":
                    v4 += 18446744073709551615
                case "a":
                    for fi44 := 0; fi44 < 15; fi44++ {
                        fmt.Println(v3, v24)
                        var v45 bool = false
                        _ = fi44
                        _ = v45
                    }
                }
            }
            v5 -= v1
            var v46 uint = 858
            _ = v46
        }
        _ = fi28
    }
    var v47 uint64 = (v4 << ((65535) & 63))
    switch "\n\t" {
    case "\n\t":
        v23 -= (v24 + (v24 - 129))
    case "foo bar":
        var v48 int16 = v5
        _ = v48
    case "a":
        switch v3 {
        case 127:
            v49 := GenS1{F0: 318, F1: (v3 << ((0) & 63))}
            _ = v49
        }
    }
    v2 = (-1 != -1)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    switch "hello" {
    case "\n\t":
        v23 = 0
    default:
        if true {
            v6 = append(v6, 431)
            var v50 int = (v3 << ((663) & 63))
            v51 := GenS5{F0: 18446744073709551615, F1: 0, F2: (v2 && (!(((("" + "foo bar") + "a") < "\n\t") && v2)))}
            _ = v50
            _ = v51
        } else {
            v4 -= (^(v47))
            var v52 uint32 = 4294967295
            _ = v52
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
