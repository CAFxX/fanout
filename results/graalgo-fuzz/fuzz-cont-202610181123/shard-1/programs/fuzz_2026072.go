package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int64 }

func main() {
    var v1 int = -32
    v2 := []uint32{29, 42, 191}
    v3 := []string{("\n\t" + (("" + ("" + ((("foo bar" + "") + "") + ("" + "hello")))) + "foo bar"))}
    var v4 uint64 = 220
    var v5 float32 = 0.1
    var v6 uint8 = 31
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (float64(v5) < 1e10) {
        var v7 float64 = 3.14159
        v8 := GenS1{F0: (v6 | (v6 ^ v6)), F1: 254, F2: 42}
        _ = len(v2)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        v5 = float32(float64((v5 + v5)))
        switch 209 {
        case -9223372036854775808:
            fmt.Println(v1, v4, v1)
        case 721:
            switch ((("\n\t" + "hello") + "hello") + "hello") {
            case "a":
                var v10 uint8 = v6
                _ = v10
            case "hello":
                v11 := GenS2{F0: 2, F1: 4294967295}
                _ = v11
            case "\n\t":
                v12 := GenS1{F0: v6, F1: (v6 >> ((42) & 63)), F2: 42}
                _ = v12
            }
        default:
            if true {
                var v13 int32 = -2147483648
                switch fi9 {
                case 4294967295:
                    v14 := GenS2{F0: 65535, F1: 2}
                    _ = v14
                case -334:
                    if (false && true) {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            v2 = append(v2, 65535)
                            var v16 uint = 0
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v4 -= 4294967295
                    }
                case 32768:
                    fmt.Println(v13, v13)
                }
                _ = v13
            }
        }
        if (!true) {
            switch "" {
            case "":
                var v17 uint16 = 160
                _ = v17
            case "\n\t":
                var v18 uint32 = 2147483647
                _ = v18
            case "foo bar":
                v5 = float32(0.0)
            default:
                switch -9223372036854775808 {
                case 2147483647:
                    switch -459 {
                    case 4294967295:
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            fmt.Println(fi9, v5, v6)
                        }
                    default:
                        v6 += (v6 + v6)
                    }
                }
            }
            if (42 <= 18446744073709551615) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v19 int32 = 1
                for fi20 := 0; fi20 < 3; fi20++ {
                    v6 -= v6
                    _ = fi20
                }
                _ = v19
            }
        } else {
            var v21 int64 = 42
            var v22 int = (v1 ^ (v1 ^ 533))
            _ = v21
            _ = v22
        }
        _ = fi9
    }
    var v23 int16 = -32768
    v6 += v6
    v4 = uint64(v6)
    for fi24 := 0; fi24 < 10; fi24++ {
        if false {
            for fi25 := 0; fi25 < 14; fi25++ {
                if len(v2) > 0 {
                    v2[0] = 2147483647
                }
                _ = fi25
            }
        } else {
            var v26 float32 = v5
            for fi27 := 0; fi27 < 5; fi27++ {
                v23 -= 0
                v28 := GenS1{F0: 128, F1: v6, F2: 0}
                v23 = v23
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        fmt.Println(v1, v4)
        if false {
            var v29 uint64 = v4
            v2 = append(v2, 2)
            _ = v29
        } else {
            v5 -= (v5 - 1e10)
        }
        _ = fi24
    }
    fmt.Println(v4, v23)
    v30 := GenS3{F0: 255, F1: -9223372036854775808}
    for fi31 := 0; fi31 < 8; fi31++ {
        var v32 uint64 = (v4 + v4)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        v6 = (v6 % ((uint8(v4)) | 1))
        if true {
            fmt.Println(v5, v1, v5)
        } else {
            var v34 int64 = -9223372036854775808
            _ = v34
        }
        fmt.Println(v1, v4)
        _ = fi33
    }
    v6 -= v6
    switch v1 {
    case -778:
        fmt.Println(v6, v1)
    }
    v6 -= (^(v6))
    var v35 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
