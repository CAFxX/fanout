package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }

func main() {
    var v1 int64 = 86
    v2 := make(map[string]uint32)
    var v3 bool = false
    v4 := make(map[string]int)
    var v5 bool = v3
    var v6 uint = 42
    v7 := GenS1{F0: "a", F1: -758}
    v6 -= (v6 << ((v6) & 63))
    fmt.Println(v5, v3, v3)
    _ = len(v4)
    v6 -= (v6 >> (((v6 & v6)) & 63))
    v1 = -737
    var v8 int16 = -32768
    if (("hello" != ((("hello" + "") + "") + ("foo bar" + "hello"))) || ((-2147483648 >= 32768) && v5)) {
        var v9 int32 = 2147483647
        v1 -= 721
        _ = v9
    }
    var v10 int8 = 0
    v6 = (v6 & v6)
    switch ("\n\t" + ("foo bar" + "")) {
    case "foo bar":
        v8 = (-(v8))
    case "":
        v5 = (!(1e10 <= -2.25))
    case "a":
        fmt.Println(v10)
    }
    if v3 {
        if false {
            v1 = -196
        } else {
            fmt.Println(v8, v1)
        }
    } else {
        var v11 bool = v5
        switch ("" + "\n\t") {
        case "\n\t":
            var v12 float32 = 0.1
            _ = v12
        case "a":
            v3 = v3
        case "hello":
            var v13 int16 = v8
            _ = v13
        }
        _ = v11
    }
    var v14 string = ("foo bar" + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
