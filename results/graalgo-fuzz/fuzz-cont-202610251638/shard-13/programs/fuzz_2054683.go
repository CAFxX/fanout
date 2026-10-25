package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint8 }
func genFn0() int {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{}
        _ = len(v8)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -95, F1: 0}
    var v10 int32 = -2147483648
    switch 128 {
    case 2147483647:
        if false {
            if (147 < 2) {
                fmt.Println(v10)
            } else {
                v11 := GenS2{F0: 128}
                _ = v11
            }
            var v12 int32 = (v10 * 42)
            _ = v12
        } else {
            if false {
                var v13 float64 = 3.14159
                v13 = (v13 * v13)
                v13 = ((((((0.1 + v13) / v13) - 3.14159) * float64(3.14159)) * v13) * (v13 - v13))
                _ = v13
            }
            v10 += v10
        }
    case -572:
        v14 := GenS2{F0: 179}
        _ = v14
    default:
        v15 := []int{0, 0, -9223372036854775808}
        _ = v15
    }
    _ = v10
    _ = v9
    return -637
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 18446744073709551615
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    var v5 uint16 = 65535
    var v6 int = 32768
    var v16 string = "\n\t"
    switch v16 {
    case "a":
        v17 := GenS1{F0: 114, F1: 879}
        _ = v17
    case "foo bar":
        v6 = 9223372036854775807
    }
    v18 := []int{genFn0(), v6, (v6 | 2147483647), (v6 << (((v2 | (v2 + (v2 - v2)))) & 63))}
    var v19 int16 = 780
    v19 += v19
    switch v6 {
    case -1:
        v19 = (v19 & v19)
    case 945:
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS2{F0: 0}
            v22 := GenS1{F0: 9223372036854775807, F1: 752}
            v23 := GenS1{F0: -911, F1: 42}
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
    }
    _ = len(v18)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v24 float32 = 1.5
    fmt.Println(v16)
    var v25 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
