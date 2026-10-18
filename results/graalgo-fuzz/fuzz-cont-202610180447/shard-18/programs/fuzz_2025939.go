package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 string }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint = 2
    v3 := make(map[int]bool)
    if (false && false) {
        var v4 uint = v2
        v5 := GenS1{F0: false, F1: ("hello" + ""), F2: ""}
        _ = v4
        _ = v5
    } else {
        var v6 float64 = 3.14159
        _ = v6
    }
    switch "\n\t" {
    case "":
        var v7 int = 127
        _ = v7
    case "a":
        v2 -= v2
    default:
        v2 = v2
    }
    v2 = (v2 - v2)
    var v8 int64 = 9223372036854775807
    var v9 uint64 = 2
    v10 := GenS1{F0: ((!true) && false), F1: "hello", F2: ("hello" + "a")}
    v11 := v3[2147483647]
    fmt.Println(v2)
    switch -9223372036854775808 {
    case -2147483648:
        v11 = (!(3.14159 < 0.0))
    case 0:
        v12 := GenS1{F0: (255 != 498), F1: ((("a" + "hello") + (("hello" + "") + (("" + "a") + ("hello" + "\n\t")))) + ((("hello" + "\n\t") + "") + "a")), F2: "\n\t"}
        _ = v12
    default:
        var v13 float32 = 1e10
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
