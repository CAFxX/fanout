package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint = 1
    v2 := []int{32767, -9223372036854775808}
    v3 := []uint32{65535, 2, 2147483647}
    var v4 int64 = -9223372036854775808
    var v5 string = "hello"
    v6 := []uint32{169, 42, 346, 42}
    v7 := GenS1{F0: 0}
    var v8 float64 = (-2.25 - 0.1)
    switch 32768 {
    case 32768:
        var v9 uint32 = 1
        _ = v9
    case 9223372036854775807:
        if false {
            fmt.Println(v4, v4)
        } else {
            if false {
                v10 := GenS1{F0: 2}
                for fi11 := 0; fi11 < 20; fi11++ {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        if len(v6) > 0 {
                            v6[0] = 0
                        }
                        _ = fi12
                    }
                    v13 := GenS1{F0: 4294967295}
                    if (false && true) {
                        v1 = (v1 & 2)
                        switch "\n\t" {
                        case "":
                            var v14 int64 = (v4 - 1)
                            _ = v14
                        }
                    }
                    _ = fi11
                    _ = v13
                }
                fmt.Println(v8)
                _ = v10
            }
        }
    default:
        v6 = append(v6, 371)
    }
    var v15 uint16 = 2
    v16 := GenS1{F0: 65535}
    v17 := GenS2{F0: (v5 != ("\n\t" + "")), F1: v15, F2: 255}
    for fi18 := 0; fi18 < 10; fi18++ {
        switch fi18 {
        case -9223372036854775808:
            var v19 int = (^(fi18))
            _ = v19
        default:
            if (true || ((1 == 0) && true)) {
                var v20 bool = true
                if v20 {
                    if (1 == -672) {
                        var v21 uint = v1
                        _ = v21
                    }
                    v5 = "\n\t"
                } else {
                    v22 := GenS2{F0: true, F1: (v15 & v15), F2: 127}
                    if v20 {
                        v4 -= 9223372036854775807
                    }
                    var v23 uint32 = 42
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        }
        v8 = (-2.25 - 0.0)
        var v24 int16 = 691
        _ = fi18
        _ = v24
    }
    v4 += 9223372036854775807
    _ = len(v6)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        switch fi25 {
        case 32768:
            v4 -= int64(v1)
        case 42:
            var v26 bool = false
            _ = v26
        case 1:
            if len(v6) > 0 {
                v6[0] = 4294967295
            }
        default:
            switch (((v5 + "hello") + v5) + (v5 + (v5 + v5))) {
            case "":
                for fi27 := 0; fi27 < 7; fi27++ {
                    switch ((("\n\t" + (v5 + v5)) + (v5 + "")) + v5) {
                    case "hello":
                        v28 := GenS1{F0: 65535}
                        _ = v28
                    case "a":
                        var v29 uint16 = (v15 ^ 553)
                        _ = v29
                    }
                    v30 := GenS1{F0: 319}
                    var v31 int = fi25
                    if false {
                        fmt.Println(fi25, v8)
                        _ = len(v3)
                    }
                    if false {
                        fmt.Println(fi27, v31)
                        if true {
                            v32 := GenS3{F0: v4, F1: 2, F2: 875}
                            _ = v32
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v33 bool = false
                            if (false && (((-32768 < -1) || (2 < 284)) || (v33 && v33))) {
                                var v34 int16 = -1
                                v35 := GenS2{F0: false, F1: v15, F2: 2}
                                var v36 float32 = ((1e10 / 3.14159) + (-2.25 - (0.0 - 0.1)))
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                        }
                        if true {
                            var v37 string = "a"
                            if (true && true) {
                                v8 = ((v8 * ((v8 + v8) - (v8 / v8))) - v8)
                                v8 = float64(0.1)
                            }
                            fmt.Println(v37, v5, v8)
                            _ = v37
                        } else {
                            var v38 int16 = int16(v15)
                            switch (-(v31)) {
                            case 9223372036854775807:
                                switch "\n\t" {
                                case "a":
                                    _ = len(v6)
                                case "hello":
                                    switch v5 {
                                    case "":
                                        if len(v3) > 0 {
                                            v3[0] = 1
                                        }
                                    case "hello":
                                        v38 -= -1
                                    }
                                case "\n\t":
                                    if false {
                                        v4 += (v4 + v4)
                                        fmt.Println(fi27, fi27, v1)
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    } else {
                                        v2 = append(v2, fi25)
                                        v8 += (3.14159 - (1.5 / (1e10 / 0.1)))
                                    }
                                default:
                                    v39 := GenS3{F0: -1, F1: 42, F2: 2147483647}
                                    _ = v39
                                }
                            default:
                                var v40 float32 = 1.5
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v40
                            }
                            switch fi25 {
                            case 127:
                                v5 = v5
                            case -361:
                                var v41 uint64 = 4294967295
                                _ = v41
                            case -396:
                                v15 += (v15 + (v15 / (((v15 + v15)) | 1)))
                                v38 += v38
                                if true {
                                    v31 += fi25
                                    v8 += (float64(0.1) - (v8 + float64(float32(v8))))
                                    v42 := GenS1{F0: 4294967295}
                                    var v43 uint16 = uint16(v1)
                                    var v44 int8 = 86
                                    var v45 int32 = 1
                                    v4 = int64(v45)
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                    _ = v45
                                } else {
                                    if (false || (!(100 < -1))) {
                                        if true {
                                            var v46 uint8 = 127
                                            _ = v46
                                        } else {
                                            fmt.Println(v1, v15)
                                        }
                                    }
                                    var v47 float32 = float32(v8)
                                    if (((v38 - (v38 * (v38 % ((v38) | 1)))) < v38) || (v15 == (^(v15)))) {
                                        if false {
                                            if (!false) {
                                                v48 := GenS3{F0: v4, F1: 302, F2: 947}
                                                var v49 int16 = -32768
                                                _ = v48
                                                _ = v49
                                            }
                                        }
                                        var v50 int = fi25
                                        v38 = v38
                                        if (1 < 2147483648) {
                                            fmt.Println(v31)
                                            var v51 float64 = v8
                                            _ = v51
                                        }
                                        var v52 uint64 = uint64(v1)
                                        _ = v50
                                        _ = v52
                                    }
                                    _ = v47
                                }
                                fmt.Println(v15)
                                var v53 int8 = -81
                                v1 = uint(v15)
                                _ = v53
                            }
                            v15 = v15
                            switch 32767 {
                            case -1:
                                fmt.Println(v8, fi25, v4)
                            case 32767:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            _ = v38
                        }
                    }
                    fmt.Println(fi25)
                    _ = fi27
                    _ = v30
                    _ = v31
                }
            }
        }
        for fi54 := 0; fi54 < 13; fi54++ {
            fmt.Println(v1, v1)
            _ = fi54
        }
        _ = fi25
    }
    v5 = v5
    for fi55 := 0; fi55 < 6; fi55++ {
        v56 := GenS3{F0: v4, F1: 2, F2: 263}
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 17; fi57++ {
        var v58 uint16 = 605
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        var v60 int = 127
        var v61 uint8 = 42
        v62 := GenS4{F0: -765, F1: -1}
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
