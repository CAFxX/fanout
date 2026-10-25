package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 bool }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int8 = 42
    var v2 uint8 = 42
    v3 := []int{2147483647, 127, 9223372036854775807}
    v1 += (v1 / ((78) | 1))
    fmt.Println(v1)
    if true {
        var v4 uint8 = v2
        v1 += -128
        _ = v4
    }
    var v5 int32 = 42
    var v6 uint64 = 862
    fmt.Println(v1, v2)
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 int = fi7
            var v9 uint64 = v6
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int64 = 1
        _ = v10
    }
    fmt.Println(v5, v5, v6)
    v5 -= v5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v6 += v6
    for fi11 := 0; fi11 < 9; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: (v1 % ((v1) | 1))}
            v14 := GenS2{F0: ((v5 & v5) != (v5 >> ((484) & 63))), F1: 1000, F2: (^(v5))}
            _ = fi12
            _ = v13
            _ = v14
        }
        if false {
            v6 += (v6 >> ((18446744073709551615) & 63))
            var v15 int64 = -1
            _ = v15
        } else {
            switch -2147483648 {
            case -108:
                for fi16 := 0; fi16 < 17; fi16++ {
                    v1 = (v1 << ((891) & 63))
                    _ = fi16
                }
            case 9223372036854775807:
                v17 := GenS2{F0: (true || true), F1: 1, F2: v5}
                _ = v17
            case -139:
                _ = len(v3)
            }
        }
        _ = fi11
    }
    v18 := GenS3{F0: v5, F1: 42, F2: false}
    _ = len(v3)
    v19 := GenS4{F0: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
