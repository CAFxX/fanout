package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 int16) int {
    fmt.Println(p1, p0)
    v5 := GenS1{F0: 65535}
    v6 := []int{}
    switch 42 {
    case -747:
        v7 := GenS1{F0: 18446744073709551615}
        _ = v7
    case 127:
        if p0 {
            v8 := GenS1{F0: 4294967295}
            var v9 string = "hello"
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 string = ""
    var v2 float64 = -2.25
    v3 := make(map[string]int)
    v4 := []int{32767, 2147483647, -971}
    v10 := GenS1{F0: 979}
    v11 := GenS1{F0: 766}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v2 -= v2
    v1 = (v1 + "foo bar")
    var v12 bool = false
    var v13 int = 0
    v2 = 3.14159
    v2 -= v2
    v12 = false
    v14 := GenS1{F0: 2}
    var v15 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
