package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint8, p2 int16) uint32 {
    v5 := []int{9223372036854775807}
    p2 -= (p2 % ((-1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 93
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{42, 245}
    var v3 int32 = 1
    v4 := make(map[int]bool)
    var v6 int8 = 0
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 42}
        _ = fi7
        _ = v8
    }
    if ("a" < "hello") {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 int64 = 1
            for fi11 := 0; fi11 < 18; fi11++ {
                switch "hello" {
                case "a":
                    switch fi11 {
                    case -145:
                        v6 = v6
                    case 9223372036854775807:
                        if true {
                            var v12 uint = 312
                            _ = v12
                        }
                    default:
                        var v13 uint16 = 32767
                        _ = v13
                    }
                    v14 := GenS1{F0: v3}
                    _ = v14
                case "\n\t":
                    v3 += int32(v10)
                }
                switch "a" {
                case "a":
                    var v15 uint = 1
                    _ = v15
                case "\n\t":
                    if (!(true || true)) {
                        if (539 < 1) {
                            v16 := GenS1{F0: v3}
                            v1 = (1.5 + (1.5 / v1))
                            switch fi9 {
                            case 9223372036854775807:
                                var v17 int32 = (v3 | v3)
                                _ = v17
                            case 32767:
                                v10 = v10
                            default:
                                var v18 float32 = -2.25
                                _ = v18
                            }
                            _ = v16
                        } else {
                            v19 := v4[fi9]
                            _ = v19
                        }
                    } else {
                        var v20 uint64 = 2
                        v20 = (v20 % (((^(v20))) | 1))
                        _ = v20
                    }
                default:
                    v10 -= (v10 % ((v10) | 1))
                }
                _ = fi11
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                var v22 uint32 = 0
                var v23 int = 4294967295
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi9
            _ = v10
        }
        _ = len(v2)
    }
    var v24 int8 = -21
    v25 := GenS1{F0: v3}
    var v26 int16 = -559
    fmt.Println(v1, v3)
    fmt.Println(v3, v1)
    for fi27 := 0; fi27 < 13; fi27++ {
        if true {
            v2 = append(v2, genFn0(350, 227, (v26 / ((-32768) | 1))))
            var v28 float32 = 0.1
            _ = v28
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 uint8 = 130
                v31 := GenS1{F0: v3}
                _ = fi29
                _ = v30
                _ = v31
            }
            switch (fi27 << ((1) & 63)) {
            case 2147483647:
                fmt.Println(v24)
            default:
                v32 := GenS1{F0: (v3 ^ 588)}
                _ = v32
            }
            switch "\n\t" {
            case "\n\t":
                var v33 int16 = v26
                _ = v33
            case "":
                fmt.Println(v26, v24, v26)
            case "a":
                v34 := GenS1{F0: v3}
                _ = v34
            default:
                var v35 int32 = (v3 / (((v3 << ((255) & 63))) | 1))
                _ = v35
            }
        }
        v24 = (^(v24))
        v36 := GenS1{F0: -1}
        _ = fi27
        _ = v36
    }
    v1 -= v1
    switch -1 {
    case 127:
        var v37 uint8 = 0
        _ = v37
    case 4294967295:
        v6 += -113
    }
    v6 = (v24 / ((117) | 1))
    var v38 bool = (923 < 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v4
    _ = v6
}
