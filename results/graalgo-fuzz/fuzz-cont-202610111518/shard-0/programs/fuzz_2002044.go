package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint }
func genFn0(p0 uint16, p1 int, p2 uint8) uint8 {
    if true {
        p1 += p1
        v7 := GenS1{F0: 255, F1: (^(p2)), F2: "a"}
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS2{F0: "hello", F1: 0}
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "\n\t":
        v10 := GenS3{F0: 1, F1: 1}
        _ = v10
    }
    p1 -= -2147483648
    for fi11 := 0; fi11 < 3; fi11++ {
        p2 = 0
        p0 += 2
        if (128 == 255) {
            p0 += 365
            var v12 bool = ((p1 - 42) <= (p1 - 32768))
            _ = v12
        } else {
            p2 = (p2 % ((p2) | 1))
            var v13 uint64 = 333
            var v14 uint32 = 2
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 + uint8(p0))
}
func genFn1(p0 int16) uint8 {
    if true {
        if false {
            switch -43 {
            case 128:
                v15 := GenS3{F0: 18446744073709551615, F1: 197}
                _ = v15
            }
        }
        var v16 uint32 = 1
        _ = v16
    }
    var v17 uint16 = 1
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 string = ("a" + "\n\t")
        _ = fi18
        _ = v19
    }
    if ((-2.25 - 1.5) >= -2.25) {
        switch (("" + ("hello" + "a")) + "foo bar") {
        case "\n\t":
            var v20 int16 = 32767
            _ = v20
        case "a":
            var v21 uint8 = 0
            _ = v21
        }
    } else {
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := []int{(fi22 << ((42) & 63)), -578}
            var v24 int8 = 127
            v25 := GenS4{F0: 2, F1: (^(v24)), F2: 725}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        v26 := []int{127}
        v17 -= (v17 ^ (v17 / ((0) | 1)))
        _ = v26
    }
    _ = p0
    _ = v17
    return 42
}

func main() {
    var v1 int = -421
    v2 := []string{((((("a" + "") + ("\n\t" + "foo bar")) + "a") + "a") + ("" + "a")), ("a" + ""), "\n\t", "foo bar"}
    var v3 uint32 = 264
    v4 := []int{(^(v1)), v1, (v1 / ((v1) | 1))}
    var v5 string = "hello"
    var v6 uint64 = 899
    var v27 uint64 = v6
    if (42 != -32768) {
        var v28 int64 = -9223372036854775808
        for fi29 := 0; fi29 < 12; fi29++ {
            fmt.Println(v28, v3, v27)
            v27 = v6
            _ = fi29
        }
        var v30 int64 = (v28 | -9223372036854775808)
        _ = v28
        _ = v30
    }
    var v31 uint32 = 4294967295
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println(v31)
    switch v1 {
    case 128:
        v6 -= v6
    case 0:
        v5 = v5
    }
    v32 := GenS3{F0: uint64(v3), F1: v27}
    fmt.Println(v27, v27)
    if (-2.25 > 3.14159) {
        var v33 float32 = 0.0
        fmt.Println(v33)
        _ = v33
    }
    var v34 int8 = -54
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
