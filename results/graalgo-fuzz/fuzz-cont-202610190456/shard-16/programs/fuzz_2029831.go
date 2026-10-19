package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 uint = 359
    v2 := []string{((("\n\t" + "\n\t") + "foo bar") + ""), (("\n\t" + ((("hello" + ((("hello" + (("" + "hello") + ("\n\t" + ""))) + "foo bar") + ("hello" + ("foo bar" + ("a" + (("a" + "a") + "foo bar")))))) + ("" + "")) + "foo bar")) + ("\n\t" + "")), "a"}
    var v3 int16 = 1000
    v4 := GenS1{F0: 0, F1: 1}
    v5 := GenS1{F0: v3, F1: 132}
    var v6 int64 = 9223372036854775807
    if ((!false) && true) {
        v7 := GenS2{F0: (^(v1))}
        _ = len(v2)
        var v8 int = 32767
        _ = v7
        _ = v8
    }
    v1 += (^(v1))
    var v9 int = 42
    v1 = v1
    v2 = append(v2, "a")
    v3 -= int16(v1)
    var v10 uint8 = 0
    var v11 int32 = 234
    var v12 int16 = int16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
