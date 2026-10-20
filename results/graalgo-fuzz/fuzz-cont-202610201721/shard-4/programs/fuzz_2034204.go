package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int; F2 int }
type GenS4 struct { F0 string; F1 uint32; F2 bool }

func main() {
    var v1 int32 = 1
    var v2 int = -300
    var v3 bool = ((float32(0.0) * float32((-2.25 + (1e10 + 3.14159)))) > 1.5)
    var v4 float64 = ((3.14159 * 1.5) - float64(3.14159))
    var v5 uint64 = 42
    switch "hello" {
    case "":
        for fi6 := 0; fi6 < 1; fi6++ {
            switch v2 {
            case 32768:
                v7 := []int{(-(v2)), 4294967295}
                _ = v7
            case 127:
                v8 := GenS1{F0: 153}
                _ = v8
            }
            v9 := []int{(fi6 << ((272) & 63)), int(v1), fi6}
            v10 := GenS2{F0: 1}
            _ = fi6
            _ = v10
            _ = v9
        }
    case "hello":
        fmt.Println(v1, v4)
    }
    var v11 int8 = -1
    v12 := GenS3{F0: 1, F1: v2, F2: (v2 - (v2 >> ((42) & 63)))}
    var v13 float64 = 3.14159
    var v14 string = "foo bar"
    v5 += v5
    v15 := GenS4{F0: v14, F1: 159, F2: true}
    v1 -= (-(v1))
    v16 := []int{v2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
