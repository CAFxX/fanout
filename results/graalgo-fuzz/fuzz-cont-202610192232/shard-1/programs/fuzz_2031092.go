package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]string)
    var v3 int16 = -899
    v4 := make(map[int]bool)
    v5 := []int{127}
    var v6 int = 160
    var v7 int64 = -9223372036854775808
    switch ("hello" + ("" + "a")) {
    case "a":
        v8 := GenS1{F0: (v7 * v7), F1: (v6 * v6)}
        _ = v8
    case "foo bar":
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: int64(v1), F1: v6}
            v3 = 42
            _ = fi9
            _ = v10
        }
    case "hello":
        if true {
            if len(v5) > 0 {
                v5[0] = 32768
            }
        }
    }
    var v11 uint32 = (^(v1))
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 uint8 = 1
        v7 -= (v7 * (v7 << ((65535) & 63)))
        var v14 int16 = (v3 | (v3 * (v3 ^ v3)))
        _ = fi12
        _ = v13
        _ = v14
    }
    fmt.Println(v6, v3, v6)
    v3 -= v3
    fmt.Println(v6)
    switch (v6 << ((2) & 63)) {
    case 880:
        delete(v2, "a")
    }
    switch v6 {
    case 128:
        v7 = v7
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
