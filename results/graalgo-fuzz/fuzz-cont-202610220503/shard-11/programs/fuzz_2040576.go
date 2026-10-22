package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint, p1 int64) int8 {
    if true {
        v6 := []int{2147483647, 174}
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                var v8 string = "foo bar"
                _ = v8
            }
            var v9 int32 = 2147483647
            if len(v6) > 0 {
                v6[0] = (fi7 % ((32767) | 1))
            }
            _ = fi7
            _ = v9
        }
        _ = v6
    } else {
        v10 := []int{268, -751, 42}
        _ = v10
    }
    v11 := []int{0, 9223372036854775807, -1, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v11
    return -1
}
func genFn1(p0 uint) int64 {
    p0 += p0
    switch ("" + ("a" + "a")) {
    case "foo bar":
        if (403 == 255) {
            p0 += (^(p0))
            var v12 uint64 = 18446744073709551615
            v13 := []int{620}
            _ = v12
            _ = v13
        }
    case "":
        var v14 uint16 = 1
        _ = v14
    case "hello":
        var v15 float64 = (((float64(0.0) - 3.14159) + 0.0) + -2.25)
        _ = v15
    }
    v16 := []int{127, 32767}
    var v17 float32 = (1e10 / 1e10)
    _ = p0
    _ = v16
    _ = v17
    return -430
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 12
    var v3 uint64 = 105
    var v4 int64 = 42
    v5 := make(map[uint32]string)
    v3 = v3
    var v18 bool = false
    v19 := GenS1{F0: 452}
    var v20 uint64 = v3
    v21 := GenS1{F0: 161}
    var v22 int32 = 42
    var v23 int8 = -128
    var v24 int = 42
    var v25 int = -9223372036854775808
    v2 -= (v2 + (v2 - 65535))
    switch 1 {
    case -764:
        if v18 {
            fmt.Println(v24, v25)
        } else {
            if (v23 != (v23 / (((v23 & (v23 + (v23 + (v23 | (v23 ^ (v23 & (-(v23))))))))) | 1))) {
                var v26 uint64 = v20
                var v27 int32 = 684
                _ = v26
                _ = v27
            }
            var v28 int16 = 1
            v25 = int(v3)
            _ = v28
        }
    case 0:
        var v29 uint16 = 0
        _ = v29
    default:
        v22 = v22
    }
    v18 = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
