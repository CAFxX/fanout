package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 bool; F1 int32; F2 string }
func genFn0(p0 uint32, p1 uint8, p2 uint8) string {
    if (false && false) {
        p2 = p2
    } else {
        p0 = p0
    }
    v4 := []int{-113, -9223372036854775808, 42, 684}
    var v5 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "\n\t"
}
func genFn1() bool {
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(v6, fi7)
        var v8 int32 = 181
        v6 -= (v6 + (v6 << ((1) & 63)))
        _ = fi7
        _ = v8
    }
    if (2147483647 == 4294967295) {
        switch ("\n\t" + "") {
        case "hello":
            v9 := GenS1{F0: true, F1: "a"}
            _ = v9
        case "foo bar":
            v10 := []int{}
            _ = v10
        case "\n\t":
            v6 = (^(v6))
        }
        fmt.Println(v6)
        switch "a" {
        case "\n\t":
            for fi11 := 0; fi11 < 3; fi11++ {
                var v12 uint16 = 0
                _ = fi11
                _ = v12
            }
        case "a":
            var v13 uint16 = 42
            _ = v13
        case "":
            v14 := []int{-9223372036854775808}
            _ = v14
        }
    }
    _ = v6
    return (!true)
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    var v2 uint64 = 2
    var v3 uint8 = 255
    v15 := GenS1{F0: ((1 > (v3 & (^(v3)))) && (true && ((float64(0.1) / (3.14159 - (1.5 * 0.0))) != 3.14159))), F1: "\n\t"}
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float32 = 1e10
        _ = fi16
        _ = v17
    }
    var v18 float32 = (((3.14159 / (3.14159 + (0.1 - -2.25))) + 0.0) / 0.1)
    var v19 float64 = 0.1
    v19 -= v19
    v20 := GenS1{F0: false, F1: "foo bar"}
    switch (("a" + ("" + "\n\t")) + (("\n\t" + "hello") + "")) {
    case "foo bar":
        var v21 int64 = 0
        _ = v21
    case "a":
        var v22 uint8 = 42
        _ = v22
    case "":
        var v23 bool = false
        _ = v23
    default:
        v24 := GenS1{F0: (v18 == v18), F1: "\n\t"}
        _ = v24
    }
    var v25 uint64 = 0
    v1 = append(v1, "foo bar")
    if false {
        if (42 <= -9223372036854775808) {
            v19 -= v19
        } else {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: true, F1: ("\n\t" + "\n\t")}
                _ = fi26
                _ = v27
            }
        }
    }
    switch 127 {
    case 42:
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS2{F0: (!true), F1: 0, F2: "hello"}
            _ = fi28
            _ = v29
        }
    default:
        v19 -= (v19 + (v19 + v19))
    }
    v19 += v19
    v2 += v25
    fmt.Println(v19, v19, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
}
