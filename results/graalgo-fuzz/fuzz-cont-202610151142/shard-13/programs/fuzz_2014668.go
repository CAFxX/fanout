package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 uint8, p1 int8) bool {
    var v7 uint16 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 = 461
            v7 = 839
            _ = fi9
        }
        p0 += p0
        _ = fi8
    }
    v10 := GenS1{F0: ((p1 ^ (-(p1))) == p1), F1: "hello"}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return (-491 > 32767)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]uint32)
    var v3 uint16 = 65535
    v4 := make(map[int]uint32)
    var v5 int = 2147483647
    v6 := make(map[string]string)
    switch (v5 % ((v5) | 1)) {
    case 127:
        fmt.Println(v3, v3)
    case -239:
        delete(v2, "foo bar")
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 127
        v13 := []int{9223372036854775807, 0, (fi11 | (v5 >> ((162) & 63))), 32767}
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v3, v3)
    v14 := GenS1{F0: false, F1: ("foo bar" + "")}
    v3 = v3
    var v15 uint32 = 19
    delete(v4, (^(v5)))
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
