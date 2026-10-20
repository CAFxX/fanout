package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 string = "\n\t"
    v3 := []int{-1, 9223372036854775807, 294}
    var v4 int = 0
    var v5 uint64 = 803
    v6 := make(map[string]int)
    v4 = -9223372036854775808
    v7 := GenS1{F0: 0, F1: 469, F2: 42}
    fmt.Println(v4)
    v2 = v2
    var v8 float32 = ((1e10 + 1.5) * float32(-2.25))
    var v9 float32 = 1e10
    v10 := GenS2{F0: (v4 + 9223372036854775807)}
    v8 += v8
    if (v4 > v4) {
        v11 := GenS2{F0: -848}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 12; fi12++ {
            switch (fi12 << ((1) & 63)) {
            case 1:
                switch v4 {
                case 128:
                    switch v2 {
                    case "foo bar":
                        var v13 string = v2
                        _ = v13
                    default:
                        var v14 int = v4
                        _ = v14
                    }
                }
            case 0:
                var v15 int16 = 161
                _ = v15
            default:
                v4 = v4
            }
            if true {
                if len(v3) > 0 {
                    v3[0] = (fi12 % (((-(v4))) | 1))
                }
            }
            v8 = ((((v8 + float32(-2.25)) - float32(((((0.0 - float64((v9 + float32((1.5 * 3.14159))))) - 3.14159) / 0.1) / 1e10))) - v9) + (3.14159 - (v8 + (float32((3.14159 / 1e10)) / (float32(0.0) - (v8 / 3.14159))))))
            _ = fi12
        }
        switch v2 {
        case "a":
            for fi16 := 0; fi16 < 19; fi16++ {
                _ = len(v3)
                if true {
                    v8 = (1e10 + v9)
                    v9 -= 3.14159
                    v2 = (v2 + v2)
                } else {
                    v17 := GenS3{F0: 0}
                    v9 = 0.1
                    _ = v17
                }
                _ = fi16
            }
        }
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            if (255 >= 2) {
                fmt.Println(fi18)
                switch (v2 + ("\n\t" + ((v2 + v2) + ("foo bar" + (v2 + v2))))) {
                case "hello":
                    _ = len(v1)
                case "foo bar":
                    v5 -= (^(v5))
                case "":
                    v20 := GenS1{F0: 953, F1: 427, F2: 1}
                    _ = v20
                }
                v8 = v9
            }
            v1[676] = (fi18 & (fi19 & -1))
            v21 := v6[v2]
            _ = fi19
            _ = v21
        }
        v5 += v5
        _ = fi18
    }
    v22 := GenS3{F0: -28}
    v9 += v9
    v5 = (v5 % ((v5) | 1))
    if true {
        switch 678 {
        case 127:
            v5 -= v5
        case 279:
            var v23 string = v2
            _ = v23
        }
        var v24 float32 = (v9 - (0.1 + v8))
        v5 -= v5
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
