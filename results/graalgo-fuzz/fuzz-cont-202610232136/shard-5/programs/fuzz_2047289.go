package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 int16 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int = 947
    var v3 bool = ((32767 == (-(v2))) && (1e10 < 3.14159))
    var v4 int8 = -1
    v5 := make(map[uint32]string)
    var v6 int = v2
    var v7 string = "\n\t"
    switch -2147483648 {
    case 128:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint64 = 42
                _ = fi9
                _ = v10
            }
            v3 = (42 > -32768)
            v11 := GenS1{F0: -9223372036854775808}
            _ = fi8
            _ = v11
        }
    case -9223372036854775808:
        v7 = v7
    }
    v2 -= v2
    for fi12 := 0; fi12 < 14; fi12++ {
        v4 = 101
        _ = fi12
    }
    var v13 uint16 = 551
    switch v6 {
    case 60:
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int8 = (^(v4))
            var v16 int64 = -1
            _ = fi14
            _ = v15
            _ = v16
        }
    case 4294967295:
        v2 = v2
    case 9223372036854775807:
        v3 = v3
    }
    v5[4294967295] = "foo bar"
    v17 := GenS2{F0: 358, F1: 255, F2: 42}
    v3 = (345 < uint(v13))
    v18 := []int{(v6 >> ((0) & 63)), (v2 & (-(v6))), 32767, 9223372036854775807}
    var v19 uint8 = 190
    v13 -= (v13 / ((v13) | 1))
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 int32 = -2147483648
        if v3 {
            if true {
                v22 := GenS3{F0: 0}
                var v23 int = (v6 ^ v2)
                _ = v22
                _ = v23
            }
            v3 = v3
        } else {
            _ = len(v5)
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
