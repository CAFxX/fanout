package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int16; F2 int16 }

func main() {
    var v1 int32 = 650
    var v2 uint8 = 255
    v3 := make(map[int]int)
    var v4 int = 32767
    var v5 float32 = 1e10
    v2 -= 1
    if true {
        v6 := []int{v4, 32767}
        _ = v6
    }
    v7 := []int{int(v2), (v4 % (((-(v4))) | 1))}
    var v8 int8 = 42
    v9 := GenS1{F0: v1}
    for fi10 := 0; fi10 < 12; fi10++ {
        v4 = fi10
        v1 = 637
        var v11 float32 = -2.25
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS2{F0: 42, F1: 4294967295, F2: v2}
        _ = fi12
        _ = v13
    }
    fmt.Println(v2, v4, v2)
    fmt.Println(v8)
    if true {
        switch "a" {
        case "foo bar":
            var v14 int8 = int8(v4)
            _ = v14
        case "a":
            v4 += v4
        }
        v15 := GenS1{F0: 801}
        _ = v15
    }
    v16 := GenS3{F0: (v2 & (v2 / (((^(v2))) | 1)))}
    v2 -= v2
    v17 := v3[(v4 * (v4 << ((1000) & 63)))]
    for fi18 := 0; fi18 < 6; fi18++ {
        switch (^(fi18)) {
        case 32767:
            v19 := GenS4{F0: (v1 % ((int32(v8)) | 1)), F1: 0, F2: -890}
            _ = v19
        case -2147483648:
            var v20 int32 = 2147483647
            _ = v20
        default:
            var v21 int = (v4 - int(v8))
            _ = v21
        }
        _ = fi18
    }
    fmt.Println(v5, v17)
    var v22 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
