package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8) uint {
    v6 := GenS1{F0: false}
    if true {
        var v7 string = ("foo bar" + "")
        var v8 uint16 = 65535
        fmt.Println(v7)
        _ = v7
        _ = v8
    }
    v9 := []int{32767, -535, 4294967295}
    _ = p0
    _ = v6
    _ = v9
    return 442
}
func genFn1(p0 uint16, p1 int32, p2 string) int64 {
    v10 := GenS1{F0: false}
    p1 = 0
    v11 := []int{127, 1}
    v11 = append(v11, -2147483648)
    v12 := GenS2{F0: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    var v1 uint = 322
    var v2 int32 = -2147483648
    v3 := []string{"\n\t", ((((("" + "hello") + "hello") + "hello") + "a") + "foo bar")}
    v4 := []uint32{0, 65535}
    var v5 uint16 = 65535
    v1 -= (v1 * v1)
    fmt.Println(v2, v5, v5)
    v13 := GenS1{F0: true}
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 float64 = 1.5
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        if len(v3) > 0 {
            v3[0] = ("hello" + "")
        }
        v1 = v1
        _ = fi16
    }
    if (!false) {
        if true {
            var v17 float32 = (3.14159 - float32(1e10))
            _ = v17
        }
        switch ("\n\t" + (("\n\t" + "\n\t") + ("a" + "hello"))) {
        case "a":
            fmt.Println(v2, v5)
        case "":
            v1 += v1
        }
        fmt.Println(v2, v5, v5)
    }
    var v18 float32 = (float32(1e10) * 1e10)
    if len(v3) > 0 {
        v3[0] = ("" + ("a" + ((((("a" + ("" + "hello")) + ("" + ("foo bar" + "\n\t"))) + "hello") + "foo bar") + "")))
    }
    var v19 uint8 = 121
    for fi20 := 0; fi20 < 1; fi20++ {
        fmt.Println(v2, v18, v1)
        _ = fi20
    }
    var v21 int64 = 0
    v5 = v5
    v22 := GenS1{F0: true}
    v23 := GenS2{F0: (v2 - (v2 % ((42) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
