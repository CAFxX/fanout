package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 bool }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
func genFn0(p0 uint16) uint8 {
    v6 := GenS1{F0: 122, F1: 0, F2: true}
    var v7 uint64 = 830
    var v8 bool = (102 > 0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 213
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 int8 = -105
    v5 := make(map[string]uint32)
    v4 -= 54
    v1 -= (v1 * (v1 | (v1 >> ((497) & 63))))
    v9 := GenS2{F0: 42, F1: -2147483648, F2: v1}
    v1 = (^(v1))
    var v10 bool = true
    v10 = v10
    switch "hello" {
    case "":
        var v11 int32 = 151
        _ = v11
    default:
        switch 32768 {
        case 4294967295:
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 uint = 148
                switch fi12 {
                case -917:
                    var v14 int32 = 0
                    _ = v14
                }
                _ = fi12
                _ = v13
            }
        case -1:
            var v15 uint64 = 18446744073709551615
            _ = v15
        case 0:
            fmt.Println(v1)
        default:
            fmt.Println(v1, v10)
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        if ((((725 <= 915) || (-603 >= -1)) || false) && v10) {
            delete(v5, ("\n\t" + "\n\t"))
            v4 += v4
        } else {
            v17 := []int{}
            if ((!(v10 || (1.5 >= float64(0.1)))) && ("hello" > "hello")) {
                if true {
                    v4 = v4
                    v1 = v1
                }
            }
            var v18 uint = 18446744073709551615
            _ = v17
            _ = v18
        }
        var v19 uint = 42
        _ = fi16
        _ = v19
    }
    var v20 bool = ((v1 & (v1 / ((v1) | 1))) != v1)
    var v21 bool = (32767 == 32767)
    var v22 uint8 = 1
    if v10 {
        fmt.Println(v20, v22)
        switch 0 {
        case -601:
            var v23 int = 815
            _ = v23
        case -9223372036854775808:
            v1 = 2147483648
        case -1:
            var v24 int64 = -220
            _ = v24
        }
    }
    var v25 uint8 = (v22 + 255)
    var v26 int64 = 0
    delete(v5, "hello")
    switch "" {
    case "hello":
        v25 -= v25
    case "foo bar":
        var v27 int32 = 1
        _ = v27
    default:
        var v28 int16 = 1
        _ = v28
    }
    v22 += v25
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
