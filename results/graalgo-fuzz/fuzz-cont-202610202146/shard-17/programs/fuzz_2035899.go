package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int64 }

func main() {
    var v1 uint8 = 42
    v2 := []int{-9223372036854775808, -2147483648, -2147483648}
    var v3 float64 = 1e10
    v4 := []string{}
    v5 := make(map[uint32]bool)
    v1 += (v1 >> ((2) & 63))
    var v6 int8 = 112
    for fi7 := 0; fi7 < 20; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                if false {
                    fmt.Println(fi8, v1, v1)
                } else {
                    v6 += v6
                    v6 -= 42
                    fmt.Println(fi7, v6, v1)
                }
                switch fi7 {
                case 32767:
                    v3 -= v3
                    v1 += v1
                case 0:
                    v3 -= float64(-2.25)
                case 714:
                    fmt.Println(v3, fi8)
                    var v9 int64 = -1
                    v6 -= v6
                    var v10 int16 = -1
                    v10 = v10
                    switch "hello" {
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            _ = len(v5)
                        case "foo bar":
                            v2 = append(v2, int(v1))
                        case "":
                            if (v1 < 0) {
                                var v11 int64 = v9
                                _ = v11
                            } else {
                                var v12 int8 = (^(v6))
                                switch (fi8 - -423) {
                                case -607:
                                    switch -1 {
                                    case -456:
                                        v12 += (v12 + v12)
                                    case -485:
                                        v13 := GenS1{F0: 1}
                                        _ = v13
                                    }
                                case 664:
                                    var v14 uint64 = 4
                                    _ = v14
                                case 9223372036854775807:
                                    v6 += v6
                                }
                                switch fi8 {
                                case 701:
                                    var v15 uint = 255
                                    _ = v15
                                case 42:
                                    v16 := GenS2{F0: "hello", F1: v9}
                                    v10 = 0
                                    switch fi8 {
                                    case 731:
                                        v3 = ((v3 + v3) - 0.1)
                                    case 42:
                                        var v17 int8 = (v6 << ((42) & 63))
                                        _ = v17
                                    case 127:
                                        switch (fi7 ^ (fi8 + (fi8 * (fi8 | (fi7 * fi8))))) {
                                        case 4294967295:
                                            v3 += v3
                                        }
                                    default:
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                    }
                                    _ = v16
                                }
                                fmt.Println(fi8, v12, v1)
                                v9 -= v9
                                _ = v12
                            }
                        }
                    default:
                        v3 -= ((v3 - 0.1) * ((v3 / ((1.5 + (((0.1 * (float64(3.14159) / 1.0)) / v3) / v3)) + v3)) - 0.1))
                    }
                    switch fi8 {
                    case 128:
                        var v18 string = "a"
                        _ = v18
                    case 2147483647:
                        var v19 string = "a"
                        _ = v19
                    case 127:
                        if (float32(v3) != 0.0) {
                            var v20 uint16 = 351
                            var v21 string = "\n\t"
                            _ = v20
                            _ = v21
                        }
                        v5[uint32(v1)] = true
                        v22 := GenS2{F0: "\n\t", F1: (v9 ^ (v9 >> ((2) & 63)))}
                        v23 := GenS2{F0: ("" + ("\n\t" + "hello")), F1: v9}
                        var v24 string = (("hello" + "foo bar") + (("a" + "foo bar") + ("foo bar" + "\n\t")))
                        if true {
                            var v25 uint8 = v1
                            var v26 float32 = (1e10 / (3.14159 - (0.0 * float32((float64(0.0) - ((v3 - 3.14159) + ((-2.25 * 3.14159) * float64((1e10 / -2.25)))))))))
                            _ = v25
                            _ = v26
                        }
                        _ = v22
                        _ = v23
                        _ = v24
                    default:
                        var v27 int32 = 73
                        _ = v27
                    }
                    _ = v10
                    _ = v9
                }
            }
            v3 += float64(-2.25)
            switch fi7 {
            case 32768:
                var v28 int8 = 127
                _ = v28
            case 128:
                if (0 >= 303) {
                    fmt.Println(v3)
                }
            }
            _ = fi8
        }
        for fi29 := 0; fi29 < 20; fi29++ {
            fmt.Println(fi7)
            switch fi7 {
            case -2147483648:
                v3 = 1e10
            case 9223372036854775807:
                v6 -= (v6 << ((2) & 63))
            }
            _ = fi29
        }
        v6 -= (v6 << ((255) & 63))
        _ = fi7
    }
    v30 := GenS1{F0: 2}
    v5[70] = (false || false)
    _ = len(v5)
    var v31 int16 = -32768
    fmt.Println(v1, v6, v1)
    var v32 int16 = (v31 * (v31 / ((v31) | 1)))
    for fi33 := 0; fi33 < 5; fi33++ {
        for fi34 := 0; fi34 < 11; fi34++ {
            var v35 int8 = 30
            var v36 uint64 = 1
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    v37 := GenS2{F0: "\n\t", F1: 42}
    fmt.Println(v3, v1)
    v38 := GenS2{F0: ("\n\t" + ""), F1: -1}
    if (true || false) {
        fmt.Println(v3, v6, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
