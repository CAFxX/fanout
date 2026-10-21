package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0() string {
    v5 := []int{32767, 127, 0, 32768}
    var v6 float64 = 0.1
    v6 = 3.14159
    var v7 float64 = ((1.5 + v6) * ((float64(1.5) / ((v6 * v6) / v6)) / 1.5))
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    _ = v5
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint, p1 int8) int32 {
    for fi8 := 0; fi8 < 10; fi8++ {
        p1 += p1
        _ = fi8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int = fi9
        p1 = p1
        _ = fi9
        _ = v10
    }
    p0 = (p0 | (p0 + 2))
    p1 = p1
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 int64 = 302
    var v2 string = ("a" + "")
    v3 := make(map[string]uint32)
    var v4 string = "\n\t"
    var v11 string = v4
    for fi12 := 0; fi12 < 7; fi12++ {
        v3["hello"] = 129
        _ = fi12
    }
    fmt.Println(v4, v4)
    if false {
        fmt.Println(v2)
        if (((float32(0.0) <= float32((float64(0.0) + ((1e10 - 1.5) / 1.0)))) && (!((float64(-2.25) * (0.1 + float64(1e10))) < 0.0))) || true) {
            fmt.Println(v2)
            v4 = v11
            var v13 bool = true
            _ = v13
        }
        var v14 uint8 = 128
        _ = v14
    }
    v15 := GenS1{F0: 128}
    fmt.Println(v11)
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(v2)
        _ = fi16
    }
    switch 602 {
    case 670:
        switch "a" {
        case "hello":
            var v17 uint16 = 574
            _ = v17
        case "a":
            switch 1 {
            case 4294967295:
                v18 := []int{}
                _ = v18
            case 32767:
                v2 = ("foo bar" + v2)
            case 558:
                v19 := []int{}
                _ = v19
            default:
                fmt.Println(v11, v11, v1)
            }
        }
    }
    if false {
        switch 127 {
        case -9223372036854775808:
            v4 = v4
        case 0:
            v20 := GenS1{F0: 0}
            _ = v20
        case -2147483648:
            var v21 float32 = 3.14159
            _ = v21
        default:
            v2 = ""
        }
    }
    v2 = v11
    fmt.Println(v11)
    for fi22 := 0; fi22 < 16; fi22++ {
        v4 = v11
        switch "\n\t" {
        case "foo bar":
            var v23 float64 = (1e10 + ((-2.25 / 3.14159) - (0.1 + (-2.25 + 1.5))))
            _ = v23
        case "\n\t":
            fmt.Println(v11, v1, v11)
        }
        _ = fi22
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
