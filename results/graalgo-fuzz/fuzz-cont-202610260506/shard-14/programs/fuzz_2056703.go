package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }

func main() {
    var v1 uint64 = 2
    v2 := []string{"", "a", "hello", "hello"}
    var v3 bool = (85 <= 104)
    var v4 bool = (288 <= 42)
    v4 = (!v3)
    v5 := GenS1{F0: 263, F1: 178, F2: 65535}
    var v6 int = 2147483647
    if true {
        v3 = true
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS2{F0: (^(fi7)), F1: 1, F2: 2}
            _ = fi8
            _ = v9
        }
        fmt.Println(v3)
        _ = fi7
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        fmt.Println(v4, v1, v6)
        for fi11 := 0; fi11 < 11; fi11++ {
            v6 += 128
            switch "\n\t" {
            case "foo bar":
                var v12 uint = 18446744073709551615
                _ = v12
            case "\n\t":
                if ((-1 == 32767) && true) {
                    if (42 >= -912) {
                        var v13 uint = 112
                        _ = v13
                    } else {
                        switch ("foo bar" + "hello") {
                        case "a":
                            v3 = (v4 && ((65535 > 255) && v4))
                        case "hello":
                            v6 -= (fi11 * 128)
                        case "":
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v4 = (1 == 0)
                                var v14 uint = 4294967295
                                _ = v14
                            }
                        default:
                            if (true || (!v3)) {
                                v15 := GenS1{F0: (-(fi10)), F1: 1, F2: 42}
                                var v16 float64 = 0.0
                                _ = v15
                                _ = v16
                            }
                        }
                    }
                    var v17 uint64 = 0
                    v2 = append(v2, "a")
                    _ = v17
                } else {
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        fmt.Println(v6, v3, v1)
                    case 4294967295:
                        v18 := GenS1{F0: -2147483648, F1: 65535, F2: 2}
                        _ = v18
                    case 127:
                        v2 = append(v2, "")
                    }
                    var v19 uint32 = 42
                    var v20 uint = 42
                    _ = v19
                    _ = v20
                }
            case "hello":
                v1 = 1
            }
            if (0 <= 31) {
                fmt.Println(fi11, v6, v1)
                if v3 {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                }
                if false {
                    switch "hello" {
                    case "foo bar":
                        var v21 int16 = 1
                        _ = v21
                    case "":
                        if v3 {
                            fmt.Println(fi10)
                            if len(v2) > 0 {
                                v2[0] = (("hello" + ("hello" + ("a" + (("hello" + ("a" + "foo bar")) + ((((("" + "\n\t") + ("a" + "\n\t")) + "a") + "hello") + ((("hello" + ("a" + "a")) + ("hello" + "")) + ("\n\t" + ("" + (("foo bar" + ("hello" + "a")) + "a"))))))))) + "a")
                            }
                        } else {
                            v22 := GenS2{F0: fi10, F1: 940, F2: 42}
                            switch "" {
                            case "\n\t":
                                switch "a" {
                                case "":
                                    if (("a" < ("a" + "\n\t")) || (((2147483647 >= 42) && ((fi10 <= 32767) && false)) && v3)) {
                                        v3 = ((!v4) || false)
                                        var v23 string = (((("\n\t" + "foo bar") + "foo bar") + ("" + "")) + (("a" + "hello") + "foo bar"))
                                        _ = v23
                                    }
                                case "hello":
                                    v3 = v3
                                }
                            case "hello":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                v3 = (!(0 > -2147483648))
                            }
                            v4 = true
                            _ = v22
                        }
                    case "hello":
                        if v3 {
                            v3 = v3
                            fmt.Println(fi10, fi10)
                            fmt.Println(fi10, v1, v1)
                        } else {
                            var v24 uint8 = 167
                            v24 = v24
                            if v3 {
                                switch ("a" + "\n\t") {
                                case "":
                                    v25 := GenS1{F0: (^(v6)), F1: 602, F2: 32767}
                                    _ = v25
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    var v26 int16 = 42
                                    _ = v26
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v6 = (fi11 >> ((65535) & 63))
                                fmt.Println(v1)
                                v3 = (v3 && (887 > 1))
                            }
                            _ = v24
                        }
                    default:
                        if v3 {
                            v2 = append(v2, ((("\n\t" + (("" + "") + "")) + ("hello" + (("a" + "\n\t") + ("foo bar" + "")))) + ""))
                            if true {
                                if false {
                                    v4 = v3
                                    var v27 int = -9223372036854775808
                                    _ = v27
                                } else {
                                    var v28 uint64 = (v1 % ((v1) | 1))
                                    v1 -= 970
                                    _ = v28
                                }
                            }
                            var v29 int32 = -485
                            _ = v29
                        }
                    }
                    var v30 int32 = 2147483647
                    v6 -= fi10
                    _ = v30
                } else {
                    switch fi11 {
                    case 32768:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -9223372036854775808:
                        var v31 int16 = 0
                        _ = v31
                    case 4294967295:
                        v3 = false
                    default:
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    }
                    var v32 int8 = -1
                    if false {
                        if (((false || v3) || (!(!(v4 || v4)))) && v4) {
                            var v33 int64 = -396
                            if (0 >= -2147483648) {
                                v34 := GenS3{F0: fi11, F1: 4294967295, F2: 65535}
                                _ = v34
                            }
                            _ = v33
                        } else {
                            if v3 {
                                fmt.Println(v3, v3, v3)
                                var v35 uint16 = 1
                                _ = v35
                            }
                            v2 = append(v2, "a")
                            var v36 int8 = 119
                            _ = v36
                        }
                        v4 = (185 > v6)
                    } else {
                        if (1e10 == 0.1) {
                            fmt.Println(v1, fi11)
                            var v37 uint64 = 729
                            v38 := GenS3{F0: v6, F1: 661, F2: 876}
                            v32 -= (v32 >> ((2) & 63))
                            _ = v37
                            _ = v38
                        } else {
                            v1 += 0
                            switch "foo bar" {
                            case "hello":
                                v32 -= (v32 & v32)
                            case "\n\t":
                                var v39 int8 = v32
                                _ = v39
                            case "a":
                                if v4 {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                } else {
                                    var v40 int16 = -444
                                    _ = v40
                                }
                                var v41 int64 = 1
                                v6 = (fi10 * -1)
                                _ = v41
                            default:
                                v4 = v3
                            }
                            v32 = v32
                        }
                    }
                    v42 := GenS4{F0: v1, F1: v32}
                    var v43 uint = 255
                    if (127 == 118) {
                        var v44 uint = v43
                        if (42 >= 1) {
                            v6 = (fi10 ^ fi10)
                        } else {
                            var v45 uint8 = 0
                            v45 -= (v45 + v45)
                            switch (v6 ^ (fi11 * (fi10 + -1))) {
                            case 32767:
                                var v46 uint8 = 34
                                _ = v46
                            case -815:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v47 int = -2147483648
                                _ = v47
                            }
                            v6 -= v6
                            var v48 int = (v6 ^ v6)
                            _ = v45
                            _ = v48
                        }
                        fmt.Println(v43, fi10)
                        _ = v44
                    }
                    v49 := GenS1{F0: (^(fi10)), F1: 2, F2: 533}
                    v6 -= fi10
                    if ((v4 || (v1 < v1)) || v4) {
                        fmt.Println(v3, fi11)
                        switch ("foo bar" + "") {
                        case "\n\t":
                            if (v3 || true) {
                                v50 := GenS1{F0: fi10, F1: 42, F2: 255}
                                var v51 float64 = 0.1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v50
                                _ = v51
                            }
                        case "a":
                            v52 := GenS4{F0: 1, F1: v32}
                            _ = v52
                        case "hello":
                            var v53 uint64 = v1
                            _ = v53
                        }
                        _ = len(v2)
                    } else {
                        v6 = (fi10 * 570)
                    }
                    _ = v32
                    _ = v42
                    _ = v43
                    _ = v49
                }
            } else {
                v54 := GenS1{F0: (fi11 - v6), F1: 32767, F2: 1}
                v3 = v3
                var v55 float32 = 0.1
                _ = v54
                _ = v55
            }
            _ = fi11
        }
        _ = fi10
    }
    switch ("a" + ("hello" + "\n\t")) {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v56 int64 = -1
    v3 = (v6 < (v6 << ((0) & 63)))
    switch (v6 ^ (v6 / (((v6 << ((42) & 63))) | 1))) {
    case 604:
        switch "" {
        case "hello":
            var v57 int8 = 6
            _ = v57
        default:
            v6 = v6
        }
    case 32767:
        v1 = v1
    default:
        if v3 {
            if v3 {
                v56 = v56
                var v58 float64 = (float64(0.1) - 0.1)
                var v59 uint = 347
                _ = v58
                _ = v59
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch "" {
            case "a":
                v4 = (true || (!(818 > 1000)))
            case "\n\t":
                v6 += v6
            default:
                v2 = append(v2, ("" + ""))
            }
        }
    }
    v4 = (v4 || (255 >= 1))
    v60 := GenS5{F0: 1, F1: -1, F2: v4}
    var v61 int32 = -10
    for fi62 := 0; fi62 < 2; fi62++ {
        v63 := GenS4{F0: v1, F1: -128}
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v60
    _ = v61
}
