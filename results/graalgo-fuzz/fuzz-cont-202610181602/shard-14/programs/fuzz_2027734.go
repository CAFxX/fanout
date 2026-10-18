package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 int, p1 uint, p2 uint32) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        p1 -= p1
        for fi7 := 0; fi7 < 5; fi7++ {
            p1 -= (p1 ^ (p1 / ((p1) | 1)))
            var v8 float64 = ((0.1 - 1e10) / 1.0)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if false {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 int64 = 9223372036854775807
            p2 += p2
            _ = fi9
            _ = v10
        }
    } else {
        switch "hello" {
        case "":
            var v11 uint16 = 42
            _ = v11
        default:
            v12 := []int{(p0 - (p0 | (p0 | p0))), p0, p0}
            _ = v12
        }
        p0 = -499
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 uint16 = 42
    var v2 int16 = -505
    v3 := make(map[int]bool)
    v4 := make(map[string]int)
    var v5 uint16 = v1
    if true {
        v2 -= (v2 - -32768)
    }
    switch "hello" {
    case "hello":
        v13 := []int{32767, -898, 127, -233}
        _ = v13
    case "\n\t":
        var v14 int = -554
        _ = v14
    default:
        fmt.Println(v2, v2, v2)
    }
    v2 = -353
    fmt.Println(v1, v1)
    v1 = v1
    if false {
        v2 += v2
        var v15 int64 = -9223372036854775808
        _ = v15
    } else {
        v1 = v1
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS1{F0: fi16}
            if (823 >= 2147483647) {
                var v18 int8 = 127
                if false {
                    v19 := []int{(-(fi16)), fi16}
                    var v20 int64 = genFn0((fi16 ^ fi16), 1, 2147483648)
                    _ = v19
                    _ = v20
                }
                _ = v18
            } else {
                _ = len(v3)
                fmt.Println(v5, v1, v1)
            }
            _ = fi16
            _ = v17
        }
        var v21 uint16 = (v5 + v5)
        _ = v21
    }
    fmt.Println(v2, v1)
    switch ("foo bar" + (("\n\t" + "") + ("hello" + (((("" + "foo bar") + "foo bar") + ("hello" + ("" + ((("\n\t" + "hello") + "") + "hello")))) + (("a" + "a") + "a"))))) {
    case "a":
        for fi22 := 0; fi22 < 17; fi22++ {
            v23 := []int{}
            var v24 string = ""
            _ = fi22
            _ = v23
            _ = v24
        }
    case "":
        var v25 uint64 = 2
        _ = v25
    default:
        var v26 float64 = (3.14159 * -2.25)
        _ = v26
    }
    if false {
        v27 := []int{}
        _ = v27
    }
    v1 -= (v5 - (v1 | v1))
    v1 = 255
    _ = len(v4)
    v5 = 520
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
