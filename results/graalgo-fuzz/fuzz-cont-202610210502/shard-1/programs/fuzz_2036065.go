package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 uint32 = 517
    var v4 uint32 = v3
    v5 := []string{("a" + "hello"), "hello", "\n\t", "a"}
    var v6 uint8 = 128
    if (!((false && ((true && false) && false)) && true)) {
        switch "foo bar" {
        case "hello":
            if false {
                v3 = 72
            }
        case "\n\t":
            fmt.Println(v4, v3)
        }
    }
    var v7 int8 = 109
    for fi8 := 0; fi8 < 11; fi8++ {
        v6 = 0
        _ = fi8
    }
    v9 := GenS1{F0: (v6 | v6)}
    var v10 string = ("a" + ((((("\n\t" + "foo bar") + "") + "a") + ("a" + ((("a" + ("a" + "\n\t")) + "hello") + (((("foo bar" + ((("foo bar" + "hello") + "foo bar") + "")) + (((("foo bar" + "\n\t") + "") + "foo bar") + ("a" + ("\n\t" + "a")))) + ("foo bar" + "")) + (("hello" + "\n\t") + ("foo bar" + ("foo bar" + (("" + "") + "foo bar")))))))) + "foo bar"))
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 1; fi12++ {
        v4 += 42
        fmt.Println(v10, v3, v6)
        v13 := v1[430]
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    var v14 uint16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
