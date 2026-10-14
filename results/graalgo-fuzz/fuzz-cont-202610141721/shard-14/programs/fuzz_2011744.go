package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int16 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 int8 }
type GenS4 struct { F0 int }
func genFn0() uint64 {
    var v4 uint32 = 4294967295
    v4 -= v4
    _ = v4
    return 0
}
func genFn1(p0 uint16) bool {
    var v5 int16 = -32768
    v6 := GenS1{F0: "a", F1: -682, F2: (v5 % (((v5 & -1)) | 1))}
    v7 := GenS2{F0: -274, F1: -44}
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (723 <= 0)
}

func main() {
    var v1 uint64 = 294
    var v2 uint = 524
    v3 := make(map[int]int)
    v8 := GenS2{F0: 393, F1: 93}
    var v9 bool = (((42 > 9223372036854775807) && true) && (42 >= 1000))
    delete(v3, 614)
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int8 = 1
        v11 += (v11 ^ v11)
        v12 := GenS3{F0: 0, F1: -367, F2: (-(v11))}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v1)
    v13 := GenS3{F0: 1, F1: -1, F2: -128}
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint8 = 243
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v1 -= (v1 ^ v1)
    switch -2 {
    case 128:
        for fi17 := 0; fi17 < 19; fi17++ {
            switch "foo bar" {
            case "":
                if (v9 && (v9 && (1 <= 2147483647))) {
                    v2 = 657
                }
            case "foo bar":
                _ = len(v3)
            default:
                fmt.Println(v9, v2)
            }
            _ = fi17
        }
    case -612:
        v18 := GenS1{F0: "", F1: 1000, F2: 560}
        _ = v18
    default:
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS4{F0: fi19}
            fmt.Println(v9, v2, v9)
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
