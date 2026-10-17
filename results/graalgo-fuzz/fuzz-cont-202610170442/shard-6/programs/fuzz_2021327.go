package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS5 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = -307
    v3 := []int{127}
    v4 := make(map[int]bool)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v2 = v2
    case "foo bar":
        var v5 int16 = 32767
        _ = v5
    case "hello":
        delete(v1, 2147483647)
    default:
        var v6 int32 = 508
        _ = v6
    }
    fmt.Println(v2)
    switch 276 {
    case -643:
        v7 := GenS1{F0: 611}
        _ = v7
    default:
        if true {
            var v8 float64 = -2.25
            _ = len(v4)
            var v9 int8 = -128
            _ = v8
            _ = v9
        }
    }
    v10 := GenS2{F0: v2}
    var v11 uint8 = 127
    for fi12 := 0; fi12 < 8; fi12++ {
        if true {
            fmt.Println(fi12)
            var v13 uint8 = 2
            _ = v13
        }
        v11 = v11
        switch "foo bar" {
        case "a":
            v14 := GenS3{F0: (false || (((((0.0 - 0.1) + (1e10 + float64((0.1 - (1.5 * 0.1))))) * float64(3.14159)) <= (1.5 * 0.0)) && (1.5 <= (3.14159 * -2.25))))}
            _ = v14
        case "\n\t":
            if false {
                var v15 uint16 = 2
                switch 32768 {
                case 4294967295:
                    v3 = append(v3, -9223372036854775808)
                case -657:
                    v16 := GenS2{F0: fi12}
                    _ = v16
                case 0:
                    v17 := GenS4{F0: -128, F1: 1, F2: v15}
                    _ = v17
                default:
                    var v18 uint64 = 4294967295
                    _ = v18
                }
                for fi19 := 0; fi19 < 19; fi19++ {
                    fmt.Println(v15, v2)
                    var v20 uint8 = 54
                    v11 += (v20 ^ v11)
                    _ = fi19
                    _ = v20
                }
                _ = v15
            }
        case "":
            fmt.Println(v11, fi12, v11)
        default:
            if len(v3) > 0 {
                v3[0] = (-(fi12))
            }
        }
        _ = fi12
    }
    v21 := GenS5{F0: "", F1: 871}
    v22 := GenS5{F0: "hello", F1: 18}
    fmt.Println(v2, v11)
    var v23 int64 = 0
    v24 := GenS2{F0: -2147483648}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
