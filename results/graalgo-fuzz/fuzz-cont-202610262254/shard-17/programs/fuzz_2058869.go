package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint32 }
func genFn0(p0 uint, p1 bool) uint32 {
    p0 = 18446744073709551615
    switch "\n\t" {
    case "a":
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 string = "\n\t"
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: 0, F1: 622, F2: 4294967295}
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    case "":
        var v11 bool = p1
        _ = v11
    case "foo bar":
        p0 += p0
    default:
        v12 := []int{}
        _ = v12
    }
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = (float64(1.5) == ((float64(-2.25) - 3.14159) / 1.5))
    v3 := make(map[int]string)
    v4 := make(map[int]bool)
    var v5 int16 = -600
    var v6 uint64 = 18446744073709551615
    if (v2 && (v2 && false)) {
        for fi13 := 0; fi13 < 2; fi13++ {
            v3[fi13] = "a"
            _ = fi13
        }
    }
    switch (("hello" + (("foo bar" + (("a" + "") + "")) + "hello")) + ("\n\t" + (("foo bar" + ("foo bar" + "hello")) + "\n\t"))) {
    case "":
        fmt.Println(v6, v6, v6)
    default:
        if v2 {
            if v2 {
                v14 := GenS1{F0: -1, F1: 4294967295, F2: 2}
                v6 = v6
                v5 = int16(v6)
                _ = v14
            }
        } else {
            v15 := []int{4294967295, -9223372036854775808, 255, -974}
            if false {
                v15 = append(v15, 765)
                switch "" {
                case "a":
                    fmt.Println(v5, v5, v2)
                }
            } else {
                v2 = v2
                v4[-165] = v2
            }
            v3[723] = ("" + "")
            _ = v15
        }
    }
    var v16 uint = 4294967295
    _ = len(v4)
    var v17 int32 = 42
    var v18 int16 = v5
    v18 = v5
    if ((v18 / ((-1) | 1)) >= (v18 | 1000)) {
        v6 = 553
        v16 -= 0
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(v18, v16)
            switch ("a" + ("a" + "a")) {
            case "foo bar":
                var v21 uint32 = uint32(v6)
                _ = v21
            }
            _ = fi20
        }
        var v22 string = "hello"
        _ = fi19
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
