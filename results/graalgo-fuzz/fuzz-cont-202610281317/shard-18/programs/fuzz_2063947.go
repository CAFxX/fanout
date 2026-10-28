package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 uint16) int {
    v5 := []int{128}
    fmt.Println(p0)
    fmt.Println(p0)
    var v6 int8 = 127
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 0
    var v3 uint64 = (v1 | 18446744073709551615)
    v4 := make(map[string]int)
    fmt.Println(v3, v1)
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 += (v3 ^ 68)
        _ = fi7
    }
    v1 -= uint64(v2)
    if ((!false) || true) {
        v8 := GenS1{F0: 255, F1: "foo bar", F2: -667}
        _ = v8
    } else {
        v2 -= (v2 / ((v2) | 1))
        v9 := GenS2{F0: "hello", F1: 32767}
        v10 := []int{-849}
        _ = v10
        _ = v9
    }
    fmt.Println(v1)
    _ = len(v4)
    fmt.Println(v3)
    for fi11 := 0; fi11 < 12; fi11++ {
        v3 -= v1
        _ = fi11
    }
    if true {
        if true {
            var v12 int8 = 0
            _ = v12
        }
        if false {
            v13 := []int{32767, genFn0(571), 4294967295, 1}
            _ = v13
        } else {
            v14 := GenS2{F0: ("hello" + ("" + "")), F1: 32767}
            v3 = v1
            _ = v14
        }
    }
    switch 845 {
    case -836:
        v3 += v1
    }
    v2 -= v2
    v2 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
