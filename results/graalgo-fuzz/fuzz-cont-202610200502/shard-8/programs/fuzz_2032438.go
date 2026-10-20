package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint16; F1 int8; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int64 }

func main() {
    var v1 int16 = 671
    var v2 uint8 = 34
    v3 := make(map[uint32]uint32)
    var v4 int64 = 42
    v5 := []int{}
    v6 := GenS1{F0: v2, F1: 32767}
    fmt.Println(v2, v1, v2)
    switch (("" + (("hello" + "foo bar") + "\n\t")) + "") {
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    default:
        fmt.Println(v1, v4)
    }
    v7 := v3[2]
    v4 = 9223372036854775807
    var v8 uint8 = (^(v2))
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            if (!false) {
                if (!true) {
                    v10 := v3[(v7 * (v7 / ((4294967295) | 1)))]
                    v11 := GenS2{F0: 855, F1: 0, F2: (v4 >> ((959) & 63))}
                    _ = v10
                    _ = v11
                } else {
                    var v12 uint32 = 0
                    v7 = v7
                    _ = v12
                }
                v4 -= (v4 << ((160) & 63))
            }
            v13 := GenS1{F0: v2, F1: fi9}
            _ = v13
        } else {
            for fi14 := 0; fi14 < 20; fi14++ {
                switch "foo bar" {
                case "a":
                    v15 := GenS2{F0: 10, F1: 127, F2: (v4 & 700)}
                    _ = v15
                case "":
                    v16 := GenS1{F0: (v2 % ((237) | 1)), F1: 32768}
                    _ = v16
                case "\n\t":
                    var v17 int16 = v1
                    _ = v17
                }
                v7 = (v7 & v7)
                v18 := v3[607]
                _ = fi14
                _ = v18
            }
        }
        _ = fi9
    }
    _ = len(v5)
    var v19 uint32 = v7
    var v20 uint8 = 97
    v21 := GenS3{F0: -9223372036854775808}
    var v22 bool = true
    switch "\n\t" {
    case "a":
        switch 271 {
        case -587:
            v23 := GenS4{F0: (v4 | 42), F1: (v4 & v4)}
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS4{F0: int64(v8), F1: v4}
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
