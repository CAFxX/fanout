package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    var v4 string = (("a" + (("\n\t" + "\n\t") + "hello")) + ("a" + "hello"))
    var v5 uint32 = v1
    fmt.Println(v4)
    fmt.Println(v3, v5)
    fmt.Println(v1)
    if false {
        var v6 uint16 = 857
        v1 = v1
        _ = v6
    }
    fmt.Println(v1, v5)
    v7 := v2[9223372036854775807]
    fmt.Println(v5)
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint64 = 4294967295
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 127
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v3 = v3
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v5, v4, v10)
    v2[int(v5)] = 2
    v13 := GenS1{F0: 100, F1: 42}
    var v14 uint8 = (v10 << ((65535) & 63))
    switch "" {
    case "foo bar":
        if (((128 > 2147483647) && (((v3 / float64(-2.25)) * (v3 / -2.25)) >= 1e10)) || false) {
            switch "foo bar" {
            case "":
                var v15 int8 = int8(v14)
                _ = v15
            }
        }
    default:
        var v16 float64 = float64(1.5)
        _ = v16
    }
    if ((v7 + (v7 * 1)) > v5) {
        v5 = v1
    }
    v17 := GenS1{F0: 127, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
