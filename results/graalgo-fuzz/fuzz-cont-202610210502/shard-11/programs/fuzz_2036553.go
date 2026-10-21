package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint8; F2 bool }
type GenS7 struct { F0 int8 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]int)
    var v3 string = "a"
    var v4 int64 = 9223372036854775807
    var v5 float32 = ((((-2.25 / float32(0.1)) + (0.1 * (1.5 / 3.14159))) * ((3.14159 - 1e10) * (1e10 / (1.5 + 1e10)))) * (float32((1e10 * float64(3.14159))) - float32(0.1)))
    var v6 int64 = (-(v4))
    fmt.Println(v4, v4)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 bool = (0 <= 21)
        v4 = v6
        fmt.Println(v1, v3, fi7)
        _ = fi7
        _ = v8
    }
    if false {
        if false {
            for fi9 := 0; fi9 < 15; fi9++ {
                var v10 uint16 = 32767
                _ = fi9
                _ = v10
            }
            fmt.Println(v6, v5, v1)
        } else {
            var v11 int32 = 2147483647
            v2[v3] = 1
            _ = v11
        }
    } else {
        switch -9223372036854775808 {
        case 363:
            _ = len(v2)
        case 127:
            v4 = (v4 | 1)
        }
        var v12 int32 = -986
        _ = v12
    }
    v5 = float32((1.5 / 1e10))
    if false {
        for fi13 := 0; fi13 < 6; fi13++ {
            v6 += (v4 | v6)
            _ = fi13
        }
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 16; fi15++ {
                v5 -= v5
                _ = fi15
            }
            _ = fi14
        }
        v5 -= v5
    }
    var v16 int16 = -32768
    switch 128 {
    case 189:
        if false {
            var v17 int32 = 42
            _ = v17
        } else {
            v6 -= (v4 % ((v6) | 1))
        }
    case -519:
        v18 := GenS1{F0: 1, F1: 42}
        _ = v18
    }
    if (!(18446744073709551615 > 208)) {
        v19 := v2[("" + "\n\t")]
        for fi20 := 0; fi20 < 19; fi20++ {
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := GenS2{F0: v3, F1: v4}
                switch v3 {
                case "foo bar":
                    v23 := []int{fi21}
                    _ = v23
                case "":
                    v24 := GenS3{F0: v16}
                    _ = v24
                default:
                    var v25 bool = true
                    switch (v19 ^ 1) {
                    case 32767:
                        if v25 {
                            v26 := GenS2{F0: (((((v3 + v3) + v3) + v3) + v3) + ((("foo bar" + (("hello" + v3) + (v3 + v3))) + "") + v3)), F1: -1}
                            if ((false && ((v19 >> ((255) & 63)) == fi21)) && (v25 || (v25 && (v25 || (!true))))) {
                                v27 := GenS4{F0: 42, F1: (v1 * v1), F2: (v16 << ((42) & 63))}
                                _ = v27
                            } else {
                                var v28 int = int(v1)
                                v25 = (1 < 255)
                                if (351 <= 32767) {
                                    v29 := GenS4{F0: 0, F1: v1, F2: v16}
                                    _ = v29
                                }
                                v25 = (1.5 >= (-2.25 / 1.0))
                                _ = v28
                            }
                            v30 := GenS5{F0: false}
                            if ((2 < 678) || ((fi20 | (fi21 | (-(fi20)))) >= v19)) {
                                fmt.Println(v3, v5, fi20)
                            }
                            fmt.Println(v16)
                            if (-2.25 != 3.14159) {
                                var v31 int32 = 2147483647
                                _ = v31
                            } else {
                                v4 -= (v4 * v6)
                                switch (v3 + ("\n\t" + ("" + v3))) {
                                case "foo bar":
                                    var v32 int64 = (v6 ^ (v4 >> ((255) & 63)))
                                    _ = v32
                                case "hello":
                                    var v33 uint8 = 255
                                    _ = v33
                                case "\n\t":
                                    if (!((v25 || (1.5 < float64(3.14159))) || false)) {
                                        v3 = v3
                                    }
                                }
                                v19 = fi20
                                var v34 uint32 = v1
                                _ = v34
                            }
                            _ = v26
                            _ = v30
                        }
                    case 0:
                        v1 += 2
                    }
                    fmt.Println(fi21, v4, v3)
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        if (float64(v5) < -2.25) {
            fmt.Println(v5)
            var v35 int32 = -1
            _ = v35
        } else {
            v5 = -2.25
        }
        _ = v19
    }
    var v36 string = ""
    for fi37 := 0; fi37 < 2; fi37++ {
        var v38 bool = true
        _ = fi37
        _ = v38
    }
    var v39 uint32 = v1
    v40 := []int{}
    var v41 uint = 527
    v16 += 1
    for fi42 := 0; fi42 < 16; fi42++ {
        var v43 bool = (!false)
        for fi44 := 0; fi44 < 15; fi44++ {
            if true {
                fmt.Println(v6)
                v45 := GenS6{F0: (^(v39)), F1: 11, F2: (v43 && true)}
                switch v3 {
                case "":
                    v46 := v2[""]
                    _ = v46
                case "foo bar":
                    v43 = (!false)
                case "\n\t":
                    _ = len(v40)
                default:
                    if (!(!v43)) {
                        v47 := GenS6{F0: (v1 >> ((42) & 63)), F1: 1, F2: false}
                        v48 := GenS4{F0: -106, F1: v39, F2: 1}
                        _ = v47
                        _ = v48
                    }
                }
                _ = v45
            } else {
                v49 := GenS7{F0: -5}
                _ = v49
            }
            v41 = v41
            v40 = append(v40, fi44)
            _ = fi44
        }
        _ = fi42
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
