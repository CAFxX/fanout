package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 bool }
type GenS4 struct { F0 int64; F1 int32; F2 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 string; F2 bool }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int64; F1 uint32 }
func genFn0(p0 bool) bool {
    if (-471 >= 800) {
        p0 = p0
        v5 := GenS1{F0: 42, F1: 2, F2: 42}
        _ = v5
    }
    p0 = p0
    var v6 int64 = -713
    if (true || p0) {
        var v7 uint32 = 206
        var v8 bool = ((v6 >> ((65535) & 63)) > v6)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        switch (fi9 * (fi9 / (((fi9 << ((696) & 63))) | 1))) {
        case 128:
            var v10 uint = 334
            _ = v10
        case 127:
            var v11 bool = (0 <= 127)
            _ = v11
        case -944:
            var v12 uint32 = 885
            _ = v12
        default:
            fmt.Println(p0, v6)
        }
        _ = fi9
    }
    _ = p0
    _ = v6
    return (!true)
}

func main() {
    var v1 float64 = (0.1 + 1.5)
    v2 := make(map[string]bool)
    var v3 uint16 = 32767
    var v4 uint8 = 2
    var v13 uint64 = 299
    switch 0 {
    case 42:
        switch -458 {
        case 128:
            v3 -= (v3 + v3)
        case 9223372036854775807:
            v4 = 0
        default:
            if ((!(0 <= (v13 << ((1) & 63)))) && true) {
                v14 := GenS2{F0: -2147483648}
                if ("a" < "hello") {
                    var v15 int16 = 138
                    var v16 int16 = (v15 / ((-32768) | 1))
                    if (-1 != -128) {
                        switch 104 {
                        case 203:
                            var v17 int32 = -1
                            _ = v17
                        case 4294967295:
                            var v18 uint64 = (^(v13))
                            _ = v18
                        case -2147483648:
                            switch 129 {
                            case 512:
                                v19 := GenS3{F0: 37, F1: false}
                                _ = v19
                            case -79:
                                var v20 int64 = -470
                                _ = v20
                            }
                        }
                        fmt.Println(v3)
                        for fi21 := 0; fi21 < 8; fi21++ {
                            var v22 int = fi21
                            _ = fi21
                            _ = v22
                        }
                    }
                    _ = v15
                    _ = v16
                } else {
                    if (!false) {
                        switch 32767 {
                        case -70:
                            for fi23 := 0; fi23 < 4; fi23++ {
                                switch 1 {
                                case 127:
                                    for fi24 := 0; fi24 < 14; fi24++ {
                                        var v25 int = fi23
                                        _ = fi24
                                        _ = v25
                                    }
                                default:
                                    var v26 bool = ((false || (-120 == 127)) && (!(("\n\t" + "foo bar") > "a")))
                                    _ = v26
                                }
                                _ = fi23
                            }
                        case -348:
                            v3 = (v3 - uint16(v4))
                        case 417:
                            for fi27 := 0; fi27 < 3; fi27++ {
                                v13 = (v13 << ((42) & 63))
                                _ = fi27
                            }
                        default:
                            v13 -= v13
                        }
                        for fi28 := 0; fi28 < 20; fi28++ {
                            var v29 uint64 = 909
                            fmt.Println(v1)
                            v29 -= v29
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        switch -581 {
                        case 9223372036854775807:
                            delete(v2, "hello")
                        case -566:
                            v30 := GenS4{F0: -148, F1: -2147483648, F2: (^(v3))}
                            _ = v30
                        }
                        v31 := GenS5{F0: 9223372036854775807}
                        _ = v31
                    }
                    var v32 uint16 = v3
                    _ = v32
                }
                _ = v14
            } else {
                if true {
                    if true {
                        var v33 float32 = 1e10
                        v34 := GenS6{F0: 0, F1: ("hello" + "hello"), F2: false}
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 11; fi35++ {
                        var v36 uint32 = 393
                        v13 = uint64(v36)
                        for fi37 := 0; fi37 < 9; fi37++ {
                            v1 += v1
                            v38 := GenS2{F0: 1}
                            _ = fi37
                            _ = v38
                        }
                        _ = fi35
                        _ = v36
                    }
                    fmt.Println(v13, v1)
                } else {
                    switch 873 {
                    case -442:
                        for fi39 := 0; fi39 < 4; fi39++ {
                            for fi40 := 0; fi40 < 14; fi40++ {
                                v41 := GenS1{F0: 463, F1: (^(v4)), F2: v13}
                                _ = fi40
                                _ = v41
                            }
                            v42 := GenS5{F0: 9223372036854775807}
                            _ = fi39
                            _ = v42
                        }
                    case 128:
                        var v43 bool = (!(-9223372036854775808 >= -573))
                        _ = v43
                    }
                }
            }
        }
    case 4294967295:
        var v44 uint16 = v3
        _ = v44
    case 985:
        v45 := v2["a"]
        _ = v45
    default:
        v4 -= v4
    }
    if true {
        fmt.Println(v3, v13)
    } else {
        var v46 string = "a"
        _ = v46
    }
    fmt.Println(v4, v3)
    var v47 bool = (((0.0 / 1.5) < 1e10) && (!false))
    var v48 int64 = 1
    for fi49 := 0; fi49 < 11; fi49++ {
        fmt.Println(v4, v47, v48)
        switch "a" {
        case "a":
            var v50 int64 = 9223372036854775807
            _ = v50
        case "foo bar":
            v47 = v47
        case "\n\t":
            v1 -= v1
        }
        switch (fi49 - fi49) {
        case 4294967295:
            v48 -= -9223372036854775808
        default:
            switch (("hello" + "foo bar") + ("foo bar" + (("hello" + "") + (("" + "") + "\n\t")))) {
            case "hello":
                v13 = v13
            }
        }
        _ = fi49
    }
    var v51 uint16 = v3
    for fi52 := 0; fi52 < 13; fi52++ {
        switch -1 {
        case 127:
            for fi53 := 0; fi53 < 14; fi53++ {
                v54 := GenS1{F0: 42, F1: 106, F2: v13}
                if v47 {
                    fmt.Println(v4, fi53)
                } else {
                    v55 := GenS4{F0: (v48 / ((v48) | 1)), F1: 42, F2: (v3 % (((v3 % (((v3 >> ((631) & 63))) | 1))) | 1))}
                    if (((0.1 / 1e10) + 0.1) > 0.0) {
                        v56 := v2[(("foo bar" + ("\n\t" + "\n\t")) + "foo bar")]
                        v3 -= (v51 - v3)
                        _ = v56
                    }
                    var v57 float64 = (v1 + v1)
                    v13 = (v13 >> ((255) & 63))
                    _ = v55
                    _ = v57
                }
                v13 += (v13 ^ (^(v13)))
                _ = fi53
                _ = v54
            }
        case 9223372036854775807:
            v58 := GenS7{F0: "\n\t", F1: 0}
            _ = v58
        case -1:
            v59 := GenS3{F0: 100, F1: ((0.0 > -2.25) || v47)}
            _ = v59
        default:
            v60 := GenS5{F0: (-(v48))}
            _ = v60
        }
        _ = fi52
    }
    v61 := v2[("" + ((("foo bar" + ("a" + "\n\t")) + "\n\t") + ""))]
    var v62 int64 = v48
    v1 += 0.1
    switch 236 {
    case 2147483647:
        v62 = v62
    case -1:
        if (!v61) {
            v63 := GenS8{F0: v62, F1: 817}
            delete(v2, ("foo bar" + ("a" + "")))
            fmt.Println(v62, v4)
            _ = v63
        } else {
            v47 = true
            v13 -= (v13 ^ v13)
            v64 := GenS7{F0: "a", F1: 1}
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 1; fi65++ {
        var v66 float64 = v1
        v47 = (0.1 > 0.0)
        _ = fi65
        _ = v66
    }
    var v67 uint16 = 255
    var v68 int8 = -128
    fmt.Println(v68)
    v1 = (float64((1.5 / float32(1e10))) / v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v51
    _ = v61
    _ = v62
    _ = v67
    _ = v68
}
