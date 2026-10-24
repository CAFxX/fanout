package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint, p1 int32, p2 int) uint32 {
    fmt.Println(p1, p0)
    p0 += p0
    var v6 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 uint16, p1 int8) int {
    var v7 uint8 = 0
    v8 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -609
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{127}
    v3 := make(map[int]string)
    var v4 int8 = 7
    v5 := make(map[int]bool)
    v9 := GenS1{F0: 0}
    fmt.Println(v4)
    if (false || ("foo bar" > "foo bar")) {
        if len(v2) > 0 {
            v2[0] = -643
        }
        fmt.Println(v4, v4)
    }
    v4 -= v4
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := v5[(-(fi10))]
        for fi12 := 0; fi12 < 1; fi12++ {
            v11 = ((true && v11) && v11)
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    v13 := GenS2{F0: "\n\t"}
    var v14 int = 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 = (v14 >> ((255) & 63))
    var v15 int = v14
    fmt.Println(v15, v14)
    if (!true) {
        fmt.Println(v4)
    } else {
        v2 = append(v2, (v14 >> ((4294967295) & 63)))
        var v16 float32 = 0.0
        var v17 string = v1
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
