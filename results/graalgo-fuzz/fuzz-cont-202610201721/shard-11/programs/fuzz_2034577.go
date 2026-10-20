package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint; F1 int; F2 uint32 }
func genFn0() uint32 {
    var v4 uint = 0
    v4 -= 0
    var v5 int16 = 1
    var v6 float32 = float32(1.5)
    _ = v4
    _ = v5
    _ = v6
    return 11
}
func genFn1() uint16 {
    var v7 int8 = 20
    fmt.Println(v7)
    _ = v7
    return 2
}

func main() {
    var v1 int32 = -9
    v2 := make(map[int]string)
    var v3 int = -2147483648
    if (!(!true)) {
        if false {
            fmt.Println(v3, v3)
        }
        var v8 int = (v3 % ((v3) | 1))
        _ = v8
    }
    var v9 float32 = float32(-2.25)
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := GenS1{F0: -2147483648, F1: "a", F2: 99}
        var v12 int8 = 127
        _ = fi10
        _ = v11
        _ = v12
    }
    v9 = (0.1 * v9)
    var v13 bool = true
    v9 = ((((((v9 - v9) * v9) - 0.1) / v9) - v9) * v9)
    v14 := []int{}
    var v15 float64 = 3.14159
    v2[(v3 * (v3 + v3))] = "a"
    v13 = (405 == 1)
    fmt.Println(v1, v15, v9)
    if (!true) {
        var v16 int64 = -525
        var v17 int64 = v16
        _ = v16
        _ = v17
    }
    v15 -= v15
    v1 = 764
    var v18 uint = 4294967295
    v19 := GenS1{F0: (-(v3)), F1: "hello", F2: 104}
    switch v3 {
    case 128:
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS2{F0: (42 >= v1), F1: 65535}
            _ = fi20
            _ = v21
        }
    case 32767:
        if len(v14) > 0 {
            v14[0] = -1
        }
    case 15:
        var v22 float32 = v9
        _ = v22
    default:
        if (int8(v18) == -87) {
            v18 = v18
            if (1 == 712) {
                var v23 string = ("foo bar" + "hello")
                _ = v23
            }
        } else {
            switch (("a" + "hello") + "a") {
            case "\n\t":
                var v24 int32 = v1
                _ = v24
            case "hello":
                v25 := GenS1{F0: (v3 | (v3 << ((1) & 63))), F1: "", F2: 1}
                _ = v25
            default:
                for fi26 := 0; fi26 < 7; fi26++ {
                    var v27 float32 = v9
                    v18 -= (v18 & v18)
                    fmt.Println(v13, v27)
                    _ = fi26
                    _ = v27
                }
            }
            if v13 {
                var v28 int = v3
                v3 = (^(v3))
                v3 = v28
                _ = v28
            } else {
                var v29 string = ("a" + "\n\t")
                var v30 string = (v29 + v29)
                if v13 {
                    var v31 string = v30
                    _ = v31
                } else {
                    v32 := GenS1{F0: v3, F1: v30, F2: -128}
                    if v13 {
                        switch "a" {
                        case "foo bar":
                            var v33 uint64 = 0
                            _ = v33
                        }
                    } else {
                        if (v13 && (!v13)) {
                            var v34 int64 = 0
                            var v35 int8 = 5
                            _ = v34
                            _ = v35
                        } else {
                            v36 := GenS3{F0: 462, F1: 2}
                            _ = v36
                        }
                    }
                    v37 := GenS4{F0: 601, F1: -9223372036854775808, F2: 1}
                    _ = v32
                    _ = v37
                }
                _ = v29
                _ = v30
            }
            v38 := v2[0]
            _ = v38
        }
    }
    var v39 uint8 = 136
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v9
}
