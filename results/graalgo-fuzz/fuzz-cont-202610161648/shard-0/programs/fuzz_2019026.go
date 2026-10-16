package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16, p2 uint16) int64 {
    var v6 uint16 = (p2 | (p1 >> ((255) & 63)))
    switch 1 {
    case -2147483648:
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := []int{(fi7 + fi7)}
            v9 := GenS1{F0: 222, F1: 108}
            _ = fi7
            _ = v8
            _ = v9
        }
    case 516:
        if (432 == -435) {
            v6 -= (v6 & (^(p1)))
        }
    case 42:
        var v10 uint64 = 490
        _ = v10
    }
    fmt.Println(p1, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 bool, p1 int64) string {
    v11 := GenS2{F0: 85}
    var v12 uint8 = 48
    fmt.Println(v12, v12)
    fmt.Println(p0, v12)
    var v13 uint16 = 2
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    v1 := []string{"a", ""}
    var v2 int = 127
    var v3 int16 = 1000
    var v4 uint8 = 128
    v5 := make(map[string]int)
    fmt.Println(v2, v2)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint64 = 316
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = fi14
        _ = v15
    }
    var v16 int16 = -32768
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v17 := GenS1{F0: 42, F1: v4}
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS2{F0: 42}
        var v20 uint16 = 255
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: (v4 - v4), F1: v4}
    var v22 uint16 = 1
    v3 = int16(v2)
    switch 1 {
    case -2147483648:
        v1 = append(v1, ("\n\t" + (("a" + "\n\t") + "a")))
    case 4294967295:
        delete(v5, ("hello" + "\n\t"))
    default:
        v23 := GenS3{F0: 264, F1: 2}
        _ = v23
    }
    fmt.Println(v16)
    v24 := GenS3{F0: 2147483648, F1: 1}
    v3 = (v16 >> ((uint(v22)) & 63))
    v5[""] = int(v22)
    if true {
        v25 := GenS4{F0: 226}
        for fi26 := 0; fi26 < 16; fi26++ {
            v16 = 0
            var v27 uint8 = 31
            _ = fi26
            _ = v27
        }
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
