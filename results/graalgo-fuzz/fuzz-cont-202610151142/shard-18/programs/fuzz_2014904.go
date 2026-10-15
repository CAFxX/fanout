package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int, p1 uint) int64 {
    v5 := []int{0, p0}
    fmt.Println(p1, p0)
    p0 = (p0 << (((^(p1))) & 63))
    _ = p0
    _ = p1
    _ = v5
    return 643
}

func main() {
    var v1 bool = ((("foo bar" + ("" + ("" + ("a" + "")))) + "hello") <= ("" + (("\n\t" + ("a" + "a")) + "foo bar")))
    var v2 uint64 = 708
    v3 := make(map[string]bool)
    v4 := make(map[uint32]int)
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 bool = (!(false || (!false)))
            var v9 uint32 = 65535
            _ = fi7
            _ = v8
            _ = v9
        }
        switch (("" + "\n\t") + ("" + "foo bar")) {
        case "\n\t":
            delete(v3, "")
        case "a":
            var v10 int16 = 1000
            _ = v10
        }
        _ = fi6
    }
    v2 = (v2 % ((v2) | 1))
    _ = len(v3)
    v1 = v1
    v1 = true
    switch -1 {
    case 1:
        v2 += (v2 + v2)
    }
    switch 1 {
    case 0:
        v2 = v2
    case 223:
        v11 := GenS1{F0: 592}
        _ = v11
    }
    var v12 int8 = 2
    v13 := GenS1{F0: 520}
    fmt.Println(v2, v12, v1)
    var v14 string = ""
    v15 := GenS1{F0: 2}
    switch v14 {
    case "":
        var v16 float32 = 0.0
        _ = v16
    case "foo bar":
        fmt.Println(v2, v12, v2)
    case "hello":
        var v17 int32 = 42
        _ = v17
    default:
        switch "\n\t" {
        case "a":
            v2 -= v2
        }
    }
    _ = len(v3)
    var v18 float64 = 1e10
    fmt.Println(v1)
    for fi19 := 0; fi19 < 15; fi19++ {
        v20 := GenS1{F0: 2147483647}
        v12 = (v12 & (v12 % ((0) | 1)))
        _ = fi19
        _ = v20
    }
    var v21 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
