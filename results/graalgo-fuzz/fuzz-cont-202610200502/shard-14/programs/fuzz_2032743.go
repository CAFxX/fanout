package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64, p1 uint8, p2 int16) uint64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        var v5 string = ("a" + ("a" + "a"))
        var v6 string = (("\n\t" + ((("" + v5) + ("foo bar" + (("hello" + ((v5 + (("" + v5) + v5)) + (("foo bar" + ((v5 + ((("hello" + v5) + v5) + ("foo bar" + (v5 + ((v5 + v5) + (v5 + v5)))))) + (v5 + (v5 + v5)))) + ((("a" + "a") + "foo bar") + v5)))) + "hello"))) + ((((v5 + v5) + v5) + "a") + (v5 + ("hello" + v5))))) + ("foo bar" + "hello"))
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint32 = 42
    p2 = (p2 * (p2 | -835))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 float32 = (float32(-2.25) / 1.0)
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    switch ("foo bar" + "foo bar") {
    case "a":
        if (0 <= 255) {
            v8 := []int{1, 42}
            _ = v8
        } else {
            var v9 int8 = 100
            v9 += v9
            _ = v9
        }
    case "hello":
        v1 += v1
    case "\n\t":
        v1 = v1
    default:
        v10 := v2[32768]
        _ = v10
    }
    fmt.Println(v1)
    v11 := GenS1{F0: 255}
    var v12 int = 999
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 float32 = v1
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    var v15 uint16 = 42
    switch (-(v12)) {
    case 956:
        var v16 float32 = v1
        _ = v16
    default:
        v12 = (v12 * v12)
    }
    if ((42 <= 1) || false) {
        v12 -= v12
    }
    fmt.Println(v15)
    var v17 int16 = 32767
    var v18 float32 = (float32(1.5) + 1e10)
    v3[v12] = 528
    v1 -= v18
    var v19 int64 = -1
    var v20 uint64 = genFn0(0, 42, v17)
    var v21 int16 = v17
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
