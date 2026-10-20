package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 uint32; F2 string }

func main() {
    var v1 float64 = (((0.1 + 1.5) - float64(1.5)) * 3.14159)
    v2 := make(map[string]uint32)
    var v3 uint16 = 221
    if false {
        switch "" {
        case "":
            if (!((!(0 == -79)) && true)) {
                var v4 uint = 42
                var v5 int8 = 49
                var v6 int64 = 55
                _ = v4
                _ = v5
                _ = v6
            } else {
                switch "a" {
                case "":
                    var v7 string = "\n\t"
                    _ = v7
                default:
                    for fi8 := 0; fi8 < 7; fi8++ {
                        var v9 int = fi8
                        if false {
                            var v10 int16 = 984
                            _ = v10
                        }
                        _ = fi8
                        _ = v9
                    }
                }
            }
        case "foo bar":
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(v3, fi11, v1)
                if false {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v2[""] = 842
                        _ = fi12
                    }
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 uint64 = 2
                        v15 := []int{(^(fi13)), 32768, fi11, fi11}
                        switch "" {
                        case "":
                            v3 += 692
                        case "hello":
                            var v16 float64 = ((v1 / 1e10) / (v1 * 1.5))
                            _ = v16
                        case "\n\t":
                            switch "\n\t" {
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    v17 := GenS1{F0: 128}
                                    _ = v17
                                default:
                                    v18 := GenS2{F0: 0, F1: 2147483648, F2: "a"}
                                    _ = v18
                                }
                            default:
                                var v19 int8 = -107
                                _ = v19
                            }
                        default:
                            v1 = 1.5
                        }
                        switch (fi13 - -816) {
                        case 127:
                            v2[("" + "\n\t")] = 2147483647
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v20 float32 = 1.5
                    _ = v20
                }
                _ = fi11
            }
        default:
            _ = len(v2)
        }
        var v21 uint64 = 0
        _ = v21
    } else {
        for fi22 := 0; fi22 < 15; fi22++ {
            if false {
                var v23 float32 = ((-2.25 - (((1.5 / -2.25) / (((1e10 / 3.14159) + (float32((3.14159 + (v1 + 0.1))) + 0.0)) - (0.0 - 1e10))) / 1e10)) / (float32((0.1 / float64(((-2.25 * (0.0 - 1.5)) / -2.25)))) * 1e10))
                for fi24 := 0; fi24 < 10; fi24++ {
                    v25 := v2[""]
                    _ = fi24
                    _ = v25
                }
                _ = v23
            }
            _ = fi22
        }
    }
    v26 := GenS2{F0: -2147483648, F1: 42, F2: "foo bar"}
    v27 := v2[""]
    v3 -= 255
    var v28 uint32 = 4294967295
    for fi29 := 0; fi29 < 6; fi29++ {
        v30 := []int{}
        switch "foo bar" {
        case "hello":
            delete(v2, "a")
        case "a":
            var v31 bool = (18446744073709551615 <= 42)
            _ = v31
        case "\n\t":
            v2["\n\t"] = 2147483647
        }
        _ = fi29
        _ = v30
    }
    var v32 int64 = 602
    var v33 int8 = 0
    switch -390 {
    case -9223372036854775808:
        for fi34 := 0; fi34 < 20; fi34++ {
            switch "\n\t" {
            case "":
                v1 = (v1 / ((v1 + (v1 + v1)) / ((v1 + ((1.5 + ((0.0 / (1.5 * (-2.25 * float64(1.5)))) / v1)) - v1)) + v1)))
            case "a":
                var v35 uint16 = 490
                _ = v35
            case "foo bar":
                v1 = (3.14159 / ((0.0 + 1.5) - float64(3.14159)))
            }
            var v36 uint64 = 2
            switch (fi34 - fi34) {
            case -9223372036854775808:
                v37 := GenS2{F0: int32(v33), F1: 1, F2: "foo bar"}
                _ = v37
            case 32768:
                var v38 uint64 = v36
                _ = v38
            case 1:
                var v39 int = -737
                _ = v39
            default:
                v3 += 285
            }
            _ = fi34
            _ = v36
        }
    case 32767:
        var v40 uint16 = (v3 & 0)
        _ = v40
    default:
        v28 = (v28 >> ((255) & 63))
    }
    var v41 float32 = (((float32(v1) + 3.14159) - ((float32(v1) / (0.0 + -2.25)) / 1.0)) / float32(v1))
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v41
}
