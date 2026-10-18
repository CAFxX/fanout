package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 string }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    v3 := make(map[string]string)
    fmt.Println(v2)
    fmt.Println(v2)
    v4 := []int{-969, -211, 0, 656}
    v4 = append(v4, 0)
    v4 = append(v4, 4294967295)
    var v5 float32 = ((1e10 - (-2.25 * 3.14159)) / (-2.25 - ((3.14159 / 3.14159) - 3.14159)))
    switch 32767 {
    case 32767:
        var v6 uint8 = (v2 % ((127) | 1))
        _ = v6
    case 9223372036854775807:
        if false {
            v5 += float32(0.0)
            v7 := GenS1{F0: 25, F1: 1, F2: ((((("hello" + ((((("foo bar" + "foo bar") + "") + "\n\t") + (("" + "") + ("hello" + "\n\t"))) + ("foo bar" + (("hello" + "hello") + "\n\t")))) + ((("hello" + "foo bar") + "\n\t") + "\n\t")) + ("hello" + "a")) + "a") + "")}
            var v8 uint64 = 0
            _ = v7
            _ = v8
        }
    default:
        v5 += v5
    }
    switch ("a" + "\n\t") {
    case "hello":
        var v9 int16 = 0
        _ = v9
    case "":
        var v10 string = "\n\t"
        _ = v10
    }
    if false {
        v11 := GenS2{F0: 42, F1: -498}
        _ = v11
    } else {
        if len(v4) > 0 {
            v4[0] = 0
        }
        var v12 int16 = 42
        var v13 int = 0
        _ = v12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
