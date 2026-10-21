package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint16 = 32767
    var v2 string = ("a" + "hello")
    var v3 float32 = 1.5
    v4 := make(map[string]string)
    v5 := []string{}
    if len(v5) > 0 {
        v5[0] = (v2 + ((("foo bar" + "a") + v2) + v2))
    }
    var v6 string = (((v2 + ("" + "hello")) + "a") + ((v2 + v2) + "foo bar"))
    for fi7 := 0; fi7 < 14; fi7++ {
        if false {
            v2 = "a"
            v8 := GenS1{F0: 359, F1: 1, F2: 127}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 16; fi9++ {
                _ = len(v4)
                _ = fi9
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                if true {
                    switch (fi10 / ((fi7) | 1)) {
                    case -853:
                        var v11 int16 = 1
                        _ = v11
                    case 32767:
                        fmt.Println(v3)
                    case 4294967295:
                        v6 = v2
                    }
                } else {
                    switch (fi10 ^ fi10) {
                    case 128:
                        v12 := GenS2{F0: true, F1: 1}
                        _ = v12
                    case 1:
                        v13 := GenS2{F0: (float64(v3) <= float64(v3)), F1: 2}
                        _ = v13
                    }
                }
                _ = fi10
            }
            if (false && (!(1 >= -172))) {
                if true {
                    if ((((("foo bar" + v2) + v6) + v2) + v2) >= v6) {
                        var v14 int32 = 94
                        v5 = append(v5, v2)
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            switch 4294967295 {
                            case 128:
                                v3 -= (v3 - v3)
                            default:
                                v16 := GenS2{F0: true, F1: 271}
                                _ = v16
                            }
                            if false {
                                v6 = v6
                            } else {
                                v3 = float32(1e10)
                            }
                            if (9223372036854775807 <= int64(fi15)) {
                                var v17 uint16 = 1
                                var v18 string = ((v2 + v6) + "")
                                switch fi15 {
                                case 32768:
                                    v17 = (v17 >> ((4294967295) & 63))
                                default:
                                    var v19 string = v2
                                    _ = v19
                                }
                                _ = v17
                                _ = v18
                            } else {
                                var v20 uint64 = 744
                                switch fi7 {
                                case 128:
                                    fmt.Println(v1)
                                case -681:
                                    var v21 int8 = 1
                                    _ = v21
                                case -2147483648:
                                    v20 -= 2
                                default:
                                    var v22 bool = (4294967295 == 2)
                                    _ = v22
                                }
                                if true {
                                    fmt.Println(v2, v3, v20)
                                }
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                _ = v20
                            }
                            switch fi7 {
                            case 672:
                                var v23 float32 = v3
                                _ = v23
                            case 2147483647:
                                v24 := GenS2{F0: (881 != 76), F1: 2147483647}
                                _ = v24
                            case 786:
                                var v25 int64 = 42
                                switch (fi7 | (fi15 / (((fi7 ^ (^(fi15)))) | 1))) {
                                case -1:
                                    v26 := v4[(v6 + v6)]
                                    _ = v26
                                }
                                _ = v25
                            }
                            fmt.Println(fi7, v3, fi7)
                            _ = fi15
                        }
                    }
                    v27 := GenS1{F0: 622, F1: 781, F2: -128}
                    if false {
                        var v28 int64 = -555
                        _ = v28
                    }
                    _ = v27
                } else {
                    var v29 int8 = -58
                    var v30 uint8 = 128
                    _ = v29
                    _ = v30
                }
                fmt.Println(v6, v3)
            }
        }
        var v31 int32 = 2147483647
        _ = fi7
        _ = v31
    }
    v32 := GenS2{F0: ((0.1 <= (v3 - v3)) || true), F1: 224}
    v1 = (v1 / (((v1 >> ((65535) & 63))) | 1))
    v33 := v4[v6]
    switch 42 {
    case 128:
        if (9223372036854775807 >= 1) {
            fmt.Println(v33, v2)
            var v34 float64 = 1.5
            _ = v34
        } else {
            var v35 int = -115
            _ = v35
        }
    case 127:
        switch v6 {
        case "hello":
            switch (v2 + "\n\t") {
            case "hello":
                for fi36 := 0; fi36 < 16; fi36++ {
                    v4[(v2 + v6)] = v33
                    v37 := GenS1{F0: 42, F1: 42, F2: 127}
                    for fi38 := 0; fi38 < 12; fi38++ {
                        v39 := GenS1{F0: 42, F1: 773, F2: 0}
                        v40 := GenS1{F0: 106, F1: 18446744073709551615, F2: 0}
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                    _ = fi36
                    _ = v37
                }
            }
        }
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        var v42 int = (fi41 << ((202) & 63))
        v4["\n\t"] = ((v2 + (v33 + v33)) + v33)
        _ = fi41
        _ = v42
    }
    v43 := GenS3{F0: -486}
    v1 = (v1 + (v1 % ((v1) | 1)))
    for fi44 := 0; fi44 < 1; fi44++ {
        fmt.Println(v33, v2, v33)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v45 := GenS1{F0: 862, F1: 4294967295, F2: -128}
        _ = fi44
        _ = v45
    }
    _ = len(v5)
    var v46 int32 = 437
    if true {
        fmt.Println(v1)
        if ((v46 / ((0) | 1)) != (v46 + v46)) {
            _ = len(v5)
        } else {
            if false {
                for fi47 := 0; fi47 < 15; fi47++ {
                    if (true || (2 > v1)) {
                        if true {
                            var v48 uint16 = v1
                            _ = v48
                        }
                    } else {
                        var v49 string = ((v2 + v6) + ("\n\t" + v2))
                        var v50 int8 = -1
                        var v51 int16 = -32768
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    v46 -= v46
                    for fi52 := 0; fi52 < 19; fi52++ {
                        fmt.Println(v2)
                        v53 := GenS1{F0: 0, F1: 2, F2: -128}
                        fmt.Println(v2)
                        _ = fi52
                        _ = v53
                    }
                    _ = fi47
                }
                v54 := GenS4{F0: false, F1: 100}
                for fi55 := 0; fi55 < 8; fi55++ {
                    if true {
                        var v56 uint32 = 2
                        _ = v56
                    }
                    var v57 uint = 884
                    var v58 float64 = 1e10
                    _ = fi55
                    _ = v57
                    _ = v58
                }
                _ = v54
            }
            v5 = append(v5, v6)
        }
        for fi59 := 0; fi59 < 10; fi59++ {
            if true {
                v60 := GenS1{F0: 2, F1: uint64(v1), F2: 1}
                v61 := GenS3{F0: -727}
                _ = v60
                _ = v61
            }
            _ = fi59
        }
    } else {
        v46 = v46
        v1 -= (v1 - v1)
    }
    v1 -= v1
    v3 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v6
}
