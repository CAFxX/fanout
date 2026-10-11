package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool) int64 {
    if (0 == 96) {
        p1 = (p1 && (float64((-2.25 - 1e10)) <= 0.0))
    }
    var v6 string = "a"
    _ = p0
    _ = p1
    _ = v6
    return 646
}
func genFn1(p0 bool, p1 uint16, p2 uint16) bool {
    v7 := []int{127, 0, 217}
    p1 = p2
    v8 := GenS1{F0: 0, F1: "foo bar", F2: 1}
    p2 = p1
    v9 := GenS1{F0: 65, F1: ("" + "\n\t"), F2: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    var v1 string = "\n\t"
    var v2 uint16 = 2
    var v3 int8 = -128
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: 0, F1: "", F2: -387}
        _ = fi10
        _ = v11
    }
    v2 = (v2 - 0)
    v12 := GenS1{F0: 42, F1: (("hello" + "a") + v1), F2: 9223372036854775807}
    fmt.Println(v2, v3, v2)
    v13 := GenS2{F0: 128}
    var v14 uint = 22
    var v15 uint8 = 128
    for fi16 := 0; fi16 < 1; fi16++ {
        if true {
            var v17 uint8 = (v15 * (v15 / (((v15 ^ (v15 >> ((v14) & 63)))) | 1)))
            fmt.Println(v14, v14)
            v18 := GenS1{F0: (v17 << ((v14) & 63)), F1: ((("" + "foo bar") + "a") + (v1 + v1)), F2: int64(v17)}
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
