package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint }
type GenS4 struct { F0 bool }
func genFn0(p0 int32, p1 uint8) int8 {
    v5 := []int{2147483647}
    var v6 int = int(p0)
    var v7 int8 = 42
    var v8 uint16 = 587
    for fi9 := 0; fi9 < 20; fi9++ {
        v6 -= (v6 | (v6 | 128))
        v5 = append(v5, v6)
        fmt.Println(v6)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (v7 | v7)
}

func main() {
    var v1 int32 = -121
    var v2 uint32 = 42
    var v3 string = "\n\t"
    var v4 float32 = 0.1
    var v10 uint8 = 255
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v3 = "hello"
            _ = fi12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: 42, F1: 417, F2: uint8(v2)}
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v1)
    if (778 > 0) {
        if false {
            v10 = v10
        } else {
            var v15 int32 = -475
            v16 := GenS1{F0: 42, F1: 9223372036854775807, F2: 231}
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: 0}
    v10 = 176
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := GenS3{F0: 32767, F1: 1, F2: 255}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := []int{fi20, -2147483648, (^(fi20))}
        v10 = (v10 / ((47) | 1))
        _ = fi20
        _ = v21
    }
    v3 = v3
    v3 = ((v3 + v3) + ("" + v3))
    var v22 uint16 = 717
    v22 = v22
    switch v3 {
    case "a":
        if (true || true) {
            v23 := GenS4{F0: (65535 >= 438)}
            _ = v23
        } else {
            if true {
                fmt.Println(v4, v3)
            } else {
                v4 = v4
            }
            if true {
                v22 = (^(v22))
            }
        }
    }
    var v24 float64 = 0.1
    var v25 int64 = 42
    var v26 uint64 = 2
    var v27 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
