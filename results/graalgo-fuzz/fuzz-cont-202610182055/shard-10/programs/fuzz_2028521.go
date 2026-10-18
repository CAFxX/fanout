package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 bool) int8 {
    var v6 uint64 = 913
    v7 := []int{2147483647, -1, -1}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint = 65535
        _ = fi8
        _ = v9
    }
    var v10 string = "\n\t"
    var v11 uint = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v6
    _ = v7
    return 100
}
func genFn1() uint64 {
    var v12 uint64 = 18446744073709551615
    v12 -= (v12 - v12)
    var v13 uint8 = 244
    _ = v12
    _ = v13
    return v12
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = 0
    v3 := make(map[uint32]int)
    var v4 int64 = -9223372036854775808
    v5 := []string{"foo bar", "\n\t", "hello", ("foo bar" + "hello")}
    var v14 bool = (false || true)
    var v15 int8 = 1
    if v14 {
        var v16 uint8 = 0
        var v17 uint16 = 142
        _ = v16
        _ = v17
    } else {
        var v18 uint = 346
        fmt.Println(v2, v1)
        var v19 int = -211
        _ = v18
        _ = v19
    }
    v20 := GenS1{F0: 0}
    fmt.Println(v15)
    switch ("a" + "a") {
    case "foo bar":
        var v21 int16 = 1000
        _ = v21
    case "hello":
        fmt.Println(v1, v14)
    default:
        var v22 uint16 = 32767
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v14 = (2 > 128)
        _ = fi23
    }
    var v24 uint32 = 2
    switch "" {
    case "a":
        var v25 int64 = v4
        _ = v25
    case "foo bar":
        var v26 int8 = v15
        _ = v26
    default:
        fmt.Println(v2)
    }
    var v27 int = -157
    v15 = v15
    for fi28 := 0; fi28 < 5; fi28++ {
        if (v27 > (fi28 % ((fi28) | 1))) {
            if ((v14 && true) || (false && v14)) {
                v3[v24] = v27
                v5 = append(v5, ("\n\t" + "foo bar"))
                fmt.Println(fi28, v27, v14)
            }
        } else {
            var v29 uint8 = 2
            if (!(false && (420 <= 586))) {
                v30 := GenS1{F0: 32767}
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                _ = v30
            }
            _ = v29
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            v15 = (v15 ^ 117)
            _ = fi31
        }
        _ = fi28
    }
    var v32 uint32 = 383
    if v14 {
        for fi33 := 0; fi33 < 13; fi33++ {
            for fi34 := 0; fi34 < 19; fi34++ {
                fmt.Println(v24)
                v5 = append(v5, ("a" + "a"))
                _ = fi34
            }
            var v35 uint8 = 42
            _ = fi33
            _ = v35
        }
        v32 += (v32 % ((v24) | 1))
    } else {
        v15 = -32
    }
    v36 := GenS2{F0: 100, F1: v15, F2: 722}
    v14 = (float64(float32(1e10)) >= -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
}
