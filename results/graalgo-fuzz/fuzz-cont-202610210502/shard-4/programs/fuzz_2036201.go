package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint64; F1 uint; F2 int16 }
type GenS3 struct { F0 uint32; F1 int; F2 int64 }
func genFn0(p0 int64, p1 bool, p2 uint16) uint32 {
    p1 = ((!p1) && true)
    var v6 uint16 = p2
    fmt.Println(p0, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 0
    var v3 int = -211
    v4 := []int{}
    var v5 uint8 = 127
    fmt.Println(v2, v1, v5)
    fmt.Println(v5, v1, v3)
    var v7 bool = (32767 > (v2 % ((v2) | 1)))
    v7 = (!true)
    var v8 int32 = -2147483648
    switch 42 {
    case 810:
        for fi9 := 0; fi9 < 5; fi9++ {
            v5 = 128
            fmt.Println(v2, v8)
            for fi10 := 0; fi10 < 5; fi10++ {
                v7 = (320 > -1)
                v8 = (^(v8))
                _ = fi10
            }
            _ = fi9
        }
    case 422:
        for fi11 := 0; fi11 < 20; fi11++ {
            v3 += -138
            if (0.0 <= float32(((1e10 * 3.14159) + 3.14159))) {
                var v12 uint64 = 4294967295
                v4 = append(v4, 32767)
                _ = v12
            }
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: 440, F1: (!true)}
                _ = fi13
                _ = v14
            }
            _ = fi11
        }
    default:
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := GenS2{F0: 725, F1: 18446744073709551615, F2: int16(v2)}
            v17 := GenS2{F0: 397, F1: 65535, F2: -781}
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    var v18 int64 = -108
    var v19 int8 = -128
    v20 := GenS3{F0: 42, F1: 42, F2: v18}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
