package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 string; F1 string; F2 bool }
func genFn0(p0 int16) uint {
    fmt.Println(p0)
    v4 := []int{32767, 127, 1}
    if (!false) {
        p0 += (-(p0))
        var v5 bool = true
        _ = v5
    }
    var v6 int16 = p0
    _ = p0
    _ = v4
    _ = v6
    return 70
}
func genFn1(p0 int8, p1 string) int {
    var v7 int64 = -739
    v7 = -48
    var v8 uint32 = 263
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32768
}

func main() {
    var v1 float32 = (3.14159 + 1.5)
    v2 := make(map[string]string)
    var v3 uint16 = 42
    if (((-2.25 * 0.0) / 1e10) == -2.25) {
        v9 := GenS1{F0: true}
        v1 += ((-2.25 * (float32(float64(v1)) + 0.0)) * (v1 + (v1 * ((v1 * 1e10) - v1))))
        _ = v9
    }
    v10 := GenS2{F0: 127, F1: 666}
    var v11 float32 = 0.1
    fmt.Println(v3, v11, v1)
    v12 := []int{42, 985}
    v13 := GenS3{F0: 1, F1: 0}
    for fi14 := 0; fi14 < 8; fi14++ {
        v2["a"] = (((("a" + "foo bar") + "\n\t") + "") + "\n\t")
        _ = fi14
    }
    v15 := GenS4{F0: ("foo bar" + ("\n\t" + "")), F1: "foo bar", F2: (true && (1e10 > ((((1.5 * 0.0) + 0.0) + ((1.5 - -2.25) + 1.5)) - 1e10)))}
    var v16 string = ("a" + "\n\t")
    if (878 <= 2) {
        _ = len(v2)
        v1 = (3.14159 / 1.0)
        var v17 uint16 = v3
        _ = v17
    }
    var v18 int32 = 596
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
