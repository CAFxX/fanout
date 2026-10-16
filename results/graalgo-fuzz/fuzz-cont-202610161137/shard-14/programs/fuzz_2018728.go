package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int32; F2 uint8 }
type GenS4 struct { F0 int }
func genFn0(p0 int32, p1 uint16, p2 bool) string {
    if (42 < 17) {
        p0 = p0
        p2 = false
    } else {
        var v7 bool = p2
        _ = v7
    }
    v8 := GenS1{F0: -185, F1: 1}
    p0 = (p0 & (p0 + p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return ("hello" + ("a" + ("\n\t" + "foo bar")))
}
func genFn1() int64 {
    v9 := []int{754, -427, -9223372036854775808, -2147483648}
    if (248 <= 1) {
        var v10 string = "foo bar"
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            fmt.Println(fi11)
            var v12 int8 = 100
            _ = fi11
            _ = v12
        }
    }
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint = 636
    var v2 int16 = 32767
    v3 := []string{}
    var v4 int64 = -776
    var v5 float32 = 0.0
    var v6 string = "hello"
    v13 := GenS2{F0: -1}
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int16 = (-(v2))
        _ = fi14
        _ = v15
    }
    if len(v3) > 0 {
        v3[0] = (v6 + v6)
    }
    if (int32(v1) > 2147483647) {
        var v16 float64 = 0.1
        _ = v16
    }
    v17 := GenS2{F0: -1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v18 := GenS3{F0: 127, F1: 2147483647, F2: 2}
    v3 = append(v3, v6)
    var v19 int64 = 3
    for fi20 := 0; fi20 < 9; fi20++ {
        v21 := GenS4{F0: fi20}
        _ = fi20
        _ = v21
    }
    v6 = (((v6 + (v6 + (v6 + (((((v6 + v6) + v6) + v6) + (v6 + ("" + v6))) + (v6 + v6))))) + "") + v6)
    var v22 float64 = 0.1
    v1 = (v1 / ((v1) | 1))
    _ = len(v3)
    var v23 uint16 = 2
    v3 = append(v3, v6)
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
