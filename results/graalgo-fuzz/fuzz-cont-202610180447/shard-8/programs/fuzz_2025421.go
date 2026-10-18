package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint }
func genFn0(p0 int8) int32 {
    for fi4 := 0; fi4 < 5; fi4++ {
        switch fi4 {
        case 127:
            switch ("hello" + (((((("" + "foo bar") + "") + ("hello" + "a")) + "hello") + "") + ("foo bar" + ""))) {
            case "hello":
                var v5 float64 = 1.5
                _ = v5
            case "foo bar":
                var v6 string = "\n\t"
                _ = v6
            }
        case 128:
            var v7 int = fi4
            _ = v7
        case 1:
            v8 := []int{}
            _ = v8
        }
        p0 -= p0
        _ = fi4
    }
    switch (((("hello" + "") + "") + ("a" + (("\n\t" + "") + "hello"))) + "hello") {
    case "\n\t":
        if false {
            p0 -= 77
        } else {
            fmt.Println(p0)
            v9 := []int{-2147483648, 785, 128, 42}
            fmt.Println(p0)
            _ = v9
        }
    }
    var v10 uint16 = 347
    p0 += (p0 / ((p0) | 1))
    v10 = (v10 + (v10 / ((1) | 1)))
    _ = p0
    _ = v10
    return 1
}
func genFn1(p0 uint32) uint32 {
    v11 := GenS1{F0: 22, F1: "a", F2: 255}
    for fi12 := 0; fi12 < 14; fi12++ {
        fmt.Println(p0, fi12)
        _ = fi12
    }
    if ((!false) && true) {
        v13 := GenS2{F0: 0}
        var v14 int8 = 42
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v11
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (float64((float32(0.0) / 1e10)) / float64(float32(-2.25)))
    v3 := []string{("\n\t" + "a"), "\n\t"}
    var v15 int32 = -2147483648
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS3{F0: (((3.14159 + float64((1e10 / (3.14159 - (((-2.25 * 0.0) * -2.25) + (-2.25 + 1e10)))))) / v2) >= float64((1.5 - 1e10))), F1: (fi16 - fi16)}
        v15 += (v15 ^ 2147483647)
        _ = fi16
        _ = v17
    }
    v18 := GenS4{F0: 127, F1: 4294967295}
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 uint16 = 606
        if false {
            if ((!(!(574 == 1))) && ((false || (!(!(!false)))) || false)) {
                v21 := GenS1{F0: -56, F1: "\n\t", F2: 0}
                v15 -= (v15 << ((4294967295) & 63))
                _ = v21
            } else {
                fmt.Println(v15)
            }
            v2 += float64(-2.25)
            v15 += 2147483647
        }
        _ = fi19
        _ = v20
    }
    v2 = v2
    v22 := GenS5{F0: 65535}
    if (2 < 171) {
        fmt.Println(v15, v2)
    }
    v2 = float64((((1.5 * -2.25) * -2.25) * 0.0))
    v23 := GenS5{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
