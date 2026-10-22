package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 int) uint8 {
    if (3.14159 >= 0.1) {
        fmt.Println(p0)
        var v7 uint64 = 2
        _ = v7
    } else {
        p0 = -1
        v8 := []int{(p0 / ((1) | 1)), p0, p0}
        v8 = append(v8, p0)
        _ = v8
    }
    var v9 bool = true
    var v10 int32 = 42
    var v11 int8 = 42
    v12 := []int{p0, p0, p0}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 2
}
func genFn1(p0 int8, p1 uint8, p2 int32) uint16 {
    fmt.Println(p1, p1, p2)
    switch ("hello" + "a") {
    case "\n\t":
        v13 := GenS1{F0: p1, F1: 420}
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 4; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := GenS1{F0: (p1 * (p1 * genFn0(fi14))), F1: 4294967295}
                var v17 int = fi15
                v18 := []int{-1, v17, 116, (v17 * 1)}
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
    }
    p2 = p2
    v19 := []int{674}
    var v20 int64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = -32768
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    var v5 string = ("a" + "")
    var v6 uint16 = 796
    var v21 int16 = v2
    if ((3.14159 != ((float32((float64(((float32(0.0) - (-2.25 + 1.5)) * (3.14159 * (3.14159 - 0.0)))) + 0.1)) / 3.14159) / 1.0)) || false) {
        switch (v5 + (((v5 + v5) + ("\n\t" + v5)) + v5)) {
        case "\n\t":
            var v22 float32 = 0.1
            _ = v22
        default:
            if false {
                v23 := v4[1]
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := []int{}
            _ = fi24
            _ = v25
        }
        var v26 uint16 = v6
        _ = v26
    }
    switch ("hello" + "a") {
    case "hello":
        var v27 uint = 1
        _ = v27
    case "\n\t":
        v28 := v1["foo bar"]
        _ = v28
    }
    fmt.Println(v6)
    v29 := []int{128, 32768}
    if (v6 <= 0) {
        var v30 string = ("\n\t" + "foo bar")
        _ = v30
    }
    if true {
        var v31 int = -2147483648
        _ = v31
    } else {
        v32 := v3[2147483648]
        _ = v32
    }
    var v33 uint8 = 29
    _ = len(v1)
    fmt.Println(v6, v2, v21)
    var v34 float64 = 3.14159
    fmt.Println(v21, v6, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
