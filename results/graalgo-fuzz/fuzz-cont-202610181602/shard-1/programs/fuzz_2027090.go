package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 uint32, p1 uint64, p2 uint8) int64 {
    v7 := []int{128, 385, 42, 32768}
    p0 += 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 bool = ((true || ((2 < 101) && true)) || false)
    var v2 int = 41
    var v3 int16 = 1000
    var v4 float64 = (0.0 * 3.14159)
    v5 := make(map[int]bool)
    v6 := []int{128, v2}
    switch ((("a" + "a") + "\n\t") + "") {
    case "foo bar":
        fmt.Println(v4, v1)
    case "":
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int16 = v3
            delete(v5, (v2 ^ -1))
            _ = fi8
            _ = v9
        }
    case "\n\t":
        var v10 float64 = 3.14159
        _ = v10
    }
    var v11 uint8 = 1
    if (-1 <= -422) {
        fmt.Println(v3, v4)
    } else {
        fmt.Println(v2, v11, v2)
        v11 = 255
    }
    switch "a" {
    case "\n\t":
        var v12 uint = 1
        _ = v12
    case "":
        var v13 uint8 = v11
        _ = v13
    }
    v6 = append(v6, 303)
    v3 += v3
    var v14 int8 = 100
    if false {
        v2 = (-(v2))
        var v15 uint16 = 2
        var v16 int16 = (v3 ^ (v3 + (v3 / (((^(v3))) | 1))))
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: (v14 - (v14 & (v14 - (v14 & 42))))}
    if false {
        var v18 int64 = 781
        var v19 uint32 = 544
        if v1 {
            fmt.Println(v11, v1, v1)
            var v20 uint16 = 42
            _ = v20
        } else {
            v18 += 1
            v18 += (v18 % (((v18 << ((1) & 63))) | 1))
            var v21 float32 = (1e10 / (3.14159 - (((3.14159 - 0.0) - 1.5) * float32(v4))))
            _ = v21
        }
        _ = v18
        _ = v19
    }
    var v22 int = v2
    v23 := GenS2{F0: (v11 >> ((0) & 63)), F1: (^(v14)), F2: (!v1)}
    var v24 int32 = 279
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 string = ("" + (((("hello" + "hello") + "a") + "\n\t") + "foo bar"))
        if ((true || v1) && v1) {
            if v1 {
                var v27 int64 = -9223372036854775808
                _ = v27
            } else {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v29 := GenS1{F0: (-(v14))}
                    _ = fi28
                    _ = v29
                }
            }
        } else {
            v1 = (!true)
            var v30 uint = 255
            v31 := v5[v22]
            _ = v30
            _ = v31
        }
        if (((0.0 - v4) + (1e10 - (0.1 / (float64(((1.5 * -2.25) + 1.5)) - 0.0)))) <= v4) {
            v32 := GenS2{F0: v11, F1: (-(v14)), F2: true}
            v33 := GenS3{F0: fi25, F1: 123}
            switch v22 {
            case 127:
                for fi34 := 0; fi34 < 15; fi34++ {
                    v22 += -9223372036854775808
                    var v35 uint32 = 42
                    var v36 int16 = v3
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case 9223372036854775807:
                var v37 int64 = -1
                _ = v37
            case -2147483648:
                var v38 float32 = -2.25
                _ = v38
            }
            _ = v32
            _ = v33
        }
        _ = fi25
        _ = v26
    }
    v39 := GenS2{F0: v11, F1: (-(v14)), F2: (v1 || (int16(v11) >= v3))}
    if v1 {
        var v40 int16 = (v3 & (v3 & (^(v3))))
        v2 = 906
        _ = v40
    } else {
        var v41 uint16 = 42
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
