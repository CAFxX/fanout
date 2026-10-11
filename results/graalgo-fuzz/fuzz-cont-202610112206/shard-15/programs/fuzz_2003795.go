package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint8; F2 int32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 string; F2 uint32 }
type GenS8 struct { F0 int8; F1 int64; F2 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("a" + "a"), "a", ""}
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    var v5 uint64 = 1
    var v6 uint32 = 500
    v5 = 4294967295
    switch "\n\t" {
    case "hello":
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = (v5 << ((0) & 63))
            _ = fi7
            _ = v8
        }
    default:
        v6 -= v6
    }
    switch "" {
    case "a":
        switch -222 {
        case 308:
            if false {
                v6 -= (v6 >> ((42) & 63))
            } else {
                var v9 bool = (true || (-282 < -135))
                if (-717 != -666) {
                    if (v9 || ((v5 * v5) > v5)) {
                        v6 = v6
                    }
                } else {
                    fmt.Println(v5, v6)
                    var v10 string = ""
                    _ = v10
                }
                switch ("\n\t" + "") {
                case "hello":
                    if (!(!(42 > 42))) {
                        fmt.Println(v9)
                        var v11 int16 = 32767
                        _ = v11
                    }
                case "foo bar":
                    for fi12 := 0; fi12 < 7; fi12++ {
                        v9 = true
                        _ = fi12
                    }
                case "\n\t":
                    v9 = true
                default:
                    fmt.Println(v6)
                }
                _ = v9
            }
        case -2147483648:
            _ = len(v1)
        default:
            for fi13 := 0; fi13 < 8; fi13++ {
                if (!true) {
                    delete(v1, fi13)
                } else {
                    var v14 uint32 = (^(v6))
                    v15 := GenS1{F0: (fi13 - fi13), F1: fi13}
                    var v16 uint64 = (v5 - (v5 % (((v5 ^ v5)) | 1)))
                    _ = v14
                    _ = v15
                    _ = v16
                }
                v4[260] = uint32(v5)
                for fi17 := 0; fi17 < 8; fi17++ {
                    if (127 > 71) {
                        delete(v3, 444)
                    } else {
                        v5 += (v5 + 453)
                        v18 := GenS1{F0: (fi13 * (fi13 >> ((65535) & 63))), F1: fi17}
                        v19 := GenS2{F0: 42, F1: ("hello" + "hello")}
                        v5 = 42
                        switch "foo bar" {
                        case "foo bar":
                            if false {
                                if (!false) {
                                    v5 -= v5
                                }
                                v5 = (v5 << ((737) & 63))
                                v6 = (v6 >> ((18446744073709551615) & 63))
                                switch ("foo bar" + ("" + "foo bar")) {
                                case "hello":
                                    _ = len(v2)
                                case "foo bar":
                                    v6 = 4294967295
                                }
                                fmt.Println(fi13, v6)
                            }
                        case "hello":
                            switch (fi17 & (fi13 / (((fi17 << ((42) & 63))) | 1))) {
                            case -9223372036854775808:
                                delete(v3, int(v5))
                                delete(v4, 32767)
                                v6 -= 2147483648
                                var v20 float32 = 0.0
                                v2 = append(v2, "hello")
                                if true {
                                    if false {
                                        fmt.Println(fi13, v6)
                                        v5 += v5
                                    } else {
                                        fmt.Println(fi17, v20)
                                        v20 = v20
                                    }
                                    v6 += (v6 % ((v6) | 1))
                                    if (!true) {
                                        v20 -= 3.14159
                                        var v21 uint = 42
                                        v6 += v6
                                        _ = v21
                                    } else {
                                        switch ("foo bar" + ("" + "a")) {
                                        case "foo bar":
                                            var v22 uint64 = 42
                                            _ = v22
                                        }
                                    }
                                    v23 := GenS3{F0: v6, F1: 155}
                                    var v24 int = fi17
                                    v25 := GenS4{F0: 42, F1: 100, F2: -2147483648}
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    v20 = -2.25
                                    _ = len(v2)
                                    fmt.Println(v6)
                                }
                                _ = v20
                            case -2147483648:
                                v6 -= v6
                            }
                        case "\n\t":
                            _ = len(v2)
                        }
                        v5 -= 1
                        _ = v18
                        _ = v19
                    }
                    _ = fi17
                }
                _ = fi13
            }
        }
    default:
        v26 := GenS3{F0: (^(v6)), F1: 2}
        _ = v26
    }
    _ = len(v2)
    var v27 uint16 = 748
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS1{F0: (-(fi28)), F1: (fi28 % (((^(fi28))) | 1))}
        _ = fi28
        _ = v29
    }
    var v30 float64 = 1.5
    for fi31 := 0; fi31 < 12; fi31++ {
        v4[(-(fi31))] = 2147483647
        v30 += (3.14159 * 0.1)
        _ = fi31
    }
    var v32 int = 127
    v5 = 1
    for fi33 := 0; fi33 < 15; fi33++ {
        v34 := v4[(fi33 % ((int(v27)) | 1))]
        var v35 int16 = 553
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint32 = (v6 >> ((702) & 63))
    switch (("foo bar" + "a") + "") {
    case "hello":
        v37 := v3[(^(v32))]
        _ = v37
    case "a":
        var v38 uint = 4294967295
        _ = v38
    default:
        if (0.0 == float32(((((float64((1e10 / -2.25)) * v30) + 1e10) - float64(1.5)) - v30))) {
            if (v30 >= (((v30 - 0.1) / v30) + (v30 / -2.25))) {
                for fi39 := 0; fi39 < 14; fi39++ {
                    switch "" {
                    case "\n\t":
                        v40 := v3[-642]
                        _ = v40
                    }
                    v6 = 799
                    var v41 uint32 = v6
                    _ = fi39
                    _ = v41
                }
                var v42 int16 = 0
                _ = v42
            } else {
                var v43 float32 = 1.5
                v32 -= (-(v32))
                var v44 uint8 = 2
                _ = v43
                _ = v44
            }
        } else {
            var v45 int64 = 1
            _ = v45
        }
    }
    v46 := GenS4{F0: 938, F1: uint8(v5), F2: 42}
    if true {
        fmt.Println(v27)
        if (true && false) {
            for fi47 := 0; fi47 < 11; fi47++ {
                if true {
                    v3[fi47] = "foo bar"
                } else {
                    var v48 bool = true
                    var v49 int = (v32 >> ((4294967295) & 63))
                    _ = v48
                    _ = v49
                }
                v30 += v30
                var v50 bool = (false || ((v27 & 143) < v27))
                _ = fi47
                _ = v50
            }
            v32 = v32
            switch "" {
            case "\n\t":
                v51 := GenS5{F0: int(v5)}
                _ = v51
            case "":
                for fi52 := 0; fi52 < 20; fi52++ {
                    if (!false) {
                        for fi53 := 0; fi53 < 5; fi53++ {
                            var v54 int32 = 2147483647
                            fmt.Println(fi52, v32, v36)
                            _ = fi53
                            _ = v54
                        }
                        v5 = v5
                        if true {
                            if true {
                                v4[49] = v6
                            } else {
                                switch v32 {
                                case 128:
                                    if true {
                                        v6 = (v36 * 4294967295)
                                    }
                                }
                                v55 := GenS6{F0: v6}
                                var v56 int32 = 2147483647
                                _ = v55
                                _ = v56
                            }
                            if true {
                                var v57 uint = 18446744073709551615
                                var v58 uint16 = v27
                                _ = v57
                                _ = v58
                            }
                        } else {
                            var v59 uint64 = (^(v5))
                            if true {
                                v36 = v6
                                fmt.Println(v27, v36, v36)
                            }
                            if ((false || (!false)) || true) {
                                delete(v3, int(v27))
                            }
                            _ = v59
                        }
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v1)
                    }
                    _ = fi52
                }
            }
        } else {
            _ = len(v2)
            fmt.Println(v5, v30)
            v32 += (^(v32))
        }
    }
    switch "a" {
    case "\n\t":
        for fi60 := 0; fi60 < 5; fi60++ {
            fmt.Println(v32)
            _ = fi60
        }
    case "hello":
        v61 := GenS7{F0: 2, F1: "a", F2: (v6 % ((376) | 1))}
        _ = v61
    case "foo bar":
        var v62 bool = ((v27 / (((^(v27))) | 1)) <= v27)
        _ = v62
    default:
        v4[v32] = v6
    }
    if true {
        if false {
            switch "" {
            case "hello":
                v63 := GenS8{F0: 127, F1: -234, F2: 1}
                _ = v63
            case "\n\t":
                switch ("hello" + "a") {
                case "hello":
                    if (658 != 255) {
                        if false {
                            v5 = v5
                            for fi64 := 0; fi64 < 8; fi64++ {
                                v65 := GenS6{F0: v6}
                                switch (fi64 << ((uint(v27)) & 63)) {
                                case 4294967295:
                                    switch fi64 {
                                    case 162:
                                        var v66 int16 = -100
                                        _ = v66
                                    case 1:
                                        for fi67 := 0; fi67 < 16; fi67++ {
                                            v68 := GenS5{F0: fi64}
                                            if (!false) {
                                                v2 = append(v2, "\n\t")
                                            } else {
                                                v69 := GenS4{F0: 0, F1: 255, F2: 849}
                                                _ = v69
                                            }
                                            _ = fi67
                                            _ = v68
                                        }
                                    case -640:
                                        v32 = fi64
                                    }
                                }
                                v36 = (v6 << ((606) & 63))
                                _ = fi64
                                _ = v65
                            }
                        } else {
                            v30 = float64(0.0)
                            var v70 int16 = 0
                            _ = v70
                        }
                        v2 = append(v2, "\n\t")
                    }
                case "\n\t":
                    switch "hello" {
                    case "foo bar":
                        var v71 uint32 = (v6 ^ v36)
                        _ = v71
                    case "hello":
                        var v72 string = ("" + ("\n\t" + ("\n\t" + "\n\t")))
                        _ = v72
                    case "\n\t":
                        v73 := GenS6{F0: 0}
                        _ = v73
                    }
                case "foo bar":
                    v36 -= v6
                default:
                    if true {
                        v6 += v36
                        switch ("foo bar" + "a") {
                        case "a":
                            v74 := GenS6{F0: v36}
                            _ = v74
                        }
                    } else {
                        var v75 int = (v32 / (((-(v32))) | 1))
                        _ = v75
                    }
                }
            case "a":
                v6 = (v6 << ((128) & 63))
            }
            var v76 uint16 = v27
            _ = v76
        } else {
            for fi77 := 0; fi77 < 18; fi77++ {
                var v78 int8 = 127
                _ = len(v2)
                switch "" {
                case "":
                    v2 = append(v2, ("hello" + ("foo bar" + "a")))
                case "\n\t":
                    v36 = (v36 * 2147483648)
                case "a":
                    v79 := GenS6{F0: v36}
                    _ = v79
                }
                _ = fi77
                _ = v78
            }
            var v80 string = ("" + "a")
            _ = v80
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v32
    _ = v36
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
