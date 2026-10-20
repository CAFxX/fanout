package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0() uint64 {
    var v4 int8 = 65
    v4 = v4
    var v5 int32 = 0
    for fi6 := 0; fi6 < 3; fi6++ {
        switch "a" {
        case "":
            if ((2 >= 1) && (!(float32(3.14159) > ((-2.25 * (0.1 * 3.14159)) + 1.5)))) {
                v7 := GenS1{F0: 225, F1: 641}
                _ = v7
            } else {
                if (2147483647 != 2147483648) {
                    v5 = (v5 ^ v5)
                    v4 += 42
                    var v8 int16 = -1
                    _ = v8
                }
                for fi9 := 0; fi9 < 12; fi9++ {
                    v5 = -967
                    _ = fi9
                }
                v10 := GenS2{F0: (v4 * (v4 | (^(v4))))}
                _ = v10
            }
        case "\n\t":
            v4 = v4
        case "foo bar":
            fmt.Println(v4, fi6)
        default:
            for fi11 := 0; fi11 < 5; fi11++ {
                v5 = v5
                _ = fi11
            }
        }
        var v12 uint64 = 4294967295
        _ = fi6
        _ = v12
    }
    _ = v4
    _ = v5
    return 53
}

func main() {
    v1 := make(map[int]int)
    var v2 string = "\n\t"
    v3 := make(map[uint32]string)
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int = (fi13 + -9223372036854775808)
        _ = fi13
        _ = v14
    }
    v15 := GenS3{F0: 797, F1: 32767, F2: 450}
    v16 := GenS1{F0: 255, F1: 42}
    _ = len(v1)
    var v17 uint64 = 18446744073709551615
    for fi18 := 0; fi18 < 8; fi18++ {
        v17 -= 800
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 uint64 = (v17 ^ (v17 * (v17 & 677)))
            fmt.Println(fi19, v20)
            v21 := v1[fi19]
            _ = fi19
            _ = v20
            _ = v21
        }
        fmt.Println(v2, fi18)
        _ = fi18
    }
    var v22 int64 = 42
    fmt.Println(v22)
    v23 := v1[9223372036854775807]
    v2 = "foo bar"
    fmt.Println(v2, v22)
    var v24 bool = true
    fmt.Println(v2)
    v25 := []int{-1, (v23 * v23), v23, -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
