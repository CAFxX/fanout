package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int32 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int32; F2 bool }
func genFn0() uint64 {
    var v4 uint64 = 103
    v5 := []int{127, 2147483647, -2147483648}
    v4 = v4
    v6 := GenS1{F0: (-114 < 42), F1: 127, F2: 1}
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1(p0 uint64, p1 uint16) uint16 {
    if true {
        p1 = (^(p1))
        v7 := GenS1{F0: true, F1: 9223372036854775807, F2: -1}
        v8 := GenS2{F0: (int32(p1) != -4), F1: 2147483647}
        _ = v7
        _ = v8
    } else {
        p0 += p0
        p1 = (p1 + 77)
    }
    p0 += 0
    p0 += genFn0()
    _ = p0
    _ = p1
    return (p1 >> ((1) & 63))
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, -625}
    var v3 uint = 2
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 string = "a"
        switch 2147483647 {
        case 9223372036854775807:
            v11 := v1[v10]
            _ = v11
        case 32767:
            _ = len(v1)
        }
        var v12 uint32 = 633
        _ = fi9
        _ = v10
        _ = v12
    }
    var v13 float64 = (float64(1e10) + (1.5 / 3.14159))
    var v14 uint16 = 468
    fmt.Println(v3, v13)
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint64 = 0
        _ = fi15
        _ = v16
    }
    v14 = 65535
    v13 = 3.14159
    v17 := GenS3{F0: "\n\t", F1: 42}
    var v18 uint16 = v14
    fmt.Println(v13, v13)
    v19 := GenS2{F0: (2 <= 4294967295), F1: 2}
    if (true || (128 >= -92)) {
        switch ("a" + ((("\n\t" + ("hello" + "a")) + "hello") + ("foo bar" + ("a" + "hello")))) {
        case "\n\t":
            if (false || false) {
                if false {
                    if (0.1 <= 1e10) {
                        v13 = (3.14159 * ((v13 + v13) / ((v13 + 1e10) / float64(((-2.25 / 1e10) + (0.1 - 1.5))))))
                        v20 := GenS4{F0: 1}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            for fi22 := 0; fi22 < 20; fi22++ {
                                v13 = 1.5
                                var v23 int32 = -416
                                _ = fi22
                                _ = v23
                            }
                            _ = fi21
                        }
                    }
                }
            }
        }
    }
    fmt.Println(v18, v14)
    var v24 int16 = 0
    v25 := GenS5{F0: uint8(v18), F1: 606, F2: (true || (!false))}
    var v26 int8 = 100
    if false {
        v3 += 221
        v13 = -2.25
    } else {
        var v27 int16 = v24
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
