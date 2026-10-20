package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0(p0 uint64, p1 int8, p2 uint32) int32 {
    var v7 int32 = int32(p1)
    v7 -= v7
    v8 := GenS1{F0: p2, F1: -32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 486
}

func main() {
    var v1 int = -2147483648
    v2 := []string{"hello", "a"}
    var v3 uint64 = 12
    v4 := make(map[string]int)
    var v5 int8 = 120
    var v6 int32 = 0
    _ = len(v2)
    v9 := v4["\n\t"]
    v6 = 2147483647
    var v10 uint32 = 1
    v10 = (^(v10))
    fmt.Println(v5)
    var v11 bool = (v10 < 2)
    if len(v2) > 0 {
        v2[0] = "a"
    }
    v5 += int8(v10)
    fmt.Println(v3)
    if true {
        v12 := GenS2{F0: 42, F1: "a"}
        if ((!(v11 && (false && v11))) && (!(("" + ("hello" + "\n\t")) != ("hello" + "hello")))) {
            delete(v4, ("hello" + "a"))
            v11 = (!v11)
            fmt.Println(v1)
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v6 -= (-(v6))
                _ = fi14
            }
            var v15 uint32 = v10
            _ = fi13
            _ = v15
        }
        _ = v12
    }
    v9 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
