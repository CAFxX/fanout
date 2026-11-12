package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 bool; F2 string }

func main() {
    var v1 int32 = 42
    var v2 float64 = (1.5 * -2.25)
    v3 := make(map[int]string)
    var v4 int8 = -128
    v5 := []int{-9223372036854775808, 1}
    v1 = (^(v1))
    v6 := GenS1{F0: ("hello" + "")}
    v4 = v4
    switch "" {
    case "hello":
        delete(v3, 775)
    }
    var v7 uint8 = 173
    v8 := GenS2{F0: v4, F1: ("hello" > ("a" + ("foo bar" + ("" + (("hello" + "\n\t") + "hello"))))), F2: "\n\t"}
    var v9 uint32 = 2
    v10 := GenS2{F0: int8(v9), F1: (!(113 == 2)), F2: ""}
    v4 = v4
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v3)
    if (!(((true && ((v2 >= (1e10 * v2)) && (true && ((("a" < "foo bar") && (-1 < -1)) && (1 < (v7 - 42)))))) || true) || false)) {
        var v11 int64 = 924
        var v12 float32 = 3.14159
        v1 = v1
        _ = v11
        _ = v12
    }
    var v13 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
