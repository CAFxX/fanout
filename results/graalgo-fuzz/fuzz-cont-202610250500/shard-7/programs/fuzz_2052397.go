package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
func genFn0(p0 string, p1 uint16, p2 int16) uint {
    if true {
        var v7 uint64 = 514
        var v8 int16 = (p2 + (p2 - p2))
        v9 := GenS1{F0: (p0 + p0), F1: 2147483647}
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 float32 = 1e10
        var v11 uint32 = 4294967295
        _ = v10
        _ = v11
    }
    p2 -= (p2 << ((2) & 63))
    if ((((!(!(3.14159 > (1.5 / -2.25)))) && true) && true) && (true && false)) {
        var v12 int64 = 1
        v13 := []int{-2147483648, -1, 32767, -498}
        _ = v12
        _ = v13
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 int16, p1 uint32, p2 uint8) bool {
    var v14 uint8 = (^(p2))
    v15 := []int{4294967295, -182, 127, 128}
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return (true && (false && false))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = -32768
    v3 := []uint32{}
    var v4 int32 = -573
    var v5 int = 994
    v6 := make(map[int]int)
    v4 = 2147483647
    var v16 int = v5
    var v17 bool = false
    _ = len(v3)
    switch ("" + "") {
    case "a":
        for fi18 := 0; fi18 < 14; fi18++ {
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(fi18)
                _ = fi19
            }
            var v20 int64 = 9223372036854775807
            var v21 uint32 = 472
            _ = fi18
            _ = v20
            _ = v21
        }
    case "hello":
        var v22 int16 = (v2 | (v2 ^ (v2 - v2)))
        _ = v22
    default:
        var v23 float64 = (1e10 - 0.0)
        _ = v23
    }
    v24 := GenS1{F0: "", F1: 65535}
    switch -2147483648 {
    case 9223372036854775807:
        v2 = 32767
    }
    fmt.Println(v4)
    delete(v1, "foo bar")
    fmt.Println(v2, v16, v4)
    v25 := GenS1{F0: (("hello" + "hello") + "a"), F1: 217}
    if false {
        v17 = (0.0 >= -2.25)
        v3 = append(v3, 42)
    }
    switch (("a" + ((("foo bar" + (("foo bar" + "hello") + ("foo bar" + ""))) + "") + "")) + "hello") {
    case "\n\t":
        var v26 int = (v5 / ((v5) | 1))
        _ = v26
    default:
        if v17 {
            v27 := GenS2{F0: 2, F1: (65535 != 90), F2: (v5 + v16)}
            _ = v27
        }
    }
    var v28 uint = 1
    for fi29 := 0; fi29 < 18; fi29++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v2, v28, v17)
        v16 = v5
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
