package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 bool }

func main() {
    var v1 string = ("" + "a")
    v2 := make(map[int]string)
    var v3 uint32 = 2147483647
    var v4 int16 = -32768
    v2[29] = ""
    switch (v1 + (v1 + "")) {
    case "\n\t":
        v3 = (v3 & v3)
    case "hello":
        switch 4294967295 {
        case 42:
            v5 := GenS1{F0: 42}
            _ = v5
        case -2147483648:
            v6 := GenS2{F0: 684, F1: 100}
            _ = v6
        }
    default:
        fmt.Println(v4)
    }
    v7 := GenS2{F0: 1, F1: 100}
    if (!false) {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        v10 := []int{9223372036854775807, 9223372036854775807}
        _ = v10
    }
    fmt.Println(v3, v1)
    switch v1 {
    case "\n\t":
        v4 = -871
    case "hello":
        v11 := []int{4294967295, -843, -2147483648}
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v3)
        var v13 int8 = -1
        _ = fi12
        _ = v13
    }
    switch "" {
    case "a":
        var v14 float32 = 0.0
        _ = v14
    }
    var v15 uint8 = 255
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 string = v1
        _ = fi16
        _ = v17
    }
    v18 := GenS3{F0: (v1 + v1), F1: (!(!(0 != 845)))}
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 bool = true
        v2[389] = v1
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
