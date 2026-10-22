package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 bool }
func genFn0(p0 int16, p1 int64) uint32 {
    if false {
        v7 := []int{}
        _ = v7
    } else {
        var v8 int16 = p0
        _ = v8
    }
    switch ((("a" + ((("" + "\n\t") + "") + ("hello" + ((("\n\t" + "") + "foo bar") + ("" + "hello"))))) + ("a" + "\n\t")) + "\n\t") {
    case "":
        v9 := GenS1{F0: 1, F1: 816, F2: true}
        _ = v9
    case "\n\t":
        p0 = (p0 >> ((0) & 63))
    case "hello":
        fmt.Println(p1, p0)
    default:
        p0 = 42
    }
    _ = p0
    _ = p1
    return 342
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 0
    var v3 bool = true
    var v4 float64 = 1e10
    v5 := make(map[int]int)
    var v6 int = 127
    if (false && v3) {
        fmt.Println(v4)
        var v10 int8 = 127
        _ = v10
    }
    var v11 uint32 = 0
    var v12 int32 = 1
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 float32 = 0.0
        if v3 {
            var v15 int32 = 1
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 int = 0
    if ((^(v2)) <= (v2 % ((0) | 1))) {
        if ((v11 != (v11 ^ (v11 / ((514) | 1)))) && (!(65535 < 4294967295))) {
            var v17 string = "foo bar"
            _ = v17
        }
        var v18 uint16 = v2
        _ = v18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v4 -= float64(1.5)
            _ = fi20
        }
        _ = fi19
    }
    var v21 int8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
