package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int32; F2 uint16 }
type GenS7 struct { F0 int }
func genFn0(p0 uint8, p1 int) int8 {
    fmt.Println(p0, p0)
    p1 = p1
    if true {
        if (4294967295 <= 2147483647) {
            if false {
                v4 := []int{(p1 - (p1 | 0)), p1, p1}
                _ = v4
            }
            v5 := GenS1{F0: -128, F1: p1}
            if false {
                fmt.Println(p0, p0)
                p1 = (p1 % ((p1) | 1))
            }
            _ = v5
        }
        var v6 uint32 = 2147483647
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{(fi7 | (fi7 - 2147483647)), fi7, (^(fi7)), 680}
            if (-333 == 1) {
                var v9 float32 = 3.14159
                var v10 int8 = 42
                _ = v10
                _ = v9
            } else {
                p0 -= p0
                v6 -= 42
            }
            var v11 uint32 = (v6 - v6)
            _ = fi7
            _ = v11
            _ = v8
        }
        _ = v6
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint, p1 uint16) uint8 {
    fmt.Println(p0, p0)
    var v12 int8 = -128
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 4294967295, F1: 114, F2: ("foo bar" + "")}
        v15 := GenS3{F0: 42}
        var v16 uint64 = 0
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    p1 = p1
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v12
    return 255
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    v3 := []string{"hello", "", "hello", "\n\t"}
    var v17 float32 = -2.25
    v17 -= -2.25
    switch 1 {
    case 248:
        v17 -= v17
    case 127:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 42:
        if false {
            v18 := GenS4{F0: 2147483648}
            if (-1 >= 127) {
                fmt.Println(v17, v1)
            }
            v1 += (v1 ^ 4294967295)
            _ = v18
        } else {
            v19 := GenS5{F0: "\n\t"}
            var v20 int64 = 293
            _ = v19
            _ = v20
        }
    default:
        var v21 uint16 = 704
        _ = v21
    }
    var v22 int16 = 1
    v23 := GenS6{F0: 42, F1: 0, F2: 0}
    v24 := GenS7{F0: 128}
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS6{F0: fi25, F1: 305, F2: 32767}
        _ = fi25
        _ = v26
    }
    var v27 uint32 = 373
    var v28 uint16 = 1
    var v29 float64 = 0.1
    v27 = v27
    for fi30 := 0; fi30 < 4; fi30++ {
        if (!false) {
            v3 = append(v3, "foo bar")
            v29 = 3.14159
        } else {
            var v31 string = ""
            if false {
                fmt.Println(v27, v22)
                v32 := v2[(v31 + (("" + ("\n\t" + v31)) + ""))]
                v33 := GenS1{F0: -128, F1: (fi30 / ((fi30) | 1))}
                _ = v32
                _ = v33
            }
            _ = v31
        }
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 float32 = v17
            _ = fi34
            _ = v35
        }
        _ = fi30
    }
    v3 = append(v3, "")
    for fi36 := 0; fi36 < 5; fi36++ {
        var v37 uint8 = 255
        if false {
            fmt.Println(v37, v27, fi36)
        }
        if (2 <= 202) {
            var v38 float64 = float64((((((float32(0.0) * 0.0) + float32((v29 / 3.14159))) + 3.14159) * 0.1) - float32(1.5)))
            _ = v38
        } else {
            fmt.Println(v22, v27, v27)
            v29 -= v29
        }
        _ = fi36
        _ = v37
    }
    var v39 uint16 = (v28 - 1)
    var v40 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v40
}
