package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32) uint16 {
    fmt.Println(p0)
    var v6 float32 = ((((3.14159 / 1e10) * -2.25) + 1.5) * ((0.1 + (float32((-2.25 / 0.1)) * 1.5)) - -2.25))
    switch "hello" {
    case "hello":
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: "hello"}
            _ = fi7
            _ = v8
        }
    case "\n\t":
        fmt.Println(p0, p0)
    default:
        var v9 string = (("hello" + "hello") + ("\n\t" + "\n\t"))
        _ = v9
    }
    _ = p0
    _ = v6
    return 4
}
func genFn1(p0 uint64, p1 bool) string {
    var v10 int = 1
    if true {
        var v11 uint32 = 933
        _ = v11
    }
    var v12 bool = (11 != 205)
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]string)
    var v3 uint = 18446744073709551615
    v4 := []string{"a", "hello", "a"}
    var v5 int32 = 1
    var v13 float64 = 0.0
    v13 += (v13 - 3.14159)
    v14 := GenS2{F0: 42}
    v3 = (v3 * v3)
    v13 -= v13
    var v15 int16 = -1
    var v16 float64 = 0.0
    fmt.Println(v13, v15)
    v1[("" + (((("foo bar" + "a") + "foo bar") + ("\n\t" + ("hello" + "foo bar"))) + (("\n\t" + (("\n\t" + "") + "a")) + "")))] = ("" + (((("a" + "\n\t") + "") + "foo bar") + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
