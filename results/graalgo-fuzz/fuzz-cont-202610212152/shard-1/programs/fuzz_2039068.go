package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int8 }
func genFn0(p0 int32) int64 {
    fmt.Println(p0)
    var v7 uint64 = 2
    switch 9223372036854775807 {
    case 42:
        var v8 uint32 = 2147483648
        _ = v8
    case 2147483647:
        for fi9 := 0; fi9 < 14; fi9++ {
            switch fi9 {
            case 1:
                v7 -= v7
            case 127:
                var v10 int8 = -1
                _ = v10
            case 32768:
                var v11 bool = (!false)
                _ = v11
            default:
                var v12 uint16 = 515
                _ = v12
            }
            v7 = v7
            _ = fi9
        }
    default:
        v13 := GenS1{F0: -299}
        _ = v13
    }
    v14 := []int{32768, 4294967295, -390}
    var v15 int32 = p0
    _ = p0
    _ = v14
    _ = v15
    _ = v7
    return -1
}
func genFn1(p0 int64, p1 string, p2 string) int16 {
    v16 := []int{4294967295, 0, 639}
    var v17 float64 = 1.5
    v17 -= v17
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 string = "\n\t"
    v3 := []uint32{65535}
    var v4 int64 = 497
    v5 := make(map[int]int)
    var v6 float64 = (1e10 - (1e10 / 1.5))
    var v18 float32 = ((-2.25 + 3.14159) + float32(v6))
    if false {
        var v19 int16 = int16(v4)
        fmt.Println(v18, v4)
        _ = v19
    } else {
        if (v6 >= 1.5) {
            for fi20 := 0; fi20 < 7; fi20++ {
                if len(v3) > 0 {
                    v3[0] = 0
                }
                fmt.Println(v6, fi20)
                _ = fi20
            }
        }
        switch v2 {
        case "\n\t":
            fmt.Println(v2, v4, v18)
        }
    }
    v18 = v18
    v18 = 1.5
    v2 = v2
    v21 := GenS1{F0: 1000}
    v2 = "hello"
    fmt.Println(v4, v18)
    v4 = v4
    delete(v5, 9223372036854775807)
    if true {
        for fi22 := 0; fi22 < 5; fi22++ {
            var v23 int16 = -1
            v4 = (v4 % ((v4) | 1))
            switch (fi22 * (fi22 / ((int(v4)) | 1))) {
            case 42:
                v24 := GenS1{F0: (v23 >> ((65535) & 63))}
                _ = v24
            case -9223372036854775808:
                v23 += genFn1((^(v4)), ("hello" + "a"), v2)
            case 556:
                switch v2 {
                case "\n\t":
                    var v25 uint16 = 2
                    _ = v25
                case "":
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                case "a":
                    v3 = append(v3, 4294967295)
                default:
                    v26 := GenS2{F0: (42 > -2147483648), F1: 42}
                    _ = v26
                }
            default:
                var v27 int32 = 1
                _ = v27
            }
            _ = fi22
            _ = v23
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 uint16 = 255
            v30 := v1[-2147483648]
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 uint32 = 375
        _ = v31
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        switch (v2 + v2) {
        case "\n\t":
            var v33 uint16 = 0
            _ = v33
        default:
            var v34 int = 983
            _ = v34
        }
        _ = fi32
    }
    if false {
        v35 := GenS1{F0: 1}
        _ = v35
    }
    v18 = (v18 * float32(v6))
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := GenS2{F0: true, F1: -1}
        v2 = ((("\n\t" + "foo bar") + v2) + (("a" + "foo bar") + v2))
        _ = fi36
        _ = v37
    }
    if (1 < 2147483647) {
        v3 = append(v3, 2147483648)
        var v38 int8 = 1
        _ = v38
    }
    fmt.Println(v4, v18, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
