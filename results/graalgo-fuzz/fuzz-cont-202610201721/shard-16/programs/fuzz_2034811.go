package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int8; F1 int32; F2 int }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64, p1 bool) uint8 {
    var v6 uint64 = p0
    v7 := []int{-849, -9223372036854775808, 88, 32767}
    if (!(!false)) {
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    var v9 float64 = (0.0 / -2.25)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v9
    return 255
}
func genFn1(p0 int64) int32 {
    p0 += p0
    p0 += p0
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{1}
    var v3 int8 = 127
    var v4 string = "foo bar"
    v5 := []int{4294967295, 128, 240}
    switch 0 {
    case 127:
        v10 := GenS1{F0: 697, F1: true}
        _ = v10
    default:
        if false {
            v3 = v3
        } else {
            if (float64(3.14159) <= 0.1) {
                _ = len(v1)
                _ = len(v1)
            } else {
                fmt.Println(v4)
            }
            v3 += v3
        }
    }
    v4 = "hello"
    var v11 string = "foo bar"
    v12 := v1[v11]
    _ = len(v1)
    if true {
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 string = v11
            var v15 int8 = 127
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 int16 = 1000
        var v17 float32 = float32((0.0 - (((float64(float32(float64(float32(float64(float32(float64(1e10))))))) + 0.0) - 0.1) * 0.1)))
        _ = v16
        _ = v17
    }
    var v18 int8 = (-(v3))
    if true {
        v19 := GenS1{F0: 2, F1: false}
        _ = v19
    } else {
        fmt.Println(v3, v4, v3)
        if false {
            if len(v2) > 0 {
                v2[0] = 969
            }
        }
        v12 = v4
    }
    switch ("" + (v12 + (v11 + v12))) {
    case "":
        v20 := v1[v12]
        _ = v20
    case "hello":
        fmt.Println(v18)
    default:
        var v21 int8 = v3
        _ = v21
    }
    v22 := GenS2{F0: 2147483648, F1: 0}
    v4 = ""
    switch 32767 {
    case 2147483647:
        var v23 uint16 = 2
        _ = v23
    case -2147483648:
        switch ((v4 + (v12 + v4)) + "foo bar") {
        case "":
            if (1 <= 9223372036854775807) {
                fmt.Println(v11, v4)
                var v24 int16 = 32767
                _ = v24
            } else {
                if (!true) {
                    v25 := GenS1{F0: 857, F1: true}
                    switch 1 {
                    case 4294967295:
                        v3 += 1
                    case 658:
                        for fi26 := 0; fi26 < 1; fi26++ {
                            var v27 int64 = 1
                            v2 = append(v2, 820)
                            _ = fi26
                            _ = v27
                        }
                    default:
                        if (!true) {
                            fmt.Println(v12)
                        }
                    }
                    _ = v25
                }
            }
        case "a":
            var v28 int = 2147483647
            _ = v28
        default:
            var v29 uint32 = 886
            _ = v29
        }
    default:
        v30 := GenS3{F0: -128, F1: 1, F2: -2147483648}
        _ = v30
    }
    for fi31 := 0; fi31 < 3; fi31++ {
        fmt.Println(fi31, v12)
        v11 = v11
        v4 = (v12 + v12)
        _ = fi31
    }
    v12 = "foo bar"
    v11 = (((v11 + (((("" + ("" + "\n\t")) + (v12 + "\n\t")) + (v12 + "foo bar")) + (v11 + "a"))) + v12) + ("a" + "hello"))
    switch 479 {
    case -9223372036854775808:
        for fi32 := 0; fi32 < 20; fi32++ {
            v33 := GenS1{F0: 468, F1: true}
            _ = fi32
            _ = v33
        }
    case 32768:
        if len(v2) > 0 {
            v2[0] = 810
        }
    case 663:
        fmt.Println(v11, v18, v18)
    default:
        for fi34 := 0; fi34 < 17; fi34++ {
            v35 := GenS4{F0: 42}
            for fi36 := 0; fi36 < 15; fi36++ {
                switch (v11 + v4) {
                case "hello":
                    var v37 uint64 = 0
                    _ = v37
                default:
                    var v38 float32 = (((((1.5 - 0.1) * float32(1e10)) / 3.14159) - 1e10) + float32(0.1))
                    _ = v38
                }
                _ = len(v2)
                _ = fi36
            }
            _ = fi34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
