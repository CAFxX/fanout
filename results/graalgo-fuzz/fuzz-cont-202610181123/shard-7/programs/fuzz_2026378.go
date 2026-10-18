package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 string }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 int32, p1 int32) int8 {
    var v7 uint32 = 2
    var v8 uint8 = 42
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 72
}
func genFn1(p0 uint8) int8 {
    switch -9223372036854775808 {
    case -611:
        v9 := []int{1, 9223372036854775807, 1}
        _ = v9
    default:
        p0 = (p0 + (p0 % (((p0 | (^(p0)))) | 1)))
    }
    if (1 < 65535) {
        p0 += 136
        v10 := []int{9223372036854775807}
        _ = v10
    } else {
        p0 = p0
        v11 := []int{604, -2147483648, 4294967295, 1}
        _ = v11
    }
    fmt.Println(p0)
    p0 = p0
    for fi12 := 0; fi12 < 19; fi12++ {
        fmt.Println(p0, p0)
        v13 := []int{fi12, fi12, fi12, 32767}
        _ = fi12
        _ = v13
    }
    _ = p0
    return 127
}

func main() {
    var v1 int = 32768
    v2 := []string{"foo bar", "foo bar", "", (("foo bar" + ("foo bar" + "\n\t")) + ("\n\t" + "\n\t"))}
    v3 := make(map[string]bool)
    var v4 string = ("" + "a")
    var v5 uint16 = 255
    var v6 bool = false
    _ = len(v3)
    fmt.Println(v1, v4, v5)
    var v14 int16 = 1
    for fi15 := 0; fi15 < 19; fi15++ {
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 string = v4
            var v18 int8 = -40
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    fmt.Println(v1)
    v5 = 1
    v19 := GenS1{F0: (v5 << ((4294967295) & 63)), F1: ((1 != 679) && v6), F2: v4}
    for fi20 := 0; fi20 < 12; fi20++ {
        v3["\n\t"] = v6
        fmt.Println(fi20, v5)
        var v21 int32 = 51
        _ = fi20
        _ = v21
    }
    v22 := GenS2{F0: v6, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
