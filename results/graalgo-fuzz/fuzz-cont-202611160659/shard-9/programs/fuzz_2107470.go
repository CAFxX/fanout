package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 int16; F1 uint8; F2 string }
type GenS4 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 int32 = 2147483647
    var v2 int = 1
    v3 := []string{"foo bar", "a", "hello"}
    var v4 uint64 = 4294967295
    var v5 int32 = v1
    v6 := []string{("\n\t" + ("\n\t" + "a"))}
    var v7 float32 = float32(1.5)
    for fi8 := 0; fi8 < 20; fi8++ {
        if len(v6) > 0 {
            v6[0] = "\n\t"
        }
        if (4294967295 < 2) {
            if len(v3) > 0 {
                v3[0] = ((("hello" + "\n\t") + "foo bar") + "a")
            }
        }
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    var v10 int32 = v5
    switch "hello" {
    case "a":
        switch ((("\n\t" + "hello") + "hello") + "") {
        case "":
            v10 += v5
        }
    case "hello":
        var v11 float32 = 0.1
        _ = v11
    case "\n\t":
        v5 = v5
    }
    var v12 int = (v2 >> ((18446744073709551615) & 63))
    for fi13 := 0; fi13 < 10; fi13++ {
        if (("foo bar" + ("\n\t" + "foo bar")) == "\n\t") {
            switch "" {
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                switch (("hello" + "a") + "foo bar") {
                case "foo bar":
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 uint16 = 1
                        _ = fi14
                        _ = v15
                    }
                case "\n\t":
                    if true {
                        var v16 float32 = ((0.1 * v7) + v7)
                        var v17 float32 = (v16 - 3.14159)
                        fmt.Println(v5, v12)
                        _ = v16
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            var v19 uint = 794
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            v7 -= v7
                            v12 = v2
                            _ = fi18
                            _ = v19
                        }
                        fmt.Println(v12, v5, v2)
                    }
                case "a":
                    v7 += (0.1 / (((1e10 / ((-2.25 / 3.14159) * ((((((v7 / (v7 + (v7 + 1e10))) * v7) / (((v7 * (1e10 + (-2.25 + ((v7 + (float32(0.1) * v7)) / 1e10)))) * ((float32((-2.25 - 0.0)) + float32(float64(((v7 / v7) + float32(((1.5 / -2.25) / (float64(v7) + 0.0))))))) - ((0.1 / ((1e10 - ((((((float32(-2.25) * v7) + float32(1e10)) + v7) * v7) * float32(3.14159)) / 1.0)) + v7)) * (v7 + v7)))) / ((0.0 / (0.1 * -2.25)) / (v7 * float32((((-2.25 / (1.5 / 1.5)) * (3.14159 * 1.5)) + -2.25)))))) + 3.14159) / 3.14159) / v7))) * (float32((1.5 / float64(v7))) / v7)) + (1.5 - 1e10)))
                }
            }
            v1 = v5
            switch (v12 % ((42) | 1)) {
            case -1:
                var v20 int32 = v10
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 3; fi21++ {
            v22 := GenS1{F0: (v10 + v5), F1: 0}
            if true {
                v23 := GenS1{F0: (v10 / ((v5) | 1)), F1: 42}
                v24 := GenS2{F0: (v2 | fi21), F1: 1, F2: 42}
                var v25 int64 = 0
                v7 = 1.5
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        _ = fi13
    }
    v10 -= v10
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            switch v2 {
            case 2147483647:
                switch v2 {
                case -2147483648:
                    switch (fi26 & (v2 + (v2 ^ v12))) {
                    case 163:
                        if (!false) {
                            var v27 float32 = v7
                            switch "\n\t" {
                            case "foo bar":
                                var v28 bool = (42 == 2)
                                _ = v28
                            case "":
                                switch v12 {
                                case 0:
                                    for fi29 := 0; fi29 < 20; fi29++ {
                                        if false {
                                            var v30 uint8 = 255
                                            _ = v30
                                        } else {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v31 int8 = 100
                                            _ = v31
                                        }
                                        _ = fi29
                                    }
                                case 128:
                                    _ = len(v3)
                                case -1:
                                    v6 = append(v6, ("\n\t" + "\n\t"))
                                default:
                                    v2 = v2
                                }
                            case "\n\t":
                                var v32 int8 = 100
                                _ = v32
                            }
                            _ = v27
                        }
                    case 0:
                        v3 = append(v3, "\n\t")
                    case -2147483648:
                        var v33 int16 = 883
                        _ = v33
                    }
                case -838:
                    v34 := GenS2{F0: 1, F1: 42, F2: 65}
                    _ = v34
                case 790:
                    v35 := GenS2{F0: (v2 ^ (v12 & v12)), F1: -674, F2: 15}
                    _ = v35
                default:
                    var v36 int16 = 42
                    _ = v36
                }
            case 32768:
                if false {
                    var v37 uint8 = 127
                    _ = v37
                } else {
                    v38 := GenS2{F0: fi26, F1: -9223372036854775808, F2: 42}
                    if (v7 == ((v7 + (0.0 / v7)) / 1.0)) {
                        if false {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                v4 += 0
                                switch "" {
                                case "":
                                    v4 += 42
                                case "foo bar":
                                    var v40 uint64 = (v4 << ((255) & 63))
                                    v4 = (v4 * v4)
                                    if true {
                                        v41 := GenS3{F0: 1, F1: 255, F2: ("\n\t" + "")}
                                        _ = v41
                                    } else {
                                        switch "hello" {
                                        case "\n\t":
                                            var v42 int8 = int8(v2)
                                            _ = v42
                                        case "":
                                            fmt.Println(v5, v2)
                                        }
                                        v43 := GenS1{F0: (^(v5)), F1: 0}
                                        v5 = 42
                                        _ = v43
                                    }
                                    _ = v40
                                case "hello":
                                    var v44 int8 = -1
                                    _ = v44
                                default:
                                    v45 := GenS4{F0: 1, F1: 42}
                                    _ = v45
                                }
                                if (2 != 820) {
                                    var v46 float32 = v7
                                    var v47 int16 = -581
                                    var v48 uint32 = 183
                                    _ = v46
                                    _ = v47
                                    _ = v48
                                }
                                var v49 int64 = 9223372036854775807
                                var v50 uint8 = 1
                                _ = fi39
                                _ = v49
                                _ = v50
                            }
                            _ = len(v3)
                            for fi51 := 0; fi51 < 1; fi51++ {
                                switch (v2 - v2) {
                                case 715:
                                    var v52 uint64 = (^(v4))
                                    _ = v52
                                default:
                                    var v53 string = "hello"
                                    _ = v53
                                }
                                _ = fi51
                            }
                        }
                    }
                    _ = v38
                }
            }
            var v54 uint64 = 1
            _ = v54
        }
        _ = fi26
    }
    var v55 string = (("foo bar" + "\n\t") + "\n\t")
    v4 += (v4 << ((1) & 63))
    fmt.Println(v10)
    switch ("a" + "") {
    case "":
        switch (v55 + "") {
        case "\n\t":
            if (127 <= 2) {
                var v56 int32 = (v10 >> ((255) & 63))
                _ = v56
            } else {
                for fi57 := 0; fi57 < 13; fi57++ {
                    fmt.Println(v12, v4, v4)
                    _ = fi57
                }
                var v58 string = v55
                var v59 uint8 = 255
                _ = v58
                _ = v59
            }
        }
    case "foo bar":
        var v60 uint8 = 0
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
