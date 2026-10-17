package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint; F2 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 int64 }
type GenS5 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 int64) int8 {
    p0 = (^(p0))
    p0 = (p0 >> ((2) & 63))
    v5 := []int{4294967295, -882, 4294967295, 128}
    _ = p0
    _ = v5
    return -128
}

func main() {
    var v1 string = ""
    var v2 float32 = (0.1 + (0.1 - 0.1))
    v3 := make(map[string]int)
    var v4 int32 = -812
    var v6 int32 = 379
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: "a", F1: false, F2: -2147483648}
        _ = fi7
        _ = v8
    }
    v9 := []int{}
    switch 42 {
    case -9223372036854775808:
        v1 = ("a" + v1)
    case 2147483647:
        var v10 float64 = 3.14159
        _ = v10
    case -80:
        fmt.Println(v4)
    default:
        var v11 bool = ((2 <= 901) || (genFn0(-574) < -1))
        _ = v11
    }
    v6 += (v4 / ((v4) | 1))
    v12 := GenS1{F0: ((("hello" + v1) + (v1 + (v1 + v1))) + v1), F1: (4294967295 >= 18446744073709551615), F2: (v6 ^ -341)}
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = ((1 == 0) || false)
        var v15 uint32 = 2
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        if (!((1e10 >= 1.5) && (2 > 2147483647))) {
            for fi16 := 0; fi16 < 3; fi16++ {
                v9 = append(v9, fi16)
                _ = fi16
            }
            fmt.Println(v1, v6, v6)
        } else {
            v17 := GenS1{F0: ("foo bar" + (v1 + ("a" + v1))), F1: true, F2: v6}
            _ = v17
        }
        switch 4294967295 {
        case -941:
            v4 = (v4 ^ (v6 >> ((65535) & 63)))
        case -871:
            if (!false) {
                for fi18 := 0; fi18 < 3; fi18++ {
                    v3[(v1 + v1)] = -9223372036854775808
                    if false {
                        if (true || false) {
                            delete(v3, v1)
                            switch (fi18 << ((65535) & 63)) {
                            case 347:
                                v19 := GenS1{F0: (v1 + v1), F1: false, F2: 42}
                                _ = v19
                            case 4294967295:
                                var v20 int16 = 32767
                                _ = v20
                            default:
                                if false {
                                    if false {
                                        for fi21 := 0; fi21 < 9; fi21++ {
                                            v2 = v2
                                            _ = fi21
                                        }
                                        var v22 string = v1
                                        _ = v22
                                    } else {
                                        var v23 float64 = 1e10
                                        v23 = ((1e10 * float64(v2)) + (v23 - (1e10 / ((1.5 - v23) - v23))))
                                        _ = v23
                                    }
                                    fmt.Println(fi18, v1, v2)
                                } else {
                                    if false {
                                        if (((fi18 >> ((65535) & 63)) <= (fi18 & fi18)) && (((((((float32((float64((float32(0.0) + v2)) + 1.5)) - ((float32(0.0) / v2) / v2)) + float32(3.14159)) + v2) + 1.5) - (v2 / 1.0)) - float32(1e10)) > v2)) {
                                            switch fi18 {
                                            case 32767:
                                                for fi24 := 0; fi24 < 20; fi24++ {
                                                    v4 = v6
                                                    v4 = (v6 << ((255) & 63))
                                                    v4 += (v4 % ((v6) | 1))
                                                    if (!((((v4 * 1) <= 2147483647) && true) || true)) {
                                                        if (false || false) {
                                                            v4 = 497
                                                            v25 := GenS1{F0: v1, F1: (!(!(42 < 127))), F2: (v6 | v4)}
                                                            _ = v25
                                                        } else {
                                                            switch v1 {
                                                            case "\n\t":
                                                                var v26 uint = 1
                                                                _ = v26
                                                            default:
                                                                var v27 int64 = -9223372036854775808
                                                                _ = v27
                                                            }
                                                            if len(v9) > 0 {
                                                                v9[0] = (fi24 ^ 0)
                                                            }
                                                        }
                                                        v4 += -2147483648
                                                        fmt.Println(v6)
                                                        var v28 uint = 738
                                                        _ = v28
                                                    }
                                                    _ = fi24
                                                }
                                            case 9223372036854775807:
                                                _ = len(v3)
                                            case 861:
                                                for fi29 := 0; fi29 < 12; fi29++ {
                                                    v2 += -2.25
                                                    v1 = "\n\t"
                                                    v2 -= (v2 - v2)
                                                    v1 = v1
                                                    _ = fi29
                                                }
                                            }
                                            v30 := GenS2{F0: "foo bar"}
                                            v4 = v4
                                            _ = v30
                                        } else {
                                            if (42 >= 0) {
                                                v6 -= v4
                                            }
                                            for fi31 := 0; fi31 < 10; fi31++ {
                                                var v32 float32 = v2
                                                v4 = 2147483647
                                                var v33 uint = 18446744073709551615
                                                _ = fi31
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                    } else {
                                        v1 = (v1 + ((v1 + "") + v1))
                                        var v34 uint64 = 2
                                        switch (fi18 & 631) {
                                        case -9223372036854775808:
                                            var v35 int32 = -1
                                            _ = v35
                                        case 128:
                                            if (!true) {
                                                for fi36 := 0; fi36 < 3; fi36++ {
                                                    if false {
                                                        switch (("\n\t" + "\n\t") + "a") {
                                                        case "":
                                                            switch (v1 + (((("a" + "a") + (v1 + "hello")) + "hello") + (v1 + v1))) {
                                                            case "\n\t":
                                                                fmt.Println(v2, v2)
                                                            }
                                                        case "foo bar":
                                                            var v37 uint16 = 32767
                                                            _ = v37
                                                        default:
                                                            var v38 int = (fi36 + 276)
                                                            _ = v38
                                                        }
                                                    }
                                                    var v39 uint16 = 65535
                                                    var v40 int = -1
                                                    _ = fi36
                                                    _ = v39
                                                    _ = v40
                                                }
                                                switch (fi18 % ((fi18) | 1)) {
                                                case 4294967295:
                                                    switch ("hello" + v1) {
                                                    case "":
                                                        switch "" {
                                                        case "":
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                        case "foo bar":
                                                            v3[v1] = (fi18 << ((65535) & 63))
                                                        case "a":
                                                            var v41 int16 = -32768
                                                            _ = v41
                                                        default:
                                                            _ = len(v9)
                                                        }
                                                    }
                                                default:
                                                    if true {
                                                        if (1 != 4294967295) {
                                                            v42 := GenS3{F0: -336, F1: 4294967295, F2: (true && false)}
                                                            v6 = (-(v6))
                                                            _ = v42
                                                        }
                                                    } else {
                                                        for fi43 := 0; fi43 < 1; fi43++ {
                                                            var v44 int = (-(fi18))
                                                            switch v1 {
                                                            case "\n\t":
                                                                if (127 >= 228) {
                                                                    v45 := GenS2{F0: v1}
                                                                    var v46 uint16 = 65535
                                                                    _ = v45
                                                                    _ = v46
                                                                }
                                                            case "":
                                                                var v47 int8 = 100
                                                                _ = v47
                                                            default:
                                                                v2 = (((3.14159 / v2) / (float32((1.5 * 0.1)) / 1.5)) * 1.5)
                                                            }
                                                            fmt.Println(v6, v4, fi18)
                                                            _ = fi43
                                                            _ = v44
                                                        }
                                                        fmt.Println(fi18, v4)
                                                    }
                                                }
                                                var v48 string = (v1 + "foo bar")
                                                _ = v48
                                            } else {
                                                if len(v9) >= 2 {
                                                    v9 = v9[1:len(v9)]
                                                }
                                                for fi49 := 0; fi49 < 9; fi49++ {
                                                    v3[v1] = (fi18 | fi18)
                                                    fmt.Println(v2, v34)
                                                    _ = fi49
                                                }
                                            }
                                        }
                                        _ = v34
                                    }
                                    for fi50 := 0; fi50 < 16; fi50++ {
                                        fmt.Println(v6, fi50)
                                        var v51 uint16 = 32767
                                        _ = fi50
                                        _ = v51
                                    }
                                }
                            }
                        } else {
                            for fi52 := 0; fi52 < 9; fi52++ {
                                v4 += (v6 << ((4294967295) & 63))
                                if true {
                                    var v53 uint64 = 539
                                    _ = v53
                                } else {
                                    var v54 int8 = 7
                                    if true {
                                        switch (fi18 * 128) {
                                        case 1:
                                            var v55 float64 = (0.1 - 0.0)
                                            _ = v55
                                        case 127:
                                            delete(v3, v1)
                                        case 32767:
                                            var v56 uint = 255
                                            _ = v56
                                        }
                                        switch (fi52 - fi52) {
                                        case 42:
                                            fmt.Println(v1)
                                        case 127:
                                            v3["foo bar"] = (fi52 ^ fi18)
                                            var v57 float32 = v2
                                            v2 -= float32(0.0)
                                            switch v1 {
                                            case "":
                                                var v58 int32 = 0
                                                _ = v58
                                            }
                                            v54 = (v54 | genFn0(42))
                                            _ = v57
                                        default:
                                            switch "hello" {
                                            case "":
                                                v2 -= v2
                                            default:
                                                if false {
                                                    v6 = v4
                                                    var v59 int8 = v54
                                                    _ = v59
                                                }
                                            }
                                        }
                                    } else {
                                        var v60 bool = false
                                        _ = v60
                                    }
                                    var v61 uint = 18446744073709551615
                                    _ = v54
                                    _ = v61
                                }
                                fmt.Println(v1)
                                _ = fi52
                            }
                            fmt.Println(v2, v6, v6)
                            for fi62 := 0; fi62 < 12; fi62++ {
                                switch 394 {
                                case 0:
                                    var v63 float64 = 3.14159
                                    _ = v63
                                }
                                v1 = ""
                                var v64 float32 = (1e10 / v2)
                                _ = fi62
                                _ = v64
                            }
                        }
                    } else {
                        fmt.Println(v6)
                        var v65 float32 = v2
                        var v66 string = ("\n\t" + ((v1 + ((v1 + "a") + (("\n\t" + "a") + v1))) + (v1 + v1)))
                        _ = v65
                        _ = v66
                    }
                    for fi67 := 0; fi67 < 1; fi67++ {
                        v68 := GenS1{F0: (v1 + (v1 + v1)), F1: (1 <= 45), F2: (v6 * (v6 - -1))}
                        _ = fi67
                        _ = v68
                    }
                    _ = fi18
                }
                v2 += float32((-2.25 * -2.25))
                switch ("hello" + v1) {
                case "foo bar":
                    fmt.Println(v4, v2, v2)
                case "a":
                    var v69 uint = 42
                    _ = v69
                default:
                    for fi70 := 0; fi70 < 16; fi70++ {
                        var v71 int64 = 615
                        v72 := GenS4{F0: -1, F1: 0, F2: 0}
                        for fi73 := 0; fi73 < 8; fi73++ {
                            fmt.Println(v71, v2, v1)
                            _ = fi73
                        }
                        _ = fi70
                        _ = v71
                        _ = v72
                    }
                }
            } else {
                switch ((v1 + v1) + v1) {
                case "a":
                    v6 = 0
                case "\n\t":
                    switch -618 {
                    case 127:
                        v6 += v4
                    case 42:
                        v6 += v6
                    }
                case "hello":
                    v6 += 1
                default:
                    fmt.Println(v6)
                }
                v74 := GenS5{F0: 9223372036854775807, F1: 100, F2: -627}
                switch 127 {
                case 32767:
                    v6 += 144
                }
                _ = v74
            }
        case 9223372036854775807:
            v2 = v2
        }
        v75 := GenS5{F0: 42, F1: 42, F2: 1}
        _ = v75
    }
    fmt.Println(v4)
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
