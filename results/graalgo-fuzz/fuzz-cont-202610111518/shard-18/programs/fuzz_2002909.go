package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int32; F2 bool }
type GenS9 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint64, p1 int8, p2 int32) uint16 {
    fmt.Println(p0, p0, p0)
    p0 += (p0 / ((917) | 1))
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = true
    var v3 uint16 = 32767
    v4 := []string{("foo bar" + "\n\t"), "", "\n\t", ("hello" + "")}
    v4 = append(v4, "a")
    _ = len(v4)
    for fi5 := 0; fi5 < 9; fi5++ {
        if (true || true) {
            v2 = (2147483647 == 127)
        }
        v1 = v1
        _ = fi5
    }
    v2 = true
    var v6 float64 = (0.1 - 0.1)
    switch 42 {
    case 128:
        if v2 {
            v7 := GenS1{F0: 77}
            _ = v7
        }
    case 32767:
        for fi8 := 0; fi8 < 18; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 uint64 = 1
                if v2 {
                    if (2147483647 > 1) {
                        fmt.Println(v2, fi9, v1)
                        fmt.Println(v1)
                        if (false || (3.14159 == 1.5)) {
                            if ((fi8 & -9223372036854775808) >= fi8) {
                                var v11 uint8 = 255
                                _ = v11
                            } else {
                                switch "" {
                                case "foo bar":
                                    v10 -= (v10 * (v10 / (((v10 + 18446744073709551615)) | 1)))
                                }
                                fmt.Println(v10, v10)
                                fmt.Println(fi8, v3)
                            }
                            var v12 int64 = (v1 * v1)
                            _ = v12
                        } else {
                            if v2 {
                                switch "hello" {
                                case "foo bar":
                                    v4 = append(v4, "")
                                default:
                                    var v13 float32 = 0.0
                                    _ = v13
                                }
                                if false {
                                    if (!(v2 && (!(2 > 255)))) {
                                        if v2 {
                                            v4 = append(v4, "\n\t")
                                            fmt.Println(fi8, fi8, fi9)
                                            var v14 int32 = 42
                                            _ = v14
                                        }
                                        var v15 string = ""
                                        _ = v15
                                    }
                                }
                                v16 := GenS2{F0: (v1 + (-(v1)))}
                                if (-1 < 42) {
                                    v3 = v3
                                    if (4294967295 > 2) {
                                        if (!(!(-1 >= 2147483647))) {
                                            if false {
                                                var v17 float32 = 1e10
                                                fmt.Println(v2)
                                                _ = len(v4)
                                                _ = v17
                                            }
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            if true {
                                                var v18 uint32 = 892
                                                _ = v18
                                            }
                                        }
                                        switch "" {
                                        case "foo bar":
                                            _ = len(v4)
                                            if ((48 > 255) && (((((-2.25 / 1e10) + (0.1 * 3.14159)) * -2.25) != (3.14159 - (0.0 / 1.5))) || (!true))) {
                                                var v19 uint32 = 42
                                                fmt.Println(v10)
                                                _ = v19
                                            }
                                        case "hello":
                                            switch "foo bar" {
                                            case "foo bar":
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case "a":
                                                if v2 {
                                                    v10 = 4294967295
                                                }
                                            case "":
                                                var v20 bool = true
                                                _ = v20
                                            default:
                                                v21 := GenS3{F0: (-161 < -1), F1: 255, F2: 2147483647}
                                                _ = v21
                                            }
                                        case "a":
                                            v3 += v3
                                        }
                                        var v22 uint32 = 871
                                        _ = v22
                                    }
                                } else {
                                    v23 := GenS1{F0: 127}
                                    v24 := GenS2{F0: (v1 << ((4294967295) & 63))}
                                    v1 = v1
                                    v2 = true
                                    _ = v23
                                    _ = v24
                                }
                                _ = v16
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = ("\n\t" + (("\n\t" + ("\n\t" + "")) + ""))
                                }
                            }
                        }
                    } else {
                        v25 := GenS1{F0: -128}
                        var v26 float64 = v6
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        _ = v25
                        _ = v26
                    }
                    if v2 {
                        v4 = append(v4, ((("\n\t" + (("foo bar" + "") + ("hello" + "a"))) + "") + ("\n\t" + "\n\t")))
                        if v2 {
                            v27 := GenS4{F0: -32768}
                            fmt.Println(v6, v3, v6)
                            _ = v27
                        } else {
                            fmt.Println(v2)
                        }
                    }
                    v28 := GenS5{F0: "foo bar"}
                    _ = v28
                } else {
                    var v29 int16 = 1000
                    v4 = append(v4, "a")
                    _ = v29
                }
                var v30 int = fi8
                _ = fi9
                _ = v10
                _ = v30
            }
            _ = fi8
        }
    default:
        if (490 <= 0) {
            fmt.Println(v6, v2)
        }
    }
    if (2 < 65535) {
        v1 = v1
        v1 -= 0
        v6 += 1e10
    } else {
        if ((738 == 32767) && v2) {
            fmt.Println(v3, v2, v2)
        } else {
            v31 := GenS2{F0: -1}
            _ = v31
        }
    }
    v32 := GenS6{F0: (v1 * -725), F1: -45}
    v6 -= v6
    _ = len(v4)
    for fi33 := 0; fi33 < 5; fi33++ {
        v3 += (v3 >> ((42) & 63))
        _ = fi33
    }
    if (586 == 1) {
        v4 = append(v4, (("hello" + "foo bar") + "a"))
        var v34 int16 = -662
        _ = v34
    } else {
        if true {
            v1 += v1
            v35 := GenS7{F0: ("" + (("hello" + "foo bar") + ("a" + (("hello" + "\n\t") + ("hello" + "")))))}
            v2 = ((-105 < -32) || v2)
            _ = v35
        }
        fmt.Println(v3, v3)
        for fi36 := 0; fi36 < 19; fi36++ {
            _ = len(v4)
            var v37 int16 = 1
            if ((1e10 == (((1e10 * (v6 / v6)) / 3.14159) + v6)) || (127 <= 2)) {
                fmt.Println(fi36, v6)
                var v38 int64 = (-(v1))
                _ = v38
            }
            _ = fi36
            _ = v37
        }
    }
    v39 := GenS8{F0: 1, F1: 1, F2: v2}
    v1 -= (v1 ^ (^(v1)))
    if (v6 == (((0.0 / float64((((1.5 * (0.1 - ((1.5 + ((3.14159 - float32(-2.25)) + (0.1 * (((-2.25 - -2.25) - float32(((-2.25 + v6) / (v6 + -2.25)))) - (3.14159 / 0.1))))) + 0.1))) / 0.1) * float32(v6)))) * v6) * (v6 - 0.1))) {
        if (v2 || (v6 != v6)) {
            v40 := GenS2{F0: v1}
            _ = v40
        }
        if true {
            if (!v2) {
                v2 = true
            } else {
                var v41 uint64 = 18446744073709551615
                for fi42 := 0; fi42 < 11; fi42++ {
                    if ((-(v1)) <= 164) {
                        v1 = 0
                        v3 += (v3 - v3)
                        var v43 float32 = (1.5 - 3.14159)
                        _ = v43
                    } else {
                        var v44 int64 = int64(v3)
                        var v45 int8 = 1
                        _ = v44
                        _ = v45
                    }
                    if (false && v2) {
                        v1 = 42
                    }
                    for fi46 := 0; fi46 < 12; fi46++ {
                        var v47 int32 = 786
                        v6 = float64(float32(0.0))
                        v41 += v41
                        _ = fi46
                        _ = v47
                    }
                    _ = fi42
                }
                _ = v41
            }
            if ((3.14159 > 0.1) && v2) {
                _ = len(v4)
                v4 = append(v4, "hello")
                v48 := GenS5{F0: "foo bar"}
                _ = v48
            }
        } else {
            v3 -= (v3 >> ((65535) & 63))
            var v49 uint32 = 2147483648
            _ = v49
        }
    } else {
        if v2 {
            for fi50 := 0; fi50 < 18; fi50++ {
                switch 4294967295 {
                case -61:
                    var v51 float64 = 1e10
                    _ = v51
                }
                _ = fi50
            }
            if v2 {
                for fi52 := 0; fi52 < 3; fi52++ {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        v54 := GenS9{F0: 2147483648, F1: 689}
                        if len(v4) > 0 {
                            v4[0] = ("a" + "hello")
                        }
                        if len(v4) > 0 {
                            v4[0] = (("" + ((("foo bar" + "\n\t") + "foo bar") + (("hello" + "a") + ("\n\t" + ((("" + "\n\t") + "a") + "a"))))) + "foo bar")
                        }
                        v1 = 1
                        var v55 uint8 = 127
                        v2 = ((-529 == 1) && v2)
                        v1 += -9223372036854775808
                        v56 := GenS3{F0: v2, F1: v55, F2: 813}
                        _ = fi53
                        _ = v54
                        _ = v55
                        _ = v56
                    }
                    for fi57 := 0; fi57 < 1; fi57++ {
                        v2 = true
                        _ = fi57
                    }
                    _ = fi52
                }
                var v58 uint64 = 42
                v59 := GenS1{F0: -128}
                _ = v58
                _ = v59
            }
        } else {
            v3 += 32767
        }
        fmt.Println(v6)
    }
    for fi60 := 0; fi60 < 2; fi60++ {
        if v2 {
            var v61 int8 = 100
            _ = v61
        }
        v4 = append(v4, "\n\t")
        _ = fi60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v6
}
