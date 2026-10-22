package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]int)
    var v2 int = -534
    v3 := []string{"\n\t", ("hello" + ("a" + ""))}
    var v4 int64 = 1
    var v5 uint = 255
    v6 := make(map[int]uint32)
    v3 = append(v3, "hello")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v4 = 9223372036854775807
    } else {
        fmt.Println(v5)
    }
    v2 = 0
    v5 -= v5
    if len(v3) > 0 {
        v3[0] = ("\n\t" + ("" + ("foo bar" + "foo bar")))
    }
    var v7 int32 = 582
    v7 -= (v7 % ((v7) | 1))
    var v8 string = "foo bar"
    v2 += v2
    var v9 float64 = 0.0
    var v10 int64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
