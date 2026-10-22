package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 int16 = 42
    var v2 int64 = 1
    v3 := []string{("a" + "")}
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint8 = 0
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 int64 = v2
            fmt.Println(v2, v1)
            if false {
                fmt.Println(v1, fi7)
                switch 9223372036854775807 {
                case 42:
                    _ = len(v4)
                case -1:
                    v9 := GenS1{F0: 951, F1: 1}
                    _ = v9
                }
                var v10 int32 = -1
                var v11 int64 = (v8 - (v8 * v2))
                _ = v10
                _ = v11
            }
            _ = len(v4)
            _ = fi7
            _ = v8
        }
        v12 := GenS2{F0: "hello"}
        _ = fi5
        _ = v12
        _ = v6
    }
    if false {
        _ = len(v3)
    }
    v1 = v1
    v13 := GenS3{F0: 229}
    v1 -= (v1 & (-(v1)))
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "hello")
    }
    if (!false) {
        switch "foo bar" {
        case "hello":
            v14 := GenS3{F0: 2}
            _ = v14
        default:
            v2 -= v2
        }
        v1 -= (v1 | v1)
        v1 = (v1 ^ v1)
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v1 += -147
        _ = len(v3)
        for fi16 := 0; fi16 < 16; fi16++ {
            var v17 uint64 = 0
            var v18 int8 = 100
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
