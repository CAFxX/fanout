package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0() string {
    var v6 bool = (false && (255 > 83))
    v7 := GenS1{F0: 255, F1: 118}
    _ = v6
    _ = v7
    return (("foo bar" + ("\n\t" + "")) + "hello")
}
func genFn1(p0 uint32, p1 uint8, p2 uint) uint8 {
    for fi8 := 0; fi8 < 9; fi8++ {
        p1 -= (p1 - uint8(p2))
        var v9 float64 = ((1e10 / 0.1) - 1.5)
        v10 := GenS1{F0: (^(p1)), F1: 78}
        _ = fi8
        _ = v10
        _ = v9
    }
    if (42 >= 0) {
        var v11 int64 = -662
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return (p1 * uint8(p0))
}

func main() {
    var v1 uint16 = 707
    var v2 float32 = (-2.25 * 1.5)
    var v3 float64 = 0.0
    var v4 int16 = 0
    v5 := []string{((("hello" + "foo bar") + ("a" + (("" + "\n\t") + "foo bar"))) + ("" + "a")), "a", "", "a"}
    var v12 bool = true
    fmt.Println(v4)
    var v13 int32 = 0
    fmt.Println(v2, v2)
    for fi14 := 0; fi14 < 2; fi14++ {
        if false {
            var v15 string = ("foo bar" + ("hello" + "a"))
            _ = v15
        }
        var v16 int = fi14
        switch "hello" {
        case "\n\t":
            _ = len(v5)
        case "a":
            v3 = (v3 / 1.5)
        case "hello":
            v13 += v13
        }
        _ = fi14
        _ = v16
    }
    if ((^(v4)) != v4) {
        for fi17 := 0; fi17 < 1; fi17++ {
            var v18 int16 = (v4 / (((v4 >> ((42) & 63))) | 1))
            switch 42 {
            case -682:
                var v19 uint = 42
                _ = v19
            case -468:
                var v20 uint16 = (v1 << ((2) & 63))
                _ = v20
            }
            fmt.Println(v13, fi17, fi17)
            _ = fi17
            _ = v18
        }
    } else {
        if (v3 < (0.1 + (v3 - 1.5))) {
            var v21 int16 = -1
            _ = v21
        } else {
            v4 -= 322
            for fi22 := 0; fi22 < 15; fi22++ {
                if len(v5) > 0 {
                    v5[0] = "\n\t"
                }
                _ = fi22
            }
        }
    }
    v23 := GenS1{F0: 1, F1: -17}
    fmt.Println(v12, v4)
    fmt.Println(v12, v2, v4)
    switch -9223372036854775808 {
    case -2147483648:
        v24 := GenS1{F0: 63, F1: 1}
        _ = v24
    case -428:
        v4 -= (-(v4))
    }
    v4 = v4
    v5 = append(v5, "")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
