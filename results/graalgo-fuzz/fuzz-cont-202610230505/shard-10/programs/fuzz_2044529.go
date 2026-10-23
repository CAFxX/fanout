package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
func genFn0(p0 uint16, p1 uint8, p2 bool) uint16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        switch (("" + ("hello" + (((("a" + "hello") + "a") + "hello") + "hello"))) + "\n\t") {
        case "\n\t":
            var v6 uint32 = 723
            _ = v6
        case "a":
            p1 = 208
        case "hello":
            v7 := GenS1{F0: -604}
            _ = v7
        }
        _ = fi5
    }
    v8 := []int{}
    _ = len(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 + (^(p0)))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = -132
    v3 := []string{("foo bar" + "foo bar"), "hello"}
    var v4 uint64 = 42
    fmt.Println(v4)
    v4 -= (v4 % (((v4 >> ((2) & 63))) | 1))
    v9 := GenS1{F0: 738}
    var v10 bool = true
    fmt.Println(v2)
    var v11 uint32 = 2147483648
    var v12 int = (-(v2))
    fmt.Println(v10)
    v13 := GenS1{F0: 0}
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    v10 = true
    v12 = v2
    var v14 float32 = (float32((0.0 * -2.25)) + (0.1 / 1.0))
    switch v12 {
    case 255:
        v15 := GenS2{F0: 42}
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
