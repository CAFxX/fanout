package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
func genFn0(p0 uint16, p1 int8) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 int8 = p1
        _ = fi4
        _ = v5
    }
    var v6 string = ""
    p1 = p1
    v7 := GenS1{F0: 0, F1: false}
    if true {
        var v8 int64 = 42
        p1 += p1
        var v9 int32 = 2147483647
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 715
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 627
    var v3 int = -908
    v10 := GenS1{F0: 127, F1: true}
    if true {
        v11 := []int{497, v3, (^(v3)), v3}
        if len(v11) > 0 {
            v11[0] = v3
        }
        _ = v11
    } else {
        v1["\n\t"] = (("a" + "") + "\n\t")
        _ = len(v1)
    }
    var v12 uint32 = (v2 / (((^(v2))) | 1))
    var v13 bool = (100 != 0)
    v3 = v3
    var v14 uint8 = 104
    v13 = ((!v13) && v13)
    if v13 {
        for fi15 := 0; fi15 < 16; fi15++ {
            delete(v1, "hello")
            v2 = 2
            var v16 float32 = (0.1 - (-2.25 - (((0.1 + 3.14159) / 1.0) * 1e10)))
            _ = fi15
            _ = v16
        }
    }
    v14 += v14
    v12 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
