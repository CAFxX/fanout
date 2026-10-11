package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int; F1 int16 }
func genFn0() int {
    switch "hello" {
    case "":
        for fi4 := 0; fi4 < 6; fi4++ {
            for fi5 := 0; fi5 < 20; fi5++ {
                if false {
                    fmt.Println(fi5, fi4)
                    v6 := GenS1{F0: 2147483648, F1: 1}
                    _ = v6
                }
                var v7 string = "\n\t"
                switch -2147483648 {
                case 127:
                    if (42 <= -1) {
                        v8 := []int{(fi4 ^ (fi4 - 0))}
                        _ = v8
                    }
                default:
                    v7 = "foo bar"
                }
                _ = fi5
                _ = v7
            }
            var v9 bool = false
            _ = fi4
            _ = v9
        }
    case "foo bar":
        v10 := GenS1{F0: 544, F1: 42}
        _ = v10
    case "\n\t":
        v11 := []int{-2147483648}
        _ = v11
    }
    var v12 float64 = 3.14159
    _ = v12
    return 1
}
func genFn1(p0 string, p1 int8, p2 uint) uint16 {
    var v13 float64 = float64((-2.25 + 3.14159))
    v13 -= v13
    v13 += -2.25
    p1 += (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 65535
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]uint32)
    var v3 int64 = -856
    switch ("a" + "hello") {
    case "\n\t":
        fmt.Println(v1, v1)
    default:
        var v14 int8 = 100
        _ = v14
    }
    v15 := GenS1{F0: 2, F1: 524}
    switch "foo bar" {
    case "":
        v16 := GenS2{F0: v3}
        _ = v16
    case "a":
        var v17 int = 4294967295
        _ = v17
    case "foo bar":
        var v18 float32 = (float32(0.0) - (1.5 / 0.1))
        _ = v18
    default:
        var v19 string = ((("a" + ("foo bar" + "")) + "foo bar") + ("foo bar" + "foo bar"))
        _ = v19
    }
    var v20 string = "hello"
    var v21 float64 = 0.1
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS3{F0: (v20 > ((v20 + v20) + v20)), F1: 75}
        var v24 bool = (v20 != "")
        _ = fi22
        _ = v23
        _ = v24
    }
    fmt.Println(v3, v1, v20)
    v25 := []int{-802, -9223372036854775808}
    if len(v25) > 0 {
        v25[0] = 1
    }
    v21 = float64(0.0)
    if (((2 != 288) || false) || (68 >= 1)) {
        if false {
            if false {
                v21 = -2.25
                if true {
                    var v26 string = v20
                    _ = v26
                }
            } else {
                var v27 int16 = 32767
                fmt.Println(v27, v20)
                _ = v27
            }
            v1 = (v1 ^ v1)
        } else {
            v28 := GenS4{F0: 1, F1: 255, F2: uint64(v1)}
            _ = len(v25)
            _ = v28
        }
        var v29 int = -2147483648
        _ = v29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS5{F0: (v3 ^ (v3 + (^(v3)))), F1: 127}
        v32 := GenS6{F0: fi30, F1: -654}
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int8 = 101
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v33
}
