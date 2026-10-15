package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint32 = 65535
    var v2 uint8 = 2
    var v3 uint16 = 32767
    v4 := make(map[string]string)
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := GenS1{F0: 1}
        _ = fi5
        _ = v6
    }
    v1 = (v1 * (v1 % (((v1 >> ((502) & 63))) | 1)))
    v7 := GenS2{F0: -106}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int32 = -1
        _ = fi8
        _ = v9
    }
    var v10 uint8 = (v2 >> ((2) & 63))
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint64 = 2
        var v13 float32 = 3.14159
        _ = fi11
        _ = v12
        _ = v13
    }
    v2 -= 190
    switch ("foo bar" + ("foo bar" + ("a" + "\n\t"))) {
    case "hello":
        for fi14 := 0; fi14 < 17; fi14++ {
            fmt.Println(v2)
            fmt.Println(v2)
            v15 := []int{(fi14 | (fi14 | (fi14 >> ((18446744073709551615) & 63)))), fi14}
            _ = fi14
            _ = v15
        }
    case "a":
        switch 127 {
        case 0:
            var v16 uint64 = 42
            _ = v16
        case 9223372036854775807:
            v10 = 42
        default:
            for fi17 := 0; fi17 < 8; fi17++ {
                v2 = (v2 * (^(v2)))
                _ = len(v4)
                _ = fi17
            }
        }
    default:
        v10 = (v2 / (((v2 / ((v10) | 1))) | 1))
    }
    delete(v4, "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
