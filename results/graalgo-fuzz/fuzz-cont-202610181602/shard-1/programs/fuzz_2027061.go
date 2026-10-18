package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 string }
func genFn0(p0 uint8) uint32 {
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 uint8 = (p0 % ((128) | 1))
        _ = fi5
        _ = v6
    }
    var v7 bool = false
    v8 := []int{-2147483648, -342, 128, -489}
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "a"
    var v3 string = v2
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(v3, fi10)
        var v11 float32 = float32(float64(((0.1 * ((float32(1e10) - -2.25) + -2.25)) - 1e10)))
        _ = fi10
        _ = v11
    }
    var v12 bool = (2 > 0)
    var v13 uint64 = 4294967295
    var v14 float64 = 0.0
    if (-177 != 2147483647) {
        v15 := GenS1{F0: 2, F1: -957}
        var v16 int32 = -86
        v12 = true
        _ = v15
        _ = v16
    }
    switch "a" {
    case "":
        if (!v12) {
            for fi17 := 0; fi17 < 17; fi17++ {
                v9 = 73
                v1[uint32(v9)] = ""
                _ = fi17
            }
        } else {
            fmt.Println(v12)
        }
    case "foo bar":
        v12 = (-61 == 42)
    }
    v18 := []int{}
    v19 := GenS1{F0: 32767, F1: 955}
    delete(v1, 2)
    if v12 {
        v20 := GenS2{F0: 62, F1: 1000}
        if false {
            fmt.Println(v12, v12, v3)
            v9 = (v9 % ((v9) | 1))
            if len(v18) > 0 {
                v18[0] = 585
            }
        }
        switch ("hello" + (v3 + "\n\t")) {
        case "hello":
            if false {
                var v21 uint = 203
                switch 42 {
                case -9223372036854775808:
                    v12 = false
                case -2147483648:
                    v22 := v1[4294967295]
                    _ = v22
                default:
                    fmt.Println(v3, v21, v3)
                }
                var v23 int16 = 306
                _ = v21
                _ = v23
            }
        case "\n\t":
            v24 := GenS3{F0: (740 < 2)}
            _ = v24
        case "a":
            for fi25 := 0; fi25 < 16; fi25++ {
                v13 += v9
                fmt.Println(v9, v3)
                _ = fi25
            }
        }
        _ = v20
    }
    v26 := GenS4{F0: "", F1: 2147483647}
    delete(v1, 2147483648)
    var v27 float64 = (v14 * (((float64((float32(1e10) / 0.1)) / (v14 * ((3.14159 + 1.5) - ((3.14159 + v14) - ((float64(1e10) - 1e10) / 0.1))))) * v14) * 3.14159))
    var v28 uint64 = v13
    var v29 int8 = int8(v13)
    v30 := GenS5{F0: ((v2 + ("foo bar" + ("" + v3))) + "hello")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v9
}
