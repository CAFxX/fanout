package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 uint8 = 255
    var v4 int16 = 1000
    var v5 int16 = (v4 ^ v4)
    if false {
        var v6 int32 = -2147483648
        _ = v6
    }
    switch 9223372036854775807 {
    case -1:
        switch -89 {
        case 426:
            var v7 float32 = (0.0 - 0.1)
            _ = v7
        case 128:
            v8 := GenS1{F0: 78, F1: 1, F2: 65535}
            _ = v8
        case -2147483648:
            var v9 uint16 = 0
            _ = v9
        }
    }
    fmt.Println(v4)
    v10 := GenS1{F0: -94, F1: -1, F2: 803}
    if true {
        switch "" {
        case "a":
            var v11 uint8 = (v3 * v3)
            _ = v11
        case "foo bar":
            v3 = (v3 / ((v3) | 1))
        case "":
            v3 = v3
        default:
            var v12 int32 = 291
            _ = v12
        }
        v13 := []int{}
        switch "a" {
        case "hello":
            var v14 int8 = 100
            _ = v14
        }
        _ = v13
    }
    var v15 string = (("" + "hello") + "a")
    var v16 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
