package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int16 }
type GenS2 struct { F0 uint8; F1 int32; F2 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 int32; F1 int8 }
func genFn0(p0 uint16, p1 int32) bool {
    v4 := []int{int(p0), 1, 42}
    v5 := GenS1{F0: 396, F1: -8, F2: -300}
    fmt.Println(p1, p0)
    var v6 uint16 = p0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return (2 < 23)
}
func genFn1(p0 int, p1 uint16, p2 uint64) int8 {
    var v7 string = ""
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case 127:
            v9 := []int{}
            _ = v9
        default:
            if true {
                v10 := GenS1{F0: 501, F1: 1, F2: 42}
                fmt.Println(fi8, p2, p1)
                _ = v10
            } else {
                var v11 float32 = float32(3.14159)
                if (3.14159 > (1e10 - 0.1)) {
                    v12 := []int{(p0 | (fi8 + 891)), (fi8 & -2147483648), -2147483648}
                    _ = v12
                }
                _ = v11
            }
        }
        _ = fi8
    }
    var v13 string = v7
    v14 := []int{}
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v7
    return 72
}

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[uint32]bool)
    var v3 bool = true
    v3 = ((32767 <= 127) && (-32768 <= 0))
    for fi15 := 0; fi15 < 19; fi15++ {
        switch (fi15 - fi15) {
        case -235:
            if v3 {
                var v16 int = -2147483648
                _ = v16
            }
        case -1:
            var v17 uint8 = 127
            _ = v17
        case 1:
            v3 = ((46 < 769) || v3)
        default:
            if v3 {
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 float64 = 0.0
                    v20 := GenS2{F0: 1, F1: 0, F2: 42}
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 uint64 = 337
                _ = v21
            }
        }
        v3 = (2 >= 42)
        v22 := GenS2{F0: 0, F1: -2147483648, F2: 0}
        _ = fi15
        _ = v22
    }
    var v23 uint8 = 2
    var v24 float32 = 1.5
    fmt.Println(v3, v3)
    if v3 {
        v24 -= ((v24 + v24) / (float32(-2.25) * (v24 * (0.0 - v24))))
    }
    var v25 int = 246
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := GenS1{F0: 4294967295, F1: -86, F2: 42}
        switch ("" + "foo bar") {
        case "a":
            switch fi26 {
            case 4294967295:
                if v3 {
                    v2[65] = (!("" <= "foo bar"))
                } else {
                    fmt.Println(v24, v24, v3)
                    for fi28 := 0; fi28 < 17; fi28++ {
                        v29 := GenS2{F0: v23, F1: 1, F2: -9223372036854775808}
                        v30 := GenS1{F0: 2, F1: -128, F2: -45}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                }
            case 695:
                var v31 uint = 42
                _ = v31
            case 646:
                v23 -= 1
            default:
                _ = len(v1)
            }
        case "hello":
            for fi32 := 0; fi32 < 9; fi32++ {
                v33 := GenS2{F0: v23, F1: int32(v23), F2: -9223372036854775808}
                var v34 uint32 = 2
                _ = fi32
                _ = v33
                _ = v34
            }
        case "foo bar":
            switch "a" {
            case "hello":
                var v35 bool = v3
                _ = v35
            case "foo bar":
                _ = len(v2)
            }
        default:
            var v36 float32 = v24
            _ = v36
        }
        _ = fi26
        _ = v27
    }
    v37 := GenS2{F0: 128, F1: 2147483647, F2: 609}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    var v38 float32 = 1.5
    if (true && v3) {
        if (722 != 969) {
            var v39 uint16 = 823
            _ = v39
        } else {
            var v40 uint = 77
            fmt.Println(v3, v3)
            _ = v40
        }
        var v41 string = ("\n\t" + "foo bar")
        for fi42 := 0; fi42 < 8; fi42++ {
            if v3 {
                switch 32767 {
                case -936:
                    v43 := GenS3{F0: v3, F1: -622, F2: int(v23)}
                    _ = v43
                case 127:
                    switch (v41 + "hello") {
                    case "\n\t":
                        v25 = (fi42 % ((v25) | 1))
                    case "hello":
                        var v44 uint32 = 0
                        _ = v44
                    case "foo bar":
                        if v3 {
                            for fi45 := 0; fi45 < 20; fi45++ {
                                fmt.Println(v25, v38)
                                v24 = float32(float64(v24))
                                if true {
                                    var v46 bool = false
                                    v47 := GenS2{F0: (^(v23)), F1: 0, F2: 33}
                                    delete(v2, 2)
                                    switch "hello" {
                                    case "\n\t":
                                        v24 = (0.1 * 1.5)
                                    default:
                                        var v48 int64 = 0
                                        _ = v48
                                    }
                                    _ = v46
                                    _ = v47
                                }
                                _ = fi45
                            }
                        } else {
                            for fi49 := 0; fi49 < 11; fi49++ {
                                v38 += float32(0.1)
                                _ = fi49
                            }
                            var v50 float64 = 0.1
                            _ = v50
                        }
                    }
                }
            } else {
                var v51 uint8 = v23
                _ = v51
            }
            v52 := GenS1{F0: 927, F1: 100, F2: 42}
            var v53 uint32 = 4294967295
            _ = fi42
            _ = v52
            _ = v53
        }
        _ = v41
    } else {
        v25 = v25
    }
    if (236 <= 2) {
        fmt.Println(v24)
        for fi54 := 0; fi54 < 14; fi54++ {
            for fi55 := 0; fi55 < 4; fi55++ {
                v56 := GenS4{F0: -300, F1: 100}
                _ = fi55
                _ = v56
            }
            delete(v2, 42)
            var v57 uint32 = 2
            _ = fi54
            _ = v57
        }
        for fi58 := 0; fi58 < 1; fi58++ {
            var v59 int32 = 370
            var v60 uint8 = (v23 % ((v23) | 1))
            for fi61 := 0; fi61 < 1; fi61++ {
                v3 = v3
                v59 = -1
                _ = fi61
            }
            _ = fi58
            _ = v59
            _ = v60
        }
    }
    fmt.Println(v23, v23, v38)
    v23 = 128
    v23 = v23
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
}
