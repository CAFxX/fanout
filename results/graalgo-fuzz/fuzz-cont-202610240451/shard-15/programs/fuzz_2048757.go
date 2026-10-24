package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }

func main() {
    var v1 bool = false
    var v2 uint8 = 127
    var v3 string = "hello"
    var v4 uint8 = (v2 + 127)
    fmt.Println(v2)
    if v1 {
        switch 0 {
        case 0:
            var v5 uint32 = 42
            _ = v5
        }
        switch v3 {
        case "a":
            v6 := GenS1{F0: 100, F1: 523}
            _ = v6
        default:
            var v7 uint8 = (v2 ^ (v4 | (v2 << ((4294967295) & 63))))
            _ = v7
        }
    }
    v2 -= 1
    v4 += (^(v4))
    v8 := GenS2{F0: v4, F1: 42}
    for fi9 := 0; fi9 < 8; fi9++ {
        v1 = (-9223372036854775808 < -9223372036854775808)
        switch (fi9 >> ((821) & 63)) {
        case 4294967295:
            v3 = "hello"
        case 1:
            v10 := GenS3{F0: 4294967295, F1: false, F2: 2}
            _ = v10
        case -187:
            v4 = v4
        }
        _ = fi9
    }
    v11 := GenS4{F0: -34}
    if (32768 > 2147483647) {
        v12 := GenS4{F0: -117}
        v4 = (^(v2))
        _ = v12
    }
    v2 = 0
    v2 += v2
    var v13 float32 = (0.0 + ((3.14159 - 0.1) * -2.25))
    v14 := GenS5{F0: -1}
    v15 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
