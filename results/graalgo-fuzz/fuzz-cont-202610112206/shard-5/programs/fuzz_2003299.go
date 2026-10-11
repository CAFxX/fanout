package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint64; F2 string }
type GenS4 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int64 }

func main() {
    v1 := []string{}
    var v2 uint = 0
    v3 := []int{128, -289, 9223372036854775807}
    var v4 uint32 = 483
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: (v4 >> ((v2) & 63)), F1: 32767, F2: ((v4 >= (^(v4))) || ("" >= "hello"))}
    fmt.Println(v2)
    v3 = append(v3, 42)
    v4 = (v4 + (v4 ^ v4))
    v7 := v5[0]
    v8 := GenS2{F0: 2, F1: v4}
    v1 = append(v1, v7)
    v4 = (v4 - v4)
    if true {
        v9 := GenS2{F0: 2, F1: uint32(v2)}
        _ = v9
    } else {
        if true {
            v5[(v4 + (v4 % ((505) | 1)))] = v7
            if true {
                v2 -= v2
            }
            v10 := GenS2{F0: 53, F1: 0}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(fi11, v2, v4)
        if true {
            for fi12 := 0; fi12 < 1; fi12++ {
                if true {
                    fmt.Println(fi11, fi11, v2)
                    var v13 string = ""
                    _ = v13
                }
                _ = fi12
            }
            var v14 float64 = float64((((3.14159 + 1.5) + 3.14159) - (1e10 - (1.5 / ((1e10 * ((1e10 - (0.0 + 0.1)) * -2.25)) - (1.5 + 1.5))))))
            v15 := GenS1{F0: (^(v4)), F1: 42, F2: false}
            _ = v14
            _ = v15
        } else {
            if (false || ((1e10 <= ((1.5 * (float32(1.5) - 1.5)) * (0.0 * (1e10 / 1e10)))) && (978 > -353))) {
                if true {
                    v16 := v5[v4]
                    var v17 string = "\n\t"
                    _ = v16
                    _ = v17
                }
            }
            for fi18 := 0; fi18 < 8; fi18++ {
                fmt.Println(v7, fi18)
                v19 := GenS3{F0: 808, F1: 461, F2: ""}
                if false {
                    v7 = ((v7 + "a") + v7)
                    switch fi18 {
                    case 127:
                        v5[v4] = v7
                        fmt.Println(v7)
                        v20 := v5[(v4 % ((4294967295) | 1))]
                        _ = v20
                    case 32768:
                        var v21 int = (^(fi18))
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
        }
        _ = fi11
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            _ = len(v3)
            var v24 int32 = 2147483647
            var v25 int = (fi22 + fi22)
            _ = fi23
            _ = v24
            _ = v25
        }
        switch fi22 {
        case 4294967295:
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 float64 = (1.5 + -2.25)
                fmt.Println(fi22)
                _ = fi26
                _ = v27
            }
        default:
            fmt.Println(v4, v4)
        }
        v7 = "hello"
        _ = fi22
    }
    v28 := GenS2{F0: 128, F1: (v4 * v4)}
    v2 -= (v2 | (^(v2)))
    for fi29 := 0; fi29 < 7; fi29++ {
        switch (^(fi29)) {
        case -2147483648:
            var v30 int8 = -1
            _ = v30
        case -9223372036854775808:
            v31 := GenS3{F0: 18, F1: 42, F2: v7}
            _ = v31
        }
        _ = fi29
    }
    switch v7 {
    case "\n\t":
        for fi32 := 0; fi32 < 18; fi32++ {
            for fi33 := 0; fi33 < 18; fi33++ {
                v34 := GenS3{F0: (^(v2)), F1: 42, F2: v7}
                if true {
                    if (((42 < 2) || false) || false) {
                        fmt.Println(v7, fi33, v2)
                        if (0.1 != -2.25) {
                            var v35 string = v7
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v36 string = v7
                            v7 = v7
                            _ = v35
                            _ = v36
                        }
                    }
                }
                switch (fi32 >> ((1) & 63)) {
                case -2147483648:
                    v2 -= (^(v2))
                case 4294967295:
                    if ((v4 == v4) || true) {
                        switch v7 {
                        case "\n\t":
                            v37 := GenS4{F0: -32768, F1: 65535, F2: 12}
                            var v38 uint8 = 42
                            v7 = "hello"
                            v2 = (v2 >> ((v2) & 63))
                            _ = v37
                            _ = v38
                        default:
                            var v39 int64 = 0
                            _ = v39
                        }
                    } else {
                        var v40 string = ("" + "hello")
                        var v41 int64 = -9223372036854775808
                        v4 = 4294967295
                        fmt.Println(fi33)
                        _ = v40
                        _ = v41
                    }
                default:
                    if (255 >= 59) {
                        switch v7 {
                        case "\n\t":
                            fmt.Println(fi32, v2)
                        default:
                            var v42 string = v7
                            _ = v42
                        }
                    }
                }
                var v43 bool = (-401 <= -1)
                v44 := GenS5{F0: -9223372036854775808}
                v5[0] = v7
                fmt.Println(v7)
                _ = fi33
                _ = v34
                _ = v43
                _ = v44
            }
            _ = fi32
        }
    default:
        switch 1 {
        case 128:
            v45 := GenS5{F0: 1}
            _ = v45
        case 4294967295:
            v4 -= (v4 * 4294967295)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
