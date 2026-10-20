package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }
func genFn0() int {
    var v6 uint8 = 66
    v7 := GenS1{F0: -475}
    var v8 uint8 = 128
    for fi9 := 0; fi9 < 7; fi9++ {
        v8 += (^(v6))
        _ = fi9
    }
    var v10 int8 = -1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    return 128
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{(("" + "") + ("foo bar" + "foo bar")), "a"}
    v3 := make(map[string]string)
    v4 := []int{1}
    v5 := []int{}
    var v11 int16 = 32767
    if (!false) {
        for fi12 := 0; fi12 < 2; fi12++ {
            v11 += v11
            delete(v1, genFn0())
            _ = fi12
        }
        var v13 int32 = -2147483648
        var v14 int = -406
        _ = v13
        _ = v14
    }
    v15 := v3["hello"]
    v15 = "a"
    v11 -= v11
    v11 = v11
    v16 := GenS2{F0: genFn0(), F1: 2, F2: (v11 & 1)}
    v17 := GenS1{F0: 42}
    if true {
        _ = len(v3)
    } else {
        var v18 string = (v15 + ("foo bar" + ((v15 + v15) + ("\n\t" + v15))))
        v19 := v3[v18]
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
