package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 int8, p1 int16) uint16 {
    v6 := []int{}
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: 837, F1: 99, F2: 42}
        v9 := GenS2{F0: p1, F1: "foo bar", F2: (-(p1))}
        if (2 < 128) {
            v10 := GenS3{F0: (p1 * (p1 & (p1 & (p1 / (((-(p1))) | 1))))), F1: false}
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 string = ("a" + (("hello" + ("" + "a")) + ("a" + ("a" + "a"))))
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            p0 = p0
            for fi13 := 0; fi13 < 8; fi13++ {
                var v14 bool = false
                _ = fi13
                _ = v14
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v15 int16 = p1
    if len(v6) > 0 {
        v6[0] = -654
    }
    _ = p0
    _ = p1
    _ = v15
    _ = v6
    return 765
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = (-2.25 / 1e10)
    var v3 uint64 = 42
    var v4 int = 127
    var v5 string = "a"
    if ("foo bar" > "hello") {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 float64 = v2
            _ = fi16
            _ = v17
        }
        for fi18 := 0; fi18 < 1; fi18++ {
            if false {
                v19 := []int{(v4 * 1), -9223372036854775808, 9223372036854775807}
                var v20 float32 = (0.0 - 0.0)
                v21 := GenS4{F0: 168}
                _ = v19
                _ = v20
                _ = v21
            }
            fmt.Println(v5, fi18)
            v22 := GenS4{F0: 127}
            _ = fi18
            _ = v22
        }
    } else {
        var v23 float32 = 1.5
        _ = v23
    }
    switch (v5 + "") {
    case "hello":
        fmt.Println(v5, v2, v5)
    case "a":
        switch "foo bar" {
        case "hello":
            v2 = (((3.14159 * (float64(((0.0 + -2.25) * 0.0)) * float64((3.14159 - 1e10)))) - (v2 * 1.5)) * 0.0)
        default:
            var v24 uint32 = 2
            _ = v24
        }
    case "\n\t":
        v5 = v5
    }
    fmt.Println(v5, v4, v5)
    fmt.Println(v3, v2, v4)
    v1[2147483647] = 424
    v25 := GenS2{F0: 323, F1: v5, F2: 343}
    v26 := []int{}
    fmt.Println(v2, v4)
    v27 := GenS5{F0: true, F1: int8(v3)}
    fmt.Println(v5, v4, v5)
    for fi28 := 0; fi28 < 10; fi28++ {
        v4 = v4
        if true {
            v4 -= fi28
        }
        _ = fi28
    }
    v29 := GenS2{F0: 1000, F1: "foo bar", F2: 1000}
    var v30 string = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
