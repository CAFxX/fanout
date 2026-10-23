package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 int8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    var v3 float64 = 0.1
    var v4 int16 = 648
    var v5 bool = false
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: 307, F1: -1, F2: 4294967295}
    var v8 uint = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint16 = 906
        fmt.Println(v5)
        _ = fi9
        _ = v10
    }
    var v11 bool = v5
    var v12 uint8 = 249
    var v13 string = (("hello" + "foo bar") + "hello")
    switch -9223372036854775808 {
    case -335:
        v14 := []int{32768, -1, -428, -385}
        _ = v14
    }
    v12 += v12
    v15 := GenS2{F0: 2, F1: 1, F2: 1}
    var v16 uint64 = 241
    for fi17 := 0; fi17 < 3; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v12 += v12
            _ = fi18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 14; fi19++ {
        v20 := []int{fi19, (fi19 % (((fi19 + (fi19 ^ (fi19 / ((fi19) | 1))))) | 1)), (fi19 << ((v8) & 63))}
        fmt.Println(v4, v11)
        var v21 uint64 = (v16 << (((v8 / ((v8) | 1))) & 63))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := []int{32768, -2147483648, 42, 32767}
    var v23 int16 = (v4 * (v4 + (v4 ^ 178)))
    for fi24 := 0; fi24 < 5; fi24++ {
        v11 = v5
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float32 = 3.14159
            if v11 {
                if true {
                    var v27 int8 = 17
                    _ = v27
                }
                switch v13 {
                case "":
                    _ = len(v22)
                default:
                    var v28 int8 = 1
                    _ = v28
                }
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            if false {
                v29 := v2[int(v12)]
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        v4 = v23
        _ = fi24
    }
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if (((v3 + 0.1) * v3) >= 0.0) {
        v13 = "foo bar"
        switch v13 {
        case "hello":
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        case "foo bar":
            fmt.Println(v16, v12)
        }
    }
    v8 = v8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
