package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 string; F2 uint64 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint = 255
    var v2 int8 = -1
    v3 := make(map[string]int)
    var v4 uint16 = 785
    v1 -= 0
    fmt.Println(v2, v4)
    var v5 int = 40
    switch "" {
    case "a":
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := GenS1{F0: 65535, F1: 127, F2: -852}
            v8 := GenS2{F0: 42, F1: (^(v4))}
            var v9 uint64 = 2
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    case "\n\t":
        v10 := GenS2{F0: 287, F1: (v4 + (v4 >> (((^(v1))) & 63)))}
        _ = v10
    }
    v11 := GenS3{F0: 128, F1: "foo bar", F2: 245}
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 float64 = 3.14159
        fmt.Println(v2, v1)
        _ = fi12
        _ = v13
    }
    if false {
        v2 = (v2 & v2)
        var v14 uint32 = 65535
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := []int{77, (v5 >> ((v1) & 63)), (fi15 * v5), -1}
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS4{F0: 463}
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(fi17, v4, v2)
            var v20 uint16 = 10
            v3["\n\t"] = 299
            _ = fi19
            _ = v20
        }
        var v21 uint64 = 184
        _ = fi17
        _ = v18
        _ = v21
    }
    v3["a"] = v5
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
