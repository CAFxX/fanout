package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 int8, p2 int16) int8 {
    v6 := GenS1{F0: 127}
    var v7 uint8 = p0
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 bool = false
    var v2 float32 = 0.0
    v3 := []string{}
    var v4 uint16 = 2
    v5 := make(map[int]int)
    var v9 float32 = float32(1.5)
    v9 = v2
    for fi10 := 0; fi10 < 4; fi10++ {
        v9 = 1.5
        _ = fi10
    }
    switch 328 {
    case -213:
        v3 = append(v3, ("a" + "hello"))
    case 2147483647:
        var v11 uint = 1
        _ = v11
    }
    fmt.Println(v9)
    if ((!((0.1 / v2) > ((v2 + (-2.25 + v9)) * v2))) || v1) {
        v1 = ((v4 >> ((405) & 63)) < 914)
        if ("a" == "hello") {
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 float32 = 0.0
                _ = fi12
                _ = v13
            }
        }
    }
    var v14 bool = v1
    delete(v5, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
