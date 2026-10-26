package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 string = "hello"
    var v3 uint16 = 0
    var v4 int16 = 503
    var v5 uint8 = 127
    var v6 uint = 1
    v4 = (-(v4))
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := GenS2{F0: 0, F1: false, F2: v6}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 4; fi11++ {
            switch fi11 {
            case -309:
                v12 := v1[v2]
                _ = v12
            case -918:
                v13 := []int{fi11, fi8, (^(fi8))}
                _ = v13
            default:
                var v14 int64 = -9223372036854775808
                v15 := []int{}
                _ = v14
                _ = v15
            }
            v4 = (v4 << ((0) & 63))
            v16 := v1["hello"]
            _ = fi11
            _ = v16
        }
        _ = fi8
    }
    fmt.Println(v4, v5, v5)
    v5 = v5
    _ = len(v1)
    switch v2 {
    case "\n\t":
        v17 := []int{}
        _ = v17
    default:
        fmt.Println(v4, v6)
    }
    v18 := GenS2{F0: 438, F1: true, F2: v6}
    if (!false) {
        v1["\n\t"] = v2
        switch 127 {
        case -9223372036854775808:
            switch "\n\t" {
            case "foo bar":
                v19 := GenS1{F0: 0}
                _ = v19
            case "\n\t":
                if (0.0 < (1e10 - -2.25)) {
                    switch "" {
                    case "foo bar":
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v21 := GenS2{F0: 2147483647, F1: (-2.25 == 3.14159), F2: 255}
                            v4 += 1
                            delete(v1, "a")
                            _ = fi20
                            _ = v21
                        }
                    }
                }
            case "hello":
                v22 := GenS1{F0: 39}
                _ = v22
            default:
                switch "\n\t" {
                case "\n\t":
                    var v23 uint = v6
                    _ = v23
                case "":
                    if false {
                        var v24 int = -96
                        _ = v24
                    }
                default:
                    v2 = v2
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
