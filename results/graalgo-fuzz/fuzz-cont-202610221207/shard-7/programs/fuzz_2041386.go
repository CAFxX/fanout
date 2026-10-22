package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]string)
    v3 := make(map[string]string)
    var v4 int16 = -32768
    _ = len(v2)
    v5 := GenS1{F0: 42, F1: -357}
    var v6 bool = (!true)
    var v7 uint64 = 626
    v8 := []int{128}
    switch "foo bar" {
    case "hello":
        var v9 uint8 = 128
        _ = v9
    case "":
        v10 := GenS1{F0: 2, F1: (v4 / ((v4) | 1))}
        _ = v10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        _ = len(v2)
        v7 = v7
        _ = fi11
    }
    var v12 uint64 = v7
    if false {
        var v13 int64 = -9223372036854775808
        v7 -= 435
        _ = v13
    } else {
        v12 = (v7 + v12)
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(fi14, v4, v12)
            for fi15 := 0; fi15 < 7; fi15++ {
                _ = len(v1)
                var v16 bool = v6
                v17 := GenS1{F0: 2, F1: (v4 % ((32767) | 1))}
                v12 -= (^(v12))
                _ = fi15
                _ = v16
                _ = v17
            }
            if v6 {
                v3["a"] = "\n\t"
                switch ("\n\t" + "hello") {
                case "hello":
                    switch (fi14 | (fi14 - fi14)) {
                    case 823:
                        if v6 {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v19 := GenS2{F0: true}
                                _ = fi18
                                _ = v19
                            }
                            switch "hello" {
                            case "\n\t":
                                _ = len(v8)
                            case "foo bar":
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    var v21 uint64 = 42
                                    _ = fi20
                                    _ = v21
                                }
                            default:
                                if ((1.5 < 3.14159) || true) {
                                    for fi22 := 0; fi22 < 1; fi22++ {
                                        _ = len(v1)
                                        fmt.Println(fi14, v6, v4)
                                        if true {
                                            v23 := GenS1{F0: 2, F1: v4}
                                            v6 = v6
                                            v12 = v7
                                            _ = v23
                                        } else {
                                            v24 := GenS3{F0: 0, F1: (v4 ^ 928)}
                                            v25 := GenS3{F0: 1, F1: (v4 - v4)}
                                            var v26 float64 = 3.14159
                                            _ = v24
                                            _ = v25
                                            _ = v26
                                        }
                                        _ = fi22
                                    }
                                }
                            }
                        } else {
                            if ((3.14159 + (3.14159 + (3.14159 * 1e10))) >= float64(float32(3.14159))) {
                                fmt.Println(v12, v6)
                                v12 = (v7 + 0)
                            } else {
                                v8 = append(v8, (fi14 | fi14))
                                for fi27 := 0; fi27 < 19; fi27++ {
                                    fmt.Println(fi27, fi14)
                                    _ = fi27
                                }
                                var v28 int32 = -991
                                _ = v28
                            }
                            fmt.Println(v12, fi14)
                        }
                    case -9223372036854775808:
                        if (2 < 0) {
                            v4 = 1000
                        }
                    default:
                        fmt.Println(v6)
                    }
                }
            } else {
                v7 = (v12 & (v12 ^ (v7 >> ((255) & 63))))
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v29 int8 = -128
                _ = v29
            }
            _ = fi14
        }
        v30 := GenS4{F0: 100}
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        v12 -= (v7 & 778)
        v12 = (v7 / ((v12) | 1))
        _ = fi31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
