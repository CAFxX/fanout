package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint8 = 42
    v2 := make(map[uint32]string)
    var v3 string = "hello"
    var v4 int32 = 2147483647
    var v5 uint16 = 192
    v6 := []uint32{}
    if (1.5 < 0.0) {
        var v7 int64 = 117
        v8 := GenS1{F0: (-1 < -1), F1: 127}
        v7 = (^(v7))
        _ = v7
        _ = v8
    } else {
        v6 = append(v6, 65535)
        v4 += -530
    }
    v4 += (v4 ^ v4)
    var v9 uint64 = 771
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 uint8 = v1
        _ = fi10
        _ = v11
    }
    if true {
        var v12 int8 = 127
        switch -9223372036854775808 {
        case 0:
            v13 := GenS1{F0: true, F1: v12}
            _ = v13
        }
        delete(v2, 269)
        _ = v12
    } else {
        switch v3 {
        case "foo bar":
            var v14 int = 128
            _ = v14
        case "\n\t":
            var v15 float64 = float64(3.14159)
            _ = v15
        }
        var v16 string = (v3 + v3)
        _ = v16
    }
    switch int(v4) {
    case 907:
        v4 = v4
    case -17:
        var v17 uint64 = v9
        _ = v17
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v2[397] = v3
        _ = fi18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 7; fi20++ {
            if (!(false || false)) {
                _ = len(v6)
                fmt.Println(v3)
                v21 := GenS1{F0: false, F1: -1}
                _ = v21
            }
            _ = fi20
        }
        _ = fi19
    }
    v1 = 183
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
