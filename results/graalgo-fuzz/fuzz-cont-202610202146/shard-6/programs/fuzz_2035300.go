package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
func genFn0(p0 int8, p1 bool, p2 uint16) bool {
    v4 := GenS1{F0: 42, F1: 2}
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float32 = ((0.0 * (-2.25 - (1e10 + 0.0))) * 1e10)
        if true {
            var v7 uint8 = 127
            var v8 uint8 = v7
            _ = v7
            _ = v8
        }
        p0 = p0
        _ = fi5
        _ = v6
    }
    switch "foo bar" {
    case "foo bar":
        for fi9 := 0; fi9 < 20; fi9++ {
            switch ("\n\t" + "hello") {
            case "hello":
                v10 := GenS1{F0: 1, F1: 4294967295}
                _ = v10
            }
            var v11 int32 = 42
            _ = fi9
            _ = v11
        }
    case "":
        p0 = p0
    case "a":
        switch 127 {
        case 2147483647:
            v12 := []int{128, -1, -459}
            _ = v12
        default:
            v13 := []int{2147483647, 127}
            _ = v13
        }
    }
    fmt.Println(p2, p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}
func genFn1(p0 uint64, p1 uint16, p2 bool) uint8 {
    for fi14 := 0; fi14 < 17; fi14++ {
        if p2 {
            var v15 uint16 = 42
            _ = v15
        }
        _ = fi14
    }
    p2 = ((true || true) && (42 > -1))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v16 := []int{495}
    var v17 int = 9223372036854775807
    var v18 string = ("\n\t" + "hello")
    delete(v2, "a")
    var v19 int = v17
    if false {
        fmt.Println(v18)
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint8 = 127
            if (false || (false || true)) {
                v18 = ("a" + "hello")
                for fi22 := 0; fi22 < 1; fi22++ {
                    v21 += v21
                    _ = fi22
                }
            }
            var v23 float32 = float32(-2.25)
            _ = fi20
            _ = v21
            _ = v23
        }
    } else {
        if (104 < genFn1(4294967295, 652, true)) {
            var v24 uint = 422
            _ = v24
        }
    }
    if len(v16) > 0 {
        v16[0] = -2147483648
    }
    var v25 int8 = 1
    v17 += -9223372036854775808
    var v26 bool = (1e10 != 1e10)
    v27 := v2[v18]
    var v28 int16 = 42
    v29 := GenS2{F0: (((0.0 >= 3.14159) && true) || true), F1: 127}
    v28 += v28
    v27 += v27
    if v26 {
        if len(v16) > 0 {
            v16[0] = v17
        }
        v18 = (v18 + (("foo bar" + v18) + "\n\t"))
        v30 := GenS2{F0: (2 != 480), F1: 2}
        _ = v30
    } else {
        var v31 int = v17
        var v32 uint16 = 1
        var v33 bool = (v25 < 100)
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
