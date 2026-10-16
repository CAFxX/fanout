package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 string; F1 uint64; F2 uint32 }
func genFn0(p0 uint16) uint8 {
    v7 := []int{220, -9223372036854775808}
    fmt.Println(p0)
    switch 32767 {
    case -35:
        var v8 int = -1
        _ = v8
    case 9223372036854775807:
        v9 := GenS1{F0: (p0 ^ 65535)}
        _ = v9
    case -2147483648:
        p0 = p0
    }
    _ = p0
    _ = v7
    return 99
}

func main() {
    v1 := []int{477}
    var v2 uint8 = 128
    var v3 int64 = 0
    v4 := make(map[string]string)
    var v5 int8 = 100
    v6 := make(map[int]int)
    v10 := GenS2{F0: -997, F1: v3, F2: (v5 * (-(v5)))}
    v11 := GenS1{F0: 625}
    fmt.Println(v3)
    for fi12 := 0; fi12 < 20; fi12++ {
        _ = len(v4)
        if (2147483647 >= -1) {
            v2 -= 1
            for fi13 := 0; fi13 < 9; fi13++ {
                v14 := GenS1{F0: 1}
                var v15 uint8 = 128
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint8 = 110
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS3{F0: v5, F1: 1}
    var v18 uint64 = 4294967295
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "hello" {
    case "hello":
        v19 := GenS1{F0: 2}
        _ = v19
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 42
        }
    default:
        v20 := v6[500]
        _ = v20
    }
    v21 := GenS4{F0: (v5 - v5), F1: 2}
    var v22 int64 = 1
    if false {
        v23 := GenS1{F0: 32767}
        _ = v23
    }
    v24 := GenS5{F0: ("hello" + "\n\t"), F1: (v18 * 681), F2: 2}
    v1 = append(v1, 128)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
