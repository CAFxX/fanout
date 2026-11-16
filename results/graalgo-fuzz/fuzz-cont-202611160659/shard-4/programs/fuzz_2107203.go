package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int; F2 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]bool)
    var v3 uint8 = 2
    var v4 int64 = 1
    fmt.Println(v1, v1, v1)
    if true {
        v5 := GenS1{F0: int(v3)}
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 float32 = 1e10
            _ = fi6
            _ = v7
        }
        _ = v5
    }
    v8 := GenS2{F0: 68, F1: 36, F2: 18446744073709551615}
    v1 = 1
    for fi9 := 0; fi9 < 7; fi9++ {
        switch "a" {
        case "\n\t":
            if (-186 >= -960) {
                v1 = v1
                v10 := GenS1{F0: (fi9 + (fi9 + fi9))}
                _ = v10
            } else {
                if false {
                    for fi11 := 0; fi11 < 18; fi11++ {
                        v3 = uint8(v1)
                        _ = fi11
                    }
                } else {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 float64 = 1.5
                        _ = fi12
                        _ = v13
                    }
                }
                var v14 uint16 = (v1 * v1)
                fmt.Println(fi9)
                _ = v14
            }
        case "foo bar":
            if false {
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 bool = (!("\n\t" <= "hello"))
                    var v17 uint8 = (v3 + (v3 >> ((0) & 63)))
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v3 -= (^(v3))
            }
        }
        v4 -= v4
        _ = fi9
    }
    v18 := GenS3{F0: false, F1: -987}
    v19 := GenS1{F0: 4294967295}
    switch 2147483647 {
    case 32768:
        if true {
            v20 := GenS1{F0: 4294967295}
            fmt.Println(v1, v1)
            v3 -= 1
            _ = v20
        } else {
            var v21 uint = 2
            _ = v21
        }
    case -9223372036854775808:
        v1 = (v1 * (v1 * v1))
    default:
        fmt.Println(v1, v1)
    }
    v22 := []int{}
    var v23 uint64 = 0
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS4{F0: v3, F1: (fi24 > (fi24 << ((45) & 63)))}
        var v26 int16 = -32768
        var v27 int64 = -1
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    _ = len(v2)
    v28 := GenS5{F0: false}
    var v29 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
