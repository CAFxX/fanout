package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int; F2 int32 }

func main() {
    var v1 uint = 42
    var v2 uint16 = 0
    var v3 float32 = (float32((float64(float32(0.0)) * -2.25)) + float32(0.0))
    var v4 float32 = v3
    var v5 int64 = -9223372036854775808
    v6 := []int{0, 9223372036854775807, 32767, 496}
    fmt.Println(v1, v5)
    if true {
        var v7 int64 = v5
        v7 = 1
        v3 += (-2.25 * (float32((-2.25 - 1.5)) * v4))
        _ = v7
    }
    if (false || (!true)) {
        if len(v6) > 0 {
            v6[0] = 128
        }
    }
    v1 = 1
    fmt.Println(v4, v2)
    v3 += v3
    v3 = 0.1
    v8 := GenS1{F0: 125}
    if len(v6) > 0 {
        v6[0] = 32767
    }
    var v9 uint16 = v2
    switch "foo bar" {
    case "hello":
        switch ("a" + "\n\t") {
        case "a":
            if false {
                var v10 int8 = -1
                v11 := GenS1{F0: 1}
                _ = v10
                _ = v11
            } else {
                var v12 string = ""
                _ = v12
            }
        }
    case "foo bar":
        var v13 int8 = 42
        _ = v13
    case "\n\t":
        var v14 uint8 = 2
        _ = v14
    default:
        switch -2147483648 {
        case -1:
            v15 := GenS2{F0: 127, F1: -9223372036854775808, F2: 42}
            _ = v15
        case 0:
            if (true || (-1 < 1)) {
                v1 = (v1 % ((v1) | 1))
                for fi16 := 0; fi16 < 7; fi16++ {
                    if (0 > v5) {
                        var v17 float64 = 0.1
                        v3 -= v4
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v6 = append(v6, (^(fi18)))
                            v9 += (v2 - (v2 ^ (v9 % (((v9 | (v9 * v2))) | 1))))
                            _ = fi18
                        }
                        _ = v17
                    }
                    var v19 int32 = 42
                    _ = fi16
                    _ = v19
                }
                fmt.Println(v2, v9, v1)
            }
        }
    }
    v1 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
