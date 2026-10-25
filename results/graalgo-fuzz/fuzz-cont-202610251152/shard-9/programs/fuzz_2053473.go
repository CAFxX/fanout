package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }

func main() {
    var v1 uint16 = 0
    v2 := []int{0}
    var v3 int64 = 9223372036854775807
    v3 += (-(v3))
    var v4 uint = 65535
    _ = len(v2)
    for fi5 := 0; fi5 < 16; fi5++ {
        v3 = (v3 | (v3 | v3))
        v6 := GenS1{F0: uint16(v4), F1: (true || (!(v4 != v4)))}
        _ = fi5
        _ = v6
    }
    fmt.Println(v3)
    var v7 int16 = 154
    v2 = append(v2, 1)
    if false {
        switch ("hello" + "hello") {
        case "\n\t":
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := GenS1{F0: (^(v1)), F1: (true && true)}
                v7 = (v7 >> ((v4) & 63))
                for fi10 := 0; fi10 < 2; fi10++ {
                    v3 = v3
                    v3 = v3
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        case "a":
            switch "hello" {
            case "hello":
                if (886 == 4294967295) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case "a":
                var v11 string = "hello"
                _ = v11
            case "\n\t":
                v3 = (v3 & v3)
            default:
                _ = len(v2)
            }
        case "":
            _ = len(v2)
        }
        var v12 string = "hello"
        _ = v12
    }
    fmt.Println(v1, v3)
    for fi13 := 0; fi13 < 16; fi13++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = fi13
    }
    v14 := GenS1{F0: v1, F1: (128 <= 2)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
