package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint }
func genFn0(p0 uint64) uint {
    fmt.Println(p0)
    var v7 int = 4294967295
    p0 = (p0 + p0)
    v7 = 9223372036854775807
    _ = p0
    _ = v7
    return uint(p0)
}
func genFn1(p0 string, p1 uint64) int64 {
    p1 = (p1 / ((p1) | 1))
    var v8 bool = true
    v9 := []int{32767}
    for fi10 := 0; fi10 < 8; fi10++ {
        if false {
            p0 = "a"
            if true {
                for fi11 := 0; fi11 < 7; fi11++ {
                    var v12 string = "foo bar"
                    v12 = (("" + "a") + v12)
                    v13 := GenS1{F0: 86}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v14 := GenS1{F0: 128}
                _ = v14
            }
            var v15 float64 = 1e10
            _ = v15
        }
        if len(v9) > 0 {
            v9[0] = fi10
        }
        switch p0 {
        case "a":
            fmt.Println(fi10)
        default:
            p0 = p0
        }
        _ = fi10
    }
    if len(v9) > 0 {
        v9[0] = 32767
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float32 = ((((((float32(0.1) + 1.5) - 1.5) * 0.1) + float32(1e10)) - -2.25) - 1e10)
    var v2 int = -2147483648
    var v3 int16 = 570
    var v4 uint8 = 1
    var v5 int = (v2 << ((426) & 63))
    v6 := []string{"", ("\n\t" + ("foo bar" + ("a" + ((("hello" + ("foo bar" + "")) + ((("a" + "hello") + ((("a" + "hello") + ("" + "")) + "\n\t")) + "foo bar")) + "a")))), "hello", "a"}
    var v16 string = "foo bar"
    v2 += v2
    if ((v5 ^ 127) != v5) {
        var v17 int16 = 1
        switch (-(v5)) {
        case -2147483648:
            v18 := GenS1{F0: (v4 & v4)}
            _ = v18
        case -377:
            for fi19 := 0; fi19 < 1; fi19++ {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v21 := GenS1{F0: 2}
                    v22 := GenS1{F0: (v4 - v4)}
                    var v23 uint16 = 65535
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi19
            }
        }
        fmt.Println(v5)
        _ = v17
    } else {
        v2 += -851
    }
    switch (v5 | 1) {
    case 32768:
        var v24 uint16 = 32767
        _ = v24
    }
    v2 = int(v3)
    fmt.Println(v4, v1, v5)
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v16, v2, v3)
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if false {
                var v27 int = (fi25 ^ -234)
                v1 += v1
                _ = v27
            } else {
                v5 -= (v5 >> ((4294967295) & 63))
            }
            if ((v1 * (0.1 / 1e10)) <= 1.5) {
                fmt.Println(v4, fi26, v1)
                if (!((!(!true)) || (430 != 978))) {
                    fmt.Println(v16, fi26, v16)
                    var v28 bool = (false || (0 != 1))
                    v6 = append(v6, "")
                    v3 += (-(v3))
                    _ = v28
                }
                v16 = ("a" + "foo bar")
            } else {
                var v29 float32 = (-2.25 / v1)
                var v30 int = -489
                _ = v29
                _ = v30
            }
            if (genFn1(v16, 948) <= 0) {
                var v31 uint64 = 18446744073709551615
                _ = v31
            } else {
                v5 = 32768
                var v32 float32 = v1
                _ = v32
            }
            _ = fi26
        }
        _ = fi25
    }
    _ = len(v6)
    if ((v2 & 150) == v5) {
        fmt.Println(v1)
        if false {
            if false {
                switch "" {
                case "a":
                    if false {
                        if (-136 < (v2 + v5)) {
                            if (0.1 > v1) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if (736 > 2147483647) {
                                    switch "" {
                                    case "foo bar":
                                        switch "hello" {
                                        case "\n\t":
                                            switch ("a" + v16) {
                                            case "hello":
                                                var v33 float32 = (float32(1.5) * (v1 - (3.14159 - 1.5)))
                                                _ = v33
                                            case "a":
                                                v16 = v16
                                            default:
                                                fmt.Println(v3)
                                            }
                                        case "":
                                            _ = len(v6)
                                        case "a":
                                            v5 = -9223372036854775808
                                        default:
                                            v1 = (((1.5 * ((-2.25 * v1) * v1)) / 1.0) / v1)
                                        }
                                    case "":
                                        if len(v6) > 0 {
                                            v6[0] = v16
                                        }
                                    case "a":
                                        v34 := GenS1{F0: v4}
                                        _ = v34
                                    default:
                                        var v35 float64 = 1.5
                                        _ = v35
                                    }
                                } else {
                                    if false {
                                        v2 = -2147483648
                                        switch v16 {
                                        case "a":
                                            switch "" {
                                            case "hello":
                                                v6 = append(v6, "")
                                            case "a":
                                                if (2147483647 != -328) {
                                                    v4 = (v4 & 210)
                                                }
                                            }
                                        case "foo bar":
                                            v5 += v5
                                        case "\n\t":
                                            var v36 uint16 = 173
                                            _ = v36
                                        default:
                                            var v37 int16 = v3
                                            _ = v37
                                        }
                                    } else {
                                        fmt.Println(v3, v16, v3)
                                    }
                                    v5 += (v5 / (((^(v5))) | 1))
                                }
                                for fi38 := 0; fi38 < 16; fi38++ {
                                    v3 += int16(v4)
                                    _ = fi38
                                }
                            } else {
                                for fi39 := 0; fi39 < 5; fi39++ {
                                    v16 = v16
                                    if ("" < v16) {
                                        v40 := GenS1{F0: (v4 / ((v4) | 1))}
                                        switch (v16 + v16) {
                                        case "a":
                                            v41 := GenS1{F0: (v4 << ((genFn0(713)) & 63))}
                                            _ = v41
                                        }
                                        _ = v40
                                    }
                                    v16 = ("" + v16)
                                    _ = fi39
                                }
                                v4 -= v4
                                var v42 int8 = 127
                                _ = v42
                            }
                        }
                        for fi43 := 0; fi43 < 16; fi43++ {
                            var v44 uint8 = v4
                            var v45 uint = 98
                            var v46 uint16 = 0
                            _ = fi43
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        v2 = v5
                    } else {
                        var v47 string = "foo bar"
                        v4 -= 127
                        if (false || true) {
                            if (18446744073709551615 == 18446744073709551615) {
                                var v48 int8 = -1
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    v6 = append(v6, ("foo bar" + v47))
                                    _ = fi49
                                }
                                fmt.Println(v4, v16, v16)
                                _ = v48
                            } else {
                                v16 = v16
                                var v50 uint = 65535
                                _ = v50
                            }
                        } else {
                            v5 += (v2 >> ((42) & 63))
                            v2 -= -9223372036854775808
                            v6 = append(v6, v47)
                        }
                        _ = v47
                    }
                case "\n\t":
                    for fi51 := 0; fi51 < 19; fi51++ {
                        var v52 uint16 = 42
                        _ = fi51
                        _ = v52
                    }
                default:
                    if len(v6) > 0 {
                        v6[0] = v16
                    }
                }
                for fi53 := 0; fi53 < 16; fi53++ {
                    var v54 bool = false
                    _ = fi53
                    _ = v54
                }
                if false {
                    fmt.Println(v3, v3, v4)
                    v3 = 1000
                }
            }
            switch v16 {
            case "a":
                for fi55 := 0; fi55 < 5; fi55++ {
                    v1 = ((1.5 - (v1 * v1)) + ((1.5 - (v1 / v1)) * 3.14159))
                    if ((v5 % (((v5 >> ((42) & 63))) | 1)) >= 358) {
                        if true {
                            switch ((v16 + "") + v16) {
                            case "":
                                var v56 uint16 = 0
                                _ = v56
                            }
                        }
                        for fi57 := 0; fi57 < 13; fi57++ {
                            v5 = v2
                            var v58 uint8 = (v4 | v4)
                            if true {
                                v1 = v1
                                fmt.Println(v58, v2)
                            } else {
                                var v59 int8 = -128
                                v16 = v16
                                switch 129 {
                                case 32767:
                                    v16 = (v16 + "hello")
                                }
                                _ = v59
                            }
                            v58 = (v4 & v58)
                            var v60 uint32 = 1
                            _ = fi57
                            _ = v58
                            _ = v60
                        }
                    }
                    _ = fi55
                }
            case "\n\t":
                v3 = v3
            case "":
                for fi61 := 0; fi61 < 7; fi61++ {
                    if true {
                        v3 = (v3 / (((v3 << ((255) & 63))) | 1))
                    } else {
                        v62 := GenS1{F0: (v4 >> ((42) & 63))}
                        _ = v62
                    }
                    fmt.Println(v1, v3, v2)
                    _ = fi61
                }
            }
            v2 -= (v5 - (v2 << ((42) & 63)))
        } else {
            var v63 int8 = 1
            v1 = ((1e10 - v1) / float32(1e10))
            _ = v63
        }
        v2 -= 42
    }
    v64 := GenS1{F0: v4}
    for fi65 := 0; fi65 < 16; fi65++ {
        for fi66 := 0; fi66 < 9; fi66++ {
            var v67 int32 = 0
            if (!true) {
                var v68 uint64 = 2
                _ = v68
            } else {
                fmt.Println(v1, v16, fi65)
                v69 := GenS2{F0: (fi66 < 9223372036854775807), F1: (v3 == -1), F2: 0}
                v5 = (v2 + 32768)
                _ = v69
            }
            fmt.Println(v4, v16)
            _ = fi66
            _ = v67
        }
        _ = fi65
    }
    v70 := GenS1{F0: 42}
    v4 = (v4 + (v4 ^ v4))
    for fi71 := 0; fi71 < 10; fi71++ {
        for fi72 := 0; fi72 < 4; fi72++ {
            var v73 float32 = ((v1 - ((1e10 * -2.25) - 1.5)) + 3.14159)
            v74 := GenS3{F0: 0, F1: 891}
            v75 := GenS4{F0: 838, F1: 1}
            _ = fi72
            _ = v73
            _ = v74
            _ = v75
        }
        var v76 string = "\n\t"
        if false {
            var v77 float64 = 3.14159
            v76 = v16
            var v78 uint32 = 4294967295
            _ = v77
            _ = v78
        }
        _ = fi71
        _ = v76
    }
    v4 = (v4 << ((0) & 63))
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v70
}
