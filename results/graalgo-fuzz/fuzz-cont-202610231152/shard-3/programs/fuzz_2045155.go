package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
func genFn0(p0 bool, p1 uint32, p2 int) uint8 {
    if (0.0 <= ((float64((-2.25 + 3.14159)) * 0.1) * 0.1)) {
        fmt.Println(p1)
        v5 := []int{(p2 | (p2 ^ (p2 * p2))), (p2 ^ 47), p2, 127}
        v6 := GenS1{F0: p0, F1: 130}
        _ = v5
        _ = v6
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            p0 = (290 < 0)
            _ = fi7
        }
        p2 += -189
    }
    switch (("" + "") + "") {
    case "hello":
        var v8 uint64 = 42
        _ = v8
    case "a":
        p0 = (677 == -1)
    }
    fmt.Println(p0)
    var v9 uint32 = (p1 / ((2147483647) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := []string{("hello" + "a")}
    v2 := []int{}
    var v3 uint8 = 128
    var v4 int = 560
    if len(v2) > 0 {
        v2[0] = v4
    }
    var v10 float32 = 0.0
    v11 := GenS1{F0: (992 > 2), F1: v3}
    var v12 string = ("" + "")
    v13 := GenS1{F0: (!true), F1: (v3 << ((138) & 63))}
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: 0, F1: 255}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int64 = -898
        if false {
            _ = len(v2)
        }
        for fi18 := 0; fi18 < 5; fi18++ {
            v3 = 34
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    switch "\n\t" {
    case "\n\t":
        var v19 uint = 4294967295
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int = (fi20 << ((773) & 63))
            _ = fi20
            _ = v21
        }
    }
    if (986 > 42) {
        fmt.Println(v12, v4, v12)
    }
    v22 := GenS3{F0: (!true), F1: -257, F2: v4}
    var v23 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
