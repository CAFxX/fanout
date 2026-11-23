package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int; F2 int16 }
type GenS5 struct { F0 int64; F1 uint; F2 uint }
func genFn0(p0 uint64) bool {
    if (!(!true)) {
        v7 := []int{4294967295, -9223372036854775808, 4294967295}
        p0 = p0
        _ = v7
    }
    p0 = (p0 | (p0 ^ (p0 & 2)))
    _ = p0
    return (2 >= 42)
}
func genFn1(p0 int64, p1 int) int64 {
    var v8 string = "foo bar"
    var v9 int8 = -128
    v10 := []int{(p1 / ((9223372036854775807) | 1))}
    var v11 bool = (false && (!true))
    var v12 float32 = (0.0 + 0.1)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []uint32{2, 42}
    var v2 uint8 = 1
    var v3 uint32 = 0
    var v4 int = 1
    var v5 float64 = 3.14159
    v6 := []uint32{(^(v3)), 0}
    var v13 string = "foo bar"
    switch v4 {
    case 32768:
        var v14 int32 = -77
        _ = v14
    case 2147483647:
        switch v4 {
        case -984:
            fmt.Println(v5, v2)
        }
    case 252:
        var v15 string = (v13 + v13)
        _ = v15
    }
    for fi16 := 0; fi16 < 20; fi16++ {
        switch (v13 + ((("\n\t" + ("hello" + "\n\t")) + (("foo bar" + ("" + "hello")) + "a")) + ("foo bar" + (v13 + "")))) {
        case "hello":
            var v17 uint64 = 4294967295
            _ = v17
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v18 uint32 = 1
            _ = v18
        }
        fmt.Println(fi16, v13)
        for fi19 := 0; fi19 < 10; fi19++ {
            v3 = (^(v3))
            v1 = append(v1, v3)
            _ = fi19
        }
        _ = fi16
    }
    switch (v13 + ((v13 + "") + "hello")) {
    case "\n\t":
        var v20 uint32 = v3
        _ = v20
    case "a":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case "hello":
        v21 := GenS1{F0: uint(v3), F1: (v13 + ""), F2: 2}
        _ = v21
    default:
        var v22 uint = 65535
        _ = v22
    }
    switch v13 {
    case "hello":
        for fi23 := 0; fi23 < 12; fi23++ {
            if true {
                fmt.Println(v3)
                v24 := GenS2{F0: -1}
                v4 = (-(v4))
                _ = v24
            } else {
                v4 += int(v3)
                if true {
                    v6 = append(v6, v3)
                    v3 -= v3
                    var v25 int32 = 2147483647
                    _ = v25
                }
                v5 = v5
            }
            _ = fi23
        }
    }
    fmt.Println(v13, v13, v5)
    var v26 uint = 2
    if (!(!(true || true))) {
        var v27 int64 = genFn1(genFn1(42, v4), (^(v4)))
        v28 := GenS3{F0: v2, F1: genFn1((^(v27)), (v4 * (v4 * v4))), F2: 65535}
        v29 := GenS4{F0: (v3 % ((v3) | 1)), F1: 42, F2: 42}
        _ = v27
        _ = v28
        _ = v29
    }
    switch v13 {
    case "a":
        switch ((((("\n\t" + v13) + v13) + (v13 + "hello")) + (v13 + v13)) + v13) {
        case "\n\t":
            v3 += v3
        case "hello":
            if (-2147483648 <= v4) {
                var v30 uint32 = 65535
                var v31 float32 = 0.0
                var v32 uint64 = 1
                _ = v30
                _ = v31
                _ = v32
            }
        default:
            fmt.Println(v3, v3, v26)
        }
    case "\n\t":
        v33 := GenS5{F0: 0, F1: v26, F2: (v26 / ((2) | 1))}
        _ = v33
    }
    var v34 uint64 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
