package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
func genFn0(p0 bool, p1 bool) int8 {
    v4 := []int{-731, -674}
    if p0 {
        fmt.Println(p1, p0)
    } else {
        var v5 int = -9223372036854775808
        switch v5 {
        case -1:
            for fi6 := 0; fi6 < 8; fi6++ {
                for fi7 := 0; fi7 < 7; fi7++ {
                    p0 = (("\n\t" + "hello") == "a")
                    if (223 <= (fi6 << ((108) & 63))) {
                        switch "\n\t" {
                        case "":
                            var v8 string = "foo bar"
                            _ = v8
                        default:
                            fmt.Println(p0)
                        }
                        _ = len(v4)
                    } else {
                        fmt.Println(p0, fi7, p1)
                        if false {
                            var v9 float32 = float32(-2.25)
                            _ = v9
                        } else {
                            var v10 uint8 = 42
                            _ = v10
                        }
                    }
                    _ = fi7
                }
                _ = fi6
            }
        }
        _ = v5
    }
    p0 = true
    v11 := GenS1{F0: 1, F1: -208}
    switch 787 {
    case 32768:
        var v12 uint64 = 18446744073709551615
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v4
    return 0
}

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[uint32]bool)
    var v3 float64 = (1e10 - (1.5 + 0.1))
    var v13 int8 = 1
    v3 = float64(0.1)
    var v14 uint16 = 645
    v3 += 1e10
    v15 := GenS1{F0: 185, F1: -32768}
    var v16 uint32 = 2
    v17 := GenS1{F0: 2, F1: -32768}
    var v18 uint32 = v16
    if (128 < 128) {
        v1 += v1
    } else {
        if true {
            v19 := GenS1{F0: 0, F1: -1}
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int = fi20
                var v22 uint = 223
                var v23 uint16 = v14
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = v19
        } else {
            v18 = v18
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 uint16 = (v14 / ((90) | 1))
                v13 = (-(v13))
                _ = fi24
                _ = v25
            }
            v26 := v2[v18]
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
