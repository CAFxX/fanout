package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
func genFn0(p0 uint64, p1 uint32) int64 {
    v4 := GenS1{F0: "hello"}
    var v5 int32 = 429
    p1 = p1
    p0 -= p0
    fmt.Println(v5, v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -71
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]string)
    var v3 uint16 = 0
    if true {
        switch -875 {
        case 1:
            var v6 uint32 = 42
            _ = v6
        case 128:
            v7 := []int{42}
            _ = v7
        case -604:
            var v8 uint = 964
            _ = v8
        default:
            v3 += v3
        }
    }
    v9 := GenS1{F0: ("hello" + "\n\t")}
    v10 := []int{}
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := GenS1{F0: ((("" + "\n\t") + ("foo bar" + ("hello" + (("foo bar" + "foo bar") + ("hello" + "a"))))) + "\n\t")}
        switch (fi11 & (fi11 + -9223372036854775808)) {
        case -262:
            var v13 float64 = ((1.5 - 0.1) + 0.1)
            _ = v13
        case 759:
            _ = len(v10)
        default:
            fmt.Println(fi11, v3)
        }
        v3 -= (v3 | (^(v3)))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    v14 := GenS1{F0: ""}
    var v15 uint16 = (v3 >> ((1) & 63))
    var v16 uint32 = 736
    if false {
        v15 = 32767
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint16 = v3
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS2{F0: v16, F1: ("hello" + "a"), F2: 32}
    var v20 float64 = 1e10
    if (1e10 < v20) {
        v15 -= 545
        v3 = 871
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
