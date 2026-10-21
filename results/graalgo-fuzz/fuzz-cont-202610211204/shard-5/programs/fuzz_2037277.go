package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 467
    var v3 uint16 = 255
    var v4 int16 = 200
    v5 := GenS1{F0: 18446744073709551615, F1: 0, F2: (((3.14159 / 1.0) == (0.1 - -2.25)) || false)}
    var v6 uint = 18446744073709551615
    var v7 bool = false
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 int32 = 0
        if v7 {
            fmt.Println(fi8, v3)
        } else {
            fmt.Println(v2, v3)
            if ("a" > ("" + "\n\t")) {
                var v10 uint16 = v3
                v2 -= 889
                _ = v10
            }
        }
        _ = fi8
        _ = v9
    }
    v2 = 337
    v6 = 42
    v6 = (v6 | (v6 | v6))
    v11 := v1[v2]
    v4 -= -487
    var v12 int16 = -32768
    v12 = 32767
    var v13 int8 = int8(v2)
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 int16 = v4
        fmt.Println(v15)
        if (255 < v3) {
            v16 := GenS2{F0: fi14, F1: (v2 & 0)}
            v3 = (v3 >> ((183) & 63))
            var v17 string = "a"
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = 18446744073709551615
            v4 = v4
            fmt.Println(v3)
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    v19 := GenS3{F0: 100, F1: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
