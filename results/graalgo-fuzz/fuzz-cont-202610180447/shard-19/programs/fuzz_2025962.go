package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint }
type GenS2 struct { F0 string; F1 uint; F2 uint32 }
func genFn0(p0 int8, p1 int64) string {
    var v7 int16 = 42
    switch "\n\t" {
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := []int{42, (fi8 >> ((442) & 63)), 42, 128}
            _ = fi8
            _ = v9
        }
    case "hello":
        var v10 string = "hello"
        _ = v10
    default:
        v11 := GenS1{F0: 255, F1: 168, F2: 18446744073709551615}
        _ = v11
    }
    if (((-2.25 - (((0.0 + 0.1) + ((0.1 + 0.1) + (3.14159 - 3.14159))) - ((((3.14159 - 3.14159) / 1.0) / -2.25) / -2.25))) > (1e10 / ((1.5 * -2.25) + (-2.25 + (0.1 * -2.25))))) || ((((("\n\t" + ((("hello" + (("" + "hello") + "")) + "") + ("hello" + ("foo bar" + (("hello" + "a") + "hello"))))) + ((("\n\t" + "") + "a") + "")) + "a") + ("hello" + (("" + "a") + "a"))) < "")) {
        var v12 uint32 = 42
        fmt.Println(v12, v12)
        _ = v12
    } else {
        v13 := []int{127, -1}
        _ = v13
    }
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}
func genFn1(p0 uint32) string {
    var v14 uint64 = 42
    if ("foo bar" != "") {
        var v15 uint16 = 42
        v16 := []int{-9223372036854775808, 42}
        var v17 int8 = 0
        _ = v15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 2, F1: 645, F2: 18446744073709551615}
    fmt.Println(p0, v14)
    var v19 string = "foo bar"
    _ = p0
    _ = v14
    _ = v18
    _ = v19
    return ""
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int = -2147483648
    var v4 int8 = -1
    v5 := []string{("a" + ("hello" + ((("" + "foo bar") + "foo bar") + ""))), "hello"}
    v6 := make(map[string]bool)
    v3 = v3
    for fi20 := 0; fi20 < 20; fi20++ {
        if ((872 <= 718) || (-2.25 > 0.1)) {
            for fi21 := 0; fi21 < 2; fi21++ {
                fmt.Println(v3, v4)
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(v4, v4, v4)
                v23 := GenS1{F0: 0, F1: 65535, F2: 1}
                if ((26 >= 438) && (2 > 42)) {
                    var v24 uint8 = 102
                    if true {
                        v4 += v4
                        var v25 float32 = (-2.25 / 3.14159)
                        _ = v25
                    } else {
                        fmt.Println(v3)
                    }
                    _ = v24
                }
                v26 := GenS2{F0: "foo bar", F1: 2, F2: 439}
                _ = fi22
                _ = v23
                _ = v26
            }
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            v4 += -128
            _ = fi27
        }
        fmt.Println(fi20, v4, v3)
        _ = fi20
    }
    v3 = v3
    for fi28 := 0; fi28 < 13; fi28++ {
        v6["a"] = true
        var v29 float64 = (1.5 + 1e10)
        _ = fi28
        _ = v29
    }
    v3 -= v3
    v3 -= (v3 ^ v3)
    _ = len(v2)
    for fi30 := 0; fi30 < 15; fi30++ {
        v3 = (v3 >> ((255) & 63))
        v4 -= (v4 * (v4 << ((4294967295) & 63)))
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 bool = (42 > 65535)
            v33 := GenS1{F0: 1, F1: 1, F2: 497}
            v32 = v32
            delete(v2, "\n\t")
            _ = fi31
            _ = v32
            _ = v33
        }
        _ = fi30
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
