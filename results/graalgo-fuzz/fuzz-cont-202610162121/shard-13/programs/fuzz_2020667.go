package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 string; F1 int; F2 int8 }
func genFn0(p0 int8, p1 int16) int {
    if false {
        for fi4 := 0; fi4 < 6; fi4++ {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := GenS1{F0: (-(fi5)), F1: (-(p0))}
                _ = fi5
                _ = v6
            }
            fmt.Println(p0, p1)
            _ = fi4
        }
        var v7 uint16 = 2
        var v8 int32 = 206
        _ = v7
        _ = v8
    }
    p1 = p1
    if false {
        v9 := []int{918, 781}
        _ = v9
    }
    v10 := []int{2147483647, 0}
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 int64 = 1
    v2 := []string{("foo bar" + ("foo bar" + "\n\t")), "a", "a"}
    v3 := []string{"", "", "hello", "a"}
    v1 += (v1 ^ v1)
    if false {
        var v11 int8 = 1
        _ = v11
    }
    fmt.Println(v1)
    var v12 int16 = 1
    for fi13 := 0; fi13 < 12; fi13++ {
        v12 = v12
        v14 := GenS1{F0: (fi13 << ((2) & 63)), F1: -68}
        if true {
            v2 = append(v2, "foo bar")
            fmt.Println(v12, fi13, fi13)
        }
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: 2147483647, F1: 1000}
    v1 += v1
    fmt.Println(v1, v1)
    if (654 != 255) {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 float64 = 3.14159
            var v18 uint64 = 2
            _ = fi16
            _ = v17
            _ = v18
        }
    } else {
        var v19 uint16 = 315
        _ = v19
    }
    switch "hello" {
    case "":
        v12 -= v12
    case "a":
        var v20 int = -2147483648
        _ = v20
    case "foo bar":
        fmt.Println(v1, v12)
    default:
        var v21 int64 = int64(v12)
        _ = v21
    }
    var v22 int32 = 42
    fmt.Println(v1, v1, v1)
    switch -1 {
    case 511:
        v23 := GenS3{F0: 868, F1: -1}
        _ = v23
    case 315:
        if (((660 < 0) || true) && true) {
            var v24 uint64 = 80
            v25 := GenS4{F0: (v1 / ((v1) | 1)), F1: 42}
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 1; fi26++ {
                var v27 int = (fi26 >> ((268) & 63))
                if (true && (true && (v12 <= v12))) {
                    v28 := GenS2{F0: 2, F1: v12}
                    if (2147483648 > 2147483648) {
                        v29 := GenS5{F0: (((("" + "\n\t") + "\n\t") + "foo bar") + "foo bar"), F1: v27, F2: 110}
                        var v30 int = v27
                        _ = v29
                        _ = v30
                    } else {
                        v22 -= (-(v22))
                    }
                    var v31 uint8 = 255
                    _ = v28
                    _ = v31
                } else {
                    var v32 float32 = ((((((3.14159 / (1e10 / (((1.5 / float32(1e10)) * -2.25) - 0.1))) * (-2.25 + float32(((-2.25 / 1.0) / 1e10)))) - 1.5) + -2.25) + 3.14159) * 1.5)
                    switch (-(v27)) {
                    case 9223372036854775807:
                        var v33 float32 = 0.1
                        _ = v33
                    }
                    _ = v32
                }
                var v34 uint = 65535
                _ = fi26
                _ = v27
                _ = v34
            }
            var v35 int8 = -109
            _ = v35
        }
    default:
        var v36 int32 = -1
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v22
    _ = v3
}
