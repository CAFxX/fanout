package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 int64 }
func genFn0(p0 uint32, p1 uint32, p2 int8) uint16 {
    switch -2147483648 {
    case 9223372036854775807:
        v5 := GenS1{F0: ((("foo bar" + "foo bar") + "\n\t") < "a"), F1: ("" + "")}
        _ = v5
    }
    var v6 int16 = 1
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 169
}
func genFn1(p0 uint8, p1 bool, p2 string) bool {
    var v8 int8 = -128
    var v9 int8 = (v8 >> ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (-601 > -483)
}

func main() {
    var v1 int16 = 136
    var v2 uint32 = 65535
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    v10 := GenS1{F0: false, F1: ("hello" + "hello")}
    var v11 uint16 = 42
    v1 += 504
    switch 32767 {
    case 32767:
        v12 := []int{}
        _ = v12
    default:
        var v13 int8 = 96
        _ = v13
    }
    var v14 uint16 = v11
    v1 = (v1 - (^(v1)))
    fmt.Println(v1)
    if (0 > 127) {
        switch (("hello" + "hello") + "a") {
        case "hello":
            v2 = v2
        case "":
            v2 += (v2 + (v2 * v2))
        default:
            v14 = v14
        }
        v2 = v2
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                if true {
                    v2 = (v2 << ((2) & 63))
                }
                _ = fi16
            }
            _ = fi15
        }
    }
    if (false && false) {
        var v17 uint16 = genFn0(v2, (v2 ^ v2), 14)
        fmt.Println(v2, v11)
        var v18 uint32 = v2
        _ = v17
        _ = v18
    }
    v19 := []int{-2147483648, 127, -9223372036854775808, 127}
    v20 := GenS1{F0: false, F1: ("a" + "\n\t")}
    v21 := GenS2{F0: ("\n\t" + "a"), F1: int64(v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
