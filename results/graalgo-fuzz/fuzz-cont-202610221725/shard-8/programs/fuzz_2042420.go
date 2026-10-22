package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int16; F1 int16; F2 uint }

func main() {
    var v1 uint = 255
    var v2 uint32 = 525
    v3 := []string{""}
    var v4 uint64 = 455
    v5 := GenS1{F0: "", F1: true}
    var v6 uint8 = 2
    switch int(v2) {
    case 32767:
        var v7 uint16 = 327
        _ = v7
    case -9223372036854775808:
        fmt.Println(v4)
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            _ = len(v3)
            if true {
                switch "\n\t" {
                case "hello":
                    v1 = (^(v1))
                default:
                    v2 += v2
                }
            }
            v9 := GenS1{F0: (("" + (("hello" + "\n\t") + "hello")) + "a"), F1: (-1 != -2147483648)}
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: "a", F1: true}
    v11 := GenS1{F0: "\n\t", F1: ((1e10 >= 0.0) && true)}
    v3 = append(v3, "foo bar")
    v12 := GenS1{F0: ("hello" + "a"), F1: true}
    var v13 int32 = -409
    var v14 uint8 = (v6 - v6)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = 953
    var v15 uint64 = 548
    v16 := GenS2{F0: 1000, F1: -32768, F2: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
