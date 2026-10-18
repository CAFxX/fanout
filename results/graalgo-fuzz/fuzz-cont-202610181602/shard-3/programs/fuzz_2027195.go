package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
func genFn0(p0 int64) string {
    var v5 float32 = 0.1
    v6 := []int{int(p0)}
    _ = p0
    _ = v5
    _ = v6
    return "\n\t"
}

func main() {
    var v1 int32 = 0
    var v2 float64 = (3.14159 + 3.14159)
    var v3 uint8 = 1
    var v4 uint = 1
    fmt.Println(v3, v1, v2)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: ("foo bar" + "hello"), F1: 127}
        var v9 int64 = -9223372036854775808
        switch (("foo bar" + "") + "") {
        case "":
            v10 := []int{(fi7 | 172), (fi7 << (((v4 >> (((^(v4))) & 63))) & 63)), fi7}
            _ = v10
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int8 = 42
        var v13 string = "foo bar"
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint32 = 1
    fmt.Println(v3)
    var v15 float32 = 0.1
    fmt.Println(v4, v2, v1)
    var v16 int = 32767
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 bool = true
        v19 := []int{0, 128, (fi17 * fi17), fi17}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int = (v16 & v16)
    v16 = (-(v20))
    fmt.Println(v14)
    v21 := GenS1{F0: "a", F1: -16}
    v22 := []int{v16, (v16 - v16)}
    if (!(!(true && true))) {
        switch int(v4) {
        case -9223372036854775808:
            if (-128 < 1) {
                v2 -= v2
                var v23 uint16 = 401
                v4 += v4
                _ = v23
            } else {
                var v24 float64 = (v2 / (v2 / -2.25))
                _ = v24
            }
        default:
            if (!false) {
                v15 += v15
            }
        }
        if true {
            v16 = (v16 * v20)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            v15 = (0.0 * ((v15 * -2.25) / ((float32(0.0) / float32((3.14159 + v2))) / (v15 * v15))))
        }
    }
    var v25 uint = (v4 | v4)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
