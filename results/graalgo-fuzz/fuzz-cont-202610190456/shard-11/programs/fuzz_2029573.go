package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }

func main() {
    var v1 bool = true
    v2 := []int{-587, 2147483647}
    v3 := []uint32{864, 65535, 769, 261}
    var v4 uint16 = 32767
    _ = len(v3)
    switch "foo bar" {
    case "":
        var v5 uint64 = 42
        _ = v5
    case "foo bar":
        v1 = ((!true) || true)
    }
    if (v1 && (42 <= 18)) {
        var v6 int = -585
        switch "foo bar" {
        case "a":
            var v7 float32 = (1e10 / (1.5 / 3.14159))
            _ = v7
        case "foo bar":
            if v1 {
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint16 = 1
                    _ = fi8
                    _ = v9
                }
                v6 -= (v6 % ((v6) | 1))
                var v10 uint = 2
                _ = v10
            }
        }
        v11 := GenS1{F0: -1, F1: ""}
        _ = v11
        _ = v6
    }
    var v12 int32 = 1
    if (-1 == 32767) {
        v4 -= (v4 | 65535)
        var v13 int64 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v13
    }
    v12 -= (v12 * v12)
    v12 += (v12 >> ((42) & 63))
    var v14 int32 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
