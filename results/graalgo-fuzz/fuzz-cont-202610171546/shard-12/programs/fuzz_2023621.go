package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 bool, p1 uint32, p2 uint) int32 {
    var v4 float32 = 0.0
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 float32 = v4
        p2 = (p2 + (^(p2)))
        _ = fi5
        _ = v6
    }
    v7 := []int{-262, -931}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 42
}
func genFn1(p0 uint16, p1 uint64) bool {
    v8 := []int{127, -9223372036854775808, -2147483648}
    if ((p1 >> ((2) & 63)) < p1) {
        p1 = (p1 ^ p1)
    }
    var v9 int = 1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 uint = 255
    var v2 bool = (!((v1 >> (((^(v1))) & 63)) > v1))
    var v3 float32 = 3.14159
    fmt.Println(v1, v1, v1)
    switch ("foo bar" + "a") {
    case "":
        v2 = (!v2)
    case "foo bar":
        for fi10 := 0; fi10 < 9; fi10++ {
            v2 = false
            v11 := []int{(fi10 - fi10), 42, (fi10 ^ (^(fi10)))}
            _ = fi10
            _ = v11
        }
    default:
        v12 := []int{-216, 4294967295, 967}
        _ = v12
    }
    fmt.Println(v2)
    switch "a" {
    case "hello":
        v1 = (v1 | 4294967295)
    case "a":
        var v13 int32 = -317
        _ = v13
    case "\n\t":
        fmt.Println(v2, v2, v1)
    default:
        var v14 uint = 42
        _ = v14
    }
    fmt.Println(v2)
    v2 = (-9223372036854775808 == 42)
    var v15 int64 = 1
    v16 := GenS1{F0: "a", F1: 100}
    v2 = true
    var v17 int16 = int16(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
