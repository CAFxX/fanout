package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint32 = 42
    v2 := []uint32{v1, 2147483648, v1}
    var v3 float64 = 3.14159
    var v4 float64 = v3
    switch 1 {
    case 9:
        if len(v2) > 0 {
            v2[0] = v1
        }
    case 350:
        if true {
            if true {
                fmt.Println(v4, v1)
            } else {
                switch -2147483648 {
                case 619:
                    var v5 uint64 = 18446744073709551615
                    _ = v5
                }
                fmt.Println(v1, v3)
            }
            v6 := GenS1{F0: v1, F1: false}
            _ = v6
        }
    case -1:
        var v7 int32 = -36
        _ = v7
    default:
        v8 := GenS1{F0: (v1 + (^(v1))), F1: false}
        _ = v8
    }
    var v9 int16 = 42
    var v10 float64 = ((((0.1 * (v3 + v4)) + (((-2.25 * (v4 + 0.0)) * v3) + v4)) / v3) - v4)
    v11 := GenS1{F0: (v1 / (((v1 % ((v1) | 1))) | 1)), F1: (true && (2147483648 <= (v1 << ((65535) & 63))))}
    v2 = append(v2, v1)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS2{F0: v9, F1: ("hello" + (("foo bar" + "hello") + "\n\t"))}
        var v14 int32 = 42
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 string = "\n\t"
    v2 = append(v2, v1)
    var v16 int64 = 0
    v17 := GenS1{F0: (v1 | 65535), F1: (false || ((v9 / ((-649) | 1)) >= -32768))}
    v4 += ((float64((-2.25 - -2.25)) * 1.5) + v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 uint32 = 65535
    v19 := GenS1{F0: v18, F1: (true && (!(0.0 == 0.1)))}
    v20 := GenS3{F0: 255}
    for fi21 := 0; fi21 < 4; fi21++ {
        v9 = (v9 - v9)
        fmt.Println(v10, v10, fi21)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
