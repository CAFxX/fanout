package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int16; F1 int16; F2 uint }

func main() {
    var v1 uint64 = 353
    var v2 int64 = -165
    v3 := make(map[uint32]string)
    var v4 uint32 = 2147483647
    v5 := make(map[string]int)
    v2 = v2
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!true) {
            v7 := []int{(fi6 - (fi6 ^ -2147483648)), 32768, (fi6 % ((42) | 1))}
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint = 18446744073709551615
                v9 = v9
                switch (((("\n\t" + ("hello" + "\n\t")) + "\n\t") + ("\n\t" + "a")) + "foo bar") {
                case "":
                    v4 = v4
                default:
                    v9 += (v9 << ((65535) & 63))
                }
                v10 := GenS1{F0: (v1 & (^(v1))), F1: 100}
                fmt.Println(v4, v2)
                fmt.Println(fi8, v2)
                _ = fi8
                _ = v10
                _ = v9
            }
            switch "\n\t" {
            case "a":
                var v11 uint32 = v4
                _ = v11
            case "":
                switch "" {
                case "hello":
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v4 = 201
                        v4 = 2
                        _ = fi12
                    }
                case "":
                    var v13 uint32 = 2
                    _ = v13
                case "foo bar":
                    switch 42 {
                    case 1:
                        var v14 uint8 = 0
                        _ = v14
                    case 0:
                        switch ("" + "") {
                        case "":
                            v15 := GenS1{F0: v1, F1: -47}
                            _ = v15
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                v16 := GenS2{F0: 42}
                                _ = v16
                            case "hello":
                                var v17 uint32 = v4
                                _ = v17
                            default:
                                v2 = (-(v2))
                            }
                        case "a":
                            var v18 string = ""
                            _ = v18
                        }
                    }
                default:
                    v19 := GenS2{F0: 2}
                    _ = v19
                }
            default:
                var v20 float32 = 3.14159
                _ = v20
            }
            _ = v7
        }
        _ = fi6
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(v2, v2, v2)
        _ = fi21
    }
    v22 := GenS3{F0: 4294967295}
    v23 := GenS3{F0: 2}
    fmt.Println(v4, v4)
    fmt.Println(v4)
    v24 := GenS3{F0: 0}
    v25 := GenS1{F0: 18446744073709551615, F1: 0}
    v1 -= (v1 % ((v1) | 1))
    var v26 int16 = 0
    var v27 int64 = (^(v2))
    v4 -= v4
    for fi28 := 0; fi28 < 12; fi28++ {
        v29 := GenS4{F0: -1, F1: (v26 ^ 42), F2: 0}
        fmt.Println(v4)
        _ = fi28
        _ = v29
    }
    switch "\n\t" {
    case "hello":
        v1 += 4294967295
    case "a":
        v4 = v4
    case "":
        v3[42] = "foo bar"
    }
    v26 -= -1
    switch "\n\t" {
    case "hello":
        switch 4294967295 {
        case -421:
            v30 := []int{-925, -9223372036854775808, 128, 4294967295}
            _ = v30
        case 0:
            var v31 bool = true
            _ = v31
        default:
            switch 4294967295 {
            case 42:
                var v32 float64 = (1.5 * (float64(0.0) + (-2.25 / (1e10 / 1.5))))
                _ = v32
            }
        }
    case "":
        fmt.Println(v26, v2, v4)
    case "a":
        v33 := GenS2{F0: 255}
        _ = v33
    default:
        fmt.Println(v4, v26, v4)
    }
    v34 := GenS1{F0: v1, F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
