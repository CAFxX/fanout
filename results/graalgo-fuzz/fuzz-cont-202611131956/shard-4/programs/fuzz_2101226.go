package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
func genFn0(p0 string, p1 bool) uint64 {
    for fi4 := 0; fi4 < 4; fi4++ {
        v5 := GenS1{F0: 2, F1: 866}
        p0 = p0
        p0 = p0
        _ = fi4
        _ = v5
    }
    var v6 float32 = 3.14159
    var v7 uint8 = 2
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int32, p1 uint8, p2 uint16) uint {
    p2 += (p2 >> ((4294967295) & 63))
    p1 = p1
    p0 += p0
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = (p0 ^ int32(fi8))
        p1 = (p1 / ((p1) | 1))
        for fi9 := 0; fi9 < 9; fi9++ {
            p2 = p2
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-2147483648}
    var v3 uint32 = 42
    var v10 uint64 = 2
    if len(v2) > 0 {
        v2[0] = 32767
    }
    fmt.Println(v10)
    var v11 uint64 = v10
    v12 := GenS1{F0: 18446744073709551615, F1: 65535}
    var v13 int64 = -793
    if false {
        if (false || true) {
            fmt.Println(v10)
        } else {
            var v14 string = ""
            if (!(-100 != 42)) {
                v15 := GenS1{F0: 255, F1: 4294967295}
                _ = v15
            }
            v16 := GenS1{F0: 18446744073709551615, F1: 307}
            _ = v14
            _ = v16
        }
    } else {
        var v17 uint8 = 1
        if true {
            var v18 int8 = 100
            fmt.Println(v17)
            v19 := GenS1{F0: 120, F1: 1}
            _ = v18
            _ = v19
        }
        _ = v17
    }
    v3 -= (v3 / ((v3) | 1))
    var v20 int8 = 77
    switch "\n\t" {
    case "foo bar":
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 string = "hello"
            _ = fi21
            _ = v22
        }
    }
    if true {
        v2 = append(v2, 1)
        v3 = 71
    } else {
        fmt.Println(v20)
        var v23 int32 = -2147483648
        _ = v23
    }
    if (!false) {
        var v24 uint16 = 886
        _ = v24
    } else {
        v25 := GenS1{F0: 317, F1: 659}
        _ = v25
    }
    var v26 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v26
    _ = v3
}
