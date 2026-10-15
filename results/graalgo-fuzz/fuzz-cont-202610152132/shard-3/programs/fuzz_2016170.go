package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint64, p1 uint32, p2 uint32) int32 {
    if (true || false) {
        var v7 uint16 = 643
        _ = v7
    }
    if (-8 <= int8(p0)) {
        if true {
            var v8 string = "a"
            var v9 uint64 = p0
            switch "\n\t" {
            case "hello":
                p0 += v9
            case "foo bar":
                fmt.Println(p0, p2, p0)
            case "a":
                fmt.Println(p1)
            default:
                v10 := []int{}
                _ = v10
            }
            _ = v8
            _ = v9
        }
        p1 += p1
        var v11 string = ("foo bar" + "hello")
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    v1 := []int{}
    v2 := []uint32{}
    var v3 int8 = -95
    var v4 uint8 = 2
    v5 := []string{"hello"}
    v6 := make(map[string]uint32)
    switch "foo bar" {
    case "\n\t":
        v12 := GenS1{F0: "foo bar"}
        _ = v12
    case "a":
        if len(v1) > 0 {
            v1[0] = -2147483648
        }
    case "hello":
        v4 -= 127
    }
    var v13 float64 = -2.25
    v4 += 2
    var v14 uint64 = 18446744073709551615
    if (((^(v4)) < v4) && ((!(-9223372036854775808 > -9223372036854775808)) && false)) {
        var v15 int = -491
        _ = v15
    } else {
        fmt.Println(v13, v3, v3)
        v13 -= 0.1
    }
    _ = len(v1)
    delete(v6, "hello")
    var v16 int32 = 0
    switch "" {
    case "\n\t":
        var v17 int = 0
        _ = v17
    case "a":
        var v18 uint = 458
        _ = v18
    }
    var v19 uint8 = (v4 * v4)
    for fi20 := 0; fi20 < 8; fi20++ {
        if false {
            v4 = uint8(v14)
            v21 := GenS1{F0: (("hello" + "") + "")}
            _ = v21
        } else {
            if true {
                v6["foo bar"] = 863
                var v22 float64 = (-2.25 / 0.1)
                _ = v22
            }
        }
        v6[("\n\t" + "foo bar")] = 33
        _ = fi20
    }
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
