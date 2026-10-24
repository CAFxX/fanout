package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
func genFn0(p0 int8, p1 int32, p2 uint8) uint64 {
    if (true || true) {
        if (3.14159 == 3.14159) {
            v6 := GenS1{F0: (("foo bar" + "") + "hello"), F1: (!false)}
            v7 := GenS1{F0: "a", F1: false}
            p1 = p1
            _ = v6
            _ = v7
        } else {
            p0 -= (p0 & (p0 & (-(p0))))
        }
        for fi8 := 0; fi8 < 14; fi8++ {
            p2 -= (p2 | (p2 % ((42) | 1)))
            v9 := []int{32767, fi8, (fi8 - (fi8 / ((2147483647) | 1))), -2147483648}
            if len(v9) > 0 {
                v9[0] = (-(fi8))
            }
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    switch 476 {
    case 0:
        var v10 uint = 2
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 693
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[uint32]int)
    var v3 int = 128
    v4 := []string{"a", "a", "foo bar"}
    var v5 uint8 = 0
    for fi11 := 0; fi11 < 6; fi11++ {
        if (false || (false || false)) {
            if (!(false || (18446744073709551615 >= genFn0(-128, -1, (v5 | (v1 % ((v1) | 1))))))) {
                switch fi11 {
                case 42:
                    var v12 bool = (18446744073709551615 < 1)
                    _ = v12
                default:
                    fmt.Println(v1, fi11, v1)
                }
            }
        } else {
            var v13 int8 = 42
            _ = v13
        }
        v4 = append(v4, "")
        if false {
            v5 = v5
            if false {
                for fi14 := 0; fi14 < 5; fi14++ {
                    fmt.Println(fi14, fi14)
                    var v15 uint8 = (v5 % ((128) | 1))
                    _ = fi14
                    _ = v15
                }
            }
        } else {
            fmt.Println(v5, fi11)
        }
        _ = fi11
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        var v16 bool = (!false)
        _ = v16
    default:
        switch "\n\t" {
        case "foo bar":
            var v17 int = 1
            _ = v17
        case "a":
            var v18 int8 = 100
            _ = v18
        }
    }
    var v19 int = (^(v3))
    v20 := GenS1{F0: (("" + "\n\t") + "foo bar"), F1: false}
    v1 = 112
    if len(v4) > 0 {
        v4[0] = "\n\t"
    }
    v1 = 108
    if (-1 == 42) {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 int = v19
            _ = fi21
            _ = v22
        }
    }
    fmt.Println(v19, v19, v1)
    v19 = 771
    for fi23 := 0; fi23 < 13; fi23++ {
        for fi24 := 0; fi24 < 12; fi24++ {
            v5 -= v1
            if true {
                v2[uint32(v5)] = v3
                v5 = 99
            } else {
                v19 = (^(v19))
            }
            _ = fi24
        }
        fmt.Println(fi23)
        var v25 string = "hello"
        _ = fi23
        _ = v25
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        _ = len(v2)
        fmt.Println(fi26, v19)
        v3 = (^(v19))
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
