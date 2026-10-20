package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int64; F1 string; F2 int64 }

func main() {
    var v1 uint8 = 2
    v2 := make(map[string]bool)
    var v3 int64 = -150
    v4 := make(map[string]bool)
    var v5 bool = true
    v6 := []uint32{656, 958, 323}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint32 = 2
        v5 = true
        if (2147483647 >= 0) {
            v9 := GenS1{F0: 1, F1: ("a" + "\n\t")}
            _ = v9
        } else {
            _ = len(v6)
            fmt.Println(v1)
        }
        _ = fi7
        _ = v8
    }
    if v5 {
        v5 = (!(2 <= 243))
    }
    v10 := GenS1{F0: 1, F1: ""}
    v5 = v5
    v11 := GenS1{F0: -1, F1: ("hello" + ("\n\t" + "hello"))}
    _ = len(v4)
    fmt.Println(v5, v1)
    v6 = append(v6, 2)
    v12 := GenS1{F0: -1, F1: "foo bar"}
    switch "" {
    case "\n\t":
        v13 := GenS2{F0: v3, F1: "a", F2: -9223372036854775808}
        _ = v13
    case "hello":
        v1 = (v1 << ((65535) & 63))
    case "foo bar":
        v3 += (v3 % (((v3 * v3)) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
