package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int32; F1 int8; F2 bool }
func genFn0(p0 int8, p1 bool) string {
    var v4 string = ("hello" + "a")
    var v5 bool = true
    v6 := []int{-80}
    switch (v4 + "\n\t") {
    case "a":
        if v5 {
            if len(v6) > 0 {
                v6[0] = 302
            }
        } else {
            var v7 float64 = ((-2.25 * (0.0 / (0.1 - 1.5))) - 3.14159)
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 int32 = -1
    v3 := make(map[uint32]string)
    v8 := GenS1{F0: 4294967295, F1: 2147483647}
    var v9 uint8 = 128
    v10 := []int{9223372036854775807, 4294967295, -9223372036854775808, -2147483648}
    switch -1 {
    case 0:
        switch "\n\t" {
        case "foo bar":
            switch "foo bar" {
            case "hello":
                fmt.Println(v9, v2)
            case "":
                v11 := GenS2{F0: (42 > v9), F1: -1}
                _ = v11
            default:
                v12 := GenS3{F0: 0, F1: 42, F2: (1000 < -32768)}
                _ = v12
            }
        case "hello":
            v9 -= v9
        case "a":
            if false {
                var v13 int32 = 1
                v14 := GenS1{F0: 1, F1: 65535}
                _ = v13
                _ = v14
            }
        }
    default:
        fmt.Println(v1, v1, v2)
    }
    switch 32767 {
    case -9223372036854775808:
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case -518:
        v2 = 1
    case -572:
        if true {
            v2 = v2
            fmt.Println(v1, v9, v2)
            var v15 int8 = 42
            _ = v15
        }
    }
    _ = len(v10)
    var v16 int16 = -32768
    switch "a" {
    case "foo bar":
        var v17 bool = false
        _ = v17
    default:
        switch -9223372036854775808 {
        case 114:
            v10 = append(v10, 628)
        default:
            for fi18 := 0; fi18 < 18; fi18++ {
                switch "hello" {
                case "hello":
                    _ = len(v3)
                case "foo bar":
                    switch "" {
                    case "a":
                        v2 += v2
                    default:
                        if (false || (3.14159 != ((1.5 * float32((3.14159 * -2.25))) / float32(1.5)))) {
                            var v19 uint8 = (v9 % ((v9) | 1))
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = v19
                        }
                    }
                case "":
                    v16 = v16
                }
                var v20 int16 = (v16 | 1)
                v2 -= 935
                _ = fi18
                _ = v20
            }
        }
    }
    var v21 float32 = 1e10
    for fi22 := 0; fi22 < 12; fi22++ {
        v21 = (1.5 / v21)
        fmt.Println(v16, v16)
        _ = fi22
    }
    v23 := GenS1{F0: 1, F1: 2147483648}
    for fi24 := 0; fi24 < 16; fi24++ {
        if true {
            _ = len(v3)
            v1 = ((v1 - ((3.14159 - v21) / v1)) / (v1 - v21))
            var v25 int16 = int16(v9)
            _ = v25
        } else {
            v26 := GenS2{F0: (2147483648 != 44), F1: v16}
            var v27 uint16 = 2
            if true {
                var v28 float64 = (1e10 + ((float64(float32(0.1)) / 1e10) * (1e10 + -2.25)))
                fmt.Println(v16)
                _ = v28
            }
            _ = v26
            _ = v27
        }
        var v29 float64 = -2.25
        for fi30 := 0; fi30 < 12; fi30++ {
            fmt.Println(v1)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v31 int64 = 0
            _ = fi30
            _ = v31
        }
        _ = fi24
        _ = v29
    }
    var v32 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v32
    _ = v8
    _ = v9
}
