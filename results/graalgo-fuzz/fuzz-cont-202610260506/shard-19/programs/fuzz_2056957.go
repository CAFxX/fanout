package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 int8 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint16) uint {
    var v6 uint32 = 1
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 -= 91
        var v8 int32 = 0
        v9 := GenS1{F0: 2, F1: "foo bar", F2: 65535}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 843, F1: "\n\t", F2: 65535}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v6
    return 255
}

func main() {
    var v1 string = ("" + "foo bar")
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 uint8 = 162
    var v5 float32 = 1e10
    if len(v2) > 0 {
        v2[0] = 42
    }
    v1 = v1
    fmt.Println(v5, v1)
    var v11 uint = 4294967295
    v4 = uint8(v11)
    var v12 int8 = -125
    fmt.Println(v4)
    var v13 int16 = -32768
    v13 = v13
    var v14 string = v1
    if (false || false) {
        if true {
            for fi15 := 0; fi15 < 16; fi15++ {
                v11 = (^(v11))
                switch 32767 {
                case 32768:
                    var v16 float64 = 0.1
                    _ = v16
                case 276:
                    for fi17 := 0; fi17 < 6; fi17++ {
                        switch (fi17 / (((^(fi15))) | 1)) {
                        case -584:
                            var v18 uint32 = 1
                            _ = v18
                        case 128:
                            v13 = (v13 + v13)
                        default:
                            fmt.Println(fi15, v5, v12)
                        }
                        _ = fi17
                    }
                case -328:
                    v19 := GenS1{F0: 1, F1: "\n\t", F2: (v11 >> ((v11) & 63))}
                    _ = v19
                }
                _ = fi15
            }
            v20 := GenS2{F0: 244}
            _ = v20
        } else {
            v12 = 42
        }
        v21 := GenS3{F0: v1, F1: 1, F2: v12}
        _ = v21
    } else {
        switch v14 {
        case "hello":
            _ = len(v3)
        case "":
            v22 := GenS2{F0: v4}
            _ = v22
        case "foo bar":
            var v23 uint = v11
            _ = v23
        }
    }
    v3[-2147483648] = 0
    v14 = (v1 + "a")
    var v24 uint64 = 2
    var v25 uint64 = 4294967295
    v26 := GenS4{F0: (v12 ^ v12)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
