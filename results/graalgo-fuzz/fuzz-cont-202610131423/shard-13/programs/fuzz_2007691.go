package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int16 }
func genFn0() uint {
    v6 := []int{4294967295, 128, 42, -1}
    var v7 float64 = 0.1
    _ = v6
    _ = v7
    return 403
}

func main() {
    var v1 int16 = 1000
    var v2 int32 = 102
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]uint32)
    var v5 bool = true
    v8 := GenS1{F0: 1, F1: "hello"}
    var v9 int8 = 127
    switch (("foo bar" + "\n\t") + "\n\t") {
    case "hello":
        fmt.Println(v9)
    }
    v9 += (v9 * 0)
    v10 := GenS2{F0: v5, F1: (v2 - (v2 << ((42) & 63)))}
    v5 = v5
    var v11 uint8 = 208
    v9 += (-(v9))
    v12 := GenS1{F0: 500, F1: ""}
    fmt.Println(v2, v9, v2)
    if v5 {
        var v13 int = 121
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v9, v1, v1)
        case "foo bar":
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS2{F0: v5, F1: 0}
                v2 = (v2 ^ int32(v1))
                for fi16 := 0; fi16 < 15; fi16++ {
                    v17 := GenS3{F0: "", F1: -2147483648}
                    _ = fi16
                    _ = v17
                }
                _ = fi14
                _ = v15
            }
        case "":
            fmt.Println(v2, v13, v5)
        }
        v18 := GenS2{F0: ((^(v11)) != v11), F1: v2}
        _ = v13
        _ = v18
    }
    v19 := GenS4{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
