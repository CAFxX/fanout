package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = -564
    v2 := []uint32{2147483648, 0, 4294967295, 2}
    v3 := make(map[int]string)
    var v4 uint = 2
    v1 = (v1 + (^(v1)))
    var v5 float32 = 1.5
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int64 = (v1 | (v1 % ((v1) | 1)))
        _ = len(v2)
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint32 = 1
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    v11 := GenS1{F0: 0}
    if true {
        v1 -= v1
        var v12 bool = ((((3.14159 - (float64(3.14159) + 1.5)) * 1e10) / (0.0 - 1.5)) > 0.1)
        v4 = (v4 >> ((714) & 63))
        _ = v12
    } else {
        v1 -= v1
    }
    var v13 uint8 = 255
    fmt.Println(v5, v13)
    for fi14 := 0; fi14 < 7; fi14++ {
        if false {
            switch "a" {
            case "":
                v4 = (v4 | (v4 ^ (v4 * v4)))
            case "hello":
                fmt.Println(v4, v13)
            case "\n\t":
                switch -852 {
                case 2147483647:
                    switch "foo bar" {
                    case "a":
                        fmt.Println(v5)
                    }
                case 0:
                    delete(v3, (fi14 | (fi14 & fi14)))
                case 32767:
                    v15 := GenS2{F0: 255}
                    _ = v15
                }
            default:
                var v16 uint8 = 153
                _ = v16
            }
            var v17 float32 = v5
            _ = v17
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        delete(v3, 0)
        v4 -= (v4 << ((65535) & 63))
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
