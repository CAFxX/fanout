package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int }
type GenS2 struct { F0 uint }

func main() {
    var v1 int8 = 0
    v2 := []string{""}
    var v3 int16 = 24
    v4 := []string{("hello" + "a")}
    var v5 uint32 = 2147483648
    var v6 uint = 326
    for fi7 := 0; fi7 < 11; fi7++ {
        if false {
            switch "hello" {
            case "a":
                var v8 string = ("a" + "\n\t")
                _ = v8
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = "a"
                }
            case "hello":
                v9 := GenS1{F0: -368, F1: 4294967295, F2: 2147483647}
                _ = v9
            }
        } else {
            v5 -= v5
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            var v11 int8 = (v1 >> (((v6 + (v6 ^ v6))) & 63))
            v12 := GenS1{F0: -1, F1: 2, F2: (fi10 | -711)}
            fmt.Println(fi10)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v1 = v1
        _ = fi13
    }
    _ = len(v2)
    v14 := GenS2{F0: (v6 - 644)}
    if (!((3.14159 + (1.5 * (0.1 / 1.5))) != -2.25)) {
        v5 = (v5 | 350)
        v15 := GenS1{F0: 2147483647, F1: 42, F2: -2147483648}
        _ = v15
    }
    if len(v4) > 0 {
        v4[0] = "hello"
    }
    v1 -= (v1 - 100)
    v6 += v6
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
