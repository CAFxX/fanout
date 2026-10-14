package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483648
    v3 := make(map[int]string)
    var v4 bool = (false || false)
    v5 := GenS1{F0: 42, F1: (573 < 2)}
    v6 := v3[-337]
    if v4 {
        v2 = (^(v2))
        var v7 uint8 = 29
        _ = v7
    }
    fmt.Println(v6, v4, v6)
    for fi8 := 0; fi8 < 11; fi8++ {
        if (161 <= 1) {
            v6 = "foo bar"
            if v4 {
                v6 = "hello"
            } else {
                fmt.Println(v6, v2)
            }
            for fi9 := 0; fi9 < 4; fi9++ {
                var v10 int32 = 499
                v11 := GenS1{F0: -1, F1: (!false)}
                v12 := GenS1{F0: -50, F1: (!v4)}
                _ = len(v3)
                v10 -= 42
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        if v4 {
            var v13 int = fi8
            var v14 int32 = -2147483648
            if (v4 || ((255 == 255) || (853 == 0))) {
                var v15 uint64 = 18446744073709551615
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := v3[(-(v13))]
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            _ = v13
            _ = v14
        } else {
            var v18 int16 = -970
            v18 += v18
            _ = v18
        }
        _ = fi8
    }
    v6 = v6
    for fi19 := 0; fi19 < 14; fi19++ {
        switch (fi19 & (fi19 / ((fi19) | 1))) {
        case 32768:
            v6 = v6
        case 18:
            v3[(fi19 + (fi19 >> ((2) & 63)))] = "hello"
        case -212:
            switch "a" {
            case "foo bar":
                v6 = ((v6 + ("a" + "\n\t")) + v6)
            default:
                if v4 {
                    v4 = (42 == 1)
                }
            }
        }
        _ = fi19
    }
    _ = len(v1)
    fmt.Println(v4)
    var v20 int16 = 227
    v20 = (v20 & (^(v20)))
    var v21 uint16 = 255
    var v22 int32 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
