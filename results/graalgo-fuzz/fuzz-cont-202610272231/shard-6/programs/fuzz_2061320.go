package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int16; F1 bool; F2 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]uint32)
    var v3 int32 = 1
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    var v6 int = 32768
    fmt.Println(v6, v3)
    v6 += v6
    v3 += v3
    v6 -= v6
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v3, fi7, v3)
        v8 := GenS1{F0: 2, F1: 1000}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 85, F1: 1}
    fmt.Println(v3, v3)
    switch "hello" {
    case "foo bar":
        v3 += (v3 * (v3 << ((595) & 63)))
    case "hello":
        for fi10 := 0; fi10 < 3; fi10++ {
            v3 = v3
            _ = fi10
        }
    case "a":
        v11 := GenS2{F0: 128, F1: (v6 >> ((255) & 63))}
        _ = v11
    default:
        if (0 == 18446744073709551615) {
            for fi12 := 0; fi12 < 19; fi12++ {
                if (v3 <= (v3 / ((int32(v6)) | 1))) {
                    if (false && ((true || false) || ((((float64(((-2.25 - 0.1) / 0.1)) / 1e10) * 0.0) / -2.25) != 0.0))) {
                        v13 := []int{}
                        switch "hello" {
                        case "foo bar":
                            v14 := GenS2{F0: 127, F1: -9223372036854775808}
                            _ = v14
                        case "\n\t":
                            v6 -= 1
                        case "":
                            var v15 bool = true
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 7; fi16++ {
                            var v17 uint = 42
                            if true {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    switch 1 {
                                    case -100:
                                        var v18 uint16 = 997
                                        _ = v18
                                    default:
                                        if (-2.25 > 1.5) {
                                            v19 := GenS3{F0: -480, F1: true, F2: -1}
                                            _ = v19
                                        }
                                    }
                                case "hello":
                                    v20 := GenS3{F0: -1, F1: (!(3.14159 > 0.0)), F2: 1}
                                    _ = v20
                                }
                                switch "" {
                                case "":
                                    if true {
                                        v21 := GenS4{F0: 991, F1: 70}
                                        v22 := GenS3{F0: 364, F1: false, F2: -43}
                                        _ = v21
                                        _ = v22
                                    } else {
                                        var v23 int8 = -99
                                        v24 := GenS5{F0: 42}
                                        if (!true) {
                                            var v25 int32 = (v3 & v3)
                                            _ = v25
                                        }
                                        switch (^(v6)) {
                                        case 170:
                                            var v26 int8 = v23
                                            _ = v26
                                        case 2147483647:
                                            switch fi16 {
                                            case 40:
                                                if (255 < (v17 | v17)) {
                                                    var v27 int64 = 1
                                                    var v28 uint32 = 2
                                                    v28 -= (^(v28))
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    v2["\n\t"] = 255
                                                }
                                            case 9223372036854775807:
                                                var v29 int8 = (v23 >> ((505) & 63))
                                                _ = v29
                                            case 4294967295:
                                                v17 = 620
                                            default:
                                                var v30 uint8 = 71
                                                _ = v30
                                            }
                                        case 617:
                                            switch ("\n\t" + "foo bar") {
                                            case "hello":
                                                v6 = v6
                                            case "foo bar":
                                                var v31 uint8 = 3
                                                _ = v31
                                            case "\n\t":
                                                switch ("a" + "foo bar") {
                                                case "foo bar":
                                                    var v32 uint32 = 65535
                                                    _ = v32
                                                case "a":
                                                    if (false && true) {
                                                        v2[(("hello" + "a") + ("hello" + "hello"))] = uint32(v17)
                                                    }
                                                case "\n\t":
                                                    switch 1 {
                                                    case 9223372036854775807:
                                                        var v33 int8 = (v23 / ((v23) | 1))
                                                        _ = v33
                                                    default:
                                                        if (9223372036854775807 < -9223372036854775808) {
                                                            switch fi16 {
                                                            case 2147483647:
                                                                _ = len(v2)
                                                            }
                                                            var v34 string = ""
                                                            var v35 int = 0
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    }
                                                default:
                                                    var v36 uint32 = 1
                                                    _ = v36
                                                }
                                            }
                                        }
                                        if true {
                                            v6 += (^(fi12))
                                            v3 = v3
                                        }
                                        v37 := v4[(("hello" + "\n\t") + ("hello" + "foo bar"))]
                                        _ = v23
                                        _ = v24
                                        _ = v37
                                    }
                                case "\n\t":
                                    var v38 int8 = -6
                                    _ = v38
                                case "hello":
                                    v39 := GenS6{F0: 127, F1: 0}
                                    _ = v39
                                }
                            } else {
                                v6 = fi12
                            }
                            _ = fi16
                            _ = v17
                        }
                        _ = v13
                    }
                    v3 -= (v3 ^ v3)
                    for fi40 := 0; fi40 < 16; fi40++ {
                        switch 502 {
                        case 32767:
                            v41 := GenS3{F0: -896, F1: false, F2: -46}
                            _ = v41
                        case 92:
                            v42 := []int{v6, (fi40 + v6), v6}
                            _ = v42
                        }
                        _ = fi40
                    }
                } else {
                    var v43 int16 = 1
                    fmt.Println(v43, v43, v3)
                    _ = v43
                }
                v3 = -652
                _ = fi12
            }
        } else {
            var v44 uint8 = 255
            fmt.Println(v44)
            _ = v44
        }
    }
    if ((800 < 0) || true) {
        v3 = 0
        switch (("\n\t" + "hello") + "") {
        case "a":
            var v45 float64 = 1.5
            _ = v45
        default:
            v6 = (v6 | 1)
        }
        var v46 bool = false
        _ = v46
    } else {
        v6 -= (v6 & 2147483647)
    }
    v47 := []int{(v6 | (^(v6))), (v6 >> ((508) & 63)), -226, v6}
    for fi48 := 0; fi48 < 2; fi48++ {
        switch v6 {
        case 2147483647:
            v49 := GenS5{F0: 2}
            _ = v49
        case -2147483648:
            v3 = (v3 / ((-713) | 1))
        case 927:
            switch ("" + (("a" + "foo bar") + (("hello" + (("a" + "a") + "a")) + "hello"))) {
            case "\n\t":
                if true {
                    v50 := GenS5{F0: 411}
                    switch v6 {
                    case 133:
                        var v51 int = -756
                        _ = v51
                    case 32767:
                        v2["hello"] = 636
                    case 663:
                        v3 += v3
                    default:
                        fmt.Println(fi48)
                    }
                    _ = v50
                }
            case "hello":
                v3 -= v3
            case "":
                v6 = (fi48 - (fi48 & int(v3)))
            default:
                for fi52 := 0; fi52 < 13; fi52++ {
                    v53 := GenS7{F0: "hello"}
                    _ = fi52
                    _ = v53
                }
            }
        default:
            v54 := GenS8{F0: v3}
            _ = v54
        }
        for fi55 := 0; fi55 < 10; fi55++ {
            fmt.Println(v3)
            v4[("\n\t" + "foo bar")] = (v6 ^ 32767)
            _ = fi55
        }
        _ = fi48
    }
    if len(v47) >= 2 {
        v47 = v47[1:len(v47)]
    }
    v56 := GenS2{F0: 128, F1: v6}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v56
    _ = v6
    _ = v9
}
