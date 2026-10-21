package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint16 }
func genFn0(p0 uint8) uint16 {
    if false {
        v4 := GenS1{F0: 45, F1: 0, F2: 2}
        _ = v4
    } else {
        switch "foo bar" {
        case "\n\t":
            var v5 bool = (381 <= 42)
            _ = v5
        }
        fmt.Println(p0)
    }
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 -= p0
        _ = fi6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: 42, F1: 32767, F2: 32767}
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    return 32767
}
func genFn1(p0 string, p1 string) uint8 {
    v10 := GenS1{F0: 42, F1: 1, F2: genFn0(115)}
    p0 = "hello"
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    return 2
}

func main() {
    v1 := []string{"hello"}
    var v2 int32 = 42
    v3 := make(map[int]int)
    v11 := GenS1{F0: 1, F1: genFn0(2), F2: 918}
    v12 := GenS1{F0: 18446744073709551615, F1: 0, F2: 226}
    var v13 int16 = -1
    v14 := GenS1{F0: 336, F1: 678, F2: 255}
    var v15 float32 = -2.25
    var v16 int32 = int32(v13)
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 uint64 = 1
        v2 += (v16 & v16)
        _ = fi17
        _ = v18
    }
    var v19 int = 1
    if (("foo bar" + "a") < "a") {
        fmt.Println(v13, v19, v13)
        var v20 bool = false
        _ = v20
    }
    v16 += v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
