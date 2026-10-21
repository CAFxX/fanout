package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 string, p1 uint) int8 {
    v5 := GenS1{F0: "\n\t"}
    var v6 float64 = (0.0 * float64(1e10))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := []int{}
    v2 := make(map[uint32]string)
    var v3 float64 = float64(1e10)
    var v4 uint16 = 0
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint8 = 255
        v9 := GenS1{F0: "\n\t"}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: "a", F1: (!true), F2: v4}
    var v11 uint16 = v4
    var v12 int8 = -1
    fmt.Println(v11, v12)
    v3 += v3
    var v13 string = "hello"
    v4 = v11
    v2[1] = v13
    var v14 float64 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!(!true)) {
        v15 := GenS1{F0: v13}
        _ = v15
    } else {
        var v16 float64 = 1e10
        _ = v16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(v14, v3)
        v13 = ("a" + "foo bar")
        if true {
            for fi18 := 0; fi18 < 9; fi18++ {
                fmt.Println(v12)
                _ = fi18
            }
        }
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
