package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint8 }

func main() {
    var v1 float64 = -2.25
    var v2 bool = (true && false)
    v3 := make(map[int]uint32)
    var v4 float32 = -2.25
    v5 := make(map[string]int)
    v6 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v1, v2)
        v8 := []int{fi7}
        v9 := GenS1{F0: 464, F1: 0}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        delete(v6, 0)
        _ = fi10
    }
    var v11 string = ("a" + "foo bar")
    v12 := GenS2{F0: 879, F1: (!(2147483647 > 42)), F2: 72}
    if v2 {
        var v13 uint8 = 227
        v11 = v11
        _ = v13
    } else {
        v1 = 0.0
        var v14 int = 2147483647
        v2 = v2
        _ = v14
    }
    fmt.Println(v1, v1, v4)
    v15 := v6[4294967295]
    v4 -= (v4 - v4)
    var v16 uint32 = 1
    v17 := GenS2{F0: 2, F1: (884 != 32768), F2: 0}
    fmt.Println(v1)
    if (false && v2) {
        var v18 int = 1
        if v2 {
            var v19 bool = (2 < 42)
            for fi20 := 0; fi20 < 13; fi20++ {
                v21 := GenS2{F0: 255, F1: false, F2: 198}
                fmt.Println(v4, v2, v15)
                _ = fi20
                _ = v21
            }
            _ = v19
        }
        var v22 int8 = 127
        _ = v18
        _ = v22
    } else {
        v23 := []int{2147483647, 2147483647, -1, 1}
        v11 = v11
        _ = v23
    }
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
