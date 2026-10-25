package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 6; fi4++ {
        var v5 string = ("\n\t" + "hello")
        _ = fi4
        _ = v5
    }
    var v6 string = "a"
    v7 := []int{249, 139, -1, 2147483647}
    _ = v6
    _ = v7
    return -658
}

func main() {
    var v1 uint = 909
    v2 := make(map[int]int)
    var v3 bool = ((0.1 * (1e10 + 1.5)) == 1e10)
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int32 = 229
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    v10 := GenS1{F0: v1, F1: 331}
    for fi11 := 0; fi11 < 11; fi11++ {
        v3 = v3
        v3 = ("" > "foo bar")
        _ = fi11
    }
    fmt.Println(v3, v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        fmt.Println(v3, v1, fi12)
        _ = fi12
    }
    v2[-1] = -9223372036854775808
    v13 := GenS1{F0: 0, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
}
