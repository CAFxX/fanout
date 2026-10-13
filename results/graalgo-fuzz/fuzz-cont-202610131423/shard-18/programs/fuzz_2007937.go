package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 int8 }
func genFn0() bool {
    var v4 bool = (2147483647 == 0)
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: 2, F1: "a", F2: 155}
        _ = fi5
        _ = v6
    }
    _ = v4
    return v4
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = -998
    var v3 uint64 = uint64(v1)
    fmt.Println(v2, v2, v2)
    var v7 bool = ((v3 ^ (^(v3))) == (^(v3)))
    v3 = v3
    v8 := GenS2{F0: 42}
    var v9 string = "a"
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS3{F0: (v3 / ((v3) | 1)), F1: 1}
        v2 -= v2
        var v12 int = (fi10 ^ (-(fi10)))
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := []int{32767, -9223372036854775808, -2147483648, 563}
    fmt.Println(v7)
    for fi14 := 0; fi14 < 16; fi14++ {
        v7 = (-1 < 127)
        _ = fi14
    }
    v3 = v3
    v2 = v2
    v1 = (v1 + v1)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
