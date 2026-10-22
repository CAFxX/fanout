package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 bool, p1 int8, p2 string) int {
    var v4 int16 = 717
    switch -1 {
    case 1000:
        var v5 int16 = v4
        _ = v5
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 32767
}
func genFn1(p0 uint32, p1 int16, p2 uint16) uint16 {
    v6 := GenS1{F0: -922}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 / ((255) | 1))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    fmt.Println(v3)
    v3 = 1
    switch "" {
    case "foo bar":
        v3 -= v3
    case "hello":
        var v7 float32 = float32(1.5)
        _ = v7
    case "":
        v8 := GenS1{F0: 32768}
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        switch ("a" + "") {
        case "":
            v10 := GenS1{F0: genFn0(false, -1, (("a" + (("foo bar" + "hello") + "\n\t")) + "a"))}
            _ = v10
        case "a":
            v11 := GenS1{F0: fi9}
            _ = v11
        default:
            v1[2147483647] = ("a" + ("" + "hello"))
        }
        _ = fi9
    }
    if true {
        var v12 int16 = 1000
        _ = v12
    }
    var v13 uint64 = uint64(v3)
    v13 = v13
    fmt.Println(v3, v13)
    for fi14 := 0; fi14 < 4; fi14++ {
        v13 += v13
        var v15 bool = true
        _ = fi14
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        var v16 uint8 = (v3 | 19)
        _ = v16
    case "a":
        var v17 uint8 = v3
        _ = v17
    }
    delete(v2, 1)
    var v18 uint64 = v13
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            v21 := GenS1{F0: 42}
            var v22 int64 = -1
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi19
    }
    var v23 int32 = 1
    for fi24 := 0; fi24 < 2; fi24++ {
        fmt.Println(v23, v3, v23)
        _ = fi24
    }
    v25 := GenS1{F0: 0}
    v13 -= (^(v13))
    var v26 int = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
}
