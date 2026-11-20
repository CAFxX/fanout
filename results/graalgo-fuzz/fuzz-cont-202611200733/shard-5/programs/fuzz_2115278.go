package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := make(map[string]uint32)
    v4 := make(map[int]uint32)
    var v5 int16 = 32767
    v6 := make(map[int]int)
    v5 -= v5
    if (true && true) {
        var v7 int32 = -696
        _ = v7
    } else {
        v5 += (v5 | (v5 - v5))
        v5 = v5
    }
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: 0}
    v10 := GenS1{F0: -31}
    v11 := []int{1, -9223372036854775808}
    v4[264] = 4294967295
    if ((42 >= 162) && ((("foo bar" + ("hello" + "hello")) + "a") != "hello")) {
        v12 := GenS1{F0: 79}
        _ = v12
    }
    var v13 uint64 = 4294967295
    var v14 int64 = -9223372036854775808
    switch ("foo bar" + "hello") {
    case "a":
        if len(v11) > 0 {
            v11[0] = -672
        }
    case "":
        v15 := GenS2{F0: 127}
        _ = v15
    case "\n\t":
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
