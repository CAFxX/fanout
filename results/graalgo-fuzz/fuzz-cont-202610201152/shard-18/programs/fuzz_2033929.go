package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 uint8 = 128
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    v6 := GenS1{F0: 2, F1: -1}
    var v7 int16 = 1000
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := GenS2{F0: ("foo bar" + "foo bar"), F1: 1}
            var v11 float64 = -2.25
            var v12 float32 = 0.1
            var v13 bool = (!false)
            var v14 float64 = float64(v12)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
            _ = v13
            _ = v14
        }
        v7 -= 1000
        v15 := []int{(fi8 & fi8), (fi8 - (fi8 / ((fi8) | 1)))}
        _ = fi8
        _ = v15
    }
    _ = len(v2)
    v16 := GenS2{F0: "foo bar", F1: -1}
    v7 = (v7 | 280)
    switch (("foo bar" + "\n\t") + "") {
    case "\n\t":
        var v17 string = "hello"
        _ = v17
    }
    if (100 <= 1) {
        v18 := v4[-9223372036854775808]
        for fi19 := 0; fi19 < 8; fi19++ {
            if (!true) {
                var v20 uint64 = 42
                _ = v20
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 uint32 = 65535
                v7 = int16(fi21)
                switch (fi19 & fi21) {
                case 32768:
                    if false {
                        v23 := []int{(fi19 & -2147483648), fi19}
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v3)
                        case "foo bar":
                            switch fi19 {
                            case 0:
                                v22 -= v18
                            case 2147483647:
                                if (false || (18446744073709551615 < 0)) {
                                    v24 := v2[fi19]
                                    _ = v24
                                } else {
                                    v25 := v4[9223372036854775807]
                                    var v26 int64 = 1
                                    _ = v25
                                    _ = v26
                                }
                            default:
                                if (!false) {
                                    _ = len(v23)
                                    v27 := v5[(fi21 / ((fi19) | 1))]
                                    _ = len(v23)
                                    _ = v27
                                }
                            }
                        default:
                            v28 := GenS2{F0: "a", F1: 42}
                            fmt.Println(v18, v22)
                            _ = v28
                        }
                        _ = v23
                    }
                case 32767:
                    v29 := []int{(fi21 >> ((584) & 63)), -2147483648, (fi19 / ((-2147483648) | 1)), (fi21 + (fi19 / ((-689) | 1)))}
                    _ = v29
                case -854:
                    v30 := GenS2{F0: ((("a" + "hello") + ("foo bar" + ("hello" + "foo bar"))) + "\n\t"), F1: -31}
                    _ = v30
                }
                _ = fi21
                _ = v22
            }
            _ = fi19
        }
        switch -614 {
        case 32768:
            if true {
                fmt.Println(v7)
                fmt.Println(v7, v7)
            }
        }
        _ = v18
    }
    var v31 int64 = -9223372036854775808
    delete(v5, -1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
