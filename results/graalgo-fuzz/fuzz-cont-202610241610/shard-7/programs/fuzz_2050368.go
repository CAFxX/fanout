package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint16 = 65535
    var v3 float32 = ((-2.25 + 0.0) / 1.0)
    v4 := make(map[string]bool)
    v3 += float32(((0.1 / 1.0) / 1e10))
    switch "hello" {
    case "\n\t":
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 string = ""
            var v7 int8 = -128
            if false {
                if false {
                    v2 = 65535
                    fmt.Println(v7)
                }
                var v8 float64 = 0.1
                v9 := GenS1{F0: (v2 ^ (v2 ^ (v2 << ((255) & 63)))), F1: -128, F2: 954}
                _ = v8
                _ = v9
            } else {
                v6 = "a"
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    case "foo bar":
        var v10 int = 128
        _ = v10
    case "a":
        var v11 int8 = 1
        _ = v11
    default:
        v3 -= v3
    }
    v12 := []int{128, 32768, 127}
    v2 -= v2
    _ = len(v12)
    var v13 uint64 = 763
    v12 = append(v12, -204)
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: v2, F1: 42, F2: 0}
        v13 -= v13
        v13 += 1
        _ = fi14
        _ = v15
    }
    var v16 float64 = 0.0
    _ = len(v12)
    v13 -= (v13 >> ((255) & 63))
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint = 559
        for fi19 := 0; fi19 < 11; fi19++ {
            if len(v12) > 0 {
                v12[0] = (fi17 + 4294967295)
            }
            fmt.Println(fi17)
            _ = fi19
        }
        _ = fi17
        _ = v18
    }
    v4["\n\t"] = (int(v13) < 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
