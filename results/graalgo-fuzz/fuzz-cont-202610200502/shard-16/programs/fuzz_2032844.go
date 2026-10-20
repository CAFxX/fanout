package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -1
    var v3 int64 = 42
    v3 = (^(v3))
    var v4 uint16 = 538
    fmt.Println(v3)
    fmt.Println(v2, v2, v2)
    if (v4 >= v4) {
        var v5 string = "hello"
        _ = v5
    } else {
        v6 := GenS1{F0: v2}
        _ = v6
    }
    v7 := GenS1{F0: v2}
    v8 := GenS1{F0: (v2 * v2)}
    for fi9 := 0; fi9 < 16; fi9++ {
        v4 += v4
        var v10 bool = (52 != 255)
        _ = fi9
        _ = v10
    }
    var v11 int32 = -736
    v12 := []int{2147483647}
    var v13 uint8 = 0
    for fi14 := 0; fi14 < 2; fi14++ {
        v3 += (v3 ^ (v3 | 0))
        _ = fi14
    }
    if (false && (0.1 <= ((0.1 + 0.1) * float64((1.5 + 1.5))))) {
        fmt.Println(v11, v3)
        fmt.Println(v4, v13, v13)
        switch ("foo bar" + "foo bar") {
        case "hello":
            var v15 int64 = int64(v11)
            _ = v15
        case "":
            v16 := v1[149]
            _ = v16
        }
    } else {
        var v17 int8 = 1
        if ((^(v13)) <= (v13 << ((247) & 63))) {
            v18 := GenS1{F0: (v2 / ((629) | 1))}
            v2 = 32767
            delete(v1, 0)
            _ = v18
        } else {
            var v19 uint16 = v4
            _ = v19
        }
        _ = v17
    }
    v11 -= 976
    if (true || ("hello" >= "\n\t")) {
        var v20 int = v2
        _ = v20
    }
    if (0.0 != 0.1) {
        var v21 int32 = -1
        if (((v3 + (^(v3))) < 1) || false) {
            var v22 bool = ((v3 | 0) >= v3)
            _ = v22
        }
        _ = v21
    }
    v23 := GenS1{F0: 128}
    v13 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
