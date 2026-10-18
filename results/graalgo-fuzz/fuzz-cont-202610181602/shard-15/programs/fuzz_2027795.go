package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }

func main() {
    var v1 string = ""
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 uint8 = 128
    var v5 uint = 4294967295
    var v6 int32 = 981
    if false {
        var v7 uint32 = 2147483648
        _ = v7
    } else {
        fmt.Println(v1, v5, v1)
        var v8 uint32 = 42
        var v9 uint64 = 1
        _ = v8
        _ = v9
    }
    fmt.Println(v1)
    switch v1 {
    case "a":
        if (!(v5 != v5)) {
            var v10 uint8 = v4
            v5 = (v5 & 1)
            _ = v10
        } else {
            v11 := GenS1{F0: -9223372036854775808, F1: (v5 * (v5 ^ (v5 * v5))), F2: 127}
            v1 = (v1 + v1)
            _ = len(v3)
            _ = v11
        }
    case "\n\t":
        var v12 int16 = -1
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{42, (fi13 >> (((v5 & uint(v4))) & 63)), 32768, (^(fi13))}
        v15 := v3[2147483647]
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 = v4
    var v16 float64 = ((0.1 - (-2.25 * 1.5)) * 0.0)
    for fi17 := 0; fi17 < 3; fi17++ {
        var v18 uint8 = v4
        if ((!(433 <= 1)) && false) {
            var v19 float64 = 1e10
            _ = v19
        }
        var v20 int8 = -128
        _ = fi17
        _ = v18
        _ = v20
    }
    var v21 bool = (((!false) || (42 <= 1000)) || ((3.14159 / 1.0) <= 1.5))
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS2{F0: (0 < 2), F1: 1000, F2: 65535}
        var v24 bool = ((0 <= (^(v5))) || v21)
        if (v6 < v6) {
            v5 -= 65535
            _ = len(v2)
        }
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int = 42
    if v21 {
        if v21 {
            if false {
                var v26 float64 = v16
                v27 := GenS1{F0: (v25 / (((v25 + (v25 % (((v25 % (((v25 << ((v5) & 63))) | 1))) | 1)))) | 1)), F1: v5, F2: (v25 >> ((v5) & 63))}
                var v28 string = ("a" + v1)
                _ = v26
                _ = v27
                _ = v28
            } else {
                var v29 int8 = 127
                v4 = (v4 % ((255) | 1))
                _ = v29
            }
        } else {
            v16 += -2.25
        }
        fmt.Println(v5, v16)
    } else {
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := []int{(-(v25)), -2147483648}
            _ = fi30
            _ = v31
        }
        var v32 int64 = -9223372036854775808
        switch v1 {
        case "foo bar":
            v33 := GenS2{F0: (65535 < 985), F1: -1, F2: 506}
            _ = v33
        case "":
            if (v21 || (95 <= 0)) {
                v6 = v6
                var v34 float32 = (3.14159 * (3.14159 + 0.1))
                _ = v34
            } else {
                var v35 uint64 = 234
                switch (("foo bar" + v1) + (((v1 + v1) + "a") + "")) {
                case "hello":
                    var v36 uint16 = 136
                    _ = v36
                case "foo bar":
                    var v37 uint32 = 1
                    _ = v37
                case "a":
                    fmt.Println(v25, v16)
                default:
                    v32 -= 115
                }
                v5 = (^(v5))
                _ = v35
            }
        case "\n\t":
            var v38 uint8 = v4
            _ = v38
        default:
            if ((^(v6)) != v6) {
                v21 = ((float64(((-2.25 + ((float32(((v16 / 0.1) + (v16 / v16))) - 0.1) / 3.14159)) / ((0.0 + float32(float64(0.1))) * float32(v16)))) + 1.5) != 1.5)
            }
        }
        _ = v32
    }
    v39 := GenS2{F0: ((((v1 + v1) > "\n\t") && (true || ((v6 - 42) != (-(v6))))) && (1 != 1)), F1: 0, F2: 65535}
    v1 = v1
    var v40 int32 = v6
    _ = len(v2)
    v41 := GenS1{F0: (v25 << ((uint(v4)) & 63)), F1: 265, F2: v25}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
