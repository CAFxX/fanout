package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint8) uint64 {
    var v5 uint32 = 965
    fmt.Println(p1, p1, p1)
    v6 := GenS1{F0: 406, F1: 0, F2: 42}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}
func genFn1() int {
    var v7 bool = (!true)
    if v7 {
        var v8 float64 = (0.0 - -2.25)
        _ = v8
    }
    var v9 int32 = -637
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 string = ("\n\t" + ("foo bar" + ""))
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 161, F1: -835, F2: (v9 << ((753) & 63))}
    _ = v12
    _ = v7
    _ = v9
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 int = 4294967295
    var v4 bool = (true || true)
    fmt.Println(v4)
    v3 = (v3 << ((138) & 63))
    v2[9223372036854775807] = v3
    var v13 uint32 = 2
    var v14 uint64 = 194
    var v15 int16 = 0
    fmt.Println(v14, v3)
    var v16 int64 = 9223372036854775807
    v17 := GenS2{F0: -814}
    var v18 bool = (((0.1 + 0.0) != (1.5 / float32(0.1))) || true)
    v16 -= -1
    if v18 {
        if false {
            for fi19 := 0; fi19 < 13; fi19++ {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 int16 = v15
                    fmt.Println(v16, fi19, v4)
                    _ = fi20
                    _ = v21
                }
                _ = fi19
            }
            v15 = (v15 & v15)
            var v22 float64 = -2.25
            _ = v22
        }
    } else {
        var v23 uint64 = (v14 % ((v14) | 1))
        _ = v23
    }
    fmt.Println(v15, v18, v18)
    v24 := GenS1{F0: 255, F1: 998, F2: 557}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
