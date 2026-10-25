package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 string; F1 string; F2 int32 }

func main() {
    v1 := []uint32{42, 2147483647}
    v2 := []int{-2147483648, 923}
    var v3 float32 = (((1.5 * float32((((0.0 / 3.14159) + 0.0) + (1.5 + (3.14159 * 1.5))))) - (1.5 + 1e10)) - 1.5)
    var v4 uint16 = 65535
    var v5 uint32 = 2
    var v6 uint16 = (v4 % (((^(v4))) | 1))
    if (995 >= -1) {
        switch "hello" {
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        case "a":
            fmt.Println(v5, v5, v6)
        case "hello":
            for fi7 := 0; fi7 < 4; fi7++ {
                for fi8 := 0; fi8 < 6; fi8++ {
                    switch (fi8 | fi7) {
                    case 32768:
                        v3 -= 3.14159
                    case -9223372036854775808:
                        v5 = (v5 ^ v5)
                    case 170:
                        switch "\n\t" {
                        case "a":
                            fmt.Println(v5)
                        default:
                            fmt.Println(v6, fi7)
                        }
                    }
                    v9 := GenS1{F0: 152, F1: (true || true)}
                    v10 := GenS2{F0: ((^(v5)) != (v5 / (((v5 / (((v5 % ((v5) | 1))) | 1))) | 1)))}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi7
            }
        }
        switch "" {
        case "\n\t":
            if (!(-9223372036854775808 > -488)) {
                fmt.Println(v3)
            }
        default:
            if false {
                for fi11 := 0; fi11 < 13; fi11++ {
                    fmt.Println(v3, v6, v3)
                    _ = fi11
                }
            } else {
                if (127 == -128) {
                    switch -485 {
                    case 128:
                        v4 -= (v6 << ((65535) & 63))
                    }
                }
                _ = len(v1)
                var v12 int = 66
                _ = v12
            }
        }
    }
    v13 := GenS3{F0: 97, F1: -254}
    if ((!((1.5 < (float64(v3) - (((((0.0 * 3.14159) + (0.1 - 1e10)) / 1.0) / 1.0) / 1e10))) || false)) || (0 != 514)) {
        v6 += (v4 & (v4 * 2))
    }
    var v14 bool = true
    for fi15 := 0; fi15 < 17; fi15++ {
        v4 -= v4
        for fi16 := 0; fi16 < 20; fi16++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            fmt.Println(v3, v3)
            var v17 int64 = -209
            v17 = v17
            fmt.Println(v14)
            if (!(510 > 18446744073709551615)) {
                var v18 int64 = 423
                switch (fi15 | 532) {
                case 32767:
                    var v19 int8 = -1
                    _ = v19
                case 414:
                    v1 = append(v1, (v5 / (((v5 + (v5 & uint32(v4)))) | 1)))
                }
                _ = v18
            }
            _ = fi16
            _ = v17
        }
        if v14 {
            fmt.Println(v14, v3, v3)
            fmt.Println(v3, v6, v5)
            v14 = (1 > 0)
        } else {
            var v20 int16 = 1000
            _ = v20
        }
        _ = fi15
    }
    if v14 {
        v4 = 804
        v21 := GenS4{F0: -128}
        _ = v21
    } else {
        if (!v14) {
            _ = len(v1)
        }
        fmt.Println(v4)
    }
    v1 = append(v1, 4294967295)
    if (v4 <= (v6 - (v4 % (((v6 | (v4 + 42))) | 1)))) {
        for fi22 := 0; fi22 < 19; fi22++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi22
        }
    } else {
        v23 := GenS1{F0: 32767, F1: (true && v14)}
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        v25 := GenS5{F0: ""}
        v6 -= 2
        _ = fi24
        _ = v25
    }
    if ((!(v14 || (!(-9223372036854775808 < -997)))) || v14) {
        var v26 int32 = 1
        _ = v26
    } else {
        for fi27 := 0; fi27 < 5; fi27++ {
            switch ((("a" + "foo bar") + "a") + "\n\t") {
            case "hello":
                var v28 int32 = -1
                _ = v28
            case "\n\t":
                v6 = uint16(v5)
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = v5
                }
            default:
                v4 -= v4
            }
            _ = fi27
        }
        for fi29 := 0; fi29 < 15; fi29++ {
            if len(v2) > 0 {
                v2[0] = (-(fi29))
            }
            v3 = 3.14159
            fmt.Println(fi29)
            _ = fi29
        }
    }
    v5 += v5
    if len(v2) > 0 {
        v2[0] = -651
    }
    var v30 uint8 = 42
    v31 := GenS6{F0: -1, F1: (float32(0.1) == (v3 + 3.14159))}
    v32 := GenS7{F0: "\n\t", F1: ("\n\t" + ((("" + "hello") + "foo bar") + (((("\n\t" + "") + "a") + (("hello" + (("hello" + "a") + "")) + "\n\t")) + ("" + ("hello" + "\n\t"))))), F2: -331}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
