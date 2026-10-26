package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 int16 = 1
    v2 := []uint32{794}
    v3 := make(map[int]int)
    var v4 string = "foo bar"
    v5 := []int{}
    var v6 bool = true
    v1 = (v1 | v1)
    v6 = v6
    v7 := GenS1{F0: "foo bar"}
    v8 := GenS2{F0: -220, F1: 0, F2: 691}
    for fi9 := 0; fi9 < 10; fi9++ {
        var v10 float32 = (3.14159 - -2.25)
        _ = fi9
        _ = v10
    }
    v11 := GenS3{F0: 255}
    v4 = "hello"
    v1 -= 32767
    v12 := v3[1]
    if ((2 == 2147483647) && true) {
        v13 := GenS4{F0: v1, F1: (("" + v4) + (((v4 + v4) + v4) + v4))}
        var v14 uint32 = 0
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        _ = len(v2)
        if v6 {
            v1 = v1
            v4 = v4
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                if (1.5 < float64(3.14159)) {
                    v17 := GenS1{F0: (v4 + v4)}
                    _ = v17
                } else {
                    v18 := GenS5{F0: 1}
                    var v19 int16 = 42
                    if (-1 != -910) {
                        if false {
                            _ = len(v3)
                            _ = len(v3)
                        }
                    }
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                v12 = 4294967295
                var v21 uint32 = 42
                v1 = 1000
                _ = fi20
                _ = v21
            }
        }
        _ = fi15
    }
    v22 := GenS6{F0: 834}
    if (false || (42 <= 42)) {
        var v23 uint = 4294967295
        _ = v23
    }
    var v24 uint8 = 1
    v24 -= (v24 >> ((18446744073709551615) & 63))
    switch v4 {
    case "foo bar":
        v12 = -688
    }
    var v25 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
