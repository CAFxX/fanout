package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int; F1 int; F2 string }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint; F2 uint64 }

func main() {
    var v1 float32 = (float32(-2.25) / (1e10 - (3.14159 * 0.0)))
    v2 := make(map[string]int)
    var v3 int64 = -1
    v4 := []int{2147483647, -794}
    v3 += (v3 & (v3 / ((v3) | 1)))
    if true {
        v5 := GenS1{F0: 2, F1: 65535}
        if (true && (127 < -128)) {
            fmt.Println(v1)
            if ("" >= (("hello" + ("foo bar" + "")) + ((("a" + "") + "foo bar") + ("\n\t" + "a")))) {
                var v6 bool = ((!true) || true)
                var v7 float64 = 1.5
                _ = v6
                _ = v7
            }
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            v3 -= (v3 ^ (v3 ^ 9223372036854775807))
            _ = fi8
        }
        _ = v5
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: "foo bar", F1: -891}
            fmt.Println(v1)
            if true {
                var v11 float32 = ((v1 * (-2.25 + float32((1.5 * ((((0.1 + ((3.14159 + (1.5 / 1.0)) * (float64(float32((((1e10 / 1.5) * float64(((v1 - v1) - v1))) * 3.14159))) * ((-2.25 / 0.1) * 0.0)))) - 1.5) + 3.14159) / 3.14159))))) + -2.25)
                if (!true) {
                    var v12 uint64 = 18446744073709551615
                    if (!(!false)) {
                        v12 += (^(v12))
                        v11 -= (v1 / (v1 - v1))
                    } else {
                        var v13 uint32 = 631
                        v3 = (^(v3))
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 float64 = (1e10 + (float64(v11) * -2.25))
                            if (!false) {
                                if true {
                                    var v16 int8 = -50
                                    v12 -= (v12 - 4294967295)
                                    v17 := GenS3{F0: fi14, F1: (^(fi14)), F2: ("a" + "a")}
                                    var v18 uint = 1
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 uint32 = v13
                                    _ = v19
                                }
                                delete(v2, "a")
                                fmt.Println(v12)
                            }
                            if (-11 > 22) {
                                if (true || false) {
                                    var v20 uint8 = 69
                                    v11 -= (((1e10 - (((1.5 - ((((0.1 - 1e10) - 1.5) - 1.5) + 3.14159)) - ((3.14159 + v1) / (float32(0.1) / 1.0))) / float32(((((v15 / -2.25) / (1.5 / v15)) / v15) - ((v15 / float64(float32((v15 - (0.0 / 1.0))))) * 1e10))))) * v1) - v11)
                                    _ = v20
                                } else {
                                    v13 += v13
                                }
                                if false {
                                    v21 := GenS3{F0: (fi9 / ((fi9) | 1)), F1: (fi14 * 32767), F2: "a"}
                                    fmt.Println(v11, v3, fi14)
                                    var v22 int = fi9
                                    _ = v21
                                    _ = v22
                                } else {
                                    switch -913 {
                                    case -366:
                                        v23 := v2["hello"]
                                        _ = v23
                                    default:
                                        var v24 bool = ((423 < 255) || false)
                                        _ = v24
                                    }
                                }
                            } else {
                                delete(v2, "\n\t")
                            }
                            v25 := GenS3{F0: 85, F1: (^(fi14)), F2: "\n\t"}
                            _ = len(v4)
                            v26 := GenS1{F0: 42, F1: 2}
                            delete(v2, "a")
                            if ((255 > 42) && (true || (!(!true)))) {
                                v4 = append(v4, -2147483648)
                                v12 += (v12 + (^(v12)))
                                fmt.Println(v1, v3)
                            } else {
                                var v27 bool = true
                                _ = v27
                            }
                            _ = fi14
                            _ = v15
                            _ = v25
                            _ = v26
                        }
                        _ = v13
                    }
                    for fi28 := 0; fi28 < 18; fi28++ {
                        _ = len(v2)
                        _ = fi28
                    }
                    _ = v12
                } else {
                    switch fi9 {
                    case 879:
                        for fi29 := 0; fi29 < 2; fi29++ {
                            var v30 uint = 65535
                            v31 := GenS3{F0: fi9, F1: int(v3), F2: "\n\t"}
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    case -9223372036854775808:
                        var v32 uint32 = 1
                        _ = v32
                    default:
                        if false {
                            v33 := GenS2{F0: "a", F1: 0}
                            var v34 float64 = 0.1
                            _ = v33
                            _ = v34
                        }
                    }
                    v1 = (float32(0.0) * ((3.14159 + ((v11 - v11) + v1)) + (v1 / 0.1)))
                }
                _ = v11
            } else {
                var v35 float32 = ((-2.25 + 0.1) - (v1 * (float32((float64(1.5) / float64(v1))) + ((v1 - v1) / 3.14159))))
                _ = v35
            }
            _ = fi9
            _ = v10
        }
        var v36 int8 = 87
        for fi37 := 0; fi37 < 1; fi37++ {
            var v38 int8 = v36
            for fi39 := 0; fi39 < 17; fi39++ {
                switch ("\n\t" + "\n\t") {
                case "\n\t":
                    if (((1e10 * 1e10) >= (1e10 - (0.0 + (float64(v1) * float64(v1))))) || (!(!(1 == 2)))) {
                        var v40 int16 = -978
                        _ = v40
                    } else {
                        delete(v2, "")
                    }
                case "a":
                    v4 = append(v4, 207)
                default:
                    v41 := v2["foo bar"]
                    _ = v41
                }
                _ = fi39
            }
            _ = fi37
            _ = v38
        }
        _ = v36
    }
    switch 128 {
    case -2147483648:
        var v42 string = "a"
        _ = v42
    case -1:
        var v43 uint = 537
        _ = v43
    default:
        var v44 uint64 = 18446744073709551615
        _ = v44
    }
    v45 := GenS2{F0: ("foo bar" + "hello"), F1: -503}
    if false {
        v3 += (^(v3))
        v3 = (v3 ^ v3)
    } else {
        fmt.Println(v3, v3)
    }
    v46 := GenS2{F0: "foo bar", F1: 0}
    if false {
        for fi47 := 0; fi47 < 19; fi47++ {
            delete(v2, "hello")
            for fi48 := 0; fi48 < 16; fi48++ {
                v1 -= ((float32(1e10) * v1) + v1)
                switch fi48 {
                case 9223372036854775807:
                    v2[("\n\t" + "")] = fi47
                case -2147483648:
                    switch "a" {
                    case "a":
                        if ((float64(0.0) >= 1.5) || ((v1 + v1) != v1)) {
                            if true {
                                v49 := GenS4{F0: 2, F1: 181, F2: 1}
                                v3 -= (v3 ^ v3)
                                switch fi47 {
                                case 1:
                                    var v50 uint16 = 13
                                    _ = v50
                                }
                                _ = v49
                            }
                            fmt.Println(v3, fi47)
                        }
                    }
                case 251:
                    v2["\n\t"] = 494
                }
                v1 += 3.14159
                _ = fi48
            }
            _ = len(v4)
            _ = fi47
        }
        fmt.Println(v1)
        switch 4294967295 {
        case 0:
            var v51 float32 = 1e10
            _ = v51
        case 127:
            v1 -= (float32(float64((0.0 * 0.1))) - (-2.25 + v1))
        default:
            var v52 int32 = 258
            _ = v52
        }
    }
    v3 = -9223372036854775808
    v1 = v1
    fmt.Println(v3, v1)
    if true {
        switch "foo bar" {
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            var v53 int32 = 1
            _ = v53
        case "a":
            _ = len(v2)
        }
        v3 = v3
        var v54 uint32 = 65535
        _ = v54
    }
    v55 := GenS5{F0: 2, F1: 731, F2: 484}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v55
}
