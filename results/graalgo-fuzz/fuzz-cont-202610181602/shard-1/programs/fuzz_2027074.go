package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    v4 := GenS1{F0: 1, F1: "foo bar", F2: 1}
    v5 := GenS1{F0: 0, F1: "foo bar", F2: 73}
    var v6 int8 = 0
    v7 := []int{-1, 127}
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint16 = 72
    v2 := make(map[string]string)
    v3 := make(map[int]int)
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 17; fi9++ {
            v3[fi8] = (fi9 / (((fi9 + -1)) | 1))
            v1 = (^(v1))
            _ = fi9
        }
        _ = fi8
    }
    var v10 int64 = -1
    v11 := []int{}
    v12 := GenS1{F0: genFn0(), F1: "", F2: 65535}
    switch "foo bar" {
    case "":
        fmt.Println(v10, v10)
    default:
        v13 := GenS1{F0: 1, F1: "", F2: 18446744073709551615}
        _ = v13
    }
    v14 := v2["\n\t"]
    if true {
        if false {
            var v15 int = 632
            _ = v15
        } else {
            var v16 uint16 = (v1 + 65535)
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int8 = 127
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        v19 := GenS1{F0: 65535, F1: v14, F2: 65535}
        _ = v19
    } else {
        var v20 uint = 18446744073709551615
        if false {
            var v21 int16 = 42
            switch "a" {
            case "":
                var v22 uint16 = v1
                _ = v22
            }
            v10 -= (v10 << ((53) & 63))
            _ = v21
        } else {
            fmt.Println(v1)
        }
        _ = v20
    }
    var v23 int16 = 56
    if false {
        fmt.Println(v14, v10)
        var v24 uint32 = 2147483647
        _ = v24
    } else {
        var v25 uint8 = 10
        for fi26 := 0; fi26 < 14; fi26++ {
            switch (fi26 + fi26) {
            case -858:
                var v27 uint = uint(v25)
                _ = v27
            case -2147483648:
                switch fi26 {
                case 970:
                    if ("hello" == "a") {
                        delete(v2, (("" + v14) + "a"))
                        var v28 uint = 0
                        _ = v28
                    }
                default:
                    v29 := GenS1{F0: 4294967295, F1: v14, F2: 65535}
                    _ = v29
                }
            case 128:
                if ((!true) && (!true)) {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        switch (v14 + ("a" + (((("foo bar" + v14) + v14) + "a") + "hello"))) {
                        case "foo bar":
                            v31 := GenS1{F0: 2147483647, F1: "\n\t", F2: 2}
                            _ = v31
                        case "a":
                            var v32 uint8 = v25
                            _ = v32
                        case "":
                            var v33 int64 = (v10 - v10)
                            switch 127 {
                            case -112:
                                var v34 uint64 = 42
                                _ = v34
                            case 42:
                                v35 := GenS2{F0: v25, F1: -85, F2: ((2 > 42) && true)}
                                _ = v35
                            case 32768:
                                fmt.Println(v23)
                            }
                            v36 := v2["foo bar"]
                            _ = v33
                            _ = v36
                        }
                        if (!true) {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            if (v1 != (^(v1))) {
                                v10 -= int64(fi30)
                                switch fi26 {
                                case -2147483648:
                                    v11 = append(v11, 32768)
                                }
                                v14 = v14
                            } else {
                                var v37 float64 = float64(((0.1 - ((1e10 * 1.5) / 1.5)) - ((1e10 * 1.5) * (0.0 - (float32((1.5 + 3.14159)) / 1.0)))))
                                switch (fi30 + -2147483648) {
                                case -514:
                                    fmt.Println(v14)
                                case 650:
                                    _ = len(v11)
                                    v14 = v14
                                case 273:
                                    if (3.14159 <= 0.0) {
                                        var v38 uint64 = 1
                                        var v39 int16 = v23
                                        if (false || (!(v14 <= v14))) {
                                            var v40 uint16 = v1
                                            _ = v40
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                default:
                                    _ = len(v11)
                                }
                                var v41 int64 = v10
                                _ = v37
                                _ = v41
                            }
                            if true {
                                v23 += v23
                            }
                            var v42 float32 = 0.0
                            v42 -= (0.1 * (((v42 / float32(((float64((float32((1e10 + ((float64((((-2.25 - 0.1) - float32(1e10)) / (((((float32((1.5 - ((((-2.25 * 0.1) / 3.14159) / 1.0) + 0.1))) / v42) / 3.14159) + -2.25) / v42) + ((3.14159 + (float32(1.5) * 0.1)) * v42)))) - (3.14159 / (-2.25 - 0.0))) + 3.14159))) * -2.25)) / 1.5) * float64(float32((0.0 * 1e10)))))) - v42) - v42))
                            _ = v42
                        }
                        _ = fi30
                    }
                }
            }
            _ = fi26
        }
        var v43 string = v14
        _ = v25
        _ = v43
    }
    var v44 int8 = 96
    if ((v23 - (-(v23))) < int16(v44)) {
        if true {
            var v45 uint16 = v1
            v11 = append(v11, 2147483647)
            _ = v45
        }
        fmt.Println(v1)
    } else {
        var v46 uint16 = (v1 & v1)
        if true {
            v14 = v14
        }
        v47 := GenS3{F0: 128}
        _ = v46
        _ = v47
    }
    if (0 <= 2) {
        v10 -= (v10 % ((v10) | 1))
        var v48 int8 = (v44 + (v44 / ((v44) | 1)))
        _ = v48
    } else {
        switch -9223372036854775808 {
        case 127:
            var v49 uint8 = 1
            _ = v49
        }
        var v50 int32 = 2147483647
        if true {
            for fi51 := 0; fi51 < 6; fi51++ {
                switch v14 {
                case "foo bar":
                    var v52 uint8 = 45
                    _ = v52
                case "\n\t":
                    if len(v11) > 0 {
                        v11[0] = fi51
                    }
                default:
                    for fi53 := 0; fi53 < 16; fi53++ {
                        v54 := GenS1{F0: 452, F1: "a", F2: 42}
                        _ = fi53
                        _ = v54
                    }
                }
                _ = fi51
            }
            fmt.Println(v44)
        } else {
            v55 := v2[v14]
            _ = v55
        }
        _ = v50
    }
    v23 -= v23
    if false {
        v56 := v3[1]
        _ = v56
    } else {
        _ = len(v3)
        for fi57 := 0; fi57 < 5; fi57++ {
            var v58 int16 = 872
            _ = fi57
            _ = v58
        }
    }
    var v59 int64 = (v10 % (((v10 * 723)) | 1))
    var v60 float32 = (3.14159 - -2.25)
    v61 := GenS2{F0: 2, F1: (v44 - v44), F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v44
    _ = v59
    _ = v60
    _ = v61
}
