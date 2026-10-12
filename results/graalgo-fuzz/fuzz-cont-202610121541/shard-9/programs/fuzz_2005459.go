package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{2, 454, 4294967295, 2}
    v3 := make(map[int]bool)
    var v4 int32 = 0
    var v5 uint = 0
    v5 = 780
    var v6 float64 = -2.25
    v4 = (v4 * 933)
    v5 -= v5
    if false {
        if (!(false || false)) {
            if true {
                v6 -= (float64(1e10) + v6)
            } else {
                if true {
                    fmt.Println(v5)
                    v7 := GenS1{F0: 9223372036854775807}
                    _ = v7
                }
                for fi8 := 0; fi8 < 15; fi8++ {
                    v9 := GenS1{F0: 9223372036854775807}
                    for fi10 := 0; fi10 < 16; fi10++ {
                        v3[2147483647] = false
                        v5 = v5
                        var v11 uint64 = 18446744073709551615
                        _ = fi10
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
            }
        }
    } else {
        fmt.Println(v6)
    }
    v6 = (((v6 / v6) + v6) + 0.1)
    var v12 uint64 = 514
    v5 = (^(v5))
    if false {
        if false {
            var v13 uint64 = v12
            for fi14 := 0; fi14 < 18; fi14++ {
                v5 -= v5
                v15 := GenS2{F0: -9223372036854775808, F1: 0}
                v4 = (^(v4))
                _ = fi14
                _ = v15
            }
            if false {
                switch (("a" + "\n\t") + "hello") {
                case "a":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v12 = (^(v12))
                        v2 = append(v2, 1)
                        _ = fi16
                    }
                case "hello":
                    var v17 int64 = -1
                    _ = v17
                default:
                    var v18 uint16 = 600
                    _ = v18
                }
            } else {
                switch 0 {
                case 32768:
                    fmt.Println(v6)
                case 128:
                    v12 += (v12 | (^(v13)))
                }
            }
            _ = v13
        }
        if true {
            switch "foo bar" {
            case "foo bar":
                v12 -= v12
            case "hello":
                v12 += v12
            default:
                v4 += v4
            }
        }
        v19 := GenS1{F0: int(v5)}
        _ = v19
    }
    fmt.Println(v6)
    switch "hello" {
    case "a":
        if (!false) {
            for fi20 := 0; fi20 < 12; fi20++ {
                fmt.Println(v5)
                switch -9223372036854775808 {
                case 32768:
                    var v21 uint32 = 2
                    _ = v21
                case 0:
                    v12 += v12
                case -1:
                    v22 := GenS3{F0: 0, F1: 42}
                    _ = v22
                default:
                    _ = len(v2)
                }
                for fi23 := 0; fi23 < 1; fi23++ {
                    v4 -= (v4 / (((v4 | 478)) | 1))
                    _ = fi23
                }
                _ = fi20
            }
        }
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v25 := v3[fi24]
        for fi26 := 0; fi26 < 18; fi26++ {
            v25 = true
            _ = fi26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
