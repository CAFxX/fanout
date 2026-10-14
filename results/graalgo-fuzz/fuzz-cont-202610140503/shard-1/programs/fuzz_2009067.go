package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 int }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 143
    v3 := make(map[string]uint32)
    var v4 uint64 = 42
    var v5 float64 = -2.25
    var v6 int16 = -32768
    fmt.Println(v5)
    delete(v1, 1)
    var v7 uint = 65535
    fmt.Println(v7, v4, v7)
    v5 = 0.1
    v8 := GenS1{F0: (("foo bar" + ("" + "foo bar")) + ("\n\t" + "hello")), F1: (v4 % (((v4 >> ((v7) & 63))) | 1)), F2: 448}
    var v9 uint = (^(v7))
    fmt.Println(v7, v4)
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        v10 := GenS1{F0: "", F1: 295, F2: 160}
        _ = v10
    case "foo bar":
        var v11 string = ("" + (((("hello" + "\n\t") + ((("\n\t" + "") + "\n\t") + ("hello" + "foo bar"))) + "") + ((("a" + "foo bar") + ("\n\t" + "hello")) + "a")))
        _ = v11
    case "":
        v2 -= (v2 * (v2 % (((v2 & v2)) | 1)))
    }
    switch ("\n\t" + "foo bar") {
    case "a":
        v12 := GenS2{F0: v4}
        _ = v12
    default:
        v13 := []int{}
        _ = v13
    }
    var v14 int32 = 42
    var v15 float64 = v5
    if ((v4 + v4) <= (v4 / (((v4 % ((v4) | 1))) | 1))) {
        v4 = v4
        if (1 > (v7 * v9)) {
            var v16 bool = (-166 <= 913)
            var v17 float32 = -2.25
            _ = v16
            _ = v17
        }
    } else {
        fmt.Println(v7, v5, v2)
    }
    v15 -= 0.1
    var v18 uint8 = 127
    for fi19 := 0; fi19 < 2; fi19++ {
        _ = len(v1)
        _ = fi19
    }
    v20 := GenS2{F0: (v4 >> ((v7) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
