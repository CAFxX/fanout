package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint32; F2 int16 }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint; F2 uint8 }

func main() {
    var v1 uint32 = 767
    v2 := []uint32{2147483648, (v1 & (v1 * v1)), v1}
    var v3 uint = uint(v1)
    var v4 bool = (0 != -128)
    v5 := make(map[string]uint32)
    var v6 uint16 = 1
    v7 := GenS1{F0: ("foo bar" > ("\n\t" + "hello"))}
    var v8 bool = (!(uint8(v1) >= 2))
    if v8 {
        v8 = v4
        switch 32768 {
        case 180:
            v4 = ((1e10 * ((3.14159 + (0.1 - ((1.5 * -2.25) * 1.5))) * (0.1 * (-2.25 / 3.14159)))) != -2.25)
        default:
            v9 := GenS2{F0: 42, F1: (^(v1)), F2: 0}
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            if (0 <= (fi10 + fi10)) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v6 = v6
                    switch "\n\t" {
                    case "\n\t":
                        switch "a" {
                        case "":
                            var v12 float32 = 0.1
                            _ = v12
                        default:
                            fmt.Println(v1, v3, v4)
                        }
                    }
                    v5["foo bar"] = (v1 % (((v1 ^ 65535)) | 1))
                    v2 = append(v2, (v1 + 300))
                    var v13 int64 = 1
                    _ = fi11
                    _ = v13
                }
                fmt.Println(v6)
            } else {
                var v14 uint32 = v1
                v3 = v3
                _ = v14
            }
            v5["hello"] = (v1 % ((v1) | 1))
            fmt.Println(fi10, v3, v8)
            _ = fi10
        }
    }
    v4 = (42 > 81)
    if v4 {
        for fi15 := 0; fi15 < 18; fi15++ {
            var v16 uint64 = 0
            _ = fi15
            _ = v16
        }
        v17 := GenS3{F0: true, F1: v6, F2: -1}
        var v18 int8 = 100
        _ = v17
        _ = v18
    }
    v19 := GenS4{F0: (v6 + uint16(v3)), F1: (^(v3)), F2: 128}
    v20 := GenS1{F0: false}
    v3 += (v3 << ((110) & 63))
    var v21 bool = (!(!v8))
    var v22 int = -9223372036854775808
    v23 := v5["foo bar"]
    fmt.Println(v8, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v24 int8 = -116
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
