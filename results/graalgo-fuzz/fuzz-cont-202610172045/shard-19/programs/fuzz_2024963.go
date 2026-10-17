package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 bool }
func genFn0() int16 {
    v4 := []int{9223372036854775807, 4294967295, 1}
    if ("\n\t" == ((("hello" + "foo bar") + ("foo bar" + "\n\t")) + "foo bar")) {
        var v5 uint16 = 385
        _ = v5
    }
    v6 := GenS1{F0: 4294967295, F1: 1, F2: (true || (0 < 0))}
    _ = v4
    _ = v6
    return 32767
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint8 = 0
    var v7 uint16 = 2
    v7 -= (v7 + v7)
    v8 := GenS1{F0: 65535, F1: 121, F2: (false || (1.5 == float64(-2.25)))}
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := []int{fi9}
        for fi11 := 0; fi11 < 5; fi11++ {
            v7 = 2
            _ = fi11
        }
        for fi12 := 0; fi12 < 7; fi12++ {
            v7 -= (v7 >> ((0) & 63))
            _ = fi12
        }
        _ = fi9
        _ = v10
    }
    v13 := []int{32767, 791, -221}
    var v14 float32 = 0.1
    var v15 uint8 = 106
    fmt.Println(v14, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
