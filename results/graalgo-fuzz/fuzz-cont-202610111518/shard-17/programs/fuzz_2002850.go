package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 string }
type GenS2 struct { F0 string; F1 uint32 }
func genFn0(p0 int32, p1 uint8) uint64 {
    p0 += p0
    p0 -= 0
    p0 = (p0 >> ((134) & 63))
    for fi5 := 0; fi5 < 6; fi5++ {
        var v6 string = "foo bar"
        var v7 string = v6
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                fmt.Println(p1, p0, p1)
                p0 = -2147483648
                _ = fi9
            }
            var v10 string = ("\n\t" + "\n\t")
            _ = fi8
            _ = v10
        }
        p1 -= (^(p1))
    }
    _ = p0
    _ = p1
    return 688
}
func genFn1(p0 int16, p1 uint) string {
    var v11 int16 = 530
    v12 := GenS1{F0: -2147483648, F1: "hello", F2: (("a" + ("foo bar" + "a")) + "\n\t")}
    v13 := GenS2{F0: "\n\t", F1: 855}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    var v1 int32 = -2147483648
    var v2 int8 = 69
    var v3 float64 = float64((1.5 * (1e10 / 0.1)))
    var v4 string = "foo bar"
    fmt.Println(v1)
    for fi14 := 0; fi14 < 6; fi14++ {
        if (true || false) {
            var v15 bool = (!(false || (!true)))
            _ = v15
        } else {
            fmt.Println(v1)
            v16 := GenS1{F0: v1, F1: "\n\t", F2: "hello"}
            _ = v16
        }
        _ = fi14
    }
    var v17 int64 = 42
    if true {
        v18 := []int{2147483647, 128, -1}
        var v19 uint64 = 659
        var v20 bool = false
        _ = v18
        _ = v19
        _ = v20
    }
    var v21 float64 = 1e10
    v22 := GenS2{F0: ("hello" + v4), F1: 42}
    fmt.Println(v3, v21)
    var v23 uint = 883
    fmt.Println(v21)
    switch ((v4 + v4) + ("foo bar" + "")) {
    case "foo bar":
        v24 := GenS2{F0: "\n\t", F1: 0}
        _ = v24
    case "hello":
        switch 32768 {
        case 0:
            var v25 uint32 = 494
            _ = v25
        }
    }
    var v26 uint64 = 2
    v2 += v2
    v27 := []int{4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
