package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{}
    var v2 int8 = -1
    var v3 uint = 0
    v3 -= (v3 ^ 255)
    var v4 int16 = -32768
    var v5 int32 = 829
    var v6 int8 = (v2 << ((2) & 63))
    for fi7 := 0; fi7 < 16; fi7++ {
        if false {
            var v8 float64 = float64((3.14159 - 0.0))
            fmt.Println(v4)
            v4 += 32767
            _ = v8
        }
        if false {
            switch "" {
            case "\n\t":
                switch "" {
                case "foo bar":
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v1 = append(v1, fi9)
                        var v10 string = "foo bar"
                        _ = fi9
                        _ = v10
                    }
                }
            case "hello":
                _ = len(v1)
            default:
                switch (("hello" + "a") + "a") {
                case "a":
                    switch fi7 {
                    case 2147483647:
                        v2 -= v2
                    default:
                        v2 -= v2
                    }
                case "":
                    var v11 float32 = (0.0 * (3.14159 - (0.1 / 0.1)))
                    _ = v11
                case "hello":
                    v1 = append(v1, fi7)
                }
            }
            v12 := GenS1{F0: 255, F1: uint16(v3), F2: v6}
            _ = v12
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 uint = 1
                if (24 <= (v14 << (((v14 + (v14 * 20))) & 63))) {
                    var v15 uint32 = 0
                    _ = v15
                }
                var v16 int16 = v4
                v16 = v4
                fmt.Println(fi13, v6, v4)
                var v17 int16 = (-(v16))
                v14 += 0
                v6 = v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v6 += v6
                _ = fi13
                _ = v14
                _ = v16
                _ = v17
            }
            v5 -= (v5 / (((-(v5))) | 1))
        }
        _ = fi7
    }
    v6 = (v2 % (((v2 & int8(v4))) | 1))
    switch 32768 {
    case 4294967295:
        for fi18 := 0; fi18 < 17; fi18++ {
            fmt.Println(v4, v3, v2)
            _ = fi18
        }
    case 128:
        v1 = append(v1, 0)
    default:
        var v19 float32 = float32(1.5)
        _ = v19
    }
    var v20 int64 = -1
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 = append(v1, fi21)
        _ = fi21
    }
    var v22 uint8 = 42
    v23 := GenS2{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
