package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 uint8; F1 bool; F2 bool }
func genFn0(p0 uint64) int32 {
    v7 := GenS1{F0: 100, F1: 838, F2: -22}
    p0 -= (^(p0))
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = (fi8 % ((fi8) | 1))
        var v10 float64 = ((0.0 + 3.14159) / float64(3.14159))
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = (p0 * (p0 % ((p0) | 1)))
    _ = p0
    _ = v7
    return -686
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = -392
    var v3 uint = 65535
    var v4 uint16 = 449
    var v5 bool = true
    v6 := []int{}
    var v11 float32 = (3.14159 * (-2.25 / ((((3.14159 - (0.1 - (float32(-2.25) * 0.1))) + 1e10) / (3.14159 / 0.1)) - (1e10 / (1.5 + 1.5)))))
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v6 = append(v6, (fi12 % ((777) | 1)))
            fmt.Println(v5, v3, v11)
            switch "" {
            case "hello":
                v14 := GenS1{F0: 66, F1: 18446744073709551615, F2: -1}
                _ = v14
            case "\n\t":
                var v15 int16 = (v2 & v2)
                _ = v15
            case "a":
                switch "foo bar" {
                case "\n\t":
                    var v16 int = fi13
                    _ = v16
                case "a":
                    v17 := GenS1{F0: 1, F1: 1, F2: 1}
                    _ = v17
                }
            default:
                v5 = v5
            }
            _ = fi13
        }
        if false {
            if true {
                if ((!(!(!(0.1 < ((1.5 * -2.25) + (float64(v11) - (-2.25 * 1.5))))))) && true) {
                    v18 := GenS1{F0: 0, F1: (v3 & v3), F2: 99}
                    var v19 uint8 = 255
                    var v20 uint8 = v19
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    var v21 int = fi12
                    v22 := GenS2{F0: "\n\t"}
                    _ = v21
                    _ = v22
                }
            }
            fmt.Println(v11, fi12)
            var v23 uint64 = 4294967295
            _ = v23
        } else {
            fmt.Println(v3, v11, v11)
            if (v5 || v5) {
                delete(v1, "")
                v4 -= 65535
                for fi24 := 0; fi24 < 15; fi24++ {
                    switch (-(fi12)) {
                    case -1:
                        if (!("\n\t" < "foo bar")) {
                            v2 = (v2 + int16(fi12))
                            v25 := GenS2{F0: (((("a" + ("" + "")) + "") + "a") + "hello")}
                            _ = v25
                        }
                    case -2147483648:
                        v1[("foo bar" + ("" + "a"))] = "\n\t"
                    case -862:
                        v26 := GenS2{F0: "hello"}
                        _ = v26
                    }
                    if v5 {
                        v27 := GenS3{F0: -103}
                        switch "\n\t" {
                        case "foo bar":
                            var v28 uint32 = 65535
                            _ = v28
                        case "a":
                            var v29 float64 = -2.25
                            _ = v29
                        case "hello":
                            v30 := GenS2{F0: "\n\t"}
                            _ = v30
                        default:
                            var v31 uint32 = 0
                            _ = v31
                        }
                        if len(v6) > 0 {
                            v6[0] = 32768
                        }
                        if (!(17 >= 2)) {
                            v3 -= v3
                            v32 := GenS4{F0: v4, F1: 65535, F2: (v4 | (v4 | v4))}
                            if (v5 && (!true)) {
                                var v33 uint16 = (v4 >> ((255) & 63))
                                _ = v33
                            } else {
                                v34 := GenS5{F0: "\n\t", F1: (v4 / (((v4 ^ v4)) | 1))}
                                v6 = append(v6, (fi12 ^ fi12))
                                _ = v34
                            }
                            _ = v32
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        _ = v27
                    } else {
                        fmt.Println(fi12, fi12, fi24)
                    }
                    _ = fi24
                }
            } else {
                fmt.Println(v11, v2, fi12)
            }
        }
        switch fi12 {
        case 4294967295:
            if (4294967295 == fi12) {
                v2 -= -32768
            } else {
                var v35 uint64 = 1
                v36 := GenS2{F0: ("\n\t" + ("hello" + "hello"))}
                v35 = v35
                _ = v35
                _ = v36
            }
        case 9223372036854775807:
            v37 := GenS4{F0: v4, F1: 845, F2: (v4 + v4)}
            _ = v37
        default:
            var v38 uint8 = 58
            _ = v38
        }
        _ = fi12
    }
    v11 += 0.1
    if (((!v5) || ((v3 & v3) >= (v3 >> (((v3 | v3)) & 63)))) && v5) {
        var v39 float32 = 1.5
        _ = v39
    }
    switch -823 {
    case 1:
        for fi40 := 0; fi40 < 14; fi40++ {
            var v41 uint16 = v4
            for fi42 := 0; fi42 < 16; fi42++ {
                v43 := GenS2{F0: ("" + "\n\t")}
                v4 += v4
                _ = fi42
                _ = v43
            }
            _ = fi40
            _ = v41
        }
    default:
        v5 = (!v5)
    }
    var v44 int16 = v2
    v1[(("" + "a") + "a")] = ""
    _ = len(v1)
    v45 := GenS4{F0: (^(v4)), F1: 170, F2: 612}
    for fi46 := 0; fi46 < 11; fi46++ {
        for fi47 := 0; fi47 < 16; fi47++ {
            v3 = v3
            _ = fi47
        }
        _ = fi46
    }
    v48 := GenS2{F0: "a"}
    var v49 uint64 = 1
    if (uint(v49) <= v3) {
        v50 := GenS2{F0: ((("foo bar" + "a") + "") + "\n\t")}
        if len(v6) > 0 {
            v6[0] = 862
        }
        for fi51 := 0; fi51 < 14; fi51++ {
            fmt.Println(v5)
            _ = fi51
        }
        _ = v50
    } else {
        var v52 uint16 = v4
        v53 := GenS6{F0: 180, F1: v5, F2: false}
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 4; fi54++ {
        v55 := v1["foo bar"]
        _ = fi54
        _ = v55
    }
    var v56 uint = v3
    var v57 int8 = -46
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v48
    _ = v49
    _ = v5
    _ = v56
    _ = v57
    _ = v6
}
