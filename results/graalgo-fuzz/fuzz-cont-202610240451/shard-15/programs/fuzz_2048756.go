package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{(("\n\t" + ("a" + ("\n\t" + "\n\t"))) + (("hello" + "\n\t") + "\n\t")), "foo bar", "\n\t", "hello"}
    var v2 uint16 = 784
    var v3 uint32 = 1
    v4 := []string{}
    v5 := make(map[string]string)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    switch "hello" {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v6 int64 = -1
    var v7 uint = 18446744073709551615
    v8 := GenS1{F0: (v2 & 42)}
    var v9 uint64 = 4294967295
    var v10 float64 = (0.0 - -2.25)
    v3 -= 567
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
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
