package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    v3 := []int{-75, 42}
    var v4 uint8 = 1
    var v5 float32 = 0.0
    var v6 uint16 = 2
    v6 = (v6 | uint16(v4))
    v7 := GenS1{F0: (v6 - (v6 >> ((v2) & 63)))}
    fmt.Println(v4)
    v8 := v1["foo bar"]
    if ((v8 < (v8 + (v8 % (((v8 + v8)) | 1)))) && ("a" <= ("foo bar" + ("a" + "hello")))) {
        v9 := GenS1{F0: (^(v6))}
        fmt.Println(v6, v5)
        if false {
            if (!true) {
                if (18446744073709551615 > 42) {
                    v2 = 1
                } else {
                    v3 = append(v3, 0)
                    _ = len(v3)
                }
                v10 := GenS2{F0: 65535, F1: v6}
                _ = v10
            }
        } else {
            switch ("a" + "foo bar") {
            case "":
                var v11 int = (-(v8))
                _ = v11
            case "hello":
                fmt.Println(v8)
            }
            if len(v3) > 0 {
                v3[0] = v8
            }
            fmt.Println(v8, v4)
        }
        _ = v9
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        switch fi12 {
        case -1:
            var v13 int32 = 0
            _ = v13
        }
        _ = fi12
    }
    v6 -= (v6 << ((v2) & 63))
    var v14 bool = (329 >= -1)
    var v15 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
