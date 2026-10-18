package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 0
    v3 := []uint32{2147483648}
    var v4 uint8 = 0
    fmt.Println(v4, v2)
    fmt.Println(v2)
    switch 0 {
    case 32768:
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 string = "a"
            v2 += 127
            var v7 uint16 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    switch 1 {
    case 4294967295:
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: -1, F1: 0, F2: 0}
            v10 := GenS1{F0: -32768, F1: 1, F2: 215}
            _ = fi8
            _ = v10
            _ = v9
        }
    case -743:
        switch 0 {
        case 2147483647:
            for fi11 := 0; fi11 < 4; fi11++ {
                v12 := GenS1{F0: 1, F1: 0, F2: 4294967295}
                _ = fi11
                _ = v12
            }
        case -9223372036854775808:
            v13 := GenS1{F0: 0, F1: 42, F2: 42}
            _ = v13
        }
    }
    fmt.Println(v2)
    v2 = (v2 << ((18446744073709551615) & 63))
    var v14 uint16 = 682
    v15 := GenS2{F0: 0, F1: 32768}
    for fi16 := 0; fi16 < 18; fi16++ {
        v1[fi16] = fi16
        switch (^(fi16)) {
        case 128:
            var v17 int16 = 636
            _ = v17
        case 193:
            if false {
                if true {
                    fmt.Println(v4, v2)
                }
                if len(v3) > 0 {
                    v3[0] = 42
                }
                v18 := GenS3{F0: (fi16 % ((fi16) | 1)), F1: 255}
                _ = v18
            }
        default:
            var v19 bool = ((!false) && (false || false))
            _ = v19
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
