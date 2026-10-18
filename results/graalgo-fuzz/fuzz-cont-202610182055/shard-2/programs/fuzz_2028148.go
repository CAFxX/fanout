package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS4 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint, p1 bool) uint64 {
    p0 -= (^(p0))
    v5 := []int{-608, 32767, 9223372036854775807, 128}
    if len(v5) > 0 {
        v5[0] = 0
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v5
    return 687
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[string]int)
    var v3 uint32 = 2147483648
    var v4 string = "foo bar"
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: v1, F1: 212, F2: false}
        _ = fi6
        _ = v7
    }
    v8 := GenS2{F0: 128, F1: -1, F2: false}
    v9 := GenS3{F0: 32767, F1: 1, F2: 42}
    var v10 int64 = 761
    v3 -= v3
    v11 := GenS3{F0: -32768, F1: 18446744073709551615, F2: 32767}
    v12 := GenS4{F0: false, F1: 32767}
    v3 += (v3 & 2147483647)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS3{F0: 429, F1: 18446744073709551615, F2: 32767}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: v1, F1: 127, F2: true}
    var v16 string = v4
    for fi17 := 0; fi17 < 4; fi17++ {
        var v18 bool = (("foo bar" + "") == (v16 + v4))
        _ = fi17
        _ = v18
    }
    var v19 int32 = int32(v10)
    var v20 int16 = 42
    v16 = (((v4 + v4) + ("foo bar" + "foo bar")) + v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
