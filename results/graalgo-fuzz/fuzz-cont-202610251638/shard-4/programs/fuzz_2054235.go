package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 80
    v3 := make(map[int]string)
    var v4 int16 = -32768
    v5 := []int{2147483647, -730}
    if (0 != 65535) {
        var v6 string = "foo bar"
        v7 := v3[0]
        _ = v6
        _ = v7
    }
    _ = len(v3)
    var v8 uint64 = 18446744073709551615
    v5 = append(v5, -9223372036854775808)
    if (((float32((1e10 * 0.1)) / ((3.14159 + (-2.25 + ((0.0 + ((1.5 * 1e10) * (float32(0.1) - (-2.25 / (1.5 + 1e10))))) - ((0.0 / 1e10) / -2.25)))) + -2.25)) + -2.25) <= 0.0) {
        v9 := GenS1{F0: (86 > -1), F1: (!false)}
        v3[0] = ("a" + ("\n\t" + "\n\t"))
        _ = len(v5)
        _ = v9
    } else {
        v8 -= v8
        switch ("\n\t" + (("a" + (("\n\t" + ((("foo bar" + ("foo bar" + ("" + "foo bar"))) + "") + "")) + ("foo bar" + ("foo bar" + ("hello" + ""))))) + "a")) {
        case "a":
            var v10 int16 = (^(v4))
            _ = v10
        case "foo bar":
            _ = len(v1)
        }
    }
    v3[127] = "foo bar"
    var v11 string = ("a" + (("foo bar" + "\n\t") + (((("" + (("a" + "\n\t") + "foo bar")) + "hello") + "\n\t") + (("" + "\n\t") + "a"))))
    var v12 uint8 = 255
    var v13 uint16 = uint16(v8)
    switch -1 {
    case -1:
        var v14 uint32 = 1
        _ = v14
    case -938:
        v1["a"] = (!true)
    default:
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint = 18446744073709551615
            _ = fi15
            _ = v16
        }
    }
    v2 = v12
    v17 := GenS1{F0: (!((81 == 63) || true)), F1: (0 != 0)}
    var v18 int = 0
    v13 = (v13 % ((0) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
