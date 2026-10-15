package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int64; F1 int64; F2 bool }
type GenS4 struct { F0 int8; F1 int8 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{42, 65535, 1}
    v3 := make(map[int]int)
    var v4 string = v1
    var v5 int32 = -1
    v6 := v3[1]
    for fi7 := 0; fi7 < 8; fi7++ {
        v1 = v4
        v8 := GenS1{F0: 2}
        _ = fi7
        _ = v8
    }
    v4 = v4
    v6 -= v6
    v5 = (v5 + v5)
    v5 -= v5
    v9 := GenS1{F0: 41}
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v1)
        _ = fi10
    }
    v4 = "foo bar"
    fmt.Println(v5, v5, v5)
    if true {
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 uint16 = 255
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        delete(v3, v6)
        switch v6 {
        case -1:
            switch (-(v6)) {
            case 32768:
                var v14 float64 = (float64(0.0) * 1.5)
                _ = v14
            case 32767:
                switch (v1 + v4) {
                case "\n\t":
                    fmt.Println(v6, v5, v1)
                }
            }
        case 128:
            v15 := GenS2{F0: 127, F1: 42}
            _ = v15
        default:
            var v16 uint16 = 0
            _ = v16
        }
    } else {
        v5 += -1
        switch v1 {
        case "":
            v6 += (v6 - 42)
        case "a":
            v17 := GenS3{F0: -1, F1: -9223372036854775808, F2: true}
            _ = v17
        default:
            v18 := GenS2{F0: 96, F1: 18446744073709551615}
            _ = v18
        }
    }
    v1 = v4
    v19 := GenS1{F0: 2}
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 2; fi21++ {
        v20 -= 4294967295
        v22 := GenS4{F0: 0, F1: 0}
        _ = fi21
        _ = v22
    }
    v6 += (v6 ^ v6)
    v23 := GenS3{F0: 754, F1: -1, F2: (!((-(v5)) != (v5 + 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
