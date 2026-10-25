package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 uint) string {
    var v6 float32 = 1e10
    v7 := GenS1{F0: -1, F1: 0, F2: -1}
    var v8 int = -2147483648
    v8 = v8
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return ("" + ("a" + "foo bar"))
}
func genFn1(p0 uint, p1 int32, p2 bool) int8 {
    if p2 {
        for fi9 := 0; fi9 < 18; fi9++ {
            v10 := GenS2{F0: 4294967295, F1: 326, F2: 134}
            var v11 float64 = 0.0
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = 1
            _ = fi12
            _ = v13
        }
    }
    v14 := GenS1{F0: 127, F1: 32767, F2: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 38
}

func main() {
    var v1 uint16 = 38
    v2 := make(map[int]string)
    var v3 uint = 0
    v4 := make(map[int]int)
    var v5 int8 = -12
    v5 = v5
    v3 = 1
    var v15 int32 = -2147483648
    v3 -= v3
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v3)
        _ = fi16
    }
    switch 1 {
    case 32768:
        var v17 uint16 = (v1 + 0)
        _ = v17
    case 2147483647:
        v15 += (-(v15))
    }
    switch "" {
    case "foo bar":
        fmt.Println(v5, v1)
    default:
        if (true && true) {
            v18 := GenS3{F0: v3, F1: -9223372036854775808, F2: 65535}
            v19 := []int{32768, -963}
            fmt.Println(v15, v15)
            _ = v18
            _ = v19
        } else {
            v15 = v15
        }
    }
    _ = len(v2)
    if (-2.25 >= 1.5) {
        var v20 int8 = v5
        v21 := []int{298, 128, 42}
        _ = v20
        _ = v21
    }
    var v22 int32 = (^(v15))
    v1 = v1
    var v23 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
