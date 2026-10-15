package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 int32) uint8 {
    var v4 bool = (1 > -1)
    var v5 uint32 = 2
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 128
}
func genFn1(p0 int64, p1 int8, p2 int32) string {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint8 = 0
        for fi8 := 0; fi8 < 12; fi8++ {
            p2 += p2
            _ = fi8
        }
        var v9 float32 = 3.14159
        _ = fi6
        _ = v7
        _ = v9
    }
    var v10 bool = true
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 float32 = (((((1e10 + -2.25) * 1.5) - 3.14159) / 0.1) * 3.14159)
    var v11 uint16 = 0
    v12 := v1[("foo bar" + "a")]
    v3 = v3
    for fi13 := 0; fi13 < 20; fi13++ {
        delete(v2, "a")
        switch "\n\t" {
        case "a":
            var v14 uint16 = 96
            _ = v14
        }
        _ = fi13
    }
    var v15 int = 534
    v15 = 128
    v16 := v2[((("hello" + ("a" + "hello")) + ("foo bar" + "foo bar")) + "foo bar")]
    delete(v1, (("hello" + "foo bar") + ("hello" + "foo bar")))
    fmt.Println(v16)
    for fi17 := 0; fi17 < 1; fi17++ {
        if (!true) {
            v16 = (fi17 / ((-9223372036854775808) | 1))
        }
        _ = fi17
    }
    if true {
        v18 := GenS1{F0: -9223372036854775808}
        var v19 int64 = -1
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        v3 = v3
        _ = fi20
    }
    fmt.Println(v16, v12)
    var v21 bool = (true || v12)
    if v21 {
        for fi22 := 0; fi22 < 4; fi22++ {
            for fi23 := 0; fi23 < 4; fi23++ {
                var v24 string = "\n\t"
                _ = fi23
                _ = v24
            }
            if (genFn0(410, 1) != 145) {
                v25 := GenS1{F0: -614}
                fmt.Println(v11, v11, fi22)
                for fi26 := 0; fi26 < 13; fi26++ {
                    switch fi22 {
                    case -433:
                        var v27 uint8 = uint8(v11)
                        _ = v27
                    default:
                        v28 := []int{}
                        _ = v28
                    }
                    v29 := GenS1{F0: 9223372036854775807}
                    v16 += (-(v15))
                    _ = fi26
                    _ = v29
                }
                _ = v25
            }
            for fi30 := 0; fi30 < 15; fi30++ {
                v3 += v3
                var v31 int16 = 314
                v32 := GenS2{F0: 808}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi22
        }
        v33 := GenS2{F0: 453}
        var v34 int64 = -78
        _ = v33
        _ = v34
    }
    var v35 string = "foo bar"
    if ((!v21) && false) {
        fmt.Println(v3)
    }
    for fi36 := 0; fi36 < 20; fi36++ {
        v37 := []int{32767, v15}
        _ = fi36
        _ = v37
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v35
}
