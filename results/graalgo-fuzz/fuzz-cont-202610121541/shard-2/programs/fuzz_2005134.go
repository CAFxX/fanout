package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
func genFn0(p0 uint16, p1 uint16, p2 uint) bool {
    v6 := []int{-9223372036854775808, 9223372036854775807, 127}
    var v7 float64 = 3.14159
    fmt.Println(p1)
    v8 := GenS1{F0: (42 < 9223372036854775807), F1: 18446744073709551615, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ((1e10 + 1e10) != float32((3.14159 - 0.0)))
}
func genFn1(p0 int, p1 int16) string {
    v9 := GenS2{F0: -1, F1: 4294967295, F2: 0}
    if (true || false) {
        p1 = (p1 << ((255) & 63))
    }
    if false {
        v10 := GenS3{F0: 1}
        v11 := []int{(p0 >> ((0) & 63))}
        if len(v11) > 0 {
            v11[0] = (p0 << ((502) & 63))
        }
        _ = v10
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return ""
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = -2.25
    var v3 float32 = 1.5
    var v4 int8 = -69
    v5 := make(map[string]string)
    var v13 float32 = (v3 / v3)
    if (977 != 0) {
        fmt.Println(v3, v3)
    } else {
        switch "a" {
        case "\n\t":
            if (false && true) {
                for fi14 := 0; fi14 < 16; fi14++ {
                    delete(v5, "")
                    v15 := GenS4{F0: -527}
                    _ = fi14
                    _ = v15
                }
            } else {
                if (!false) {
                    fmt.Println(v2, v4, v3)
                    delete(v5, "\n\t")
                    var v16 int8 = (v4 ^ v4)
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    v2 = (v2 * v2)
                    _ = fi17
                }
            }
        case "foo bar":
            var v18 int32 = 42
            _ = v18
        default:
            fmt.Println(v2, v3)
        }
    }
    fmt.Println(v3, v13, v2)
    v2 -= float64((v13 * v13))
    var v19 uint64 = 42
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 string = "foo bar"
        fmt.Println(fi20)
        _ = fi20
        _ = v21
    }
    var v22 uint64 = (v19 % ((847) | 1))
    for fi23 := 0; fi23 < 10; fi23++ {
        var v24 string = (("hello" + "a") + "foo bar")
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 int8 = -128
            switch "foo bar" {
            case "":
                if false {
                    var v27 int64 = -1
                    fmt.Println(v13, fi25)
                    _ = v27
                }
            case "hello":
                v26 = (v26 << ((580) & 63))
            default:
                var v28 float64 = (float64(v13) + v2)
                v19 += 1
                var v29 float32 = v13
                _ = v28
                _ = v29
            }
            var v30 string = v24
            _ = fi25
            _ = v26
            _ = v30
        }
        fmt.Println(fi23, v22)
        _ = fi23
        _ = v24
    }
    v13 = ((1e10 * 3.14159) / (((1.5 / float32(v2)) * v3) * v3))
    v31 := GenS5{F0: 0}
    fmt.Println(v19)
    if (32767 > -518) {
        v32 := v1["foo bar"]
        var v33 int64 = 1
        _ = v32
        _ = v33
    }
    v13 -= 3.14159
    v13 -= 0.0
    v34 := GenS1{F0: false, F1: 2, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
