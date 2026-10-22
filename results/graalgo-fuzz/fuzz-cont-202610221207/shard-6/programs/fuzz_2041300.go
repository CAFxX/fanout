package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 857
    v2 := make(map[int]uint32)
    var v3 int16 = 42
    v4 := make(map[int]uint32)
    v5 := []string{("hello" + "a"), (("hello" + "a") + "\n\t"), ((("a" + "\n\t") + "\n\t") + ("hello" + (("\n\t" + "hello") + ("hello" + "\n\t"))))}
    var v6 float32 = 3.14159
    var v7 int64 = -1
    var v8 uint8 = 127
    v1 -= (v1 * (^(v1)))
    v8 = 128
    delete(v2, 673)
    fmt.Println(v6)
    v7 -= (v7 ^ (v7 * (v7 ^ -511)))
    var v9 string = (((("" + ("a" + "hello")) + ("" + ("foo bar" + ""))) + ("hello" + (((("a" + "foo bar") + "a") + "a") + (("hello" + "a") + "foo bar")))) + "\n\t")
    switch 221 {
    case 2147483647:
        v3 -= 42
    case 0:
        var v10 uint32 = 2147483648
        _ = v10
    }
    fmt.Println(v8, v1, v8)
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v3)
        _ = fi11
    }
    v12 := GenS1{F0: v9}
    var v13 bool = (!(true && ((v1 - v1) != v1)))
    var v14 int16 = (v3 | v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
