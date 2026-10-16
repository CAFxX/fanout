package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = 0.1
    v3 := []int{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 := GenS1{F0: "foo bar"}
    var v5 bool = ((1 == 65535) && ("\n\t" < "hello"))
    var v6 int = -145
    v2 = (0.0 - v2)
    var v7 uint8 = 1
    _ = len(v3)
    fmt.Println(v7)
    _ = len(v3)
    if false {
        var v8 int32 = -1
        _ = v8
    }
    delete(v1, "\n\t")
    v5 = ((v6 << ((uint(v7)) & 63)) > (-(v6)))
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float64 = (v2 - 3.14159)
        _ = fi9
        _ = v10
    }
    v1["\n\t"] = ("hello" + "foo bar")
    var v11 uint = 350
    var v12 uint8 = (^(v7))
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    if v5 {
        var v13 int64 = 1
        var v14 uint64 = 921
        _ = v13
        _ = v14
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
