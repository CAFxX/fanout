package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 uint16, p1 uint32) int8 {
    v7 := []int{229, 762, 42}
    fmt.Println(p1, p1)
    v7 = append(v7, -1)
    if (!((1e10 - ((1.5 / 1.0) + (((3.14159 + 1e10) + float64(0.0)) + 0.0))) == float64((0.0 + (1e10 / 1e10))))) {
        switch 0 {
        case 0:
            p1 = (p1 & p1)
        }
        var v8 uint32 = (p1 % (((^(p1))) | 1))
        for fi9 := 0; fi9 < 8; fi9++ {
            p0 = p0
            v10 := GenS1{F0: 42}
            v8 += (p1 << ((939) & 63))
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    if false {
        var v11 bool = false
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[string]string)
    v3 := []string{("foo bar" + "\n\t"), "foo bar", "\n\t", (("hello" + "foo bar") + "\n\t")}
    var v4 uint8 = 128
    var v5 int = 32768
    var v6 int8 = 100
    v4 -= 2
    v12 := GenS1{F0: 1}
    if ((v6 | (v6 % ((v6) | 1))) >= (-(v6))) {
        v13 := GenS1{F0: 886}
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v6)
        _ = fi14
    }
    var v15 uint16 = 477
    var v16 uint8 = v4
    v17 := GenS2{F0: v6, F1: 0, F2: 610}
    switch (-(v5)) {
    case 697:
        _ = len(v2)
    case -9223372036854775808:
        delete(v2, "hello")
    }
    v3 = append(v3, ("hello" + ((("hello" + "") + (("a" + "\n\t") + "\n\t")) + ("hello" + "foo bar"))))
    v3 = append(v3, "")
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
