package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 string) int64 {
    v5 := GenS1{F0: 2, F1: 2147483648, F2: -376}
    v6 := GenS2{F0: 42}
    if (734 >= 42) {
        switch ((p1 + (p1 + "hello")) + p1) {
        case "\n\t":
            p1 = p1
        default:
            switch p1 {
            case "a":
                v7 := []int{-467}
                _ = v7
            default:
                v8 := []int{-2147483648, -1}
                _ = v8
            }
        }
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 0
    var v2 uint = 255
    var v3 bool = false
    v4 := make(map[int]string)
    v3 = (false && (!(128 > 246)))
    var v9 uint64 = v1
    fmt.Println(v2, v1)
    for fi10 := 0; fi10 < 18; fi10++ {
        v11 := []int{fi10, (fi10 ^ (-(fi10)))}
        _ = fi10
        _ = v11
    }
    v1 = (v9 & (v9 | (v1 + v1)))
    v3 = false
    var v12 uint32 = 318
    v9 = 1
    v13 := GenS3{F0: 240, F1: -1}
    switch int(v12) {
    case 263:
        _ = len(v4)
    }
    if ((!(255 <= 63)) && true) {
        var v14 uint64 = v9
        v15 := GenS4{F0: v2}
        _ = v14
        _ = v15
    }
    v9 = v9
    fmt.Println(v9, v3, v2)
    var v16 uint = 4294967295
    var v17 uint64 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
