package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }

func main() {
    var v1 float32 = 3.14159
    var v2 uint32 = 59
    var v3 int8 = -112
    v4 := make(map[int]bool)
    var v5 float64 = 0.1
    fmt.Println(v5)
    switch 32768 {
    case 4294967295:
        var v6 uint16 = 42
        _ = v6
    }
    var v7 uint8 = 42
    v2 = (v2 / ((v2) | 1))
    fmt.Println(v7, v2)
    fmt.Println(v2, v1)
    v4[-1] = true
    var v8 uint8 = (v7 ^ v7)
    if (!true) {
        if true {
            if true {
                fmt.Println(v1)
            } else {
                var v9 string = "foo bar"
                v10 := GenS1{F0: 2, F1: v8, F2: -1}
                v5 += 1.5
                _ = v10
                _ = v9
            }
            if (("a" + "a") <= "") {
                for fi11 := 0; fi11 < 9; fi11++ {
                    delete(v4, (^(fi11)))
                    _ = fi11
                }
                v12 := []int{}
                _ = v12
            }
        } else {
            fmt.Println(v8, v8, v1)
        }
        switch 2147483647 {
        case 2147483647:
            for fi13 := 0; fi13 < 16; fi13++ {
                v3 -= (v3 % (((v3 / ((-128) | 1))) | 1))
                _ = fi13
            }
        default:
            v14 := GenS1{F0: 2, F1: (v7 % ((127) | 1)), F2: 346}
            _ = v14
        }
    }
    var v15 int16 = -32768
    switch "hello" {
    case "hello":
        v16 := GenS2{F0: 504, F1: (v3 << ((18446744073709551615) & 63)), F2: (v7 << ((2) & 63))}
        _ = v16
    case "a":
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int32 = -794
            var v19 float64 = (v5 * v5)
            _ = fi17
            _ = v18
            _ = v19
        }
    case "":
        var v20 float32 = (3.14159 - (v1 + 0.0))
        _ = v20
    }
    var v21 int8 = v3
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 float64 = (v5 * float64(0.0))
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
