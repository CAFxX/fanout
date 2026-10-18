package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }

func main() {
    var v1 float32 = -2.25
    var v2 int32 = -2147483648
    var v3 uint = 255
    var v4 int64 = -1
    v5 := GenS1{F0: 42, F1: 987}
    v6 := GenS2{F0: 236, F1: 65535, F2: 32768}
    for fi7 := 0; fi7 < 18; fi7++ {
        v3 = (v3 << (((v3 & (v3 / ((134) | 1)))) & 63))
        _ = fi7
    }
    v2 = (v2 & (-(v2)))
    v8 := []int{0}
    if (!false) {
        for fi9 := 0; fi9 < 14; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint8 = 42
                var v12 bool = true
                v13 := GenS3{F0: ("foo bar" + ("" + "foo bar"))}
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            var v14 float64 = (1.5 / float64(v1))
            _ = fi9
            _ = v14
        }
        v15 := GenS4{F0: "foo bar", F1: 4}
        _ = v15
    }
    v2 -= (-(v2))
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "hello":
        switch 32768 {
        case 144:
            v1 += 3.14159
        }
    case "a":
        v16 := GenS5{F0: -385, F1: true, F2: -683}
        _ = v16
    case "\n\t":
        if ((v2 - v2) <= 2147483647) {
            v17 := GenS6{F0: (v3 / ((v3) | 1)), F1: 1000, F2: 576}
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            _ = v17
        }
    default:
        v1 = (1e10 / float32(0.1))
    }
    var v18 int32 = (v2 & -1)
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
