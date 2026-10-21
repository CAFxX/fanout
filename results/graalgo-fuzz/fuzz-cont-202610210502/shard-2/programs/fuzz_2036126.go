package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint64; F1 int16; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]bool)
    var v3 float64 = ((1e10 / 1.0) / ((1e10 * 0.0) - ((3.14159 - 0.1) + (0.0 / float64(3.14159)))))
    v4 := make(map[uint32]bool)
    v5 := []int{4294967295, 128, 128, 396}
    v3 -= (1.5 + v3)
    v3 += (v3 - 1e10)
    var v6 bool = false
    switch -44 {
    case 128:
        var v7 int = 765
        _ = v7
    case 299:
        for fi8 := 0; fi8 < 16; fi8++ {
            if (!true) {
                var v9 float32 = float32(v3)
                var v10 int32 = -2147483648
                if false {
                    if (!(!false)) {
                        var v11 string = "hello"
                        for fi12 := 0; fi12 < 14; fi12++ {
                            var v13 uint8 = 128
                            v13 += v13
                            v9 -= v9
                            v11 = (v11 + "")
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    } else {
                        switch -2147483648 {
                        case -555:
                            _ = len(v5)
                        }
                    }
                    var v14 int64 = -690
                    _ = v14
                }
                _ = v10
                _ = v9
            } else {
                for fi15 := 0; fi15 < 12; fi15++ {
                    if v6 {
                        if v6 {
                            fmt.Println(v6)
                            var v16 int16 = -1
                            if (v3 >= 0.0) {
                                v17 := GenS1{F0: 127, F1: 255}
                                _ = v17
                            }
                            _ = v16
                        } else {
                            if (v6 && (v6 || false)) {
                                v3 += v3
                                v6 = false
                            }
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v6 = v6
                        switch "\n\t" {
                        case "hello":
                            v3 = v3
                            var v18 int16 = 1
                            _ = v18
                        case "\n\t":
                            v6 = ((!(((v3 + 0.0) == v3) || (v6 || (!false)))) && (v6 || (-1 > 42)))
                        case "foo bar":
                            v19 := GenS2{F0: "a", F1: (42 != 2)}
                            _ = v19
                        default:
                            v3 += -2.25
                        }
                        var v20 int = (fi8 - (-(fi15)))
                        _ = v20
                    }
                    v3 = ((0.1 * v3) + v3)
                    _ = fi15
                }
            }
            var v21 int8 = 121
            delete(v1, "a")
            _ = fi8
            _ = v21
        }
    }
    fmt.Println(v6, v3)
    v5 = append(v5, -184)
    v3 = v3
    _ = len(v4)
    v6 = v6
    for fi22 := 0; fi22 < 14; fi22++ {
        switch (fi22 >> ((0) & 63)) {
        case 9223372036854775807:
            switch ("hello" + "\n\t") {
            case "\n\t":
                var v23 int = (fi22 % (((fi22 << ((65535) & 63))) | 1))
                _ = v23
            case "hello":
                v3 = 1.5
            default:
                v3 = (v3 / (v3 / float64((1e10 - (1e10 + -2.25)))))
            }
        case 0:
            for fi24 := 0; fi24 < 9; fi24++ {
                v6 = v6
                var v25 int8 = 127
                _ = fi24
                _ = v25
            }
        case 4294967295:
            var v26 uint8 = 128
            _ = v26
        }
        var v27 uint8 = 128
        _ = fi22
        _ = v27
    }
    v6 = (v6 && (false || (true && v6)))
    fmt.Println(v6, v6)
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 float32 = 0.0
        if ((-2.25 >= v29) || (!true)) {
            switch (((("a" + ("\n\t" + "a")) + "a") + "") + ("hello" + "hello")) {
            case "":
                v29 = (((0.0 / (float32(v3) / -2.25)) - (0.0 + v29)) + v29)
            default:
                fmt.Println(v29)
            }
            v6 = false
        }
        v29 = v29
        _ = fi28
        _ = v29
    }
    fmt.Println(v6, v3)
    switch "" {
    case "\n\t":
        v3 = v3
    case "foo bar":
        delete(v1, (("" + "\n\t") + "hello"))
    case "a":
        v30 := v4[995]
        _ = v30
    }
    if v6 {
        v31 := GenS3{F0: 869, F1: 225, F2: 0}
        _ = v31
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
