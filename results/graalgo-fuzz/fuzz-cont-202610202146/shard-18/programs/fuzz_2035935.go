package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint64 }

func main() {
    var v1 int = 42
    var v2 float32 = (1e10 + 1e10)
    var v3 bool = true
    var v4 int = (^(v1))
    var v5 int16 = 42
    var v6 int8 = 0
    var v7 uint64 = 2
    fmt.Println(v2, v3, v6)
    var v8 uint8 = 27
    var v9 int64 = 42
    switch "a" {
    case "hello":
        switch (v4 & -796) {
        case 32767:
            for fi10 := 0; fi10 < 1; fi10++ {
                v3 = v3
                _ = fi10
            }
        case 127:
            if v3 {
                for fi11 := 0; fi11 < 19; fi11++ {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        v13 := GenS1{F0: v9, F1: (v8 + (v8 - v8)), F2: 451}
                        if true {
                            var v14 uint8 = 240
                            v5 += (v5 - (v5 << ((65535) & 63)))
                            _ = v14
                        }
                        v4 = fi12
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 float64 = 1.5
                    var v17 uint8 = 128
                    fmt.Println(v2, v17)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            } else {
                if true {
                    switch 32768 {
                    case 325:
                        v18 := GenS1{F0: v9, F1: (v8 << ((108) & 63)), F2: v7}
                        _ = v18
                    case 32767:
                        v3 = false
                    case -2147483648:
                        var v19 int = v1
                        _ = v19
                    }
                    var v20 uint8 = 184
                    _ = v20
                } else {
                    var v21 uint = 659
                    _ = v21
                }
            }
        case 128:
            if v3 {
                if v3 {
                    switch "\n\t" {
                    case "hello":
                        v22 := GenS1{F0: (v9 << ((34) & 63)), F1: 128, F2: v7}
                        _ = v22
                    default:
                        var v23 float64 = 0.0
                        _ = v23
                    }
                } else {
                    v24 := []int{-9223372036854775808, (v1 >> ((367) & 63)), (v1 ^ (-(v4)))}
                    v24 = append(v24, v4)
                    v25 := GenS1{F0: 1, F1: (v8 >> ((65535) & 63)), F2: v7}
                    _ = v24
                    _ = v25
                }
                v3 = v3
            }
        default:
            v2 = (float32(1.5) - (v2 - float32(1e10)))
        }
    }
    v26 := []int{(v1 * v1), (v1 ^ (v1 >> ((4294967295) & 63)))}
    var v27 uint = 344
    v9 = (v9 * (v9 % ((v9) | 1)))
    var v28 uint = v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
