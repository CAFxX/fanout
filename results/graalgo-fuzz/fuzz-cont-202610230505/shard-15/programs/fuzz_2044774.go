package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 uint8, p1 uint64, p2 int64) string {
    switch "\n\t" {
    case "hello":
        p0 = p0
    case "":
        v4 := GenS1{F0: 65535}
        _ = v4
    }
    p2 = -754
    p1 += 566
    _ = p0
    _ = p1
    _ = p2
    return "a"
}
func genFn1() uint8 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int32 = -72
        fmt.Println(fi5, v6)
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 255
    var v8 float32 = (0.0 * 0.0)
    var v9 uint16 = v7
    _ = v7
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ""
    var v3 bool = (!((-2.25 - 0.1) <= 3.14159))
    if v3 {
        v3 = v3
    }
    v10 := []int{}
    var v11 int8 = -1
    var v12 float32 = (float32(3.14159) + float32(1e10))
    fmt.Println(v11, v11, v2)
    var v13 string = v2
    v12 = ((v12 - (-2.25 / v12)) / (3.14159 + float32((((1.5 * 1e10) * 3.14159) - (3.14159 - 0.1)))))
    var v14 string = v2
    var v15 uint64 = 2
    v16 := v1[906]
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint = 734
        _ = fi17
        _ = v18
    }
    delete(v1, 4294967295)
    for fi19 := 0; fi19 < 7; fi19++ {
        fmt.Println(v15, v12)
        _ = fi19
    }
    if (!v3) {
        v20 := GenS1{F0: 184}
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
