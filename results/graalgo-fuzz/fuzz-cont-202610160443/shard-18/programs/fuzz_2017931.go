package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int }
func genFn0(p0 bool, p1 int8) uint64 {
    switch "" {
    case "foo bar":
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := []int{127, fi5, (fi5 & (^(fi5))), fi5}
            _ = fi5
            _ = v6
        }
    case "hello":
        fmt.Println(p0, p0)
    }
    var v7 uint64 = 42
    var v8 uint64 = v7
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = (0.0 - (0.0 / 3.14159))
    var v3 uint64 = 347
    var v4 string = "a"
    var v9 uint = 656
    if false {
        fmt.Println(v2)
        if false {
            v10 := GenS1{F0: 58}
            for fi11 := 0; fi11 < 19; fi11++ {
                fmt.Println(v4, v1, v9)
                var v12 uint64 = v3
                v13 := GenS2{F0: (-653 == -872), F1: v3, F2: fi11}
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
        switch (v4 + v4) {
        case "":
            var v14 int8 = -1
            _ = v14
        case "foo bar":
            v1 += (v1 + v1)
        case "hello":
            if (615 < uint16(v9)) {
                if (((v4 + v4) + (v4 + "hello")) > "foo bar") {
                    v2 = 0.0
                } else {
                    var v15 int64 = (v1 / ((v1) | 1))
                    v16 := []int{-9223372036854775808, 42, 4294967295}
                    var v17 int = 32767
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
        default:
            v9 = v9
        }
    }
    v2 = v2
    var v18 uint64 = (^(v3))
    if (-2147483648 >= 231) {
        fmt.Println(v3, v3)
        for fi19 := 0; fi19 < 7; fi19++ {
            switch (("" + (v4 + (v4 + v4))) + ((v4 + ("hello" + ((v4 + ((v4 + ((v4 + v4) + v4)) + ("" + v4))) + ((v4 + ((v4 + ("\n\t" + v4)) + v4)) + (v4 + ((v4 + "a") + "foo bar")))))) + (v4 + ("\n\t" + v4)))) {
            case "foo bar":
                if true {
                    var v20 bool = false
                    if (v2 <= v2) {
                        var v21 bool = false
                        var v22 int32 = -1
                        _ = v21
                        _ = v22
                    }
                    _ = v20
                } else {
                    var v23 uint64 = (v18 + (v18 * (v18 & v3)))
                    var v24 int16 = int16(fi19)
                    _ = v23
                    _ = v24
                }
            case "hello":
                var v25 int16 = -106
                _ = v25
            case "":
                var v26 uint8 = 128
                _ = v26
            }
            _ = fi19
        }
    }
    switch (("foo bar" + v4) + v4) {
    case "":
        v2 += v2
    }
    fmt.Println(v18)
    for fi27 := 0; fi27 < 10; fi27++ {
        v28 := []int{(^(fi27))}
        v2 -= v2
        v28 = append(v28, (fi27 << (((v9 % (((v9 / ((v9) | 1))) | 1))) & 63)))
        _ = fi27
        _ = v28
    }
    v29 := []int{-257, 826}
    var v30 int64 = (v1 ^ (^(v1)))
    if (((^(v9)) < (^(v9))) && (!true)) {
        fmt.Println(v18, v30, v4)
    }
    v29 = append(v29, -1)
    if (((1.5 + ((float64(0.1) - 0.0) + 1e10)) + float64((1e10 * 1e10))) > 0.0) {
        var v31 float64 = (float64((float32(-2.25) - v2)) + float64(v2))
        _ = v31
    }
    v32 := GenS1{F0: 128}
    var v33 uint8 = 42
    _ = len(v29)
    var v34 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v9
}
