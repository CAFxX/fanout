package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint64) bool {
    v6 := GenS1{F0: -32768, F1: ((("hello" + "a") + "a") + (("hello" + "\n\t") + "foo bar"))}
    switch 936 {
    case 0:
        fmt.Println(p0)
    }
    var v7 float64 = 0.0
    p0 += (p0 / ((631) | 1))
    _ = p0
    _ = v6
    _ = v7
    return true
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{9223372036854775807, -57, -9223372036854775808, 32768}
    v3 := make(map[int]int)
    var v4 int8 = 0
    v5 := make(map[uint32]int)
    for fi8 := 0; fi8 < 16; fi8++ {
        switch (fi8 | (-(fi8))) {
        case 0:
            _ = len(v2)
        }
        _ = fi8
    }
    var v9 uint32 = 681
    fmt.Println(v4, v9)
    fmt.Println(v4, v9)
    v4 = v4
    fmt.Println(v4, v4)
    fmt.Println(v4)
    if true {
        fmt.Println(v4, v4)
        if (255 != 140) {
            delete(v3, 2147483647)
        }
    }
    v10 := GenS1{F0: 1000, F1: ((((((("" + "") + "a") + ("\n\t" + "foo bar")) + (("hello" + ("foo bar" + "hello")) + "")) + ("hello" + "")) + "a") + "foo bar")}
    v4 -= 75
    if ((65535 <= 0) || (-591 <= 1)) {
        v11 := GenS2{F0: 845}
        _ = v11
    }
    if true {
        v4 = v4
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
