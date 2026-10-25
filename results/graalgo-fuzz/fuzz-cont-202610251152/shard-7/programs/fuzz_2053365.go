package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }

func main() {
    var v1 uint8 = 127
    var v2 int32 = 1
    var v3 int64 = 1
    v2 = v2
    switch "a" {
    case "hello":
        for fi4 := 0; fi4 < 4; fi4++ {
            v2 = int32(fi4)
            _ = fi4
        }
    case "foo bar":
        v3 = v3
    }
    var v5 bool = ((!(false && ("a" <= "\n\t"))) || (2 == 255))
    var v6 bool = v5
    v1 = v1
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{fi7, -2147483648}
        v5 = true
        _ = fi7
        _ = v8
    }
    v9 := []int{4294967295, 42, 66, 42}
    v2 = (v2 ^ v2)
    var v10 bool = v5
    if (-2.25 > (0.0 * 3.14159)) {
        v11 := GenS1{F0: -122, F1: "hello", F2: -97}
        _ = len(v9)
        _ = v11
    }
    v6 = true
    v2 = v2
    switch "\n\t" {
    case "\n\t":
        v9 = append(v9, 151)
    case "hello":
        v1 = (^(v1))
    case "a":
        var v12 int64 = v3
        _ = v12
    }
    v13 := GenS2{F0: 65535, F1: 125, F2: 2}
    if len(v9) > 0 {
        v9[0] = 0
    }
    var v14 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
