package main

import (
    "fmt"
)


func main() {
    var v1 uint8 = 42
    var v2 string = "hello"
    v3 := []uint32{2}
    var v4 uint16 = 396
    _ = len(v3)
    var v5 uint16 = 32767
    var v6 int8 = -76
    v2 = (v2 + v2)
    var v7 uint64 = 4294967295
    fmt.Println(v5)
    for fi8 := 0; fi8 < 8; fi8++ {
        switch "" {
        case "":
            _ = len(v3)
        case "\n\t":
            _ = len(v3)
        default:
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        _ = fi8
    }
    v5 += v5
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
