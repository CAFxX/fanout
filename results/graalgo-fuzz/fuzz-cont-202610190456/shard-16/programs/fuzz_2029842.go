package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 uint8, p1 uint8, p2 int32) uint32 {
    fmt.Println(p0, p1)
    if (p1 <= (p1 & 128)) {
        p1 -= 0
        var v6 int = -2147483648
        v7 := []int{}
        _ = v6
        _ = v7
    }
    p0 = 135
    var v8 uint = 234
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 2147483647
}
func genFn1(p0 uint) string {
    p0 += 940
    var v9 int = 0
    var v10 uint16 = 65535
    v10 = (v10 & (^(v10)))
    if true {
        v11 := []int{(v9 >> ((p0) & 63))}
        var v12 uint16 = (v10 ^ (v10 / ((32767) | 1)))
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v9
    return "\n\t"
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]bool)
    var v3 string = ((("foo bar" + "\n\t") + "a") + ("hello" + ((("hello" + "") + "a") + "hello")))
    var v4 uint32 = 42
    v5 := []uint32{v4}
    fmt.Println(v1, v1, v4)
    v13 := GenS1{F0: 128}
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(v3)
            fmt.Println(v1)
            if len(v5) > 0 {
                v5[0] = 121
            }
            switch "a" {
            case "hello":
                v3 = "hello"
            case "\n\t":
                if false {
                    fmt.Println(fi15)
                } else {
                    var v16 int32 = 0
                    _ = v16
                }
            case "a":
                v4 -= (v4 | (v4 ^ v4))
            default:
                if (0 <= 2) {
                    v5 = append(v5, v4)
                    v17 := v2[fi15]
                    var v18 uint = 840
                    _ = v17
                    _ = v18
                }
            }
            v4 = (v4 - v4)
            _ = fi15
        }
        switch (v3 + (("foo bar" + "foo bar") + v3)) {
        case "":
            v5 = append(v5, 42)
        case "hello":
            var v19 float32 = 1e10
            _ = v19
        case "\n\t":
            v1 -= (v1 - 0.0)
        }
        switch (fi14 / ((-9223372036854775808) | 1)) {
        case 9223372036854775807:
            v20 := GenS1{F0: 19}
            _ = v20
        case 32767:
            var v21 int8 = 127
            _ = v21
        default:
            _ = len(v2)
        }
        _ = fi14
    }
    var v22 int64 = 809
    for fi23 := 0; fi23 < 16; fi23++ {
        v24 := v2[9223372036854775807]
        if (2147483647 == 869) {
            fmt.Println(fi23)
            var v25 uint16 = 0
            var v26 int32 = 42
            _ = v25
            _ = v26
        } else {
            v5 = append(v5, v4)
        }
        _ = len(v2)
        _ = fi23
        _ = v24
    }
    if (false && true) {
        for fi27 := 0; fi27 < 3; fi27++ {
            v28 := GenS1{F0: 150}
            var v29 int8 = 22
            _ = fi27
            _ = v28
            _ = v29
        }
        v22 += 487
    }
    v30 := GenS1{F0: 127}
    v31 := GenS2{F0: ((true && true) && (false || (-54 >= 1))), F1: 0, F2: 2}
    var v32 uint8 = 22
    var v33 int = 1
    var v34 string = v3
    v3 = (v3 + v34)
    switch "a" {
    case "\n\t":
        v3 = "a"
    }
    for fi35 := 0; fi35 < 5; fi35++ {
        v34 = v34
        var v36 int = 9223372036854775807
        for fi37 := 0; fi37 < 15; fi37++ {
            if len(v5) > 0 {
                v5[0] = genFn0(v32, uint8(v4), 42)
            }
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    v32 -= 255
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
