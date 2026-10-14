package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint16 }

func main() {
    v1 := []string{"a"}
    var v2 uint = 255
    v3 := make(map[string]string)
    v4 := []uint32{2147483647, 42, 2147483648, 4294967295}
    var v5 bool = (0 < (v2 | 2))
    for fi6 := 0; fi6 < 20; fi6++ {
        if false {
            fmt.Println(fi6, v5, v2)
            v2 += v2
        }
        _ = fi6
    }
    fmt.Println(v5)
    for fi7 := 0; fi7 < 12; fi7++ {
        v2 = (v2 - 42)
        if true {
            switch fi7 {
            case -2147483648:
                v8 := GenS1{F0: 398, F1: 0, F2: 2}
                _ = v8
            }
        } else {
            v5 = (((!true) && true) && v5)
            var v9 string = "hello"
            v3[v9] = v9
            _ = v9
        }
        _ = fi7
    }
    v10 := GenS1{F0: -393, F1: 1, F2: 65535}
    var v11 uint64 = 566
    v2 = v2
    switch 128 {
    case -2147483648:
        fmt.Println(v11, v11, v5)
    case 32767:
        v5 = (0 <= 1)
    default:
        v5 = (!v5)
    }
    v11 = 151
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
