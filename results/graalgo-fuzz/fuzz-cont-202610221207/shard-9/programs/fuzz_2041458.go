package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint32 }
func genFn0(p0 int32) int16 {
    p0 = (p0 * (p0 & (^(p0))))
    v6 := GenS1{F0: 9223372036854775807}
    p0 += p0
    _ = p0
    _ = v6
    return 51
}
func genFn1(p0 int16, p1 uint8) int64 {
    if true {
        if false {
            var v7 int64 = 1
            var v8 int32 = -1
            v9 := GenS1{F0: (v7 % (((v7 - (v7 >> ((65535) & 63)))) | 1))}
            _ = v7
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 755}
        _ = v10
    } else {
        fmt.Println(p1, p1)
        v11 := GenS1{F0: 232}
        _ = v11
    }
    switch ("foo bar" + "\n\t") {
    case "":
        v12 := []int{2147483647, -797, 826, 4294967295}
        _ = v12
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := []string{"a", "a", "a", ("" + ("foo bar" + "foo bar"))}
    var v2 float32 = ((0.0 - 1e10) / 1.5)
    v3 := []uint32{}
    var v4 int64 = -328
    v5 := []string{("a" + "hello"), "foo bar"}
    var v13 uint = 42
    for fi14 := 0; fi14 < 10; fi14++ {
        v3 = append(v3, 2)
        _ = fi14
    }
    v13 = 633
    var v15 float32 = float32((((0.1 / ((3.14159 / -2.25) * float64(v2))) + (3.14159 + 0.0)) / 1.5))
    var v16 uint = v13
    fmt.Println(v15, v15)
    var v17 float64 = (0.1 - 3.14159)
    fmt.Println(v17, v16, v17)
    for fi18 := 0; fi18 < 6; fi18++ {
        switch (("hello" + "") + "\n\t") {
        case "\n\t":
            var v19 uint64 = 42
            _ = v19
        default:
            fmt.Println(v15, fi18)
        }
        var v20 float32 = v15
        v21 := GenS1{F0: v4}
        _ = fi18
        _ = v20
        _ = v21
    }
    var v22 int8 = -128
    for fi23 := 0; fi23 < 8; fi23++ {
        if (449 >= 42) {
            _ = len(v1)
            fmt.Println(fi23)
        }
        _ = fi23
    }
    var v24 int8 = (v22 << ((18446744073709551615) & 63))
    fmt.Println(v13, v2)
    var v25 uint = v13
    for fi26 := 0; fi26 < 5; fi26++ {
        fmt.Println(v15, v25)
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS1{F0: 0}
            fmt.Println(v17, v13, fi26)
            switch "a" {
            case "a":
                v24 = (v24 & (v22 | -1))
            case "foo bar":
                v29 := GenS1{F0: v4}
                switch ("" + "") {
                case "\n\t":
                    v30 := GenS1{F0: v4}
                    v2 = (0.1 + 1.5)
                    var v31 float32 = (v2 * 0.1)
                    v32 := GenS1{F0: (v4 - 719)}
                    _ = v30
                    _ = v31
                    _ = v32
                case "a":
                    if ((((v4 * (v4 >> (((v13 * 4294967295)) & 63))) <= -827) && false) && true) {
                        v17 += (3.14159 - ((0.1 + (v17 * (float64((-2.25 + (((1e10 - v15) / float32(v17)) * 1e10))) * (((1e10 - (((1e10 / ((float64((v15 + (0.1 - 1e10))) + 1.5) - (((float64(((v15 + ((0.0 / v2) / v2)) - (float32(v17) * v15))) * v17) - 0.0) / v17))) - (1e10 / ((v17 * (3.14159 * v17)) * (v17 + v17)))) + v17)) / v17) - ((v17 * (v17 * -2.25)) + v17))))) * v17))
                        var v33 int16 = -32768
                        _ = v33
                    } else {
                        var v34 int = (fi27 & fi27)
                        fmt.Println(v13)
                        _ = v34
                    }
                    fmt.Println(v4)
                    switch fi26 {
                    case -2147483648:
                        fmt.Println(fi26, v24)
                    }
                default:
                    var v35 int64 = v4
                    _ = v35
                }
                v36 := GenS2{F0: (v24 & v22), F1: 864}
                _ = v29
                _ = v36
            case "hello":
                var v37 uint8 = 128
                _ = v37
            }
            _ = fi27
            _ = v28
        }
        var v38 uint = (v25 * v13)
        _ = fi26
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
