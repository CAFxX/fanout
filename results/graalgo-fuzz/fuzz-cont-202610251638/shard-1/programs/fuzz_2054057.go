package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float32 = float32(0.0)
        _ = fi5
        _ = v6
    }
    var v7 int16 = 32767
    var v8 int64 = 1
    v8 = v8
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := []int{}
    var v2 uint = 120
    var v3 string = (("a" + "\n\t") + "\n\t")
    var v4 int = 42
    v4 += 4294967295
    fmt.Println(v2, v4)
    var v9 bool = ((!false) || (0 != 2147483647))
    v10 := GenS1{F0: (v2 != 2), F1: -1}
    if (v9 && v9) {
        var v11 uint32 = 42
        if true {
            v3 = (("hello" + v3) + ((("" + v3) + v3) + (("hello" + ((v3 + "foo bar") + "")) + v3)))
            fmt.Println(v11)
        }
        _ = v11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v12 := GenS1{F0: v9, F1: 9223372036854775807}
    fmt.Println(v9, v9, v3)
    for fi13 := 0; fi13 < 2; fi13++ {
        v4 = -1
        var v14 string = "foo bar"
        v9 = (213 == 2)
        _ = fi13
        _ = v14
    }
    var v15 int64 = 791
    var v16 int32 = -2147483648
    if len(v1) > 0 {
        v1[0] = (v4 / ((v4) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
