package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int32; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
func genFn0() uint8 {
    if ((((1e10 + -2.25) / ((1.5 + 3.14159) + float64(0.1))) * -2.25) == (1e10 - 1.5)) {
        var v4 int64 = 0
        v5 := []int{351, 127, 1, 32767}
        _ = v4
        _ = v5
    } else {
        var v6 uint16 = 349
        _ = v6
    }
    if true {
        if (127 > 211) {
            v7 := []int{-9223372036854775808, -657, -1}
            _ = v7
        } else {
            var v8 string = "a"
            _ = v8
        }
        v9 := []int{-223, -2147483648}
        switch "\n\t" {
        case "\n\t":
            v10 := GenS1{F0: true, F1: 2147483648}
            _ = v10
        }
        _ = v9
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: ((-1 < -1) || false), F1: 65535}
        _ = fi11
        _ = v12
    }
    return 145
}
func genFn1(p0 int16, p1 uint) uint64 {
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := []int{(fi13 - fi13), (fi13 << ((93) & 63)), -9223372036854775808}
        _ = fi13
        _ = v14
    }
    p1 = 806
    if (-2.25 > 0.0) {
        fmt.Println(p0)
        v15 := GenS1{F0: false, F1: 65535}
        v16 := GenS1{F0: true, F1: 0}
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int8 = 127
    var v2 int64 = -1
    v3 := []string{("\n\t" + (("\n\t" + "foo bar") + "foo bar")), "hello", "a", "a"}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v17 int = 32767
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v2, fi18, fi18)
        v17 = v17
        if len(v3) > 0 {
            v3[0] = (("\n\t" + "") + "")
        }
        _ = fi18
    }
    if (47 != -913) {
        for fi19 := 0; fi19 < 13; fi19++ {
            v1 = -1
            for fi20 := 0; fi20 < 15; fi20++ {
                v21 := GenS1{F0: false, F1: 2147483647}
                v17 -= (fi20 << ((457) & 63))
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
    } else {
        v22 := GenS1{F0: false, F1: 65535}
        var v23 string = (("a" + "\n\t") + ("\n\t" + "hello"))
        _ = v22
        _ = v23
    }
    if true {
        v24 := GenS2{F0: 100}
        v1 = (v1 | v1)
        _ = v24
    }
    switch "\n\t" {
    case "":
        v17 = (^(v17))
    default:
        var v25 uint32 = 2147483648
        _ = v25
    }
    if (false || true) {
        v26 := GenS3{F0: "a", F1: 1, F2: 1}
        v1 -= (v1 | -128)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v26
    } else {
        var v27 int8 = (-(v1))
        var v28 string = ("foo bar" + ("hello" + "a"))
        switch (("\n\t" + "foo bar") + ((("a" + ((v28 + v28) + (v28 + ("" + v28)))) + "hello") + ("" + "hello"))) {
        case "a":
            fmt.Println(v28, v27, v17)
        }
        _ = v27
        _ = v28
    }
    v29 := GenS4{F0: 42}
    var v30 int = (v17 - 42)
    var v31 uint64 = 2
    for fi32 := 0; fi32 < 1; fi32++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v31, fi32, v30)
        v17 = v17
        _ = fi32
    }
    v1 = (-(v1))
    v2 += v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v30 += (^(v17))
    var v33 uint64 = genFn1(-339, 18446744073709551615)
    v34 := GenS5{F0: (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
}
