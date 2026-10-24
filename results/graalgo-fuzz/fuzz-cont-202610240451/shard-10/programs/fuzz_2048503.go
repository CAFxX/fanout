package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }

func main() {
    v1 := []string{"hello", "foo bar"}
    v2 := []string{("" + "a"), "\n\t"}
    var v3 string = "a"
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    switch 864 {
    case 8:
        switch "\n\t" {
        case "\n\t":
            if true {
                v3 = v3
                _ = len(v1)
                var v4 float32 = 1e10
                _ = v4
            } else {
                var v5 int = 514
                v1 = append(v1, ("\n\t" + ("hello" + "\n\t")))
                _ = v5
            }
        case "":
            var v6 uint32 = 65535
            _ = v6
        case "hello":
            fmt.Println(v3)
        }
    case 786:
        fmt.Println(v3)
    default:
        var v7 uint16 = 65535
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        v3 = ((v3 + ("foo bar" + v3)) + "")
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 42
    if true {
        v11 := GenS1{F0: 65535, F1: 1000}
        _ = v11
    } else {
        if (530 < 1) {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := GenS1{F0: 18446744073709551615, F1: 0}
                _ = fi12
                _ = v13
            }
            var v14 int32 = 42
            _ = v14
        }
    }
    v10 = v10
    if true {
        var v15 float32 = (1.5 + 0.0)
        _ = v15
    }
    switch 4294967295 {
    case -2147483648:
        var v16 uint8 = 0
        _ = v16
    case 32767:
        switch (v3 + v3) {
        case "a":
            switch v3 {
            case "hello":
                v10 = v10
            case "foo bar":
                _ = len(v2)
            case "":
                v3 = v3
            default:
                switch "hello" {
                case "":
                    v2 = append(v2, "a")
                }
            }
        default:
            fmt.Println(v10, v10)
        }
    case -455:
        var v17 uint32 = 959
        _ = v17
    default:
        v18 := GenS2{F0: "hello", F1: 0, F2: 2}
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
