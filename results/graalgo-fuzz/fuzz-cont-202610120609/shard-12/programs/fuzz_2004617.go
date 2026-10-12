package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 int16 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int64 = 147
    var v2 int8 = 1
    v3 := make(map[int]bool)
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 2147483647, F1: 65535, F2: 42}
            _ = fi5
            _ = v6
        }
        v7 := GenS2{F0: -39}
        _ = fi4
        _ = v7
    }
    var v8 string = "foo bar"
    fmt.Println(v2, v2, v8)
    fmt.Println(v8, v1, v2)
    var v9 string = (("foo bar" + ("hello" + ((v8 + v8) + "\n\t"))) + "")
    v1 = (-(v1))
    var v10 bool = (!(251 < 127))
    var v11 string = (v9 + v8)
    var v12 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
