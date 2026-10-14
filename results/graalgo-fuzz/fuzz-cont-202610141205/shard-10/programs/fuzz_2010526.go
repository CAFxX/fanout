package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string }

func main() {
    var v1 uint8 = 42
    var v2 uint = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 int = 32767
    var v5 int16 = 1
    var v6 string = ("" + "a")
    var v7 uint16 = 0
    delete(v3, 2147483647)
    v8 := []int{128, v4}
    var v9 uint = (v2 | 42)
    var v10 uint32 = 536
    if ((v5 + v5) == int16(v1)) {
        if true {
            v9 = (v9 - (v2 + v2))
            fmt.Println(v2, v6)
        }
    }
    _ = len(v8)
    var v11 int16 = v5
    v12 := GenS1{F0: v6, F1: v10}
    switch v4 {
    case 921:
        switch (v6 + v6) {
        case "":
            v13 := GenS2{F0: 32768, F1: (^(v7))}
            _ = v13
        case "\n\t":
            v7 += v7
        case "hello":
            v6 = ((v6 + v6) + (v6 + "hello"))
        }
    case 2147483647:
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := v3[(v10 >> ((v2) & 63))]
                fmt.Println(v4)
                _ = fi15
                _ = v16
            }
            switch "" {
            case "a":
                if true {
                    fmt.Println(v11, v1, v1)
                    v11 = (v11 / ((292) | 1))
                } else {
                    switch (v4 * 127) {
                    case 128:
                        v7 += v7
                    }
                }
            case "foo bar":
                v1 -= 0
            default:
                v17 := GenS3{F0: v6}
                _ = v17
            }
            v9 = v9
            _ = fi14
        }
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
