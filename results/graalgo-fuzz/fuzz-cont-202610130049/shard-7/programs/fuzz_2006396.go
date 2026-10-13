package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int8) string {
    if (-1 <= -613) {
        var v6 uint16 = 42
        v7 := []int{}
        _ = v6
        _ = v7
    }
    var v8 bool = true
    _ = p0
    _ = v8
    return ""
}

func main() {
    v1 := []string{}
    var v2 float64 = (0.0 + -2.25)
    var v3 int32 = 2147483647
    v4 := make(map[int]string)
    v5 := []uint32{1, 0}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v9 string = ""
    fmt.Println(v3)
    fmt.Println(v3, v9, v2)
    v2 -= v2
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint8 = 168
        _ = fi10
        _ = v11
    }
    v9 = ((v9 + "foo bar") + v9)
    switch "foo bar" {
    case "hello":
        var v12 int16 = 667
        _ = v12
    }
    var v13 float64 = v2
    fmt.Println(v2)
    v13 = ((float64((float32(v13) / -2.25)) / float64(1.5)) - float64((1.5 - (((0.1 + (((3.14159 + 3.14159) + 3.14159) * (float32(-2.25) / -2.25))) * 3.14159) - (0.1 / -2.25)))))
    var v14 bool = (int8(v3) >= -1)
    fmt.Println(v2)
    v15 := GenS1{F0: 33}
    var v16 uint = 81
    v9 = ("\n\t" + ("hello" + "a"))
    var v17 uint32 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
