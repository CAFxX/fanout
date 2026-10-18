package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 int16, p1 bool) uint64 {
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 float64 = 0.0
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{}
            _ = len(v11)
            _ = fi10
            _ = v11
        }
        v12 := []int{(fi8 << ((255) & 63)), -88, 127}
        _ = fi8
        _ = v12
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}
func genFn1() int {
    v13 := []int{-356, -1, 42}
    var v14 uint16 = 469
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    var v1 float32 = 0.0
    v2 := []string{"", "hello", "a"}
    var v3 float32 = v1
    var v4 bool = false
    v5 := []uint32{4294967295, 4294967295}
    var v6 float32 = (float32(-2.25) / (1.5 - (((v1 / 1.5) / (v1 * v1)) - 0.0)))
    if len(v5) > 0 {
        v5[0] = 98
    }
    if (1 >= 255) {
        var v15 uint8 = 42
        var v16 int = 42
        _ = v15
        _ = v16
    } else {
        if ((!(32767 <= -32768)) && true) {
            for fi17 := 0; fi17 < 13; fi17++ {
                if true {
                    if len(v2) > 0 {
                        v2[0] = (("foo bar" + ("hello" + "")) + "foo bar")
                    }
                }
                _ = fi17
            }
            fmt.Println(v6)
        }
    }
    var v18 int8 = 120
    fmt.Println(v3, v18)
    if (9223372036854775807 <= 561) {
        switch "hello" {
        case "foo bar":
            v4 = v4
        default:
            v19 := GenS1{F0: 0, F1: "foo bar"}
            _ = v19
        }
    }
    var v20 float32 = ((v3 + (((3.14159 + v1) / v1) - v1)) / v1)
    switch 1 {
    case 155:
        switch genFn1() {
        case 2147483647:
            if true {
                v21 := GenS1{F0: 65535, F1: ((("" + ("a" + ("\n\t" + "a"))) + "hello") + "\n\t")}
                _ = v21
            }
        case 42:
            switch 9223372036854775807 {
            case 127:
                var v22 string = "\n\t"
                _ = v22
            }
        case 51:
            fmt.Println(v6)
        default:
            var v23 uint = 255
            _ = v23
        }
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -2147483648:
        var v24 uint16 = 65535
        _ = v24
    default:
        fmt.Println(v6)
    }
    v6 += (v6 + ((v1 + v3) * (v3 * v3)))
    if v4 {
        fmt.Println(v3)
        v6 = float32(3.14159)
        fmt.Println(v6)
    } else {
        v1 = 1.5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
