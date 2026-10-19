package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8, p1 string, p2 int) uint8 {
    v6 := []int{32768}
    v7 := GenS1{F0: 42}
    var v8 uint8 = p0
    fmt.Println(p0)
    fmt.Println(p2, v8, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    v3 := make(map[int]int)
    var v4 uint = v2
    var v5 float64 = -2.25
    var v9 int = -906
    v9 = (v9 << (((v2 % (((^(v4))) | 1))) & 63))
    v9 -= v9
    var v10 uint8 = 42
    v11 := v3[v9]
    v12 := GenS2{F0: (!false), F1: 48}
    v13 := GenS2{F0: ((110 <= -1) || false), F1: 2147483648}
    switch (((("" + "foo bar") + ((("foo bar" + (("" + "\n\t") + "a")) + "\n\t") + ("foo bar" + "hello"))) + "\n\t") + "hello") {
    case "a":
        v14 := v3[-837]
        _ = v14
    }
    var v15 string = ""
    var v16 uint32 = 1
    var v17 uint8 = v10
    var v18 float64 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
