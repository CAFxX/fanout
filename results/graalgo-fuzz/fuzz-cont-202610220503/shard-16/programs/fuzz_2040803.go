package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []uint32{2, 971}
    v2 := []int{9223372036854775807, 1}
    v3 := make(map[uint32]int)
    var v4 uint = 2
    for fi5 := 0; fi5 < 3; fi5++ {
        switch ("a" + "\n\t") {
        case "hello":
            fmt.Println(v4)
        default:
            if false {
                v4 -= (v4 + 0)
                if ((((!true) || false) && (!(-331 > -1))) && true) {
                    v4 = v4
                } else {
                    var v6 int32 = 2147483647
                    v6 = -2147483648
                    _ = v6
                }
            }
        }
        v7 := GenS1{F0: 42}
        _ = fi5
        _ = v7
    }
    switch -991 {
    case 42:
        var v8 int16 = -1
        _ = v8
    case 32767:
        switch "foo bar" {
        case "hello":
            var v9 int8 = 100
            _ = v9
        case "foo bar":
            v4 = (v4 - v4)
        default:
            v10 := GenS2{F0: 592, F1: 42, F2: 62}
            _ = v10
        }
    }
    v4 -= (v4 ^ (v4 ^ 871))
    v11 := GenS2{F0: -1, F1: 18446744073709551615, F2: 128}
    v4 -= v4
    v4 = v4
    var v12 int32 = 663
    fmt.Println(v12, v4)
    var v13 int16 = -32768
    var v14 bool = (-2 <= 13)
    for fi15 := 0; fi15 < 12; fi15++ {
        v14 = (1 != 954)
        switch "\n\t" {
        case "a":
            var v16 int = fi15
            _ = v16
        default:
            v17 := GenS2{F0: 0, F1: 4294967295, F2: 42}
            _ = v17
        }
        _ = len(v3)
        _ = fi15
    }
    var v18 int64 = 9223372036854775807
    if false {
        if v14 {
            v12 = (v12 ^ v12)
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS3{F0: 2}
                _ = fi19
                _ = v20
            }
            var v21 bool = (v18 != (v18 % ((514) | 1)))
            _ = v21
        }
    }
    v22 := GenS3{F0: 2}
    v18 -= 535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
