package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint16) string {
    v7 := []int{42, 32768, 9223372036854775807, 111}
    var v8 uint32 = 2
    var v9 int16 = -32768
    v7 = append(v7, 0)
    p0 = 102
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ""
}
func genFn1(p0 uint8) bool {
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := GenS1{F0: 42, F1: 1, F2: (!true)}
        var v12 int = (fi10 >> ((2) & 63))
        _ = fi10
        _ = v11
        _ = v12
    }
    p0 += (p0 * 42)
    _ = p0
    return (!false)
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 46
    var v5 bool = (("" + ("foo bar" + "")) >= "hello")
    v6 := []uint32{2, 399, 4294967295}
    fmt.Println(v5)
    fmt.Println(v4)
    if (0.1 != 1.5) {
        v2[-9223372036854775808] = 631
        v13 := GenS1{F0: 351, F1: 65535, F2: (2 > 99)}
        _ = v13
    }
    v14 := GenS1{F0: 4294967295, F1: 0, F2: (2 <= 128)}
    if true {
        switch "" {
        case "a":
            v5 = ((0 >= 18446744073709551615) || v5)
        case "\n\t":
            var v15 uint8 = 1
            _ = v15
        }
        v5 = (2 >= 128)
    } else {
        var v16 int32 = 1
        for fi17 := 0; fi17 < 5; fi17++ {
            switch "a" {
            case "":
                v5 = v5
            }
            fmt.Println(fi17, v16, fi17)
            _ = fi17
        }
        _ = v16
    }
    switch "a" {
    case "hello":
        switch "hello" {
        case "foo bar":
            var v18 int32 = -1
            _ = v18
        case "hello":
            if false {
                v4 = (-(v4))
            } else {
                var v19 int16 = 619
                fmt.Println(v5)
                _ = v19
            }
        case "":
            v2[-1] = 2
        }
    case "foo bar":
        for fi20 := 0; fi20 < 12; fi20++ {
            v5 = (true && true)
            for fi21 := 0; fi21 < 2; fi21++ {
                v22 := GenS1{F0: 541, F1: 18446744073709551615, F2: v5}
                var v23 int32 = -2147483648
                var v24 uint = 2
                if (!((v4 - 1) < -10)) {
                    fmt.Println(v5, v4)
                } else {
                    if true {
                        v3[42] = 866
                        var v25 uint8 = 255
                        var v26 float64 = 1.5
                        _ = v25
                        _ = v26
                    } else {
                        var v27 float64 = 0.0
                        _ = v27
                    }
                }
                v28 := GenS1{F0: 635, F1: 65535, F2: v5}
                v5 = v5
                var v29 uint8 = 127
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
                _ = v28
                _ = v29
            }
            _ = fi20
        }
    case "":
        v30 := GenS2{F0: 42}
        _ = v30
    default:
        switch (("foo bar" + "\n\t") + "a") {
        case "hello":
            v4 = (v4 % ((7) | 1))
        case "\n\t":
            var v31 uint16 = 0
            _ = v31
        }
    }
    var v32 int16 = -849
    fmt.Println(v4, v4, v4)
    fmt.Println(v5, v32, v32)
    var v33 bool = ((0.1 * float32((((0.1 - 0.0) / ((1.5 * 3.14159) / 3.14159)) - 1e10))) <= 0.0)
    var v34 int8 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
