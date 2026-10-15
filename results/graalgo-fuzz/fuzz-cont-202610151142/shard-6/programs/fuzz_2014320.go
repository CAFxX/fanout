package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 int32; F1 bool; F2 string }
type GenS3 struct { F0 uint; F1 int16; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int }
func genFn0() bool {
    var v5 int16 = -32768
    if true {
        switch -121 {
        case 32768:
            v5 = v5
        default:
            v6 := GenS1{F0: 65535, F1: "hello", F2: 9223372036854775807}
            _ = v6
        }
        v5 -= (v5 * (-(v5)))
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        switch "foo bar" {
        case "\n\t":
            if (!(-9223372036854775808 < 249)) {
                switch (fi7 ^ fi7) {
                case -354:
                    if false {
                        var v8 float64 = float64(0.0)
                        v9 := []int{}
                        _ = v8
                        _ = v9
                    } else {
                        fmt.Println(v5)
                        v10 := []int{(fi7 % (((fi7 ^ (fi7 / ((fi7) | 1)))) | 1))}
                        _ = v10
                    }
                default:
                    v11 := []int{128, (fi7 % (((-(fi7))) | 1))}
                    _ = v11
                }
                v5 -= v5
                v12 := GenS2{F0: -1, F1: (555 > 4294967295), F2: ""}
                _ = v12
            } else {
                v5 = (v5 << ((492) & 63))
                var v13 uint16 = 369
                if (!false) {
                    var v14 int16 = v5
                    v14 -= (v5 - 1)
                    _ = v14
                }
                _ = v13
            }
        case "a":
            for fi15 := 0; fi15 < 2; fi15++ {
                v5 += (v5 ^ 815)
                _ = fi15
            }
        }
        v16 := GenS3{F0: 2, F1: 1000, F2: (fi7 % ((fi7) | 1))}
        _ = fi7
        _ = v16
    }
    _ = v5
    return (!(0.0 != ((0.1 / 3.14159) + 0.0)))
}

func main() {
    var v1 uint = 978
    v2 := []uint32{}
    var v3 uint = (^(v1))
    v4 := make(map[int]int)
    v17 := GenS1{F0: v3, F1: "\n\t", F2: 9223372036854775807}
    v1 += v3
    v18 := GenS4{F0: 42}
    var v19 bool = (false && true)
    v20 := GenS3{F0: 18446744073709551615, F1: 32767, F2: 2147483647}
    delete(v4, -1)
    if v19 {
        if true {
            for fi21 := 0; fi21 < 16; fi21++ {
                v4[(fi21 << (((v1 % (((v3 * (v1 % ((18446744073709551615) | 1)))) | 1))) & 63))] = fi21
                var v22 uint32 = 65535
                switch fi21 {
                case 128:
                    var v23 float32 = 0.1
                    _ = v23
                default:
                    v2 = append(v2, v22)
                }
                _ = fi21
                _ = v22
            }
            fmt.Println(v3)
            v19 = (!(((1 < 42) || ((8 == 32767) && v19)) && (!v19)))
        }
    } else {
        var v24 float64 = (0.1 * ((3.14159 / -2.25) / 1.5))
        for fi25 := 0; fi25 < 13; fi25++ {
            for fi26 := 0; fi26 < 9; fi26++ {
                v24 += v24
                var v27 int64 = -285
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
        _ = v24
    }
    if (-9223372036854775808 < 454) {
        switch 32768 {
        case -26:
            v1 += v3
        case -713:
            v28 := GenS5{F0: -2147483648, F1: 32767}
            _ = v28
        }
        switch "" {
        case "hello":
            v1 += 2
        case "foo bar":
            if false {
                fmt.Println(v19)
            } else {
                var v29 int64 = 9223372036854775807
                _ = v29
            }
        case "":
            switch "\n\t" {
            case "":
                for fi30 := 0; fi30 < 6; fi30++ {
                    v3 -= (v3 * v3)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v31 int16 = 1000
                    _ = fi30
                    _ = v31
                }
            case "\n\t":
                v3 += v3
            case "foo bar":
                for fi32 := 0; fi32 < 5; fi32++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = fi32
                }
            }
        }
        switch "" {
        case "foo bar":
            for fi33 := 0; fi33 < 14; fi33++ {
                var v34 float32 = (1e10 / 1e10)
                var v35 uint8 = 128
                _ = fi33
                _ = v34
                _ = v35
            }
        }
    }
    var v36 uint32 = 65535
    var v37 int8 = 100
    for fi38 := 0; fi38 < 19; fi38++ {
        v39 := GenS5{F0: 2147483647, F1: fi38}
        fmt.Println(v36, v3)
        _ = fi38
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
