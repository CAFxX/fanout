package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = 0
    var v2 int64 = 42
    var v3 int8 = 45
    var v4 int = 2147483647
    v5 := make(map[uint32]string)
    var v6 int16 = 586
    v2 = (-(v2))
    var v7 int16 = (v6 | (v6 ^ 0))
    v8 := GenS1{F0: ("foo bar" + (("" + ("foo bar" + ("foo bar" + (("" + "\n\t") + ("a" + ""))))) + ("hello" + ((("foo bar" + "hello") + ("foo bar" + "foo bar")) + (("\n\t" + "") + "\n\t")))))}
    fmt.Println(v4, v3, v6)
    var v9 int64 = 9223372036854775807
    fmt.Println(v7, v6, v7)
    v10 := GenS1{F0: ""}
    v4 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
