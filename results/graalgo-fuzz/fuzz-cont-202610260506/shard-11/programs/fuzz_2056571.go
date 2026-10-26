package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 int = 32768
    var v2 uint64 = 18446744073709551615
    var v3 uint64 = v2
    v4 := make(map[string]string)
    var v5 string = "foo bar"
    var v6 float64 = 1e10
    delete(v4, v5)
    var v7 string = "hello"
    var v8 int8 = 42
    if false {
        var v9 float32 = 0.0
        var v10 float32 = (1e10 + 1.5)
        _ = v10
        _ = v9
    } else {
        var v11 int = (v1 % ((-9223372036854775808) | 1))
        _ = v11
    }
    v12 := v4[(v5 + v5)]
    v8 = v8
    if (true || true) {
        v13 := GenS1{F0: (v2 | v2), F1: -611}
        for fi14 := 0; fi14 < 11; fi14++ {
            fmt.Println(v1, fi14, v12)
            _ = fi14
        }
        var v15 float64 = v6
        _ = v13
        _ = v15
    } else {
        var v16 int = (v1 >> ((54) & 63))
        _ = len(v4)
        var v17 uint64 = (v2 ^ 18446744073709551615)
        _ = v16
        _ = v17
    }
    fmt.Println(v12, v7, v2)
    fmt.Println(v3)
    fmt.Println(v5)
    if (65535 != 2147483647) {
        v3 = v3
        v8 -= 1
    } else {
        var v18 bool = ((10 < (v8 / ((v8) | 1))) && (("hello" + "foo bar") > v5))
        if (!v18) {
            v19 := GenS1{F0: v2, F1: -1}
            v7 = v12
            _ = v19
        } else {
            var v20 uint32 = 4294967295
            switch (("foo bar" + v7) + v12) {
            case "a":
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 string = ("foo bar" + ("" + v5))
                    _ = fi21
                    _ = v22
                }
            }
            v4[(v5 + v7)] = v7
            _ = v20
        }
        _ = v18
    }
    var v23 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
