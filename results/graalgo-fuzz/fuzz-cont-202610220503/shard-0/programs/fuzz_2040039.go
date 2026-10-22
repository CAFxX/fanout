package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 int }
type GenS2 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0(p0 int32) int {
    var v7 string = ("foo bar" + "a")
    v8 := []int{-1}
    _ = p0
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int16) string {
    switch "hello" {
    case "a":
        switch "foo bar" {
        case "\n\t":
            p0 += (p0 % ((p0) | 1))
        }
    case "hello":
        var v9 float32 = 1.5
        _ = v9
    case "":
        var v10 uint64 = 490
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        var v12 uint = 65535
        v13 := GenS1{F0: 2147483648, F1: 28, F2: genFn0(361)}
        _ = fi11
        _ = v12
        _ = v13
    }
    p0 = (-(p0))
    v14 := GenS2{F0: 255, F1: 400, F2: 2}
    _ = p0
    _ = v14
    return ("" + (("foo bar" + ("foo bar" + "a")) + "\n\t"))
}

func main() {
    var v1 uint = 1
    var v2 string = "foo bar"
    var v3 int32 = 42
    v4 := make(map[string]string)
    v5 := []string{"\n\t", ("a" + "a")}
    var v6 uint32 = 2147483648
    v1 -= v1
    v3 -= v3
    fmt.Println(v6, v2)
    var v15 uint32 = v6
    v16 := GenS2{F0: 615, F1: 65535, F2: 170}
    v4[v2] = "\n\t"
    var v17 string = v2
    fmt.Println(v1)
    var v18 float32 = float32(float64(float32(3.14159)))
    v3 = v3
    var v19 string = (v2 + (((((("hello" + "\n\t") + (v2 + v2)) + v2) + v2) + v17) + v17))
    v1 = (v1 % ((v1) | 1))
    var v20 float32 = (v18 - float32(0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
