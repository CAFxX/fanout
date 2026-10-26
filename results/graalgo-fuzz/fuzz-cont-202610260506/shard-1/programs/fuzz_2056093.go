package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint8; F2 int8 }
type GenS4 struct { F0 int8; F1 int32; F2 int16 }

func main() {
    v1 := []int{0}
    var v2 string = ("foo bar" + (("\n\t" + "foo bar") + "a"))
    var v3 float64 = float64(float32(0.0))
    v1 = append(v1, 997)
    v4 := GenS1{F0: -32768, F1: (!(128 < 255)), F2: -507}
    v2 = (((v2 + (v2 + ((("hello" + v2) + v2) + (("a" + (v2 + v2)) + "")))) + v2) + "foo bar")
    fmt.Println(v2, v2)
    v5 := GenS2{F0: ("a" + v2)}
    v6 := GenS3{F0: -476, F1: 0, F2: 55}
    v7 := GenS4{F0: 100, F1: -1, F2: -1}
    var v8 uint = 371
    var v9 int32 = 397
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
