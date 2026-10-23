package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 int32; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 int; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    var v1 uint8 = 145
    var v2 uint = 1
    var v3 float32 = float32(1.5)
    v1 -= v1
    if (true && ((true && true) || (-1 == 4294967295))) {
        v4 := GenS1{F0: 19, F1: 42}
        _ = v4
    } else {
        v5 := GenS2{F0: "hello"}
        v6 := []int{467, 127, 32768}
        switch 1 {
        case -638:
            v3 = -2.25
        default:
            for fi7 := 0; fi7 < 9; fi7++ {
                switch "foo bar" {
                case "a":
                    v3 += 0.1
                case "\n\t":
                    fmt.Println(v2)
                case "":
                    for fi8 := 0; fi8 < 2; fi8++ {
                        var v9 uint16 = 0
                        if false {
                            var v10 bool = (203 < 863)
                            _ = v10
                        } else {
                            if false {
                                v1 = 128
                            } else {
                                switch "" {
                                case "":
                                    switch (fi8 >> ((v2) & 63)) {
                                    case 684:
                                        _ = len(v6)
                                    case 4294967295:
                                        var v11 int64 = -9223372036854775808
                                        _ = v11
                                    case -247:
                                        v3 -= v3
                                    }
                                case "a":
                                    v12 := GenS3{F0: (v3 > v3), F1: -418}
                                    _ = v12
                                case "\n\t":
                                    var v13 uint32 = 0
                                    _ = v13
                                default:
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                            }
                            switch (-(fi7)) {
                            case 42:
                                var v14 uint16 = v9
                                _ = v14
                            case 4294967295:
                                v2 = 42
                            default:
                                switch fi8 {
                                case -742:
                                    var v15 float32 = ((v3 - v3) + 0.0)
                                    _ = v15
                                case -9223372036854775808:
                                    var v16 uint16 = (v9 % ((v9) | 1))
                                    _ = v16
                                case 128:
                                    if ("\n\t" != "a") {
                                        fmt.Println(v1, v1, v1)
                                        var v17 int64 = -865
                                        _ = v17
                                    }
                                    var v18 bool = (true || true)
                                    switch "\n\t" {
                                    case "hello":
                                        switch fi8 {
                                        case -731:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case 2147483647:
                                            v3 = (0.0 / 1.0)
                                        default:
                                            var v19 int16 = -32768
                                            _ = v19
                                        }
                                    default:
                                        switch 127 {
                                        case -2147483648:
                                            v3 = 0.0
                                        case 128:
                                            var v20 uint32 = 42
                                            _ = v20
                                        case -1:
                                            switch -568 {
                                            case -9223372036854775808:
                                                if v18 {
                                                    v21 := GenS1{F0: int8(v9), F1: 42}
                                                    _ = v21
                                                } else {
                                                    if v18 {
                                                        if v18 {
                                                            var v22 float64 = 1.5
                                                            _ = v22
                                                        }
                                                        if (1 <= 306) {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                fmt.Println(v1)
                                                            case "a":
                                                                v23 := GenS4{F0: int32(v2), F1: -29, F2: v9}
                                                                _ = v23
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            fmt.Println(v1, fi7, v1)
                                                        } else {
                                                            var v24 int8 = 127
                                                            _ = v24
                                                        }
                                                        fmt.Println(v18)
                                                    }
                                                }
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                var v25 int16 = 0
                                                _ = v25
                                            }
                                        default:
                                            if v18 {
                                                fmt.Println(v18, v9, v9)
                                            }
                                            v1 += (^(v1))
                                        }
                                    }
                                    _ = v18
                                default:
                                    var v26 string = ""
                                    v27 := GenS5{F0: (v1 | (v1 / (((^(v1))) | 1)))}
                                    _ = v26
                                    _ = v27
                                }
                            }
                        }
                        v3 = 0.0
                        _ = fi8
                        _ = v9
                    }
                default:
                    v3 += (float32(0.0) + v3)
                }
                if (uint(v1) <= 4294967295) {
                    var v28 int = (fi7 + (fi7 << (((v2 ^ (v2 * (v2 ^ v2)))) & 63)))
                    var v29 string = "foo bar"
                    var v30 bool = (v28 < (fi7 ^ v28))
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        if (541 > 18446744073709551615) {
                            var v32 uint16 = 1
                            v1 = 92
                            switch (fi31 << (((v2 & v2)) & 63)) {
                            case -721:
                                var v33 uint16 = (v32 & v32)
                                _ = v33
                            case 42:
                                fmt.Println(fi7)
                            case -1:
                                v32 = (v32 & 399)
                            }
                            _ = v32
                        }
                        var v34 uint32 = 0
                        _ = fi31
                        _ = v34
                    }
                    var v35 uint64 = 42
                    _ = v35
                }
                _ = fi7
            }
        }
        _ = v5
        _ = v6
    }
    var v36 int32 = 42
    v37 := []int{-2147483648, 32767}
    if true {
        if (((true && false) && false) && true) {
            if true {
                var v38 float32 = (((float32(3.14159) * (v3 / v3)) - v3) - v3)
                _ = v38
            } else {
                var v39 int16 = -1
                _ = v39
            }
            for fi40 := 0; fi40 < 17; fi40++ {
                v41 := GenS6{F0: 0, F1: -705, F2: -128}
                switch ("foo bar" + ("hello" + "hello")) {
                case "hello":
                    for fi42 := 0; fi42 < 19; fi42++ {
                        switch (fi42 | 751) {
                        case 53:
                            switch fi40 {
                            case 2147483647:
                                var v43 uint = v2
                                _ = v43
                            case -2147483648:
                                var v44 uint16 = 32767
                                _ = v44
                            case 32767:
                                var v45 int = (fi42 - 4294967295)
                                _ = v45
                            }
                        case -2147483648:
                            v37 = append(v37, (-(fi40)))
                        }
                        switch -325 {
                        case -482:
                            if true {
                                fmt.Println(v3)
                                v46 := GenS7{F0: (v36 - (v36 / ((v36) | 1)))}
                                var v47 int8 = 42
                                _ = v46
                                _ = v47
                            } else {
                                if true {
                                    var v48 int = fi40
                                    _ = v48
                                }
                            }
                        case -2147483648:
                            fmt.Println(v1, v36)
                        }
                        _ = fi42
                    }
                case "":
                    v49 := GenS2{F0: "hello"}
                    _ = v49
                case "\n\t":
                    switch (fi40 ^ -1) {
                    case 128:
                        v36 = (v36 >> ((748) & 63))
                    case 42:
                        fmt.Println(v3, v3)
                    default:
                        v50 := GenS6{F0: 2, F1: (fi40 >> ((65535) & 63)), F2: 127}
                        _ = v50
                    }
                default:
                    for fi51 := 0; fi51 < 3; fi51++ {
                        v1 += (v1 >> (((v2 + 255)) & 63))
                        v36 += (^(v36))
                        if len(v37) > 0 {
                            v37[0] = (fi40 / ((fi40) | 1))
                        }
                        _ = fi51
                    }
                }
                var v52 string = "\n\t"
                _ = fi40
                _ = v41
                _ = v52
            }
            if false {
                if (-291 > -9223372036854775808) {
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                }
            }
        }
        var v53 uint = v2
        v2 = (v2 / (((v2 * (v2 & (^(v53))))) | 1))
        _ = v53
    }
    if (v1 == 255) {
        switch 838 {
        case 4294967295:
            var v54 int32 = -161
            _ = v54
        }
        v3 += (((v3 - (v3 * ((float32((((3.14159 - 1.5) - 1.5) * 1.5)) / ((v3 + 1e10) / v3)) / v3))) / v3) * (0.1 * v3))
    }
    for fi55 := 0; fi55 < 9; fi55++ {
        fmt.Println(v1, v36)
        for fi56 := 0; fi56 < 7; fi56++ {
            var v57 uint64 = 2
            if true {
                var v58 uint8 = 42
                var v59 int8 = int8(v36)
                var v60 int16 = int16(v57)
                _ = v58
                _ = v59
                _ = v60
            } else {
                v61 := GenS8{F0: "hello"}
                _ = v61
            }
            _ = fi56
            _ = v57
        }
        if true {
            var v62 float32 = v3
            _ = v62
        }
        _ = fi55
    }
    var v63 int64 = -9223372036854775808
    v64 := GenS7{F0: v36}
    if true {
        var v65 int16 = -1
        _ = v65
    } else {
        v63 = (-(v63))
        if true {
            fmt.Println(v36)
        } else {
            fmt.Println(v63, v36, v3)
            switch 833 {
            case -2147483648:
                for fi66 := 0; fi66 < 8; fi66++ {
                    fmt.Println(v3)
                    _ = fi66
                }
            case 687:
                var v67 uint16 = 1
                _ = v67
            case 1:
                v1 = (v1 >> ((v2) & 63))
            default:
                v37 = append(v37, -13)
            }
        }
        v36 -= (v36 ^ (-(v36)))
    }
    if false {
        var v68 string = "hello"
        v69 := GenS2{F0: (v68 + v68)}
        v70 := GenS1{F0: -90, F1: 42}
        _ = v68
        _ = v69
        _ = v70
    }
    var v71 int64 = v63
    _ = len(v37)
    v72 := GenS9{F0: "\n\t", F1: 128}
    v73 := GenS10{F0: "hello", F1: true, F2: 1}
    if (true || (1 != 2147483648)) {
        v74 := GenS4{F0: v36, F1: -128, F2: 1}
        v3 = 1.5
        _ = v74
    } else {
        v2 = (v2 - (v2 + (v2 & 42)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v63:", v63)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v63
    _ = v64
    _ = v71
    _ = v72
    _ = v73
}
