package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int }

func main() {
    var v1 float64 = (float64(float32(-2.25)) / 1.0)
    var v2 uint64 = 244
    var v3 uint64 = 1
    var v4 bool = false
    var v5 uint16 = 32767
    if v4 {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint32 = 2
            v1 = ((float64(3.14159) / ((float64(0.1) * float64(1e10)) * v1)) / 0.1)
            var v8 uint32 = 42
            _ = fi6
            _ = v7
            _ = v8
        }
        if ("" == ("foo bar" + "foo bar")) {
            var v9 float32 = 3.14159
            switch -1 {
            case -586:
                var v10 uint64 = 18446744073709551615
                _ = v10
            case 2147483647:
                switch "hello" {
                case "hello":
                    if v4 {
                        var v11 uint = 18446744073709551615
                        v1 = v1
                        for fi12 := 0; fi12 < 9; fi12++ {
                            var v13 bool = (!((v2 / ((v3) | 1)) < v2))
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    }
                case "a":
                    for fi14 := 0; fi14 < 10; fi14++ {
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v1 -= ((float64(v9) + 0.0) - v1)
                            var v16 int16 = 895
                            var v17 uint = 0
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v2 = v3
                        _ = fi14
                    }
                }
            case 32768:
                switch "hello" {
                case "\n\t":
                    switch (("\n\t" + ("" + "\n\t")) + "\n\t") {
                    case "foo bar":
                        v18 := GenS1{F0: 248, F1: -461, F2: 2147483647}
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := []int{0, fi19}
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            default:
                v21 := []int{-1, 4294967295, 32768}
                _ = v21
            }
            switch 32767 {
            case 127:
                v22 := []int{1, 42, -949}
                _ = v22
            default:
                var v23 uint8 = 127
                _ = v23
            }
            _ = v9
        } else {
            fmt.Println(v3, v2, v3)
        }
        fmt.Println(v2, v3)
    }
    v24 := []int{}
    v4 = false
    fmt.Println(v1, v3, v3)
    v25 := GenS1{F0: 2, F1: -1, F2: 32767}
    for fi26 := 0; fi26 < 3; fi26++ {
        v27 := GenS1{F0: 42, F1: -9223372036854775808, F2: 0}
        v2 = (^(v3))
        v24 = append(v24, (fi26 - fi26))
        _ = fi26
        _ = v27
    }
    fmt.Println(v5)
    var v28 uint16 = (v5 / (((v5 & (v5 | v5))) | 1))
    var v29 bool = ((65535 > 263) || v4)
    switch 128 {
    case 1:
        v1 += v1
    default:
        fmt.Println(v1, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
