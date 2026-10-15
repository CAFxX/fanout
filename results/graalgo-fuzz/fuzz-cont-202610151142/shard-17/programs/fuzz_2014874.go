package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }
func genFn0() string {
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(fi8, fi8)
            var v9 int32 = -1
            v9 = 111
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            v11 := GenS1{F0: 1000, F1: 2}
            var v12 float64 = 3.14159
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    var v13 bool = (!(1 == 42))
    if (!v13) {
        fmt.Println(v13)
        var v14 uint32 = 491
        fmt.Println(v14)
        _ = v14
    } else {
        for fi15 := 0; fi15 < 19; fi15++ {
            if true {
                v16 := []int{(fi15 | (^(fi15))), 2147483647}
                _ = v16
            } else {
                v13 = (-44 == 365)
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v13 = v13
                v18 := []int{fi17}
                var v19 uint32 = 2147483648
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint64 = 987
            _ = fi15
            _ = v20
        }
        v13 = (2 > 255)
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v13 = v13
        var v22 uint64 = 0
        _ = fi21
        _ = v22
    }
    _ = v13
    return ("foo bar" + "foo bar")
}

func main() {
    v1 := []int{}
    var v2 uint16 = 32767
    var v3 int32 = 870
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := make(map[uint32]uint32)
    v3 += v3
    v23 := GenS2{F0: 1, F1: 0, F2: (^(v2))}
    v24 := GenS2{F0: 42, F1: 1, F2: (v2 >> ((4294967295) & 63))}
    v2 += (v2 >> ((255) & 63))
    var v25 float32 = 3.14159
    var v26 string = ((("hello" + "a") + "foo bar") + ("a" + "hello"))
    v26 = v26
    _ = len(v5)
    v3 = v3
    var v27 uint16 = (^(v2))
    fmt.Println(v2, v3)
    var v28 int32 = (v3 ^ (-(v3)))
    v29 := GenS1{F0: 42, F1: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
