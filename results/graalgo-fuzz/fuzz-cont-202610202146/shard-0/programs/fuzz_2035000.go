package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 uint16; F1 uint32; F2 bool }
type GenS9 struct { F0 uint; F1 bool }

func main() {
    var v1 string = (("hello" + "\n\t") + "")
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 int8 = -1
    v5 := []int{32768, 9223372036854775807}
    switch 2147483647 {
    case 32768:
        v4 -= 27
    default:
        v6 := GenS1{F0: -1, F1: 127, F2: 0}
        _ = v6
    }
    v4 = (v4 * v4)
    var v7 int64 = -990
    if (!(!((v7 * (v7 - -1)) <= v7))) {
        var v8 uint16 = 42
        _ = len(v5)
        _ = v8
    }
    if false {
        if (4294967295 != 2147483648) {
            fmt.Println(v7, v1, v7)
            var v9 uint16 = 780
            _ = v9
        } else {
            var v10 int64 = (v7 << ((980) & 63))
            var v11 uint = 18446744073709551615
            v12 := GenS1{F0: v7, F1: 152, F2: uint64(v11)}
            _ = v10
            _ = v11
            _ = v12
        }
        switch v1 {
        case "a":
            _ = len(v2)
        case "\n\t":
            var v13 string = (v1 + (v1 + v1))
            _ = v13
        default:
            v1 = (v1 + v1)
        }
    } else {
        v1 = "foo bar"
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        if (false && ((!false) && (!true))) {
            for fi15 := 0; fi15 < 10; fi15++ {
                v4 = v4
                v16 := v3[65535]
                if true {
                    var v17 bool = (!((42 > 1) || (439 == 42)))
                    fmt.Println(fi14, v7, v7)
                    switch v16 {
                    case "foo bar":
                        var v18 int16 = 347
                        _ = v18
                    case "\n\t":
                        v19 := GenS1{F0: (v7 ^ v7), F1: 191, F2: 18446744073709551615}
                        v7 = (v7 / ((1) | 1))
                        _ = v19
                    case "hello":
                        v20 := GenS2{F0: (v7 * 686), F1: 255}
                        _ = v20
                    default:
                        _ = len(v3)
                    }
                    _ = v17
                } else {
                    switch "hello" {
                    case "a":
                        if true {
                            var v21 int32 = 0
                            v22 := GenS1{F0: v7, F1: 128, F2: 4294967295}
                            _ = v21
                            _ = v22
                        }
                    case "":
                        if false {
                            fmt.Println(v7, v4)
                            var v23 bool = false
                            if (v16 <= v1) {
                                if ("hello" < (v16 + v1)) {
                                    delete(v3, 2147483648)
                                } else {
                                    var v24 float32 = float32(float64(((float32((0.0 * 1.5)) + -2.25) / 1.0)))
                                    switch v1 {
                                    case "":
                                        if true {
                                            v7 = -366
                                        }
                                    default:
                                        var v25 int = (-(fi15))
                                        v25 = fi14
                                        switch (fi14 & 9223372036854775807) {
                                        case -1:
                                            v26 := GenS1{F0: v7, F1: 255, F2: 42}
                                            _ = v26
                                        case 32768:
                                            if (149 != 42) {
                                                if false {
                                                    _ = len(v5)
                                                    v4 = v4
                                                    v27 := GenS3{F0: (("hello" + "") + ("a" + "\n\t"))}
                                                    var v28 int64 = v7
                                                    _ = v27
                                                    _ = v28
                                                } else {
                                                    v29 := GenS2{F0: int64(v4), F1: 140}
                                                    fmt.Println(v24, v16)
                                                    var v30 float64 = 1.5
                                                    _ = v29
                                                    _ = v30
                                                }
                                                var v31 bool = (float64(v24) < (0.0 / -2.25))
                                                var v32 int = fi14
                                                _ = v31
                                                _ = v32
                                            }
                                        }
                                        _ = v25
                                    }
                                    _ = v24
                                }
                                v4 -= v4
                                if len(v5) > 0 {
                                    v5[0] = fi15
                                }
                            }
                            fmt.Println(fi15, v1)
                            _ = v23
                        }
                    case "foo bar":
                        if false {
                            if (255 > 150) {
                                v33 := v2[v1]
                                fmt.Println(v16)
                                _ = v33
                            } else {
                                switch int(v4) {
                                case -1:
                                    var v34 uint16 = 65535
                                    _ = v34
                                case 148:
                                    switch ((v16 + v1) + v16) {
                                    case "\n\t":
                                        var v35 float64 = ((1e10 - 0.0) / ((((1.5 * ((1e10 * (1.5 + 3.14159)) / (float64((3.14159 - 1e10)) + 3.14159))) - 0.1) / 0.1) / float64(-2.25)))
                                        _ = v35
                                    case "":
                                        v36 := GenS4{F0: -32768, F1: 2147483648, F2: 347}
                                        _ = v36
                                    default:
                                        v5 = append(v5, (fi14 * fi15))
                                    }
                                case 32767:
                                    var v37 uint8 = 2
                                    if (1 <= 4294967295) {
                                        v38 := GenS5{F0: (fi15 << ((255) & 63)), F1: v7}
                                        if (!false) {
                                            var v39 bool = false
                                            v39 = true
                                            _ = v39
                                        } else {
                                            v40 := GenS1{F0: v7, F1: v37, F2: 560}
                                            _ = v40
                                        }
                                        _ = v38
                                    }
                                    _ = len(v3)
                                    v41 := GenS6{F0: 39}
                                    _ = v37
                                    _ = v41
                                }
                            }
                            v3[65535] = "foo bar"
                            v42 := GenS7{F0: 945, F1: 0}
                            _ = v42
                        } else {
                            switch (-(fi14)) {
                            case -183:
                                var v43 uint = 65535
                                v44 := GenS2{F0: (v7 >> (((v43 & (v43 + v43))) & 63)), F1: 0}
                                _ = v43
                                _ = v44
                            default:
                                v16 = "\n\t"
                            }
                            v7 = 42
                        }
                        switch v16 {
                        case "a":
                            v16 = v16
                        case "\n\t":
                            v3[231] = "a"
                        }
                        v4 -= 31
                    }
                    var v45 uint64 = 804
                    _ = v45
                }
                if len(v5) > 0 {
                    v5[0] = fi15
                }
                _ = fi15
                _ = v16
            }
            var v46 int = 637
            _ = v46
        }
        switch "" {
        case "":
            v7 += v7
        default:
            if true {
                v47 := GenS8{F0: 0, F1: 2, F2: true}
                _ = v47
            }
        }
        _ = fi14
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        v4 = -128
        for fi49 := 0; fi49 < 16; fi49++ {
            v7 += v7
            v5 = append(v5, 928)
            _ = fi49
        }
        _ = fi48
    }
    v7 += v7
    var v50 bool = true
    v51 := GenS9{F0: 277, F1: true}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v7
}
