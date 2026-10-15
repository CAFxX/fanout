package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 uint8; F1 int32; F2 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 uint16 }
type GenS4 struct { F0 uint }

func main() {
    var v1 bool = (9223372036854775807 <= 32767)
    var v2 uint32 = 2
    var v3 uint32 = (^(v2))
    var v4 int = -482
    var v5 float64 = (((0.0 + 3.14159) * (1e10 / 1.0)) / 1.5)
    var v6 float64 = v5
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: -1, F1: 0, F2: 2}
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int = (v4 ^ fi9)
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v4 = v4
    fmt.Println(v1, v5)
    var v11 int = (v4 * 1)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int32 = 0
        switch v11 {
        case 42:
            v2 = v3
        case 12:
            if ((2 < 42) && true) {
                v14 := []int{fi12, (^(fi12)), (fi12 | (v11 % ((-1) | 1)))}
                _ = v14
            } else {
                var v15 string = "foo bar"
                v16 := GenS2{F0: 143, F1: 1, F2: 1000}
                _ = v15
                _ = v16
            }
        case -9223372036854775808:
            v1 = (false && (172 <= (v2 << ((18446744073709551615) & 63))))
        default:
            v4 = fi12
        }
        _ = fi12
        _ = v13
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v3)
        var v18 int = fi17
        if (((^(v2)) <= v3) || (v1 && (-838 < -1))) {
            if (2 < 1) {
                if v1 {
                    for fi19 := 0; fi19 < 9; fi19++ {
                        if v1 {
                            var v20 uint = 65535
                            if (v1 && true) {
                                var v21 float64 = -2.25
                                v22 := []int{32768, (fi17 | (v11 ^ 128)), -1, (v11 ^ v4)}
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v21, v20, v11)
                                }
                                _ = v21
                                _ = v22
                            }
                            var v23 uint32 = v2
                            _ = v20
                            _ = v23
                        }
                        var v24 uint64 = 11
                        _ = fi19
                        _ = v24
                    }
                    var v25 int = 128
                    fmt.Println(v6, v18)
                    _ = v25
                } else {
                    v26 := GenS1{F0: 1, F1: 255, F2: 0}
                    v1 = (0.0 <= 1.5)
                    _ = v26
                }
                v27 := GenS3{F0: -9223372036854775808, F1: -2147483648, F2: 2}
                for fi28 := 0; fi28 < 7; fi28++ {
                    v29 := GenS2{F0: 42, F1: 2147483647, F2: 1000}
                    if false {
                        v30 := []int{v11, (v4 * v18)}
                        v31 := GenS4{F0: 896}
                        if len(v30) >= 2 {
                            v30 = v30[1:len(v30)]
                        }
                        _ = v30
                        _ = v31
                    } else {
                        var v32 float64 = v6
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        }
        _ = fi17
        _ = v18
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v4 -= fi33
        _ = fi33
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v3, v3)
    case "a":
        for fi34 := 0; fi34 < 18; fi34++ {
            v4 += fi34
            switch "\n\t" {
            case "\n\t":
                v11 = v11
            case "a":
                v35 := []int{(v11 % ((v11) | 1)), (v4 << ((18446744073709551615) & 63)), (fi34 + fi34)}
                _ = v35
            case "foo bar":
                var v36 uint32 = (v2 ^ v3)
                _ = v36
            default:
                var v37 int16 = 1000
                _ = v37
            }
            _ = fi34
        }
    case "\n\t":
        fmt.Println(v4, v1, v6)
    default:
        v3 = (v3 >> ((2) & 63))
    }
    var v38 uint = 398
    v39 := GenS4{F0: v38}
    for fi40 := 0; fi40 < 14; fi40++ {
        v41 := GenS1{F0: 23, F1: 255, F2: (v38 << ((v38) & 63))}
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        v6 -= v6
        _ = fi42
    }
    if v1 {
        fmt.Println(v5)
    }
    v43 := GenS1{F0: -68, F1: 127, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
