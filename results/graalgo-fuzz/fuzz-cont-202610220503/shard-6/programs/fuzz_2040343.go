package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    v3 := []uint32{0, 2147483648, 919, 445}
    var v4 bool = true
    var v5 uint8 = 135
    v6 := []uint32{2147483648, 42}
    v5 += v5
    _ = len(v3)
    for fi7 := 0; fi7 < 12; fi7++ {
        v5 -= v5
        v8 := GenS1{F0: 255, F1: fi7, F2: 42}
        if v4 {
            var v9 uint = 255
            for fi10 := 0; fi10 < 13; fi10++ {
                if ((1e10 * ((3.14159 / 1.0) - 0.1)) == 1.5) {
                    var v11 int16 = 1000
                    _ = v11
                }
                _ = fi10
            }
            _ = v9
        } else {
            var v12 string = "\n\t"
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(v5)
                var v14 int8 = 0
                fmt.Println(v12)
                _ = fi13
                _ = v14
            }
            for fi15 := 0; fi15 < 11; fi15++ {
                var v16 bool = false
                _ = fi15
                _ = v16
            }
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    switch ("a" + ("" + ("foo bar" + "foo bar"))) {
    case "hello":
        var v17 uint8 = v5
        _ = v17
    case "foo bar":
        v4 = v4
    case "":
        v18 := GenS2{F0: "a", F1: 7}
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        _ = len(v1)
        v5 += v5
        _ = fi19
    }
    var v20 uint32 = 296
    v20 += uint32(v5)
    v21 := GenS3{F0: -32768, F1: "a"}
    v4 = v4
    var v22 int16 = -32768
    if (0 != (v22 * v22)) {
        fmt.Println(v22)
        var v23 int8 = 42
        _ = v23
    }
    v4 = ((v5 + (v5 - v5)) != 119)
    fmt.Println(v4, v20, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
