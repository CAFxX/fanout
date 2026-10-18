package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 int32 = 42
    var v5 float32 = (-2.25 / (float32(-2.25) + (0.1 + 1e10)))
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        v4 += (v4 - 0)
        for fi7 := 0; fi7 < 9; fi7++ {
            v5 = 3.14159
            v8 := GenS1{F0: 0}
            var v9 uint64 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    if (0 < 18446744073709551615) {
        v10 := v3[2147483648]
        v5 += (v5 + 1.5)
        var v11 uint64 = 914
        _ = v10
        _ = v11
    }
    v5 = float32(-2.25)
    var v12 float32 = ((0.1 - v5) * (3.14159 * -2.25))
    fmt.Println(v12, v5)
    switch -350 {
    case 128:
        switch ("\n\t" + "\n\t") {
        case "":
            for fi13 := 0; fi13 < 14; fi13++ {
                v1 = int64(v4)
                _ = fi13
            }
        }
    case 2147483647:
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 19; fi15++ {
                v16 := GenS2{F0: 0}
                _ = fi15
                _ = v16
            }
            var v17 bool = false
            if v17 {
                v18 := GenS3{F0: 0}
                v19 := []int{192, (fi14 * (fi14 - fi14)), 32767}
                switch "foo bar" {
                case "foo bar":
                    switch (fi14 + (fi14 ^ (fi14 << ((4294967295) & 63)))) {
                    case 1:
                        for fi20 := 0; fi20 < 5; fi20++ {
                            switch 4294967295 {
                            case 9223372036854775807:
                                v17 = v17
                            case -990:
                                var v21 float64 = float64(0.0)
                                _ = v21
                            }
                            switch "a" {
                            case "foo bar":
                                v1 += v1
                            case "hello":
                                v22 := GenS2{F0: 0}
                                _ = v22
                            }
                            fmt.Println(v5)
                            var v23 float64 = (0.0 - (-2.25 / 3.14159))
                            _ = fi20
                            _ = v23
                        }
                    case 0:
                        switch "" {
                        case "":
                            switch "hello" {
                            case "\n\t":
                                if v17 {
                                    var v24 uint64 = 2
                                    if (!false) {
                                        if v17 {
                                            v12 = v5
                                            var v25 float64 = (1.5 * (1.5 + 0.0))
                                            for fi26 := 0; fi26 < 9; fi26++ {
                                                if (-86 > -128) {
                                                    v27 := GenS2{F0: 4294967295}
                                                    switch fi26 {
                                                    case 510:
                                                        fmt.Println(v4)
                                                        fmt.Println(v24)
                                                        v12 += v5
                                                    }
                                                    delete(v2, 849)
                                                    if len(v19) > 0 {
                                                        v19[0] = (fi26 & fi14)
                                                    }
                                                    var v28 int32 = v4
                                                    _ = v27
                                                    _ = v28
                                                }
                                                _ = fi26
                                            }
                                            _ = v25
                                        } else {
                                            switch (fi14 ^ fi14) {
                                            case 0:
                                                fmt.Println(fi14, fi14)
                                            case 42:
                                                delete(v3, 42)
                                            case 841:
                                                var v29 int16 = -974
                                                _ = v29
                                            default:
                                                v30 := v3[691]
                                                _ = v30
                                            }
                                            var v31 int64 = (v1 >> ((859) & 63))
                                            switch ("a" + "") {
                                            case "foo bar":
                                                v2[65535] = (fi14 >> ((18446744073709551615) & 63))
                                            default:
                                                v1 -= (v1 - (v1 % ((0) | 1)))
                                            }
                                            _ = v31
                                        }
                                        v32 := v2[455]
                                        for fi33 := 0; fi33 < 20; fi33++ {
                                            fmt.Println(v1, v1)
                                            _ = fi33
                                        }
                                        _ = v32
                                    }
                                    var v34 float64 = 0.0
                                    _ = v24
                                    _ = v34
                                } else {
                                    v35 := GenS4{F0: 0, F1: fi14}
                                    var v36 float32 = v5
                                    var v37 uint32 = 65535
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                            case "foo bar":
                                var v38 uint32 = 94
                                _ = v38
                            }
                        }
                    case 603:
                        v17 = (1 < 0)
                    }
                }
                _ = v18
                _ = v19
            }
            _ = fi14
            _ = v17
        }
    case 796:
        fmt.Println(v1, v12)
    }
    fmt.Println(v1)
    v5 -= (0.1 / -2.25)
    var v39 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
