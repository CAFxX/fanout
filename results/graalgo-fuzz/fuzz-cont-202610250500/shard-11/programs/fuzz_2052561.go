package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint32 = 42
    var v2 string = ""
    var v3 int16 = 1000
    v4 := []string{""}
    v5 := make(map[uint32]uint32)
    if (v2 <= v2) {
        v3 += (-(v3))
    }
    var v6 uint64 = 42
    var v7 int16 = 93
    if false {
        v3 = v3
        v8 := GenS1{F0: (("a" + v2) + ("a" + (("a" + ((v2 + (((v2 + "") + (v2 + "foo bar")) + "\n\t")) + v2)) + v2)))}
        fmt.Println(v1)
        _ = v8
    } else {
        v5[2147483648] = v1
        _ = len(v4)
    }
    var v9 float64 = 3.14159
    var v10 uint64 = (v6 & 4294967295)
    if (true || true) {
        var v11 int16 = (v3 % ((v7) | 1))
        fmt.Println(v10, v11, v10)
        _ = len(v5)
        _ = v11
    } else {
        v12 := GenS1{F0: (v2 + v2)}
        _ = v12
    }
    fmt.Println(v7, v6, v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v13 uint = 65535
    v14 := GenS1{F0: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
