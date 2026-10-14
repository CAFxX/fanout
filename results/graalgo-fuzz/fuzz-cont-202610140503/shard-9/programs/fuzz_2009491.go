package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }
func genFn0() int16 {
    if true {
        var v5 int = 42
        _ = v5
    }
    var v6 int8 = 100
    _ = v6
    return -1
}
func genFn1(p0 int, p1 uint16) int16 {
    var v7 uint32 = 2
    var v8 uint = 65535
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = 0
    var v3 int = 42
    v4 := []string{"a"}
    if true {
        var v9 uint64 = (v1 / (((v1 % (((v1 | (v1 / ((v1) | 1)))) | 1))) | 1))
        v10 := GenS1{F0: "", F1: -872, F2: ""}
        var v11 string = "a"
        _ = v10
        _ = v11
        _ = v9
    } else {
        fmt.Println(v3)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        fmt.Println(v3, v1, v2)
        v1 -= 1
    }
    fmt.Println(v1, v2)
    var v12 bool = (863 > 1)
    v13 := GenS2{F0: v2, F1: 723, F2: (v1 % ((v1) | 1))}
    var v14 float64 = ((3.14159 - (1e10 * 3.14159)) + 0.1)
    var v15 int8 = 1
    switch ("\n\t" + "hello") {
    case "a":
        v3 = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
