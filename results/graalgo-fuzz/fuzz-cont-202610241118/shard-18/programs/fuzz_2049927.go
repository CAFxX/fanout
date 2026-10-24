package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0(p0 uint8, p1 uint) int16 {
    fmt.Println(p1, p0)
    switch 0 {
    case -421:
        if true {
            var v5 float32 = 3.14159
            p0 -= (^(p0))
            _ = v5
        } else {
            if true {
                v6 := []int{1}
                _ = v6
            }
            v7 := GenS1{F0: 1}
            _ = v7
        }
    }
    _ = p0
    _ = p1
    return 1000
}

func main() {
    var v1 uint8 = 116
    var v2 int64 = -1
    var v3 int8 = -1
    var v4 string = "foo bar"
    var v8 float64 = (1e10 / ((3.14159 / 1.0) * -2.25))
    v8 = 1.5
    v9 := GenS2{F0: 1}
    v4 = ("hello" + v4)
    fmt.Println(v1, v1)
    v3 = (^(v3))
    var v10 uint = 1
    switch "" {
    case "hello":
        v11 := GenS2{F0: 0}
        _ = v11
    case "foo bar":
        v3 = (v3 << (((v10 ^ (v10 / ((v10) | 1)))) & 63))
    case "\n\t":
        fmt.Println(v1, v3, v3)
    default:
        fmt.Println(v10, v2, v2)
    }
    v8 -= v8
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{(fi12 ^ (fi12 & (fi12 % (((^(fi12))) | 1)))), fi12, fi12, 9223372036854775807}
        for fi14 := 0; fi14 < 2; fi14++ {
            v8 = v8
            v15 := GenS3{F0: (v2 * int64(v3)), F1: 541, F2: v3}
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v10, v10, v10)
    v16 := GenS2{F0: genFn0(v1, uint(v1))}
    for fi17 := 0; fi17 < 7; fi17++ {
        v18 := GenS4{F0: uint32(v1), F1: true}
        fmt.Println(fi17)
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
