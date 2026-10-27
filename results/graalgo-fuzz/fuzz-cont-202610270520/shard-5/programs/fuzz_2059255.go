package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 bool; F1 int8; F2 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 float32 = 0.0
    var v4 float64 = -2.25
    v5 := make(map[string]bool)
    v6 := GenS1{F0: 0, F1: -128}
    if true {
        if (-1 > -9223372036854775808) {
            for fi7 := 0; fi7 < 7; fi7++ {
                switch fi7 {
                case 42:
                    v8 := []int{}
                    _ = v8
                case 114:
                    if (359 == 0) {
                        v9 := []int{(fi7 << ((543) & 63)), 32768, (fi7 / (((fi7 << ((18446744073709551615) & 63))) | 1))}
                        v9 = append(v9, fi7)
                        fmt.Println(v3)
                        _ = v9
                    }
                case 127:
                    var v10 int8 = -128
                    _ = v10
                }
                v11 := []int{fi7, fi7, fi7}
                _ = fi7
                _ = v11
            }
            v12 := []int{-332, 32768, 0}
            _ = v12
        }
        if false {
            var v13 uint = 0
            _ = v13
        }
        v14 := []int{}
        _ = v14
    } else {
        v15 := []int{4294967295, 324, -9223372036854775808, -887}
        delete(v1, 2147483648)
        _ = v15
    }
    v16 := []int{9223372036854775807}
    fmt.Println(v3, v4)
    v17 := GenS2{F0: true, F1: -1, F2: "\n\t"}
    switch 848 {
    case 1:
        for fi18 := 0; fi18 < 6; fi18++ {
            switch (((("\n\t" + "a") + ("a" + "foo bar")) + "a") + ("a" + "\n\t")) {
            case "a":
                switch (fi18 & (fi18 ^ fi18)) {
                case -1:
                    if (42 != 255) {
                        for fi19 := 0; fi19 < 7; fi19++ {
                            var v20 uint64 = 0
                            if len(v16) > 0 {
                                v16[0] = 42
                            }
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(v3)
                            default:
                                switch "foo bar" {
                                case "foo bar":
                                    v20 = 551
                                case "\n\t":
                                    v20 -= (v20 % ((v20) | 1))
                                    var v21 float64 = v4
                                    _ = v21
                                }
                            }
                            v22 := GenS3{F0: "foo bar"}
                            _ = fi19
                            _ = v20
                            _ = v22
                        }
                        var v23 float64 = v4
                        var v24 int8 = 1
                        _ = v23
                        _ = v24
                    } else {
                        if true {
                            v25 := GenS2{F0: true, F1: -128, F2: ""}
                            var v26 uint = 2
                            v4 -= v4
                            _ = v25
                            _ = v26
                        } else {
                            v4 -= v4
                            v27 := GenS4{F0: 0, F1: 1}
                            v3 = v3
                            _ = v27
                        }
                        fmt.Println(v3)
                    }
                case 2147483647:
                    v3 += v3
                case 164:
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                }
            case "foo bar":
                if true {
                    v3 = (float32((0.0 * 1.5)) - (v3 - v3))
                    v28 := GenS1{F0: (fi18 * 9223372036854775807), F1: 52}
                    _ = v28
                }
            case "\n\t":
                var v29 int16 = -125
                _ = v29
            }
            _ = fi18
        }
    default:
        v30 := GenS1{F0: 32767, F1: 127}
        _ = v30
    }
    v3 += v3
    fmt.Println(v4, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
