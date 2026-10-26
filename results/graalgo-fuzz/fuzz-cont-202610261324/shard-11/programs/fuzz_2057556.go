package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0() string {
    v6 := GenS1{F0: "\n\t", F1: 206}
    var v7 float64 = -2.25
    fmt.Println(v7)
    if (float64(((1.5 * (1e10 * ((-2.25 / 1e10) - (3.14159 * 3.14159)))) + (1e10 - 1e10))) >= float64((0.0 * 0.1))) {
        var v8 int8 = -88
        var v9 uint8 = 255
        _ = v8
        _ = v9
    } else {
        fmt.Println(v7)
    }
    fmt.Println(v7)
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint8 = 127
    var v2 uint16 = 988
    v3 := []uint32{0}
    v4 := []string{"a", "foo bar"}
    var v5 bool = (v1 < (^(v1)))
    v5 = v5
    var v10 float32 = ((-2.25 + -2.25) / 3.14159)
    v3 = append(v3, 2147483647)
    var v11 uint8 = (v1 % ((2) | 1))
    switch "a" {
    case "\n\t":
        v5 = v5
    }
    var v12 bool = false
    var v13 bool = ((0.0 + float64(v10)) == ((float64(v10) * (1e10 - 0.1)) * (0.1 - (0.1 / 1.5))))
    v14 := GenS1{F0: ("a" + ""), F1: 42}
    fmt.Println(v5, v2, v12)
    switch 2147483647 {
    case 4294967295:
        if false {
            v15 := GenS1{F0: ("foo bar" + "\n\t"), F1: 0}
            _ = v15
        }
    case 2147483647:
        var v16 int32 = -2147483648
        _ = v16
    }
    if v5 {
        v17 := GenS1{F0: (("foo bar" + ("" + "\n\t")) + ((("" + (("\n\t" + "hello") + "hello")) + (((("hello" + "hello") + "foo bar") + "foo bar") + "hello")) + ((("\n\t" + "a") + (("a" + "hello") + "")) + ("a" + "a")))), F1: 2147483647}
        fmt.Println(v11, v12)
        _ = len(v3)
        _ = v17
    } else {
        var v18 uint64 = 4294967295
        v19 := GenS2{F0: 65535}
        _ = v18
        _ = v19
    }
    var v20 float64 = 1.5
    v12 = true
    fmt.Println(v2, v5, v13)
    var v21 uint16 = 65535
    var v22 uint64 = 42
    v23 := GenS3{F0: 42, F1: (^(v21)), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
