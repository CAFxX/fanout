package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }

func main() {
    v1 := []string{"\n\t", "hello"}
    v2 := []int{-658, 704, 1}
    var v3 float32 = ((0.1 + ((1.5 * 1.5) - -2.25)) * (1e10 * ((-2.25 + 1.5) * 0.1)))
    v4 := make(map[string]uint32)
    var v5 bool = (false && (4294967295 < 18446744073709551615))
    fmt.Println(v5, v3)
    var v6 string = ("foo bar" + "foo bar")
    v7 := GenS1{F0: ((("hello" + ("\n\t" + ("\n\t" + ("a" + (v6 + v6))))) + "hello") + v6), F1: 1}
    v3 = v3
    switch ((v6 + (v6 + ((v6 + v6) + ("hello" + "\n\t")))) + ("\n\t" + v6)) {
    case "hello":
        var v8 int32 = -1
        _ = v8
    }
    fmt.Println(v5, v3, v3)
    var v9 int8 = 127
    fmt.Println(v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
