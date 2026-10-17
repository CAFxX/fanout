package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 bool }
func genFn0() uint {
    v5 := GenS1{F0: 4294967295, F1: 815, F2: (!true)}
    var v6 bool = false
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 string = ("hello" + ("a" + "a"))
    var v4 string = ((v3 + ("\n\t" + (v3 + v3))) + "foo bar")
    v1 = append(v1, 32)
    var v7 uint32 = 4294967295
    v4 = "hello"
    var v8 bool = (!true)
    if ((0.0 - (1.5 - float32(0.0))) < -2.25) {
        v3 = (v3 + v4)
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 int16 = 1000
            fmt.Println(v10, v8, v4)
            _ = fi9
            _ = v10
        }
    } else {
        var v11 uint32 = v7
        _ = v11
    }
    var v12 string = "\n\t"
    var v13 float32 = (float32((((-2.25 - 0.0) + (-2.25 - 0.0)) + 1e10)) - ((3.14159 - (0.1 - (1e10 + 0.1))) + 3.14159))
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int64 = -257
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi14
        _ = v15
    }
    delete(v2, 42)
    var v16 int64 = -848
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
