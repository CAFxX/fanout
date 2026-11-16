package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint32; F2 uint16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 float32 = 0.1
    var v2 uint32 = 65535
    var v3 uint = 65535
    v4 := make(map[int]string)
    v4[128] = ("a" + "foo bar")
    var v5 uint32 = 1
    v2 -= v2
    v3 = (v3 - 103)
    v5 -= v5
    if false {
        if true {
            v6 := GenS1{F0: 785}
            fmt.Println(v2, v1, v3)
            switch -2147483648 {
            case 0:
                var v7 int64 = 42
                _ = v7
            case -829:
                v8 := []int{690, -143, -1}
                _ = v8
            case 996:
                var v9 int = 9223372036854775807
                _ = v9
            default:
                fmt.Println(v3, v1)
            }
            _ = v6
        }
        if (true && (!false)) {
            v5 = v2
            v10 := v4[127]
            var v11 uint32 = (v2 / ((v2) | 1))
            _ = v10
            _ = v11
        }
        v12 := []int{}
        _ = v12
    } else {
        v13 := GenS1{F0: -519}
        if false {
            delete(v4, 127)
            v14 := GenS2{F0: -556, F1: (v5 / (((^(v2))) | 1)), F2: 2}
            v15 := v4[-1]
            _ = v14
            _ = v15
        } else {
            var v16 string = ("" + "")
            var v17 float32 = (v1 / v1)
            var v18 float64 = float64(float32(1.5))
            _ = v16
            _ = v17
            _ = v18
        }
        switch "foo bar" {
        case "a":
            delete(v4, 0)
        }
        _ = v13
    }
    fmt.Println(v1, v3)
    var v19 uint = (v3 * v3)
    var v20 uint64 = 2
    v21 := []int{127}
    var v22 float64 = ((-2.25 / -2.25) / (1.5 + float64(3.14159)))
    v23 := GenS3{F0: (v19 << ((835) & 63))}
    v22 = v22
    v19 = (v19 ^ (v19 >> (((v3 ^ (v19 - 20))) & 63)))
    for fi24 := 0; fi24 < 20; fi24++ {
        if true {
            v25 := GenS3{F0: (^(v19))}
            _ = v25
        }
        v26 := GenS3{F0: v19}
        _ = fi24
        _ = v26
    }
    v3 -= (v3 / (((v3 >> ((v19) & 63))) | 1))
    v5 += v2
    if true {
        _ = len(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
