package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]uint32)
    var v3 int32 = 0
    var v4 int = 127
    v5 := []string{("" + "a"), ("" + "foo bar")}
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v3, v3)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    fmt.Println(v4)
    var v7 int8 = 127
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := v2[""]
            if false {
                var v11 uint = 42
                _ = len(v2)
                var v12 uint8 = 217
                _ = v11
                _ = v12
            }
            v4 = 4294967295
            fmt.Println(v1)
            v13 := GenS1{F0: v1, F1: (v3 & 1)}
            if (!(!false)) {
                v1 = (^(v1))
                fmt.Println(v10, fi9)
            } else {
                v14 := GenS1{F0: v1, F1: v3}
                var v15 uint16 = 2
                var v16 int = v4
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi9
            _ = v10
            _ = v13
        }
        _ = fi8
    }
    v7 = 1
    for fi17 := 0; fi17 < 18; fi17++ {
        switch "\n\t" {
        case "\n\t":
            for fi18 := 0; fi18 < 17; fi18++ {
                fmt.Println(fi18, fi17)
                _ = fi18
            }
        case "a":
            switch 588 {
            case 128:
                v19 := GenS1{F0: 2, F1: 1}
                _ = v19
            case -16:
                var v20 uint8 = 1
                _ = v20
            }
        case "hello":
            v4 += -425
        default:
            for fi21 := 0; fi21 < 16; fi21++ {
                v4 = (fi21 << ((18446744073709551615) & 63))
                _ = fi21
            }
        }
        _ = fi17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v22 uint32 = 866
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
