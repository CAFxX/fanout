package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
func genFn0(p0 uint16, p1 bool, p2 int) int32 {
    if ("\n\t" != "hello") {
        for fi5 := 0; fi5 < 10; fi5++ {
            var v6 bool = true
            v7 := GenS1{F0: 42}
            _ = fi5
            _ = v6
            _ = v7
        }
        fmt.Println(p0, p2)
        var v8 int8 = -128
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 string = "\n\t"
        var v12 uint16 = 65535
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 int = (p2 % ((p2) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 255
    v2 := []string{}
    var v3 float64 = -2.25
    var v4 int32 = 2147483647
    var v14 uint16 = v1
    v15 := GenS1{F0: 42}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 uint = 42
    fmt.Println(v3, v1)
    v17 := GenS2{F0: 18446744073709551615, F1: 225, F2: 18446744073709551615}
    fmt.Println(v16, v16)
    for fi18 := 0; fi18 < 6; fi18++ {
        v3 = (0.0 - v3)
        fmt.Println(v3)
        v19 := GenS3{F0: (^(v14)), F1: 9223372036854775807}
        _ = fi18
        _ = v19
    }
    v3 -= (v3 + v3)
    v20 := GenS4{F0: 5, F1: 1, F2: 2147483648}
    v2 = append(v2, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
