package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
func genFn0() string {
    if true {
        var v4 int64 = 0
        _ = v4
    } else {
        var v5 uint8 = 119
        var v6 uint16 = 0
        switch "\n\t" {
        case "":
            v7 := []int{0}
            _ = v7
        case "\n\t":
            v6 -= v6
        case "hello":
            v6 = 0
        }
        _ = v5
        _ = v6
    }
    var v8 int32 = 1
    v9 := []int{-9223372036854775808}
    _ = v8
    _ = v9
    return (("a" + "hello") + ("\n\t" + "foo bar"))
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    var v3 int16 = -1
    switch -310 {
    case 42:
        var v10 uint = 2
        _ = v10
    case 0:
        v11 := v2[700]
        _ = v11
    default:
        var v12 string = "hello"
        _ = v12
    }
    v13 := []int{}
    var v14 int16 = v3
    if (3.14159 <= ((1e10 / 3.14159) * -2.25)) {
        v3 += (v3 % ((1000) | 1))
        switch v1 {
        case "":
            var v15 uint8 = 220
            _ = v15
        default:
            if false {
                v16 := GenS1{F0: 255, F1: 100, F2: 18}
                _ = v16
            } else {
                _ = len(v2)
            }
        }
    }
    var v17 int64 = -854
    v1 = v1
    v1 = (("foo bar" + ((v1 + "") + "foo bar")) + "\n\t")
    v17 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
}
