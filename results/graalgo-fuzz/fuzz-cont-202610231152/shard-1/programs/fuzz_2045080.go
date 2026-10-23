package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 42
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    v5 := make(map[string]int)
    if ((0 > 0) || false) {
        var v6 int64 = (v2 / ((v2) | 1))
        for fi7 := 0; fi7 < 4; fi7++ {
            delete(v1, 227)
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 42, F1: 42, F2: 65535}
                v10 := []int{(-(fi8)), (fi7 + fi7)}
                var v11 int32 = -1
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := []int{299}
            _ = fi7
            _ = v12
        }
        _ = v6
    } else {
        v2 += (-(v2))
        v13 := GenS1{F0: 112, F1: 42, F2: 255}
        _ = v13
    }
    var v14 int64 = v2
    var v15 int = 42
    v16 := v5["a"]
    v2 = (^(v2))
    v16 = v16
    v14 = (v14 << ((4294967295) & 63))
    switch (v16 << ((933) & 63)) {
    case 42:
        v16 -= v16
    default:
        var v17 float32 = float32(3.14159)
        _ = v17
    }
    v16 += (v15 << ((65535) & 63))
    switch ("a" + ("a" + ("" + "hello"))) {
    case "hello":
        if true {
            for fi18 := 0; fi18 < 16; fi18++ {
                v2 = v2
                v19 := []int{(fi18 << ((572) & 63)), (^(v16)), (v15 >> ((0) & 63)), (v16 * 1)}
                var v20 int16 = 205
                _ = fi18
                _ = v19
                _ = v20
            }
            switch "hello" {
            case "foo bar":
                var v21 int16 = 1
                _ = v21
            }
        }
    case "\n\t":
        v22 := GenS1{F0: 18446744073709551615, F1: -1, F2: 255}
        _ = v22
    default:
        var v23 uint64 = 4294967295
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
