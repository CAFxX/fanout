package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }

func main() {
    v1 := []string{"\n\t", "hello", "a"}
    v2 := []uint32{2147483647, 42, 4}
    v3 := []uint32{2, 0}
    var v4 uint16 = 91
    v4 = 255
    var v5 uint64 = 18446744073709551615
    v2 = append(v2, 65535)
    switch 42 {
    case 32767:
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 uint8 = 127
            _ = fi6
            _ = v7
        }
    case 490:
        v5 -= v5
    }
    v8 := GenS1{F0: 1, F1: 42, F2: ""}
    var v9 float64 = (3.14159 / float64(-2.25))
    v4 = v4
    var v10 uint64 = (v5 >> ((0) & 63))
    v11 := GenS1{F0: 1, F1: 4294967295, F2: "hello"}
    var v12 int64 = 248
    if (42 <= 42) {
        var v13 float64 = v9
        var v14 uint32 = 2147483648
        _ = v13
        _ = v14
    }
    if (1e10 > ((float32(v9) * (((-2.25 / 1.0) + ((3.14159 + (1.5 - 0.1)) + -2.25)) + (-2.25 - -2.25))) / (1.5 + (1.5 * 1.5)))) {
        switch ((("foo bar" + "hello") + "") + ("foo bar" + "a")) {
        case "hello":
            for fi15 := 0; fi15 < 9; fi15++ {
                v16 := GenS1{F0: -1, F1: 2, F2: "\n\t"}
                _ = fi15
                _ = v16
            }
        }
        if len(v3) > 0 {
            v3[0] = 420
        }
    }
    var v17 uint64 = 4294967295
    v4 = v4
    if (v4 == 0) {
        if false {
            v9 = (v9 + 1.5)
        }
        fmt.Println(v9)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
