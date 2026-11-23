package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int8; F2 uint8 }

func main() {
    var v1 int = 42
    var v2 uint16 = 32767
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    fmt.Println(v2, v1)
    v5 := GenS1{F0: 2}
    var v6 uint16 = (v2 * v2)
    v7 := GenS1{F0: 2}
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            fmt.Println(v1, v2, fi8)
            var v10 uint32 = 2
            switch ("hello" + "") {
            case "\n\t":
                v1 += v1
            case "hello":
                var v11 uint = 255
                _ = v11
            default:
                v10 = v10
            }
            _ = fi9
            _ = v10
        }
        v12 := GenS2{F0: (("a" + ("" + ("" + ""))) + (("" + ("\n\t" + "")) + ("a" + "\n\t"))), F1: -128, F2: 125}
        _ = fi8
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        switch "" {
        case "hello":
            fmt.Println(fi13, v1)
        case "\n\t":
            v2 = 32767
        case "":
            v6 -= v6
        }
        fmt.Println(v6, v6)
        _ = fi13
    }
    var v14 int8 = 127
    v2 += v6
    var v15 int16 = -836
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v14, v14)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
