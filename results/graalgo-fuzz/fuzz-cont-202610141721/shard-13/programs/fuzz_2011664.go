package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16; F2 int8 }
type GenS3 struct { F0 string; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }

func main() {
    var v1 uint = 0
    v2 := make(map[string]bool)
    v3 := []int{-264, 32767}
    _ = len(v2)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    if false {
        v4 := GenS1{F0: 0, F1: 0}
        _ = v4
    } else {
        v3 = append(v3, -1)
        for fi5 := 0; fi5 < 8; fi5++ {
            if len(v3) > 0 {
                v3[0] = 0
            }
            var v6 int8 = 0
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS2{F0: 1, F1: 294, F2: 1}
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS3{F0: "hello", F1: (v1 | 2), F2: 32767}
    switch 128 {
    case 9223372036854775807:
        fmt.Println(v1)
    case 51:
        v3 = append(v3, 128)
    case 130:
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v1, fi10)
            _ = len(v3)
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "\n\t":
                for fi11 := 0; fi11 < 5; fi11++ {
                    var v12 int16 = 42
                    v1 = (v1 << ((v1) & 63))
                    v12 = int16(v1)
                    _ = fi11
                    _ = v12
                }
            case "hello":
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 string = ("hello" + "a")
                    _ = fi13
                    _ = v14
                }
            default:
                for fi15 := 0; fi15 < 13; fi15++ {
                    v16 := GenS3{F0: "", F1: v1, F2: 65535}
                    v17 := GenS1{F0: -1, F1: 42}
                    var v18 int32 = 1
                    fmt.Println(v18, v18)
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
            _ = fi10
        }
    default:
        var v19 float32 = 0.0
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(fi20, v1)
            _ = fi21
        }
        var v22 float32 = (((float32(1e10) / -2.25) / -2.25) - ((3.14159 - float32((0.0 * (1.5 * float64((0.0 - 1.5)))))) * 0.0))
        v1 = (v1 & v1)
        _ = fi20
        _ = v22
    }
    var v23 bool = ((2 >= 148) && (!false))
    v24 := GenS3{F0: (("hello" + "hello") + ("\n\t" + "\n\t")), F1: v1, F2: 1}
    v1 = 4294967295
    v25 := GenS1{F0: 9223372036854775807, F1: 128}
    v26 := GenS4{F0: 42, F1: 42, F2: 0}
    var v27 float64 = ((float64((((3.14159 * 1e10) / float32(-2.25)) / -2.25)) * 1.5) + (float64(((float32((((float64(1.5) * (1e10 - 1e10)) * (0.0 + 1e10)) - 0.1)) + (-2.25 - 1.5)) * 1e10)) / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v9
}
