package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint32) bool {
    v6 := []int{32767}
    switch (("foo bar" + "hello") + ("hello" + "a")) {
    case "a":
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(fi7, fi7)
            var v8 int32 = int32(fi7)
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: 127, F1: 2, F2: 1}
    _ = p0
    _ = v6
    _ = v9
    return (!false)
}

func main() {
    var v1 int16 = 1
    v2 := make(map[int]string)
    var v3 int = 127
    v4 := make(map[uint32]string)
    var v5 int8 = 1
    v10 := GenS2{F0: 444, F1: 2147483647}
    v3 += v3
    v4[610] = ""
    var v11 int64 = 1
    v4[683] = "\n\t"
    fmt.Println(v3, v1)
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v3, v3)
        for fi13 := 0; fi13 < 20; fi13++ {
            v11 += v11
            _ = fi13
        }
        _ = fi12
    }
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 bool = (((true && (!false)) && true) || (true || false))
            v16 := []int{}
            v1 -= (-(v1))
            _ = fi14
            _ = v15
            _ = v16
        }
        v11 -= v11
        var v17 float32 = 1.5
        _ = v17
    } else {
        fmt.Println(v3, v5)
        v3 -= (v3 - v3)
        var v18 string = "foo bar"
        _ = v18
    }
    fmt.Println(v11, v5)
    v2[(v3 << ((255) & 63))] = "\n\t"
    v4[693] = ""
    var v19 uint = 1
    delete(v4, 693)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
