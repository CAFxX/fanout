package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
type GenS2 struct { F0 bool; F1 int16; F2 bool }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = true
    var v2 uint64 = 126
    v3 := []uint32{}
    var v4 uint16 = 0
    switch (("foo bar" + "hello") + "a") {
    case "a":
        var v5 int8 = 127
        _ = v5
    default:
        v1 = v1
    }
    switch ("" + ("hello" + ("hello" + "\n\t"))) {
    case "\n\t":
        var v6 int = 4294967295
        _ = v6
    case "a":
        v2 += 0
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            v1 = (1 < int8(fi7))
            _ = fi7
        }
    default:
        var v8 uint8 = 2
        _ = v8
    }
    var v9 float64 = 0.1
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: 1, F1: (fi10 % ((127) | 1)), F2: (fi10 * fi10)}
        var v12 int = (fi10 << ((2) & 63))
        if false {
            for fi13 := 0; fi13 < 5; fi13++ {
                if (uint(v2) < 210) {
                    var v14 int64 = -361
                    if true {
                        v15 := GenS1{F0: 42, F1: 391, F2: 32768}
                        var v16 string = ("hello" + "\n\t")
                        _ = len(v3)
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v12, fi10)
                    _ = v14
                } else {
                    v12 = (fi13 ^ (fi13 / (((fi10 % ((v12) | 1))) | 1)))
                    var v17 int8 = -1
                    _ = v17
                }
                _ = fi13
            }
        } else {
            v18 := GenS2{F0: true, F1: 1000, F2: v1}
            _ = v18
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    v19 := GenS1{F0: uint8(v2), F1: -476, F2: 42}
    v3 = append(v3, 2147483647)
    var v20 uint = 1
    var v21 float32 = (0.1 * (3.14159 - (((((((0.0 * 1.5) + 1e10) / float32(1.5)) / 0.1) / 1.5) + 1.5) * 3.14159)))
    v22 := GenS2{F0: (((!(42 >= 328)) && v1) && ((42 < 0) || false)), F1: 1000, F2: (v20 == (v20 | 495))}
    fmt.Println(v20, v4)
    var v23 int32 = -885
    fmt.Println(v23)
    fmt.Println(v9)
    if (v21 != v21) {
        v24 := GenS1{F0: 0, F1: -974, F2: -605}
        _ = v24
    } else {
        v21 = 3.14159
        for fi25 := 0; fi25 < 8; fi25++ {
            if v1 {
                var v26 uint32 = 819
                _ = v26
            }
            v3 = append(v3, 2147483647)
            v27 := GenS3{F0: 0}
            _ = fi25
            _ = v27
        }
        v28 := GenS2{F0: v1, F1: 953, F2: (74 >= 0)}
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
