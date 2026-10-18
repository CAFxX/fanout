package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 float32 = float32(3.14159)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 42
    var v4 int64 = int64(v3)
    if true {
        switch -18 {
        case -645:
            for fi5 := 0; fi5 < 1; fi5++ {
                if false {
                    var v6 int8 = 122
                    _ = v6
                }
                _ = fi5
            }
        case -941:
            var v7 int32 = -1
            _ = v7
        case -9223372036854775808:
            v3 = v3
        }
    } else {
        var v8 float32 = v1
        var v9 int = 32767
        if true {
            v10 := GenS1{F0: 1, F1: 65535}
            _ = v10
        }
        _ = v8
        _ = v9
    }
    var v11 float32 = v1
    v12 := GenS1{F0: 18446744073709551615, F1: 2}
    v13 := GenS1{F0: 0, F1: 873}
    v3 = v3
    v14 := []int{127, 9223372036854775807}
    v11 += (0.0 / ((-2.25 + 0.1) / v1))
    delete(v2, 4294967295)
    if (false || ((0 != 18446744073709551615) || (!(((0 != -1) || false) || (!((!(false || (true || false))) || false)))))) {
        fmt.Println(v3, v3)
        v15 := GenS1{F0: 1, F1: 121}
        switch 32767 {
        case -9223372036854775808:
            _ = len(v2)
        case 1:
            var v16 uint64 = 2
            _ = v16
        case 128:
            for fi17 := 0; fi17 < 9; fi17++ {
                v1 += float32(1e10)
                _ = fi17
            }
        default:
            v4 = v4
        }
        _ = v15
    }
    v3 -= (v3 | v3)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v18 := GenS1{F0: 42, F1: 141}
    v2[0] = ("hello" >= "\n\t")
    v1 = v11
    v3 += v3
    v3 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
