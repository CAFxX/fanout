package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint8 = 136
    var v2 int32 = -1
    var v3 int8 = -86
    var v4 uint8 = 127
    v5 := []int{-434, -689, 9223372036854775807}
    v3 = v3
    v6 := GenS1{F0: v2, F1: true}
    if false {
        v3 -= 42
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 int32 = (v2 - (v2 & (v2 ^ v2)))
            switch fi7 {
            case 127:
                v2 = (^(v8))
            case 834:
                if false {
                    v2 += v2
                }
            default:
                v5 = append(v5, (fi7 % ((fi7) | 1)))
            }
            v9 := GenS1{F0: (v2 % ((v8) | 1)), F1: (!false)}
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v1 -= (v1 & v1)
        var v10 float64 = ((float64(1.5) * float64(3.14159)) * 1.5)
        _ = v10
    }
    var v11 int8 = v3
    switch "hello" {
    case "\n\t":
        switch ("foo bar" + "") {
        case "hello":
            var v12 int64 = 0
            _ = v12
        default:
            var v13 int32 = (v2 << ((735) & 63))
            _ = v13
        }
    case "foo bar":
        switch 0 {
        case 32768:
            v14 := GenS2{F0: 32767, F1: 1}
            _ = v14
        case -9223372036854775808:
            v11 -= 42
        case -497:
            fmt.Println(v11, v1)
        default:
            var v15 int16 = -688
            _ = v15
        }
    }
    var v16 string = "foo bar"
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: v2, F1: (-929 <= 1)}
        switch (v16 + (v16 + "hello")) {
        case "a":
            v16 = "hello"
        case "":
            fmt.Println(v16, v2, fi17)
        }
        _ = fi17
        _ = v18
    }
    var v19 uint8 = v1
    v5 = append(v5, -2147483648)
    var v20 float64 = 0.1
    var v21 int32 = (v2 / ((515) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
