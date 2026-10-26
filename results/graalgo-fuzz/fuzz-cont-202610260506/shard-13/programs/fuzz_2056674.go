package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 string }

func main() {
    v1 := []uint32{}
    var v2 int64 = 9223372036854775807
    var v3 string = (("hello" + "\n\t") + "")
    v4 := make(map[int]int)
    v5 := []string{}
    var v6 bool = (!(!(((((v3 + ("hello" + v3)) + v3) + ((v3 + v3) + (v3 + ""))) > (v3 + "foo bar")) && true)))
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            for fi8 := 0; fi8 < 12; fi8++ {
                v3 = v3
                _ = fi8
            }
        }
        v9 := GenS1{F0: 753}
        fmt.Println(v3, v2, v6)
        _ = fi7
        _ = v9
    }
    v3 = ""
    v10 := GenS1{F0: 255}
    var v11 uint16 = 255
    v2 = v2
    v2 -= (v2 + v2)
    var v12 float64 = ((0.0 * (1.5 * (0.1 * (0.0 * (0.0 / 1.5))))) - float64(1.5))
    for fi13 := 0; fi13 < 5; fi13++ {
        _ = len(v1)
        v14 := GenS2{F0: -101, F1: "\n\t"}
        switch -549 {
        case 9223372036854775807:
            if (-128 == 100) {
                for fi15 := 0; fi15 < 9; fi15++ {
                    delete(v4, (fi13 ^ 128))
                    _ = fi15
                }
            } else {
                v2 -= (^(v2))
            }
        case 4294967295:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -172:
            fmt.Println(v11, v12, fi13)
        }
        _ = fi13
        _ = v14
    }
    fmt.Println(v11)
    _ = len(v4)
    v2 -= (v2 - v2)
    for fi16 := 0; fi16 < 14; fi16++ {
        switch fi16 {
        case 32768:
            v12 = (v12 - v12)
        default:
            switch v3 {
            case "\n\t":
                var v17 float64 = 1.5
                _ = v17
            }
        }
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
