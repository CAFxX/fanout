package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS4 struct { F0 int16; F1 uint; F2 bool }
func genFn0(p0 bool, p1 int) int32 {
    p1 = 110
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 bool = p0
        var v7 uint16 = 853
        v8 := GenS1{F0: 18446744073709551615}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 string, p1 int64) int16 {
    var v9 int16 = 1
    var v10 string = "a"
    v10 = v10
    p0 = "a"
    v11 := GenS2{F0: 0, F1: 263}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return v9
}

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 2
    var v3 uint32 = 2147483648
    var v4 int16 = 0
    v3 += v3
    v3 = (v3 * 2147483648)
    v12 := GenS2{F0: -177, F1: 0}
    var v13 uint8 = 183
    v4 = v4
    fmt.Println(v4)
    v14 := GenS3{F0: v13, F1: -9223372036854775808, F2: (v13 / ((v13) | 1))}
    for fi15 := 0; fi15 < 2; fi15++ {
        delete(v1, "hello")
        var v16 string = ""
        v17 := GenS4{F0: (v4 - int16(fi15)), F1: 18446744073709551615, F2: true}
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
