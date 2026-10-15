package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int8; F2 int8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint16 = 61
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    fmt.Println(v1)
    v1 = (^(v1))
    fmt.Println(v1)
    switch 128 {
    case 4294967295:
        var v4 uint32 = 2
        _ = v4
    case -635:
        v5 := []int{42, 127}
        _ = v5
    case -100:
        if (false || true) {
            switch -835 {
            case 127:
                v6 := GenS1{F0: ((true || true) && false)}
                _ = v6
            case 2147483647:
                v7 := []int{}
                _ = v7
            case -656:
                var v8 int64 = 9223372036854775807
                _ = v8
            }
        } else {
            var v9 bool = false
            switch 386 {
            case 995:
                var v10 int8 = -1
                _ = v10
            case 9223372036854775807:
                for fi11 := 0; fi11 < 12; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS2{F0: -15, F1: 0, F2: -62}
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
            default:
                v3[473] = (!(65535 <= v1))
            }
            _ = v9
        }
    default:
        var v14 uint = 7
        _ = v14
    }
    switch -2147483648 {
    case 42:
        if true {
            v1 -= (v1 - (v1 | (v1 << ((491) & 63))))
            switch "foo bar" {
            case "\n\t":
                var v15 float64 = (((1.5 - float64(0.1)) / 1.0) + float64((-2.25 - ((((0.0 + ((3.14159 - (1.5 * 3.14159)) / -2.25)) / -2.25) * 3.14159) * ((0.0 * 1e10) + (1e10 - 3.14159))))))
                _ = v15
            case "":
                v16 := GenS3{F0: "a", F1: -1}
                _ = v16
            }
        }
    case 32767:
        var v17 uint = 18446744073709551615
        _ = v17
    default:
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := []int{(fi18 % ((fi18) | 1)), -2147483648, fi18, (fi18 + fi19)}
                switch "" {
                case "a":
                    var v21 int64 = 0
                    _ = v21
                default:
                    if (551 < (fi19 * fi18)) {
                        if true {
                            v1 -= 255
                        } else {
                            var v22 string = ""
                            v23 := v2[(v22 + (v22 + v22))]
                            var v24 int16 = 1000
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        v1 += 795
                    }
                }
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    if false {
        switch 2147483647 {
        case 2147483647:
            v25 := []int{-1, 2147483647, 1}
            _ = v25
        default:
            if true {
                var v26 int = 410
                fmt.Println(v26, v1)
                v1 = v1
                _ = v26
            }
        }
        var v27 float64 = 0.0
        if (42 < 255) {
            fmt.Println(v27)
            delete(v2, "foo bar")
        }
        _ = v27
    } else {
        if true {
            v28 := GenS3{F0: "", F1: 2147483647}
            for fi29 := 0; fi29 < 5; fi29++ {
                switch fi29 {
                case -1:
                    for fi30 := 0; fi30 < 9; fi30++ {
                        var v31 int16 = -166
                        var v32 int32 = -2147483648
                        v1 += 447
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                case -939:
                    var v33 int = 49
                    _ = v33
                case -908:
                    if (42 > 2) {
                        fmt.Println(v1, v1)
                        var v34 bool = (873 != -2147483648)
                        v1 = v1
                        _ = v34
                    }
                }
                _ = fi29
            }
            _ = v28
        }
        v1 = (v1 / (((v1 << ((4294967295) & 63))) | 1))
        v35 := GenS4{F0: 18446744073709551615}
        _ = v35
    }
    v36 := GenS4{F0: 1}
    var v37 float32 = 1e10
    v37 = 3.14159
    v37 -= (v37 * float32(float64(((v37 * 3.14159) - 1.5))))
    v38 := []int{127, -9223372036854775808, 127}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
