package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]bool)
    var v3 int64 = -1
    v4 := []string{"\n\t", "foo bar", "a"}
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(v1, fi5, v1)
        var v6 int64 = v3
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 255
    delete(v2, 573)
    if (float64(3.14159) > 1.5) {
        v2[2] = true
    }
    switch -1 {
    case -525:
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                switch "hello" {
                case "a":
                    var v10 bool = (127 > 1)
                    _ = v10
                case "":
                    var v11 float64 = (0.1 / 1e10)
                    _ = v11
                case "foo bar":
                    switch (fi8 / ((-1) | 1)) {
                    case 9223372036854775807:
                        var v12 uint64 = 42
                        _ = v12
                    }
                default:
                    v13 := GenS1{F0: true, F1: 42}
                    _ = v13
                }
                _ = fi9
            }
            var v14 int8 = 47
            _ = fi8
            _ = v14
        }
    case 4294967295:
        switch -873 {
        case 42:
            if true {
                fmt.Println(v7)
                switch 2147483647 {
                case -9223372036854775808:
                    _ = len(v4)
                default:
                    var v15 uint32 = 1
                    _ = v15
                }
            }
        default:
            v2[1] = true
        }
    }
    fmt.Println(v3)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 uint16 = 46
        if len(v4) > 0 {
            v4[0] = "hello"
        }
        switch "a" {
        case "\n\t":
            for fi18 := 0; fi18 < 11; fi18++ {
                v3 += v3
                _ = fi18
            }
        default:
            var v19 uint32 = 2147483648
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    var v20 uint8 = 255
    var v21 float64 = 0.1
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
