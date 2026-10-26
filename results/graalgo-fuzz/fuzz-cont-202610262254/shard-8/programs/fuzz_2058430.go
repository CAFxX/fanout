package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 string; F2 uint }
type GenS3 struct { F0 int; F1 int32; F2 uint64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint = 0
    v2 := []int{0}
    var v3 bool = (9223372036854775807 > 9223372036854775807)
    v4 := make(map[int]int)
    var v5 uint = 994
    v5 = (v5 % (((v1 / ((v5) | 1))) | 1))
    v3 = (v3 && (65535 > 140))
    switch 2147483647 {
    case 415:
        var v6 bool = (2147483647 >= -178)
        _ = v6
    case -9223372036854775808:
        fmt.Println(v5)
    }
    fmt.Println(v3, v3, v1)
    if len(v2) > 0 {
        v2[0] = 128
    }
    if len(v2) > 0 {
        v2[0] = 32768
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        switch (("foo bar" + ((("" + "") + "") + "foo bar")) + ("foo bar" + "a")) {
        case "hello":
            if (v3 || v3) {
                v8 := GenS1{F0: 42}
                _ = v8
            } else {
                v9 := GenS2{F0: -117, F1: "", F2: 18446744073709551615}
                v10 := GenS2{F0: 1, F1: "a", F2: (v5 + (v1 ^ v1))}
                _ = len(v4)
                _ = v10
                _ = v9
            }
        case "foo bar":
            if len(v2) > 0 {
                v2[0] = (fi7 % ((128) | 1))
            }
        }
        fmt.Println(v5)
        if (4294967295 <= 65535) {
            v5 = (v1 ^ v1)
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS2{F0: -1, F1: "foo bar", F2: v1}
                var v13 float64 = (0.0 + 1e10)
                fmt.Println(fi11)
                _ = len(v2)
                var v14 int32 = 1
                var v15 int32 = v14
                if (398 != 255) {
                    v16 := GenS3{F0: (fi11 ^ (fi11 - fi11)), F1: (v14 ^ (^(v15))), F2: 909}
                    v17 := GenS4{F0: 668, F1: 0}
                    _ = v16
                    _ = v17
                }
                switch 2147483647 {
                case 721:
                    v13 = (0.1 + ((v13 - 3.14159) - 0.1))
                default:
                    fmt.Println(fi7)
                }
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
                _ = v15
            }
        } else {
            var v18 string = ""
            if ((^(v1)) == (v1 << ((4294967295) & 63))) {
                for fi19 := 0; fi19 < 14; fi19++ {
                    delete(v4, (-(fi7)))
                    if false {
                        var v20 int = fi7
                        if (228 <= 1) {
                            var v21 float32 = float32(((0.1 * (0.1 + 3.14159)) - 3.14159))
                            _ = v21
                        }
                        v18 = v18
                        var v22 int16 = -32768
                        v20 += (v20 ^ (fi7 - v20))
                        _ = v20
                        _ = v22
                    } else {
                        v3 = (2147483647 <= 2147483647)
                    }
                    _ = fi19
                }
                for fi23 := 0; fi23 < 11; fi23++ {
                    var v24 string = "\n\t"
                    var v25 string = ("foo bar" + (((("" + v18) + "a") + v24) + ("" + ("a" + ((("" + ("a" + ("\n\t" + v18))) + v18) + "hello")))))
                    var v26 float64 = -2.25
                    v27 := v4[(fi7 >> ((v5) & 63))]
                    var v28 float32 = 1e10
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                if v3 {
                    v1 += (v5 + 18446744073709551615)
                    var v29 float64 = 1.5
                    _ = v29
                }
            } else {
                var v30 uint = v5
                fmt.Println(v1, v5, v3)
                _ = v30
            }
            var v31 bool = v3
            _ = v18
            _ = v31
        }
        _ = fi7
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 uint64 = 415
        fmt.Println(v33, v3, v1)
        _ = fi32
        _ = v33
    }
    v1 -= v1
    switch ("a" + "") {
    case "hello":
        switch 32767 {
        case 308:
            var v34 string = ("foo bar" + ((("foo bar" + ((("" + "\n\t") + "") + "hello")) + ("a" + ("\n\t" + (("\n\t" + "") + "\n\t")))) + ""))
            _ = v34
        case 2147483647:
            v35 := v4[2147483647]
            _ = v35
        case 32767:
            v3 = v3
        default:
            if v3 {
                v5 += v5
                if (968 > 42) {
                    var v36 float64 = -2.25
                    v36 += (v36 / (3.14159 - ((1e10 / (v36 * ((0.1 * (v36 - (1e10 / float64((3.14159 - 0.0))))) / (((float64((0.1 / -2.25)) - (v36 / v36)) + 0.1) * (v36 + -2.25))))) / (((float64(((3.14159 * 0.0) * float32(float64((0.0 * 0.0))))) * -2.25) + (v36 / v36)) / (1e10 + -2.25)))))
                    if (!(v3 || (!((^(v1)) > v5)))) {
                        v37 := GenS2{F0: -50, F1: "a", F2: v5}
                        v38 := GenS2{F0: 100, F1: ("foo bar" + ""), F2: v5}
                        _ = v37
                        _ = v38
                    }
                    _ = v36
                }
            } else {
                var v39 int32 = 494
                if v3 {
                    if ((9223372036854775807 != -9223372036854775808) || true) {
                        switch 31 {
                        case -782:
                            v3 = v3
                        case 42:
                            for fi40 := 0; fi40 < 7; fi40++ {
                                v41 := v4[fi40]
                                var v42 uint32 = 1
                                switch "foo bar" {
                                case "foo bar":
                                    var v43 float64 = 1.5
                                    _ = v43
                                case "":
                                    fmt.Println(fi40)
                                }
                                _ = fi40
                                _ = v41
                                _ = v42
                            }
                        case -70:
                            v39 = (v39 | (v39 & 0))
                        }
                        v5 = (v5 * 0)
                    }
                    var v44 bool = true
                    v44 = (42 < 0)
                    _ = v44
                }
                for fi45 := 0; fi45 < 2; fi45++ {
                    v46 := GenS5{F0: 42}
                    v1 -= v1
                    _ = fi45
                    _ = v46
                }
                _ = v39
            }
        }
    case "a":
        v3 = v3
    }
    for fi47 := 0; fi47 < 18; fi47++ {
        for fi48 := 0; fi48 < 6; fi48++ {
            v5 -= (v1 & v1)
            _ = fi48
        }
        for fi49 := 0; fi49 < 6; fi49++ {
            v3 = false
            var v50 int = fi49
            if false {
                var v51 uint8 = 28
                _ = v51
            }
            _ = fi49
            _ = v50
        }
        _ = fi47
    }
    v52 := GenS4{F0: 32767, F1: 42}
    if false {
        v53 := GenS6{F0: (v1 % ((v1) | 1)), F1: -486}
        var v54 int16 = -32768
        _ = v53
        _ = v54
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
