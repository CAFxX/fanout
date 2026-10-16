package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0(p0 uint16, p1 int) string {
    for fi6 := 0; fi6 < 15; fi6++ {
        switch "a" {
        case "\n\t":
            var v7 string = ""
            _ = v7
        case "a":
            var v8 int8 = 22
            _ = v8
        case "":
            var v9 uint8 = 128
            _ = v9
        }
        v10 := GenS1{F0: 0, F1: (((("foo bar" + ("" + "")) + ((("hello" + ("a" + "\n\t")) + "a") + "hello")) + ("hello" + ("" + ""))) + ("\n\t" + "hello"))}
        _ = fi6
        _ = v10
    }
    v11 := GenS2{F0: 42, F1: 127, F2: 1}
    v12 := []int{p1, (p1 * -2147483648), 0, (p1 - p1)}
    p0 = p0
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return "hello"
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 42
    v3 := []string{"hello", "hello", "foo bar", ("foo bar" + "a")}
    var v4 int = -9223372036854775808
    v5 := make(map[uint32]uint32)
    var v13 int16 = -472
    v2 = v2
    if false {
        var v14 uint8 = 28
        v15 := GenS2{F0: -106, F1: 127, F2: 670}
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 int = -2147483648
        _ = fi16
        _ = v17
    }
    v13 = (v13 * 0)
    fmt.Println(v4, v13)
    var v18 int8 = 13
    var v19 uint32 = 0
    fmt.Println(v2, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
