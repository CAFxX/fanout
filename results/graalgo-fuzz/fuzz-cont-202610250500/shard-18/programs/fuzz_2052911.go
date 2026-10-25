package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = true
    var v3 int = 1
    v4 := make(map[string]uint32)
    var v5 int = v3
    for fi6 := 0; fi6 < 13; fi6++ {
        if false {
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int16 = 42
                var v9 int = fi6
                var v10 bool = v2
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            v11 := GenS1{F0: 0, F1: 959, F2: 42}
            _ = v11
        }
        _ = fi6
    }
    var v12 int64 = 0
    v2 = false
    v3 += (v3 + -402)
    v3 = (v5 << ((4294967295) & 63))
    var v13 int8 = 30
    if (false || v2) {
        v3 -= 115
        if ((v13 < (v13 / (((v13 + v13)) | 1))) || ((v3 >= (v5 << ((492) & 63))) || (!v2))) {
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint64 = 2
                _ = fi14
                _ = v15
            }
        }
    } else {
        fmt.Println(v3)
        if v2 {
            v16 := GenS2{F0: v5, F1: 212}
            _ = v16
        }
        _ = len(v4)
    }
    for fi17 := 0; fi17 < 18; fi17++ {
        v18 := []int{}
        var v19 uint16 = 32767
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 1}
    v21 := []int{(v3 ^ 9223372036854775807)}
    var v22 uint32 = 2
    var v23 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
