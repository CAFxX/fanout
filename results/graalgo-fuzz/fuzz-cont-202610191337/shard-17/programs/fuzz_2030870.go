package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string; F1 int }
func genFn0(p0 int8, p1 int64, p2 int64) string {
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            switch (("foo bar" + "hello") + ("a" + "\n\t")) {
            case "foo bar":
                var v7 uint32 = 2147483647
                var v8 string = "foo bar"
                _ = v7
                _ = v8
            }
            p2 -= p1
            p2 -= (p2 << ((0) & 63))
            v9 := GenS1{F0: 0, F1: 42}
            _ = fi6
            _ = v9
        }
        if false {
            v10 := []int{fi5, (fi5 % ((fi5) | 1)), fi5}
            _ = v10
        }
        p2 -= (p1 ^ p1)
        _ = fi5
    }
    if ((p2 - 42) != p2) {
        v11 := GenS1{F0: 592, F1: 1}
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if (18446744073709551615 > 655) {
            v13 := []int{fi12}
            _ = v13
        }
        _ = fi12
    }
    v14 := []int{}
    p0 = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return ((("a" + "hello") + "foo bar") + ("foo bar" + ""))
}
func genFn1(p0 int, p1 bool, p2 int32) int16 {
    fmt.Println(p0)
    v15 := []int{-9223372036854775808, 39, -2147483648, p0}
    if p1 {
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int64 = 489
            p0 -= (fi16 + 32768)
            _ = fi16
            _ = v17
        }
    }
    var v18 int = (p0 / ((p0) | 1))
    var v19 int = v18
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v18
    _ = v19
    return 32767
}

func main() {
    var v1 int16 = -1
    var v2 bool = (!true)
    var v3 int8 = 127
    var v4 bool = (0 == 749)
    fmt.Println(v1)
    if ((v2 || v2) && v2) {
        switch -504 {
        case -322:
            v4 = v4
        case 411:
            v20 := GenS2{F0: v2, F1: -198}
            _ = v20
        }
        fmt.Println(v3)
        var v21 float32 = 1.5
        _ = v21
    }
    v22 := GenS1{F0: 478, F1: (v1 * v1)}
    v23 := GenS1{F0: 508, F1: genFn1(825, (1 <= 0), 64)}
    if (18446744073709551615 > 564) {
        v24 := []int{-373}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 int8 = v3
            v27 := []int{32768, (-(fi25)), -369, (fi25 % (((fi25 | (-(fi25)))) | 1))}
            v2 = v4
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 string = "foo bar"
            v4 = (v29 >= "a")
            fmt.Println(v29, v3)
            _ = fi28
            _ = v29
        }
        v4 = (1e10 >= (3.14159 / (3.14159 * (float64(-2.25) / 0.1))))
    }
    switch "\n\t" {
    case "a":
        v1 -= (^(v1))
    case "foo bar":
        v2 = v2
    case "\n\t":
        fmt.Println(v1, v3)
    default:
        switch 684 {
        case 128:
            var v30 bool = (-2147483648 < 42)
            _ = v30
        case 2147483647:
            v31 := GenS1{F0: 262, F1: (v1 * (-(v1)))}
            _ = v31
        }
    }
    var v32 int32 = 757
    v33 := []int{956, -9223372036854775808, 1, -2147483648}
    v1 = genFn1(0, v2, -1)
    v34 := GenS2{F0: (0 > 540), F1: 0}
    v3 = 20
    var v35 uint8 = 181
    if (!v2) {
        for fi36 := 0; fi36 < 12; fi36++ {
            v37 := GenS3{F0: "a", F1: (^(fi36))}
            v2 = v2
            _ = fi36
            _ = v37
        }
    }
    v33 = append(v33, 0)
    v38 := GenS1{F0: 547, F1: 427}
    if (v1 == genFn1(4294967295, v2, v32)) {
        var v39 float32 = 1e10
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
