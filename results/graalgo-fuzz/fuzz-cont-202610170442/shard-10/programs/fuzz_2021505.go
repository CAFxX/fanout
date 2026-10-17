package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 uint32 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[string]string)
    var v3 int = 379
    v4 := make(map[string]bool)
    v5 := make(map[int]bool)
    var v6 int16 = 0
    var v7 uint16 = 687
    v3 = 42
    v7 = (v7 % ((650) | 1))
    var v8 bool = (false && (!(1 <= 684)))
    if (!v8) {
        var v9 uint8 = 42
        var v10 int = 32768
        _ = v10
        _ = v9
    }
    switch -2147483648 {
    case 42:
        for fi11 := 0; fi11 < 16; fi11++ {
            if (!(3.14159 < ((0.1 / 1.0) + ((1e10 - -2.25) + 0.0)))) {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                v12 := GenS1{F0: 105, F1: 65535}
                _ = v12
            } else {
                v6 = v6
                var v13 int16 = (v6 / ((v6) | 1))
                v5[0] = (31 == 84)
                _ = v13
            }
            _ = len(v2)
            _ = fi11
        }
    case 4294967295:
        fmt.Println(v8, v7, v7)
    default:
        fmt.Println(v6)
    }
    var v14 uint64 = 2
    if (42 < 108) {
        var v15 uint64 = v14
        _ = v15
    }
    var v16 uint = 1
    v17 := GenS1{F0: 42, F1: 42}
    v18 := GenS2{F0: ((("a" + "foo bar") + (("foo bar" + ("" + "")) + (("a" + "") + "\n\t"))) + "\n\t"), F1: 0, F2: 65535}
    switch (("foo bar" + "hello") + "foo bar") {
    case "\n\t":
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    var v20 uint64 = 0
    v20 += uint64(v16)
    if v8 {
        switch "" {
        case "foo bar":
            v21 := GenS2{F0: "a", F1: 0, F2: 193}
            _ = v21
        case "\n\t":
            switch (("\n\t" + "") + "\n\t") {
            case "":
                v1 = append(v1, "a")
            case "a":
                v14 -= (v14 << (((v16 % ((1) | 1))) & 63))
            default:
                v7 = v7
            }
        case "hello":
            v22 := GenS2{F0: "\n\t", F1: 901, F2: 65535}
            _ = v22
        }
    }
    var v23 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
