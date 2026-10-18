package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint64; F1 int8; F2 int64 }
type GenS5 struct { F0 int; F1 int; F2 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 int64; F1 int }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]int)
    var v3 int64 = 9223372036854775807
    v4 := []uint32{65535, 0, 619}
    var v5 int32 = 2147483647
    v6 := make(map[string]int)
    fmt.Println(v3)
    if (!(!("a" >= "foo bar"))) {
        if ((-2.25 - 1.5) <= (1e10 + (float32((0.0 - 1.5)) / 1.0))) {
            v7 := GenS1{F0: 94, F1: v5}
            var v8 bool = (1 <= 127)
            _ = v7
            _ = v8
        } else {
            if len(v4) > 0 {
                v4[0] = 42
            }
            v1 -= (v1 + 923)
        }
        var v9 float32 = 1e10
        _ = v9
    }
    var v10 float32 = ((0.0 - (((0.1 - 3.14159) / 1.0) + 3.14159)) * ((3.14159 - (1e10 / 3.14159)) - -2.25))
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 bool = true
        _ = fi11
        _ = v12
    }
    if true {
        var v13 float32 = (((3.14159 + float32((float64((v10 * (((v10 / (v10 * v10)) - v10) / 1e10))) - ((1.5 * (0.0 - 1e10)) + 1e10)))) - v10) / (v10 + 1e10))
        v14 := GenS2{F0: -32768, F1: "foo bar"}
        for fi15 := 0; fi15 < 16; fi15++ {
            v3 = v3
            _ = fi15
        }
        _ = v13
        _ = v14
    } else {
        var v16 int64 = (v3 | (v3 % (((-(v3))) | 1)))
        var v17 uint64 = 4294967295
        if true {
            v18 := GenS3{F0: 884, F1: -979, F2: v17}
            var v19 bool = false
            v1 = v1
            _ = v18
            _ = v19
        } else {
            var v20 string = "foo bar"
            _ = v20
        }
        _ = v16
        _ = v17
    }
    var v21 bool = (42 < v5)
    v1 -= 42
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 float32 = ((v10 * float32(0.0)) + v10)
        var v24 int8 = 42
        if ((!(v23 > v10)) || v21) {
            if (v21 && (0 <= (^(v5)))) {
                var v25 float32 = float32(0.1)
                var v26 uint32 = 238
                v3 = (v3 + 1)
                _ = v25
                _ = v26
            } else {
                fmt.Println(fi22, v10)
                switch "" {
                case "\n\t":
                    for fi27 := 0; fi27 < 4; fi27++ {
                        var v28 int16 = 1000
                        v29 := GenS2{F0: (v28 * 1), F1: "\n\t"}
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            if false {
                                var v30 uint64 = 912
                                v31 := GenS1{F0: 100, F1: v5}
                                if true {
                                    if v21 {
                                        v32 := GenS2{F0: (v28 + (v28 & v28)), F1: "\n\t"}
                                        switch fi27 {
                                        case -71:
                                            var v33 string = "a"
                                            v6[(v33 + v33)] = 0
                                            v5 = (-(v5))
                                            if v21 {
                                                fmt.Println(fi27, v23)
                                                var v34 float32 = v23
                                                v2[2] = fi27
                                                _ = v34
                                            }
                                            var v35 int16 = (v28 * (v28 ^ (v28 / ((v28) | 1))))
                                            _ = v33
                                            _ = v35
                                        case -1:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v36 int8 = -1
                                            fmt.Println(v24, v24)
                                            _ = v36
                                        case 42:
                                            v37 := GenS4{F0: (v30 >> ((v1) & 63)), F1: (v24 * 1), F2: -783}
                                            _ = v37
                                        default:
                                            v38 := GenS3{F0: v1, F1: (v28 + 987), F2: 925}
                                            _ = v38
                                        }
                                        _ = len(v6)
                                        v39 := GenS5{F0: fi27, F1: fi27, F2: 32767}
                                        _ = v32
                                        _ = v39
                                    }
                                    switch -860 {
                                    case 128:
                                        var v40 uint = (^(v1))
                                        _ = v40
                                    case -53:
                                        v41 := GenS6{F0: 100}
                                        _ = v41
                                    default:
                                        v42 := GenS7{F0: "hello", F1: "\n\t"}
                                        _ = v42
                                    }
                                }
                                _ = v30
                                _ = v31
                            } else {
                                v24 = (-(v24))
                            }
                        case "foo bar":
                            switch fi22 {
                            case -9223372036854775808:
                                _ = len(v6)
                            case -437:
                                v2[65535] = fi22
                            default:
                                v5 = v5
                            }
                        case "hello":
                            if v21 {
                                if (v21 || v21) {
                                    _ = len(v2)
                                    var v43 int64 = (v3 / ((9223372036854775807) | 1))
                                    v28 = 0
                                    v6[("hello" + ("a" + (("\n\t" + "foo bar") + "foo bar")))] = 4294967295
                                    _ = v43
                                }
                                switch fi27 {
                                case 771:
                                    var v44 int64 = (^(v3))
                                    _ = v44
                                default:
                                    v45 := GenS7{F0: "foo bar", F1: ((("hello" + ("\n\t" + "a")) + "hello") + ((("" + ("foo bar" + "")) + ("a" + "\n\t")) + "\n\t"))}
                                    v46 := GenS8{F0: (v3 & (v3 + (v3 ^ v3))), F1: 128}
                                    var v47 float32 = ((float32(1.5) + v10) * 0.1)
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                            }
                        default:
                            v24 = 100
                        }
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    switch 128 {
                    case 4294967295:
                        v5 = v5
                    case -452:
                        var v48 uint8 = 42
                        _ = v48
                    }
                }
            }
            if (v21 && true) {
                for fi49 := 0; fi49 < 12; fi49++ {
                    if (v24 >= v24) {
                        var v50 float32 = v10
                        var v51 int64 = (v3 - (^(v3)))
                        _ = v50
                        _ = v51
                    }
                    v3 += v3
                    v24 += v24
                    _ = fi49
                }
            } else {
                var v52 uint16 = 255
                v23 += float32(1e10)
                _ = len(v4)
                _ = v52
            }
            var v53 float32 = 0.1
            _ = v53
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v54 uint64 = 18446744073709551615
    v10 -= ((v10 * (0.0 * -2.25)) / 0.1)
    fmt.Println(v10, v3, v10)
    for fi55 := 0; fi55 < 12; fi55++ {
        var v56 string = ""
        if v21 {
            if (0 == v54) {
                v56 = (("" + ("" + v56)) + "\n\t")
                for fi57 := 0; fi57 < 6; fi57++ {
                    v3 += (v3 | v3)
                    var v58 uint32 = 4294967295
                    fmt.Println(fi55, v1, v1)
                    _ = fi57
                    _ = v58
                }
            } else {
                for fi59 := 0; fi59 < 14; fi59++ {
                    fmt.Println(v21, fi55)
                    v10 = (v10 / float32(((3.14159 + 1.5) + 1e10)))
                    fmt.Println(v1, v54)
                    _ = fi59
                }
            }
            v3 = v3
            for fi60 := 0; fi60 < 16; fi60++ {
                var v61 float32 = (1.5 / (float32(0.1) * v10))
                _ = len(v2)
                _ = fi60
                _ = v61
            }
        }
        var v62 uint32 = 1
        _ = fi55
        _ = v56
        _ = v62
    }
    var v63 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v63
}
