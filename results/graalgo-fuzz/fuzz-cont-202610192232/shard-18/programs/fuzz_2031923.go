package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }

func main() {
    var v1 int64 = 1
    var v2 int16 = 132
    var v3 float32 = 3.14159
    v4 := []int{845, 0, 42}
    for fi5 := 0; fi5 < 10; fi5++ {
        if len(v4) > 0 {
            v4[0] = fi5
        }
        v2 = int16(v1)
        _ = fi5
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= 1
    var v6 int8 = 88
    var v7 uint64 = 4294967295
    var v8 bool = (255 != 65535)
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: -347, F1: (v1 / ((v1) | 1)), F2: 0}
        v6 = v6
        _ = fi9
        _ = v10
    }
    var v11 int = -2147483648
    if v8 {
        switch v11 {
        case 90:
            switch "" {
            case "foo bar":
                v2 += 67
            }
        }
    } else {
        switch "\n\t" {
        case "a":
            v6 = v6
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
