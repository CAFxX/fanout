package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 uint32 }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int64 = 1
    v2 := []uint32{0, 383, 1, 327}
    var v3 uint = 39
    if false {
        _ = len(v2)
        if (!(42 >= 127)) {
            switch "" {
            case "":
                var v4 bool = (!true)
                _ = v4
            case "hello":
                var v5 uint32 = 2
                _ = v5
            case "\n\t":
                v6 := GenS1{F0: 255, F1: 1, F2: true}
                _ = v6
            default:
                switch 32768 {
                case 1:
                    v1 -= v1
                }
            }
            var v7 int32 = 1
            _ = v7
        }
        if false {
            v1 = (-(v1))
            for fi8 := 0; fi8 < 14; fi8++ {
                if (false && (0.0 < (0.0 * 0.1))) {
                    if (0 == 4294967295) {
                        v1 += (v1 & v1)
                        if false {
                            var v9 uint64 = 413
                            fmt.Println(fi8)
                            _ = v9
                        } else {
                            v3 = (v3 - 439)
                            var v10 int32 = 851
                            _ = v10
                        }
                    }
                    v11 := GenS1{F0: 9, F1: 824, F2: true}
                    for fi12 := 0; fi12 < 11; fi12++ {
                        fmt.Println(v3, fi12, v3)
                        _ = fi12
                    }
                    _ = v11
                }
                switch ("" + "foo bar") {
                case "":
                    switch "" {
                    case "a":
                        fmt.Println(v3)
                    case "foo bar":
                        var v13 int32 = 0
                        _ = v13
                    default:
                        var v14 uint32 = 4294967295
                        _ = v14
                    }
                case "\n\t":
                    var v15 uint16 = 32767
                    _ = v15
                }
                _ = fi8
            }
        }
    }
    v1 = (v1 - 212)
    if (v3 <= v3) {
        v1 = v1
    } else {
        for fi16 := 0; fi16 < 2; fi16++ {
            if false {
                v3 = v3
            }
            var v17 string = "\n\t"
            if ("" == v17) {
                if true {
                    var v18 int8 = 1
                    var v19 int32 = 560
                    if ((((-2.25 > (0.0 - ((-2.25 + 0.0) - 0.0))) && (1 >= 127)) && (true || (v3 > v3))) && true) {
                        v3 = v3
                        var v20 int = fi16
                        _ = v20
                    }
                    _ = v18
                    _ = v19
                }
                v3 = (v3 >> ((65535) & 63))
            } else {
                if true {
                    if len(v2) > 0 {
                        v2[0] = 2147483648
                    }
                    switch (-(fi16)) {
                    case -2147483648:
                        var v21 uint = (^(v3))
                        _ = v21
                    case 9223372036854775807:
                        fmt.Println(fi16, v17)
                    }
                }
                v1 += v1
            }
            _ = fi16
            _ = v17
        }
        v1 += -1
        var v22 uint8 = 46
        _ = v22
    }
    if false {
        if (false || true) {
            if len(v2) > 0 {
                v2[0] = 2
            }
            var v23 uint64 = 42
            _ = v23
        }
    } else {
        switch ("foo bar" + "a") {
        case "\n\t":
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint8 = 194
                v26 := GenS1{F0: v25, F1: fi24, F2: false}
                v27 := GenS1{F0: v25, F1: (fi24 / ((fi24) | 1)), F2: (!true)}
                _ = fi24
                _ = v25
                _ = v26
                _ = v27
            }
        case "a":
            var v28 int = 128
            _ = v28
        case "foo bar":
            if (1 != 109) {
                v3 += (v3 / (((v3 / ((v3) | 1))) | 1))
            }
        }
    }
    for fi29 := 0; fi29 < 10; fi29++ {
        v30 := GenS2{F0: 65535, F1: false}
        _ = fi29
        _ = v30
    }
    var v31 int32 = -22
    v31 = (v31 * v31)
    v32 := GenS1{F0: 1, F1: -245, F2: false}
    for fi33 := 0; fi33 < 16; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            v3 += (v3 + v3)
            _ = fi34
        }
        v35 := GenS2{F0: 510, F1: (62 > 2)}
        _ = fi33
        _ = v35
    }
    for fi36 := 0; fi36 < 5; fi36++ {
        v3 = (v3 & v3)
        var v37 uint64 = 2
        var v38 string = "\n\t"
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 uint16 = 42
    v40 := GenS3{F0: ("" + "foo bar"), F1: v1}
    v3 += 2
    var v41 string = ""
    v42 := GenS4{F0: -128, F1: 88, F2: 2147483647}
    switch 2147483647 {
    case -133:
        var v43 float64 = (3.14159 / 1.5)
        _ = v43
    case 0:
        v1 -= 954
    case -48:
        if false {
            if (!true) {
                if true {
                    if ((((false && ((!true) && (true && (54 < 1)))) || (128 < 155)) && (float32(0.1) < (((float32(((0.1 / 1.5) + (0.1 / 1e10))) / 1.5) * 0.1) - (1e10 * (1e10 * (3.14159 + (1.5 * float32((float64(float32(1.5)) * 0.1))))))))) || (v41 > (v41 + v41))) {
                        v31 += (-(v31))
                        for fi44 := 0; fi44 < 15; fi44++ {
                            v45 := GenS4{F0: 0, F1: int8(v1), F2: 2147483647}
                            _ = fi44
                            _ = v45
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        v46 := GenS5{F0: v41, F1: -32768}
                        _ = v46
                    case "hello":
                        v47 := GenS6{F0: 32767, F1: 2}
                        _ = v47
                    }
                } else {
                    var v48 uint32 = 65535
                    for fi49 := 0; fi49 < 10; fi49++ {
                        v2 = append(v2, (^(v48)))
                        _ = len(v2)
                        _ = fi49
                    }
                    _ = v48
                }
            }
            v1 = (v1 >> (((v3 >> (((v3 * (v3 * (v3 % ((v3) | 1))))) & 63))) & 63))
            v3 = 276
        }
    default:
        var v50 uint32 = 42
        _ = v50
    }
    var v51 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v51
}
