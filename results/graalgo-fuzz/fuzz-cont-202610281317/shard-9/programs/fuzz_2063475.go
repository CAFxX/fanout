package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
func genFn0(p0 bool, p1 int16, p2 int16) int16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: fi5, F1: 65535}
        if p0 {
            p2 = p2
            var v7 int = fi5
            _ = v7
        } else {
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := []int{(fi8 * 42), fi8, (fi8 - fi5)}
                _ = fi8
                _ = v9
            }
            p0 = (!("hello" == "a"))
        }
        _ = fi5
        _ = v6
    }
    v10 := GenS1{F0: -1, F1: 0}
    switch (("foo bar" + "foo bar") + ("" + "a")) {
    case "\n\t":
        switch "hello" {
        case "hello":
            p2 = (p1 - (p2 / ((p2) | 1)))
        }
    case "hello":
        v11 := GenS1{F0: -9223372036854775808, F1: 5}
        _ = v11
    }
    var v12 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    return p1
}
func genFn1(p0 uint, p1 int) string {
    v13 := []int{(^(p1)), 1, p1, -9223372036854775808}
    var v14 uint32 = 576
    _ = len(v13)
    if false {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        var v15 bool = (float32(1e10) <= 3.14159)
        var v16 float32 = 0.1
        _ = v15
        _ = v16
    } else {
        p0 += (p0 + p0)
        var v17 uint16 = 1
        var v18 uint8 = 127
        _ = v17
        _ = v18
    }
    if false {
        if (1 < -126) {
            v13 = append(v13, (p1 << (((p0 * (p0 | 992))) & 63)))
        }
        var v19 int32 = 2147483647
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return (("a" + "hello") + "foo bar")
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 65
    var v3 bool = false
    var v4 int = 32767
    v1[""] = v4
    for fi20 := 0; fi20 < 4; fi20++ {
        fmt.Println(fi20)
        v1[("hello" + ("a" + "foo bar"))] = v4
        _ = fi20
    }
    v2 -= (v2 | 0)
    var v21 int8 = (v2 / ((-56) | 1))
    var v22 int32 = 716
    if (23 != 255) {
        v1[""] = 48
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        if (("foo bar" <= "hello") && v3) {
            switch (v4 | (fi23 << ((4294967295) & 63))) {
            case -9223372036854775808:
                v24 := GenS1{F0: 0, F1: 265}
                _ = v24
            case 9223372036854775807:
                v3 = (-1 > 506)
            case 4294967295:
                v25 := GenS1{F0: 128, F1: 97}
                _ = v25
            default:
                if ((v21 > v21) || true) {
                    delete(v1, "a")
                }
            }
        }
        v21 = (^(v2))
        _ = fi23
    }
    if (((v2 <= (^(v21))) || v3) || v3) {
        v26 := v1[("\n\t" + "")]
        fmt.Println(v4, v3, v2)
        switch "hello" {
        case "":
            v22 = v22
        }
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
