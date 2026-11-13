package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8; F2 int32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 uint8; F1 int8; F2 uint64 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 int16 = 0
    var v5 int = 1
    v6 := make(map[int]bool)
    switch v5 {
    case 9223372036854775807:
        v7 := GenS1{F0: false}
        _ = v7
    case 2147483647:
        v4 -= v4
    case -361:
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float32 = float32((0.0 / (1e10 + 3.14159)))
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS2{F0: 4294967295, F1: 569, F2: v1}
    if true {
        var v11 int64 = 0
        v1 += (^(v1))
        var v12 bool = true
        _ = v11
        _ = v12
    } else {
        switch v5 {
        case -1:
            if ((v5 == v5) || true) {
                v13 := GenS3{F0: 465, F1: int8(v5), F2: -1}
                if true {
                    v14 := GenS4{F0: -73, F1: v1}
                    if false {
                        var v15 int16 = (v4 + v4)
                        _ = v15
                    }
                    v1 = v1
                    _ = v14
                }
                _ = v13
            } else {
                v5 = (v5 ^ 2147483647)
                var v16 float64 = ((-2.25 * 0.0) / ((float64(0.0) - 0.1) * ((((float64(0.0) - (1.5 - 0.1)) - ((0.1 - (-2.25 / -2.25)) + (1.5 / (1.5 / 1.5)))) * (3.14159 / 0.1)) * 1.5)))
                v17 := GenS2{F0: 0, F1: 0, F2: v1}
                _ = v16
                _ = v17
            }
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint16 = (v1 >> ((570) & 63))
        v20 := GenS4{F0: v5, F1: 2}
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 string = "foo bar"
    v22 := GenS5{F0: 128, F1: int8(v5), F2: 42}
    fmt.Println(v5, v1)
    if (2 == 255) {
        var v23 int8 = 1
        if true {
            fmt.Println(v21, v21)
        } else {
            v24 := []int{(^(v5)), v5}
            _ = v24
        }
        switch (("foo bar" + v21) + v21) {
        case "hello":
            for fi25 := 0; fi25 < 13; fi25++ {
                fmt.Println(v23, fi25)
                var v26 int16 = v4
                _ = fi25
                _ = v26
            }
        case "":
            switch -1 {
            case 42:
                v21 = "\n\t"
            case -2147483648:
                var v27 int32 = 42
                _ = v27
            default:
                switch -670 {
                case 42:
                    fmt.Println(v1)
                case 32768:
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 uint64 = 1
                        if true {
                            v30 := []int{(v5 % (((fi28 >> ((18446744073709551615) & 63))) | 1)), 4294967295}
                            _ = v30
                        } else {
                            v29 -= (v29 & (v29 / ((v29) | 1)))
                        }
                        for fi31 := 0; fi31 < 18; fi31++ {
                            v32 := []int{}
                            if (1.5 >= 3.14159) {
                                v23 = (v23 + 0)
                            }
                            v4 -= v4
                            _ = fi31
                            _ = v32
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            }
        default:
            fmt.Println(v21)
        }
        _ = v23
    } else {
        v4 = (-(v4))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
