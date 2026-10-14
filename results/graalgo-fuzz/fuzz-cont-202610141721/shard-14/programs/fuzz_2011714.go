package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{0, 42, 1, 1}
    v3 := []uint32{0, 65535}
    v4 := make(map[string]uint32)
    _ = len(v1)
    v5 := GenS1{F0: true}
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            var v7 int32 = 2147483647
            _ = v7
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 float64 = (3.14159 * 3.14159)
                _ = fi8
                _ = v9
            }
            v10 := v1[2147483648]
            for fi11 := 0; fi11 < 18; fi11++ {
                v10 = ("a" + (("a" + "") + v10))
                fmt.Println(v10)
                v10 = v10
                _ = fi11
            }
            _ = v10
        }
        _ = fi6
    }
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            var v12 float32 = (1e10 / (0.1 - -2.25))
            _ = v12
        }
    case "a":
        v13 := GenS2{F0: 645}
        _ = v13
    }
    switch 42 {
    case -506:
        v3 = append(v3, 2147483648)
    }
    var v14 int16 = -30
    var v15 uint16 = 320
    var v16 uint16 = (^(v15))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
