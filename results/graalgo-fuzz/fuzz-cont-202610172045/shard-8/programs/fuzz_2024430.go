package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 uint8 }
func genFn0(p0 bool) int64 {
    p0 = p0
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p0)
    case "":
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: "hello", F1: 42, F2: 30}
            _ = fi6
            _ = v7
        }
    case "\n\t":
        switch "hello" {
        case "hello":
            p0 = (!(p0 && true))
        }
    }
    switch ("foo bar" + (("a" + "a") + "a")) {
    case "a":
        var v8 uint8 = 2
        _ = v8
    case "hello":
        var v9 float64 = 1.5
        _ = v9
    }
    p0 = true
    _ = p0
    return 42
}
func genFn1(p0 int8, p1 int16) int {
    if true {
        p0 += (p0 * 83)
        if (!(float64((3.14159 * 1e10)) > ((1.5 / (0.0 - ((((0.1 * (1e10 / (0.1 / 1.5))) - ((3.14159 * 0.1) - 3.14159)) + 3.14159) / 3.14159))) + (0.1 * 3.14159)))) {
            v10 := GenS1{F0: ("foo bar" + "hello"), F1: 2, F2: 0}
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 float64 = -2.25
                if (((v12 - v12) - v12) <= (v12 * v12)) {
                    for fi13 := 0; fi13 < 12; fi13++ {
                        if false {
                            var v14 uint64 = 1
                            var v15 int16 = p1
                            _ = v14
                            _ = v15
                        }
                        _ = fi13
                    }
                } else {
                    v12 -= (v12 / (1e10 - (v12 - (((v12 + (v12 * ((((float64(3.14159) * v12) / -2.25) / -2.25) / v12))) * ((v12 - 0.1) - 3.14159)) / ((v12 - float64((((((1.5 / 1.5) + 1e10) + ((-2.25 * (0.1 + 3.14159)) / 1.0)) - 0.0) / (1.5 - (-2.25 - 1.5))))) / v12)))))
                }
                _ = fi11
                _ = v12
            }
            switch "a" {
            case "\n\t":
                p1 -= (p1 | (^(p1)))
            case "a":
                p0 = p0
            case "hello":
                v16 := GenS1{F0: "a", F1: 0, F2: 117}
                _ = v16
            }
            _ = v10
        }
    }
    p0 = (p0 % ((p0) | 1))
    var v17 int8 = (^(p0))
    var v18 uint = 1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return -2147483648
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{((("hello" + "foo bar") + "foo bar") + "hello"), "hello"}
    var v3 int8 = -1
    var v4 float32 = ((0.1 + 1e10) / float32(-2.25))
    v5 := make(map[int]bool)
    if true {
        v19 := GenS1{F0: "foo bar", F1: 1, F2: 255}
        var v20 int = 633
        _ = v19
        _ = v20
    }
    var v21 uint16 = 255
    v21 = v21
    switch "a" {
    case "foo bar":
        var v22 uint = 4294967295
        _ = v22
    case "\n\t":
        var v23 uint32 = 549
        _ = v23
    default:
        v4 = (float32(1.5) - (v4 / (v4 * v4)))
    }
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v21, v4)
    case "\n\t":
        v24 := v5[1]
        _ = v24
    }
    var v25 int8 = 1
    fmt.Println(v3, v21)
    for fi26 := 0; fi26 < 9; fi26++ {
        v3 += (v25 + -85)
        for fi27 := 0; fi27 < 13; fi27++ {
            var v28 int8 = (v3 * v3)
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    switch 32768 {
    case -9223372036854775808:
        var v29 string = "a"
        _ = v29
    }
    if false {
        delete(v1, "foo bar")
    }
    v30 := GenS1{F0: ((("hello" + "hello") + ("" + "a")) + ("" + ("foo bar" + "\n\t"))), F1: (v21 << ((255) & 63)), F2: 89}
    fmt.Println(v4)
    v31 := GenS2{F0: v3, F1: -1, F2: 1}
    v32 := GenS1{F0: "\n\t", F1: (v21 / ((32767) | 1)), F2: 127}
    v21 -= 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
