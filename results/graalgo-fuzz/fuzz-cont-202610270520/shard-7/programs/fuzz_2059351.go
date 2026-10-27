package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }

func main() {
    var v1 int16 = 32767
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    switch "foo bar" {
    case "\n\t":
        var v4 uint32 = 0
        _ = v4
    }
    if false {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 int16 = 1
            v1 = 32767
            _ = fi5
            _ = v6
        }
        var v7 int8 = -99
        _ = v7
    } else {
        switch "a" {
        case "a":
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 int64 = 169
                v10 := GenS1{F0: -149, F1: v1, F2: fi8}
                _ = fi8
                _ = v10
                _ = v9
            }
        default:
            var v11 int8 = 1
            _ = v11
        }
    }
    fmt.Println(v1)
    var v12 int = 1
    v12 = v12
    if (1 >= 42) {
        v1 = -32768
        v1 = v1
    }
    v13 := v2[0]
    var v14 bool = (!(v13 || v13))
    for fi15 := 0; fi15 < 5; fi15++ {
        v1 = 0
        switch v12 {
        case -9223372036854775808:
            var v16 uint16 = 2
            _ = v16
        default:
            v14 = v13
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        switch 2147483647 {
        case 9223372036854775807:
            var v18 string = "\n\t"
            _ = v18
        default:
            switch ("\n\t" + (((("a" + "\n\t") + ((("a" + "foo bar") + ("" + "\n\t")) + "\n\t")) + ("foo bar" + "foo bar")) + ("" + "\n\t"))) {
            case "\n\t":
                var v19 int = v12
                _ = v19
            case "foo bar":
                if false {
                    v14 = (1 > -128)
                    if false {
                        var v20 uint64 = 455
                        _ = v20
                    }
                } else {
                    var v21 uint = 42
                    fmt.Println(v14, v1, fi17)
                    _ = v21
                }
            default:
                fmt.Println(v14, v13)
            }
        }
        v22 := v2[4294967295]
        _ = fi17
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
