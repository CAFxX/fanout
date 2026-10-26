package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint16; F2 int16 }
type GenS6 struct { F0 int32 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    var v3 int16 = 0
    var v4 uint = 2
    v5 := []uint32{0}
    v6 := make(map[uint32]int)
    if true {
        delete(v6, 149)
        var v7 int = -44
        for fi8 := 0; fi8 < 2; fi8++ {
            v3 = 42
            var v9 int8 = -128
            fmt.Println(v9, v3, v4)
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    if true {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int = (-(fi10))
            v5 = append(v5, 520)
            _ = fi10
            _ = v11
        }
        if true {
            var v12 bool = false
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            switch -1 {
            case 9223372036854775807:
                for fi14 := 0; fi14 < 5; fi14++ {
                    v4 = (^(v4))
                    v15 := GenS1{F0: ((128 == 1) && false), F1: (v3 - v3), F2: 42}
                    _ = fi14
                    _ = v15
                }
            default:
                for fi16 := 0; fi16 < 2; fi16++ {
                    fmt.Println(fi16)
                    _ = fi16
                }
            }
            _ = fi13
        }
    }
    if (true && false) {
        v3 = v3
        var v17 int32 = -205
        for fi18 := 0; fi18 < 13; fi18++ {
            for fi19 := 0; fi19 < 4; fi19++ {
                switch ("\n\t" + "foo bar") {
                case "a":
                    var v20 uint32 = 4294967295
                    _ = v20
                case "foo bar":
                    var v21 uint = v4
                    _ = v21
                }
                switch "foo bar" {
                case "hello":
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                case "":
                    v17 -= v17
                }
                _ = fi19
            }
            v3 -= int16(v17)
            var v22 int16 = v3
            _ = fi18
            _ = v22
        }
        _ = v17
    }
    var v23 float64 = 0.0
    v23 = v23
    v24 := GenS2{F0: v4}
    v23 += float64(-2.25)
    var v25 int = 2147483647
    var v26 uint16 = 91
    v26 = (v26 ^ (^(v26)))
    v23 -= float64(1e10)
    for fi27 := 0; fi27 < 10; fi27++ {
        fmt.Println(v26, v23, v23)
        var v28 uint8 = 197
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        if true {
            v3 += v3
            for fi30 := 0; fi30 < 2; fi30++ {
                delete(v1, "a")
                _ = fi30
            }
        }
        _ = fi29
    }
    switch "" {
    case "\n\t":
        var v31 uint16 = v26
        _ = v31
    }
    for fi32 := 0; fi32 < 14; fi32++ {
        v5 = append(v5, 0)
        if (((float64((((((float32(3.14159) * 1.5) + float32((v23 * 0.0))) - float32(v23)) + 1.5) / (float32((v23 * ((v23 - 0.1) * 3.14159))) / 1.5))) - v23) + ((1e10 - (-2.25 / float64(1.5))) + 1.5)) > v23) {
            var v33 uint8 = uint8(v4)
            _ = v33
        } else {
            for fi34 := 0; fi34 < 5; fi34++ {
                switch (v25 + (-(fi34))) {
                case 2147483647:
                    v25 = (v25 & v25)
                case 1:
                    switch ("a" + "\n\t") {
                    case "\n\t":
                        switch "" {
                        case "hello":
                            _ = len(v6)
                        case "\n\t":
                            v25 -= fi34
                        case "a":
                            var v35 int32 = -1
                            _ = v35
                        default:
                            if (((v23 * -2.25) + 3.14159) == float64((1.5 * (3.14159 * (float32((float64(float32(((1e10 / v23) - v23))) / 3.14159)) - 1.5))))) {
                                v36 := GenS3{F0: v26}
                                v37 := GenS2{F0: (^(v4))}
                                _ = v36
                                _ = v37
                            } else {
                                v38 := GenS4{F0: -9223372036854775808, F1: 0, F2: 0}
                                v4 += v4
                                v3 = -236
                                v26 = 32767
                                v26 += uint16(v4)
                                v39 := GenS4{F0: -1, F1: 42, F2: 255}
                                _ = v38
                                _ = v39
                            }
                            switch fi34 {
                            case 2147483647:
                                v26 += (v26 + (v26 + v26))
                            default:
                                var v40 int = (v25 >> ((64) & 63))
                                _ = v40
                            }
                            v25 = 128
                        }
                    case "":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "a":
                        switch (-(fi34)) {
                        case -1:
                            var v41 uint32 = 593
                            _ = v41
                        case 32768:
                            var v42 uint8 = 255
                            _ = v42
                        case 32767:
                            switch fi32 {
                            case 32767:
                                if true {
                                    if false {
                                        fmt.Println(fi34, v26, fi32)
                                        fmt.Println(v25)
                                        v25 = fi32
                                    }
                                    _ = len(v5)
                                } else {
                                    v43 := GenS5{F0: true, F1: v26, F2: v3}
                                    if false {
                                        var v44 int8 = 0
                                        var v45 float32 = 1e10
                                        _ = v44
                                        _ = v45
                                    } else {
                                        v1["\n\t"] = "a"
                                    }
                                    if (78 == 2) {
                                        var v46 uint32 = 2147483648
                                        v6[(v46 - 65535)] = 4294967295
                                        _ = v46
                                    } else {
                                        var v47 uint8 = 1
                                        if false {
                                            fmt.Println(v25, v26)
                                            v48 := GenS5{F0: true, F1: (v26 / (((v26 | (v26 << (((^(v4))) & 63)))) | 1)), F2: v3}
                                            var v49 uint = v4
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = v47
                                    }
                                    _ = v43
                                }
                            case 919:
                                v50 := GenS2{F0: 0}
                                _ = v50
                            default:
                                v51 := GenS6{F0: -2147483648}
                                _ = v51
                            }
                        }
                    }
                }
                fmt.Println(v23, v26)
                v23 -= v23
                switch (fi32 + (fi32 & 657)) {
                case 122:
                    switch 1 {
                    case 32767:
                        var v52 int16 = (^(v3))
                        _ = v52
                    }
                }
                var v53 int = int(v26)
                _ = len(v2)
                var v54 uint32 = 610
                fmt.Println(v26, v26, v4)
                _ = len(v1)
                v53 = (fi34 * -9223372036854775808)
                v26 -= (v26 & v26)
                fmt.Println(v26, v53)
                _ = fi34
                _ = v53
                _ = v54
            }
            var v55 uint = 2
            _ = v55
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
