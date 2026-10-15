package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 uint; F1 uint16 }
func genFn0(p0 int16) int16 {
    p0 = -1
    p0 = p0
    switch ("hello" + "a") {
    case "a":
        p0 -= 150
    case "\n\t":
        var v4 bool = true
        _ = v4
    case "":
        v5 := []int{128}
        _ = v5
    default:
        p0 = (p0 ^ (p0 * (p0 | p0)))
    }
    _ = p0
    return (p0 & -549)
}

func main() {
    var v1 uint16 = 439
    v2 := []uint32{42, 42}
    var v3 int8 = 42
    fmt.Println(v3, v3)
    var v6 int = 32767
    if len(v2) > 0 {
        v2[0] = 977
    }
    switch "hello" {
    case "foo bar":
        if ((180 != 1) || (true || (true && ((!(!true)) && (!((!false) || false)))))) {
            var v7 int64 = -447
            _ = v7
        } else {
            var v8 uint32 = 1
            fmt.Println(v8, v1)
            fmt.Println(v8, v6)
            _ = v8
        }
    case "hello":
        fmt.Println(v3, v6)
    case "a":
        var v9 float32 = (0.0 * (0.1 / -2.25))
        _ = v9
    }
    if (!false) {
        if true {
            switch (v6 % ((v6) | 1)) {
            case -821:
                fmt.Println(v6, v6, v3)
            case 1:
                v1 -= v1
            default:
                if false {
                    var v10 bool = (-9223372036854775808 >= 0)
                    var v11 uint32 = 2147483647
                    v3 = (v3 - -45)
                    _ = v10
                    _ = v11
                }
            }
            if false {
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            }
            _ = len(v2)
        } else {
            v6 = (v6 + v6)
        }
        var v12 string = ("a" + "a")
        _ = v12
    } else {
        v2 = append(v2, 2)
        var v13 int = v6
        if true {
            if (726 == 2147483647) {
                if (!true) {
                    _ = len(v2)
                }
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := GenS1{F0: (v1 / ((v1) | 1)), F1: -1, F2: fi14}
                    _ = fi14
                    _ = v15
                }
                for fi16 := 0; fi16 < 18; fi16++ {
                    v17 := GenS1{F0: 571, F1: int16(v13), F2: (-(v6))}
                    _ = fi16
                    _ = v17
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v18 int32 = -358
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            var v19 uint64 = 18446744073709551615
            _ = v18
            _ = v19
        }
        _ = v13
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        if false {
            for fi21 := 0; fi21 < 5; fi21++ {
                var v22 uint8 = 0
                var v23 float32 = float32(0.1)
                _ = fi21
                _ = v22
                _ = v23
            }
            var v24 int = v6
            fmt.Println(v1, v6)
            _ = v24
        }
        v3 -= v3
        _ = fi20
    }
    v6 += (v6 + (v6 / (((^(v6))) | 1)))
    _ = len(v2)
    switch "hello" {
    case "a":
        v2 = append(v2, 65535)
    case "hello":
        var v25 uint64 = 847
        _ = v25
    }
    v6 = 1
    switch 127 {
    case -1:
        switch "" {
        case "hello":
            fmt.Println(v1)
        }
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS1{F0: (v1 / ((65535) | 1)), F1: 1, F2: (^(v6))}
        var v28 int16 = 1
        _ = fi26
        _ = v27
        _ = v28
    }
    _ = len(v2)
    if false {
        v29 := GenS2{F0: 925, F1: v1}
        _ = len(v2)
        v1 = (v1 & (v1 ^ 32767))
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
}
