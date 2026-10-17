package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 int16, p2 int8) bool {
    v5 := []int{127, -2147483648, -702, 273}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, v7)
            switch fi8 {
            case -9223372036854775808:
                var v9 uint32 = 46
                _ = v9
            case 499:
                var v10 float64 = ((-2.25 / 0.1) * 0.0)
                _ = v10
            case 0:
                var v11 uint = 18446744073709551615
                _ = v11
            }
            _ = fi8
        }
        fmt.Println(v7, p0, fi6)
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return true
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{}
    var v3 uint16 = 42
    v4 := []string{}
    var v12 float64 = 0.0
    switch "a" {
    case "hello":
        var v13 uint64 = 62
        _ = v13
    }
    var v14 int64 = v1
    var v15 int64 = (v1 * v1)
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 float32 = 0.0
            _ = len(v4)
            _ = fi17
            _ = v18
        }
        v2 = append(v2, ("foo bar" + ("a" + ("hello" + "a"))))
        switch fi16 {
        case 1:
            fmt.Println(v15, v15)
        case -288:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case 2147483647:
            for fi19 := 0; fi19 < 18; fi19++ {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                v20 := GenS1{F0: 65535, F1: v14}
                v1 += 1
                v15 -= (v14 % ((42) | 1))
                var v21 float32 = 1.5
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = fi16
    }
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 int16 = 0
        _ = fi22
        _ = v23
    }
    v24 := GenS1{F0: 2147483648, F1: -624}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
