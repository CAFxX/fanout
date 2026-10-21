package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    var v1 uint = 65535
    var v2 uint = (v1 * (v1 + v1))
    v3 := make(map[string]bool)
    var v4 int16 = -415
    var v5 bool = (643 < 0)
    fmt.Println(v1)
    var v6 string = "\n\t"
    v7 := []int{-2147483648, 127, 32768}
    v8 := GenS1{F0: int(v4), F1: ((((v6 + "hello") + "") + ("hello" + ((v6 + v6) + (v6 + ((v6 + "a") + (((v6 + v6) + v6) + v6)))))) + v6)}
    v2 += (v1 ^ v1)
    v4 += v4
    v4 = (v4 - v4)
    v5 = true
    var v9 int = 42
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint = (^(v2))
        switch fi10 {
        case 127:
            v12 := GenS1{F0: v9, F1: ("a" + v6)}
            _ = v12
        case 42:
            v4 -= v4
        case 4294967295:
            v2 += (^(v11))
        default:
            for fi13 := 0; fi13 < 20; fi13++ {
                v4 = (v4 & -168)
                v4 += v4
                v14 := GenS1{F0: -9223372036854775808, F1: "\n\t"}
                _ = fi13
                _ = v14
            }
        }
        var v15 int8 = 44
        _ = fi10
        _ = v11
        _ = v15
    }
    delete(v3, (v6 + (v6 + (("" + v6) + v6))))
    v5 = (-1 != 980)
    for fi16 := 0; fi16 < 8; fi16++ {
        if (255 != 127) {
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS1{F0: fi16, F1: ((("hello" + v6) + "") + v6)}
                _ = len(v3)
                v19 := GenS1{F0: fi16, F1: v6}
                _ = fi17
                _ = v18
                _ = v19
            }
            fmt.Println(v1, v1, v6)
            v4 -= (v4 << (((v2 << (((v1 * 255)) & 63))) & 63))
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
