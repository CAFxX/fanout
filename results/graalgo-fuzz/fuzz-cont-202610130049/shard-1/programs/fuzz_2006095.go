package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint }

func main() {
    v1 := []string{"", "a", (((((("" + "") + ("\n\t" + ("" + "foo bar"))) + "hello") + "\n\t") + ("foo bar" + "\n\t")) + ("hello" + "foo bar"))}
    v2 := []uint32{42, 2, 65535, 285}
    var v3 int64 = -1
    v4 := make(map[string]uint32)
    var v5 uint = 255
    v6 := make(map[int]int)
    v3 += v3
    var v7 int = int(v3)
    if len(v2) > 0 {
        v2[0] = 807
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint8 = uint8(v5)
        _ = fi8
        _ = v9
    }
    switch (((("hello" + "\n\t") + "a") + "") + "\n\t") {
    case "\n\t":
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 float64 = 3.14159
            v12 := GenS1{F0: 367, F1: 0, F2: (v5 * v5)}
            _ = fi10
            _ = v11
            _ = v12
        }
    case "a":
        fmt.Println(v3, v3, v5)
    case "foo bar":
        var v13 int32 = 0
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v15 := GenS1{F0: 298, F1: 361, F2: v5}
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: 2147483648, F1: 42, F2: 65535}
    fmt.Println(v5, v3)
    for fi17 := 0; fi17 < 10; fi17++ {
        v5 += v5
        delete(v4, ("a" + "a"))
        _ = fi17
    }
    v5 = (v5 % ((255) | 1))
    fmt.Println(v3, v5, v7)
    fmt.Println(v7)
    v18 := GenS1{F0: 2, F1: 2, F2: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
