package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int; F2 int }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 string = "a"
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi5, v6, fi5)
            var v8 string = v6
            fmt.Println(v6)
            _ = fi7
            _ = v8
        }
        v9 := []int{(fi5 >> ((0) & 63)), (fi5 / (((fi5 % ((fi5) | 1))) | 1)), fi5, -476}
        _ = fi5
        _ = v6
        _ = v9
    }
    switch 326 {
    case 1:
        switch -869 {
        case 32767:
            v10 := GenS1{F0: 18446744073709551615, F1: "a", F2: -580}
            _ = v10
        default:
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi11)
                v12 := GenS1{F0: 65535, F1: (((("\n\t" + "\n\t") + "\n\t") + "a") + ("foo bar" + ("" + "\n\t"))), F2: 1}
                _ = fi11
                _ = v12
            }
        }
    default:
        var v13 int16 = 0
        _ = v13
    }
    if (false && false) {
        v14 := GenS1{F0: 42, F1: "hello", F2: -259}
        if (65535 > 255) {
            v15 := []int{391, -626, 127}
            if len(v15) > 0 {
                v15[0] = -2147483648
            }
            var v16 uint64 = 0
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 65535, F1: "a", F2: 1}
            for fi18 := 0; fi18 < 15; fi18++ {
                var v19 float64 = (3.14159 + 3.14159)
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        v20 := []int{0, 0, 32767, 128}
        _ = v14
        _ = v20
    }
    return 255
}

func main() {
    v1 := []int{37, 32767}
    v2 := make(map[string]bool)
    var v3 uint = 0
    v4 := make(map[uint32]int)
    var v21 uint8 = 244
    if (!(401 > v3)) {
        var v22 uint8 = (v21 >> ((v3) & 63))
        _ = v22
    } else {
        v1 = append(v1, -1)
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 int16 = 42
        var v25 int16 = -1
        for fi26 := 0; fi26 < 4; fi26++ {
            delete(v2, "hello")
            v24 = v24
            switch (fi26 / ((fi23) | 1)) {
            case 32768:
                if len(v1) > 0 {
                    v1[0] = fi23
                }
                v21 = (v21 | 1)
            case 0:
                var v27 uint = v3
                _ = v27
            default:
                v28 := GenS2{F0: v21, F1: "\n\t"}
                _ = v28
            }
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    switch 32767 {
    case -263:
        var v29 uint8 = genFn0()
        _ = v29
    }
    switch 1 {
    case 32768:
        var v30 uint16 = 255
        _ = v30
    case 32767:
        switch -9223372036854775808 {
        case 42:
            fmt.Println(v3, v21)
        }
    case -9223372036854775808:
        fmt.Println(v3, v21)
    default:
        var v31 float32 = (float32(float64(0.0)) * (3.14159 / 1e10))
        _ = v31
    }
    v32 := GenS3{F0: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v3)
    for fi33 := 0; fi33 < 15; fi33++ {
        for fi34 := 0; fi34 < 12; fi34++ {
            switch 134 {
            case 290:
                var v35 int8 = -1
                _ = v35
            default:
                v21 -= v21
            }
            _ = fi34
        }
        _ = len(v1)
        _ = fi33
    }
    var v36 float64 = -2.25
    v37 := GenS4{F0: 749, F1: 1}
    v21 += v21
    if false {
        switch 1 {
        case 2147483647:
            fmt.Println(v36, v3)
        case -1:
            if false {
                var v38 uint = 404
                var v39 string = ""
                switch v39 {
                case "":
                    v3 -= 65535
                case "a":
                    if (!false) {
                        v40 := GenS3{F0: 42}
                        v38 = (v38 ^ (v38 * v38))
                        fmt.Println(v36)
                        _ = v40
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if ((351 == 42) || (!(!false))) {
                            if true {
                                if false {
                                    for fi41 := 0; fi41 < 16; fi41++ {
                                        fmt.Println(v39)
                                        if ((-2.25 / (3.14159 * float32(-2.25))) > 0.1) {
                                            v42 := GenS5{F0: 42}
                                            var v43 string = "\n\t"
                                            switch fi41 {
                                            case 4294967295:
                                                var v44 string = "foo bar"
                                                _ = v44
                                            case -804:
                                                for fi45 := 0; fi45 < 2; fi45++ {
                                                    var v46 uint8 = v21
                                                    _ = fi45
                                                    _ = v46
                                                }
                                            case -839:
                                                var v47 uint8 = uint8(v3)
                                                _ = v47
                                            }
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v48 := v4[2147483648]
                                            v38 = 316
                                            _ = v48
                                        }
                                        v21 -= v21
                                        _ = fi41
                                    }
                                    v38 -= (v38 | 18446744073709551615)
                                } else {
                                    fmt.Println(v21, v39)
                                    v21 += (v21 % (((v21 - v21)) | 1))
                                    var v49 int = -2147483648
                                    _ = v49
                                }
                                fmt.Println(v39)
                                v50 := GenS5{F0: -68}
                                _ = v50
                            }
                            v39 = (v39 + (v39 + "a"))
                            for fi51 := 0; fi51 < 1; fi51++ {
                                var v52 uint = (v3 | (^(v38)))
                                _ = fi51
                                _ = v52
                            }
                        }
                    }
                case "hello":
                    v21 -= (v21 | (v21 - genFn0()))
                }
                _ = v38
                _ = v39
            }
        case 1:
            v4[42] = 4294967295
        default:
            if (!false) {
                v53 := GenS4{F0: 2, F1: 2147483647}
                _ = v53
            }
        }
        for fi54 := 0; fi54 < 6; fi54++ {
            v21 = (^(v21))
            if true {
                var v55 uint32 = 359
                _ = v55
            } else {
                v56 := GenS2{F0: (v21 / ((v21) | 1)), F1: "foo bar"}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v56
            }
            v21 = 207
            _ = fi54
        }
        if (2147483647 != 2147483647) {
            if len(v1) > 0 {
                v1[0] = 0
            }
            for fi57 := 0; fi57 < 17; fi57++ {
                v21 -= v21
                var v58 float64 = (0.1 - (v36 / (v36 - (1.5 - v36))))
                _ = fi57
                _ = v58
            }
        }
    }
    delete(v4, 2)
    v21 = uint8(v3)
    if (-67 < -2147483648) {
        v59 := GenS6{F0: v21, F1: -951, F2: -1}
        _ = v59
    } else {
        var v60 int = int(v21)
        _ = v60
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
}
