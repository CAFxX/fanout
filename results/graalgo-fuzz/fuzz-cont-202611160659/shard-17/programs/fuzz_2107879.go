package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }

func main() {
    var v1 int64 = 42
    v2 := []uint32{2147483647}
    v3 := []int{}
    if (704 < 0) {
        var v4 uint16 = 255
        _ = v4
    }
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v5 := GenS1{F0: 0, F1: 42}
    v1 = v1
    fmt.Println(v1)
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch ("a" + "a") {
        case "\n\t":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            for fi6 := 0; fi6 < 13; fi6++ {
                _ = len(v3)
                if ((65535 >= 611) && true) {
                    var v7 int8 = 1
                    v8 := GenS1{F0: 951, F1: 4294967295}
                    _ = v7
                    _ = v8
                }
                _ = fi6
            }
        }
    }
    v9 := GenS1{F0: 837, F1: 42}
    var v10 uint = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
