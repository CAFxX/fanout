package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint16) int8 {
    fmt.Println(p1, p1)
    for fi5 := 0; fi5 < 9; fi5++ {
        p0 = uint32(p1)
        _ = fi5
    }
    switch "foo bar" {
    case "\n\t":
        v6 := []int{259, 127}
        _ = v6
    }
    v7 := GenS1{F0: -30, F1: (^(p1))}
    v8 := []int{415, 127, -2147483648, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{}
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483648
    var v4 float64 = 0.1
    v4 = (0.0 + float64(-2.25))
    _ = len(v1)
    v3 = (v3 % ((v3) | 1))
    var v9 uint = 255
    v10 := GenS1{F0: -98, F1: 0}
    v11 := v2[4294967295]
    v12 := GenS1{F0: 127, F1: 32767}
    switch ("" + (("hello" + "\n\t") + "")) {
    case "a":
        v13 := GenS1{F0: 1, F1: 2}
        _ = v13
    case "foo bar":
        fmt.Println(v3)
    case "\n\t":
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 string = (("a" + ("" + ("foo bar" + (("" + "foo bar") + "\n\t")))) + "")
            if (v11 || ((v4 * (v4 / -2.25)) <= v4)) {
                v16 := GenS1{F0: 39, F1: 742}
                v17 := GenS1{F0: 0, F1: 399}
                v18 := GenS2{F0: -128}
                _ = v16
                _ = v17
                _ = v18
            }
            switch v15 {
            case "":
                for fi19 := 0; fi19 < 13; fi19++ {
                    v11 = v11
                    _ = fi19
                }
            }
            _ = fi14
            _ = v15
        }
    }
    var v20 string = ("foo bar" + "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
