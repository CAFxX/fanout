package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 0.0
    v3 := []int{67, -1, 0, -448}
    var v4 int8 = -128
    var v5 string = ""
    var v6 int = 1
    _ = len(v3)
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            v8 := GenS1{F0: 32767, F1: 65535, F2: 2}
            var v9 float32 = -2.25
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 float32 = ((v9 * v9) * ((((((v9 * ((1e10 / (((3.14159 - (-2.25 - (float32(3.14159) - (float32((((float64(v9) * (float64(v9) - v2)) - 1e10) + -2.25)) / 0.1)))) / (v9 + 3.14159)) * float32(v2))) - 1e10)) / ((1e10 * v9) - (v9 - ((float32(v2) - (v9 * v9)) * float32((v2 * 0.1)))))) * (-2.25 + 1.5)) * 0.0) - 3.14159) / float32(1.5)))
                var v12 uint = 113
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v8
            _ = v9
        } else {
            if len(v3) > 0 {
                v3[0] = (fi7 << ((1) & 63))
            }
            v13 := GenS1{F0: 1000, F1: uint16(v1), F2: 128}
            var v14 int64 = 90
            _ = v13
            _ = v14
        }
        _ = fi7
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v5 = v5
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi15
    }
    switch (v5 + "\n\t") {
    case "hello":
        var v16 uint = 1
        _ = v16
    case "foo bar":
        if false {
            _ = len(v3)
            v3 = append(v3, (v6 ^ 9223372036854775807))
            v2 -= 1.5
        } else {
            if true {
                switch v6 {
                case 127:
                    if true {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if false {
                                v18 := GenS1{F0: 482, F1: 42, F2: 1}
                                _ = len(v3)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v18
                            } else {
                                var v19 int8 = -128
                                var v20 int32 = 407
                                switch v6 {
                                case 9223372036854775807:
                                    v21 := GenS2{F0: 1, F1: v20, F2: 429}
                                    _ = v21
                                case 32767:
                                    var v22 int = fi17
                                    _ = v22
                                case 631:
                                    if (1e10 <= (-2.25 * ((0.1 + 1.5) * 1.5))) {
                                        var v23 bool = (1 >= 128)
                                        var v24 string = v5
                                        switch fi17 {
                                        case -238:
                                            for fi25 := 0; fi25 < 4; fi25++ {
                                                if v23 {
                                                    if len(v3) > 0 {
                                                        v3[0] = 806
                                                    }
                                                    v2 = (-2.25 / (float64((1e10 - 3.14159)) - (((1e10 - v2) * ((0.0 * v2) - 0.0)) - v2)))
                                                } else {
                                                    var v26 bool = (v6 < fi25)
                                                    _ = v26
                                                }
                                                _ = len(v3)
                                                _ = len(v3)
                                                _ = fi25
                                            }
                                        default:
                                            var v27 int64 = int64(v1)
                                            _ = v27
                                        }
                                        _ = v23
                                        _ = v24
                                    } else {
                                        switch int(v1) {
                                        case 650:
                                            v5 = ((v5 + "a") + (v5 + "\n\t"))
                                        default:
                                            fmt.Println(v2, v20)
                                        }
                                    }
                                default:
                                    var v28 uint16 = 42
                                    _ = v28
                                }
                                _ = v19
                                _ = v20
                            }
                            for fi29 := 0; fi29 < 7; fi29++ {
                                var v30 string = v5
                                _ = fi29
                                _ = v30
                            }
                            _ = fi17
                        }
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v5 = (v5 + v5)
                            v32 := GenS2{F0: 610, F1: 223, F2: 1}
                            v3 = append(v3, -9223372036854775808)
                            _ = fi31
                            _ = v32
                        }
                    }
                default:
                    v3 = append(v3, v6)
                }
                var v33 uint32 = 42
                var v34 uint16 = 801
                _ = v33
                _ = v34
            }
            if ((true || (97 < 65535)) && false) {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v35 int8 = (v4 - (v4 + v4))
                for fi36 := 0; fi36 < 2; fi36++ {
                    if len(v3) > 0 {
                        v3[0] = (fi36 + (fi36 * (v6 ^ (-(v6)))))
                    }
                    v3 = append(v3, 396)
                    v5 = (v5 + v5)
                    _ = fi36
                }
                _ = v35
            }
        }
    case "\n\t":
        var v37 string = v5
        _ = v37
    }
    if false {
        switch v6 {
        case 0:
            switch (v5 + "") {
            case "hello":
                fmt.Println(v5, v2)
            case "":
                switch 0 {
                case 976:
                    if true {
                        v1 -= (v1 << ((898) & 63))
                    }
                case 699:
                    v2 = v2
                case 32767:
                    if (42 != 401) {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            var v39 int = fi38
                            v3 = append(v3, v39)
                            _ = fi38
                            _ = v39
                        }
                    }
                default:
                    if (2147483647 > 2147483647) {
                        v5 = ("" + "hello")
                        if (-1 <= 342) {
                            v4 = (v4 % (((^(v4))) | 1))
                            fmt.Println(v1, v2, v4)
                            switch -9223372036854775808 {
                            case -985:
                                v4 = v4
                            }
                        } else {
                            for fi40 := 0; fi40 < 11; fi40++ {
                                var v41 uint32 = 2147483648
                                v4 -= v4
                                v3 = append(v3, fi40)
                                _ = fi40
                                _ = v41
                            }
                            switch -9223372036854775808 {
                            case -542:
                                switch v6 {
                                case 127:
                                    v3 = append(v3, 2147483647)
                                case 128:
                                    v2 = (v2 * 1e10)
                                case -755:
                                    if (!false) {
                                        if (688 == 349) {
                                            v6 = 32767
                                            v1 -= v1
                                        }
                                        var v42 uint64 = 1
                                        _ = v42
                                    } else {
                                        fmt.Println(v1, v5, v1)
                                        v1 = v1
                                    }
                                }
                            default:
                                v6 += -1
                            }
                            v43 := GenS2{F0: 2, F1: 2147483647, F2: 18446744073709551615}
                            _ = v43
                        }
                        var v44 uint16 = 99
                        _ = v44
                    } else {
                        _ = len(v3)
                        _ = len(v3)
                    }
                }
            case "a":
                v2 -= v2
            }
        default:
            v2 = v2
        }
    }
    var v45 string = v5
    v45 = ("" + "hello")
    for fi46 := 0; fi46 < 1; fi46++ {
        fmt.Println(v45)
        _ = fi46
    }
    var v47 uint16 = 1
    if false {
        for fi48 := 0; fi48 < 13; fi48++ {
            if len(v3) > 0 {
                v3[0] = 1
            }
            v47 = 1
            v4 -= v4
            _ = fi48
        }
        v5 = "foo bar"
    } else {
        if (9223372036854775807 < -1) {
            var v49 int32 = -2147483648
            _ = v49
        } else {
            v6 = 4294967295
            if false {
                v2 = -2.25
                v50 := GenS2{F0: 361, F1: 2147483647, F2: 66}
                if false {
                    var v51 int = (v6 + v6)
                    for fi52 := 0; fi52 < 17; fi52++ {
                        var v53 uint16 = uint16(v1)
                        v51 = 614
                        var v54 int64 = -1
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    _ = v51
                }
                _ = v50
            }
            var v55 int64 = 854
            _ = v55
        }
    }
    var v56 int64 = -1
    if false {
        if false {
            v57 := GenS3{F0: uint8(v1)}
            _ = v57
        }
    } else {
        var v58 uint64 = (v1 * (v1 / ((v1) | 1)))
        v58 = v1
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v47
    _ = v5
    _ = v56
    _ = v6
}
