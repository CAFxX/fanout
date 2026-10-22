package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 string }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
func genFn0() int {
    var v4 int8 = -128
    var v5 int8 = v4
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint, p1 string) int16 {
    fmt.Println(p0, p1)
    if (4294967295 >= uint64(p0)) {
        if true {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 string = (p1 + ((((p1 + "hello") + ("" + "")) + (p1 + p1)) + (p1 + "hello")))
                var v8 int = (fi6 & (fi6 * fi6))
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: p1}
            v10 := []int{0, 0, 0, -704}
            _ = v10
            _ = v9
        } else {
            if (0 == (p0 ^ p0)) {
                if false {
                    p1 = (p1 + (p1 + p1))
                }
            } else {
                v11 := GenS1{F0: 0, F1: (p0 >> ((p0) & 63)), F2: p1}
                var v12 float64 = 3.14159
                _ = v11
                _ = v12
            }
            v13 := []int{42, 32768}
            _ = v13
        }
        var v14 bool = false
        _ = v14
    } else {
        if true {
            p0 += p0
            v15 := GenS1{F0: 42, F1: 498, F2: p1}
            _ = v15
        }
        fmt.Println(p1, p0)
        if (127 < 61) {
            var v16 int = 1
            _ = v16
        }
    }
    var v17 float32 = (1.5 / -2.25)
    var v18 float32 = 1e10
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return -201
}

func main() {
    var v1 bool = (0 >= 1)
    v2 := []int{32768, 127, 2147483647}
    v3 := []uint32{485, 1, 172, 1}
    v1 = true
    var v19 uint16 = 155
    v19 -= 291
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v20 float32 = 3.14159
    var v21 bool = false
    v22 := GenS2{F0: 2, F1: 845}
    switch "\n\t" {
    case "a":
        v23 := GenS2{F0: 255, F1: -9223372036854775808}
        _ = v23
    case "foo bar":
        var v24 uint = 0
        _ = v24
    }
    if v21 {
        if len(v3) > 0 {
            v3[0] = 0
        }
        var v25 uint16 = 2
        var v26 float64 = (-2.25 + ((0.0 + 0.1) / (1e10 + (0.0 + (((1e10 + (((3.14159 * 0.1) + 1.5) + 1.5)) - (1.5 * ((1.5 / -2.25) * -2.25))) + 0.0)))))
        _ = v25
        _ = v26
    }
    v27 := GenS3{F0: 9223372036854775807, F1: 2, F2: 782}
    fmt.Println(v1, v21, v20)
    v19 = 1
    var v28 int64 = 1
    v20 += ((1.5 - -2.25) + (0.0 / v20))
    if (true && true) {
        if (100 != -1) {
            v20 = float32(float64(v20))
            for fi29 := 0; fi29 < 2; fi29++ {
                v30 := GenS1{F0: (fi29 / (((fi29 + -170)) | 1)), F1: 269, F2: (("foo bar" + "\n\t") + "hello")}
                var v31 int = (fi29 ^ 127)
                switch (-(v31)) {
                case 916:
                    var v32 int = 377
                    _ = v32
                case 651:
                    if ((0.1 + 1e10) > float32((3.14159 + 0.0))) {
                        var v33 uint8 = 92
                        _ = v33
                    } else {
                        for fi34 := 0; fi34 < 16; fi34++ {
                            v35 := GenS2{F0: 255, F1: (fi34 << ((722) & 63))}
                            _ = fi34
                            _ = v35
                        }
                    }
                default:
                    v1 = v21
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        } else {
            _ = len(v3)
        }
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        v3 = append(v3, 241)
        _ = fi36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
}
