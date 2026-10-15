package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 int8; F1 string }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{247, 65535, 42, 319}
    var v3 float64 = 3.14159
    v4 := []int{42, 774, -1}
    switch 42 {
    case 897:
        v5 := GenS1{F0: "\n\t"}
        _ = v5
    case 127:
        v3 = v3
    case 1:
        var v6 float32 = -2.25
        _ = v6
    default:
        var v7 bool = ((false && ((1 != 966) && (false || true))) && true)
        _ = v7
    }
    var v8 bool = false
    if true {
        if v8 {
            var v9 int = -9223372036854775808
            _ = v9
        }
        v10 := GenS1{F0: "\n\t"}
        if v8 {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS1{F0: ("foo bar" + ("a" + ("" + "")))}
                _ = fi11
                _ = v12
            }
            for fi13 := 0; fi13 < 3; fi13++ {
                fmt.Println(v3, v3)
                for fi14 := 0; fi14 < 3; fi14++ {
                    v8 = (1 >= 42)
                    v8 = false
                    if (!false) {
                        var v15 int64 = -401
                        if len(v4) > 0 {
                            v4[0] = fi13
                        }
                        v15 -= (v15 * (-(v15)))
                        v1[fi13] = ("foo bar" + "")
                        _ = v15
                    }
                    _ = fi14
                }
                _ = fi13
            }
            if (-1 != -44) {
                switch ("hello" + "") {
                case "hello":
                    var v16 uint16 = 42
                    _ = v16
                }
                switch ("foo bar" + ("foo bar" + "hello")) {
                case "a":
                    if (1 < 767) {
                        var v17 uint = 0
                        v1[42] = "\n\t"
                        _ = v17
                    }
                case "foo bar":
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v19 := GenS1{F0: "hello"}
                        _ = fi18
                        _ = v19
                    }
                default:
                    if v8 {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v8 = v8
                            v3 -= (1.5 / (float64((3.14159 * 3.14159)) / (-2.25 / 1.0)))
                            _ = fi20
                        }
                        for fi21 := 0; fi21 < 11; fi21++ {
                            fmt.Println(v8, v8)
                            if len(v4) > 0 {
                                v4[0] = fi21
                            }
                            _ = fi21
                        }
                        var v22 uint64 = 42
                        _ = v22
                    }
                }
                switch "hello" {
                case "\n\t":
                    if (!true) {
                        v8 = (v8 && ((32767 < 786) && ((!false) || v8)))
                        v1[-558] = "\n\t"
                        var v23 uint32 = 9
                        _ = v23
                    } else {
                        fmt.Println(v3, v3)
                        for fi24 := 0; fi24 < 18; fi24++ {
                            for fi25 := 0; fi25 < 11; fi25++ {
                                if ((9223372036854775807 < 4294967295) || false) {
                                    fmt.Println(fi24, fi24, v8)
                                }
                                _ = fi25
                            }
                            _ = fi24
                        }
                        var v26 bool = (v8 && (v8 || (214 >= 735)))
                        _ = v26
                    }
                case "a":
                    if false {
                        if false {
                            v27 := GenS2{F0: 2, F1: -1}
                            var v28 string = ""
                            _ = v27
                            _ = v28
                        }
                    } else {
                        var v29 bool = false
                        _ = v29
                    }
                default:
                    fmt.Println(v3)
                }
            } else {
                if v8 {
                    if true {
                        for fi30 := 0; fi30 < 20; fi30++ {
                            delete(v1, fi30)
                            var v31 int32 = -1
                            v32 := GenS3{F0: 0, F1: "hello"}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        var v33 bool = true
                        _ = v33
                    }
                    fmt.Println(v8, v8)
                } else {
                    var v34 string = "\n\t"
                    v8 = (v8 && ((0.0 / 0.1) < (((((0.0 / ((1.5 * 1.5) - float32((0.0 / v3)))) * (0.0 / -2.25)) * ((1.5 - (3.14159 - 1e10)) + 1.5)) / 1e10) / 1.0)))
                    v35 := GenS2{F0: 134, F1: 908}
                    _ = v34
                    _ = v35
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    for fi37 := 0; fi37 < 13; fi37++ {
                        var v38 int = -1
                        _ = len(v4)
                        var v39 uint = 0
                        _ = fi37
                        _ = v38
                        _ = v39
                    }
                    v40 := GenS3{F0: -1, F1: ("a" + "foo bar")}
                    _ = fi36
                    _ = v40
                }
            }
        }
        _ = v10
    } else {
        v3 = v3
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v1)
    }
    var v41 bool = (v8 && v8)
    var v42 uint = 255
    if (!((1 > -385) && ((222 > 0) || (v3 != float64(1.5))))) {
        var v43 int = 32768
        v43 = v43
        for fi44 := 0; fi44 < 16; fi44++ {
            var v45 int = v43
            if true {
                v43 = 2147483647
                if v8 {
                    v41 = false
                }
            } else {
                v41 = false
                v41 = (!v41)
            }
            _ = fi44
            _ = v45
        }
        _ = v43
    }
    fmt.Println(v41)
    for fi46 := 0; fi46 < 18; fi46++ {
        for fi47 := 0; fi47 < 15; fi47++ {
            v3 = ((3.14159 + (v3 / ((((v3 + v3) * float64(3.14159)) * 0.1) * float64(1e10)))) * v3)
            var v48 string = "foo bar"
            _ = fi47
            _ = v48
        }
        for fi49 := 0; fi49 < 1; fi49++ {
            v41 = (0 > 476)
            fmt.Println(fi46, fi46, v41)
            if v41 {
                v8 = (!(1 >= -1))
                v50 := v1[(fi46 ^ 32767)]
                _ = v50
            } else {
                v8 = (-2.25 < -2.25)
                switch (fi49 >> ((4294967295) & 63)) {
                case 1:
                    var v51 uint32 = 420
                    _ = v51
                case 454:
                    v3 += float64(float32(-2.25))
                case 128:
                    switch (fi49 - fi49) {
                    case 1:
                        if (!v41) {
                            var v52 int64 = -308
                            v42 = v42
                            v53 := GenS2{F0: v42, F1: 755}
                            _ = v52
                            _ = v53
                        }
                    default:
                        fmt.Println(v41, v41)
                    }
                    var v54 bool = v8
                    var v55 uint8 = 1
                    _ = v54
                    _ = v55
                default:
                    v42 += v42
                }
                switch 453 {
                case 4294967295:
                    v42 = 306
                case 371:
                    var v56 int32 = 1
                    _ = v56
                case 2147483647:
                    if (!(true && ((v41 || v8) && v41))) {
                        fmt.Println(fi46)
                    }
                }
            }
            _ = fi49
        }
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v8
}
