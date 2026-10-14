package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int32, p1 uint8, p2 int8) int {
    p2 -= (p2 - (p2 / (((-(p2))) | 1)))
    v7 := []int{2147483647}
    p1 += 2
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -154
}
func genFn1(p0 uint16, p1 uint8) int {
    if true {
        switch -996 {
        case -1:
            if ((^(p0)) <= p0) {
                p1 += 42
                p0 -= (p0 + (^(p0)))
            }
        default:
            switch "a" {
            case "":
                switch "a" {
                case "foo bar":
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 float64 = float64(1.5)
                        _ = fi8
                        _ = v9
                    }
                case "":
                    fmt.Println(p0)
                }
            }
        }
        var v10 int8 = 127
        _ = v10
    }
    v11 := GenS1{F0: (2 != 2147483648)}
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 int16 = 1000
    v2 := []string{"foo bar"}
    var v3 uint32 = 63
    var v4 uint32 = v3
    var v5 float64 = ((((0.1 + 1.5) / 1.0) + -2.25) / float64((-2.25 + 1.5)))
    v6 := []string{}
    v5 = v5
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    v5 += v5
    v3 += (v3 | (v4 & v4))
    var v12 int16 = (^(v1))
    v6 = append(v6, "foo bar")
    var v13 int64 = 1
    v4 -= v3
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v3, v13, fi14)
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 string = (("hello" + ("hello" + ("" + ("foo bar" + "a")))) + "a")
            _ = fi15
            _ = v16
        }
        var v17 int32 = 1
        _ = fi14
        _ = v17
    }
    var v18 int = 2147483647
    fmt.Println(v12, v12)
    v12 += (v1 / ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
