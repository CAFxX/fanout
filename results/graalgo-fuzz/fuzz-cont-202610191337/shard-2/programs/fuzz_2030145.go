package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 uint8, p1 int8, p2 int) uint32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{356}
        for fi8 := 0; fi8 < 10; fi8++ {
            v9 := GenS1{F0: 18446744073709551615, F1: ((^(p1)) >= 0)}
            var v10 uint = 315
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v11 string = "foo bar"
    p2 = (p2 + (-(p2)))
    v12 := []int{(p2 | p2), p2, (p2 % ((32767) | 1)), 308}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 2147483647
}
func genFn1(p0 uint16, p1 int64, p2 int16) uint {
    p1 = (p1 | p1)
    switch "hello" {
    case "hello":
        p0 = p0
    case "foo bar":
        p2 = p2
    case "":
        if ("foo bar" < "") {
            fmt.Println(p0)
            for fi13 := 0; fi13 < 17; fi13++ {
                for fi14 := 0; fi14 < 10; fi14++ {
                    fmt.Println(p1)
                    if (true || true) {
                        fmt.Println(p2, fi14, p0)
                        p0 = (p0 | p0)
                        var v15 uint = 42
                        _ = v15
                    }
                    v16 := []int{1, 4294967295}
                    _ = fi14
                    _ = v16
                }
                _ = fi13
            }
        }
    }
    var v17 float64 = ((-2.25 - 1.5) + -2.25)
    fmt.Println(p1, v17)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 1
}

func main() {
    var v1 int8 = 100
    var v2 int16 = 0
    v3 := []string{"\n\t", ""}
    v4 := []int{}
    var v5 int32 = 1
    if len(v4) > 0 {
        v4[0] = 32767
    }
    v2 = -52
    fmt.Println(v5, v1)
    var v18 int8 = v1
    v3 = append(v3, "foo bar")
    v5 = v5
    fmt.Println(v2, v18)
    if false {
        v1 = v18
        v19 := GenS1{F0: genFn1(2, -9223372036854775808, 42), F1: false}
        _ = v19
    } else {
        v1 += 127
        var v20 int64 = 0
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 bool = false
                if v23 {
                    var v24 int64 = -645
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            var v25 float64 = (-2.25 + 1.5)
            _ = fi21
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
