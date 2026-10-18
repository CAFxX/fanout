package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int32 = 1
    v3 := []uint32{}
    var v4 uint = 18446744073709551615
    var v5 uint16 = 0
    if false {
        var v6 uint64 = 0
        v7 := GenS1{F0: -1}
        _ = v6
        _ = v7
    }
    if true {
        v1 += v1
    }
    v2 = v2
    fmt.Println(v1, v5, v1)
    v2 += v2
    for fi8 := 0; fi8 < 13; fi8++ {
        if (!false) {
            switch ("" + ("a" + "\n\t")) {
            case "foo bar":
                if true {
                    v2 = v2
                    var v9 int64 = -9223372036854775808
                    _ = len(v3)
                    _ = v9
                }
            case "":
                var v10 uint8 = 177
                _ = v10
            case "hello":
                var v11 uint8 = 255
                _ = v11
            }
        } else {
            _ = len(v3)
        }
        var v12 int8 = -82
        var v13 uint8 = 1
        _ = fi8
        _ = v12
        _ = v13
    }
    var v14 uint8 = 255
    var v15 int8 = 100
    v5 += (v5 ^ v5)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
