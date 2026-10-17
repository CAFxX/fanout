package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := []uint32{42, 336, 0}
    var v2 uint64 = 42
    v3 := make(map[string]string)
    var v4 string = "a"
    var v5 int32 = -2147483648
    if false {
        var v6 bool = true
        _ = v6
    } else {
        if true {
            var v7 bool = false
            if (-63 < -128) {
                for fi8 := 0; fi8 < 4; fi8++ {
                    switch (^(fi8)) {
                    case -9223372036854775808:
                        v1 = append(v1, 26)
                    case 128:
                        v2 += v2
                    }
                    _ = fi8
                }
                var v9 int = -972
                _ = v9
            }
            _ = v7
        }
        switch v4 {
        case "":
            v10 := GenS1{F0: uint32(v2), F1: 0}
            _ = v10
        case "\n\t":
            v4 = ""
        case "a":
            _ = len(v3)
        }
        fmt.Println(v4)
    }
    v11 := v3[v4]
    var v12 uint64 = v2
    if (!false) {
        var v13 uint = 18446744073709551615
        _ = v13
    }
    v5 -= 2147483647
    v1 = append(v1, 392)
    var v14 int16 = 1000
    fmt.Println(v5, v5, v14)
    var v15 uint8 = 128
    v12 = (v2 ^ v2)
    var v16 float64 = (3.14159 * -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
