package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint64; F1 int; F2 int32 }

func main() {
    var v1 uint16 = 0
    var v2 bool = false
    var v3 int16 = -981
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: 127, F1: v1}
    v6 := GenS1{F0: 128, F1: 65535}
    v1 = 641
    var v7 uint8 = 1
    switch "a" {
    case "foo bar":
        var v8 uint16 = (v1 * 65535)
        _ = v8
    }
    switch 715 {
    case 128:
        var v9 uint = 18446744073709551615
        _ = v9
    }
    var v10 int32 = -1
    _ = len(v4)
    if v2 {
        v11 := GenS2{F0: 2147483648, F1: true}
        for fi12 := 0; fi12 < 15; fi12++ {
            switch (((("foo bar" + ("foo bar" + "\n\t")) + "") + "\n\t") + ("a" + "a")) {
            case "a":
                var v13 float32 = -2.25
                _ = v13
            }
            v14 := GenS3{F0: 0, F1: 2147483647, F2: 42}
            _ = fi12
            _ = v14
        }
        _ = v11
    }
    var v15 int = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
