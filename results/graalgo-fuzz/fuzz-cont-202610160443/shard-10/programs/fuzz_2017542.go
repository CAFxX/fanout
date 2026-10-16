package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int8, p1 bool) int8 {
    var v6 int8 = (p0 | p0)
    var v7 bool = (!p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (p0 >> ((42) & 63))
}

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 0.1
    var v3 int8 = -15
    v4 := make(map[string]bool)
    var v5 uint = 4294967295
    v4[("a" + ("foo bar" + "hello"))] = (147 > 10)
    var v8 int64 = -914
    v2 = float64((((v1 + v1) / v1) + v1))
    v2 -= (v2 - float64((0.1 + float32(v2))))
    switch "\n\t" {
    case "a":
        var v9 uint8 = 42
        _ = v9
    case "hello":
        var v10 float32 = v1
        _ = v10
    case "foo bar":
        var v11 int = 518
        _ = v11
    }
    v12 := GenS1{F0: 65535}
    if false {
        v8 -= v8
        v5 = 166
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 float32 = 1e10
                v15 = 0.0
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    } else {
        v5 -= v5
        v1 = v1
        v16 := GenS1{F0: 255}
        _ = v16
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        if true {
            var v18 bool = ("\n\t" < (("" + ("" + "foo bar")) + ("\n\t" + "\n\t")))
            _ = v18
        }
        delete(v4, ("\n\t" + ("hello" + "a")))
        var v19 float64 = v2
        _ = fi17
        _ = v19
    }
    fmt.Println(v2)
    v2 += v2
    v20 := []int{127, 9223372036854775807}
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
