package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
func genFn0() uint8 {
    var v7 int64 = 1
    if false {
        var v8 int32 = 692
        _ = v8
    }
    _ = v7
    return 128
}
func genFn1(p0 int32, p1 int32, p2 int16) int32 {
    var v9 uint32 = 2147483647
    var v10 uint8 = 158
    switch "foo bar" {
    case "":
        p2 = 0
    }
    var v11 uint = 2
    var v12 float64 = (((0.0 * 3.14159) * (-2.25 * 1e10)) * -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return (^(p1))
}

func main() {
    v1 := []uint32{2, 1, 2, 42}
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 int64 = 42
    var v5 uint8 = 1
    v6 := []string{}
    var v13 bool = v2
    delete(v3, -2147483648)
    if (100 > 85) {
        delete(v3, -1)
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: -457, F1: fi14}
            var v16 float32 = -2.25
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v4)
    var v17 int8 = -1
    v5 += (v5 % ((v5) | 1))
    switch ("hello" + "foo bar") {
    case "hello":
        v2 = true
    case "a":
        if ((!v13) && true) {
            var v18 int16 = 0
            var v19 uint32 = 4294967295
            switch ("\n\t" + ("" + (((("\n\t" + ((("hello" + "") + ("a" + ("a" + "\n\t"))) + ("hello" + "a"))) + "hello") + "") + (("foo bar" + "foo bar") + "foo bar")))) {
            case "a":
                for fi20 := 0; fi20 < 19; fi20++ {
                    v1 = append(v1, (v19 / ((v19) | 1)))
                    switch fi20 {
                    case 9223372036854775807:
                        v21 := GenS1{F0: fi20, F1: (^(fi20))}
                        _ = v21
                    case 32768:
                        v22 := GenS2{F0: (v5 | genFn0()), F1: (v4 & (v4 * (v4 << ((366) & 63)))), F2: int(v18)}
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 14; fi23++ {
                        v6 = append(v6, (("hello" + "") + "foo bar"))
                        v18 = (v18 + (v18 % ((42) | 1)))
                        fmt.Println(v17, fi23, v19)
                        _ = fi23
                    }
                    _ = fi20
                }
            case "\n\t":
                v24 := GenS2{F0: (v5 >> ((18446744073709551615) & 63)), F1: v4, F2: 225}
                _ = v24
            default:
                _ = len(v3)
            }
            _ = v18
            _ = v19
        } else {
            v25 := GenS1{F0: 32768, F1: int(v4)}
            var v26 int32 = -2147483648
            v2 = (-2.25 <= float64(1e10))
            _ = v25
            _ = v26
        }
    case "\n\t":
        switch "foo bar" {
        case "foo bar":
            var v27 int64 = -243
            _ = v27
        case "":
            v28 := GenS1{F0: 42, F1: -210}
            _ = v28
        }
    }
    v13 = v2
    fmt.Println(v17)
    switch "a" {
    case "hello":
        fmt.Println(v17, v5, v5)
    case "\n\t":
        if v2 {
            if (!false) {
                var v29 int64 = v4
                v2 = (!(!v13))
                _ = v29
            }
            fmt.Println(v17, v13)
            var v30 string = "\n\t"
            _ = v30
        } else {
            v2 = (42 != 519)
        }
    case "a":
        if ((2147483648 != 42) || v13) {
            v13 = (float64(0.0) <= 3.14159)
        }
    }
    fmt.Println(v13, v2, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
