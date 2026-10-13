package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int64; F2 string }
type GenS7 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    var v5 int = -2147483648
    switch v4 {
    case "":
        v4 = "hello"
    default:
        v5 = v5
    }
    v4 = ("" + ("\n\t" + ""))
    v6 := GenS1{F0: 0, F1: 2}
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 = "\n\t"
        for fi8 := 0; fi8 < 10; fi8++ {
            if true {
                v9 := GenS2{F0: 1, F1: (1 == 0)}
                var v10 int64 = 1
                _ = v10
                _ = v9
            } else {
                v11 := GenS3{F0: 935, F1: 2}
                if true {
                    v4 = (v4 + (v4 + "a"))
                    if true {
                        v5 += (fi8 * (fi8 ^ 2147483647))
                        switch -9223372036854775808 {
                        case 0:
                            v12 := GenS4{F0: 0, F1: ((v5 >> ((18446744073709551615) & 63)) < 655)}
                            _ = v12
                        case 9223372036854775807:
                            var v13 float64 = (3.14159 / 3.14159)
                            _ = v13
                        default:
                            v5 += (v5 << ((1) & 63))
                        }
                    } else {
                        var v14 float32 = float32((((float64(3.14159) * ((0.0 - (1.5 + (((-2.25 + 1.5) + (((1.5 * ((float64(3.14159) * float64(0.1)) - 3.14159)) + (3.14159 - (-2.25 * (0.1 * float64((3.14159 - -2.25)))))) - (3.14159 * -2.25))) + 1.5))) - float64(3.14159))) + 0.0) * ((0.0 - ((1.5 - 3.14159) / ((float64(((1e10 * (0.0 * 0.1)) + 3.14159)) / (3.14159 / 3.14159)) + (1e10 + 1e10)))) - (1e10 / -2.25))))
                        if (-2147483648 != 0) {
                            v15 := []int{(fi7 << ((0) & 63))}
                            v14 = float32(float64((v14 / v14)))
                            v16 := GenS4{F0: -1, F1: true}
                            v5 -= -869
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                    v17 := []int{v5, 1}
                    _ = v17
                }
                _ = v11
            }
            var v18 int16 = 337
            _ = fi8
            _ = v18
        }
        _ = fi7
    }
    var v19 uint32 = 4294967295
    v19 += v19
    var v20 uint32 = v19
    _ = len(v2)
    if (((1.5 + -2.25) / ((3.14159 - 1.5) - 1.5)) < float32((3.14159 * -2.25))) {
        v2[v20] = v20
    }
    v19 -= (v20 & 0)
    fmt.Println(v5, v5, v5)
    var v21 string = "\n\t"
    v22 := []int{}
    for fi23 := 0; fi23 < 9; fi23++ {
        if (!((65535 <= 18446744073709551615) || (42 < 2))) {
            fmt.Println(v20, v20)
            v4 = "\n\t"
            fmt.Println(fi23)
        }
        _ = fi23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        v20 = (v20 * v20)
        switch "foo bar" {
        case "foo bar":
            v20 += (v19 << ((1) & 63))
        }
        v25 := GenS5{F0: 728}
        _ = fi24
        _ = v25
    }
    v21 = ((v4 + "\n\t") + "")
    for fi26 := 0; fi26 < 13; fi26++ {
        switch (v5 + fi26) {
        case 9223372036854775807:
            switch "" {
            case "":
                if false {
                    v27 := GenS6{F0: ((float64((0.0 / 1.5)) != float64(-2.25)) && (86 > 42)), F1: -304, F2: v4}
                    _ = v27
                }
            case "foo bar":
                fmt.Println(v4)
            case "\n\t":
                v4 = "hello"
            }
        }
        v28 := GenS7{F0: 807, F1: true}
        if (!false) {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 int64 = -851
                fmt.Println(fi29)
                _ = fi29
                _ = v30
            }
            switch (fi26 + fi26) {
            case 877:
                var v31 uint32 = 2
                _ = v31
            }
            switch (^(v5)) {
            case 432:
                v5 += 413
            case 127:
                v32 := GenS5{F0: v19}
                _ = v32
            }
        } else {
            _ = len(v2)
        }
        _ = fi26
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
