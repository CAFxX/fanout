package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 0.1
    var v3 int16 = 758
    switch ("foo bar" + ("foo bar" + ("a" + ("\n\t" + ("foo bar" + "foo bar"))))) {
    case "a":
        v1 = v1
    case "":
        fmt.Println(v2, v3)
    case "foo bar":
        if ((v1 < (v1 ^ (v1 * (v1 ^ v1)))) && false) {
            v4 := []int{127, 128, -9223372036854775808}
            v5 := GenS1{F0: v3, F1: 127}
            _ = v4
            _ = v5
        }
    }
    v6 := GenS1{F0: v3, F1: 1}
    v7 := GenS2{F0: 1, F1: (127 <= -128)}
    if true {
        v8 := []int{32768, 962}
        var v9 int8 = -15
        _ = v8
        _ = v9
    }
    fmt.Println(v3, v2)
    switch ("a" + "a") {
    case "":
        v10 := GenS1{F0: v3, F1: 61}
        _ = v10
    }
    var v11 uint32 = 449
    v1 = 9223372036854775807
    v2 = float64(3.14159)
    for fi12 := 0; fi12 < 13; fi12++ {
        v11 += 1
        fmt.Println(v2, v3)
        if false {
            v13 := GenS3{F0: fi12}
            _ = v13
        }
        _ = fi12
    }
    v14 := []int{128}
    var v15 int16 = (v3 * 42)
    var v16 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
