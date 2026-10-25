package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint8 = 67
    var v2 int32 = 2147483647
    v3 := []uint32{0, 4294967295, 2}
    v4 := []int{1, 9223372036854775807, -9223372036854775808, 127}
    v5 := []uint32{627}
    v6 := GenS1{F0: 4294967295, F1: 0}
    v2 -= v2
    var v7 int16 = 0
    v7 += -1000
    switch 9223372036854775807 {
    case 42:
        for fi8 := 0; fi8 < 4; fi8++ {
            v1 = (^(v1))
            _ = fi8
        }
    default:
        _ = len(v5)
    }
    var v9 bool = false
    var v10 uint16 = uint16(v1)
    if (!((!((false || ((v10 * (v10 * 255)) != (v10 - v10))) || true)) || (("hello" == ((("hello" + "\n\t") + "") + ("hello" + (("hello" + "hello") + (("foo bar" + ("hello" + ("hello" + ("foo bar" + ((("\n\t" + "hello") + "\n\t") + "\n\t"))))) + ""))))) || (72 < 4294967295)))) {
        v7 = (v7 % ((v7) | 1))
        switch "a" {
        case "foo bar":
            v3 = append(v3, uint32(v10))
        case "":
            v10 += 440
        }
        v5 = append(v5, 538)
    } else {
        if (461 != 125) {
            fmt.Println(v10, v7)
            v10 += (v10 / (((v10 << ((0) & 63))) | 1))
            v1 = v1
        } else {
            switch 32768 {
            case 9223372036854775807:
                switch "a" {
                case "foo bar":
                    fmt.Println(v10, v1)
                case "\n\t":
                    fmt.Println(v10, v2, v9)
                }
            }
        }
        var v11 uint64 = 1
        _ = v11
    }
    var v12 string = ("a" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
