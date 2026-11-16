package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int8 }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 1
    var v3 uint32 = 4294967295
    var v4 uint8 = 64
    v5 := []string{"a"}
    var v6 bool = (((0.0 >= 0.1) && false) || (808 != 32767))
    for fi7 := 0; fi7 < 8; fi7++ {
        v3 -= v3
        v8 := v1[fi7]
        for fi9 := 0; fi9 < 3; fi9++ {
            switch fi9 {
            case 104:
                var v10 bool = ((v4 - (^(v4))) >= v4)
                _ = v10
            case -285:
                v11 := GenS1{F0: uint16(v3), F1: -1, F2: -128}
                _ = v11
            default:
                if false {
                    var v12 int64 = int64(v2)
                    _ = v12
                } else {
                    if (true && v6) {
                        v4 += 255
                    } else {
                        v6 = (0.1 >= ((0.0 + 1.5) + (1.5 + 1.5)))
                        v6 = v6
                        v1[(fi9 * fi7)] = v8
                    }
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    var v13 string = ("" + "hello")
    _ = len(v1)
    var v14 float32 = 1.5
    v13 = "\n\t"
    var v15 uint64 = 4294967295
    var v16 int = -1
    fmt.Println(v2)
    switch v13 {
    case "foo bar":
        fmt.Println(v6, v3, v13)
    case "hello":
        v6 = true
    default:
        _ = len(v1)
    }
    v17 := GenS1{F0: 479, F1: int16(v15), F2: -17}
    if false {
        fmt.Println(v3)
        fmt.Println(v14)
    }
    fmt.Println(v15, v15)
    v18 := GenS2{F0: ("hello" + ((("foo bar" + v13) + v13) + v13))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
