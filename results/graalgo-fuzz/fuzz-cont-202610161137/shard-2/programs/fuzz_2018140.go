package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint8, p1 string) int64 {
    v5 := []int{2147483647, 128, 604}
    p1 = p1
    p1 = (p1 + (p1 + p1))
    for fi6 := 0; fi6 < 13; fi6++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 558
    v3 := []uint32{v2, v2}
    var v4 uint32 = v2
    delete(v1, "foo bar")
    var v7 int64 = 169
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint32 = 1
        _ = fi8
        _ = v9
    }
    v7 = v7
    switch 462 {
    case 32768:
        v3 = append(v3, v4)
    case 0:
        var v10 uint64 = 4294967295
        _ = v10
    }
    fmt.Println(v2, v4, v4)
    _ = len(v3)
    v11 := GenS1{F0: 18446744073709551615, F1: (654 >= 524)}
    var v12 uint64 = 200
    v12 = (v12 & (v12 | (v12 / ((0) | 1))))
    fmt.Println(v12)
    v13 := GenS1{F0: 4294967295, F1: false}
    for fi14 := 0; fi14 < 15; fi14++ {
        if ((^(fi14)) < (^(fi14))) {
            v4 += 2147483648
            v15 := GenS2{F0: (((("a" + ("\n\t" + "\n\t")) + (((((("a" + ("\n\t" + ((("\n\t" + "\n\t") + "a") + ("a" + "foo bar")))) + ("a" + ("hello" + (("a" + "") + ("\n\t" + "hello"))))) + ("\n\t" + "a")) + "foo bar") + "hello") + "")) + "a") + (("a" + "hello") + "hello")), F1: -1}
            v16 := GenS3{F0: 255, F1: v12}
            _ = v15
            _ = v16
        }
        v17 := GenS4{F0: 127, F1: ("" + "\n\t")}
        _ = fi14
        _ = v17
    }
    if ((((3.14159 * 0.1) - 3.14159) * ((1e10 * -2.25) + 1.5)) >= 0.1) {
        v1[("hello" + "foo bar")] = v4
    }
    var v18 uint64 = v12
    v4 = 984
    v19 := GenS4{F0: 32, F1: "\n\t"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
