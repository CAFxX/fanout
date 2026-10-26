package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int16 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }

func main() {
    var v1 uint = 42
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v1 = 521
    var v4 string = (("hello" + (("\n\t" + (("hello" + "foo bar") + "\n\t")) + "foo bar")) + ("" + "a"))
    v4 = ("foo bar" + "hello")
    v5 := GenS1{F0: 378, F1: 0, F2: 86}
    fmt.Println(v4)
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 int = (fi6 ^ fi6)
        _ = fi6
        _ = v7
    }
    switch (("hello" + v4) + (v4 + "")) {
    case "\n\t":
        delete(v2, ("hello" + v4))
    case "foo bar":
        v8 := GenS2{F0: 2147483647, F1: 0, F2: 18446744073709551615}
        _ = v8
    case "":
        for fi9 := 0; fi9 < 4; fi9++ {
            if true {
                v10 := []int{(fi9 >> ((v1) & 63)), fi9, -263}
                _ = v10
            } else {
                v11 := []int{(-(fi9))}
                _ = v11
            }
            _ = fi9
        }
    default:
        v12 := GenS3{F0: (int8(v1) < 0), F1: (v1 % ((0) | 1)), F2: 42}
        _ = v12
    }
    v13 := v2[(v4 + "")]
    var v14 string = v13
    v1 = (^(v1))
    fmt.Println(v14, v13, v14)
    v15 := GenS1{F0: 978, F1: 1, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
