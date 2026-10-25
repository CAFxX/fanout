package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []uint32{507, 830}
    var v2 int64 = -181
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    var v5 uint16 = 31
    var v6 float64 = 1.5
    if (false || (!true)) {
        var v7 uint32 = 0
        var v8 string = "foo bar"
        v6 = v6
        _ = v7
        _ = v8
    }
    var v9 float32 = 0.1
    fmt.Println(v5)
    _ = len(v1)
    v10 := GenS1{F0: (v2 % ((v2) | 1)), F1: 42}
    var v11 uint32 = 2147483647
    switch (("a" + ("" + (("foo bar" + "\n\t") + "hello"))) + "foo bar") {
    case "foo bar":
        for fi12 := 0; fi12 < 14; fi12++ {
            v2 = v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v9 += (v9 + 0.0)
            _ = fi12
        }
    case "\n\t":
        v6 -= float64((1.5 / float32(float64(((v9 / (((v9 + float32(float64((v9 * float32(float64((v9 + v9))))))) * v9) * (v9 - float32(v6)))) - v9)))))
    case "hello":
        if (!false) {
            v1 = append(v1, 1)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
