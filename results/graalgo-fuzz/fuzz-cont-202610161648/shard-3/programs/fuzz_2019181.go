package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint = 983
    var v2 uint64 = 18446744073709551615
    v3 := make(map[int]int)
    v4 := []string{(("" + "foo bar") + ("\n\t" + "\n\t"))}
    var v5 string = ""
    v6 := make(map[int]bool)
    var v7 int64 = -642
    fmt.Println(v5)
    v5 = v5
    v8 := GenS1{F0: (("" != "foo bar") && (0 != 32767))}
    if (false && (true && (!(-32768 <= 42)))) {
        v1 = 254
        var v9 uint16 = 2
        _ = v9
    }
    v5 = "\n\t"
    v4 = append(v4, "")
    v1 = (v1 | 255)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
