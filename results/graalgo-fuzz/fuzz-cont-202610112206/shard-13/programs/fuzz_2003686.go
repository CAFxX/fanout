package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 int8; F2 string }
type GenS3 struct { F0 uint8; F1 int32; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 string }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    if false {
        p0 = (^(p0))
    } else {
        var v7 int32 = -1
        var v8 uint32 = p0
        var v9 bool = true
        _ = v7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 1, F1: true, F2: 32767}
    _ = p0
    _ = v10
    return 727
}

func main() {
    v1 := []int{0, -872, 131, 32767}
    v2 := make(map[uint32]int)
    v3 := []uint32{42, 2147483648, 0}
    var v4 int32 = -2147483648
    v5 := make(map[int]int)
    var v6 int8 = 127
    var v11 uint16 = 590
    v11 = genFn0(4294967295)
    for fi12 := 0; fi12 < 8; fi12++ {
        switch "foo bar" {
        case "\n\t":
            v13 := GenS1{F0: 4294967295, F1: (1e10 == ((1.5 - 0.1) + 3.14159)), F2: (v11 / ((0) | 1))}
            _ = v13
        default:
            v14 := GenS1{F0: 625, F1: (1.5 > (float32(((3.14159 * (-2.25 - 3.14159)) / 1.5)) * 1.5)), F2: v11}
            _ = v14
        }
        fmt.Println(fi12, fi12)
        v6 = v6
        _ = fi12
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint8 = 35
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: -1, F1: (v6 | v6), F2: "foo bar"}
    fmt.Println(v11)
    if true {
        switch 1 {
        case -369:
            v18 := GenS3{F0: 127, F1: 565, F2: -1}
            _ = v18
        case 9223372036854775807:
            v19 := GenS4{F0: -9223372036854775808}
            _ = v19
        default:
            delete(v5, 4294967295)
        }
        delete(v5, 128)
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS3{F0: 62, F1: v4, F2: 9223372036854775807}
        _ = fi20
        _ = v21
    }
    switch 2147483647 {
    case 0:
        var v22 uint64 = 2
        _ = v22
    case 4294967295:
        var v23 int8 = (v6 * 100)
        _ = v23
    default:
        v6 -= v6
    }
    v5[4294967295] = -6
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, -826)
    v6 = v6
    if (true && (false && false)) {
        v6 = -90
        switch (("hello" + ("" + "foo bar")) + "\n\t") {
        case "hello":
            switch -616 {
            case 32768:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 127:
                v4 = (v4 * v4)
            default:
                v6 = (v6 ^ 75)
            }
        case "":
            for fi24 := 0; fi24 < 9; fi24++ {
                v11 += genFn0(623)
                if false {
                    v25 := GenS1{F0: 112, F1: (false || true), F2: (v11 ^ v11)}
                    var v26 float32 = 3.14159
                    _ = v25
                    _ = v26
                }
                if true {
                    if ((-204 > (v4 & (v4 | (v4 / ((-1) | 1))))) || true) {
                        for fi27 := 0; fi27 < 13; fi27++ {
                            switch (fi27 % ((fi24) | 1)) {
                            case -9223372036854775808:
                                if false {
                                    var v28 int8 = int8(v4)
                                    if false {
                                        v28 = v28
                                        var v29 uint = 355
                                        _ = v29
                                    } else {
                                        v30 := GenS5{F0: 0}
                                        v4 += 1
                                        _ = len(v5)
                                        _ = v30
                                    }
                                    _ = v28
                                }
                            default:
                                var v31 float64 = 1.5
                                _ = v31
                            }
                            v32 := GenS3{F0: 223, F1: v4, F2: 900}
                            fmt.Println(fi24, v6)
                            _ = fi27
                            _ = v32
                        }
                        v33 := GenS6{F0: 0, F1: 2147483647}
                        if len(v1) > 0 {
                            v1[0] = -2147483648
                        }
                        _ = v33
                    } else {
                        var v34 string = "foo bar"
                        v11 += v11
                        var v35 string = "a"
                        _ = v34
                        _ = v35
                    }
                    v36 := GenS7{F0: "a"}
                    v1 = append(v1, fi24)
                    _ = v36
                } else {
                    v5[fi24] = 1
                    v37 := GenS4{F0: 42}
                    var v38 uint = 292
                    _ = v37
                    _ = v38
                }
                _ = fi24
            }
        default:
            var v39 bool = (65535 > 4294967295)
            _ = v39
        }
    } else {
        v11 += genFn0(572)
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    v40 := GenS4{F0: 9223372036854775807}
    var v41 int8 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
