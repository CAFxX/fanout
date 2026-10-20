package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 string, p1 uint, p2 bool) uint {
    v4 := []int{}
    if len(v4) > 0 {
        v4[0] = 1
    }
    switch ((p0 + p0) + p0) {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        p0 = ("" + "hello")
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 278
}

func main() {
    var v1 uint32 = 498
    v2 := make(map[uint32]bool)
    var v3 string = "\n\t"
    for fi5 := 0; fi5 < 5; fi5++ {
        v3 = ""
        _ = fi5
    }
    if (948 <= 1) {
        var v6 uint = genFn0(("\n\t" + ((((("\n\t" + "hello") + "") + v3) + v3) + "a")), 65535, ((^(v1)) != v1))
        if true {
            v1 = 559
            v1 += v1
        }
        v7 := GenS1{F0: v6}
        _ = v6
        _ = v7
    } else {
        fmt.Println(v1, v3)
        v8 := GenS1{F0: 165}
        _ = v8
    }
    var v9 float64 = ((((((1.5 / 1.5) - 1e10) - 0.0) + 0.0) - float64(3.14159)) * 1.5)
    v10 := []int{-1}
    var v11 int = -856
    if len(v10) > 0 {
        v10[0] = -565
    }
    v12 := GenS1{F0: 1}
    var v13 uint32 = 65535
    fmt.Println(v9)
    v1 = v13
    v10 = append(v10, v11)
    var v14 float64 = float64((1.5 + 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v9
}
