package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int8 = 100
    var v2 int = 32768
    v3 := make(map[string]int)
    var v4 uint16 = 917
    v5 := make(map[uint32]bool)
    fmt.Println(v2, v2, v2)
    var v6 int16 = int16(v2)
    switch -9223372036854775808 {
    case 42:
        fmt.Println(v1, v2, v2)
    case 32768:
        var v7 int16 = (v6 / ((v6) | 1))
        _ = v7
    case 0:
        v8 := v3[("a" + "a")]
        _ = v8
    default:
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                v2 += (-(fi9))
                switch "hello" {
                case "":
                    v4 = (v4 - (^(v4)))
                case "\n\t":
                    if false {
                        var v11 float32 = 1e10
                        v12 := v5[797]
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            var v13 int16 = 32767
                            var v14 uint32 = 2147483647
                            _ = v13
                            _ = v14
                        } else {
                            var v15 float64 = ((0.0 - 0.1) * 0.0)
                            _ = v15
                        }
                        if false {
                            var v16 float64 = -2.25
                            _ = v16
                        }
                        v2 += fi10
                    }
                }
                _ = fi10
            }
            v1 = 127
            _ = fi9
        }
    }
    fmt.Println(v1)
    v6 = v6
    switch "a" {
    case "\n\t":
        var v17 int16 = 42
        _ = v17
    default:
        v18 := GenS1{F0: 42, F1: 9223372036854775807}
        _ = v18
    }
    v19 := []int{v2, v2, v2}
    if (!(127 >= 128)) {
        if false {
            v4 = 424
            if false {
                for fi20 := 0; fi20 < 3; fi20++ {
                    v21 := GenS2{F0: 42, F1: (!(!(100 != (v1 & 42))))}
                    if len(v19) > 0 {
                        v19[0] = v2
                    }
                    _ = fi20
                    _ = v21
                }
                v22 := GenS2{F0: 42, F1: true}
                if len(v19) > 0 {
                    v19[0] = v2
                }
                _ = v22
            }
            if false {
                v1 = 0
                fmt.Println(v4, v6, v6)
            }
        } else {
            var v23 float32 = (float32(-2.25) + 1e10)
            _ = v23
        }
        v2 -= -897
    }
    v6 -= 0
    v24 := v5[554]
    if (!v24) {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS3{F0: 1}
            var v27 uint32 = 2147483647
            switch "a" {
            case "hello":
                fmt.Println(v24)
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v24 = ((^(v4)) >= 577)
    }
    var v28 uint64 = 132
    var v29 uint16 = (v4 % ((v4) | 1))
    v6 = v6
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
