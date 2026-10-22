package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 uint }
type GenS2 struct { F0 int16; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 string }

func main() {
    var v1 float32 = (1e10 / float32(-2.25))
    var v2 uint8 = 15
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 float64 = (1e10 - -2.25)
    switch 70 {
    case -131:
        v6 := GenS1{F0: 890, F1: 65535, F2: 335}
        _ = v6
    default:
        v7 := GenS2{F0: 890, F1: 614, F2: -880}
        _ = v7
    }
    v8 := GenS1{F0: 255, F1: 2147483647, F2: 4294967295}
    v9 := v3[4294967295]
    v9 = v9
    v10 := GenS2{F0: 1, F1: v9, F2: -890}
    v11 := v4[127]
    for fi12 := 0; fi12 < 3; fi12++ {
        v5 -= ((v5 + v5) / 3.14159)
        var v13 uint = 2
        fmt.Println(v5)
        _ = fi12
        _ = v13
    }
    var v14 int64 = 1
    switch 805 {
    case 1:
        for fi15 := 0; fi15 < 3; fi15++ {
            for fi16 := 0; fi16 < 19; fi16++ {
                switch "\n\t" {
                case "foo bar":
                    var v17 uint8 = v2
                    _ = v17
                default:
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        _ = len(v4)
                    case "hello":
                        var v18 int = fi15
                        _ = v18
                    }
                    var v19 int8 = 42
                    var v20 uint = 255
                    _ = v19
                    _ = v20
                }
                v9 -= 42
                v21 := GenS2{F0: 42, F1: v9, F2: (v14 | v14)}
                v1 -= (v1 + (((v1 + float32(((v5 * (v5 / -2.25)) + v5))) - (v1 + v1)) / v1))
                var v22 int = 127
                _ = fi16
                _ = v21
                _ = v22
            }
            v23 := GenS1{F0: 758, F1: v9, F2: 4294967295}
            _ = fi15
            _ = v23
        }
    case 584:
        v5 = v5
    default:
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 uint = uint(v2)
            if true {
                v25 = (^(v25))
            }
            v4[(fi24 | -945)] = (!(v11 && (false || v11)))
            _ = fi24
            _ = v25
        }
    }
    v5 -= ((0.1 / v5) / v5)
    for fi26 := 0; fi26 < 5; fi26++ {
        v27 := GenS1{F0: 2, F1: 65535, F2: 824}
        v28 := GenS3{F0: v2, F1: "a"}
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
