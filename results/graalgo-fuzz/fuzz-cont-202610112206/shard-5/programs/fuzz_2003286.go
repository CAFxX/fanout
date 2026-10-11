package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint16, p1 uint16) uint64 {
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{}
        var v8 string = ("foo bar" + ((("foo bar" + "\n\t") + ("a" + ((("a" + "\n\t") + "a") + ("a" + ("hello" + (((("a" + ("hello" + "a")) + "\n\t") + "") + "foo bar")))))) + ("" + "hello")))
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int = -2147483648
    var v10 uint = uint(p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 1
}
func genFn1() uint8 {
    var v11 float32 = -2.25
    var v12 uint32 = 4294967295
    _ = v11
    _ = v12
    return 95
}

func main() {
    var v1 uint64 = 4294967295
    var v2 float64 = 0.1
    var v3 bool = false
    var v4 int = 0
    v5 := make(map[int]bool)
    v13 := GenS1{F0: 934}
    switch "hello" {
    case "foo bar":
        fmt.Println(v2, v2, v3)
    case "\n\t":
        v3 = v3
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: "foo bar"}
        delete(v5, fi14)
        delete(v5, fi14)
        _ = fi14
        _ = v15
    }
    switch v4 {
    case 128:
        var v16 int64 = 891
        _ = v16
    default:
        v17 := GenS1{F0: 2147483647}
        _ = v17
    }
    v18 := []int{}
    fmt.Println(v4, v4, v2)
    v3 = (1 >= 100)
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 string = "foo bar"
            if v3 {
                var v22 uint32 = 2147483647
                _ = v22
            }
            var v23 bool = (genFn1() >= genFn1())
            _ = fi20
            _ = v21
            _ = v23
        }
        _ = fi19
    }
    if false {
        var v24 uint8 = genFn1()
        var v25 int8 = -65
        fmt.Println(v3)
        _ = v24
        _ = v25
    }
    fmt.Println(v2)
    var v26 int = -1
    var v27 float32 = 1e10
    v28 := GenS1{F0: 698}
    if v3 {
        var v29 uint16 = 895
        var v30 uint8 = 255
        var v31 uint64 = 42
        _ = v29
        _ = v30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
