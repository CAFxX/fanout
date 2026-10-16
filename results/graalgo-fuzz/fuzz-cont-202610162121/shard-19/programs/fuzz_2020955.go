package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }

func main() {
    var v1 float64 = 3.14159
    v2 := []string{("foo bar" + ("a" + "a")), "\n\t", ((("\n\t" + "foo bar") + "hello") + ("a" + "\n\t"))}
    var v3 uint64 = 0
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 4, F1: 100}
    var v6 uint8 = 127
    v7 := GenS2{F0: 255, F1: 774, F2: 0}
    v6 -= (v6 * 0)
    _ = len(v4)
    for fi8 := 0; fi8 < 19; fi8++ {
        v9 := GenS2{F0: 167, F1: -1, F2: 1}
        for fi10 := 0; fi10 < 14; fi10++ {
            switch "a" {
            case "hello":
                if true {
                    if false {
                        if true {
                            var v11 uint16 = 1
                            fmt.Println(v11, fi8, v1)
                            var v12 uint = 65535
                            _ = v11
                            _ = v12
                        } else {
                            fmt.Println(v6, fi8, fi8)
                        }
                    } else {
                        var v13 uint = 4294967295
                        _ = v13
                    }
                    if ((^(v3)) <= (v3 | 94)) {
                        v1 = 1.5
                        v14 := GenS2{F0: 590, F1: -9223372036854775808, F2: 0}
                        _ = v14
                    } else {
                        if ("foo bar" > ("" + "hello")) {
                            var v15 int = 32767
                            var v16 uint16 = 32767
                            v17 := GenS3{F0: -2147483648, F1: (v6 * 85), F2: 97}
                            _ = v15
                            _ = v16
                            _ = v17
                        } else {
                            v6 = (v6 / ((v6) | 1))
                            fmt.Println(v3)
                        }
                    }
                } else {
                    if (1 < 2) {
                        v1 = v1
                        v3 = v3
                    } else {
                        if true {
                            var v18 int64 = 0
                            v19 := GenS2{F0: 42, F1: v18, F2: 255}
                            _ = v18
                            _ = v19
                        } else {
                            v20 := v4[fi10]
                            _ = v20
                        }
                        v2 = append(v2, "hello")
                    }
                }
            case "a":
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v6)
                }
            case "foo bar":
                v21 := GenS1{F0: 100, F1: -128}
                _ = v21
            }
            v22 := GenS4{F0: 800, F1: 42, F2: 255}
            fmt.Println(v3, v6, fi10)
            _ = fi10
            _ = v22
        }
        v6 += (v6 * 212)
        _ = fi8
        _ = v9
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        v6 += v6
        v3 = 0
        v6 = (v6 << ((42) & 63))
        _ = fi23
    }
    v1 = 1.5
    v1 += v1
    v24 := GenS3{F0: 805, F1: (v6 & 128), F2: 65535}
    var v25 int32 = 2147483647
    switch "" {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        for fi26 := 0; fi26 < 3; fi26++ {
            v27 := GenS3{F0: -2147483648, F1: v6, F2: 4294967295}
            v4[(fi26 - (fi26 ^ fi26))] = true
            _ = fi26
            _ = v27
        }
    }
    var v28 int8 = -44
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v1, v25)
        var v30 uint8 = (v6 + (v6 % ((v6) | 1)))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi29
        _ = v30
    }
    v6 -= (^(v6))
    v1 -= 0.1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
