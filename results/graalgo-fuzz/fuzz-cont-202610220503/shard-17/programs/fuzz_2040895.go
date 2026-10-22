package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 uint }

func main() {
    var v1 uint64 = 710
    var v2 bool = (1 != 42)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v5 := []uint32{uint32(v1), 0, 4294967295, 549}
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(v1, v2, fi7)
        _ = fi7
    }
    switch "foo bar" {
    case "hello":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        var v8 int16 = -32768
        _ = v8
    }
    var v9 uint32 = 4294967295
    switch "a" {
    case "a":
        v10 := GenS1{F0: v9}
        _ = v10
    }
    v11 := GenS2{F0: 479, F1: -2147483648, F2: 0}
    for fi12 := 0; fi12 < 8; fi12++ {
        if (!(!true)) {
            delete(v6, (fi12 + (^(fi12))))
            var v13 int8 = 100
            _ = v13
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                delete(v4, v9)
                if (!true) {
                    delete(v4, 65535)
                    v15 := GenS2{F0: 867, F1: 1, F2: (fi12 + (fi12 | fi12))}
                    fmt.Println(v1, fi14, fi14)
                    v2 = (128 <= 2)
                    _ = v15
                } else {
                    v16 := GenS3{F0: 638, F1: -32768}
                    _ = v16
                }
                var v17 uint16 = 2
                _ = v17
            }
            var v18 uint = 65535
            if false {
                v19 := GenS2{F0: -1, F1: 0, F2: fi14}
                if true {
                    v1 += (v1 * 0)
                    fmt.Println(fi14, fi12)
                    v20 := GenS4{F0: (^(v18)), F1: 1}
                    var v21 float64 = (0.1 + (3.14159 * 1e10))
                    _ = v20
                    _ = v21
                }
                var v22 float64 = float64(-2.25)
                _ = v19
                _ = v22
            } else {
                var v23 int16 = 1
                v24 := GenS1{F0: v9}
                _ = v23
                _ = v24
            }
            _ = fi14
            _ = v18
        }
        if (v2 || (65535 <= 1)) {
            v25 := v3[(v9 % ((0) | 1))]
            _ = v25
        } else {
            v2 = v2
        }
        _ = fi12
    }
    if v2 {
        v9 -= (v9 / ((65535) | 1))
    }
    v2 = (42 <= -1)
    v5 = append(v5, uint32(v1))
    var v26 uint8 = 2
    var v27 uint8 = (v26 << ((65535) & 63))
    var v28 uint = 259
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
