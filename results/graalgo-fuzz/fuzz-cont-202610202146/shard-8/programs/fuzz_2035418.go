package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 int32 }
type GenS4 struct { F0 int64; F1 bool; F2 int16 }
type GenS5 struct { F0 int16; F1 string }
func genFn0() uint {
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint16 = 255
        switch (fi7 - 128) {
        case 1:
            v9 := []int{(fi7 & fi7), 127, fi7}
            _ = v9
        case 187:
            v8 -= (^(v8))
        case 655:
            switch "" {
            case "hello":
                var v10 int = fi7
                _ = v10
            case "":
                fmt.Println(fi7)
            }
        default:
            switch "foo bar" {
            case "\n\t":
                v8 -= v8
            default:
                var v11 int = -9223372036854775808
                _ = v11
            }
        }
        _ = fi7
        _ = v8
    }
    if (885 <= 9223372036854775807) {
        v12 := []int{128, -541, -9223372036854775808, 2147483647}
        _ = v12
    }
    var v13 uint = 742
    v13 = (v13 + v13)
    v13 = 1
    _ = v13
    return 463
}
func genFn1(p0 uint16, p1 int32) uint {
    for fi14 := 0; fi14 < 13; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            if (32767 >= 519) {
                switch "\n\t" {
                case "\n\t":
                    switch "hello" {
                    case "a":
                        var v16 float64 = 1e10
                        _ = v16
                    case "":
                        p1 -= (p1 / (((p1 | p1)) | 1))
                    case "hello":
                        var v17 int = 42
                        _ = v17
                    }
                }
                v18 := []int{fi14}
                _ = v18
            }
            p0 -= 1
            _ = fi15
        }
        var v19 uint64 = 42
        _ = fi14
        _ = v19
    }
    p0 = p0
    switch 0 {
    case 42:
        if false {
            var v20 uint64 = 0
            _ = v20
        } else {
            v21 := []int{32768, 127}
            p0 -= (p0 >> ((0) & 63))
            _ = v21
        }
    case -1:
        v22 := GenS1{F0: (p0 % ((p0) | 1)), F1: p0, F2: 1}
        _ = v22
    case 32767:
        p0 = 1
    }
    if false {
        p0 -= (^(p0))
        v23 := []int{int(p1), -872, 0}
        _ = v23
    } else {
        p0 = p0
        var v24 uint64 = 893
        _ = v24
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int8 = -79
    v2 := make(map[uint32]string)
    v3 := []uint32{464}
    v4 := make(map[int]string)
    var v5 float64 = 1e10
    v6 := []uint32{65535, 2147483647, 0, 0}
    v25 := GenS1{F0: 65535, F1: 200, F2: 42}
    _ = len(v3)
    v26 := GenS1{F0: 42, F1: 255, F2: 1}
    _ = len(v2)
    var v27 bool = (42 > 1)
    v5 -= float64(3.14159)
    v27 = v27
    switch -538 {
    case -37:
        switch 42 {
        case -2147483648:
            if true {
                v27 = (!false)
                if false {
                    var v28 uint16 = 1
                    _ = v28
                }
                fmt.Println(v1)
            }
        case 2:
            if len(v3) > 0 {
                v3[0] = 712
            }
        case -1:
            if (4294967295 <= 0) {
                if (763 < 240) {
                    _ = len(v2)
                    if (32767 == 235) {
                        v29 := GenS1{F0: 1, F1: 0, F2: 551}
                        _ = v29
                    } else {
                        v1 -= 50
                    }
                }
                if (-2147483648 < 635) {
                    v6 = append(v6, 42)
                } else {
                    if v27 {
                        var v30 int32 = -1
                        _ = len(v2)
                        fmt.Println(v27, v1, v27)
                        _ = v30
                    } else {
                        switch ("foo bar" + "hello") {
                        case "":
                            v27 = (1 < 4294967295)
                        case "foo bar":
                            if v27 {
                                var v31 uint8 = 255
                                v32 := GenS2{F0: (v1 | (v1 - (v1 | v1))), F1: 318}
                                var v33 float64 = float64(float32((float64(3.14159) * v5)))
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                v34 := GenS3{F0: 505, F1: -781, F2: 42}
                                var v35 int64 = 1
                                _ = v34
                                _ = v35
                            }
                        default:
                            v2[2147483647] = ("" + "\n\t")
                        }
                        _ = len(v4)
                    }
                    v36 := GenS4{F0: 42, F1: (42 != 32768), F2: -32768}
                    _ = v36
                }
                for fi37 := 0; fi37 < 1; fi37++ {
                    fmt.Println(v27)
                    _ = fi37
                }
            }
        }
    case -629:
        _ = len(v4)
    case 128:
        v5 -= float64(1.5)
    default:
        v38 := GenS1{F0: 1, F1: 42, F2: 0}
        _ = v38
    }
    var v39 int16 = 1
    var v40 uint = 4294967295
    var v41 float32 = -2.25
    _ = len(v2)
    if (v27 && v27) {
        v42 := GenS4{F0: -984, F1: (v5 >= 3.14159), F2: v39}
        for fi43 := 0; fi43 < 16; fi43++ {
            fmt.Println(fi43)
            _ = fi43
        }
        _ = v42
    } else {
        var v44 uint = 18446744073709551615
        _ = v44
    }
    if true {
        v40 += (v40 / (((v40 & v40)) | 1))
        fmt.Println(v27, v5)
        if v27 {
            var v45 uint8 = 127
            if v27 {
                if v27 {
                    _ = len(v6)
                    v41 += (v41 - v41)
                }
            }
            _ = v45
        }
    }
    v41 += (v41 * 1e10)
    v46 := GenS5{F0: (v39 ^ (^(v39))), F1: ("\n\t" + ("foo bar" + "a"))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v46
    _ = v5
    _ = v6
}
