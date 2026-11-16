package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
func genFn0() uint8 {
    var v6 uint32 = 0
    v6 += v6
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 481
    var v2 bool = false
    var v3 uint64 = (^(v1))
    var v4 float32 = (0.0 - 0.0)
    v5 := []int{-9223372036854775808, -9223372036854775808, 128}
    var v7 bool = false
    v1 = (v3 + v1)
    switch -353 {
    case 2147483647:
        v8 := GenS1{F0: 682, F1: (v3 + v3)}
        _ = v8
    }
    v9 := GenS2{F0: 309, F1: 463, F2: (-1 >= -47)}
    v10 := GenS3{F0: -1, F1: v1, F2: 799}
    var v11 uint64 = (v3 & v3)
    switch ("hello" + "hello") {
    case "hello":
        for fi12 := 0; fi12 < 5; fi12++ {
            v13 := GenS1{F0: 0, F1: 18446744073709551615}
            v14 := GenS1{F0: (^(v1)), F1: (v11 | (^(v3)))}
            _ = fi12
            _ = v13
            _ = v14
        }
    case "":
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v5) > 0 {
                v5[0] = fi15
            }
            _ = fi15
        }
    }
    v1 -= 1
    switch "foo bar" {
    case "a":
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 uint8 = 1
            var v18 int64 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch "a" {
    case "\n\t":
        var v19 int16 = 1
        _ = v19
    case "hello":
        v11 += (v11 + (v3 % (((^(v11))) | 1)))
    default:
        v3 = 152
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
