package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 string; F2 uint32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float32 = ((0.0 - ((3.14159 + float32(-2.25)) / -2.25)) * 1.5)
    var v2 uint16 = 0
    v3 := []int{-72}
    var v4 int = 327
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 bool = false
        v2 += v2
        fmt.Println(v1, fi5, v4)
        _ = fi5
        _ = v6
    }
    var v7 float64 = 0.1
    if true {
        if (!true) {
            v8 := GenS1{F0: (v4 ^ (v4 & v4))}
            for fi9 := 0; fi9 < 3; fi9++ {
                if (v2 > v2) {
                    v10 := GenS1{F0: 128}
                    _ = v10
                }
                _ = fi9
            }
            if true {
                var v11 uint16 = v2
                switch "hello" {
                case "foo bar":
                    fmt.Println(v7)
                case "\n\t":
                    for fi12 := 0; fi12 < 16; fi12++ {
                        if true {
                            v13 := GenS1{F0: (fi12 * (-(v4)))}
                            fmt.Println(v11)
                            _ = v13
                        } else {
                            if ((!(255 > 13)) || true) {
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    fmt.Println(v2, v4)
                                    _ = fi14
                                }
                                v15 := GenS2{F0: (((v7 / (((v7 + float64(v1)) * 1e10) + (0.0 + float64(v1)))) + v7) < float64(v1)), F1: "hello", F2: 65535}
                                var v16 int32 = -491
                                _ = v15
                                _ = v16
                            } else {
                                var v17 int32 = 737
                                v7 = v7
                                fmt.Println(v4, v7, v11)
                                _ = v17
                            }
                            v18 := GenS1{F0: v4}
                            _ = v18
                        }
                        _ = fi12
                    }
                case "a":
                    v19 := GenS1{F0: v4}
                    _ = v19
                }
                _ = v11
            } else {
                var v20 int32 = int32(v2)
                _ = v20
            }
            _ = v8
        }
        fmt.Println(v4, v7, v2)
    } else {
        fmt.Println(v1, v4)
        _ = len(v3)
    }
    switch 127 {
    case -672:
        fmt.Println(v1, v7, v7)
    case 1:
        if true {
            v4 -= (v4 ^ v4)
        }
    default:
        var v21 uint32 = 0
        _ = v21
    }
    _ = len(v3)
    var v22 string = "a"
    switch "\n\t" {
    case "foo bar":
        if (false || (-1 > 42)) {
            fmt.Println(v1)
            v4 += 4294967295
        } else {
            v23 := GenS3{F0: ((!((v4 > (v4 % ((v4) | 1))) || (!false))) && (!false))}
            _ = v23
        }
    default:
        var v24 int8 = -18
        _ = v24
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
}
