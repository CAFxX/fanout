package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = ("a" + ("foo bar" + "a"))
    v3 := make(map[string]uint32)
    var v4 int = 2147483647
    v5 := []uint32{4294967295, 583, 2147483648}
    v1[2147483648] = (v4 & v4)
    fmt.Println(v2, v4)
    v6 := GenS1{F0: -135, F1: 0, F2: 42}
    var v7 int16 = -1
    switch v2 {
    case "":
        switch (^(v4)) {
        case 643:
            var v8 uint = 255
            _ = v8
        case 2147483647:
            for fi9 := 0; fi9 < 3; fi9++ {
                if (-1 == 733) {
                    v7 = (v7 * v7)
                    var v10 bool = true
                    _ = len(v1)
                    _ = v10
                } else {
                    v11 := GenS1{F0: (v4 << ((67) & 63)), F1: 42, F2: 255}
                    v12 := GenS1{F0: fi9, F1: 1, F2: 230}
                    fmt.Println(v2)
                    _ = v11
                    _ = v12
                }
                v4 = fi9
                fmt.Println(v7, v2)
                _ = fi9
            }
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int64 = 9223372036854775807
        _ = fi13
        _ = v14
    }
    v2 = v2
    switch (v4 | v4) {
    case 127:
        fmt.Println(v4, v4, v4)
    }
    v15 := GenS2{F0: 127}
    v5 = append(v5, 65535)
    _ = len(v3)
    v16 := GenS3{F0: ((0.0 > float64(0.1)) && (!true))}
    _ = len(v3)
    fmt.Println(v2, v2, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
