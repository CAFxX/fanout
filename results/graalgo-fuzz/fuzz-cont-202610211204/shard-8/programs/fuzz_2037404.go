package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 float64 = (0.0 + 1e10)
    v5 := []uint32{0, 4294967295, 0}
    v6 := make(map[int]int)
    var v7 int16 = 202
    if false {
        if true {
            var v8 uint = 596
            var v9 uint = (v8 >> ((255) & 63))
            var v10 int32 = -1
            _ = v10
            _ = v8
            _ = v9
        } else {
            v11 := GenS1{F0: 18446744073709551615, F1: 255, F2: 1}
            v12 := GenS1{F0: 1, F1: 46, F2: 807}
            switch ((((("a" + "") + "") + ("hello" + "\n\t")) + "foo bar") + "foo bar") {
            case "hello":
                var v13 uint16 = 1
                _ = v13
            }
            _ = v11
            _ = v12
        }
    } else {
        for fi14 := 0; fi14 < 20; fi14++ {
            v7 = 1000
            var v15 uint32 = 102
            _ = fi14
            _ = v15
        }
        if (602 <= 42) {
            fmt.Println(v7)
            v16 := GenS2{F0: 127, F1: 1}
            v7 -= 1000
            _ = v16
        } else {
            if true {
                var v17 int64 = 0
                switch "a" {
                case "hello":
                    var v18 uint16 = 0
                    _ = v18
                case "foo bar":
                    v17 -= (v17 | v17)
                default:
                    v19 := GenS2{F0: 7, F1: 799}
                    _ = v19
                }
                _ = len(v1)
                _ = v17
            }
            v20 := GenS1{F0: 18446744073709551615, F1: 656, F2: 32767}
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v7 = (v7 & v7)
        if false {
            var v22 float32 = float32(v4)
            for fi23 := 0; fi23 < 14; fi23++ {
                v4 = (3.14159 - v4)
                var v24 uint8 = 1
                v24 -= (^(v24))
                _ = fi23
                _ = v24
            }
            _ = v22
        }
        _ = fi21
    }
    var v25 uint16 = 42
    var v26 float64 = float64(0.0)
    if len(v5) > 0 {
        v5[0] = 1
    }
    fmt.Println(v4, v4, v4)
    var v27 uint32 = 369
    v28 := GenS1{F0: 271, F1: (v25 | (v25 & (v25 / (((^(v25))) | 1)))), F2: 42}
    v25 -= (^(v25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
