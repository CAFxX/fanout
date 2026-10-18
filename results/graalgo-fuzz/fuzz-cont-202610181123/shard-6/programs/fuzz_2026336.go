package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 int64; F2 int8 }
func genFn0(p0 string) bool {
    v7 := []int{4294967295, 4294967295, 32768, 313}
    v8 := GenS1{F0: 2147483647}
    if false {
        fmt.Println(p0)
    }
    p0 = "hello"
    v7 = append(v7, -716)
    _ = p0
    _ = v7
    _ = v8
    return (32768 >= -2147483648)
}

func main() {
    var v1 uint8 = 215
    var v2 string = "a"
    v3 := []int{-9223372036854775808}
    var v4 uint32 = 65535
    var v5 bool = false
    v6 := []string{}
    if len(v6) > 0 {
        v6[0] = v2
    }
    var v9 int32 = 1
    for fi10 := 0; fi10 < 20; fi10++ {
        v6 = append(v6, "hello")
        _ = fi10
    }
    var v11 int8 = 1
    var v12 int32 = -2147483648
    v11 -= 1
    if (!v5) {
        fmt.Println(v12, v11)
        v2 = v2
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                v15 := GenS2{F0: -128, F1: v1}
                if (((1.5 * ((0.0 * ((1.5 - (3.14159 - -2.25)) / 3.14159)) / 1.0)) != float64(0.1)) && false) {
                    var v16 int16 = 1000
                    switch (fi14 & (fi14 >> ((897) & 63))) {
                    case 127:
                        var v17 uint64 = 42
                        _ = v17
                    case 42:
                        v18 := GenS3{F0: 255, F1: 2147483647}
                        _ = v18
                    default:
                        fmt.Println(v5)
                    }
                    _ = v16
                }
                v19 := GenS4{F0: uint8(v4), F1: fi14}
                _ = fi14
                _ = v15
                _ = v19
            }
            _ = fi13
        }
        var v20 int8 = (v11 - v11)
        _ = v20
    }
    v21 := GenS5{F0: (v11 * v11), F1: 255, F2: 0}
    for fi22 := 0; fi22 < 12; fi22++ {
        v3 = append(v3, (fi22 + -468))
        v2 = "hello"
        _ = fi22
    }
    var v23 uint32 = v4
    var v24 float32 = 3.14159
    if (4294967295 == 2) {
        v9 = v12
        v25 := GenS6{F0: 2, F1: (v11 << ((255) & 63))}
        v26 := GenS7{F0: 32767}
        _ = v25
        _ = v26
    }
    v24 += ((v24 * v24) * ((v24 * v24) / (float32(1.5) + 0.0)))
    if v5 {
        var v27 bool = ((v11 * v11) <= v11)
        if (!true) {
            v28 := GenS5{F0: v11, F1: 65535, F2: 42}
            switch 32767 {
            case 128:
                var v29 string = v2
                _ = v29
            case 32768:
                v30 := GenS8{F0: ("a" > "")}
                _ = v30
            }
            _ = v28
        } else {
            for fi31 := 0; fi31 < 15; fi31++ {
                v32 := GenS9{F0: 0, F1: 9223372036854775807, F2: 100}
                v5 = (!false)
                _ = fi31
                _ = v32
            }
        }
        _ = v27
    }
    v5 = ((v1 * 180) <= v1)
    var v33 int32 = v12
    v12 = v9
    var v34 string = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
