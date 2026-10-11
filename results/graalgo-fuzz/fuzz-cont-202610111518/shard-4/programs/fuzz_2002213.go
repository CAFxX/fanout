package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8; F1 int8; F2 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 1; fi4++ {
        if (2 <= 0) {
            for fi5 := 0; fi5 < 12; fi5++ {
                if false {
                    fmt.Println(fi5, fi5)
                    switch 9223372036854775807 {
                    case 42:
                        if false {
                            var v6 bool = false
                            if v6 {
                                var v7 int = 42
                                v6 = v6
                                fmt.Println(v7, v7, v7)
                                _ = v7
                            }
                            _ = v6
                        }
                    default:
                        v8 := GenS1{F0: (fi4 / ((fi4) | 1))}
                        _ = v8
                    }
                }
                _ = fi5
            }
        } else {
            var v9 float32 = ((1.5 / 0.1) / 1.5)
            v10 := GenS2{F0: 9223372036854775807, F1: 246}
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    var v11 float32 = float32((0.0 - 0.0))
    var v12 uint64 = 500
    v13 := v2[350]
    v3[0] = 2
    for fi14 := 0; fi14 < 16; fi14++ {
        v12 = (v12 - 18446744073709551615)
        fmt.Println(v13, fi14)
        _ = fi14
    }
    switch v13 {
    case "":
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS3{F0: 379, F1: 42}
            v11 = float32(0.1)
            _ = fi15
            _ = v16
        }
    case "a":
        var v17 float32 = (1e10 / (0.0 + (0.0 * v11)))
        _ = v17
    case "hello":
        if false {
            for fi18 := 0; fi18 < 5; fi18++ {
                v11 -= ((0.0 / 1.0) / 0.1)
                if true {
                    var v19 int16 = 32767
                    v20 := []int{-1, (fi18 / (((fi18 | (-(fi18)))) | 1)), fi18, (fi18 * 795)}
                    fmt.Println(v11, v11)
                    _ = v19
                    _ = v20
                } else {
                    delete(v1, (fi18 * fi18))
                }
                _ = fi18
            }
        } else {
            switch v13 {
            case "":
                if false {
                    v12 = 272
                }
            case "foo bar":
                v12 -= 18446744073709551615
            default:
                if (42 <= 128) {
                    v21 := []int{317, -439, -9223372036854775808, 128}
                    _ = v21
                }
            }
        }
    default:
        v22 := GenS4{F0: uint16(v12)}
        _ = v22
    }
    v23 := GenS1{F0: -816}
    v1[-591] = 769
    for fi24 := 0; fi24 < 2; fi24++ {
        delete(v2, 1)
        var v25 float32 = v11
        var v26 int64 = 9223372036854775807
        _ = fi24
        _ = v25
        _ = v26
    }
    switch v13 {
    case "":
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 uint8 = 128
            v29 := GenS5{F0: (fi27 ^ (fi27 / ((fi27) | 1))), F1: v28}
            _ = fi27
            _ = v28
            _ = v29
        }
    case "hello":
        v13 = v13
    default:
        switch ("a" + (v13 + "hello")) {
        case "hello":
            v30 := GenS2{F0: -9223372036854775808, F1: 1}
            _ = v30
        case "a":
            fmt.Println(v12)
        default:
            v31 := GenS6{F0: true}
            _ = v31
        }
    }
    v32 := []int{0, 112}
    _ = len(v2)
    v33 := GenS1{F0: 4294967295}
    switch 4294967295 {
    case 32767:
        fmt.Println(v11, v11)
    case 4294967295:
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
    case -1:
        if len(v32) > 0 {
            v32[0] = 128
        }
    }
    var v34 string = (v13 + v13)
    v35 := GenS7{F0: 42, F1: -1, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
