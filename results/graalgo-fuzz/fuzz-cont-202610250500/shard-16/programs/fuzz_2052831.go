package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 int8 }
func genFn0(p0 bool, p1 uint64) int64 {
    var v7 int = -9223372036854775808
    p1 += (p1 & p1)
    fmt.Println(v7)
    for fi8 := 0; fi8 < 13; fi8++ {
        v9 := GenS1{F0: 270, F1: "", F2: 32767}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return -62
}

func main() {
    var v1 int16 = -651
    var v2 float64 = 1.5
    v3 := make(map[int]int)
    var v4 int8 = -128
    var v5 int64 = 1
    var v6 int16 = 1
    v10 := GenS1{F0: 42, F1: "", F2: (v6 / (((^(v6))) | 1))}
    v11 := []int{4294967295, 127}
    if (-2147483648 <= 0) {
        v12 := GenS1{F0: 4294967295, F1: "a", F2: 0}
        v1 -= (^(v1))
        v13 := GenS1{F0: 255, F1: ("\n\t" + "\n\t"), F2: 1}
        _ = v12
        _ = v13
    }
    v4 = 127
    var v14 int32 = 42
    for fi15 := 0; fi15 < 3; fi15++ {
        v2 = (v2 / (v2 / float64(1.5)))
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        _ = fi15
    }
    var v16 bool = true
    fmt.Println(v4, v14)
    if ((("\n\t" + "foo bar") + "") <= ("foo bar" + ("\n\t" + ("\n\t" + "foo bar")))) {
        v2 = (float64(-2.25) * (float64(-2.25) * ((1.5 / v2) + v2)))
    }
    if ((v5 >> ((18446744073709551615) & 63)) != v5) {
        _ = len(v11)
        v17 := GenS2{F0: 1, F1: v5, F2: (v14 & -863)}
        _ = v17
    } else {
        var v18 int32 = v14
        var v19 float64 = 1e10
        switch 1 {
        case 4294967295:
            fmt.Println(v4, v4, v4)
        default:
            var v20 uint16 = 2
            _ = v20
        }
        _ = v18
        _ = v19
    }
    _ = len(v11)
    var v21 float32 = 3.14159
    var v22 float32 = v21
    if (false || true) {
        _ = len(v3)
        var v23 int64 = genFn0(true, 5)
        _ = v23
    }
    v21 = v21
    v24 := GenS3{F0: 2, F1: 1, F2: v4}
    fmt.Println(v21, v5)
    var v25 uint = 874
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
