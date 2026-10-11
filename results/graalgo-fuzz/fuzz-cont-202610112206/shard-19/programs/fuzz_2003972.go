package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int8, p1 string, p2 string) uint8 {
    p2 = p1
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 255
}
func genFn1(p0 int8, p1 int, p2 uint32) int32 {
    v4 := []int{(^(p1)), p1, p1, p1}
    if (!(!((true && (65535 > 240)) || (("\n\t" == ("" + "")) || (-284 > 261))))) {
        fmt.Println(p0, p1)
        for fi5 := 0; fi5 < 5; fi5++ {
            fmt.Println(fi5, p2, fi5)
            _ = fi5
        }
        p1 = -293
    }
    p2 += 271
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 2147483647
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 0.0
    var v3 uint16 = 2
    if (!false) {
        for fi6 := 0; fi6 < 17; fi6++ {
            v7 := GenS1{F0: 255, F1: 726, F2: "foo bar"}
            v8 := GenS1{F0: 127, F1: 1, F2: "hello"}
            if true {
                var v9 uint = 18446744073709551615
                if true {
                    v3 -= (v3 >> ((v9) & 63))
                    for fi10 := 0; fi10 < 3; fi10++ {
                        if (true && (32767 > (fi10 ^ (fi6 % (((fi6 | fi6)) | 1))))) {
                            switch (fi6 / ((fi10) | 1)) {
                            case 972:
                                v11 := GenS1{F0: 255, F1: 722, F2: ""}
                                _ = v11
                            case 1:
                                if ("" > ("a" + "")) {
                                    var v12 int8 = 0
                                    if false {
                                        v9 += (v9 | (v9 % (((v9 << ((v9) & 63))) | 1)))
                                        var v13 bool = false
                                        _ = v13
                                    } else {
                                        var v14 int8 = v12
                                        var v15 uint32 = 65535
                                        _ = v14
                                        _ = v15
                                    }
                                    v16 := GenS1{F0: 2, F1: 0, F2: ""}
                                    _ = v12
                                    _ = v16
                                }
                                var v17 int16 = 1000
                                _ = v17
                            default:
                                var v18 uint = v9
                                _ = v18
                            }
                        }
                        fmt.Println(fi10, v3)
                        _ = fi10
                    }
                    switch fi6 {
                    case 346:
                        var v19 uint = (v9 ^ v9)
                        _ = v19
                    case 846:
                        v20 := []int{fi6, int(v3)}
                        _ = v20
                    case 127:
                        for fi21 := 0; fi21 < 13; fi21++ {
                            v9 += (^(v9))
                            _ = fi21
                        }
                    default:
                        switch (-(fi6)) {
                        case -463:
                            switch (fi6 / (((fi6 << (((^(v9))) & 63))) | 1)) {
                            case 4294967295:
                                v2 = (3.14159 / ((-2.25 + v2) + (1.5 + 3.14159)))
                            case 1:
                                v22 := GenS1{F0: 0, F1: 99, F2: ""}
                                _ = v22
                            case 834:
                                var v23 bool = (1 >= -37)
                                _ = v23
                            }
                        case 4294967295:
                            v24 := GenS1{F0: 1, F1: 1, F2: ""}
                            _ = v24
                        default:
                            var v25 uint = (v9 + (v9 / ((255) | 1)))
                            _ = v25
                        }
                    }
                } else {
                    if true {
                        fmt.Println(v3, v3, fi6)
                        v3 = 1
                    } else {
                        var v26 uint16 = uint16(v9)
                        _ = v26
                    }
                }
                _ = v9
            } else {
                for fi27 := 0; fi27 < 9; fi27++ {
                    if true {
                        var v28 string = "\n\t"
                        var v29 int8 = 100
                        _ = v28
                        _ = v29
                    }
                    _ = fi27
                }
                switch "\n\t" {
                case "a":
                    var v30 uint64 = 125
                    _ = v30
                case "foo bar":
                    for fi31 := 0; fi31 < 20; fi31++ {
                        if false {
                            v32 := GenS2{F0: -32768, F1: "foo bar"}
                            _ = v32
                        } else {
                            var v33 uint = 0
                            var v34 float32 = (3.14159 / (((1.5 + (float32(0.1) - -2.25)) + 1e10) * (((float32(v2) * ((1e10 - 3.14159) * float32(((v2 / float64((1.5 * 3.14159))) * v2)))) * float32(v2)) * 3.14159)))
                            _ = v33
                            _ = v34
                        }
                        v35 := v1[(("" + "foo bar") + "")]
                        _ = fi31
                        _ = v35
                    }
                default:
                    fmt.Println(fi6, fi6, v2)
                }
                if (true || (v3 >= 556)) {
                    v36 := []int{(fi6 | (-(fi6))), 32768}
                    if (true && (int(v3) != 127)) {
                        for fi37 := 0; fi37 < 8; fi37++ {
                            v38 := GenS3{F0: 2}
                            var v39 bool = false
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        v40 := GenS1{F0: genFn0(58, "foo bar", "foo bar"), F1: 42, F2: ((("hello" + "foo bar") + "a") + ("hello" + "\n\t"))}
                        var v41 uint32 = 2147483647
                        _ = v40
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 4; fi42++ {
                            _ = len(v1)
                            v1["foo bar"] = (fi6 << ((410) & 63))
                            _ = fi42
                        }
                        if len(v36) > 0 {
                            v36[0] = 4294967295
                        }
                        switch ("foo bar" + (("\n\t" + "foo bar") + ("\n\t" + ""))) {
                        case "a":
                            var v43 int32 = 2147483647
                            _ = v43
                        case "hello":
                            var v44 int64 = 231
                            _ = v44
                        default:
                            switch fi6 {
                            case -9223372036854775808:
                                if (0.0 < 3.14159) {
                                    if true {
                                        v3 = v3
                                    }
                                    if true {
                                        v36 = append(v36, fi6)
                                        v3 = v3
                                        _ = len(v36)
                                    }
                                } else {
                                    v2 -= float64(((float32(v2) * 0.0) / -2.25))
                                    switch "hello" {
                                    case "hello":
                                        v3 += (^(v3))
                                    case "foo bar":
                                        v3 += 594
                                    case "\n\t":
                                        var v45 uint16 = (v3 / ((v3) | 1))
                                        _ = v45
                                    default:
                                        fmt.Println(fi6, v3)
                                    }
                                    var v46 int16 = 1000
                                    _ = v46
                                }
                            case 929:
                                var v47 uint64 = 2
                                _ = v47
                            case 2147483647:
                                v2 = 1.5
                            default:
                                fmt.Println(v3, fi6, v2)
                            }
                        }
                    }
                    var v48 uint16 = (v3 << ((18446744073709551615) & 63))
                    _ = v36
                    _ = v48
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    fmt.Println(v3, v3)
    v3 += v3
    if (1 >= -1) {
        var v49 int16 = 0
        var v50 int32 = -588
        _ = v49
        _ = v50
    }
    switch "" {
    case "a":
        if true {
            if false {
                if (678 < 9223372036854775807) {
                    v51 := GenS1{F0: genFn0(113, ("a" + "hello"), ("a" + "hello")), F1: 817, F2: ""}
                    _ = v51
                } else {
                    var v52 uint = 65535
                    _ = v52
                }
                fmt.Println(v2, v2)
                fmt.Println(v3)
            } else {
                v53 := []int{}
                delete(v1, ("\n\t" + "foo bar"))
                _ = v53
            }
            fmt.Println(v2, v2)
        } else {
            if true {
                var v54 float32 = (((float32(v2) - float32((((float64(1.5) * v2) - float64(1e10)) / ((v2 * v2) - (v2 - v2))))) + 1e10) * (0.0 / -2.25))
                _ = v54
            }
            var v55 int16 = -462
            if false {
                var v56 bool = (-1 < 9223372036854775807)
                _ = v56
            }
            _ = v55
        }
    }
    var v57 int = 42
    v58 := GenS1{F0: 1, F1: 1, F2: ("a" + "foo bar")}
    v1[("" + ("" + ("\n\t" + "foo bar")))] = v57
    v59 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v59:", len(v59), v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v57
    _ = v58
    _ = v59
}
