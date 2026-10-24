package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int16; F2 string }
func genFn0(p0 uint16, p1 int) string {
    var v6 int16 = 1
    v7 := []int{p1}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ((("a" + "foo bar") + (("foo bar" + "a") + "")) + "a")
}
func genFn1() int64 {
    var v8 string = "\n\t"
    v9 := GenS1{F0: (0.1 > 1.5)}
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    var v3 string = ""
    v4 := make(map[int]bool)
    var v5 uint8 = 123
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := GenS2{F0: (196 >= 0), F1: 1, F2: "foo bar"}
        _ = fi10
        _ = v11
    }
    var v12 float64 = 1e10
    var v13 float32 = ((-2.25 * -2.25) - 1e10)
    var v14 int32 = 2147483647
    var v15 int32 = 1
    var v16 int8 = v2
    v17 := []int{42}
    if false {
        var v18 int16 = 1
        _ = v18
    } else {
        switch 11 {
        case 0:
            _ = len(v17)
        case -2147483648:
            var v19 int8 = (v16 + v2)
            _ = v19
        case -9223372036854775808:
            v2 -= (v2 | (v16 + v16))
        }
    }
    fmt.Println(v14, v5, v2)
    var v20 int16 = 791
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
