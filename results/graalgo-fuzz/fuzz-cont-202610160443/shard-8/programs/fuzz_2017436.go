package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
func genFn0(p0 uint, p1 uint8, p2 uint) uint {
    var v7 bool = true
    fmt.Println(p2, p2)
    v7 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 959
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = 100
    var v3 bool = true
    v4 := make(map[uint32]int)
    var v5 float64 = 1.5
    var v6 string = ("" + "a")
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 4294967295, F1: 42}
        v10 := GenS1{F0: 18446744073709551615, F1: 4294967295}
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 += v2
    v1[709] = ("\n\t" + v6)
    v6 = v6
    fmt.Println(v6, v6, v3)
    v6 = v6
    fmt.Println(v3, v5)
    fmt.Println(v3, v2, v6)
    var v11 int32 = -644
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := GenS2{F0: (v2 % (((v2 + 1)) | 1)), F1: 9223372036854775807, F2: v3}
        var v14 int8 = 30
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
