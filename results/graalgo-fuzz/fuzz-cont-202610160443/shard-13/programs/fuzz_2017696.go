package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int16; F1 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 int8 = 42
    var v3 uint = 42
    var v4 int64 = 451
    var v5 int16 = 0
    var v6 uint8 = 42
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            var v8 float64 = 0.1
            if false {
                for fi9 := 0; fi9 < 19; fi9++ {
                    v4 = v4
                    _ = fi9
                }
            } else {
                var v10 uint = 0
                _ = v10
            }
            _ = v8
        } else {
            v11 := GenS1{F0: 1, F1: 938}
            if false {
                v1 = (v1 - uint(v6))
            } else {
                v12 := GenS2{F0: (^(v5)), F1: 22}
                _ = v12
            }
            _ = v11
        }
        _ = fi7
    }
    var v13 uint64 = 4294967295
    v14 := []int{-251, 42}
    if len(v14) > 0 {
        v14[0] = 32768
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 float32 = 0.1
    var v16 int = -2147483648
    v17 := GenS2{F0: int16(v1), F1: 735}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
