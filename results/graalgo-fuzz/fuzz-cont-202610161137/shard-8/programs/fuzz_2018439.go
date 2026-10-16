package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0(p0 int, p1 int16, p2 int32) uint {
    v6 := []int{p0}
    var v7 int16 = 0
    var v8 float32 = 3.14159
    if ((false || true) && false) {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint = 255
            var v11 int = p0
            _ = fi9
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            if len(v6) > 0 {
                v6[0] = -2147483648
            }
            fmt.Println(p0)
            if (-11 < -1) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int8 = 127
                    var v15 uint16 = 31
                    var v16 int32 = -1
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                var v17 uint = 205
                var v18 float32 = (v8 - float32((0.0 / 1.0)))
                _ = v17
                _ = v18
            }
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 262
}
func genFn1(p0 uint16, p1 uint16) uint64 {
    v19 := []int{4294967295, 32767}
    switch 174 {
    case 746:
        var v20 int16 = 0
        _ = v20
    }
    _ = len(v19)
    fmt.Println(p1, p0)
    if (float64((1e10 + (-2.25 - float32(0.1)))) != (0.1 - 3.14159)) {
        _ = len(v19)
    }
    _ = p0
    _ = p1
    _ = v19
    return 18446744073709551615
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 887
    var v3 uint16 = 42
    v4 := make(map[string]string)
    v5 := []string{v1}
    v4[("foo bar" + ("" + v1))] = ""
    v2 = -968
    v3 = v3
    v21 := GenS1{F0: 97, F1: (!((^(v3)) != (v3 - 65535)))}
    var v22 int32 = 42
    switch 32768 {
    case -512:
        if (!(true && false)) {
            if false {
                fmt.Println(v2, v2)
                var v23 uint32 = 48
                _ = v23
            }
            v1 = v1
        } else {
            v24 := GenS1{F0: 18446744073709551615, F1: true}
            v2 = (v2 / ((42) | 1))
            _ = v24
        }
    case -9223372036854775808:
        var v25 int64 = 9223372036854775807
        _ = v25
    case 9223372036854775807:
        var v26 int = -9223372036854775808
        _ = v26
    default:
        var v27 uint = 255
        _ = v27
    }
    v28 := GenS2{F0: 0, F1: 790, F2: -12}
    var v29 string = v1
    var v30 int = -2147483648
    v31 := GenS3{F0: 0, F1: 4294967295}
    _ = len(v5)
    v32 := GenS1{F0: 4294967295, F1: (255 >= 255)}
    v33 := GenS4{F0: v30, F1: 0}
    var v34 int64 = 0
    var v35 uint32 = 4294967295
    if (!false) {
        v36 := v4[v29]
        var v37 string = (v1 + v29)
        for fi38 := 0; fi38 < 4; fi38++ {
            var v39 uint16 = v3
            _ = fi38
            _ = v39
        }
        _ = v36
        _ = v37
    } else {
        v2 = (v2 * (v2 ^ v2))
        v30 = (v30 * (-(v30)))
        switch v30 {
        case 741:
            v3 = 42
        case -9223372036854775808:
            v40 := GenS4{F0: 32767, F1: 449}
            _ = v40
        case 32767:
            v41 := GenS2{F0: 1, F1: 849, F2: 1}
            _ = v41
        default:
            for fi42 := 0; fi42 < 20; fi42++ {
                v29 = "\n\t"
                _ = fi42
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
