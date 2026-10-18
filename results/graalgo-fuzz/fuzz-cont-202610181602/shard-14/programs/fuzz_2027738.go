package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8) int16 {
    p0 += 1
    p0 = p0
    fmt.Println(p0)
    v4 := []int{4294967295, 4294967295}
    _ = p0
    _ = v4
    return -991
}

func main() {
    var v1 float32 = 3.14159
    var v2 int64 = -1
    v3 := []int{42, 1}
    if (1.5 <= 3.14159) {
        switch 2147483647 {
        case 2147483647:
            if (!true) {
                v5 := GenS1{F0: 127, F1: "a"}
                _ = v5
            }
        case 1:
            var v6 uint = 42
            _ = v6
        case -1:
            var v7 string = ("a" + "a")
            _ = v7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 string = ""
            v2 += -9223372036854775808
            v1 -= 1.5
            _ = fi8
            _ = v9
        }
        fmt.Println(v1)
    } else {
        v2 -= v2
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v10 float32 = float32(0.0)
    switch "hello" {
    case "foo bar":
        fmt.Println(v10, v10, v1)
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := GenS1{F0: fi12, F1: "a"}
            _ = fi12
            _ = v13
        }
        _ = len(v3)
        _ = fi11
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        v1 = float32(((-2.25 + float64((v10 / (v1 / 1e10)))) + (0.1 + 1e10)))
        _ = fi14
    }
    var v15 int32 = -1
    switch "hello" {
    case "a":
        v1 = 1.5
    case "":
        if true {
            v16 := GenS2{F0: 42, F1: (v2 - v2)}
            v17 := GenS2{F0: v2, F1: v2}
            _ = v16
            _ = v17
        } else {
            var v18 uint64 = 4294967295
            switch (("foo bar" + ("foo bar" + "hello")) + ((("\n\t" + "a") + "foo bar") + "hello")) {
            case "":
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int8 = -128
                    _ = fi19
                    _ = v20
                }
            }
            _ = v18
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v21 float64 = ((0.0 * 0.0) - (-2.25 + 0.0))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v21
    _ = v3
}
