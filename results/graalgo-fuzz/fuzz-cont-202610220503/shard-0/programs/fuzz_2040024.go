package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint8 = 127
    var v4 uint8 = v3
    v4 = (v3 << ((42) & 63))
    var v5 int = 0
    switch "" {
    case "\n\t":
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                v7 := []int{(-(v5)), v5}
                _ = v7
            }
            _ = fi6
        }
    case "a":
        var v8 uint64 = 325
        _ = v8
    case "":
        switch v5 {
        case -2147483648:
            v9 := GenS1{F0: 0}
            _ = v9
        case -9223372036854775808:
            fmt.Println(v5, v4)
        }
    }
    fmt.Println(v3, v5, v4)
    var v10 float32 = -2.25
    v3 = v4
    v11 := GenS1{F0: -58}
    if ((1.5 < ((v10 + v10) * (float32(0.0) / v10))) || (9223372036854775807 < -1)) {
        var v12 uint16 = 0
        switch "foo bar" {
        case "hello":
            switch v5 {
            case 397:
                var v13 uint16 = v12
                _ = v13
            case 2147483647:
                for fi14 := 0; fi14 < 16; fi14++ {
                    for fi15 := 0; fi15 < 6; fi15++ {
                        switch "hello" {
                        case "":
                            fmt.Println(v5)
                        default:
                            v3 = v4
                        }
                        if false {
                            fmt.Println(v4, v4, v5)
                            fmt.Println(v12, fi14)
                        }
                        _ = fi15
                    }
                    v12 = 663
                    _ = fi14
                }
            }
        default:
            v12 = 1
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            v10 -= ((v10 / 1e10) / 1.0)
            v4 = 0
            _ = fi16
        }
        _ = v12
    } else {
        fmt.Println(v5, v3, v4)
        fmt.Println(v4, v10)
    }
    switch 4294967295 {
    case -2147483648:
        v17 := GenS1{F0: 0}
        _ = v17
    default:
        fmt.Println(v4, v5, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
