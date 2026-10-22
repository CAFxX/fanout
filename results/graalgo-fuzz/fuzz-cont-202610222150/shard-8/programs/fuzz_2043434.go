package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 string; F1 uint64; F2 bool }
func genFn0() uint16 {
    var v6 int32 = -1
    v7 := []int{203, 32768}
    var v8 float32 = float32(0.1)
    var v9 int = 2147483647
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 string = "a"
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    var v4 uint8 = 134
    v5 := make(map[uint32]int)
    v1 = (("a" + "hello") + v1)
    fmt.Println(v1, v2, v4)
    v10 := v5[1]
    switch (v10 | v10) {
    case -1000:
        var v11 int32 = -2147483648
        _ = v11
    }
    var v12 int = (^(v10))
    v12 += (v12 & v12)
    if (!(!false)) {
        v5[42] = 32767
    } else {
        var v13 int16 = -304
        switch (v1 + (v1 + "")) {
        case "\n\t":
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS1{F0: v1}
                _ = fi14
                _ = v15
            }
        }
        switch v1 {
        case "foo bar":
            if (272 <= 2) {
                for fi16 := 0; fi16 < 20; fi16++ {
                    switch (fi16 - (v10 * (v12 + 128))) {
                    case 42:
                        var v17 string = "hello"
                        _ = v17
                    case 32768:
                        var v18 int32 = -2147483648
                        _ = v18
                    default:
                        fmt.Println(v13, v12)
                    }
                    _ = fi16
                }
                fmt.Println(v13)
            }
        case "hello":
            fmt.Println(v2, v1)
        default:
            _ = len(v3)
        }
        _ = v13
    }
    v19 := []int{v10}
    v20 := GenS1{F0: ""}
    fmt.Println(v12, v4)
    v4 = v4
    v12 = v10
    v21 := GenS2{F0: 2147483647, F1: 127}
    var v22 uint16 = 1
    var v23 int32 = -1
    if (true || true) {
        var v24 float64 = ((float64(1e10) / (1.5 * (1.5 * (-2.25 - 0.0)))) - 0.1)
        if false {
            v25 := GenS3{F0: v1, F1: 42, F2: false}
            var v26 int16 = 1000
            _ = v25
            _ = v26
        } else {
            fmt.Println(v10)
            v2 = (v2 / ((v2) | 1))
            v10 = (v10 << ((v2) & 63))
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 uint32 = 2
            var v29 uint64 = 2
            var v30 uint64 = v29
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        _ = v24
    }
    var v31 uint = (v2 & 255)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
