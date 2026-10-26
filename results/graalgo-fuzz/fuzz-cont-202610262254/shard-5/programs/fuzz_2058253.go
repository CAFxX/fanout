package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0(p0 string, p1 uint8) bool {
    v6 := GenS1{F0: 0, F1: 0}
    p1 -= 42
    v7 := GenS1{F0: 1, F1: 1}
    p0 = ("" + "hello")
    p0 = (p0 + "")
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return false
}
func genFn1() int8 {
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int32 = -1
        var v11 int = (fi9 - (fi9 / (((fi9 ^ 0)) | 1)))
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint16 = 1
    v8 = 1
    _ = v12
    _ = v8
    return -1
}

func main() {
    var v1 int8 = -128
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    var v4 int8 = 42
    var v5 uint16 = 65535
    var v13 uint64 = 37
    switch (("foo bar" + "foo bar") + "hello") {
    case "\n\t":
        for fi14 := 0; fi14 < 8; fi14++ {
            v2["hello"] = 4294967295
            fmt.Println(v5, v4, v1)
            _ = fi14
        }
    case "foo bar":
        v13 = v13
    }
    v4 = v4
    if true {
        var v15 bool = false
        _ = v15
    } else {
        switch "" {
        case "hello":
            if false {
                var v16 uint32 = 1
                _ = v16
            } else {
                for fi17 := 0; fi17 < 3; fi17++ {
                    delete(v3, fi17)
                    _ = fi17
                }
            }
        }
        v18 := []int{}
        _ = v18
    }
    var v19 bool = false
    switch "hello" {
    case "hello":
        var v20 uint16 = (^(v5))
        _ = v20
    case "\n\t":
        v21 := []int{128, 128, 426, -4}
        _ = v21
    case "a":
        switch 42 {
        case 1:
            v2["a"] = 820
        case 811:
            var v22 int16 = -1
            _ = v22
        default:
            var v23 float64 = (3.14159 - 0.0)
            _ = v23
        }
    }
    v24 := GenS2{F0: 430, F1: -1}
    if ((874 >= 1000) || v19) {
        var v25 int64 = -9223372036854775808
        var v26 int64 = v25
        _ = v25
        _ = v26
    }
    v13 = v13
    switch ("hello" + ((("\n\t" + "hello") + ("hello" + "")) + "")) {
    case "foo bar":
        var v27 int16 = -32768
        _ = v27
    case "\n\t":
        v28 := v2[("a" + ((("" + ("foo bar" + "\n\t")) + "\n\t") + "\n\t"))]
        _ = v28
    }
    switch 32768 {
    case -239:
        v29 := GenS2{F0: (^(v13)), F1: 9223372036854775807}
        _ = v29
    case 0:
        var v30 bool = (!true)
        _ = v30
    }
    fmt.Println(v13, v19)
    v31 := GenS2{F0: v13, F1: 1}
    v4 += (v1 / ((v4) | 1))
    for fi32 := 0; fi32 < 13; fi32++ {
        v33 := []int{(fi32 << ((321) & 63)), (fi32 - 4294967295), fi32, (^(fi32))}
        _ = fi32
        _ = v33
    }
    var v34 uint32 = 437
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
}
