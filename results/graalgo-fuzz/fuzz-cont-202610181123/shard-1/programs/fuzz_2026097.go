package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64, p1 uint64) int32 {
    v6 := []int{128, -69}
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return -1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    v3 := []string{("a" + "a")}
    var v4 string = (("foo bar" + (("foo bar" + "foo bar") + "a")) + "")
    v5 := []int{128, 695, 42, 32768}
    v7 := GenS1{F0: 42, F1: 126}
    fmt.Println(v2)
    if true {
        v8 := v1[42]
        var v9 bool = false
        _ = v8
        _ = v9
    } else {
        v2 -= v2
    }
    var v10 uint = 2
    if len(v5) > 0 {
        v5[0] = 128
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint64 = 664
        fmt.Println(v10, fi11)
        _ = fi11
        _ = v12
    }
    v4 = ((v4 + v4) + (("\n\t" + v4) + "\n\t"))
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
