package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int8 = 42
    v2 := make(map[string]uint32)
    var v3 int64 = 42
    var v4 uint = 1
    var v5 uint32 = 87
    if true {
        var v6 int = int(v5)
        _ = v6
    }
    v7 := GenS1{F0: (-(v3)), F1: 325}
    _ = len(v2)
    v1 -= 0
    var v8 bool = (42 == -2147483648)
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v9 float32 = -2.25
        _ = v9
    }
    v5 += v5
    var v10 int8 = (v1 << ((65535) & 63))
    if false {
        switch 563 {
        case 42:
            v2[("\n\t" + ("a" + "foo bar"))] = 2147483648
        default:
            switch ("\n\t" + "hello") {
            case "\n\t":
                v5 += (v5 | uint32(v4))
            case "hello":
                v11 := []int{1, 9223372036854775807, 32768, -1}
                _ = v11
            case "a":
                v1 -= int8(v5)
            }
        }
        v8 = (3.14159 > float64(float32(float64(3.14159))))
        v3 = (v3 % ((v3) | 1))
    }
    fmt.Println(v5, v1)
    var v12 int64 = 1
    _ = len(v2)
    if ((!((true && false) || v8)) || v8) {
        var v13 int16 = -255
        _ = v13
    }
    v1 = v10
    v14 := GenS2{F0: 433}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
