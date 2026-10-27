package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 int32, p1 int16) int32 {
    v6 := []int{32767}
    var v7 string = "foo bar"
    if (2 <= 238) {
        v7 = v7
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}

func main() {
    v1 := []string{}
    var v2 int64 = -9223372036854775808
    var v3 uint = 316
    var v4 uint8 = 128
    var v5 uint16 = 0
    var v8 uint = uint(v4)
    fmt.Println(v2, v8)
    v9 := GenS1{F0: v4, F1: 32767}
    _ = len(v1)
    var v10 int64 = v2
    if ((^(v5)) != 1) {
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                v10 = v10
                fmt.Println(v3)
                fmt.Println(v10, v8)
                var v13 int16 = -865
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v1 = append(v1, "foo bar")
        v10 = 0
    }
    v14 := GenS2{F0: "hello", F1: "a"}
    v2 -= 1
    v8 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
