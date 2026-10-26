package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 int; F2 int }

func main() {
    var v1 uint16 = 65535
    v2 := []int{-681, 1, 4294967295, 2147483647}
    v3 := make(map[uint32]bool)
    var v4 string = "a"
    if (9223372036854775807 == -1) {
        v5 := GenS1{F0: int64(v1)}
        v1 = 0
        _ = v5
    } else {
        var v6 uint8 = 51
        switch ((v4 + "hello") + v4) {
        case "foo bar":
            for fi7 := 0; fi7 < 3; fi7++ {
                v8 := GenS2{F0: -1}
                v4 = v4
                v9 := GenS2{F0: -2147483648}
                _ = fi7
                _ = v8
                _ = v9
            }
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int16 = 42
            _ = fi10
            _ = v11
        }
        _ = v6
    }
    v4 = ("hello" + "\n\t")
    delete(v3, 2)
    fmt.Println(v1, v1)
    v12 := GenS3{F0: -912, F1: 128, F2: 42}
    switch 38 {
    case 32768:
        v13 := GenS2{F0: 2147483647}
        _ = v13
    case -9223372036854775808:
        fmt.Println(v4, v4)
    case 32767:
        v4 = (v4 + ("foo bar" + "\n\t"))
    default:
        v14 := GenS2{F0: -680}
        _ = v14
    }
    var v15 bool = ((-2.25 <= 1.5) && false)
    fmt.Println(v1)
    switch "a" {
    case "":
        if (-2.25 <= (0.1 - (0.1 + (1e10 / 0.1)))) {
            v15 = v15
            var v16 float64 = ((3.14159 / 0.1) - (((0.1 - 1.5) / 3.14159) - (0.1 - (1.5 + 3.14159))))
            _ = v16
        }
    case "hello":
        var v17 int = 128
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
