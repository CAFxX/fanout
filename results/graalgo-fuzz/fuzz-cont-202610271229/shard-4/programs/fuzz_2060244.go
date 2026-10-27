package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 int32; F2 int8 }
type GenS4 struct { F0 int16; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint8 }

func main() {
    var v1 bool = true
    var v2 string = ""
    var v3 uint32 = 465
    v4 := []uint32{(v3 / ((65535) | 1)), v3, v3}
    v5 := make(map[uint32]string)
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v1 = v1
            if (9223372036854775807 < (fi6 + fi7)) {
                v1 = (203 < 255)
                switch "" {
                case "":
                    if ((!v1) && v1) {
                        if v1 {
                            v2 = (v2 + v2)
                            v8 := GenS1{F0: 42}
                            var v9 uint16 = 65535
                            _ = v8
                            _ = v9
                        }
                        v10 := GenS1{F0: -607}
                        if false {
                            v2 = "\n\t"
                            v11 := v5[(v3 >> ((65535) & 63))]
                            switch (^(fi7)) {
                            case 4294967295:
                                _ = len(v4)
                            }
                            _ = v11
                        }
                        _ = v10
                    }
                }
                v1 = (1 >= -1)
            } else {
                v1 = (32767 != 0)
            }
            var v12 uint64 = 18446744073709551615
            var v13 uint16 = 2
            v14 := GenS1{F0: -576}
            _ = fi7
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    v15 := GenS2{F0: 42, F1: 131, F2: ((((3.14159 - ((1.5 / 3.14159) - 1e10)) / 1.0) - 1.5) <= -2.25)}
    if (!(42 == -1)) {
        fmt.Println(v3, v1)
    }
    if len(v4) > 0 {
        v4[0] = v3
    }
    v16 := GenS1{F0: 243}
    _ = len(v5)
    if v1 {
        if v1 {
            for fi17 := 0; fi17 < 2; fi17++ {
                switch v2 {
                case "":
                    v18 := GenS1{F0: 2147483647}
                    _ = v18
                }
                _ = fi17
            }
        } else {
            var v19 uint64 = 1
            if false {
                var v20 int = 709
                for fi21 := 0; fi21 < 19; fi21++ {
                    for fi22 := 0; fi22 < 13; fi22++ {
                        var v23 float32 = 1.5
                        _ = fi22
                        _ = v23
                    }
                    if (v3 < v3) {
                        for fi24 := 0; fi24 < 17; fi24++ {
                            fmt.Println(fi21, fi24, v2)
                            var v25 uint = 18446744073709551615
                            _ = fi24
                            _ = v25
                        }
                        v3 -= v3
                        for fi26 := 0; fi26 < 18; fi26++ {
                            v27 := GenS3{F0: -296, F1: -1, F2: 42}
                            _ = fi26
                            _ = v27
                        }
                    }
                    _ = fi21
                }
                var v28 uint16 = 32767
                _ = v20
                _ = v28
            } else {
                var v29 int64 = -729
                v2 = v2
                var v30 int32 = 0
                _ = v29
                _ = v30
            }
            if (1 > 939) {
                switch v2 {
                case "a":
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 uint32 = 65535
                        var v33 int = fi31
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                case "\n\t":
                    v2 = ((("hello" + "a") + "a") + "\n\t")
                }
            } else {
                v2 = "\n\t"
                var v34 int8 = -26
                _ = v34
            }
            _ = v19
        }
    }
    var v35 uint8 = 0
    delete(v5, v3)
    var v36 float64 = ((0.1 - 1.5) - 3.14159)
    fmt.Println(v35, v2)
    if (!true) {
        switch -2147483648 {
        case -780:
            if ((0 <= 838) || (42 == 623)) {
                for fi37 := 0; fi37 < 17; fi37++ {
                    var v38 int32 = 2147483647
                    _ = fi37
                    _ = v38
                }
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 uint = 42
                    _ = fi39
                    _ = v40
                }
                if (607 < uint64(v35)) {
                    if v1 {
                        var v41 bool = ((int8(v3) > 0) || (1e10 == (((1.5 * (float64((1e10 + ((-2.25 * (1e10 - 0.1)) + (-2.25 - float32(v36))))) + 0.0)) * (((((v36 * v36) + v36) / (0.0 + float64(1.5))) + v36) / (1e10 * 3.14159))) - 3.14159)))
                        _ = v41
                    } else {
                        for fi42 := 0; fi42 < 11; fi42++ {
                            switch fi42 {
                            case 4294967295:
                                var v43 int = (fi42 - (fi42 / (((fi42 * (fi42 >> ((0) & 63)))) | 1)))
                                _ = v43
                            case 2147483647:
                                v2 = v2
                            case 9223372036854775807:
                                if v1 {
                                    var v44 int64 = -524
                                    v44 -= -686
                                    _ = len(v5)
                                    _ = v44
                                } else {
                                    v3 += v3
                                    v1 = (1 > 1)
                                }
                            }
                            if v1 {
                                var v45 uint16 = 871
                                v46 := GenS4{F0: 81, F1: 65535, F2: 42}
                                _ = v45
                                _ = v46
                            }
                            v47 := GenS5{F0: (v3 * (v3 << ((2) & 63)))}
                            _ = fi42
                            _ = v47
                        }
                        v48 := GenS6{F0: (^(v3)), F1: v3, F2: 0}
                        _ = v48
                    }
                }
            } else {
                v36 = v36
                var v49 uint = 844
                v50 := GenS3{F0: 42, F1: 343, F2: 127}
                _ = v49
                _ = v50
            }
        }
    } else {
        fmt.Println(v3, v35, v2)
        fmt.Println(v1, v3, v36)
        v51 := GenS7{F0: 1, F1: (v3 & v3), F2: 1}
        _ = v51
    }
    var v52 uint8 = v35
    v53 := GenS2{F0: 1, F1: 42, F2: v1}
    for fi54 := 0; fi54 < 10; fi54++ {
        for fi55 := 0; fi55 < 8; fi55++ {
            fmt.Println(v2, v2)
            var v56 uint32 = v3
            _ = fi55
            _ = v56
        }
        for fi57 := 0; fi57 < 16; fi57++ {
            v58 := GenS2{F0: 788, F1: 255, F2: true}
            _ = fi57
            _ = v58
        }
        _ = fi54
    }
    v35 = (^(v35))
    var v59 int8 = -128
    switch 0 {
    case -355:
        for fi60 := 0; fi60 < 6; fi60++ {
            v3 = 2147483648
            v36 = v36
            v5[(v3 + (v3 - 2147483648))] = "a"
            _ = fi60
        }
    case 2147483647:
        v61 := GenS6{F0: 0, F1: 2, F2: 782}
        _ = v61
    default:
        v3 -= 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v59
}
