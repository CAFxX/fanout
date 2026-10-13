package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint32; F2 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }
func genFn0() uint16 {
    var v6 int8 = -1
    v7 := []int{9223372036854775807, 1, -566, 1}
    v8 := GenS1{F0: 0, F1: 1}
    fmt.Println(v6)
    fmt.Println(v6)
    _ = v6
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1e10
    v3 := make(map[string]bool)
    v4 := []int{-2147483648, 32767, -169, 32767}
    var v5 float32 = 1.5
    if (-9223372036854775808 < 128) {
        v5 = 1e10
    }
    v2 = 0.0
    if true {
        var v9 uint64 = 0
        _ = v9
    }
    var v10 int = 2147483647
    fmt.Println(v2)
    var v11 float32 = 0.1
    v12 := GenS2{F0: 1000, F1: 0}
    for fi13 := 0; fi13 < 12; fi13++ {
        delete(v3, "foo bar")
        fmt.Println(v10, v5, v10)
        for fi14 := 0; fi14 < 17; fi14++ {
            v2 = v2
            var v15 float64 = v2
            if true {
                v16 := GenS3{F0: ""}
                _ = v16
            } else {
                var v17 uint8 = 247
                v18 := GenS2{F0: -1, F1: 65535}
                var v19 bool = false
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    if false {
        fmt.Println(v10, v2)
        switch "\n\t" {
        case "\n\t":
            if true {
                switch v10 {
                case 42:
                    v20 := GenS4{F0: "a", F1: 2, F2: 1}
                    _ = v20
                case 690:
                    v5 -= ((v11 * v11) / (v5 + v11))
                default:
                    v21 := GenS2{F0: -134, F1: 98}
                    _ = v21
                }
                if false {
                    fmt.Println(v2, v5)
                    var v22 int = v10
                    v11 = v5
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 8; fi23++ {
                        if (!true) {
                            _ = len(v3)
                        }
                        var v24 bool = (2147483647 == 1)
                        _ = fi23
                        _ = v24
                    }
                }
                if true {
                    var v25 bool = (2147483647 <= 42)
                    _ = v25
                }
            }
        case "foo bar":
            v26 := GenS3{F0: "hello"}
            _ = v26
        default:
            var v27 uint64 = 18446744073709551615
            _ = v27
        }
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi28 := 0; fi28 < 7; fi28++ {
                for fi29 := 0; fi29 < 11; fi29++ {
                    fmt.Println(fi29)
                    if false {
                        if false {
                            switch -9223372036854775808 {
                            case 344:
                                if ((0 <= -2147483648) && false) {
                                    var v30 bool = (!(false || (668 != 641)))
                                    _ = v30
                                }
                            default:
                                fmt.Println(v11, v5)
                            }
                        }
                    }
                    var v31 uint = 2
                    if (!((true || (1 < -1)) || (false || (42 > -9223372036854775808)))) {
                        v31 = (v31 + v31)
                        if len(v4) > 0 {
                            v4[0] = v10
                        }
                        var v32 uint32 = 106
                        _ = v32
                    } else {
                        if false {
                            if false {
                                if (!false) {
                                    v33 := v1[2147483647]
                                    if true {
                                        v10 += 359
                                        var v34 int = (-(fi28))
                                        v11 = 3.14159
                                        _ = v34
                                    } else {
                                        v33 = v33
                                        v35 := GenS5{F0: ("" + (("a" + ((v33 + (v33 + "")) + "a")) + v33))}
                                        v10 -= (fi28 ^ (fi29 + fi29))
                                        _ = v35
                                    }
                                    _ = v33
                                }
                            } else {
                                var v36 uint64 = 1
                                if true {
                                    v5 = ((0.1 + v11) + float32(0.1))
                                    var v37 int64 = 977
                                    var v38 int16 = -1
                                    _ = v37
                                    _ = v38
                                }
                                fmt.Println(v36)
                                _ = v36
                            }
                        }
                        if (0 < 2) {
                            var v39 uint32 = 863
                            _ = v39
                        } else {
                            switch (("" + (("" + "") + "hello")) + "a") {
                            case "foo bar":
                                var v40 uint64 = 756
                                _ = v40
                            }
                            if ((134 < 114) || (!(880 < -500))) {
                                if true {
                                    v1[353] = "a"
                                    v41 := v1[4294967295]
                                    _ = v41
                                }
                                var v42 uint8 = 233
                                _ = v42
                            }
                            delete(v1, 65535)
                        }
                        var v43 int64 = -227
                        v5 = (v11 / ((float32(((float64(3.14159) + v2) / float64(v11))) + float32(1.5)) - (((float32(v2) + float32(float64(0.1))) / v5) * (v5 * v5))))
                        _ = v43
                    }
                    v2 -= (v2 + v2)
                    _ = fi29
                    _ = v31
                }
                _ = fi28
            }
        }
    }
    var v44 uint = 2
    switch (v10 + int(v44)) {
    case -9223372036854775808:
        fmt.Println(v44)
    case 143:
        if true {
            v45 := GenS6{F0: 67}
            v3["foo bar"] = true
            _ = v45
        }
    }
    var v46 uint = v44
    for fi47 := 0; fi47 < 10; fi47++ {
        var v48 float64 = float64(v11)
        v4 = append(v4, -517)
        fmt.Println(v10)
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 13; fi49++ {
        v44 -= v46
        v50 := GenS2{F0: 0, F1: 255}
        var v51 uint64 = 757
        _ = fi49
        _ = v50
        _ = v51
    }
    if true {
        if true {
            switch (v10 >> (((^(v46))) & 63)) {
            case 9223372036854775807:
                for fi52 := 0; fi52 < 16; fi52++ {
                    if false {
                        v53 := GenS4{F0: "", F1: 2147483648, F2: 65535}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v54 float32 = v11
                        _ = v53
                        _ = v54
                    }
                    _ = fi52
                }
            case -434:
                if len(v4) > 0 {
                    v4[0] = (v10 - 0)
                }
            case -840:
                fmt.Println(v10, v2, v46)
            default:
                v46 = (v44 * v46)
            }
        }
        for fi55 := 0; fi55 < 14; fi55++ {
            v56 := GenS4{F0: (("\n\t" + ((("hello" + "a") + "foo bar") + "foo bar")) + "foo bar"), F1: 198, F2: 32767}
            v11 -= v11
            _ = fi55
            _ = v56
        }
        var v57 uint = (v46 << ((v46) & 63))
        _ = v57
    }
    for fi58 := 0; fi58 < 10; fi58++ {
        for fi59 := 0; fi59 < 14; fi59++ {
            v5 += (v11 / 3.14159)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v44)
            _ = fi59
        }
        v60 := v1[1]
        v10 = (v10 ^ (fi58 ^ -594))
        _ = fi58
        _ = v60
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v5
}
