package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []string{"", ("hello" + ("" + "")), ""}
    v2 := make(map[uint32]string)
    var v3 uint32 = 0
    var v4 uint16 = 1
    var v5 uint64 = 660
    v6 := GenS1{F0: v3, F1: int64(v5), F2: 18446744073709551615}
    v5 -= v5
    v7 := GenS2{F0: v3, F1: -167}
    if false {
        fmt.Println(v4)
        var v8 uint = 255
        _ = v8
    }
    if ((1 == 255) && true) {
        v4 = 42
    } else {
        var v9 int64 = 1
        switch 525 {
        case -9223372036854775808:
            v10 := GenS2{F0: v3, F1: -419}
            _ = v10
        }
        _ = v9
    }
    v5 += (v5 % (((v5 | 767)) | 1))
    switch ("" + "\n\t") {
    case "":
        for fi11 := 0; fi11 < 5; fi11++ {
            v5 += v5
            v3 = (v3 / ((v3) | 1))
            fmt.Println(v5, v4)
            _ = fi11
        }
    }
    var v12 uint16 = v4
    var v13 int = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
