package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
type GenS3 struct { F0 uint64; F1 string }
func genFn0(p0 int32, p1 int32, p2 int32) uint8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 string = ("\n\t" + "")
        fmt.Println(v7)
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 201, F1: p0, F2: ("foo bar" + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 int = 761
    var v2 float32 = 0.1
    var v3 int16 = 1000
    v4 := make(map[int]bool)
    var v5 uint64 = 2
    v1 += 125
    if (true || (!false)) {
        v9 := GenS1{F0: 2, F1: 594, F2: "\n\t"}
        v10 := GenS2{F0: "", F1: -1, F2: -1}
        if (!(0 != 53)) {
            v11 := []int{(^(v1))}
            v12 := GenS2{F0: ((((("foo bar" + "\n\t") + "foo bar") + "\n\t") + "foo bar") + "\n\t"), F1: 32767, F2: 0}
            _ = v11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    v13 := GenS1{F0: 1, F1: int32(v3), F2: ("a" + ("foo bar" + "a"))}
    if false {
        fmt.Println(v1)
        switch (v1 % ((4294967295) | 1)) {
        case 127:
            var v14 int16 = (v3 << ((0) & 63))
            _ = v14
        default:
            v15 := GenS2{F0: "hello", F1: v3, F2: 112}
            _ = v15
        }
        fmt.Println(v2)
    } else {
        v16 := GenS2{F0: ("\n\t" + "\n\t"), F1: (v3 * -651), F2: 100}
        _ = v16
    }
    var v17 string = ("a" + (("" + "hello") + ("a" + "")))
    var v18 float32 = v2
    v1 = v1
    if (42 >= -225) {
        v19 := GenS3{F0: 42, F1: v17}
        v20 := GenS3{F0: v5, F1: "a"}
        fmt.Println(v17, v1, v5)
        _ = v19
        _ = v20
    }
    v1 = (v1 / (((-(v1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
