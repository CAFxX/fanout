package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16; F1 int8; F2 bool }
func genFn0() int32 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 100}
        _ = fi5
        _ = v6
    }
    v7 := GenS2{F0: 515, F1: 1, F2: 32768}
    var v8 uint16 = 742
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []uint32{981, 2147483647, 34, 2}
    v2 := make(map[int]int)
    var v3 uint32 = 2
    var v4 bool = (127 <= -93)
    var v9 float32 = (float32(float64(1.5)) - 0.1)
    if false {
        var v10 float64 = (0.0 - (0.1 / 1.5))
        if len(v1) > 0 {
            v1[0] = v3
        }
        _ = v10
    }
    v3 -= 42
    if (-2.25 == 3.14159) {
        if ((93 != -80) || v4) {
            v3 -= v3
        } else {
            var v11 uint8 = 1
            _ = v11
        }
        var v12 uint16 = 65535
        if (!v4) {
            var v13 uint64 = 42
            _ = v13
        } else {
            fmt.Println(v3, v12, v3)
            var v14 uint8 = 1
            _ = v14
        }
        _ = v12
    } else {
        fmt.Println(v9, v4, v9)
        v15 := GenS3{F0: (v4 && v4), F1: 18446744073709551615, F2: 42}
        v2[356] = 128
        _ = v15
    }
    var v16 uint16 = 1
    var v17 uint16 = (v16 * v16)
    v18 := GenS3{F0: true, F1: 18446744073709551615, F2: (v17 * v16)}
    v4 = (!v4)
    var v19 int32 = 42
    if (52 == 127) {
        for fi20 := 0; fi20 < 13; fi20++ {
            if len(v1) > 0 {
                v1[0] = v3
            }
            v21 := GenS2{F0: (v17 >> ((4294967295) & 63)), F1: (v16 | v17), F2: fi20}
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 19; fi22++ {
            if true {
                if (!((v3 ^ v3) != (v3 | 298))) {
                    for fi23 := 0; fi23 < 19; fi23++ {
                        var v24 int8 = 127
                        switch "foo bar" {
                        case "hello":
                            v4 = true
                        case "a":
                            var v25 uint32 = (v3 + v3)
                            _ = v25
                        default:
                            fmt.Println(v9, v24, v16)
                        }
                        _ = fi23
                        _ = v24
                    }
                    switch (fi22 + 2147483647) {
                    case 9223372036854775807:
                        for fi26 := 0; fi26 < 8; fi26++ {
                            v4 = true
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v27 := GenS4{F0: 232, F1: 773}
                            v16 += v17
                            delete(v2, (fi26 >> ((1) & 63)))
                            _ = fi26
                            _ = v27
                        }
                    default:
                        switch (fi22 + (-(fi22))) {
                        case 518:
                            for fi28 := 0; fi28 < 14; fi28++ {
                                fmt.Println(fi22)
                                v3 += (^(v3))
                                switch 0 {
                                case -52:
                                    v9 = v9
                                default:
                                    var v29 bool = v4
                                    _ = v29
                                }
                                _ = fi28
                            }
                        }
                    }
                    v30 := GenS5{F0: -32768, F1: 14, F2: v4}
                    _ = v30
                }
            } else {
                if (!(0.1 != float64((v9 + v9)))) {
                    switch (fi22 << ((1) & 63)) {
                    case 32767:
                        v3 += (v3 ^ uint32(v17))
                    default:
                        var v31 int64 = 0
                        _ = v31
                    }
                    v32 := GenS2{F0: 32767, F1: (v16 + (v16 >> ((2) & 63))), F2: 0}
                    _ = v32
                }
            }
            _ = fi22
        }
        var v33 int8 = -128
        _ = v33
    }
    delete(v2, 9223372036854775807)
    var v34 uint = 0
    var v35 int8 = 127
    fmt.Println(v16)
    if (("\n\t" + "foo bar") > "a") {
        var v36 int = 42
        if false {
            var v37 uint64 = 274
            switch ("" + (((("" + "a") + "hello") + "a") + "a")) {
            case "":
                delete(v2, 164)
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                v38 := GenS1{F0: v35}
                _ = v38
            }
            var v39 uint64 = v37
            _ = v37
            _ = v39
        }
        _ = v36
    } else {
        switch "foo bar" {
        case "hello":
            var v40 int64 = -635
            _ = v40
        case "":
            v41 := GenS1{F0: -1}
            _ = v41
        case "a":
            var v42 int8 = int8(v19)
            _ = v42
        }
    }
    v3 = uint32(v16)
    fmt.Println(v3, v17)
    switch 667 {
    case 967:
        if false {
            var v43 float64 = (0.0 * 0.1)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v43
        } else {
            var v44 uint32 = uint32(v16)
            v19 -= (v19 - (v19 / (((v19 / ((v19) | 1))) | 1)))
            _ = v44
        }
    case 4294967295:
        v45 := GenS1{F0: v35}
        _ = v45
    default:
        var v46 bool = true
        _ = v46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v9
}
