package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 bool; F2 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint; F2 uint }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 int16, p1 int64) uint8 {
    fmt.Println(p1, p1)
    p0 = (p0 << ((1) & 63))
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint8 = 0
    var v2 bool = (42 > 1)
    v3 := []string{"hello", "a", "a", "foo bar"}
    v4 := make(map[string]int)
    if true {
        for fi5 := 0; fi5 < 12; fi5++ {
            v2 = false
            var v6 uint = 292
            _ = fi5
            _ = v6
        }
        v2 = true
        switch "a" {
        case "foo bar":
            v2 = (!v2)
        default:
            v3 = append(v3, ("\n\t" + "a"))
        }
    }
    switch "\n\t" {
    case "a":
        for fi7 := 0; fi7 < 6; fi7++ {
            var v8 string = "\n\t"
            fmt.Println(v2)
            var v9 int8 = 64
            _ = fi7
            _ = v8
            _ = v9
        }
    default:
        v1 -= 1
    }
    if v2 {
        var v10 int8 = 100
        if (!v2) {
            delete(v4, "foo bar")
            var v11 float64 = 1e10
            v1 += (v1 / ((genFn0(1000, -987)) | 1))
            _ = v11
        } else {
            switch "\n\t" {
            case "\n\t":
                var v12 float32 = (-2.25 + 0.1)
                _ = v12
            }
        }
        _ = v10
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := GenS1{F0: 9223372036854775807}
        var v15 int8 = -1
        fmt.Println(v1, fi13, fi13)
        _ = fi13
        _ = v14
        _ = v15
    }
    if (1e10 == -2.25) {
        var v16 int8 = -128
        _ = v16
    }
    v1 -= 1
    _ = len(v3)
    if v2 {
        var v17 int64 = 1
        delete(v4, "\n\t")
        _ = v17
    }
    v18 := GenS2{F0: 2, F1: (900 > 0), F2: -655}
    v19 := GenS2{F0: 1, F1: v2, F2: 281}
    v1 += 255
    v1 = (v1 & v1)
    v20 := GenS3{F0: 604}
    if (false || v2) {
        v2 = (!true)
        if (("foo bar" + ("foo bar" + ("" + "hello"))) <= "a") {
            v1 -= (v1 * 128)
            var v21 int64 = 1
            _ = v21
        } else {
            if v2 {
                if (true || v2) {
                    v2 = (!((40 == 127) || v2))
                    v22 := GenS4{F0: "hello", F1: 2, F2: 782}
                    v2 = v2
                    _ = v22
                }
                v2 = false
                fmt.Println(v2, v2)
            }
        }
    } else {
        if false {
            var v23 int = -401
            v2 = (1e10 <= (0.1 / (((1.5 - (0.1 + float64(1.5))) - 1e10) + float64((((float32(0.1) + 3.14159) + (-2.25 + 0.1)) / 0.1)))))
            v24 := GenS3{F0: -9223372036854775808}
            _ = v23
            _ = v24
        }
    }
    v2 = (!(v2 && v2))
    if ((3.14159 * (-2.25 + 0.0)) != (0.1 + (3.14159 - (float32((1.5 - 1e10)) * 3.14159)))) {
        v25 := GenS5{F0: (v1 ^ (^(v1))), F1: v2}
        v26 := GenS2{F0: 2, F1: true, F2: -645}
        _ = v25
        _ = v26
    } else {
        v1 = v1
        v1 = (v1 / ((genFn0(1, 0)) | 1))
    }
    v2 = v2
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 int32 = 2147483647
        var v29 float32 = (1.5 + 0.0)
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
