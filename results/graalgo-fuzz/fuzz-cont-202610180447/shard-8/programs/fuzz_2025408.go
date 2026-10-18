package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 2
    var v3 float64 = ((-2.25 + 1e10) * 0.0)
    var v4 string = ("foo bar" + "\n\t")
    var v5 int8 = -21
    v6 := make(map[uint32]string)
    var v7 int8 = v5
    v8 := []int{-2147483648, -715}
    var v9 uint32 = 760
    var v10 uint64 = 473
    var v11 uint32 = (v9 | v9)
    switch v4 {
    case "hello":
        v11 -= v11
    case "\n\t":
        if false {
            if false {
                v5 -= 127
                if len(v8) > 0 {
                    v8[0] = 975
                }
            } else {
                if len(v8) > 0 {
                    v8[0] = 9223372036854775807
                }
            }
            switch 264 {
            case 1:
                var v12 int8 = v5
                _ = v12
            case 127:
                var v13 int = 128
                _ = v13
            default:
                var v14 float64 = 1e10
                _ = v14
            }
        } else {
            v5 += v5
            switch v4 {
            case "hello":
                var v15 int16 = -1
                _ = v15
            case "a":
                v5 = (v5 << ((0) & 63))
            }
        }
    case "a":
        for fi16 := 0; fi16 < 18; fi16++ {
            switch ((v4 + v4) + v4) {
            case "hello":
                var v17 int = 1
                _ = v17
            case "\n\t":
                switch (fi16 << ((124) & 63)) {
                case 981:
                    for fi18 := 0; fi18 < 8; fi18++ {
                        var v19 int64 = 42
                        _ = fi18
                        _ = v19
                    }
                case 42:
                    var v20 int16 = 0
                    _ = v20
                default:
                    v1[(fi16 * fi16)] = fi16
                }
            case "":
                v21 := GenS1{F0: v2}
                _ = v21
            default:
                var v22 int32 = -839
                _ = v22
            }
            if false {
                v5 -= (-(v7))
                delete(v1, (fi16 + (fi16 << ((255) & 63))))
                v8 = append(v8, (fi16 / (((fi16 - (fi16 + (^(fi16))))) | 1)))
            } else {
                v23 := v1[(fi16 + -9223372036854775808)]
                _ = v23
            }
            _ = fi16
        }
    default:
        v4 = v4
    }
    v10 += v10
    var v24 int8 = (v7 & v5)
    var v25 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
