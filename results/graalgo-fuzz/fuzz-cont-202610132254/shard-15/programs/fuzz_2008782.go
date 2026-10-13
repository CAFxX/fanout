package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 uint = 1
    if (!true) {
        var v5 int16 = int16(v1)
        v1 = 128
        v1 += v1
        _ = v5
    } else {
        var v6 int64 = 844
        _ = v6
    }
    v1 -= v1
    v7 := GenS1{F0: 118, F1: 699, F2: 2}
    v8 := GenS1{F0: 1, F1: 306, F2: 18446744073709551615}
    fmt.Println(v1)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: true, F1: -389, F2: 2}
        if true {
            v4 = 255
            var v11 uint8 = (v1 & 170)
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(v1)
    fmt.Println(v4, v1)
    if (!(0 > -1)) {
        if true {
            v12 := []int{-9223372036854775808, -604}
            fmt.Println(v1, v4)
            if (((-677 > 1) || (-1 != -2147483648)) || (false || true)) {
                for fi13 := 0; fi13 < 3; fi13++ {
                    v1 = 74
                    if false {
                        fmt.Println(v1)
                        v4 = v4
                    } else {
                        v1 -= v1
                    }
                    if (!((1.5 / 0.1) == ((3.14159 + 1.5) * 1.5))) {
                        if (!false) {
                            var v14 int32 = 42
                            _ = v14
                        }
                        fmt.Println(v1, v4)
                        if (!("hello" != "\n\t")) {
                            fmt.Println(v1)
                            v4 = 0
                        }
                    } else {
                        var v15 int16 = 264
                        v16 := GenS1{F0: 100, F1: 65535, F2: 2}
                        _ = v15
                        _ = v16
                    }
                    _ = fi13
                }
                v17 := GenS3{F0: (((("foo bar" + ("hello" + "hello")) + ("a" + "")) + (("\n\t" + "\n\t") + "hello")) + ("foo bar" + ("\n\t" + "a")))}
                _ = v17
            } else {
                for fi18 := 0; fi18 < 16; fi18++ {
                    v12 = append(v12, -2147483648)
                    _ = fi18
                }
                if (1 < 2) {
                    var v19 uint64 = 1
                    switch 183 {
                    case 4294967295:
                        v1 = 81
                    default:
                        var v20 float32 = ((1.5 + 1.5) + 0.1)
                        _ = v20
                    }
                    _ = v19
                } else {
                    v1 -= v1
                }
            }
            _ = v12
        } else {
            fmt.Println(v4, v4)
        }
        v21 := GenS2{F0: true, F1: 9223372036854775807, F2: 42}
        _ = v21
    }
    var v22 uint32 = 0
    if true {
        switch "\n\t" {
        case "\n\t":
            v22 = v22
        case "":
            for fi23 := 0; fi23 < 9; fi23++ {
                v22 = v22
                v24 := []int{}
                v1 = 2
                _ = fi23
                _ = v24
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
