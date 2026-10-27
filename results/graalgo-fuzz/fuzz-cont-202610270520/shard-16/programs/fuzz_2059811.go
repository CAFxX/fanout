package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 string }
func genFn0(p0 int16, p1 int, p2 uint) int8 {
    p1 -= (p1 - (-(p1)))
    var v7 int32 = 1
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := []int{9223372036854775807, 0}
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint16 = 65535
            v7 -= (v7 & v7)
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    var v12 float64 = -2.25
    for fi13 := 0; fi13 < 13; fi13++ {
        p1 = (p1 + 42)
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 float32 = (-2.25 * float32(((v12 / (((v12 * (-2.25 / float64(1.5))) - v12) / (-2.25 * (v12 + float64(3.14159))))) * 3.14159)))
            var v16 uint32 = 2
            v17 := GenS1{F0: 2, F1: true, F2: (!("\n\t" < ((("" + "a") + ("hello" + "hello")) + ("foo bar" + "\n\t"))))}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        v18 := GenS1{F0: 710, F1: false, F2: true}
        _ = fi13
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return -128
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{0, 361, 4294967295, 4294967295}
    v3 := []int{128}
    v4 := make(map[uint32]bool)
    v5 := make(map[int]string)
    var v6 bool = true
    v19 := GenS1{F0: 0, F1: (true && v6), F2: (1.5 <= 1e10)}
    if (2 > 161) {
        var v20 float32 = 1e10
        _ = v20
    }
    var v21 int16 = -32768
    switch "foo bar" {
    case "\n\t":
        v22 := GenS1{F0: 250, F1: (v6 || (float32(1.5) != 1.5)), F2: (830 >= 65535)}
        _ = v22
    case "a":
        var v23 uint64 = 42
        _ = v23
    case "hello":
        for fi24 := 0; fi24 < 3; fi24++ {
            switch "hello" {
            case "\n\t":
                if (966 < 750) {
                    var v25 float64 = (3.14159 / 1.0)
                    v6 = (("hello" + ("" + ("hello" + "a"))) < "foo bar")
                    _ = v25
                }
            case "foo bar":
                v26 := v5[fi24]
                _ = v26
            case "a":
                fmt.Println(v6, v6, v6)
            default:
                v27 := GenS1{F0: 0, F1: (!true), F2: v6}
                _ = v27
            }
            if (2 != 1) {
                v21 -= (v21 % ((-32768) | 1))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                var v28 int8 = 42
                for fi29 := 0; fi29 < 1; fi29++ {
                    v4[4294967295] = (-45 == genFn0(v21, (fi24 << ((1) & 63)), 607))
                    v6 = v6
                    _ = fi29
                }
                _ = v28
            }
            if (127 < 0) {
                if true {
                    fmt.Println(v21, v6)
                    switch "hello" {
                    case "\n\t":
                        var v30 string = ("hello" + "a")
                        _ = v30
                    case "foo bar":
                        v5[fi24] = "a"
                    case "a":
                        fmt.Println(v6)
                    }
                }
                v31 := GenS2{F0: 1, F1: "a"}
                _ = v31
            }
            _ = fi24
        }
    }
    if true {
        v6 = (!true)
        v21 -= (v21 | (v21 ^ (v21 & v21)))
        v21 -= -32768
    }
    v2 = append(v2, 42)
    for fi32 := 0; fi32 < 5; fi32++ {
        for fi33 := 0; fi33 < 12; fi33++ {
            fmt.Println(v6, fi32)
            v6 = false
            _ = fi33
        }
        v34 := GenS1{F0: 481, F1: false, F2: (3.14159 > ((float64(1e10) * -2.25) - ((1e10 + 0.1) + 3.14159)))}
        _ = fi32
        _ = v34
    }
    if ((v6 || false) || false) {
        var v35 int = -2147483648
        for fi36 := 0; fi36 < 12; fi36++ {
            v21 -= v21
            _ = fi36
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := GenS1{F0: 1, F1: ((v6 && (18446744073709551615 < 2)) && v6), F2: (42 > 1)}
            var v39 uint = 903
            _ = fi37
            _ = v38
            _ = v39
        }
        _ = v35
    }
    var v40 uint = 4294967295
    fmt.Println(v40)
    v6 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
