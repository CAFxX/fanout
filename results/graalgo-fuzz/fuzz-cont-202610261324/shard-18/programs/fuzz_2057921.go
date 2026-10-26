package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 int; F2 uint16 }
func genFn0() int16 {
    var v7 uint64 = 18446744073709551615
    var v8 int8 = 82
    v7 = (v7 % ((42) | 1))
    var v9 int16 = -86
    _ = v7
    _ = v8
    _ = v9
    return (^(v9))
}
func genFn1() int8 {
    var v10 uint16 = 2
    v10 = 1
    v11 := GenS1{F0: 2147483648, F1: ("\n\t" + "foo bar")}
    v12 := GenS2{F0: 4294967295}
    if (!(float32(1e10) >= 0.0)) {
        v10 = 626
        var v13 int16 = genFn0()
        var v14 uint32 = 184
        _ = v13
        _ = v14
    }
    _ = v10
    _ = v11
    _ = v12
    return 100
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{32768}
    var v3 float32 = 0.0
    v4 := []uint32{669, 291, 992, 4294967295}
    var v5 uint64 = 1
    v6 := []uint32{}
    v15 := GenS3{F0: 42, F1: 127}
    v3 = -2.25
    if (false || false) {
        v3 = v3
        var v16 int16 = genFn0()
        if (32767 < -1) {
            var v17 uint16 = 83
            _ = v17
        }
        _ = v16
    } else {
        var v18 int64 = -9223372036854775808
        if false {
            if (2 > 439) {
                var v19 float32 = ((0.0 / -2.25) / v3)
                v5 = 839
                if (2147483647 == uint32(v5)) {
                    fmt.Println(v18)
                    var v20 float64 = 0.1
                    v18 = (v18 - v18)
                    _ = v20
                }
                _ = v19
            } else {
                v21 := GenS2{F0: v5}
                _ = v21
            }
            v5 += v5
        }
        _ = v18
    }
    var v22 uint32 = 0
    v23 := GenS4{F0: true}
    if (1 != 253) {
        for fi24 := 0; fi24 < 7; fi24++ {
            for fi25 := 0; fi25 < 16; fi25++ {
                v26 := GenS1{F0: v22, F1: ("foo bar" + "hello")}
                v1 = v1
                v6 = append(v6, (^(v22)))
                var v27 int64 = 1
                _ = fi25
                _ = v26
                _ = v27
            }
            switch "\n\t" {
            case "a":
                var v28 int64 = 42
                _ = v28
            case "\n\t":
                var v29 int16 = 32767
                _ = v29
            default:
                v30 := GenS5{F0: ("foo bar" + "\n\t"), F1: (fi24 & fi24), F2: 255}
                _ = v30
            }
            _ = fi24
        }
        var v31 float32 = (((0.1 * (v3 - 3.14159)) * v3) - v3)
        _ = v31
    }
    v32 := GenS1{F0: v22, F1: (((("foo bar" + "hello") + "") + "") + "\n\t")}
    v4 = append(v4, 600)
    v3 = v3
    switch "a" {
    case "\n\t":
        var v33 int16 = genFn0()
        _ = v33
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        v1 -= (v1 - v1)
        fmt.Println(v22, v22, v1)
        _ = fi34
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        v22 = v22
        if false {
            v5 = (^(v5))
            fmt.Println(v3, v1)
        } else {
            if false {
                var v36 float32 = v3
                var v37 bool = false
                var v38 bool = v37
                _ = v36
                _ = v37
                _ = v38
            }
            if len(v6) > 0 {
                v6[0] = (v22 * (v22 | 0))
            }
            fmt.Println(fi35, v1)
        }
        _ = fi35
    }
    v3 = ((v3 + -2.25) + (float32(0.0) / v3))
    var v39 uint = 4294967295
    var v40 uint = v39
    var v41 float32 = float32(((1.5 + float64(v3)) + 1e10))
    v42 := GenS4{F0: ((!(true && (!(32767 != 4294967295)))) && false)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
