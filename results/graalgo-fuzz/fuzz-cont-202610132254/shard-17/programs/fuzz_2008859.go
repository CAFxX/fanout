package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 bool; F1 string }

func main() {
    var v1 bool = (!true)
    var v2 int64 = 42
    var v3 int = 647
    var v4 uint = 283
    v5 := []int{1, (v3 * (v3 >> ((v4) & 63)))}
    var v6 int16 = -1
    for fi7 := 0; fi7 < 11; fi7++ {
        switch "" {
        case "a":
            v2 = (v2 % ((int64(v3)) | 1))
        case "foo bar":
            v1 = ((v3 % ((32768) | 1)) != 665)
        }
        v2 -= (v2 * v2)
        v8 := GenS1{F0: (^(v4)), F1: 42, F2: v6}
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
    case "foo bar":
        _ = len(v5)
    case "hello":
        var v9 uint64 = 0
        _ = v9
    }
    var v10 int8 = -1
    fmt.Println(v2)
    v2 = int64(v10)
    v11 := GenS1{F0: 78, F1: 42, F2: (v6 & 423)}
    var v12 uint32 = 587
    if v1 {
        _ = len(v5)
        if len(v5) > 0 {
            v5[0] = int(v6)
        }
    }
    v13 := GenS2{F0: (!(2147483647 < (v12 >> ((v4) & 63)))), F1: "hello"}
    var v14 uint8 = 255
    v15 := GenS1{F0: (v4 << ((v4) & 63)), F1: (^(v14)), F2: 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, (v3 / ((v3) | 1)))
    v6 += 1000
    v2 += v2
    v5 = append(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
