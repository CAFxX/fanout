package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[uint32]int)
    v3 := []string{"a", ("\n\t" + ("" + "\n\t")), ("foo bar" + "\n\t"), ("\n\t" + "a")}
    v4 := make(map[int]string)
    v5 := make(map[string]uint32)
    var v6 int16 = -975
    var v7 uint32 = 42
    v1 = -221
    v8 := GenS1{F0: ((v7 ^ 2) > 2147483647), F1: uint8(v7), F2: 947}
    if (179 <= 549) {
        v9 := GenS1{F0: true, F1: 178, F2: 2147483648}
        _ = v9
    } else {
        v10 := GenS2{F0: 255, F1: 42}
        _ = v10
    }
    var v11 int = -2147483648
    delete(v2, v7)
    for fi12 := 0; fi12 < 17; fi12++ {
        v1 -= v1
        _ = fi12
    }
    switch 2147483647 {
    case 1:
        v11 += (v11 + v11)
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        delete(v4, fi13)
        _ = fi13
    }
    switch "hello" {
    case "":
        switch (v11 % ((v11) | 1)) {
        case 42:
            var v14 uint = 0
            _ = v14
        }
    case "\n\t":
        var v15 uint32 = v7
        _ = v15
    case "foo bar":
        var v16 int = (v11 | (v11 ^ (v11 * (v11 | v11))))
        _ = v16
    }
    var v17 uint8 = 143
    v11 = (-(v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
