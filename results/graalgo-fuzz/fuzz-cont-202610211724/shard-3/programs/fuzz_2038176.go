package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int32; F2 bool }
func genFn0(p0 uint32, p1 int, p2 uint32) string {
    var v7 int32 = 1
    if false {
        v8 := []int{p1, p1, p1, p1}
        for fi9 := 0; fi9 < 2; fi9++ {
            p2 = 2147483648
            _ = fi9
        }
        _ = v8
    }
    var v10 int64 = 1
    v11 := []int{p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v7
    return "foo bar"
}
func genFn1() uint64 {
    var v12 uint8 = 2
    if (false || (!true)) {
        var v13 int8 = 121
        _ = v13
    } else {
        v14 := []int{}
        var v15 int8 = 106
        v15 += v15
        _ = v14
        _ = v15
    }
    _ = v12
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 109
    v3 := make(map[int]bool)
    var v4 bool = false
    var v5 int64 = 1
    v6 := make(map[int]string)
    if v4 {
        fmt.Println(v2)
        var v16 uint32 = 2
        switch (("a" + "a") + "hello") {
        case "":
            var v17 int = 853
            _ = v17
        case "hello":
            var v18 string = "\n\t"
            _ = v18
        default:
            var v19 float32 = (1e10 - 1.5)
            _ = v19
        }
        _ = v16
    }
    v4 = v4
    for fi20 := 0; fi20 < 9; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v4 = false
            var v22 bool = (1 <= v2)
            switch ("a" + ("\n\t" + "foo bar")) {
            case "\n\t":
                v23 := v6[fi20]
                _ = v23
            case "hello":
                v2 += 128
            case "a":
                if ((-(v5)) != 366) {
                    v22 = (32767 != 158)
                }
                var v24 uint8 = (v2 | (^(v2)))
                v25 := GenS1{F0: v5}
                _ = v24
                _ = v25
            }
            _ = fi21
            _ = v22
        }
        v26 := GenS2{F0: true, F1: -731, F2: v4}
        _ = fi20
        _ = v26
    }
    var v27 int = -9223372036854775808
    if (!(2 != 223)) {
        v2 -= (v2 % (((v2 + 128)) | 1))
    }
    var v28 int8 = 127
    var v29 float32 = ((1e10 + -2.25) + 3.14159)
    for fi30 := 0; fi30 < 11; fi30++ {
        switch (fi30 | (fi30 / (((v27 >> ((65535) & 63))) | 1))) {
        case -345:
            if (genFn1() < 2) {
                v28 += (v28 << ((255) & 63))
                _ = len(v1)
                v31 := []int{}
                _ = v31
            } else {
                v32 := []int{(fi30 / ((v27) | 1)), v27}
                _ = v32
            }
        }
        _ = fi30
    }
    v5 += v5
    v27 = (v27 - (v27 & (v27 / ((v27) | 1))))
    v4 = ((float64(v29) * float64(v29)) < ((1e10 - float64(-2.25)) * -2.25))
    var v33 uint32 = 2147483647
    fmt.Println(v27)
    v2 = (^(v2))
    if true {
        for fi34 := 0; fi34 < 20; fi34++ {
            var v35 float64 = ((-2.25 - 0.0) - (float64(v29) / -2.25))
            _ = fi34
            _ = v35
        }
    }
    v36 := []int{(v27 >> ((4294967295) & 63))}
    fmt.Println(v27, v33, v27)
    v37 := GenS1{F0: (v5 - v5)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
