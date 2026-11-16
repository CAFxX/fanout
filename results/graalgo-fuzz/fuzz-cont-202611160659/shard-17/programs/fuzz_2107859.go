package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int8; F1 int; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 string; F2 int64 }

func main() {
    var v1 float64 = 0.1
    var v2 int32 = -1
    var v3 uint8 = 42
    var v4 string = ("a" + "foo bar")
    var v5 uint16 = 767
    var v6 bool = false
    v7 := GenS1{F0: 883, F1: 2147483647, F2: 42}
    v8 := []int{9223372036854775807, -428}
    var v9 bool = false
    for fi10 := 0; fi10 < 16; fi10++ {
        v8 = append(v8, 9223372036854775807)
        for fi11 := 0; fi11 < 16; fi11++ {
            v3 -= v3
            if v6 {
                v12 := GenS2{F0: 127, F1: 42, F2: (("" + (v4 + v4)) + v4)}
                fmt.Println(fi10, v4, v9)
                _ = v12
            } else {
                fmt.Println(v9)
                v13 := GenS1{F0: 42, F1: 483, F2: 506}
                _ = v13
            }
            var v14 int8 = 42
            _ = fi11
            _ = v14
        }
        fmt.Println(v5, v5, v1)
        _ = fi10
    }
    var v15 uint64 = uint64(v5)
    v16 := GenS3{F0: 11}
    switch -2147483648 {
    case 266:
        switch ((v4 + "\n\t") + v4) {
        case "\n\t":
            if len(v8) > 0 {
                v8[0] = 0
            }
        case "hello":
            var v17 int = 4294967295
            _ = v17
        }
    }
    var v18 uint16 = (^(v5))
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 uint64 = v15
        _ = fi19
        _ = v20
    }
    v1 += -2.25
    v21 := GenS4{F0: -9223372036854775808, F1: "a", F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
