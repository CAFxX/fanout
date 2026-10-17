package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0(p0 bool) int16 {
    v4 := GenS1{F0: 2, F1: 89}
    var v5 uint16 = 65535
    v5 += v5
    _ = p0
    _ = v4
    _ = v5
    return 1000
}

func main() {
    var v1 int8 = -1
    v2 := make(map[string]uint32)
    var v3 uint = 1
    v1 = 71
    var v6 string = ("hello" + ("a" + "\n\t"))
    v7 := []int{}
    if true {
        v6 = v6
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        if (!(-1 != 1)) {
            v1 = (v1 - -56)
            if true {
                fmt.Println(v3, v3, v6)
                var v8 string = "foo bar"
                _ = v8
            }
        }
    }
    var v9 int32 = 0
    v1 += (v1 ^ (^(v1)))
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 0, F1: (-(v1))}
        fmt.Println(v3)
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 bool = true
            var v14 int64 = -9223372036854775808
            var v15 uint = (v3 / (((v3 % ((v3) | 1))) | 1))
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    v3 -= v3
    var v16 bool = (2 != 394)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch 885 {
    case 4294967295:
        _ = len(v2)
    case 127:
        switch (v6 + "hello") {
        case "\n\t":
            v3 += v3
        case "hello":
            var v17 uint16 = 1
            _ = v17
        case "foo bar":
            var v18 float64 = 0.1
            _ = v18
        }
    case 42:
        var v19 uint8 = 1
        _ = v19
    default:
        switch "a" {
        case "\n\t":
            v7 = append(v7, 32768)
        }
    }
    v20 := GenS1{F0: 255, F1: (v1 - v1)}
    fmt.Println(v3, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
