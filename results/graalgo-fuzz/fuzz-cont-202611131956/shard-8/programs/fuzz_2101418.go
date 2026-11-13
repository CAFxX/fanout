package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 uint16 }
func genFn0() int8 {
    switch "foo bar" {
    case "a":
        if true {
            var v6 string = "\n\t"
            var v7 uint = 105
            _ = v6
            _ = v7
        }
    case "foo bar":
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint64 = 0
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    default:
        switch 53 {
        case 9223372036854775807:
            v11 := GenS1{F0: 100, F1: (32767 < 2), F2: 2}
            _ = v11
        case -86:
            var v12 int = 127
            _ = v12
        default:
            if true {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS2{F0: -128, F1: 255}
                    _ = fi13
                    _ = v14
                }
            } else {
                var v15 int32 = -2147483648
                _ = v15
            }
        }
    }
    var v16 uint8 = 255
    fmt.Println(v16)
    _ = v16
    return 42
}

func main() {
    var v1 int16 = 1
    v2 := make(map[string]int)
    var v3 uint = 4294967295
    var v4 uint = v3
    var v5 int32 = -358
    switch 1 {
    case 959:
        v3 += v3
    default:
        switch "" {
        case "hello":
            switch 127 {
            case 32767:
                v5 += (-(v5))
            case 2147483647:
                var v17 float64 = float64(-2.25)
                _ = v17
            }
        default:
            if true {
                if (!true) {
                    var v18 uint32 = 2
                    var v19 uint32 = 0
                    _ = v18
                    _ = v19
                }
            }
        }
    }
    var v20 uint32 = 42
    var v21 uint = v4
    var v22 uint = v4
    for fi23 := 0; fi23 < 10; fi23++ {
        v24 := GenS1{F0: genFn0(), F1: true, F2: v3}
        var v25 float32 = 1.5
        var v26 uint32 = (v20 >> ((v3) & 63))
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 0, F1: true, F2: v3}
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 uint16 = 2
        _ = fi28
        _ = v29
    }
    var v30 uint8 = 127
    fmt.Println(v3)
    var v31 int16 = (v1 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
