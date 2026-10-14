package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 string, p1 string, p2 uint32) int {
    p1 = (p0 + p0)
    switch 32768 {
    case 2147483647:
        var v6 bool = true
        _ = v6
    case 32767:
        fmt.Println(p0)
    case 4294967295:
        p2 += p2
    default:
        var v7 string = p0
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            p1 = p0
        } else {
            v9 := []int{(-(fi8)), 0, (fi8 & fi8), (^(fi8))}
            v9 = append(v9, fi8)
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}
func genFn1(p0 bool) int8 {
    var v10 int64 = -563
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(p0, p0, v10)
        var v12 float32 = 3.14159
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    v5 := []int{4294967295, 42, 32767, 2147483647}
    fmt.Println(v1)
    switch (("" + (("" + ("\n\t" + "a")) + "a")) + "\n\t") {
    case "\n\t":
        v1 -= 813
    case "hello":
        v1 -= v1
    case "a":
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 float32 = 1.5
            v15 := GenS1{F0: 18446744073709551615}
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v1 -= (^(v1))
    v1 += (v1 - v1)
    v2[genFn0(("" + "\n\t"), "foo bar", 4294967295)] = false
    var v16 string = "a"
    v1 -= (v1 % ((1) | 1))
    if len(v5) > 0 {
        v5[0] = 192
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
