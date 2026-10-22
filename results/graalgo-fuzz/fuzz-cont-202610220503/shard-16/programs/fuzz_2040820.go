package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int32, p1 int16) int8 {
    p1 += p1
    var v7 uint16 = 65535
    var v8 uint = 1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 float32 = 3.14159
    var v2 int = 128
    var v3 string = "hello"
    var v4 uint8 = 166
    v5 := make(map[string]int)
    v6 := []int{(-(v2)), (^(v2)), (v2 ^ v2)}
    if len(v6) > 0 {
        v6[0] = v2
    }
    var v9 uint32 = 2147483648
    switch v2 {
    case 709:
        var v10 int16 = -187
        _ = v10
    case -1:
        v3 = (v3 + "\n\t")
    }
    var v11 int16 = 419
    v12 := GenS1{F0: ("\n\t" + (((v3 + v3) + "\n\t") + (v3 + (v3 + ((v3 + ("" + v3)) + v3)))))}
    switch (-(v2)) {
    case 995:
        fmt.Println(v2)
    default:
        v5[(v3 + v3)] = v2
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v13 int32 = 2147483647
    var v14 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
