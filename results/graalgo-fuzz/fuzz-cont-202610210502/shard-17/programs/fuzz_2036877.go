package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
func genFn0() string {
    switch 9223372036854775807 {
    case 32768:
        v5 := GenS1{F0: -1, F1: 1}
        _ = v5
    }
    v6 := []int{9223372036854775807}
    var v7 int64 = -9223372036854775808
    _ = v6
    _ = v7
    return "foo bar"
}
func genFn1() uint8 {
    var v8 int = -964
    var v9 float64 = float64(float32((3.14159 + (1e10 / (-2.25 - 1.5)))))
    v10 := GenS1{F0: 42, F1: 1}
    _ = v10
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]string)
    v4 := []string{"\n\t", "foo bar", "\n\t", "\n\t"}
    fmt.Println(v2)
    var v11 uint8 = genFn1()
    v11 = v11
    var v12 int = 0
    v1 = append(v1, (("a" + "") + "hello"))
    for fi13 := 0; fi13 < 20; fi13++ {
        v12 = v12
        v2 -= v2
        var v14 int = (fi13 ^ fi13)
        _ = fi13
        _ = v14
    }
    _ = len(v3)
    var v15 uint32 = 4294967295
    var v16 string = ("hello" + (("a" + "\n\t") + "hello"))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        switch v16 {
        case "foo bar":
            var v17 int8 = -104
            _ = v17
        default:
            for fi18 := 0; fi18 < 19; fi18++ {
                v2 += (v2 * v2)
                var v19 bool = true
                _ = fi18
                _ = v19
            }
        }
        v16 = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
