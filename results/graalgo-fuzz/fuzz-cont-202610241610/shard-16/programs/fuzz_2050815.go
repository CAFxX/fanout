package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool }
func genFn0(p0 int, p1 int, p2 uint16) string {
    switch 2147483647 {
    case -2147483648:
        var v7 uint64 = 0
        _ = v7
    case 32768:
        switch ("\n\t" + ("\n\t" + (("" + "a") + ("\n\t" + ("hello" + "\n\t"))))) {
        case "hello":
            var v8 int8 = 42
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(p0)
                p0 = (-(p0))
                var v10 int64 = 498
                _ = fi9
                _ = v10
            }
        case "\n\t":
            v11 := []int{}
            _ = v11
        default:
            v12 := GenS1{F0: true}
            _ = v12
        }
    case 786:
        v13 := GenS1{F0: false}
        _ = v13
    default:
        for fi14 := 0; fi14 < 11; fi14++ {
            p1 += (p0 + (p0 | fi14))
            v15 := GenS2{F0: true, F1: 42, F2: 42}
            v16 := GenS3{F0: -1, F1: "foo bar"}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    fmt.Println(p0)
    p2 -= (p2 << ((42) & 63))
    var v17 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return "a"
}
func genFn1(p0 int, p1 int) int {
    var v18 uint32 = 554
    var v19 bool = ((("foo bar" + "\n\t") + "") >= "foo bar")
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := []int{(p0 | (p0 ^ (p0 << ((0) & 63)))), (p1 & p0)}
        _ = fi20
        _ = v21
    }
    switch 2147483647 {
    case -911:
        v22 := []int{}
        _ = v22
    case -204:
        v23 := GenS4{F0: 18446744073709551615, F1: 255}
        _ = v23
    default:
        p0 = (-(p1))
    }
    var v24 string = ((((("foo bar" + "a") + "hello") + "foo bar") + "hello") + "a")
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    _ = v24
    return (p0 ^ (p0 * p0))
}

func main() {
    v1 := []uint32{}
    var v2 uint32 = 2147483647
    v3 := make(map[int]string)
    v4 := make(map[int]int)
    var v5 uint64 = 721
    var v6 bool = (("a" + "foo bar") != ("hello" + "a"))
    delete(v4, 879)
    for fi25 := 0; fi25 < 17; fi25++ {
        v2 += (v2 * 665)
        var v26 int8 = -58
        var v27 int16 = -905
        _ = fi25
        _ = v26
        _ = v27
    }
    v6 = (!true)
    v28 := GenS4{F0: 2, F1: 0}
    for fi29 := 0; fi29 < 16; fi29++ {
        v30 := v4[fi29]
        v6 = v6
        fmt.Println(v2)
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 2; fi31++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v6, v6, v2)
        _ = fi31
    }
    fmt.Println(v5, v6, v6)
    v6 = v6
    var v32 bool = true
    if (2 >= 234) {
        v33 := GenS5{F0: v6}
        v34 := GenS3{F0: -128, F1: ("" + "a")}
        var v35 int = 916
        _ = v33
        _ = v34
        _ = v35
    } else {
        v5 += (v5 - (v5 << ((4294967295) & 63)))
        v2 = v2
    }
    var v36 uint32 = (v2 ^ v2)
    fmt.Println(v36, v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
