package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []uint32{65535, 2147483648, 42, 2147483647}
    v2 := make(map[int]string)
    var v3 uint16 = 32767
    v4 := GenS1{F0: 4294967295, F1: (9223372036854775807 >= 42), F2: -992}
    v3 = v3
    v5 := GenS2{F0: 0}
    if (18446744073709551615 != 18446744073709551615) {
        var v6 string = ((("foo bar" + "a") + "a") + "\n\t")
        var v7 uint8 = 109
        _ = v6
        _ = v7
    }
    v3 = (v3 & (v3 / ((0) | 1)))
    var v8 uint = 986
    v9 := GenS2{F0: 506}
    for fi10 := 0; fi10 < 1; fi10++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi10
    }
    fmt.Println(v3, v3)
    var v11 uint16 = (v3 | 32767)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
