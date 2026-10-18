package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint, p2 string) int16 {
    var v7 int8 = -49
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 int64 = -452
    var v2 bool = ("foo bar" <= "\n\t")
    v3 := []string{}
    var v4 string = "hello"
    var v5 int8 = 100
    v6 := make(map[string]string)
    if (!(!(!("" > v4)))) {
        v5 -= v5
        v10 := GenS1{F0: (^(v1))}
        _ = v10
    }
    v11 := GenS1{F0: v1}
    v12 := GenS1{F0: v1}
    v5 = v5
    var v13 bool = v2
    delete(v6, ((v4 + ("foo bar" + "")) + v4))
    v13 = false
    v1 = (v1 << ((255) & 63))
    v5 += (v5 - v5)
    var v14 uint32 = 155
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
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
}
