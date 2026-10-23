package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64, p1 int, p2 uint32) uint {
    v4 := GenS1{F0: -32768, F1: false}
    switch "foo bar" {
    case "":
        v5 := GenS1{F0: 1, F1: (255 >= 2)}
        _ = v5
    case "\n\t":
        v6 := []int{}
        _ = v6
    case "a":
        switch "\n\t" {
        case "foo bar":
            switch (-(p1)) {
            case 128:
                var v7 uint8 = 1
                _ = v7
            case 330:
                v8 := []int{p1, -9223372036854775808, p1, p1}
                _ = v8
            default:
                if false {
                    v9 := []int{p1, 639, p1}
                    _ = len(v9)
                    _ = v9
                } else {
                    p0 = (p0 | (p0 << ((18446744073709551615) & 63)))
                }
            }
        case "\n\t":
            var v10 float32 = 3.14159
            _ = v10
        }
    }
    var v11 uint8 = 42
    var v12 float32 = (1.5 * (1.5 * 3.14159))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return 181
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    fmt.Println(v1)
    var v13 uint16 = 183
    v1 = (^(v1))
    var v14 int = 4294967295
    v15 := GenS1{F0: 340, F1: (!((float32(1.5) < (3.14159 - (1e10 + -2.25))) || true))}
    var v16 string = ("\n\t" + (("hello" + "hello") + ("a" + "")))
    fmt.Println(v13, v1, v1)
    v13 -= v13
    v17 := GenS2{F0: 2147483648, F1: 279}
    v13 = v13
    v1 = -837
    v13 = v13
    fmt.Println(v14, v1)
    var v18 uint64 = 4294967295
    var v19 float32 = (0.1 / 1e10)
    delete(v3, 4294967295)
    v16 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
