package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 0.0
    var v3 float32 = 0.0
    var v4 int = -781
    var v5 float64 = 3.14159
    var v6 uint64 = 804
    var v7 bool = (128 < 69)
    v6 = (v6 % ((v6) | 1))
    v8 := GenS1{F0: 2}
    v3 = v2
    if ((!v7) && false) {
        switch (v4 + 0) {
        case 1:
            fmt.Println(v2, v3, v5)
        }
    }
    fmt.Println(v5)
    var v9 uint8 = 66
    fmt.Println(v3)
    v10 := GenS2{F0: 9223372036854775807, F1: -1}
    for fi11 := 0; fi11 < 13; fi11++ {
        switch (fi11 | -1) {
        case 754:
            for fi12 := 0; fi12 < 4; fi12++ {
                var v13 float32 = v2
                v14 := GenS2{F0: 42, F1: 100}
                v2 = v3
                _ = fi12
                _ = v13
                _ = v14
            }
        case -348:
            v15 := GenS3{F0: -933, F1: v6}
            _ = v15
        case 32767:
            v16 := GenS4{F0: 18446744073709551615, F1: 1000, F2: 100}
            _ = v16
        default:
            if v7 {
                v17 := GenS2{F0: -1000, F1: 100}
                if true {
                    var v18 string = ("" + "hello")
                    _ = v18
                }
                v2 = 3.14159
                _ = v17
            }
        }
        var v19 uint64 = v6
        if (1 == 32767) {
            var v20 uint16 = 1
            _ = v20
        }
        _ = fi11
        _ = v19
    }
    var v21 string = ("foo bar" + "\n\t")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
