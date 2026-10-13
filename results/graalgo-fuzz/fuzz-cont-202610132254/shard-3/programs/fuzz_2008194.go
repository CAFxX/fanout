package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]int)
    var v3 string = "foo bar"
    v4 := make(map[uint32]bool)
    var v5 string = (("hello" + (v3 + v3)) + v3)
    if true {
        switch v3 {
        case "a":
            v1 = v1
        case "\n\t":
            var v6 int16 = -32768
            _ = v6
        }
        switch 128 {
        case 32767:
            var v7 float32 = (((((0.1 * (1e10 / -2.25)) + (((0.0 - float32((1e10 / float64((-2.25 * (-2.25 * (1.5 - 0.0))))))) + 0.0) / 3.14159)) / float32((1.5 / 1.0))) + 0.1) - (0.1 - 3.14159))
            _ = v7
        default:
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int64 = 82
                _ = fi8
                _ = v9
            }
        }
    }
    fmt.Println(v1)
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint16 = 32767
            var v13 uint = 26
            switch "foo bar" {
            case "":
                var v14 uint8 = 220
                _ = v14
            default:
                var v15 int = fi11
                _ = v15
            }
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    if true {
        var v16 bool = true
        _ = v16
    } else {
        for fi17 := 0; fi17 < 13; fi17++ {
            v18 := GenS1{F0: 2147483647, F1: 293, F2: v1}
            _ = fi17
            _ = v18
        }
    }
    v3 = "\n\t"
    fmt.Println(v5)
    if true {
        _ = len(v4)
    }
    fmt.Println(v1, v5)
    _ = len(v4)
    if false {
        v3 = ((v5 + v3) + "\n\t")
        v3 = v3
        v5 = "a"
    }
    fmt.Println(v5, v1, v1)
    fmt.Println(v3, v1)
    var v19 string = (v5 + v5)
    v1 += 1
    switch "foo bar" {
    case "a":
        for fi20 := 0; fi20 < 1; fi20++ {
            v1 += (v1 & v1)
            _ = len(v2)
            if ((71 <= 127) || (42 < 4294967295)) {
                var v21 uint32 = (v1 & (v1 | 565))
                _ = v21
            }
            _ = fi20
        }
    case "\n\t":
        var v22 uint8 = 207
        _ = v22
    default:
        fmt.Println(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
