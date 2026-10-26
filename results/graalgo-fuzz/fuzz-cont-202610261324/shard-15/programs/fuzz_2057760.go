package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 uint8 = 255
    var v4 float32 = 0.1
    switch "" {
    case "a":
        v4 = float32(1.5)
    case "":
        v4 -= v4
    default:
        if true {
            var v5 float64 = (3.14159 / 1.5)
            _ = v5
        }
    }
    var v6 string = "a"
    if false {
        var v7 int8 = 42
        fmt.Println(v4, v7, v6)
        _ = v7
    } else {
        v8 := GenS1{F0: 255, F1: 32767, F2: 95}
        if (true || (26 <= v3)) {
            v3 = v3
        } else {
            for fi9 := 0; fi9 < 20; fi9++ {
                switch 2147483647 {
                case 42:
                    v3 -= (^(v3))
                case 942:
                    v3 = (v3 | (v3 * (v3 & 255)))
                default:
                    v10 := []int{}
                    _ = v10
                }
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint = 42
                    var v13 int16 = 1
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi9
            }
            var v14 int8 = 0
            switch (v6 + v6) {
            case "hello":
                fmt.Println(v3, v4, v3)
            }
            _ = v14
        }
        _ = v8
    }
    v3 -= v3
    if true {
        v15 := []int{692, 1}
        _ = v15
    }
    if (-1 >= 42) {
        fmt.Println(v4)
        v16 := GenS1{F0: 2, F1: -32768, F2: (v3 % (((v3 & (v3 * (v3 << ((0) & 63))))) | 1))}
        v6 = v6
        _ = v16
    } else {
        _ = len(v1)
        v17 := GenS2{F0: 65535, F1: 862}
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS2{F0: 2, F1: (fi18 >> ((42) & 63))}
        _ = fi18
        _ = v19
    }
    delete(v1, -860)
    if (32767 < 129) {
        v6 = v6
        var v20 int8 = 127
        _ = v20
    }
    if (!(v6 != "hello")) {
        var v21 float64 = (-2.25 - 1e10)
        v22 := []int{-1}
        var v23 uint16 = 255
        _ = v21
        _ = v22
        _ = v23
    } else {
        var v24 int16 = -504
        _ = v24
    }
    fmt.Println(v6, v6)
    var v25 float64 = ((float64(v4) - 3.14159) / 1.0)
    v26 := v1[0]
    for fi27 := 0; fi27 < 13; fi27++ {
        if true {
            v3 = (^(v3))
        } else {
            v3 = v3
        }
        var v28 int32 = -1
        _ = fi27
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v6
}
