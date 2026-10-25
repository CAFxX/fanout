package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []string{(("\n\t" + "hello") + "a"), ("foo bar" + "\n\t"), ""}
    var v2 bool = true
    var v3 float32 = (-2.25 * 3.14159)
    v4 := []uint32{2147483647, 2147483647, 4294967295}
    var v5 int64 = -1
    var v6 uint64 = 1
    if true {
        var v7 int16 = -556
        _ = v7
    } else {
        var v8 uint16 = 1
        _ = v8
    }
    fmt.Println(v5, v2)
    for fi9 := 0; fi9 < 16; fi9++ {
        fmt.Println(v3)
        v3 = (v3 - v3)
        _ = fi9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        switch "foo bar" {
        case "a":
            var v11 int16 = 42
            _ = v11
        case "":
            v12 := GenS1{F0: 127}
            _ = v12
        case "hello":
            switch ("" + ("" + "a")) {
            case "\n\t":
                v13 := GenS1{F0: 101}
                _ = v13
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
        }
        v6 = v6
        v4 = append(v4, 42)
        _ = fi10
    }
    if (!((1 != -123) && v2)) {
        v14 := GenS1{F0: 56}
        v2 = v2
        _ = v14
    }
    v2 = true
    if len(v4) > 0 {
        v4[0] = 687
    }
    v5 += v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
