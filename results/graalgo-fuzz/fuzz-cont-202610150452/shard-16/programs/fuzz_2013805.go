package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 int16, p2 string) int16 {
    switch p2 {
    case "":
        var v6 float64 = -2.25
        _ = v6
    case "a":
        v7 := GenS1{F0: -2147483648, F1: 359, F2: 32767}
        _ = v7
    }
    p1 += (p1 - p1)
    fmt.Println(p2, p1, p2)
    v8 := GenS1{F0: 530, F1: 853, F2: 9223372036854775807}
    v9 := []int{127}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (p1 / (((p1 * -32768)) | 1))
}
func genFn1(p0 int32, p1 uint8) string {
    var v10 int32 = p0
    for fi11 := 0; fi11 < 2; fi11++ {
        p0 -= (-(v10))
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return ""
}

func main() {
    var v1 string = ""
    var v2 string = "hello"
    var v3 uint64 = 0
    var v4 string = ((v1 + (("a" + ((((("hello" + (v2 + v2)) + v1) + ((v2 + v1) + "\n\t")) + v1) + v1)) + "hello")) + v1)
    v5 := make(map[int]uint32)
    v12 := GenS1{F0: 0, F1: 4294967295, F2: 128}
    fmt.Println(v1, v2)
    var v13 int = 9223372036854775807
    v1 = v2
    fmt.Println(v4, v3)
    for fi14 := 0; fi14 < 9; fi14++ {
        if false {
            v15 := []int{v13, (fi14 % ((-900) | 1)), fi14, v13}
            v2 = (v2 + "foo bar")
            _ = v15
        } else {
            var v16 uint16 = 1
            switch ("a" + v1) {
            case "\n\t":
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := []int{}
                    _ = fi17
                    _ = v18
                }
            case "hello":
                fmt.Println(v2, fi14)
            }
            _ = v16
        }
        var v19 uint = 4294967295
        _ = fi14
        _ = v19
    }
    fmt.Println(v3, v4)
    var v20 bool = false
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v2)
        _ = fi21
    }
    v22 := GenS2{F0: 18446744073709551615}
    if (!v20) {
        fmt.Println(v1, v2, v2)
    } else {
        v23 := []int{v13, (v13 % ((v13) | 1)), 1, (v13 / ((v13) | 1))}
        _ = v23
    }
    fmt.Println(v3, v2)
    v24 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
