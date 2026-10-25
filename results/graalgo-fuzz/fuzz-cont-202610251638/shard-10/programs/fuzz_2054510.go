package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (1.5 - float32(1.5))
    v3 := []string{("foo bar" + ("foo bar" + "\n\t")), (("hello" + ("" + "hello")) + "a"), "hello"}
    var v4 uint8 = 167
    v5 := make(map[int]int)
    var v6 int32 = 42
    v4 = (v4 % (((v4 | 42)) | 1))
    _ = len(v3)
    switch (("\n\t" + "a") + ("a" + ("a" + "foo bar"))) {
    case "hello":
        v2 += ((v2 * v2) / (v2 + (1.5 + 1e10)))
    }
    fmt.Println(v4, v4, v4)
    v5[1] = int(v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        v2 = float32(float64(v2))
        fmt.Println(v6, v2, v2)
        _ = fi7
    }
    v3 = append(v3, "hello")
    v8 := GenS1{F0: 42, F1: int64(v6)}
    var v9 int8 = 100
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = (v4 - v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
