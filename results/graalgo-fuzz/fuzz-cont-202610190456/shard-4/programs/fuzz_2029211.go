package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{42, 341}
    var v3 uint16 = 32767
    var v4 string = "foo bar"
    var v5 float32 = 1e10
    v6 := make(map[uint32]string)
    v2 = append(v2, 2147483648)
    var v7 uint = 1
    if (false && false) {
        _ = len(v2)
        v8 := GenS1{F0: -609}
        _ = v8
    }
    v5 += v5
    var v9 float32 = 1.5
    for fi10 := 0; fi10 < 5; fi10++ {
        v9 = ((-2.25 + (float32(1.5) + -2.25)) / float32(1e10))
        var v11 int16 = 1000
        _ = fi10
        _ = v11
    }
    fmt.Println(v9, v4, v7)
    v4 = "foo bar"
    v3 = 32767
    v5 += (((float32(1e10) - (3.14159 * (v5 / v9))) * ((v5 * v9) * v5)) - v5)
    for fi12 := 0; fi12 < 7; fi12++ {
        switch 32768 {
        case 128:
            v13 := GenS2{F0: -1}
            _ = v13
        }
        switch v4 {
        case "hello":
            fmt.Println(fi12, v4)
        case "\n\t":
            var v14 int32 = -45
            _ = v14
        default:
            var v15 string = "foo bar"
            _ = v15
        }
        var v16 float64 = 1.5
        _ = fi12
        _ = v16
    }
    var v17 uint8 = 225
    _ = len(v2)
    var v18 uint16 = (v3 << (((v7 ^ (v7 | (^(v7))))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
