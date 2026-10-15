package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int64; F1 uint; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 uint }

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 160
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]int)
    var v5 uint = 2
    v1 = (v1 << (((v5 >> ((v5) & 63))) & 63))
    if (-830 < 1) {
        v2 = v2
    }
    var v6 int = -789
    switch (v6 | 128) {
    case -9223372036854775808:
        for fi7 := 0; fi7 < 8; fi7++ {
            for fi8 := 0; fi8 < 14; fi8++ {
                if ((float64((0.0 + 1.5)) - -2.25) != 1e10) {
                    v9 := GenS1{F0: (^(v6))}
                    _ = v9
                }
                _ = fi8
            }
            if false {
                for fi10 := 0; fi10 < 14; fi10++ {
                    v5 = v5
                    _ = fi10
                }
            }
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v5, fi7)
                _ = fi11
            }
            _ = fi7
        }
    case 32768:
        v12 := GenS2{F0: 42, F1: -36}
        _ = v12
    }
    switch "hello" {
    case "\n\t":
        v1 = (v1 - 320)
    case "foo bar":
        v4[(^(v2))] = -382
    }
    var v13 int = (v6 - int(v5))
    for fi14 := 0; fi14 < 7; fi14++ {
        switch (v13 >> (((v5 % (((^(v5))) | 1))) & 63)) {
        case -1:
            if true {
                switch (v6 * fi14) {
                case -2147483648:
                    switch (("" + "foo bar") + "hello") {
                    case "a":
                        v15 := GenS3{F0: 100, F1: (176 <= 127)}
                        _ = v15
                    }
                default:
                    if (!true) {
                        v16 := GenS4{F0: 2147483647, F1: v13}
                        switch "hello" {
                        case "a":
                            switch (fi14 >> ((v5) & 63)) {
                            case 32768:
                                v6 = -2147483648
                            }
                        case "foo bar":
                            var v17 uint = v5
                            _ = v17
                        }
                        v2 -= 2
                        _ = v16
                    }
                }
                if ((v1 % ((v1) | 1)) != v1) {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        var v19 float32 = 0.1
                        if (!(true || (!(!true)))) {
                            v20 := []int{(v13 << ((841) & 63))}
                            if false {
                                v21 := GenS3{F0: -104, F1: true}
                                var v22 string = ""
                                _ = len(v20)
                                _ = v21
                                _ = v22
                            }
                            fmt.Println(v2, fi18, v6)
                            _ = v20
                        } else {
                            var v23 string = ((("foo bar" + ("foo bar" + "\n\t")) + ((("foo bar" + "") + "a") + "\n\t")) + "hello")
                            _ = v23
                        }
                        _ = fi18
                        _ = v19
                    }
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if false {
                            fmt.Println(v2, v6, v6)
                        } else {
                            switch (fi14 & v13) {
                            case 161:
                                var v25 string = (((("a" + "hello") + "foo bar") + "") + ("\n\t" + (("hello" + "hello") + (((("foo bar" + (((("foo bar" + "") + ("" + "hello")) + "\n\t") + "a")) + ("a" + "\n\t")) + "") + "hello"))))
                                _ = v25
                            case 42:
                                var v26 float32 = 3.14159
                                _ = v26
                            case -826:
                                var v27 float32 = (3.14159 - ((1e10 / (float32((1e10 / -2.25)) - ((float32(1.5) - -2.25) * ((1.5 + 0.0) + (((0.0 + 3.14159) / float32((3.14159 + (float64(1e10) + 0.1)))) / (float32((0.0 - (((float64((0.0 / 1.0)) + (1e10 + 1.5)) / (((-2.25 / 1.0) * (1.5 * (-2.25 * 1e10))) - 0.1)) * float64(0.1)))) + 1e10)))))) - 1e10))
                                v28 := []int{(^(fi24)), (fi14 << ((v5) & 63)), (v6 % (((fi14 | 803)) | 1)), v13}
                                _ = v27
                                _ = v28
                            default:
                                var v29 bool = false
                                _ = v29
                            }
                        }
                        _ = fi24
                    }
                    v6 -= (-(fi14))
                } else {
                    v30 := GenS1{F0: v13}
                    v1 -= (v1 ^ 0)
                    switch (fi14 ^ v13) {
                    case 32767:
                        var v31 int32 = -1
                        _ = v31
                    case -967:
                        fmt.Println(v13)
                    case 2147483647:
                        fmt.Println(v1)
                    }
                    _ = v30
                }
            } else {
                var v32 uint8 = 42
                switch 42 {
                case 1:
                    switch v13 {
                    case 630:
                        v33 := GenS4{F0: 1, F1: (v6 + (v13 ^ (v13 >> ((v5) & 63))))}
                        _ = v33
                    case 127:
                        var v34 uint = 18446744073709551615
                        _ = v34
                    case 38:
                        var v35 string = "a"
                        _ = v35
                    }
                }
                var v36 float64 = 0.0
                _ = v32
                _ = v36
            }
        }
        var v37 bool = true
        _ = fi14
        _ = v37
    }
    for fi38 := 0; fi38 < 4; fi38++ {
        for fi39 := 0; fi39 < 15; fi39++ {
            v40 := []int{fi39, (v6 / (((fi39 >> ((v5) & 63))) | 1))}
            _ = len(v3)
            if (0 <= 127) {
                v41 := GenS5{F0: 710, F1: v5, F2: (v1 % (((v1 >> ((v5) & 63))) | 1))}
                fmt.Println(fi39, v2, v2)
                v42 := GenS2{F0: -71, F1: -16}
                _ = v41
                _ = v42
            }
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v43 := GenS6{F0: 42}
    if ((true || ((true && (127 == 42)) && true)) || (3.14159 != (float32(float64(-2.25)) + 1e10))) {
        if false {
            v2 = 2147483647
            var v44 float32 = (1e10 - (float32(3.14159) + (0.1 / 0.1)))
            v5 = 1
            _ = v44
        }
        v45 := GenS3{F0: -128, F1: true}
        _ = v45
    }
    if (!(100 == -128)) {
        var v46 int8 = -54
        _ = v46
    } else {
        v6 += v13
        v47 := GenS7{F0: (v1 / ((v1) | 1)), F1: 127}
        _ = v47
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        fmt.Println(v13, v13)
        v49 := GenS8{F0: v5}
        _ = fi48
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
