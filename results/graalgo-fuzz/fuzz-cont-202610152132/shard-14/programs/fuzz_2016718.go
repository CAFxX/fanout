package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := []string{"", ""}
    var v2 uint = 18446744073709551615
    var v3 int32 = 0
    var v4 int8 = 0
    v4 = v4
    v5 := GenS1{F0: 0, F1: 42}
    v6 := GenS1{F0: 112, F1: 99}
    v7 := GenS1{F0: 4294967295, F1: 42}
    fmt.Println(v2)
    switch ((("hello" + (("" + "hello") + "\n\t")) + "a") + ("\n\t" + "a")) {
    case "foo bar":
        _ = len(v1)
    }
    _ = len(v1)
    v1 = append(v1, ((("" + "\n\t") + "hello") + ("\n\t" + ("\n\t" + ("" + (("hello" + "") + "a"))))))
    v8 := GenS1{F0: 4294967295, F1: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
