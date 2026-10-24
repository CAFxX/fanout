package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 bool, p1 int8) int8 {
    v6 := []int{32767, 127, 127}
    p0 = ((("" + "hello") + ("" + ("\n\t" + ((("" + "foo bar") + "foo bar") + "hello")))) > "\n\t")
    switch "\n\t" {
    case "":
        if ((2 > 255) && (p1 < p1)) {
            v7 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
            _ = v7
        }
    case "hello":
        switch 226 {
        case -9223372036854775808:
            v8 := GenS1{F0: -853, F1: 0}
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return p1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    var v3 bool = (2147483647 > 42)
    var v4 int16 = 32767
    var v5 uint64 = v2
    if (!("a" <= "hello")) {
        if (-237 != 32767) {
            v3 = (42 >= 475)
            switch ((((("hello" + "\n\t") + "a") + "hello") + "\n\t") + (("\n\t" + "a") + ("a" + "foo bar"))) {
            case "foo bar":
                v3 = ("" > ("a" + "a"))
            case "a":
                fmt.Println(v3)
            }
        }
    }
    fmt.Println(v5, v2, v2)
    v3 = (0 == genFn0(v3, 35))
    v9 := GenS2{F0: 897, F1: 390, F2: 161}
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint16 = 0
        var v12 float64 = ((float64(float32(0.1)) + 1.5) * ((1.5 * -2.25) * 3.14159))
        _ = fi10
        _ = v11
        _ = v12
    }
    if (!(-2147483648 > -375)) {
        _ = len(v1)
        fmt.Println(v2, v4, v5)
    }
    var v13 uint32 = 0
    var v14 uint64 = (v2 << ((372) & 63))
    var v15 string = ""
    for fi16 := 0; fi16 < 10; fi16++ {
        v15 = "hello"
        var v17 string = v15
        _ = fi16
        _ = v17
    }
    v18 := []int{-415}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
