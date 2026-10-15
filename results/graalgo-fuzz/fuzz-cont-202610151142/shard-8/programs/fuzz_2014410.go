package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
func genFn0(p0 uint32, p1 int8) uint8 {
    p0 += p0
    p1 = p1
    v6 := []int{668, -9223372036854775808, -189, 0}
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1(p0 uint, p1 bool, p2 uint) int16 {
    fmt.Println(p2)
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := GenS1{F0: ("foo bar" + ("" + "a")), F1: 4294967295}
        _ = fi7
        _ = v8
    }
    v9 := []int{128, 9223372036854775807}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(p0, p1, p0)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -1
    var v3 float32 = float32(1.5)
    v4 := make(map[uint32]uint32)
    var v5 bool = true
    if v5 {
        switch (("a" + "\n\t") + "a") {
        case "":
            var v11 int8 = -92
            _ = v11
        }
        v2 -= 0
        delete(v4, 42)
    } else {
        v12 := GenS1{F0: "hello", F1: 1}
        _ = v12
    }
    v13 := GenS1{F0: "\n\t", F1: 65535}
    v1[32767] = "foo bar"
    fmt.Println(v3, v3, v2)
    v2 = (-(v2))
    fmt.Println(v2)
    v14 := []int{263, 4294967295, 9223372036854775807}
    if (18446744073709551615 < 4294967295) {
        var v15 int64 = v2
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 float64 = (0.0 - 3.14159)
        for fi18 := 0; fi18 < 7; fi18++ {
            v4[21] = 800
            if v5 {
                switch 431 {
                case 42:
                    _ = len(v4)
                }
            }
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    v19 := GenS1{F0: ("" + "foo bar"), F1: 2147483647}
    var v20 float32 = v3
    fmt.Println(v3)
    var v21 int16 = -54
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
