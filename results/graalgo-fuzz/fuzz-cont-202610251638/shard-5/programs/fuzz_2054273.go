package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 int8 }

func main() {
    v1 := []int{814, 4294967295, -1, -2147483648}
    v2 := make(map[int]int)
    v3 := []string{"a", "foo bar", "\n\t", ("\n\t" + "\n\t")}
    var v4 uint32 = 42
    var v5 uint64 = 42
    v6 := make(map[int]bool)
    v5 = (v5 ^ (^(v5)))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("hello" + (("foo bar" + "\n\t") + "hello")) {
    case "hello":
        v4 = (^(v4))
    case "a":
        if true {
            _ = len(v6)
        }
    case "\n\t":
        v4 = (v4 & (v4 % (((v4 >> ((2) & 63))) | 1)))
    }
    switch 32768 {
    case -1:
        v7 := GenS1{F0: 2, F1: -107, F2: -45}
        _ = v7
    default:
        fmt.Println(v4, v5)
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v4)
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        var v8 uint64 = (v5 | (^(v5)))
        _ = v8
    }
    var v9 float32 = 1e10
    switch 179 {
    case -996:
        fmt.Println(v9, v5)
    case 50:
        switch 1 {
        case 32767:
            v3 = append(v3, "a")
        case 9223372036854775807:
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 10; fi11++ {
                    v1 = append(v1, -1)
                    var v12 uint8 = uint8(v4)
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v5, v5, v4)
                v4 += (v4 * 717)
                _ = fi10
            }
        }
    default:
        if true {
            var v13 uint32 = v4
            _ = v13
        } else {
            for fi14 := 0; fi14 < 6; fi14++ {
                v4 = v4
                _ = fi14
            }
        }
    }
    v15 := GenS2{F0: 0}
    v4 += 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
