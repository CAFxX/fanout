package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    var v3 int16 = -32768
    var v4 uint32 = v2
    var v5 uint8 = 60
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        fmt.Println(v2, v4, v4)
    }
    var v6 uint64 = 2
    v3 -= v3
    v7 := GenS1{F0: true, F1: true, F2: -72}
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v5, v3, v3)
        delete(v1, v2)
        _ = fi8
    }
    var v9 int8 = 0
    var v10 int64 = 1
    var v11 float64 = float64((float32((1.5 / -2.25)) - (float32(3.14159) - 3.14159)))
    v12 := GenS1{F0: (!(!(!((false && (true && true)) || true)))), F1: false, F2: v9}
    var v13 float64 = float64(1e10)
    v14 := GenS2{F0: (v4 + (v4 | (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))))}
    for fi15 := 0; fi15 < 15; fi15++ {
        for fi16 := 0; fi16 < 13; fi16++ {
            v4 = 0
            _ = fi16
        }
        v17 := GenS3{F0: 72, F1: (^(v4)), F2: (v6 ^ (v6 >> ((1) & 63)))}
        fmt.Println(v9, v5, v11)
        _ = fi15
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
