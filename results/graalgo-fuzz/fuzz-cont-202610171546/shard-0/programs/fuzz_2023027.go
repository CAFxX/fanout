package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    var v1 uint = 1
    v2 := []string{"foo bar", ("\n\t" + "hello")}
    var v3 int16 = -32768
    v4 := make(map[uint32]int)
    var v5 int64 = 0
    if true {
        fmt.Println(v1, v1)
    } else {
        fmt.Println(v3)
        var v6 int64 = v5
        switch "\n\t" {
        case "":
            var v7 uint16 = 42
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 5; fi8++ {
                switch (fi8 ^ fi8) {
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 4294967295:
                    var v9 uint16 = 373
                    _ = v9
                case 18:
                    v5 += (-(v5))
                default:
                    for fi10 := 0; fi10 < 12; fi10++ {
                        var v11 uint16 = 32767
                        var v12 float32 = 1e10
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi8
            }
        }
        _ = v6
    }
    v13 := GenS1{F0: "hello", F1: 2}
    v1 += (v1 & v1)
    v3 -= int16(v5)
    var v14 uint = v1
    switch 32767 {
    case 0:
        _ = len(v4)
    case -9223372036854775808:
        _ = len(v4)
    case 657:
        var v15 int16 = (v3 * -618)
        _ = v15
    default:
        _ = len(v4)
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = v3
        fmt.Println(v14, fi16, v5)
        _ = fi16
        _ = v17
    }
    var v18 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
