package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 uint64) uint64 {
    var v7 uint = 291
    switch "" {
    case "a":
        v8 := []int{4294967295, 953}
        _ = v8
    case "":
        v9 := GenS1{F0: 0}
        _ = v9
    case "\n\t":
        v10 := []int{42, 0}
        _ = v10
    }
    fmt.Println(p0, v7)
    switch 9223372036854775807 {
    case 0:
        switch ("a" + "a") {
        case "foo bar":
            fmt.Println(p0)
        case "\n\t":
            var v11 uint16 = 265
            _ = v11
        default:
            if ((true || false) && (139 > 372)) {
                if (42 < 0) {
                    var v12 float64 = 3.14159
                    var v13 uint64 = p0
                    _ = v12
                    _ = v13
                } else {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 uint = (v7 * (v7 + v7))
                        v16 := []int{fi14}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
                var v17 uint = v7
                v18 := []int{454, 575, 128}
                _ = v17
                _ = v18
            } else {
                var v19 string = "foo bar"
                var v20 int = 186
                var v21 uint = v7
                _ = v19
                _ = v20
                _ = v21
            }
        }
    }
    v7 -= v7
    _ = p0
    _ = v7
    return p0
}
func genFn1() uint16 {
    switch "a" {
    case "":
        v22 := []int{127, 0}
        _ = v22
    case "\n\t":
        var v23 int = -315
        _ = v23
    default:
        var v24 uint = 1
        _ = v24
    }
    var v25 uint8 = 128
    if false {
        var v26 uint64 = 4294967295
        _ = v26
    } else {
        v25 += v25
        var v27 uint32 = 1
        _ = v27
    }
    var v28 int64 = -9223372036854775808
    _ = v25
    _ = v28
    return 1
}

func main() {
    var v1 int16 = -477
    var v2 int = 42
    v3 := make(map[int]uint32)
    var v4 uint8 = 2
    v5 := make(map[string]int)
    v6 := make(map[string]uint32)
    fmt.Println(v1, v1, v1)
    v29 := GenS1{F0: -733}
    for fi30 := 0; fi30 < 17; fi30++ {
        v4 -= v4
        _ = fi30
    }
    switch (("\n\t" + ("" + (("" + "foo bar") + ("a" + "hello")))) + ("a" + "a")) {
    case "a":
        for fi31 := 0; fi31 < 15; fi31++ {
            v3[(v2 + (v2 ^ fi31))] = 862
            if true {
                v1 = v1
            }
            _ = fi31
        }
    case "foo bar":
        for fi32 := 0; fi32 < 16; fi32++ {
            var v33 uint8 = (v4 / (((v4 / ((v4) | 1))) | 1))
            v2 = fi32
            v34 := []int{v2, (v2 << ((0) & 63))}
            _ = fi32
            _ = v33
            _ = v34
        }
    case "hello":
        var v35 uint64 = 42
        _ = v35
    default:
        for fi36 := 0; fi36 < 15; fi36++ {
            v4 -= v4
            delete(v3, (v2 - 0))
            v1 += (v1 % (((^(v1))) | 1))
            _ = fi36
        }
    }
    for fi37 := 0; fi37 < 1; fi37++ {
        v2 += (-(fi37))
        _ = fi37
    }
    if (((!false) || true) && ((v4 / ((255) | 1)) != (v4 | (v4 & v4)))) {
        fmt.Println(v4, v2)
    } else {
        fmt.Println(v4)
    }
    v38 := v5[("foo bar" + (((("foo bar" + ("foo bar" + ("\n\t" + ""))) + (("" + "") + "")) + "foo bar") + "\n\t"))]
    delete(v6, "hello")
    v39 := []int{-742}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
