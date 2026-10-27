package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }

func main() {
    var v1 int = -1
    var v2 bool = false
    v3 := []string{"\n\t"}
    var v4 int = -2147483648
    v5 := make(map[string]int)
    var v6 uint = 839
    var v7 float64 = -2.25
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int16 = 64
        _ = fi8
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v6 = (v6 - (v6 + 2))
        _ = fi10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 bool = (13 <= 68)
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v2)
            fmt.Println(v4)
            v4 = 292
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    if (1 == 42) {
        if true {
            switch 2147483647 {
            case 32767:
                v3 = append(v3, ("hello" + ""))
            case 4294967295:
                fmt.Println(v4, v1, v2)
            case -2147483648:
                fmt.Println(v2, v6)
            default:
                switch 959 {
                case 1:
                    var v14 uint32 = 0
                    _ = v14
                case -2147483648:
                    v15 := GenS1{F0: 127, F1: (v6 % ((v6) | 1)), F2: 9223372036854775807}
                    _ = v15
                default:
                    v16 := GenS1{F0: int(v6), F1: (v6 * (v6 * v6)), F2: -9223372036854775808}
                    _ = v16
                }
            }
            switch "foo bar" {
            case "hello":
                var v17 string = ""
                _ = v17
            default:
                fmt.Println(v4, v6)
            }
            v18 := GenS1{F0: (^(v1)), F1: (^(v6)), F2: 1}
            _ = v18
        }
    } else {
        fmt.Println(v1, v2)
    }
    v6 = (v6 * v6)
    v4 = v4
    var v19 uint8 = 2
    var v20 bool = false
    var v21 int16 = -1
    if false {
        v1 = (v4 | (-(v1)))
        var v22 uint32 = 4294967295
        _ = v22
    } else {
        var v23 uint32 = 116
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
